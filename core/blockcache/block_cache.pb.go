// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: core/blockcache/block_cache.proto

package blockcache

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

type BcMessageType int32

const (
	BcMessageType_LinkType                    BcMessageType = 0
	BcMessageType_UpdateActiveType            BcMessageType = 1
	BcMessageType_UpdateLinkedRootWitnessType BcMessageType = 2
)

// Enum value maps for BcMessageType.
var (
	BcMessageType_name = map[int32]string{
		0: "LinkType",
		1: "UpdateActiveType",
		2: "UpdateLinkedRootWitnessType",
	}
	BcMessageType_value = map[string]int32{
		"LinkType":                    0,
		"UpdateActiveType":            1,
		"UpdateLinkedRootWitnessType": 2,
	}
)

func (x BcMessageType) Enum() *BcMessageType {
	p := new(BcMessageType)
	*p = x
	return p
}

func (x BcMessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BcMessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_core_blockcache_block_cache_proto_enumTypes[0].Descriptor()
}

func (BcMessageType) Type() protoreflect.EnumType {
	return &file_core_blockcache_block_cache_proto_enumTypes[0]
}

func (x BcMessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BcMessageType.Descriptor instead.
func (BcMessageType) EnumDescriptor() ([]byte, []int) {
	return file_core_blockcache_block_cache_proto_rawDescGZIP(), []int{0}
}

type BcMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte        `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Type BcMessageType `protobuf:"varint,2,opt,name=type,proto3,enum=blockcache.BcMessageType" json:"type,omitempty"`
}

func (x *BcMessage) Reset() {
	*x = BcMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_blockcache_block_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BcMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BcMessage) ProtoMessage() {}

func (x *BcMessage) ProtoReflect() protoreflect.Message {
	mi := &file_core_blockcache_block_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BcMessage.ProtoReflect.Descriptor instead.
func (*BcMessage) Descriptor() ([]byte, []int) {
	return file_core_blockcache_block_cache_proto_rawDescGZIP(), []int{0}
}

func (x *BcMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BcMessage) GetType() BcMessageType {
	if x != nil {
		return x.Type
	}
	return BcMessageType_LinkType
}

type BlockCacheRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockBytes  []byte       `protobuf:"bytes,1,opt,name=blockBytes,proto3" json:"blockBytes,omitempty"`
	WitnessList *WitnessList `protobuf:"bytes,2,opt,name=witnessList,proto3" json:"witnessList,omitempty"`
	SerialNum   int64        `protobuf:"varint,3,opt,name=serialNum,proto3" json:"serialNum,omitempty"`
}

func (x *BlockCacheRaw) Reset() {
	*x = BlockCacheRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_blockcache_block_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockCacheRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockCacheRaw) ProtoMessage() {}

func (x *BlockCacheRaw) ProtoReflect() protoreflect.Message {
	mi := &file_core_blockcache_block_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockCacheRaw.ProtoReflect.Descriptor instead.
func (*BlockCacheRaw) Descriptor() ([]byte, []int) {
	return file_core_blockcache_block_cache_proto_rawDescGZIP(), []int{1}
}

func (x *BlockCacheRaw) GetBlockBytes() []byte {
	if x != nil {
		return x.BlockBytes
	}
	return nil
}

func (x *BlockCacheRaw) GetWitnessList() *WitnessList {
	if x != nil {
		return x.WitnessList
	}
	return nil
}

func (x *BlockCacheRaw) GetSerialNum() int64 {
	if x != nil {
		return x.SerialNum
	}
	return 0
}

type UpdateActiveRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHashBytes []byte       `protobuf:"bytes,1,opt,name=blockHashBytes,proto3" json:"blockHashBytes,omitempty"`
	WitnessList    *WitnessList `protobuf:"bytes,2,opt,name=witnessList,proto3" json:"witnessList,omitempty"`
}

func (x *UpdateActiveRaw) Reset() {
	*x = UpdateActiveRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_blockcache_block_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateActiveRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateActiveRaw) ProtoMessage() {}

func (x *UpdateActiveRaw) ProtoReflect() protoreflect.Message {
	mi := &file_core_blockcache_block_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateActiveRaw.ProtoReflect.Descriptor instead.
func (*UpdateActiveRaw) Descriptor() ([]byte, []int) {
	return file_core_blockcache_block_cache_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateActiveRaw) GetBlockHashBytes() []byte {
	if x != nil {
		return x.BlockHashBytes
	}
	return nil
}

func (x *UpdateActiveRaw) GetWitnessList() *WitnessList {
	if x != nil {
		return x.WitnessList
	}
	return nil
}

type UpdateLinkedRootWitnessRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LinkedRootWitness []string `protobuf:"bytes,1,rep,name=linkedRootWitness,proto3" json:"linkedRootWitness,omitempty"`
	BlockHashBytes    []byte   `protobuf:"bytes,2,opt,name=blockHashBytes,proto3" json:"blockHashBytes,omitempty"`
}

func (x *UpdateLinkedRootWitnessRaw) Reset() {
	*x = UpdateLinkedRootWitnessRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_blockcache_block_cache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLinkedRootWitnessRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLinkedRootWitnessRaw) ProtoMessage() {}

func (x *UpdateLinkedRootWitnessRaw) ProtoReflect() protoreflect.Message {
	mi := &file_core_blockcache_block_cache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLinkedRootWitnessRaw.ProtoReflect.Descriptor instead.
func (*UpdateLinkedRootWitnessRaw) Descriptor() ([]byte, []int) {
	return file_core_blockcache_block_cache_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateLinkedRootWitnessRaw) GetLinkedRootWitness() []string {
	if x != nil {
		return x.LinkedRootWitness
	}
	return nil
}

func (x *UpdateLinkedRootWitnessRaw) GetBlockHashBytes() []byte {
	if x != nil {
		return x.BlockHashBytes
	}
	return nil
}

type WitnessList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActiveWitnessList    []string `protobuf:"bytes,1,rep,name=activeWitnessList,proto3" json:"activeWitnessList,omitempty"`
	PendingWitnessList   []string `protobuf:"bytes,2,rep,name=pendingWitnessList,proto3" json:"pendingWitnessList,omitempty"`
	PendingWitnessNumber int64    `protobuf:"varint,3,opt,name=pendingWitnessNumber,proto3" json:"pendingWitnessNumber,omitempty"`
	WitnessInfo          []string `protobuf:"bytes,4,rep,name=witnessInfo,proto3" json:"witnessInfo,omitempty"`
}

func (x *WitnessList) Reset() {
	*x = WitnessList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_blockcache_block_cache_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WitnessList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WitnessList) ProtoMessage() {}

func (x *WitnessList) ProtoReflect() protoreflect.Message {
	mi := &file_core_blockcache_block_cache_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WitnessList.ProtoReflect.Descriptor instead.
func (*WitnessList) Descriptor() ([]byte, []int) {
	return file_core_blockcache_block_cache_proto_rawDescGZIP(), []int{4}
}

func (x *WitnessList) GetActiveWitnessList() []string {
	if x != nil {
		return x.ActiveWitnessList
	}
	return nil
}

func (x *WitnessList) GetPendingWitnessList() []string {
	if x != nil {
		return x.PendingWitnessList
	}
	return nil
}

func (x *WitnessList) GetPendingWitnessNumber() int64 {
	if x != nil {
		return x.PendingWitnessNumber
	}
	return 0
}

func (x *WitnessList) GetWitnessInfo() []string {
	if x != nil {
		return x.WitnessInfo
	}
	return nil
}

type WitnessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetID string `protobuf:"bytes,1,opt,name=NetID,proto3" json:"NetID,omitempty"`
}

func (x *WitnessInfo) Reset() {
	*x = WitnessInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_blockcache_block_cache_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WitnessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WitnessInfo) ProtoMessage() {}

func (x *WitnessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_core_blockcache_block_cache_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WitnessInfo.ProtoReflect.Descriptor instead.
func (*WitnessInfo) Descriptor() ([]byte, []int) {
	return file_core_blockcache_block_cache_proto_rawDescGZIP(), []int{5}
}

func (x *WitnessInfo) GetNetID() string {
	if x != nil {
		return x.NetID
	}
	return ""
}

var File_core_blockcache_block_cache_proto protoreflect.FileDescriptor

var file_core_blockcache_block_cache_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x22,
	0x4e, 0x0a, 0x09, 0x62, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x63, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x88, 0x01, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x61,
	0x77, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x39, 0x0a, 0x0b, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x0b, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x22, 0x74, 0x0a, 0x0f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61, 0x77, 0x12, 0x26, 0x0a,
	0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0b, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x0b, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x72, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x52, 0x6f, 0x6f, 0x74, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x61, 0x77, 0x12, 0x2c,
	0x0a, 0x11, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x52, 0x6f, 0x6f, 0x74, 0x57, 0x69, 0x74, 0x6e,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x69, 0x6e, 0x6b, 0x65,
	0x64, 0x52, 0x6f, 0x6f, 0x74, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x22, 0xc1, 0x01, 0x0a, 0x0b, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x57, 0x69,
	0x74, 0x6e, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x57, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x57, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x14, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x23, 0x0a, 0x0b, 0x57, 0x69, 0x74, 0x6e,
	0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x65, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4e, 0x65, 0x74, 0x49, 0x44, 0x2a, 0x54, 0x0a,
	0x0d, 0x62, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c,
	0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x65, 0x64, 0x52, 0x6f, 0x6f, 0x74, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x10, 0x02, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6f, 0x73, 0x74, 0x2d, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x2f,
	0x67, 0x6f, 0x2d, 0x69, 0x6f, 0x73, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_core_blockcache_block_cache_proto_rawDescOnce sync.Once
	file_core_blockcache_block_cache_proto_rawDescData = file_core_blockcache_block_cache_proto_rawDesc
)

func file_core_blockcache_block_cache_proto_rawDescGZIP() []byte {
	file_core_blockcache_block_cache_proto_rawDescOnce.Do(func() {
		file_core_blockcache_block_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_blockcache_block_cache_proto_rawDescData)
	})
	return file_core_blockcache_block_cache_proto_rawDescData
}

var file_core_blockcache_block_cache_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_blockcache_block_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_core_blockcache_block_cache_proto_goTypes = []interface{}{
	(BcMessageType)(0),                 // 0: blockcache.bcMessageType
	(*BcMessage)(nil),                  // 1: blockcache.bcMessage
	(*BlockCacheRaw)(nil),              // 2: blockcache.BlockCacheRaw
	(*UpdateActiveRaw)(nil),            // 3: blockcache.UpdateActiveRaw
	(*UpdateLinkedRootWitnessRaw)(nil), // 4: blockcache.UpdateLinkedRootWitnessRaw
	(*WitnessList)(nil),                // 5: blockcache.WitnessList
	(*WitnessInfo)(nil),                // 6: blockcache.WitnessInfo
}
var file_core_blockcache_block_cache_proto_depIdxs = []int32{
	0, // 0: blockcache.bcMessage.type:type_name -> blockcache.bcMessageType
	5, // 1: blockcache.BlockCacheRaw.witnessList:type_name -> blockcache.WitnessList
	5, // 2: blockcache.UpdateActiveRaw.witnessList:type_name -> blockcache.WitnessList
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_core_blockcache_block_cache_proto_init() }
func file_core_blockcache_block_cache_proto_init() {
	if File_core_blockcache_block_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_blockcache_block_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BcMessage); i {
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
		file_core_blockcache_block_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockCacheRaw); i {
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
		file_core_blockcache_block_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateActiveRaw); i {
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
		file_core_blockcache_block_cache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLinkedRootWitnessRaw); i {
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
		file_core_blockcache_block_cache_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WitnessList); i {
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
		file_core_blockcache_block_cache_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WitnessInfo); i {
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
			RawDescriptor: file_core_blockcache_block_cache_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_blockcache_block_cache_proto_goTypes,
		DependencyIndexes: file_core_blockcache_block_cache_proto_depIdxs,
		EnumInfos:         file_core_blockcache_block_cache_proto_enumTypes,
		MessageInfos:      file_core_blockcache_block_cache_proto_msgTypes,
	}.Build()
	File_core_blockcache_block_cache_proto = out.File
	file_core_blockcache_block_cache_proto_rawDesc = nil
	file_core_blockcache_block_cache_proto_goTypes = nil
	file_core_blockcache_block_cache_proto_depIdxs = nil
}
