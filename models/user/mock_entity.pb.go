// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user/mock_entity.proto

package user

import (
	commproto "braid-demo/models/commproto"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// -----------------  entity module -----------------\
type EntityBagModule struct {
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" bson:"id"`
	// key: Dictionary ID of the item
	// value: For a given dictionary ID, there may be multiple non-stackable item instances
	Bag map[int32]*commproto.ItemList `protobuf:"bytes,2,rep,name=Bag,proto3" json:"Bag,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *EntityBagModule) Reset()         { *m = EntityBagModule{} }
func (m *EntityBagModule) String() string { return proto.CompactTextString(m) }
func (*EntityBagModule) ProtoMessage()    {}
func (*EntityBagModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f44b31d74392c8b, []int{0}
}
func (m *EntityBagModule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntityBagModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntityBagModule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntityBagModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityBagModule.Merge(m, src)
}
func (m *EntityBagModule) XXX_Size() int {
	return m.Size()
}
func (m *EntityBagModule) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityBagModule.DiscardUnknown(m)
}

var xxx_messageInfo_EntityBagModule proto.InternalMessageInfo

func (m *EntityBagModule) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *EntityBagModule) GetBag() map[int32]*commproto.ItemList {
	if m != nil {
		return m.Bag
	}
	return nil
}

type EntityTimeInfoModule struct {
	ID         string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" bson:"id"`
	LoginTime  int64  `protobuf:"varint,2,opt,name=LoginTime,proto3" json:"LoginTime,omitempty" bson:"login_time"`
	SyncTime   int64  `protobuf:"varint,3,opt,name=SyncTime,proto3" json:"SyncTime,omitempty" bson:"sync_time"`
	CreateTime int64  `protobuf:"varint,4,opt,name=CreateTime,proto3" json:"CreateTime,omitempty" bson:"create_time"`
}

func (m *EntityTimeInfoModule) Reset()         { *m = EntityTimeInfoModule{} }
func (m *EntityTimeInfoModule) String() string { return proto.CompactTextString(m) }
func (*EntityTimeInfoModule) ProtoMessage()    {}
func (*EntityTimeInfoModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f44b31d74392c8b, []int{1}
}
func (m *EntityTimeInfoModule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntityTimeInfoModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntityTimeInfoModule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntityTimeInfoModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityTimeInfoModule.Merge(m, src)
}
func (m *EntityTimeInfoModule) XXX_Size() int {
	return m.Size()
}
func (m *EntityTimeInfoModule) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityTimeInfoModule.DiscardUnknown(m)
}

var xxx_messageInfo_EntityTimeInfoModule proto.InternalMessageInfo

func (m *EntityTimeInfoModule) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *EntityTimeInfoModule) GetLoginTime() int64 {
	if m != nil {
		return m.LoginTime
	}
	return 0
}

func (m *EntityTimeInfoModule) GetSyncTime() int64 {
	if m != nil {
		return m.SyncTime
	}
	return 0
}

func (m *EntityTimeInfoModule) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type AirshipProduction struct {
	ProductionID string `protobuf:"bytes,1,opt,name=ProductionID,proto3" json:"ProductionID,omitempty"`
	EndTime      int64  `protobuf:"varint,3,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (m *AirshipProduction) Reset()         { *m = AirshipProduction{} }
func (m *AirshipProduction) String() string { return proto.CompactTextString(m) }
func (*AirshipProduction) ProtoMessage()    {}
func (*AirshipProduction) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f44b31d74392c8b, []int{2}
}
func (m *AirshipProduction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AirshipProduction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AirshipProduction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AirshipProduction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AirshipProduction.Merge(m, src)
}
func (m *AirshipProduction) XXX_Size() int {
	return m.Size()
}
func (m *AirshipProduction) XXX_DiscardUnknown() {
	xxx_messageInfo_AirshipProduction.DiscardUnknown(m)
}

var xxx_messageInfo_AirshipProduction proto.InternalMessageInfo

func (m *AirshipProduction) GetProductionID() string {
	if m != nil {
		return m.ProductionID
	}
	return ""
}

func (m *AirshipProduction) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type EntityAirshipModule struct {
	ID         string             `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" bson:"id"`
	Production *AirshipProduction `protobuf:"bytes,2,opt,name=Production,proto3" json:"Production,omitempty"`
}

func (m *EntityAirshipModule) Reset()         { *m = EntityAirshipModule{} }
func (m *EntityAirshipModule) String() string { return proto.CompactTextString(m) }
func (*EntityAirshipModule) ProtoMessage()    {}
func (*EntityAirshipModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f44b31d74392c8b, []int{3}
}
func (m *EntityAirshipModule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntityAirshipModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntityAirshipModule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntityAirshipModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityAirshipModule.Merge(m, src)
}
func (m *EntityAirshipModule) XXX_Size() int {
	return m.Size()
}
func (m *EntityAirshipModule) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityAirshipModule.DiscardUnknown(m)
}

var xxx_messageInfo_EntityAirshipModule proto.InternalMessageInfo

func (m *EntityAirshipModule) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *EntityAirshipModule) GetProduction() *AirshipProduction {
	if m != nil {
		return m.Production
	}
	return nil
}

// -----------------  entity -----------------
type EntityUserModule struct {
	ID           string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" bson:"id"`
	OpenID       string   `protobuf:"bytes,2,opt,name=OpenID,proto3" json:"OpenID,omitempty" bson:"open_id"`
	Token        string   `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty" bson:"token"`
	CurChannel   string   `protobuf:"bytes,4,opt,name=CurChannel,proto3" json:"CurChannel,omitempty" bson:"cur_channel"`
	ChatChannels []string `protobuf:"bytes,5,rep,name=ChatChannels,proto3" json:"ChatChannels,omitempty" bson:"chat_channels"`
}

func (m *EntityUserModule) Reset()         { *m = EntityUserModule{} }
func (m *EntityUserModule) String() string { return proto.CompactTextString(m) }
func (*EntityUserModule) ProtoMessage()    {}
func (*EntityUserModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f44b31d74392c8b, []int{4}
}
func (m *EntityUserModule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntityUserModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntityUserModule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntityUserModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityUserModule.Merge(m, src)
}
func (m *EntityUserModule) XXX_Size() int {
	return m.Size()
}
func (m *EntityUserModule) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityUserModule.DiscardUnknown(m)
}

var xxx_messageInfo_EntityUserModule proto.InternalMessageInfo

func (m *EntityUserModule) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *EntityUserModule) GetOpenID() string {
	if m != nil {
		return m.OpenID
	}
	return ""
}

func (m *EntityUserModule) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *EntityUserModule) GetCurChannel() string {
	if m != nil {
		return m.CurChannel
	}
	return ""
}

func (m *EntityUserModule) GetChatChannels() []string {
	if m != nil {
		return m.ChatChannels
	}
	return nil
}

func init() {
	proto.RegisterType((*EntityBagModule)(nil), "user.EntityBagModule")
	proto.RegisterMapType((map[int32]*commproto.ItemList)(nil), "user.EntityBagModule.BagEntry")
	proto.RegisterType((*EntityTimeInfoModule)(nil), "user.EntityTimeInfoModule")
	proto.RegisterType((*AirshipProduction)(nil), "user.AirshipProduction")
	proto.RegisterType((*EntityAirshipModule)(nil), "user.EntityAirshipModule")
	proto.RegisterType((*EntityUserModule)(nil), "user.EntityUserModule")
}

func init() { proto.RegisterFile("user/mock_entity.proto", fileDescriptor_6f44b31d74392c8b) }

var fileDescriptor_6f44b31d74392c8b = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0xd3, 0x4a,
	0x14, 0xc6, 0x6b, 0xbb, 0xe9, 0xad, 0x4f, 0x7b, 0x69, 0x3a, 0x0d, 0xc5, 0x8a, 0x84, 0x1d, 0xcd,
	0x02, 0x05, 0x16, 0x4e, 0xd5, 0x4a, 0x80, 0x2a, 0x36, 0x38, 0xcd, 0x22, 0xa2, 0x08, 0x30, 0x65,
	0x1d, 0xb9, 0xce, 0xe0, 0x58, 0x89, 0x67, 0x22, 0x7b, 0x8c, 0x94, 0xb7, 0xe0, 0x4d, 0x78, 0x0d,
	0x96, 0x5d, 0xb0, 0x60, 0x65, 0xa1, 0xe4, 0x0d, 0xbc, 0x65, 0x83, 0x66, 0xc6, 0x49, 0x1a, 0xd8,
	0x64, 0x63, 0xcd, 0x9c, 0xf3, 0xfb, 0xce, 0x9f, 0xf9, 0x0c, 0xa7, 0x79, 0x46, 0xd2, 0x4e, 0xc2,
	0xc2, 0xf1, 0x80, 0x50, 0x1e, 0xf3, 0x99, 0x3b, 0x4d, 0x19, 0x67, 0x68, 0x57, 0xc4, 0x9b, 0x47,
	0x11, 0x8b, 0x58, 0x47, 0x7c, 0x54, 0xb8, 0xd9, 0x08, 0x59, 0x92, 0xc8, 0x63, 0x47, 0x9c, 0x54,
	0x14, 0x7f, 0xd3, 0xe0, 0xa8, 0x27, 0xd5, 0x5e, 0x10, 0xbd, 0x65, 0xc3, 0x7c, 0x42, 0xd0, 0x63,
	0xd0, 0xfb, 0x57, 0x96, 0xd6, 0xd2, 0xda, 0xa6, 0xf7, 0x7f, 0x59, 0x38, 0xe6, 0x6d, 0xc6, 0xe8,
	0x25, 0x8e, 0x87, 0xd8, 0xd7, 0xfb, 0x57, 0xe8, 0x0c, 0x0c, 0x2f, 0x88, 0x2c, 0xbd, 0x65, 0xb4,
	0x0f, 0xce, 0x6d, 0x57, 0x74, 0x73, 0xff, 0x2a, 0xe1, 0x7a, 0x41, 0xd4, 0xa3, 0x3c, 0x9d, 0xf9,
	0x02, 0x6d, 0xbe, 0x81, 0xfd, 0x65, 0x00, 0xd5, 0xc1, 0x18, 0x93, 0x99, 0xac, 0x5e, 0xf3, 0xc5,
	0x11, 0x3d, 0x85, 0xda, 0x97, 0x60, 0x92, 0x13, 0x4b, 0x6f, 0x69, 0xed, 0x83, 0xf3, 0x13, 0x77,
	0x35, 0xa8, 0xdb, 0xe7, 0x24, 0xb9, 0x8e, 0x33, 0xee, 0x2b, 0xe2, 0x52, 0x7f, 0xa9, 0xe1, 0x1f,
	0x1a, 0x34, 0x54, 0xbb, 0x9b, 0x38, 0x21, 0x7d, 0xfa, 0x99, 0x6d, 0x37, 0xf6, 0x05, 0x98, 0xd7,
	0x2c, 0x8a, 0xa9, 0x50, 0xc9, 0x56, 0x86, 0xf7, 0xb0, 0x2c, 0x9c, 0x63, 0x45, 0x4d, 0x44, 0x6a,
	0xc0, 0xe3, 0x84, 0x60, 0x7f, 0xcd, 0xa1, 0x33, 0xd8, 0xff, 0x38, 0xa3, 0xa1, 0xd4, 0x18, 0x52,
	0xd3, 0x28, 0x0b, 0xa7, 0xae, 0x34, 0xd9, 0x8c, 0x86, 0x95, 0x64, 0x45, 0xa1, 0xe7, 0x00, 0xdd,
	0x94, 0x04, 0x9c, 0x48, 0xcd, 0xae, 0xd4, 0x9c, 0x96, 0x85, 0x83, 0x94, 0x26, 0x94, 0xb9, 0x4a,
	0x75, 0x8f, 0xc4, 0x1f, 0xe0, 0xf8, 0x75, 0x9c, 0x66, 0xa3, 0x78, 0xfa, 0x3e, 0x65, 0xc3, 0x3c,
	0xe4, 0x31, 0xa3, 0x08, 0xc3, 0xe1, 0xfa, 0xb6, 0x5c, 0xce, 0xdf, 0x88, 0x21, 0x0b, 0xfe, 0xeb,
	0xd1, 0xe1, 0x7a, 0x42, 0x7f, 0x79, 0xc5, 0x09, 0x9c, 0xa8, 0x87, 0xaa, 0x0a, 0x6f, 0xf7, 0x4e,
	0x2f, 0x00, 0xd6, 0xf5, 0x2b, 0x4f, 0x1e, 0x29, 0x97, 0xff, 0x19, 0xd0, 0xbf, 0x87, 0xe2, 0xdf,
	0x1a, 0xd4, 0x55, 0xbf, 0x4f, 0x19, 0x49, 0xb7, 0x6b, 0xf6, 0x0c, 0xf6, 0xde, 0x4d, 0x89, 0x58,
	0x4d, 0x97, 0x08, 0x2a, 0x0b, 0xe7, 0x81, 0x42, 0xd8, 0x94, 0xd0, 0x81, 0xe0, 0x2a, 0x02, 0x3d,
	0x81, 0xda, 0x0d, 0x1b, 0x13, 0x2a, 0xd7, 0x34, 0xbd, 0x7a, 0x59, 0x38, 0x87, 0x0a, 0xe5, 0x22,
	0x8c, 0x7d, 0x95, 0x96, 0x0e, 0xe4, 0x69, 0x77, 0x14, 0x50, 0x4a, 0x26, 0xd2, 0x01, 0x73, 0xc3,
	0x81, 0x3c, 0x1d, 0x84, 0x2a, 0x29, 0x1c, 0x58, 0x91, 0xe8, 0x15, 0x1c, 0x76, 0x47, 0x01, 0xaf,
	0xae, 0x99, 0x55, 0x6b, 0x19, 0x6d, 0xd3, 0xb3, 0xca, 0xc2, 0x69, 0x54, 0xca, 0x51, 0xc0, 0x97,
	0xd2, 0x0c, 0xfb, 0x1b, 0xb4, 0x67, 0x7d, 0x9f, 0xdb, 0xda, 0xdd, 0xdc, 0xd6, 0x7e, 0xcd, 0x6d,
	0xed, 0xeb, 0xc2, 0xde, 0xb9, 0x5b, 0xd8, 0x3b, 0x3f, 0x17, 0xf6, 0xce, 0xed, 0x9e, 0xfc, 0x97,
	0x2f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xe4, 0x75, 0xe7, 0xb0, 0x03, 0x00, 0x00,
}

func (m *EntityBagModule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityBagModule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntityBagModule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bag) > 0 {
		for k := range m.Bag {
			v := m.Bag[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintMockEntity(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintMockEntity(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintMockEntity(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintMockEntity(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EntityTimeInfoModule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityTimeInfoModule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntityTimeInfoModule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreateTime != 0 {
		i = encodeVarintMockEntity(dAtA, i, uint64(m.CreateTime))
		i--
		dAtA[i] = 0x20
	}
	if m.SyncTime != 0 {
		i = encodeVarintMockEntity(dAtA, i, uint64(m.SyncTime))
		i--
		dAtA[i] = 0x18
	}
	if m.LoginTime != 0 {
		i = encodeVarintMockEntity(dAtA, i, uint64(m.LoginTime))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintMockEntity(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AirshipProduction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AirshipProduction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AirshipProduction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndTime != 0 {
		i = encodeVarintMockEntity(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ProductionID) > 0 {
		i -= len(m.ProductionID)
		copy(dAtA[i:], m.ProductionID)
		i = encodeVarintMockEntity(dAtA, i, uint64(len(m.ProductionID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EntityAirshipModule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityAirshipModule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntityAirshipModule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Production != nil {
		{
			size, err := m.Production.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMockEntity(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintMockEntity(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EntityUserModule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityUserModule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntityUserModule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChatChannels) > 0 {
		for iNdEx := len(m.ChatChannels) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ChatChannels[iNdEx])
			copy(dAtA[i:], m.ChatChannels[iNdEx])
			i = encodeVarintMockEntity(dAtA, i, uint64(len(m.ChatChannels[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.CurChannel) > 0 {
		i -= len(m.CurChannel)
		copy(dAtA[i:], m.CurChannel)
		i = encodeVarintMockEntity(dAtA, i, uint64(len(m.CurChannel)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintMockEntity(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OpenID) > 0 {
		i -= len(m.OpenID)
		copy(dAtA[i:], m.OpenID)
		i = encodeVarintMockEntity(dAtA, i, uint64(len(m.OpenID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintMockEntity(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMockEntity(dAtA []byte, offset int, v uint64) int {
	offset -= sovMockEntity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EntityBagModule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovMockEntity(uint64(l))
	}
	if len(m.Bag) > 0 {
		for k, v := range m.Bag {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovMockEntity(uint64(l))
			}
			mapEntrySize := 1 + sovMockEntity(uint64(k)) + l
			n += mapEntrySize + 1 + sovMockEntity(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *EntityTimeInfoModule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovMockEntity(uint64(l))
	}
	if m.LoginTime != 0 {
		n += 1 + sovMockEntity(uint64(m.LoginTime))
	}
	if m.SyncTime != 0 {
		n += 1 + sovMockEntity(uint64(m.SyncTime))
	}
	if m.CreateTime != 0 {
		n += 1 + sovMockEntity(uint64(m.CreateTime))
	}
	return n
}

func (m *AirshipProduction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProductionID)
	if l > 0 {
		n += 1 + l + sovMockEntity(uint64(l))
	}
	if m.EndTime != 0 {
		n += 1 + sovMockEntity(uint64(m.EndTime))
	}
	return n
}

func (m *EntityAirshipModule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovMockEntity(uint64(l))
	}
	if m.Production != nil {
		l = m.Production.Size()
		n += 1 + l + sovMockEntity(uint64(l))
	}
	return n
}

func (m *EntityUserModule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovMockEntity(uint64(l))
	}
	l = len(m.OpenID)
	if l > 0 {
		n += 1 + l + sovMockEntity(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovMockEntity(uint64(l))
	}
	l = len(m.CurChannel)
	if l > 0 {
		n += 1 + l + sovMockEntity(uint64(l))
	}
	if len(m.ChatChannels) > 0 {
		for _, s := range m.ChatChannels {
			l = len(s)
			n += 1 + l + sovMockEntity(uint64(l))
		}
	}
	return n
}

func sovMockEntity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMockEntity(x uint64) (n int) {
	return sovMockEntity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EntityBagModule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMockEntity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EntityBagModule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntityBagModule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bag", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMockEntity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMockEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bag == nil {
				m.Bag = make(map[int32]*commproto.ItemList)
			}
			var mapkey int32
			var mapvalue *commproto.ItemList
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMockEntity
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMockEntity
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMockEntity
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthMockEntity
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthMockEntity
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &commproto.ItemList{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMockEntity(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthMockEntity
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Bag[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMockEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMockEntity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EntityTimeInfoModule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMockEntity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EntityTimeInfoModule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntityTimeInfoModule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoginTime", wireType)
			}
			m.LoginTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LoginTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SyncTime", wireType)
			}
			m.SyncTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SyncTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTime", wireType)
			}
			m.CreateTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMockEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMockEntity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AirshipProduction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMockEntity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AirshipProduction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AirshipProduction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProductionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMockEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMockEntity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EntityAirshipModule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMockEntity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EntityAirshipModule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntityAirshipModule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Production", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMockEntity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMockEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Production == nil {
				m.Production = &AirshipProduction{}
			}
			if err := m.Production.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMockEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMockEntity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EntityUserModule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMockEntity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EntityUserModule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntityUserModule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OpenID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChatChannels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMockEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMockEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChatChannels = append(m.ChatChannels, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMockEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMockEntity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMockEntity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMockEntity
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMockEntity
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMockEntity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMockEntity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMockEntity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMockEntity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMockEntity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMockEntity = fmt.Errorf("proto: unexpected end of group")
)
