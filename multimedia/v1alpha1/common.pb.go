// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: animeshon/multimedia/v1alpha1/common.proto

package multimedia

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PublishingType int32

const (
	// The default value. This value is used if the state is omitted.
	PublishingType_PUBLISHING_TYPE_UNSPECIFIED PublishingType = 0
	PublishingType_SELF                        PublishingType = 1
	PublishingType_CORPORATE                   PublishingType = 2
)

// Enum value maps for PublishingType.
var (
	PublishingType_name = map[int32]string{
		0: "PUBLISHING_TYPE_UNSPECIFIED",
		1: "SELF",
		2: "CORPORATE",
	}
	PublishingType_value = map[string]int32{
		"PUBLISHING_TYPE_UNSPECIFIED": 0,
		"SELF":                        1,
		"CORPORATE":                   2,
	}
)

func (x PublishingType) Enum() *PublishingType {
	p := new(PublishingType)
	*p = x
	return p
}

func (x PublishingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PublishingType) Descriptor() protoreflect.EnumDescriptor {
	return file_animeshon_multimedia_v1alpha1_common_proto_enumTypes[0].Descriptor()
}

func (PublishingType) Type() protoreflect.EnumType {
	return &file_animeshon_multimedia_v1alpha1_common_proto_enumTypes[0]
}

func (x PublishingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PublishingType.Descriptor instead.
func (PublishingType) EnumDescriptor() ([]byte, []int) {
	return file_animeshon_multimedia_v1alpha1_common_proto_rawDescGZIP(), []int{0}
}

type State int32

const (
	// The default value. This value is used if the state is omitted.
	State_STATE_UNSPECIFIED State = 0
	State_ONGOING           State = 1
	State_COMPLETED         State = 2
	State_SCHEDULED         State = 3
	State_INTERRUPTED       State = 4
	State_CANCELED          State = 5
	State_SUSPENDED         State = 6
	State_WORK_IN_PROGRESS  State = 7
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ONGOING",
		2: "COMPLETED",
		3: "SCHEDULED",
		4: "INTERRUPTED",
		5: "CANCELED",
		6: "SUSPENDED",
		7: "WORK_IN_PROGRESS",
	}
	State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ONGOING":           1,
		"COMPLETED":         2,
		"SCHEDULED":         3,
		"INTERRUPTED":       4,
		"CANCELED":          5,
		"SUSPENDED":         6,
		"WORK_IN_PROGRESS":  7,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_animeshon_multimedia_v1alpha1_common_proto_enumTypes[1].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_animeshon_multimedia_v1alpha1_common_proto_enumTypes[1]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_animeshon_multimedia_v1alpha1_common_proto_rawDescGZIP(), []int{1}
}

// Represents the metadata of the long-running operation.
type OperationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The time the operation was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The time the operation finished running.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Output only. Server-defined resource path for the target of the operation.
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	// Output only. Name of the verb executed by the operation.
	Verb string `protobuf:"bytes,4,opt,name=verb,proto3" json:"verb,omitempty"`
	// Output only. Human-readable status of the operation, if any.
	StatusMessage string `protobuf:"bytes,5,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	// Output only. Identifies whether the user has requested cancellation
	// of the operation. Operations that have successfully been cancelled
	// have [Operation.error][] value with a [google.rpc.Status.code][google.rpc.Status.code] of 1,
	// corresponding to `Code.CANCELLED`.
	RequestedCancellation bool `protobuf:"varint,6,opt,name=requested_cancellation,json=requestedCancellation,proto3" json:"requested_cancellation,omitempty"`
	// Output only. API version used to start the operation.
	ApiVersion string `protobuf:"bytes,7,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// Output only.
	ProgressPercentage int32 `protobuf:"varint,8,opt,name=progress_percentage,json=progressPercentage,proto3" json:"progress_percentage,omitempty"`
}

func (x *OperationMetadata) Reset() {
	*x = OperationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_animeshon_multimedia_v1alpha1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationMetadata) ProtoMessage() {}

func (x *OperationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_animeshon_multimedia_v1alpha1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationMetadata.ProtoReflect.Descriptor instead.
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return file_animeshon_multimedia_v1alpha1_common_proto_rawDescGZIP(), []int{0}
}

func (x *OperationMetadata) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *OperationMetadata) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *OperationMetadata) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *OperationMetadata) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *OperationMetadata) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *OperationMetadata) GetRequestedCancellation() bool {
	if x != nil {
		return x.RequestedCancellation
	}
	return false
}

func (x *OperationMetadata) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *OperationMetadata) GetProgressPercentage() int32 {
	if x != nil {
		return x.ProgressPercentage
	}
	return 0
}

var File_animeshon_multimedia_v1alpha1_common_proto protoreflect.FileDescriptor

var file_animeshon_multimedia_v1alpha1_common_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x03,
	0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x17,
	0x0a, 0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x04, 0x76, 0x65, 0x72, 0x62, 0x12, 0x2a, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x24, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x2a, 0x4a, 0x0a, 0x0e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x1b, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x53, 0x45, 0x4c, 0x46, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x52, 0x50,
	0x4f, 0x52, 0x41, 0x54, 0x45, 0x10, 0x02, 0x2a, 0x8d, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x4e, 0x47, 0x4f,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45,
	0x44, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x55, 0x50, 0x54,
	0x45, 0x44, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44,
	0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10,
	0x06, 0x12, 0x14, 0x0a, 0x10, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f,
	0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x07, 0x42, 0x88, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x50, 0x01, 0x5a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0xea, 0x02, 0x1f, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x6e, 0x3a, 0x3a, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x3a, 0x3a, 0x76, 0x31, 0x41, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_animeshon_multimedia_v1alpha1_common_proto_rawDescOnce sync.Once
	file_animeshon_multimedia_v1alpha1_common_proto_rawDescData = file_animeshon_multimedia_v1alpha1_common_proto_rawDesc
)

func file_animeshon_multimedia_v1alpha1_common_proto_rawDescGZIP() []byte {
	file_animeshon_multimedia_v1alpha1_common_proto_rawDescOnce.Do(func() {
		file_animeshon_multimedia_v1alpha1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_animeshon_multimedia_v1alpha1_common_proto_rawDescData)
	})
	return file_animeshon_multimedia_v1alpha1_common_proto_rawDescData
}

var file_animeshon_multimedia_v1alpha1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_animeshon_multimedia_v1alpha1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_animeshon_multimedia_v1alpha1_common_proto_goTypes = []interface{}{
	(PublishingType)(0),           // 0: animeshon.multimedia.v1alpha1.PublishingType
	(State)(0),                    // 1: animeshon.multimedia.v1alpha1.State
	(*OperationMetadata)(nil),     // 2: animeshon.multimedia.v1alpha1.OperationMetadata
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_animeshon_multimedia_v1alpha1_common_proto_depIdxs = []int32{
	3, // 0: animeshon.multimedia.v1alpha1.OperationMetadata.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: animeshon.multimedia.v1alpha1.OperationMetadata.end_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_animeshon_multimedia_v1alpha1_common_proto_init() }
func file_animeshon_multimedia_v1alpha1_common_proto_init() {
	if File_animeshon_multimedia_v1alpha1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_animeshon_multimedia_v1alpha1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationMetadata); i {
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
			RawDescriptor: file_animeshon_multimedia_v1alpha1_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_animeshon_multimedia_v1alpha1_common_proto_goTypes,
		DependencyIndexes: file_animeshon_multimedia_v1alpha1_common_proto_depIdxs,
		EnumInfos:         file_animeshon_multimedia_v1alpha1_common_proto_enumTypes,
		MessageInfos:      file_animeshon_multimedia_v1alpha1_common_proto_msgTypes,
	}.Build()
	File_animeshon_multimedia_v1alpha1_common_proto = out.File
	file_animeshon_multimedia_v1alpha1_common_proto_rawDesc = nil
	file_animeshon_multimedia_v1alpha1_common_proto_goTypes = nil
	file_animeshon_multimedia_v1alpha1_common_proto_depIdxs = nil
}
