// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: nflog/nflogpb/nflog.proto

package nflogpb

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

type Receiver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configured name of the receiver group.
	GroupName string `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	// Name of the integration of the receiver.
	Integration string `protobuf:"bytes,2,opt,name=integration,proto3" json:"integration,omitempty"`
	// Index of the receiver with respect to the integration.
	// Every integration in a group may have 0..N configurations.
	Idx uint32 `protobuf:"varint,3,opt,name=idx,proto3" json:"idx,omitempty"`
}

func (x *Receiver) Reset() {
	*x = Receiver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nflog_nflogpb_nflog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receiver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receiver) ProtoMessage() {}

func (x *Receiver) ProtoReflect() protoreflect.Message {
	mi := &file_nflog_nflogpb_nflog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receiver.ProtoReflect.Descriptor instead.
func (*Receiver) Descriptor() ([]byte, []int) {
	return file_nflog_nflogpb_nflog_proto_rawDescGZIP(), []int{0}
}

func (x *Receiver) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *Receiver) GetIntegration() string {
	if x != nil {
		return x.Integration
	}
	return ""
}

func (x *Receiver) GetIdx() uint32 {
	if x != nil {
		return x.Idx
	}
	return 0
}

// Entry holds information about a successful notification
// sent to a receiver.
type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key identifying the dispatching group.
	GroupKey []byte `protobuf:"bytes,1,opt,name=group_key,json=groupKey,proto3" json:"group_key,omitempty"`
	// The receiver that was notified.
	Receiver *Receiver `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// Hash over the state of the group at notification time.
	// Deprecated in favor of FiringAlerts field, but kept for compatibility.
	GroupHash []byte `protobuf:"bytes,3,opt,name=group_hash,json=groupHash,proto3" json:"group_hash,omitempty"`
	// Whether the notification was about a resolved alert.
	// Deprecated in favor of ResolvedAlerts field, but kept for compatibility.
	Resolved bool `protobuf:"varint,4,opt,name=resolved,proto3" json:"resolved,omitempty"`
	// Timestamp of the succeeding notification.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// FiringAlerts list of hashes of firing alerts at the last notification time.
	FiringAlerts []uint64 `protobuf:"varint,6,rep,packed,name=firing_alerts,json=firingAlerts,proto3" json:"firing_alerts,omitempty"`
	// ResolvedAlerts list of hashes of resolved alerts at the last notification time.
	ResolvedAlerts []uint64 `protobuf:"varint,7,rep,packed,name=resolved_alerts,json=resolvedAlerts,proto3" json:"resolved_alerts,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nflog_nflogpb_nflog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_nflog_nflogpb_nflog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_nflog_nflogpb_nflog_proto_rawDescGZIP(), []int{1}
}

func (x *Entry) GetGroupKey() []byte {
	if x != nil {
		return x.GroupKey
	}
	return nil
}

func (x *Entry) GetReceiver() *Receiver {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *Entry) GetGroupHash() []byte {
	if x != nil {
		return x.GroupHash
	}
	return nil
}

func (x *Entry) GetResolved() bool {
	if x != nil {
		return x.Resolved
	}
	return false
}

func (x *Entry) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Entry) GetFiringAlerts() []uint64 {
	if x != nil {
		return x.FiringAlerts
	}
	return nil
}

func (x *Entry) GetResolvedAlerts() []uint64 {
	if x != nil {
		return x.ResolvedAlerts
	}
	return nil
}

// MeshEntry is a wrapper message to communicate a notify log
// entry through a mesh network.
type MeshEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The original raw notify log entry.
	Entry *Entry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	// A timestamp indicating when the mesh peer should evict
	// the log entry from its state.
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *MeshEntry) Reset() {
	*x = MeshEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nflog_nflogpb_nflog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshEntry) ProtoMessage() {}

func (x *MeshEntry) ProtoReflect() protoreflect.Message {
	mi := &file_nflog_nflogpb_nflog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshEntry.ProtoReflect.Descriptor instead.
func (*MeshEntry) Descriptor() ([]byte, []int) {
	return file_nflog_nflogpb_nflog_proto_rawDescGZIP(), []int{2}
}

func (x *MeshEntry) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *MeshEntry) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

var File_nflog_nflogpb_nflog_proto protoreflect.FileDescriptor

var file_nflog_nflogpb_nflog_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6e, 0x66, 0x6c, 0x6f, 0x67, 0x2f, 0x6e, 0x66, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2f,
	0x6e, 0x66, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x66, 0x6c,
	0x6f, 0x67, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x69, 0x64, 0x78, 0x22, 0x96, 0x02, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x6e, 0x66, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x23, 0x0a, 0x0d, 0x66, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64,
	0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0e, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0x6c, 0x0a,
	0x09, 0x4d, 0x65, 0x73, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x66, 0x6c, 0x6f,
	0x67, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x42, 0x2c, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x6f, 0x63, 0x6f, 0x6f,
	0x73, 0x2f, 0x6d, 0x73, 0x67, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x6e, 0x66, 0x6c, 0x6f,
	0x67, 0x2f, 0x6e, 0x66, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_nflog_nflogpb_nflog_proto_rawDescOnce sync.Once
	file_nflog_nflogpb_nflog_proto_rawDescData = file_nflog_nflogpb_nflog_proto_rawDesc
)

func file_nflog_nflogpb_nflog_proto_rawDescGZIP() []byte {
	file_nflog_nflogpb_nflog_proto_rawDescOnce.Do(func() {
		file_nflog_nflogpb_nflog_proto_rawDescData = protoimpl.X.CompressGZIP(file_nflog_nflogpb_nflog_proto_rawDescData)
	})
	return file_nflog_nflogpb_nflog_proto_rawDescData
}

var file_nflog_nflogpb_nflog_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nflog_nflogpb_nflog_proto_goTypes = []interface{}{
	(*Receiver)(nil),              // 0: nflogpb.Receiver
	(*Entry)(nil),                 // 1: nflogpb.Entry
	(*MeshEntry)(nil),             // 2: nflogpb.MeshEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_nflog_nflogpb_nflog_proto_depIdxs = []int32{
	0, // 0: nflogpb.Entry.receiver:type_name -> nflogpb.Receiver
	3, // 1: nflogpb.Entry.timestamp:type_name -> google.protobuf.Timestamp
	1, // 2: nflogpb.MeshEntry.entry:type_name -> nflogpb.Entry
	3, // 3: nflogpb.MeshEntry.expires_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nflog_nflogpb_nflog_proto_init() }
func file_nflog_nflogpb_nflog_proto_init() {
	if File_nflog_nflogpb_nflog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nflog_nflogpb_nflog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Receiver); i {
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
		file_nflog_nflogpb_nflog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_nflog_nflogpb_nflog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshEntry); i {
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
			RawDescriptor: file_nflog_nflogpb_nflog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nflog_nflogpb_nflog_proto_goTypes,
		DependencyIndexes: file_nflog_nflogpb_nflog_proto_depIdxs,
		MessageInfos:      file_nflog_nflogpb_nflog_proto_msgTypes,
	}.Build()
	File_nflog_nflogpb_nflog_proto = out.File
	file_nflog_nflogpb_nflog_proto_rawDesc = nil
	file_nflog_nflogpb_nflog_proto_goTypes = nil
	file_nflog_nflogpb_nflog_proto_depIdxs = nil
}
