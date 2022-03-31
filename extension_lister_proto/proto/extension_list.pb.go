// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/extension_list.proto

package importpath

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_extension_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_extension_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_proto_extension_list_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_extension_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_extension_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_proto_extension_list_proto_rawDescGZIP(), []int{1}
}

func (x *AuthResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *AuthResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type ListExtensionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyContains string `protobuf:"bytes,1,opt,name=keyContains,proto3" json:"keyContains,omitempty"`
}

func (x *ListExtensionRequest) Reset() {
	*x = ListExtensionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_extension_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListExtensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExtensionRequest) ProtoMessage() {}

func (x *ListExtensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_extension_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExtensionRequest.ProtoReflect.Descriptor instead.
func (*ListExtensionRequest) Descriptor() ([]byte, []int) {
	return file_proto_extension_list_proto_rawDescGZIP(), []int{2}
}

func (x *ListExtensionRequest) GetKeyContains() string {
	if x != nil {
		return x.KeyContains
	}
	return ""
}

type ExtensionServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Company string `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
}

func (x *ExtensionServerInfo) Reset() {
	*x = ExtensionServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_extension_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionServerInfo) ProtoMessage() {}

func (x *ExtensionServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_extension_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionServerInfo.ProtoReflect.Descriptor instead.
func (*ExtensionServerInfo) Descriptor() ([]byte, []int) {
	return file_proto_extension_list_proto_rawDescGZIP(), []int{3}
}

func (x *ExtensionServerInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ExtensionServerInfo) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

type Extension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key                 string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Company             string                 `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
	Descriptionlabel    string                 `protobuf:"bytes,3,opt,name=descriptionlabel,proto3" json:"descriptionlabel,omitempty"`
	Descriptionsummary  string                 `protobuf:"bytes,4,opt,name=descriptionsummary,proto3" json:"descriptionsummary,omitempty"`
	Extensionserverinfo []*ExtensionServerInfo `protobuf:"bytes,5,rep,name=extensionserverinfo,proto3" json:"extensionserverinfo,omitempty"`
}

func (x *Extension) Reset() {
	*x = Extension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_extension_list_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extension) ProtoMessage() {}

func (x *Extension) ProtoReflect() protoreflect.Message {
	mi := &file_proto_extension_list_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extension.ProtoReflect.Descriptor instead.
func (*Extension) Descriptor() ([]byte, []int) {
	return file_proto_extension_list_proto_rawDescGZIP(), []int{4}
}

func (x *Extension) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Extension) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *Extension) GetDescriptionlabel() string {
	if x != nil {
		return x.Descriptionlabel
	}
	return ""
}

func (x *Extension) GetDescriptionsummary() string {
	if x != nil {
		return x.Descriptionsummary
	}
	return ""
}

func (x *Extension) GetExtensionserverinfo() []*ExtensionServerInfo {
	if x != nil {
		return x.Extensionserverinfo
	}
	return nil
}

var File_proto_extension_list_proto protoreflect.FileDescriptor

var file_proto_extension_list_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3e, 0x0a, 0x0c, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x41, 0x0a, 0x13, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0xe2, 0x01, 0x0a, 0x09, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2e,
	0x0a, 0x12, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x4d,
	0x0a, 0x13, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x13, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x95, 0x01,
	0x0a, 0x0f, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x13, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x30, 0x01, 0x42, 0x19, 0x5a, 0x17, 0x74, 0x68, 0x69, 0x73, 0x2f, 0x69, 0x73,
	0x2f, 0x73, 0x6f, 0x6d, 0x65, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x61, 0x74, 0x68,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_extension_list_proto_rawDescOnce sync.Once
	file_proto_extension_list_proto_rawDescData = file_proto_extension_list_proto_rawDesc
)

func file_proto_extension_list_proto_rawDescGZIP() []byte {
	file_proto_extension_list_proto_rawDescOnce.Do(func() {
		file_proto_extension_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_extension_list_proto_rawDescData)
	})
	return file_proto_extension_list_proto_rawDescData
}

var file_proto_extension_list_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_extension_list_proto_goTypes = []interface{}{
	(*AuthRequest)(nil),          // 0: lister.AuthRequest
	(*AuthResponse)(nil),         // 1: lister.AuthResponse
	(*ListExtensionRequest)(nil), // 2: lister.ListExtensionRequest
	(*ExtensionServerInfo)(nil),  // 3: lister.ExtensionServerInfo
	(*Extension)(nil),            // 4: lister.Extension
}
var file_proto_extension_list_proto_depIdxs = []int32{
	3, // 0: lister.Extension.extensionserverinfo:type_name -> lister.ExtensionServerInfo
	0, // 1: lister.ExtensionLister.Authenticate:input_type -> lister.AuthRequest
	2, // 2: lister.ExtensionLister.listExtensions:input_type -> lister.ListExtensionRequest
	1, // 3: lister.ExtensionLister.Authenticate:output_type -> lister.AuthResponse
	4, // 4: lister.ExtensionLister.listExtensions:output_type -> lister.Extension
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_extension_list_proto_init() }
func file_proto_extension_list_proto_init() {
	if File_proto_extension_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_extension_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_extension_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_extension_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListExtensionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_extension_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionServerInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_extension_list_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extension); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_extension_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_extension_list_proto_goTypes,
		DependencyIndexes: file_proto_extension_list_proto_depIdxs,
		MessageInfos:      file_proto_extension_list_proto_msgTypes,
	}.Build()
	File_proto_extension_list_proto = out.File
	file_proto_extension_list_proto_rawDesc = nil
	file_proto_extension_list_proto_goTypes = nil
	file_proto_extension_list_proto_depIdxs = nil
}
