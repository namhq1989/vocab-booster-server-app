// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: userpb/hub.proto

package userpb

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

type FindUserByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindUserByIDRequest) Reset() {
	*x = FindUserByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_hub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserByIDRequest) ProtoMessage() {}

func (x *FindUserByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_hub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserByIDRequest.ProtoReflect.Descriptor instead.
func (*FindUserByIDRequest) Descriptor() ([]byte, []int) {
	return file_userpb_hub_proto_rawDescGZIP(), []int{0}
}

func (x *FindUserByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindUserByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *FindUserByIDResponse) Reset() {
	*x = FindUserByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_hub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserByIDResponse) ProtoMessage() {}

func (x *FindUserByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_hub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserByIDResponse.ProtoReflect.Descriptor instead.
func (*FindUserByIDResponse) Descriptor() ([]byte, []int) {
	return file_userpb_hub_proto_rawDescGZIP(), []int{1}
}

func (x *FindUserByIDResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type FindUserByEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *FindUserByEmailRequest) Reset() {
	*x = FindUserByEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_hub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserByEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserByEmailRequest) ProtoMessage() {}

func (x *FindUserByEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_hub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserByEmailRequest.ProtoReflect.Descriptor instead.
func (*FindUserByEmailRequest) Descriptor() ([]byte, []int) {
	return file_userpb_hub_proto_rawDescGZIP(), []int{2}
}

func (x *FindUserByEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type FindUserByEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *FindUserByEmailResponse) Reset() {
	*x = FindUserByEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_hub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserByEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserByEmailResponse) ProtoMessage() {}

func (x *FindUserByEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_hub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserByEmailResponse.ProtoReflect.Descriptor instead.
func (*FindUserByEmailResponse) Descriptor() ([]byte, []int) {
	return file_userpb_hub_proto_rawDescGZIP(), []int{3}
}

func (x *FindUserByEmailResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email          string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Timezone       string `protobuf:"bytes,3,opt,name=timezone,proto3" json:"timezone,omitempty"`
	ProviderSource string `protobuf:"bytes,4,opt,name=providerSource,proto3" json:"providerSource,omitempty"`
	ProviderUid    string `protobuf:"bytes,5,opt,name=providerUid,proto3" json:"providerUid,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_hub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_hub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_userpb_hub_proto_rawDescGZIP(), []int{4}
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *CreateUserRequest) GetProviderSource() string {
	if x != nil {
		return x.ProviderSource
	}
	return ""
}

func (x *CreateUserRequest) GetProviderUid() string {
	if x != nil {
		return x.ProviderUid
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_hub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_hub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_userpb_hub_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_userpb_hub_proto protoreflect.FileDescriptor

var file_userpb_hub_proto_rawDesc = []byte{
	0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x1a, 0x14, 0x75, 0x73, 0x65, 0x72,
	0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x25, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x2e, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x3b, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xa3,
	0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x55, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x55, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xf7, 0x01, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x46, 0x69,
	0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x7f, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x70, 0x62, 0x42, 0x08, 0x48, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f,
	0x76, 0x6f, 0x63, 0x61, 0x62, 0x2d, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0xa2,
	0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x55, 0x73, 0x65, 0x72, 0x70, 0x62, 0xca, 0x02,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x70, 0x62, 0xe2, 0x02, 0x12, 0x55, 0x73, 0x65, 0x72, 0x70, 0x62,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userpb_hub_proto_rawDescOnce sync.Once
	file_userpb_hub_proto_rawDescData = file_userpb_hub_proto_rawDesc
)

func file_userpb_hub_proto_rawDescGZIP() []byte {
	file_userpb_hub_proto_rawDescOnce.Do(func() {
		file_userpb_hub_proto_rawDescData = protoimpl.X.CompressGZIP(file_userpb_hub_proto_rawDescData)
	})
	return file_userpb_hub_proto_rawDescData
}

var file_userpb_hub_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_userpb_hub_proto_goTypes = []interface{}{
	(*FindUserByIDRequest)(nil),     // 0: userpb.FindUserByIDRequest
	(*FindUserByIDResponse)(nil),    // 1: userpb.FindUserByIDResponse
	(*FindUserByEmailRequest)(nil),  // 2: userpb.FindUserByEmailRequest
	(*FindUserByEmailResponse)(nil), // 3: userpb.FindUserByEmailResponse
	(*CreateUserRequest)(nil),       // 4: userpb.CreateUserRequest
	(*CreateUserResponse)(nil),      // 5: userpb.CreateUserResponse
	(*User)(nil),                    // 6: userpb.User
}
var file_userpb_hub_proto_depIdxs = []int32{
	6, // 0: userpb.FindUserByIDResponse.user:type_name -> userpb.User
	6, // 1: userpb.FindUserByEmailResponse.user:type_name -> userpb.User
	0, // 2: userpb.UserService.FindUserByID:input_type -> userpb.FindUserByIDRequest
	2, // 3: userpb.UserService.FindUserByEmail:input_type -> userpb.FindUserByEmailRequest
	4, // 4: userpb.UserService.CreateUser:input_type -> userpb.CreateUserRequest
	1, // 5: userpb.UserService.FindUserByID:output_type -> userpb.FindUserByIDResponse
	3, // 6: userpb.UserService.FindUserByEmail:output_type -> userpb.FindUserByEmailResponse
	5, // 7: userpb.UserService.CreateUser:output_type -> userpb.CreateUserResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_userpb_hub_proto_init() }
func file_userpb_hub_proto_init() {
	if File_userpb_hub_proto != nil {
		return
	}
	file_userpb_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_userpb_hub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserByIDRequest); i {
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
		file_userpb_hub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserByIDResponse); i {
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
		file_userpb_hub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserByEmailRequest); i {
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
		file_userpb_hub_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserByEmailResponse); i {
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
		file_userpb_hub_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_userpb_hub_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
			RawDescriptor: file_userpb_hub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userpb_hub_proto_goTypes,
		DependencyIndexes: file_userpb_hub_proto_depIdxs,
		MessageInfos:      file_userpb_hub_proto_msgTypes,
	}.Build()
	File_userpb_hub_proto = out.File
	file_userpb_hub_proto_rawDesc = nil
	file_userpb_hub_proto_goTypes = nil
	file_userpb_hub_proto_depIdxs = nil
}
