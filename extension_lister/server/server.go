package main

import (
	"context"
	_ "fmt"
	"log"
	"net"
	"net/url"
	"time"
	"strings"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	_ "github.com/vmware/govmomi/view"
	_ "github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/session/cache"
	"google.golang.org/grpc"
	pb "dummymodule/proto"
)

type server struct {
	pb.UnimplementedExtensionListerServer
	authenticated bool
}

func (s *server) Authenticate(ctx context.Context, authRequest *pb.AuthRequest) (*pb.AuthResponse, error) {
	log.Println("authenticate called")
	if (authRequest.GetUsername() == "admin" &&
		authRequest.GetPassword() == "admin") {
		s.authenticated = true
		return &pb.AuthResponse{Result: true, Reason: "credentials OK"}, nil
	}
	return &pb.AuthResponse{Result: false, Reason: "wrong credentials"}, nil
}

func (s *server) listExtensions(ctx context.Context, keyContains string, client *vim25.Client,
		extensionServer pb.ExtensionLister_ListExtensionsServer) error {
	extensionManager, err := object.GetExtensionManager(client)
	extensions, err := extensionManager.List(ctx)
	if err != nil {
		log.Fatalf("error while fetching extension: %v", err)
	}
	log.Printf("Found %d extensions", len(extensions))
	log.Printf("Type: %T", extensions)
	for _, extension := range extensions {
		if (!strings.Contains(extension.Key, keyContains)) {
			continue
		}
		time.Sleep(100 * time.Millisecond)

		serverInfo := make([]*pb.ExtensionServerInfo, 0)
		for _, sInfo := range extension.Server {
			serverInfo = append(serverInfo, &pb.ExtensionServerInfo{
				Url: sInfo.Url,
				Company: sInfo.Company,
			})
		}

		if err := extensionServer.Send(&pb.Extension{
			Key: extension.Key,
			Company: extension.Company,
			Descriptionlabel: extension.Description.GetDescription().Label,
			Descriptionsummary: extension.Description.GetDescription().Summary,
			Extensionserverinfo: serverInfo,
		}); err != nil {
			return err
		}
		//log.Printf("description label: %v, description summary: %v, key: %v, company: %v", extension.Description.GetDescription().Label, extension.Description.GetDescription().Summary, extension.Key, extension.Company)
	}
	return nil
}

func (s *server) ListExtensions(listRequest *pb.ListExtensionRequest, extensionServer pb.ExtensionLister_ListExtensionsServer) error {
	log.Println("list_extension called")
	if s.authenticated {
		sdk_url, err := soap.ParseURL("https://127.0.0.1/sdk")
		sdk_url.User = url.UserPassword("admin", "admin")

		ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)
		defer cancel()

		var client *vim25.Client
		if err != nil {
			log.Printf("Couldn't parse vCenter url: %v", err)
		} else {
			log.Printf("url: %v", sdk_url)
			session := &cache.Session {
				URL: sdk_url,
				Insecure: true,
				Passthrough: true,
			}
			client = new(vim25.Client)
			//log.Print("Login start")
			err = session.Login(ctx, client, nil)
			//log.Print("Login end")
			if err != nil {
				log.Printf("Couldn't login to vCenter: %v", err)
				client = nil
			}
		}

		if (client != nil) {
			//log.Printf("logged to vCenter. client = %v", client)
			return s.listExtensions(ctx, listRequest.GetKeyContains(), client, extensionServer)
		}

		//return &pb.ExtensionList{Key: "verified_extension"}, nil
		return nil
	} else {
		// return &pb.ExtensionList{}, nil
		return nil
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3030")
	if err != nil {
		log.Fatal("no way. couldn't bound to port")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterExtensionListerServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("no way. grpc server couldn't serve")
	}
}
