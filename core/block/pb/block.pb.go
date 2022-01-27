// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: core/block/pb/block.proto

package pb

import (
	pb1 "github.com/iost-official/go-iost/v3/core/tx/pb"
	pb "github.com/iost-official/go-iost/v3/crypto/pb"
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

type BlockType int32

const (
	BlockType_NORMAL   BlockType = 0
	BlockType_ONLYHASH BlockType = 1
)

// Enum value maps for BlockType.
var (
	BlockType_name = map[int32]string{
		0: "NORMAL",
		1: "ONLYHASH",
	}
	BlockType_value = map[string]int32{
		"NORMAL":   0,
		"ONLYHASH": 1,
	}
)

func (x BlockType) Enum() *BlockType {
	p := new(BlockType)
	*p = x
	return p
}

func (x BlockType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockType) Descriptor() protoreflect.EnumDescriptor {
	return file_core_block_pb_block_proto_enumTypes[0].Descriptor()
}

func (BlockType) Type() protoreflect.EnumType {
	return &file_core_block_pb_block_proto_enumTypes[0]
}

func (x BlockType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockType.Descriptor instead.
func (BlockType) EnumDescriptor() ([]byte, []int) {
	return file_core_block_pb_block_proto_rawDescGZIP(), []int{0}
}

type BlockHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version             int64  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	ParentHash          []byte `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	TxMerkleHash        []byte `protobuf:"bytes,3,opt,name=txMerkleHash,proto3" json:"txMerkleHash,omitempty"`
	TxReceiptMerkleHash []byte `protobuf:"bytes,4,opt,name=txReceiptMerkleHash,proto3" json:"txReceiptMerkleHash,omitempty"`
	Info                []byte `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	Number              int64  `protobuf:"varint,6,opt,name=number,proto3" json:"number,omitempty"`
	Witness             string `protobuf:"bytes,7,opt,name=witness,proto3" json:"witness,omitempty"`
	Time                int64  `protobuf:"varint,8,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *BlockHead) Reset() {
	*x = BlockHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_block_pb_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHead) ProtoMessage() {}

func (x *BlockHead) ProtoReflect() protoreflect.Message {
	mi := &file_core_block_pb_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHead.ProtoReflect.Descriptor instead.
func (*BlockHead) Descriptor() ([]byte, []int) {
	return file_core_block_pb_block_proto_rawDescGZIP(), []int{0}
}

func (x *BlockHead) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *BlockHead) GetParentHash() []byte {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *BlockHead) GetTxMerkleHash() []byte {
	if x != nil {
		return x.TxMerkleHash
	}
	return nil
}

func (x *BlockHead) GetTxReceiptMerkleHash() []byte {
	if x != nil {
		return x.TxReceiptMerkleHash
	}
	return nil
}

func (x *BlockHead) GetInfo() []byte {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *BlockHead) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *BlockHead) GetWitness() string {
	if x != nil {
		return x.Witness
	}
	return ""
}

func (x *BlockHead) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head          *BlockHead       `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	Sign          *pb.Signature    `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
	Txs           []*pb1.Tx        `protobuf:"bytes,3,rep,name=txs,proto3" json:"txs,omitempty"`
	Receipts      []*pb1.TxReceipt `protobuf:"bytes,4,rep,name=receipts,proto3" json:"receipts,omitempty"`
	TxHashes      [][]byte         `protobuf:"bytes,5,rep,name=txHashes,proto3" json:"txHashes,omitempty"`
	ReceiptHashes [][]byte         `protobuf:"bytes,6,rep,name=receiptHashes,proto3" json:"receiptHashes,omitempty"`
	BlockType     BlockType        `protobuf:"varint,7,opt,name=blockType,proto3,enum=blockpb.BlockType" json:"blockType,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_block_pb_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_core_block_pb_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_core_block_pb_block_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetHead() *BlockHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *Block) GetSign() *pb.Signature {
	if x != nil {
		return x.Sign
	}
	return nil
}

func (x *Block) GetTxs() []*pb1.Tx {
	if x != nil {
		return x.Txs
	}
	return nil
}

func (x *Block) GetReceipts() []*pb1.TxReceipt {
	if x != nil {
		return x.Receipts
	}
	return nil
}

func (x *Block) GetTxHashes() [][]byte {
	if x != nil {
		return x.TxHashes
	}
	return nil
}

func (x *Block) GetReceiptHashes() [][]byte {
	if x != nil {
		return x.ReceiptHashes
	}
	return nil
}

func (x *Block) GetBlockType() BlockType {
	if x != nil {
		return x.BlockType
	}
	return BlockType_NORMAL
}

var File_core_block_pb_block_proto protoreflect.FileDescriptor

var file_core_block_pb_block_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x70, 0x62, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x70, 0x62, 0x1a, 0x19, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x13, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x78, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x78, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65,
	0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c,
	0x74, 0x78, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x74, 0x78, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x30, 0x0a, 0x13, 0x74, 0x78, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x4d, 0x65, 0x72,
	0x6b, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x74,
	0x78, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x92, 0x02, 0x0a,
	0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x26, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x24,
	0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73,
	0x69, 0x67, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x04,
	0x73, 0x69, 0x67, 0x6e, 0x12, 0x1a, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x74, 0x78, 0x70, 0x62, 0x2e, 0x54, 0x78, 0x52, 0x03, 0x74, 0x78, 0x73,
	0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x78, 0x70, 0x62, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x08, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12,
	0x30, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x2a, 0x25, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x4e,
	0x4c, 0x59, 0x48, 0x41, 0x53, 0x48, 0x10, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x73, 0x74, 0x2d, 0x6f, 0x66, 0x66, 0x69,
	0x63, 0x69, 0x61, 0x6c, 0x2f, 0x67, 0x6f, 0x2d, 0x69, 0x6f, 0x73, 0x74, 0x2f, 0x76, 0x33, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_block_pb_block_proto_rawDescOnce sync.Once
	file_core_block_pb_block_proto_rawDescData = file_core_block_pb_block_proto_rawDesc
)

func file_core_block_pb_block_proto_rawDescGZIP() []byte {
	file_core_block_pb_block_proto_rawDescOnce.Do(func() {
		file_core_block_pb_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_block_pb_block_proto_rawDescData)
	})
	return file_core_block_pb_block_proto_rawDescData
}

var file_core_block_pb_block_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_block_pb_block_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_core_block_pb_block_proto_goTypes = []interface{}{
	(BlockType)(0),        // 0: blockpb.BlockType
	(*BlockHead)(nil),     // 1: blockpb.BlockHead
	(*Block)(nil),         // 2: blockpb.Block
	(*pb.Signature)(nil),  // 3: sigpb.Signature
	(*pb1.Tx)(nil),        // 4: txpb.Tx
	(*pb1.TxReceipt)(nil), // 5: txpb.TxReceipt
}
var file_core_block_pb_block_proto_depIdxs = []int32{
	1, // 0: blockpb.Block.head:type_name -> blockpb.BlockHead
	3, // 1: blockpb.Block.sign:type_name -> sigpb.Signature
	4, // 2: blockpb.Block.txs:type_name -> txpb.Tx
	5, // 3: blockpb.Block.receipts:type_name -> txpb.TxReceipt
	0, // 4: blockpb.Block.blockType:type_name -> blockpb.BlockType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_core_block_pb_block_proto_init() }
func file_core_block_pb_block_proto_init() {
	if File_core_block_pb_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_block_pb_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHead); i {
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
		file_core_block_pb_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
			RawDescriptor: file_core_block_pb_block_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_block_pb_block_proto_goTypes,
		DependencyIndexes: file_core_block_pb_block_proto_depIdxs,
		EnumInfos:         file_core_block_pb_block_proto_enumTypes,
		MessageInfos:      file_core_block_pb_block_proto_msgTypes,
	}.Build()
	File_core_block_pb_block_proto = out.File
	file_core_block_pb_block_proto_rawDesc = nil
	file_core_block_pb_block_proto_goTypes = nil
	file_core_block_pb_block_proto_depIdxs = nil
}
