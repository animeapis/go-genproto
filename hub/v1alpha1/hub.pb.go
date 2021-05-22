// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.0
// source: animeshon/hub/v1alpha1/hub.proto

package hub

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Repository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Repository) Reset() {
	*x = Repository{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_hub_v1alpha1_hub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repository) ProtoMessage() {}

func (x *Repository) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_hub_v1alpha1_hub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repository.ProtoReflect.Descriptor instead.
func (*Repository) Descriptor() ([]byte, []int) {
	return file_animeshon_hub_v1alpha1_hub_proto_rawDescGZIP(), []int{0}
}

func (x *Repository) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AdvertiseReferencesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *AdvertiseReferencesRequest) Reset() {
	*x = AdvertiseReferencesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_hub_v1alpha1_hub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvertiseReferencesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvertiseReferencesRequest) ProtoMessage() {}

func (x *AdvertiseReferencesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_hub_v1alpha1_hub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvertiseReferencesRequest.ProtoReflect.Descriptor instead.
func (*AdvertiseReferencesRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_hub_v1alpha1_hub_proto_rawDescGZIP(), []int{1}
}

func (x *AdvertiseReferencesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdvertiseReferencesRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type ReceivePackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	HttpBody *httpbody.HttpBody `protobuf:"bytes,2,opt,name=http_body,json=httpBody,proto3" json:"http_body,omitempty"`
}

func (x *ReceivePackRequest) Reset() {
	*x = ReceivePackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_hub_v1alpha1_hub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceivePackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceivePackRequest) ProtoMessage() {}

func (x *ReceivePackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_hub_v1alpha1_hub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceivePackRequest.ProtoReflect.Descriptor instead.
func (*ReceivePackRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_hub_v1alpha1_hub_proto_rawDescGZIP(), []int{2}
}

func (x *ReceivePackRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReceivePackRequest) GetHttpBody() *httpbody.HttpBody {
	if x != nil {
		return x.HttpBody
	}
	return nil
}

type UploadPackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	HttpBody *httpbody.HttpBody `protobuf:"bytes,2,opt,name=http_body,json=httpBody,proto3" json:"http_body,omitempty"`
}

func (x *UploadPackRequest) Reset() {
	*x = UploadPackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_hub_v1alpha1_hub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadPackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadPackRequest) ProtoMessage() {}

func (x *UploadPackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_hub_v1alpha1_hub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadPackRequest.ProtoReflect.Descriptor instead.
func (*UploadPackRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_hub_v1alpha1_hub_proto_rawDescGZIP(), []int{3}
}

func (x *UploadPackRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadPackRequest) GetHttpBody() *httpbody.HttpBody {
	if x != nil {
		return x.HttpBody
	}
	return nil
}

type CreateRepositoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateRepositoryRequest) Reset() {
	*x = CreateRepositoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_hub_v1alpha1_hub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepositoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepositoryRequest) ProtoMessage() {}

func (x *CreateRepositoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_hub_v1alpha1_hub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepositoryRequest.ProtoReflect.Descriptor instead.
func (*CreateRepositoryRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_hub_v1alpha1_hub_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRepositoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteRepositoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteRepositoryRequest) Reset() {
	*x = DeleteRepositoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_hub_v1alpha1_hub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRepositoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRepositoryRequest) ProtoMessage() {}

func (x *DeleteRepositoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_hub_v1alpha1_hub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRepositoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteRepositoryRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_hub_v1alpha1_hub_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRepositoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_animeshon_hub_v1alpha1_hub_proto protoreflect.FileDescriptor

var file_animeshon_hub_v1alpha1_hub_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2f, 0x68, 0x75, 0x62, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x68, 0x75,
	0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0a, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x1a, 0x41,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5b, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x31, 0x0a, 0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70,
	0x42, 0x6f, 0x64, 0x79, 0x22, 0x5a, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a,
	0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79,
	0x22, 0x2d, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x2d, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xba,
	0x03, 0x0a, 0x03, 0x47, 0x69, 0x74, 0x12, 0x87, 0x01, 0x0a, 0x13, 0x41, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x32,
	0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73,
	0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x12, 0x1e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x2a, 0x2f, 0x2a, 0x7d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x72, 0x65, 0x66, 0x73,
	0x12, 0x89, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x12, 0x2a, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x68, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f,
	0x64, 0x79, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x22, 0x25, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x2a, 0x2f, 0x2a, 0x7d,
	0x2f, 0x67, 0x69, 0x74, 0x2d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x2d, 0x70, 0x61, 0x63,
	0x6b, 0x3a, 0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x86, 0x01, 0x0a,
	0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x29, 0x2e, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x37, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x31, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x2a, 0x2f, 0x2a, 0x7d, 0x2f, 0x67, 0x69, 0x74, 0x2d, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2d, 0x70, 0x61, 0x63, 0x6b, 0x3a, 0x09, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x14, 0xca, 0x41, 0x11, 0x68, 0x75, 0x62, 0x2e, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0xb3, 0x02, 0x0a, 0x06,
	0x49, 0x73, 0x65, 0x6b, 0x61, 0x69, 0x12, 0x8f, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2f, 0x2e, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x2a, 0x2f, 0x2a, 0x7d, 0x3a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x80, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2f, 0x2e,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x2a, 0x2f, 0x2a, 0x7d, 0x3a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x14, 0xca, 0x41, 0x11,
	0x68, 0x75, 0x62, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x42, 0x6c, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68,
	0x6f, 0x6e, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x50,
	0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x68, 0x75, 0x62, 0xea, 0x02, 0x18, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e,
	0x3a, 0x3a, 0x48, 0x75, 0x62, 0x3a, 0x3a, 0x76, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_animeshon_hub_v1alpha1_hub_proto_rawDescOnce sync.Once
	file_animeshon_hub_v1alpha1_hub_proto_rawDescData = file_animeshon_hub_v1alpha1_hub_proto_rawDesc
)

func file_animeshon_hub_v1alpha1_hub_proto_rawDescGZIP() []byte {
	file_animeshon_hub_v1alpha1_hub_proto_rawDescOnce.Do(func() {
		file_animeshon_hub_v1alpha1_hub_proto_rawDescData = protoimpl.X.CompressGZIP(file_animeshon_hub_v1alpha1_hub_proto_rawDescData)
	})
	return file_animeshon_hub_v1alpha1_hub_proto_rawDescData
}

var file_animeshon_hub_v1alpha1_hub_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_animeshon_hub_v1alpha1_hub_proto_goTypes = []interface{}{
	(*Repository)(nil),                 // 0: animeshon.hub.v1alpha1.Repository
	(*AdvertiseReferencesRequest)(nil), // 1: animeshon.hub.v1alpha1.AdvertiseReferencesRequest
	(*ReceivePackRequest)(nil),         // 2: animeshon.hub.v1alpha1.ReceivePackRequest
	(*UploadPackRequest)(nil),          // 3: animeshon.hub.v1alpha1.UploadPackRequest
	(*CreateRepositoryRequest)(nil),    // 4: animeshon.hub.v1alpha1.CreateRepositoryRequest
	(*DeleteRepositoryRequest)(nil),    // 5: animeshon.hub.v1alpha1.DeleteRepositoryRequest
	(*httpbody.HttpBody)(nil),          // 6: google.api.HttpBody
	(*emptypb.Empty)(nil),              // 7: google.protobuf.Empty
}
var file_animeshon_hub_v1alpha1_hub_proto_depIdxs = []int32{
	6, // 0: animeshon.hub.v1alpha1.ReceivePackRequest.http_body:type_name -> google.api.HttpBody
	6, // 1: animeshon.hub.v1alpha1.UploadPackRequest.http_body:type_name -> google.api.HttpBody
	1, // 2: animeshon.hub.v1alpha1.Git.AdvertiseReferences:input_type -> animeshon.hub.v1alpha1.AdvertiseReferencesRequest
	2, // 3: animeshon.hub.v1alpha1.Git.ReceivePack:input_type -> animeshon.hub.v1alpha1.ReceivePackRequest
	3, // 4: animeshon.hub.v1alpha1.Git.UploadPack:input_type -> animeshon.hub.v1alpha1.UploadPackRequest
	4, // 5: animeshon.hub.v1alpha1.Isekai.CreateRepository:input_type -> animeshon.hub.v1alpha1.CreateRepositoryRequest
	5, // 6: animeshon.hub.v1alpha1.Isekai.DeleteRepository:input_type -> animeshon.hub.v1alpha1.DeleteRepositoryRequest
	6, // 7: animeshon.hub.v1alpha1.Git.AdvertiseReferences:output_type -> google.api.HttpBody
	6, // 8: animeshon.hub.v1alpha1.Git.ReceivePack:output_type -> google.api.HttpBody
	6, // 9: animeshon.hub.v1alpha1.Git.UploadPack:output_type -> google.api.HttpBody
	0, // 10: animeshon.hub.v1alpha1.Isekai.CreateRepository:output_type -> animeshon.hub.v1alpha1.Repository
	7, // 11: animeshon.hub.v1alpha1.Isekai.DeleteRepository:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_animeshon_hub_v1alpha1_hub_proto_init() }
func file_animeshon_hub_v1alpha1_hub_proto_init() {
	if File_animeshon_hub_v1alpha1_hub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_animeshon_hub_v1alpha1_hub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repository); i {
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
		file_animeshon_hub_v1alpha1_hub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvertiseReferencesRequest); i {
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
		file_animeshon_hub_v1alpha1_hub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceivePackRequest); i {
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
		file_animeshon_hub_v1alpha1_hub_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadPackRequest); i {
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
		file_animeshon_hub_v1alpha1_hub_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepositoryRequest); i {
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
		file_animeshon_hub_v1alpha1_hub_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRepositoryRequest); i {
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
			RawDescriptor: file_animeshon_hub_v1alpha1_hub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_animeshon_hub_v1alpha1_hub_proto_goTypes,
		DependencyIndexes: file_animeshon_hub_v1alpha1_hub_proto_depIdxs,
		MessageInfos:      file_animeshon_hub_v1alpha1_hub_proto_msgTypes,
	}.Build()
	File_animeshon_hub_v1alpha1_hub_proto = out.File
	file_animeshon_hub_v1alpha1_hub_proto_rawDesc = nil
	file_animeshon_hub_v1alpha1_hub_proto_goTypes = nil
	file_animeshon_hub_v1alpha1_hub_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GitClient is the client API for Git service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GitClient interface {
	AdvertiseReferences(ctx context.Context, in *AdvertiseReferencesRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	ReceivePack(ctx context.Context, in *ReceivePackRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	UploadPack(ctx context.Context, in *UploadPackRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type gitClient struct {
	cc grpc.ClientConnInterface
}

func NewGitClient(cc grpc.ClientConnInterface) GitClient {
	return &gitClient{cc}
}

func (c *gitClient) AdvertiseReferences(ctx context.Context, in *AdvertiseReferencesRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/animeshon.hub.v1alpha1.Git/AdvertiseReferences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitClient) ReceivePack(ctx context.Context, in *ReceivePackRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/animeshon.hub.v1alpha1.Git/ReceivePack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitClient) UploadPack(ctx context.Context, in *UploadPackRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/animeshon.hub.v1alpha1.Git/UploadPack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitServer is the server API for Git service.
type GitServer interface {
	AdvertiseReferences(context.Context, *AdvertiseReferencesRequest) (*httpbody.HttpBody, error)
	ReceivePack(context.Context, *ReceivePackRequest) (*httpbody.HttpBody, error)
	UploadPack(context.Context, *UploadPackRequest) (*httpbody.HttpBody, error)
}

// UnimplementedGitServer can be embedded to have forward compatible implementations.
type UnimplementedGitServer struct {
}

func (*UnimplementedGitServer) AdvertiseReferences(context.Context, *AdvertiseReferencesRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdvertiseReferences not implemented")
}
func (*UnimplementedGitServer) ReceivePack(context.Context, *ReceivePackRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceivePack not implemented")
}
func (*UnimplementedGitServer) UploadPack(context.Context, *UploadPackRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPack not implemented")
}

func RegisterGitServer(s *grpc.Server, srv GitServer) {
	s.RegisterService(&_Git_serviceDesc, srv)
}

func _Git_AdvertiseReferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvertiseReferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitServer).AdvertiseReferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.hub.v1alpha1.Git/AdvertiseReferences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitServer).AdvertiseReferences(ctx, req.(*AdvertiseReferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Git_ReceivePack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceivePackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitServer).ReceivePack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.hub.v1alpha1.Git/ReceivePack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitServer).ReceivePack(ctx, req.(*ReceivePackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Git_UploadPack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadPackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitServer).UploadPack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.hub.v1alpha1.Git/UploadPack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitServer).UploadPack(ctx, req.(*UploadPackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Git_serviceDesc = grpc.ServiceDesc{
	ServiceName: "animeshon.hub.v1alpha1.Git",
	HandlerType: (*GitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdvertiseReferences",
			Handler:    _Git_AdvertiseReferences_Handler,
		},
		{
			MethodName: "ReceivePack",
			Handler:    _Git_ReceivePack_Handler,
		},
		{
			MethodName: "UploadPack",
			Handler:    _Git_UploadPack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "animeshon/hub/v1alpha1/hub.proto",
}

// IsekaiClient is the client API for Isekai service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IsekaiClient interface {
	CreateRepository(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*Repository, error)
	DeleteRepository(ctx context.Context, in *DeleteRepositoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type isekaiClient struct {
	cc grpc.ClientConnInterface
}

func NewIsekaiClient(cc grpc.ClientConnInterface) IsekaiClient {
	return &isekaiClient{cc}
}

func (c *isekaiClient) CreateRepository(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*Repository, error) {
	out := new(Repository)
	err := c.cc.Invoke(ctx, "/animeshon.hub.v1alpha1.Isekai/CreateRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *isekaiClient) DeleteRepository(ctx context.Context, in *DeleteRepositoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/animeshon.hub.v1alpha1.Isekai/DeleteRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IsekaiServer is the server API for Isekai service.
type IsekaiServer interface {
	CreateRepository(context.Context, *CreateRepositoryRequest) (*Repository, error)
	DeleteRepository(context.Context, *DeleteRepositoryRequest) (*emptypb.Empty, error)
}

// UnimplementedIsekaiServer can be embedded to have forward compatible implementations.
type UnimplementedIsekaiServer struct {
}

func (*UnimplementedIsekaiServer) CreateRepository(context.Context, *CreateRepositoryRequest) (*Repository, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepository not implemented")
}
func (*UnimplementedIsekaiServer) DeleteRepository(context.Context, *DeleteRepositoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepository not implemented")
}

func RegisterIsekaiServer(s *grpc.Server, srv IsekaiServer) {
	s.RegisterService(&_Isekai_serviceDesc, srv)
}

func _Isekai_CreateRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IsekaiServer).CreateRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.hub.v1alpha1.Isekai/CreateRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IsekaiServer).CreateRepository(ctx, req.(*CreateRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Isekai_DeleteRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IsekaiServer).DeleteRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.hub.v1alpha1.Isekai/DeleteRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IsekaiServer).DeleteRepository(ctx, req.(*DeleteRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Isekai_serviceDesc = grpc.ServiceDesc{
	ServiceName: "animeshon.hub.v1alpha1.Isekai",
	HandlerType: (*IsekaiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRepository",
			Handler:    _Isekai_CreateRepository_Handler,
		},
		{
			MethodName: "DeleteRepository",
			Handler:    _Isekai_DeleteRepository_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "animeshon/hub/v1alpha1/hub.proto",
}
