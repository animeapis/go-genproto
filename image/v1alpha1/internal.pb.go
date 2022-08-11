// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: animeshon/image/v1alpha1/internal.proto

package image

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// This is a private internal structure used to store metadata information
// about a specific Image. This structure is never exposed to the public.
type ImageEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The image resource id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// A reference to a parent resource.
	//
	// Types that are assignable to Parent:
	//	*ImageEntry_AlbumId
	Parent isImageEntry_Parent `protobuf_oneof:"parent"`
	// The uri of the image file.
	// Example: gs://my-bucket/my-user/my-album/my-image.jpeg
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// The MIME type of the image.
	MimeType string `protobuf:"bytes,4,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// The size of the image in bytes.
	FileSize int32 `protobuf:"varint,5,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	// The width of image in pixels.
	Width int32 `protobuf:"varint,6,opt,name=width,proto3" json:"width,omitempty"`
	// The height of image in pixels.
	Height int32 `protobuf:"varint,7,opt,name=height,proto3" json:"height,omitempty"`
	// When the image was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *ImageEntry) Reset() {
	*x = ImageEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_image_v1alpha1_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageEntry) ProtoMessage() {}

func (x *ImageEntry) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_image_v1alpha1_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageEntry.ProtoReflect.Descriptor instead.
func (*ImageEntry) Descriptor() ([]byte, []int) {
	return file_animeshon_image_v1alpha1_internal_proto_rawDescGZIP(), []int{0}
}

func (x *ImageEntry) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (m *ImageEntry) GetParent() isImageEntry_Parent {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (x *ImageEntry) GetAlbumId() int64 {
	if x, ok := x.GetParent().(*ImageEntry_AlbumId); ok {
		return x.AlbumId
	}
	return 0
}

func (x *ImageEntry) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *ImageEntry) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *ImageEntry) GetFileSize() int32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *ImageEntry) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ImageEntry) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ImageEntry) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type isImageEntry_Parent interface {
	isImageEntry_Parent()
}

type ImageEntry_AlbumId struct {
	// The album resource id. Example: `albums/123`.
	AlbumId int64 `protobuf:"varint,2,opt,name=album_id,json=albumId,proto3,oneof"`
}

func (*ImageEntry_AlbumId) isImageEntry_Parent() {}

// This is a private internal structure used to store metadata information
// about a specific Album. This structure is never exposed to the public.
type AlbumEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The album resource id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// A reference to a parent resource.
	//
	// Types that are assignable to Parent:
	//	*AlbumEntry_OrganizationId
	//	*AlbumEntry_UserId
	Parent isAlbumEntry_Parent `protobuf_oneof:"parent"`
	// The album display name.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The bucket where to store resources.
	Bucket string `protobuf:"bytes,5,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Whether the album is system managed.
	SystemManaged bool `protobuf:"varint,6,opt,name=system_managed,json=systemManaged,proto3" json:"system_managed,omitempty"`
	// Whether images should be available through the image search.
	EnableSearch bool `protobuf:"varint,7,opt,name=enable_search,json=enableSearch,proto3" json:"enable_search,omitempty"`
	// Whether images can be annotated.
	EnableAnnotation bool `protobuf:"varint,8,opt,name=enable_annotation,json=enableAnnotation,proto3" json:"enable_annotation,omitempty"`
	// Whether images should be client-side encrypted.
	EnableEncryption bool `protobuf:"varint,9,opt,name=enable_encryption,json=enableEncryption,proto3" json:"enable_encryption,omitempty"`
	// Whether the content of this album is highly visible by the public such as
	// profile pictures and banners.
	HighVisibility bool `protobuf:"varint,10,opt,name=high_visibility,json=highVisibility,proto3" json:"high_visibility,omitempty"`
	// When the album was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// When the album was updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *AlbumEntry) Reset() {
	*x = AlbumEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_image_v1alpha1_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumEntry) ProtoMessage() {}

func (x *AlbumEntry) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_image_v1alpha1_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumEntry.ProtoReflect.Descriptor instead.
func (*AlbumEntry) Descriptor() ([]byte, []int) {
	return file_animeshon_image_v1alpha1_internal_proto_rawDescGZIP(), []int{1}
}

func (x *AlbumEntry) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (m *AlbumEntry) GetParent() isAlbumEntry_Parent {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (x *AlbumEntry) GetOrganizationId() int64 {
	if x, ok := x.GetParent().(*AlbumEntry_OrganizationId); ok {
		return x.OrganizationId
	}
	return 0
}

func (x *AlbumEntry) GetUserId() int64 {
	if x, ok := x.GetParent().(*AlbumEntry_UserId); ok {
		return x.UserId
	}
	return 0
}

func (x *AlbumEntry) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AlbumEntry) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *AlbumEntry) GetSystemManaged() bool {
	if x != nil {
		return x.SystemManaged
	}
	return false
}

func (x *AlbumEntry) GetEnableSearch() bool {
	if x != nil {
		return x.EnableSearch
	}
	return false
}

func (x *AlbumEntry) GetEnableAnnotation() bool {
	if x != nil {
		return x.EnableAnnotation
	}
	return false
}

func (x *AlbumEntry) GetEnableEncryption() bool {
	if x != nil {
		return x.EnableEncryption
	}
	return false
}

func (x *AlbumEntry) GetHighVisibility() bool {
	if x != nil {
		return x.HighVisibility
	}
	return false
}

func (x *AlbumEntry) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *AlbumEntry) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type isAlbumEntry_Parent interface {
	isAlbumEntry_Parent()
}

type AlbumEntry_OrganizationId struct {
	// The organization resource id. Example: `organizations/456`.
	OrganizationId int64 `protobuf:"varint,2,opt,name=organization_id,json=organizationId,proto3,oneof"`
}

type AlbumEntry_UserId struct {
	// The user resource id. Example: `users/789`.
	UserId int64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3,oneof"`
}

func (*AlbumEntry_OrganizationId) isAlbumEntry_Parent() {}

func (*AlbumEntry_UserId) isAlbumEntry_Parent() {}

type RouteEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A reference to a parent resource.
	//
	// Types that are assignable to Parent:
	//	*RouteEntry_OrganizationId
	//	*RouteEntry_UserId
	Parent isRouteEntry_Parent `protobuf_oneof:"parent"`
	// The album resource id.
	AlbumId int64 `protobuf:"varint,3,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
	// The image resource id.
	ImageId int64 `protobuf:"varint,4,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
}

func (x *RouteEntry) Reset() {
	*x = RouteEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_image_v1alpha1_internal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteEntry) ProtoMessage() {}

func (x *RouteEntry) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_image_v1alpha1_internal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteEntry.ProtoReflect.Descriptor instead.
func (*RouteEntry) Descriptor() ([]byte, []int) {
	return file_animeshon_image_v1alpha1_internal_proto_rawDescGZIP(), []int{2}
}

func (m *RouteEntry) GetParent() isRouteEntry_Parent {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (x *RouteEntry) GetOrganizationId() int64 {
	if x, ok := x.GetParent().(*RouteEntry_OrganizationId); ok {
		return x.OrganizationId
	}
	return 0
}

func (x *RouteEntry) GetUserId() int64 {
	if x, ok := x.GetParent().(*RouteEntry_UserId); ok {
		return x.UserId
	}
	return 0
}

func (x *RouteEntry) GetAlbumId() int64 {
	if x != nil {
		return x.AlbumId
	}
	return 0
}

func (x *RouteEntry) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

type isRouteEntry_Parent interface {
	isRouteEntry_Parent()
}

type RouteEntry_OrganizationId struct {
	// The organization resource id. Example: `organizations/456`.
	OrganizationId int64 `protobuf:"varint,1,opt,name=organization_id,json=organizationId,proto3,oneof"`
}

type RouteEntry_UserId struct {
	// The user resource id. Example: `users/789`.
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3,oneof"`
}

func (*RouteEntry_OrganizationId) isRouteEntry_Parent() {}

func (*RouteEntry_UserId) isRouteEntry_Parent() {}

var File_animeshon_image_v1alpha1_internal_proto protoreflect.FileDescriptor

var file_animeshon_image_v1alpha1_internal_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x69, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x22, 0xf0, 0x03, 0x0a, 0x0a, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x29, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0e, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x2b, 0x0a,
	0x11, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x68, 0x69, 0x67, 0x68, 0x5f,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x68, 0x69, 0x67, 0x68, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x29, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x42,
	0x08, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x74, 0x0a, 0x1c, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0xea, 0x02, 0x1a, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x3a, 0x3a,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x3a, 0x3a, 0x76, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_animeshon_image_v1alpha1_internal_proto_rawDescOnce sync.Once
	file_animeshon_image_v1alpha1_internal_proto_rawDescData = file_animeshon_image_v1alpha1_internal_proto_rawDesc
)

func file_animeshon_image_v1alpha1_internal_proto_rawDescGZIP() []byte {
	file_animeshon_image_v1alpha1_internal_proto_rawDescOnce.Do(func() {
		file_animeshon_image_v1alpha1_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_animeshon_image_v1alpha1_internal_proto_rawDescData)
	})
	return file_animeshon_image_v1alpha1_internal_proto_rawDescData
}

var file_animeshon_image_v1alpha1_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_animeshon_image_v1alpha1_internal_proto_goTypes = []interface{}{
	(*ImageEntry)(nil),            // 0: animeshon.image.v1alpha1.ImageEntry
	(*AlbumEntry)(nil),            // 1: animeshon.image.v1alpha1.AlbumEntry
	(*RouteEntry)(nil),            // 2: animeshon.image.v1alpha1.RouteEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_animeshon_image_v1alpha1_internal_proto_depIdxs = []int32{
	3, // 0: animeshon.image.v1alpha1.ImageEntry.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: animeshon.image.v1alpha1.AlbumEntry.create_time:type_name -> google.protobuf.Timestamp
	3, // 2: animeshon.image.v1alpha1.AlbumEntry.update_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_animeshon_image_v1alpha1_internal_proto_init() }
func file_animeshon_image_v1alpha1_internal_proto_init() {
	if File_animeshon_image_v1alpha1_internal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_animeshon_image_v1alpha1_internal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageEntry); i {
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
		file_animeshon_image_v1alpha1_internal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumEntry); i {
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
		file_animeshon_image_v1alpha1_internal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteEntry); i {
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
	file_animeshon_image_v1alpha1_internal_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ImageEntry_AlbumId)(nil),
	}
	file_animeshon_image_v1alpha1_internal_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AlbumEntry_OrganizationId)(nil),
		(*AlbumEntry_UserId)(nil),
	}
	file_animeshon_image_v1alpha1_internal_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*RouteEntry_OrganizationId)(nil),
		(*RouteEntry_UserId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_animeshon_image_v1alpha1_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_animeshon_image_v1alpha1_internal_proto_goTypes,
		DependencyIndexes: file_animeshon_image_v1alpha1_internal_proto_depIdxs,
		MessageInfos:      file_animeshon_image_v1alpha1_internal_proto_msgTypes,
	}.Build()
	File_animeshon_image_v1alpha1_internal_proto = out.File
	file_animeshon_image_v1alpha1_internal_proto_rawDesc = nil
	file_animeshon_image_v1alpha1_internal_proto_goTypes = nil
	file_animeshon_image_v1alpha1_internal_proto_depIdxs = nil
}
