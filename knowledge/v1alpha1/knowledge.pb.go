// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.0
// source: animeshon/knowledge/v1alpha1/knowledge.proto

package knowledge

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Contribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the universe.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Contribution) Reset() {
	*x = Contribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contribution) ProtoMessage() {}

func (x *Contribution) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contribution.ProtoReflect.Descriptor instead.
func (*Contribution) Descriptor() ([]byte, []int) {
	return file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescGZIP(), []int{0}
}

func (x *Contribution) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateContributionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resurce payload.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *CreateContributionRequest) Reset() {
	*x = CreateContributionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContributionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContributionRequest) ProtoMessage() {}

func (x *CreateContributionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContributionRequest.ProtoReflect.Descriptor instead.
func (*CreateContributionRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContributionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateContributionRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type ListContributionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of users to return. Server may return fewer users
	// than requested. If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The value returned from the previous call.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter to be applied to results.
	Filter string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListContributionsRequest) Reset() {
	*x = ListContributionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListContributionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContributionsRequest) ProtoMessage() {}

func (x *ListContributionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContributionsRequest.ProtoReflect.Descriptor instead.
func (*ListContributionsRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescGZIP(), []int{2}
}

func (x *ListContributionsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListContributionsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListContributionsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListContributionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of contributions.
	Contributions []*Contribution `protobuf:"bytes,1,rep,name=contributions,proto3" json:"contributions,omitempty"`
	// A token to retrieve next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListContributionsResponse) Reset() {
	*x = ListContributionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListContributionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContributionsResponse) ProtoMessage() {}

func (x *ListContributionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContributionsResponse.ProtoReflect.Descriptor instead.
func (*ListContributionsResponse) Descriptor() ([]byte, []int) {
	return file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescGZIP(), []int{3}
}

func (x *ListContributionsResponse) GetContributions() []*Contribution {
	if x != nil {
		return x.Contributions
	}
	return nil
}

func (x *ListContributionsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetContributionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the requested contribution.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetContributionRequest) Reset() {
	*x = GetContributionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContributionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContributionRequest) ProtoMessage() {}

func (x *GetContributionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContributionRequest.ProtoReflect.Descriptor instead.
func (*GetContributionRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescGZIP(), []int{4}
}

func (x *GetContributionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ApproveContributionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ApproveContributionRequest) Reset() {
	*x = ApproveContributionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveContributionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveContributionRequest) ProtoMessage() {}

func (x *ApproveContributionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveContributionRequest.ProtoReflect.Descriptor instead.
func (*ApproveContributionRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescGZIP(), []int{5}
}

func (x *ApproveContributionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RejectContributionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RejectContributionRequest) Reset() {
	*x = RejectContributionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectContributionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectContributionRequest) ProtoMessage() {}

func (x *RejectContributionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectContributionRequest.ProtoReflect.Descriptor instead.
func (*RejectContributionRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescGZIP(), []int{6}
}

func (x *RejectContributionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_animeshon_knowledge_v1alpha1_knowledge_proto protoreflect.FileDescriptor

var file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2f, 0x6b, 0x6e, 0x6f, 0x77,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x0c,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x3a, 0x47, 0xea, 0x41, 0x44, 0x0a, 0x24, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x22, 0x49, 0x0a, 0x19, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x6e, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x22, 0x95, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5a, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x6b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0xe0,
	0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x1a, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x19, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xca, 0x06, 0x0a, 0x09,
	0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x37, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xa5, 0x01, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x36, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73,
	0x68, 0x6f, 0x6e, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x9d, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f,
	0x6e, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12,
	0x20, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a,
	0x7d, 0x12, 0x9c, 0x01, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x2e, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x33, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2d, 0x22, 0x28, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x99, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73,
	0x68, 0x6f, 0x6e, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c,
	0x22, 0x27, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x2a, 0x7d, 0x3a, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x1a, 0x1a, 0xca, 0x41,
	0x17, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x84, 0x01, 0x0a, 0x20, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x50, 0x01, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x3b, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0xea, 0x02,
	0x1e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x3a, 0x3a, 0x4b, 0x6e, 0x6f, 0x77,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x3a, 0x3a, 0x76, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescOnce sync.Once
	file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescData = file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDesc
)

func file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescGZIP() []byte {
	file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescOnce.Do(func() {
		file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescData = protoimpl.X.CompressGZIP(file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescData)
	})
	return file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDescData
}

var file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_animeshon_knowledge_v1alpha1_knowledge_proto_goTypes = []interface{}{
	(*Contribution)(nil),               // 0: animeshon.knowledge.v1alpha1.Contribution
	(*CreateContributionRequest)(nil),  // 1: animeshon.knowledge.v1alpha1.CreateContributionRequest
	(*ListContributionsRequest)(nil),   // 2: animeshon.knowledge.v1alpha1.ListContributionsRequest
	(*ListContributionsResponse)(nil),  // 3: animeshon.knowledge.v1alpha1.ListContributionsResponse
	(*GetContributionRequest)(nil),     // 4: animeshon.knowledge.v1alpha1.GetContributionRequest
	(*ApproveContributionRequest)(nil), // 5: animeshon.knowledge.v1alpha1.ApproveContributionRequest
	(*RejectContributionRequest)(nil),  // 6: animeshon.knowledge.v1alpha1.RejectContributionRequest
	(*emptypb.Empty)(nil),              // 7: google.protobuf.Empty
}
var file_animeshon_knowledge_v1alpha1_knowledge_proto_depIdxs = []int32{
	0, // 0: animeshon.knowledge.v1alpha1.ListContributionsResponse.contributions:type_name -> animeshon.knowledge.v1alpha1.Contribution
	1, // 1: animeshon.knowledge.v1alpha1.Knowledge.CreateContribution:input_type -> animeshon.knowledge.v1alpha1.CreateContributionRequest
	2, // 2: animeshon.knowledge.v1alpha1.Knowledge.ListContributions:input_type -> animeshon.knowledge.v1alpha1.ListContributionsRequest
	4, // 3: animeshon.knowledge.v1alpha1.Knowledge.GetContribution:input_type -> animeshon.knowledge.v1alpha1.GetContributionRequest
	5, // 4: animeshon.knowledge.v1alpha1.Knowledge.ApproveContribution:input_type -> animeshon.knowledge.v1alpha1.ApproveContributionRequest
	6, // 5: animeshon.knowledge.v1alpha1.Knowledge.RejectContribution:input_type -> animeshon.knowledge.v1alpha1.RejectContributionRequest
	0, // 6: animeshon.knowledge.v1alpha1.Knowledge.CreateContribution:output_type -> animeshon.knowledge.v1alpha1.Contribution
	3, // 7: animeshon.knowledge.v1alpha1.Knowledge.ListContributions:output_type -> animeshon.knowledge.v1alpha1.ListContributionsResponse
	0, // 8: animeshon.knowledge.v1alpha1.Knowledge.GetContribution:output_type -> animeshon.knowledge.v1alpha1.Contribution
	7, // 9: animeshon.knowledge.v1alpha1.Knowledge.ApproveContribution:output_type -> google.protobuf.Empty
	7, // 10: animeshon.knowledge.v1alpha1.Knowledge.RejectContribution:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_animeshon_knowledge_v1alpha1_knowledge_proto_init() }
func file_animeshon_knowledge_v1alpha1_knowledge_proto_init() {
	if File_animeshon_knowledge_v1alpha1_knowledge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contribution); i {
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
		file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContributionRequest); i {
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
		file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListContributionsRequest); i {
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
		file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListContributionsResponse); i {
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
		file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContributionRequest); i {
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
		file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveContributionRequest); i {
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
		file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejectContributionRequest); i {
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
			RawDescriptor: file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_animeshon_knowledge_v1alpha1_knowledge_proto_goTypes,
		DependencyIndexes: file_animeshon_knowledge_v1alpha1_knowledge_proto_depIdxs,
		MessageInfos:      file_animeshon_knowledge_v1alpha1_knowledge_proto_msgTypes,
	}.Build()
	File_animeshon_knowledge_v1alpha1_knowledge_proto = out.File
	file_animeshon_knowledge_v1alpha1_knowledge_proto_rawDesc = nil
	file_animeshon_knowledge_v1alpha1_knowledge_proto_goTypes = nil
	file_animeshon_knowledge_v1alpha1_knowledge_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KnowledgeClient is the client API for Knowledge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KnowledgeClient interface {
	CreateContribution(ctx context.Context, in *CreateContributionRequest, opts ...grpc.CallOption) (*Contribution, error)
	ListContributions(ctx context.Context, in *ListContributionsRequest, opts ...grpc.CallOption) (*ListContributionsResponse, error)
	GetContribution(ctx context.Context, in *GetContributionRequest, opts ...grpc.CallOption) (*Contribution, error)
	ApproveContribution(ctx context.Context, in *ApproveContributionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RejectContribution(ctx context.Context, in *RejectContributionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type knowledgeClient struct {
	cc grpc.ClientConnInterface
}

func NewKnowledgeClient(cc grpc.ClientConnInterface) KnowledgeClient {
	return &knowledgeClient{cc}
}

func (c *knowledgeClient) CreateContribution(ctx context.Context, in *CreateContributionRequest, opts ...grpc.CallOption) (*Contribution, error) {
	out := new(Contribution)
	err := c.cc.Invoke(ctx, "/animeshon.knowledge.v1alpha1.Knowledge/CreateContribution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledgeClient) ListContributions(ctx context.Context, in *ListContributionsRequest, opts ...grpc.CallOption) (*ListContributionsResponse, error) {
	out := new(ListContributionsResponse)
	err := c.cc.Invoke(ctx, "/animeshon.knowledge.v1alpha1.Knowledge/ListContributions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledgeClient) GetContribution(ctx context.Context, in *GetContributionRequest, opts ...grpc.CallOption) (*Contribution, error) {
	out := new(Contribution)
	err := c.cc.Invoke(ctx, "/animeshon.knowledge.v1alpha1.Knowledge/GetContribution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledgeClient) ApproveContribution(ctx context.Context, in *ApproveContributionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/animeshon.knowledge.v1alpha1.Knowledge/ApproveContribution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledgeClient) RejectContribution(ctx context.Context, in *RejectContributionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/animeshon.knowledge.v1alpha1.Knowledge/RejectContribution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KnowledgeServer is the server API for Knowledge service.
type KnowledgeServer interface {
	CreateContribution(context.Context, *CreateContributionRequest) (*Contribution, error)
	ListContributions(context.Context, *ListContributionsRequest) (*ListContributionsResponse, error)
	GetContribution(context.Context, *GetContributionRequest) (*Contribution, error)
	ApproveContribution(context.Context, *ApproveContributionRequest) (*emptypb.Empty, error)
	RejectContribution(context.Context, *RejectContributionRequest) (*emptypb.Empty, error)
}

// UnimplementedKnowledgeServer can be embedded to have forward compatible implementations.
type UnimplementedKnowledgeServer struct {
}

func (*UnimplementedKnowledgeServer) CreateContribution(context.Context, *CreateContributionRequest) (*Contribution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContribution not implemented")
}
func (*UnimplementedKnowledgeServer) ListContributions(context.Context, *ListContributionsRequest) (*ListContributionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContributions not implemented")
}
func (*UnimplementedKnowledgeServer) GetContribution(context.Context, *GetContributionRequest) (*Contribution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContribution not implemented")
}
func (*UnimplementedKnowledgeServer) ApproveContribution(context.Context, *ApproveContributionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveContribution not implemented")
}
func (*UnimplementedKnowledgeServer) RejectContribution(context.Context, *RejectContributionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectContribution not implemented")
}

func RegisterKnowledgeServer(s *grpc.Server, srv KnowledgeServer) {
	s.RegisterService(&_Knowledge_serviceDesc, srv)
}

func _Knowledge_CreateContribution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContributionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeServer).CreateContribution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.knowledge.v1alpha1.Knowledge/CreateContribution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeServer).CreateContribution(ctx, req.(*CreateContributionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Knowledge_ListContributions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContributionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeServer).ListContributions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.knowledge.v1alpha1.Knowledge/ListContributions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeServer).ListContributions(ctx, req.(*ListContributionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Knowledge_GetContribution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContributionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeServer).GetContribution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.knowledge.v1alpha1.Knowledge/GetContribution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeServer).GetContribution(ctx, req.(*GetContributionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Knowledge_ApproveContribution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveContributionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeServer).ApproveContribution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.knowledge.v1alpha1.Knowledge/ApproveContribution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeServer).ApproveContribution(ctx, req.(*ApproveContributionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Knowledge_RejectContribution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectContributionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeServer).RejectContribution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.knowledge.v1alpha1.Knowledge/RejectContribution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeServer).RejectContribution(ctx, req.(*RejectContributionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Knowledge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "animeshon.knowledge.v1alpha1.Knowledge",
	HandlerType: (*KnowledgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContribution",
			Handler:    _Knowledge_CreateContribution_Handler,
		},
		{
			MethodName: "ListContributions",
			Handler:    _Knowledge_ListContributions_Handler,
		},
		{
			MethodName: "GetContribution",
			Handler:    _Knowledge_GetContribution_Handler,
		},
		{
			MethodName: "ApproveContribution",
			Handler:    _Knowledge_ApproveContribution_Handler,
		},
		{
			MethodName: "RejectContribution",
			Handler:    _Knowledge_RejectContribution_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "animeshon/knowledge/v1alpha1/knowledge.proto",
}
