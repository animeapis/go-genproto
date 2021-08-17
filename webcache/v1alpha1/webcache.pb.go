// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: animeshon/webcache/v1alpha1/webcache.proto

package webcache

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Cache contains meta information about a specific web resource.
type Cache struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the cache, idempotently generated from the scheme and uri.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The original scheme indicating the protocol used for the original request.
	Scheme string `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	// The request uri stripped of the original scheme.
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// The response content type indicating the original media type.
	MimeType string `protobuf:"bytes,4,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// The response code indicating the status of the remote response.
	StatusCode int32 `protobuf:"varint,5,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// The absolute redirect uri indicating any permanent or temporary redirect.
	RedirectUri string `protobuf:"bytes,6,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	// The full resource name of the cached resource.
	Resource string `protobuf:"bytes,7,opt,name=resource,proto3" json:"resource,omitempty"`
	// The randomly generated revision identifier of this cache.
	// The format is an 8-character hexadecimal string.
	RevisionId string `protobuf:"bytes,8,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// The creation time indicating when this revision was created.
	RevisionCreateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=revision_create_time,json=revisionCreateTime,proto3" json:"revision_create_time,omitempty"`
	// The expiration time indicating when this revision should no longer be
	// considered valid.
	RevisionExpireTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=revision_expire_time,json=revisionExpireTime,proto3" json:"revision_expire_time,omitempty"`
}

func (x *Cache) Reset() {
	*x = Cache{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cache) ProtoMessage() {}

func (x *Cache) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cache.ProtoReflect.Descriptor instead.
func (*Cache) Descriptor() ([]byte, []int) {
	return file_animeshon_webcache_v1alpha1_webcache_proto_rawDescGZIP(), []int{0}
}

func (x *Cache) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cache) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *Cache) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Cache) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *Cache) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Cache) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *Cache) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Cache) GetRevisionId() string {
	if x != nil {
		return x.RevisionId
	}
	return ""
}

func (x *Cache) GetRevisionCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RevisionCreateTime
	}
	return nil
}

func (x *Cache) GetRevisionExpireTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RevisionExpireTime
	}
	return nil
}

type CreateCacheRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The cache to be created.
	Cache *Cache `protobuf:"bytes,1,opt,name=cache,proto3" json:"cache,omitempty"`
	// The time-to-live indicating how long this cache should be considered valid.
	// If set to zero, the cache will not have an expiration time.
	Ttl *durationpb.Duration `protobuf:"bytes,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *CreateCacheRequest) Reset() {
	*x = CreateCacheRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCacheRequest) ProtoMessage() {}

func (x *CreateCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCacheRequest.ProtoReflect.Descriptor instead.
func (*CreateCacheRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_webcache_v1alpha1_webcache_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCacheRequest) GetCache() *Cache {
	if x != nil {
		return x.Cache
	}
	return nil
}

func (x *CreateCacheRequest) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

type ListCachesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The value returned from the previous call.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter to be applied to results.
	//
	// Currently accepted filters include:
	// - uri:{absolute uri}
	// - resource:{full resource name}
	Filter string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	// Whether to return only the latest revision for each cache.
	OnlyLatestRevision bool `protobuf:"varint,4,opt,name=only_latest_revision,json=onlyLatestRevision,proto3" json:"only_latest_revision,omitempty"`
}

func (x *ListCachesRequest) Reset() {
	*x = ListCachesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCachesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCachesRequest) ProtoMessage() {}

func (x *ListCachesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCachesRequest.ProtoReflect.Descriptor instead.
func (*ListCachesRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_webcache_v1alpha1_webcache_proto_rawDescGZIP(), []int{2}
}

func (x *ListCachesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCachesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListCachesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListCachesRequest) GetOnlyLatestRevision() bool {
	if x != nil {
		return x.OnlyLatestRevision
	}
	return false
}

type ListCachesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of caches.
	Caches []*Cache `protobuf:"bytes,1,rep,name=caches,proto3" json:"caches,omitempty"`
	// A token to retrieve next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListCachesResponse) Reset() {
	*x = ListCachesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCachesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCachesResponse) ProtoMessage() {}

func (x *ListCachesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCachesResponse.ProtoReflect.Descriptor instead.
func (*ListCachesResponse) Descriptor() ([]byte, []int) {
	return file_animeshon_webcache_v1alpha1_webcache_proto_rawDescGZIP(), []int{3}
}

func (x *ListCachesResponse) GetCaches() []*Cache {
	if x != nil {
		return x.Caches
	}
	return nil
}

func (x *ListCachesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetCacheRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the requested cache.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetCacheRequest) Reset() {
	*x = GetCacheRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCacheRequest) ProtoMessage() {}

func (x *GetCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCacheRequest.ProtoReflect.Descriptor instead.
func (*GetCacheRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_webcache_v1alpha1_webcache_proto_rawDescGZIP(), []int{4}
}

func (x *GetCacheRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteCacheRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the cache to delete.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteCacheRequest) Reset() {
	*x = DeleteCacheRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCacheRequest) ProtoMessage() {}

func (x *DeleteCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCacheRequest.ProtoReflect.Descriptor instead.
func (*DeleteCacheRequest) Descriptor() ([]byte, []int) {
	return file_animeshon_webcache_v1alpha1_webcache_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCacheRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_animeshon_webcache_v1alpha1_webcache_proto protoreflect.FileDescriptor

var file_animeshon_webcache_v1alpha1_webcache_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2f, 0x77, 0x65, 0x62, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x77, 0x65,
	0x62, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x96, 0x03, 0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x05, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x4c,
	0x0a, 0x14, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x14,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3d, 0x0a, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x77, 0x65, 0x62,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x12, 0x30, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x04, 0x52, 0x03, 0x74,
	0x74, 0x6c, 0x22, 0x99, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x14,
	0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x6f, 0x6e, 0x6c, 0x79,
	0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x78,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e,
	0x2e, 0x77, 0x65, 0x62, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x06, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0xaf, 0x04, 0x0a, 0x08, 0x57, 0x65, 0x62, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x12, 0x7f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12,
	0x2f, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x77, 0x65, 0x62,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x3a, 0x01,
	0x2a, 0x12, 0x87, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x73,
	0x12, 0x2e, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x77, 0x65, 0x62,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x77, 0x65, 0x62,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x12, 0x7f, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x2c, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73,
	0x68, 0x6f, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f,
	0x6e, 0x2e, 0x77, 0x65, 0x62, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x7c, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x2f, 0x2e, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a, 0x19, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x01, 0x2a, 0x1a, 0x19, 0xca, 0x41, 0x16, 0x77,
	0x65, 0x62, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x80, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65,
	0x62, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b,
	0x77, 0x65, 0x62, 0x63, 0x61, 0x63, 0x68, 0x65, 0xea, 0x02, 0x1d, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x73, 0x68, 0x6f, 0x6e, 0x3a, 0x3a, 0x57, 0x65, 0x62, 0x43, 0x61, 0x63, 0x68, 0x65, 0x3a, 0x3a,
	0x76, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_animeshon_webcache_v1alpha1_webcache_proto_rawDescOnce sync.Once
	file_animeshon_webcache_v1alpha1_webcache_proto_rawDescData = file_animeshon_webcache_v1alpha1_webcache_proto_rawDesc
)

func file_animeshon_webcache_v1alpha1_webcache_proto_rawDescGZIP() []byte {
	file_animeshon_webcache_v1alpha1_webcache_proto_rawDescOnce.Do(func() {
		file_animeshon_webcache_v1alpha1_webcache_proto_rawDescData = protoimpl.X.CompressGZIP(file_animeshon_webcache_v1alpha1_webcache_proto_rawDescData)
	})
	return file_animeshon_webcache_v1alpha1_webcache_proto_rawDescData
}

var file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_animeshon_webcache_v1alpha1_webcache_proto_goTypes = []interface{}{
	(*Cache)(nil),                 // 0: animeshon.webcache.v1alpha1.Cache
	(*CreateCacheRequest)(nil),    // 1: animeshon.webcache.v1alpha1.CreateCacheRequest
	(*ListCachesRequest)(nil),     // 2: animeshon.webcache.v1alpha1.ListCachesRequest
	(*ListCachesResponse)(nil),    // 3: animeshon.webcache.v1alpha1.ListCachesResponse
	(*GetCacheRequest)(nil),       // 4: animeshon.webcache.v1alpha1.GetCacheRequest
	(*DeleteCacheRequest)(nil),    // 5: animeshon.webcache.v1alpha1.DeleteCacheRequest
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 7: google.protobuf.Duration
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_animeshon_webcache_v1alpha1_webcache_proto_depIdxs = []int32{
	6, // 0: animeshon.webcache.v1alpha1.Cache.revision_create_time:type_name -> google.protobuf.Timestamp
	6, // 1: animeshon.webcache.v1alpha1.Cache.revision_expire_time:type_name -> google.protobuf.Timestamp
	0, // 2: animeshon.webcache.v1alpha1.CreateCacheRequest.cache:type_name -> animeshon.webcache.v1alpha1.Cache
	7, // 3: animeshon.webcache.v1alpha1.CreateCacheRequest.ttl:type_name -> google.protobuf.Duration
	0, // 4: animeshon.webcache.v1alpha1.ListCachesResponse.caches:type_name -> animeshon.webcache.v1alpha1.Cache
	1, // 5: animeshon.webcache.v1alpha1.WebCache.CreateCache:input_type -> animeshon.webcache.v1alpha1.CreateCacheRequest
	2, // 6: animeshon.webcache.v1alpha1.WebCache.ListCaches:input_type -> animeshon.webcache.v1alpha1.ListCachesRequest
	4, // 7: animeshon.webcache.v1alpha1.WebCache.GetCache:input_type -> animeshon.webcache.v1alpha1.GetCacheRequest
	5, // 8: animeshon.webcache.v1alpha1.WebCache.DeleteCache:input_type -> animeshon.webcache.v1alpha1.DeleteCacheRequest
	0, // 9: animeshon.webcache.v1alpha1.WebCache.CreateCache:output_type -> animeshon.webcache.v1alpha1.Cache
	3, // 10: animeshon.webcache.v1alpha1.WebCache.ListCaches:output_type -> animeshon.webcache.v1alpha1.ListCachesResponse
	0, // 11: animeshon.webcache.v1alpha1.WebCache.GetCache:output_type -> animeshon.webcache.v1alpha1.Cache
	8, // 12: animeshon.webcache.v1alpha1.WebCache.DeleteCache:output_type -> google.protobuf.Empty
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_animeshon_webcache_v1alpha1_webcache_proto_init() }
func file_animeshon_webcache_v1alpha1_webcache_proto_init() {
	if File_animeshon_webcache_v1alpha1_webcache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cache); i {
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
		file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCacheRequest); i {
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
		file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCachesRequest); i {
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
		file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCachesResponse); i {
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
		file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCacheRequest); i {
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
		file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCacheRequest); i {
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
			RawDescriptor: file_animeshon_webcache_v1alpha1_webcache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_animeshon_webcache_v1alpha1_webcache_proto_goTypes,
		DependencyIndexes: file_animeshon_webcache_v1alpha1_webcache_proto_depIdxs,
		MessageInfos:      file_animeshon_webcache_v1alpha1_webcache_proto_msgTypes,
	}.Build()
	File_animeshon_webcache_v1alpha1_webcache_proto = out.File
	file_animeshon_webcache_v1alpha1_webcache_proto_rawDesc = nil
	file_animeshon_webcache_v1alpha1_webcache_proto_goTypes = nil
	file_animeshon_webcache_v1alpha1_webcache_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WebCacheClient is the client API for WebCache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WebCacheClient interface {
	CreateCache(ctx context.Context, in *CreateCacheRequest, opts ...grpc.CallOption) (*Cache, error)
	ListCaches(ctx context.Context, in *ListCachesRequest, opts ...grpc.CallOption) (*ListCachesResponse, error)
	// See https://google.aip.dev/162#referencing-revisions for more information.
	GetCache(ctx context.Context, in *GetCacheRequest, opts ...grpc.CallOption) (*Cache, error)
	DeleteCache(ctx context.Context, in *DeleteCacheRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type webCacheClient struct {
	cc grpc.ClientConnInterface
}

func NewWebCacheClient(cc grpc.ClientConnInterface) WebCacheClient {
	return &webCacheClient{cc}
}

func (c *webCacheClient) CreateCache(ctx context.Context, in *CreateCacheRequest, opts ...grpc.CallOption) (*Cache, error) {
	out := new(Cache)
	err := c.cc.Invoke(ctx, "/animeshon.webcache.v1alpha1.WebCache/CreateCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webCacheClient) ListCaches(ctx context.Context, in *ListCachesRequest, opts ...grpc.CallOption) (*ListCachesResponse, error) {
	out := new(ListCachesResponse)
	err := c.cc.Invoke(ctx, "/animeshon.webcache.v1alpha1.WebCache/ListCaches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webCacheClient) GetCache(ctx context.Context, in *GetCacheRequest, opts ...grpc.CallOption) (*Cache, error) {
	out := new(Cache)
	err := c.cc.Invoke(ctx, "/animeshon.webcache.v1alpha1.WebCache/GetCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webCacheClient) DeleteCache(ctx context.Context, in *DeleteCacheRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/animeshon.webcache.v1alpha1.WebCache/DeleteCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebCacheServer is the server API for WebCache service.
type WebCacheServer interface {
	CreateCache(context.Context, *CreateCacheRequest) (*Cache, error)
	ListCaches(context.Context, *ListCachesRequest) (*ListCachesResponse, error)
	// See https://google.aip.dev/162#referencing-revisions for more information.
	GetCache(context.Context, *GetCacheRequest) (*Cache, error)
	DeleteCache(context.Context, *DeleteCacheRequest) (*emptypb.Empty, error)
}

// UnimplementedWebCacheServer can be embedded to have forward compatible implementations.
type UnimplementedWebCacheServer struct {
}

func (*UnimplementedWebCacheServer) CreateCache(context.Context, *CreateCacheRequest) (*Cache, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCache not implemented")
}
func (*UnimplementedWebCacheServer) ListCaches(context.Context, *ListCachesRequest) (*ListCachesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCaches not implemented")
}
func (*UnimplementedWebCacheServer) GetCache(context.Context, *GetCacheRequest) (*Cache, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCache not implemented")
}
func (*UnimplementedWebCacheServer) DeleteCache(context.Context, *DeleteCacheRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCache not implemented")
}

func RegisterWebCacheServer(s *grpc.Server, srv WebCacheServer) {
	s.RegisterService(&_WebCache_serviceDesc, srv)
}

func _WebCache_CreateCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebCacheServer).CreateCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.webcache.v1alpha1.WebCache/CreateCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebCacheServer).CreateCache(ctx, req.(*CreateCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebCache_ListCaches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCachesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebCacheServer).ListCaches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.webcache.v1alpha1.WebCache/ListCaches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebCacheServer).ListCaches(ctx, req.(*ListCachesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebCache_GetCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebCacheServer).GetCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.webcache.v1alpha1.WebCache/GetCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebCacheServer).GetCache(ctx, req.(*GetCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebCache_DeleteCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebCacheServer).DeleteCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/animeshon.webcache.v1alpha1.WebCache/DeleteCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebCacheServer).DeleteCache(ctx, req.(*DeleteCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WebCache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "animeshon.webcache.v1alpha1.WebCache",
	HandlerType: (*WebCacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCache",
			Handler:    _WebCache_CreateCache_Handler,
		},
		{
			MethodName: "ListCaches",
			Handler:    _WebCache_ListCaches_Handler,
		},
		{
			MethodName: "GetCache",
			Handler:    _WebCache_GetCache_Handler,
		},
		{
			MethodName: "DeleteCache",
			Handler:    _WebCache_DeleteCache_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "animeshon/webcache/v1alpha1/webcache.proto",
}
