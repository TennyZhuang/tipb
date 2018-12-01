// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pump.proto

package binlog

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WriteBinlogReq struct {
	// The identifier of tidb-cluster, which is given at tidb startup.
	// Must specify the clusterID for each binlog to write.
	ClusterID uint64 `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	// Payload bytes can be decoded back to binlog struct by the protobuf.
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteBinlogReq) Reset()         { *m = WriteBinlogReq{} }
func (m *WriteBinlogReq) String() string { return proto.CompactTextString(m) }
func (*WriteBinlogReq) ProtoMessage()    {}
func (*WriteBinlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_pump_8952f341c4b311c3, []int{0}
}
func (m *WriteBinlogReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WriteBinlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WriteBinlogReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *WriteBinlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteBinlogReq.Merge(dst, src)
}
func (m *WriteBinlogReq) XXX_Size() int {
	return m.Size()
}
func (m *WriteBinlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteBinlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_WriteBinlogReq proto.InternalMessageInfo

func (m *WriteBinlogReq) GetClusterID() uint64 {
	if m != nil {
		return m.ClusterID
	}
	return 0
}

func (m *WriteBinlogReq) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type WriteBinlogResp struct {
	// An empty errmsg returned means a successful write.
	// Otherwise return the error description.
	Errmsg               string   `protobuf:"bytes,1,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteBinlogResp) Reset()         { *m = WriteBinlogResp{} }
func (m *WriteBinlogResp) String() string { return proto.CompactTextString(m) }
func (*WriteBinlogResp) ProtoMessage()    {}
func (*WriteBinlogResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_pump_8952f341c4b311c3, []int{1}
}
func (m *WriteBinlogResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WriteBinlogResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WriteBinlogResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *WriteBinlogResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteBinlogResp.Merge(dst, src)
}
func (m *WriteBinlogResp) XXX_Size() int {
	return m.Size()
}
func (m *WriteBinlogResp) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteBinlogResp.DiscardUnknown(m)
}

var xxx_messageInfo_WriteBinlogResp proto.InternalMessageInfo

func (m *WriteBinlogResp) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

type PullBinlogReq struct {
	// Specifies which clusterID of binlog to pull.
	ClusterID uint64 `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	// The position from which the binlog will be sent.
	StartFrom            Pos      `protobuf:"bytes,2,opt,name=startFrom" json:"startFrom"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullBinlogReq) Reset()         { *m = PullBinlogReq{} }
func (m *PullBinlogReq) String() string { return proto.CompactTextString(m) }
func (*PullBinlogReq) ProtoMessage()    {}
func (*PullBinlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_pump_8952f341c4b311c3, []int{2}
}
func (m *PullBinlogReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PullBinlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PullBinlogReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PullBinlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullBinlogReq.Merge(dst, src)
}
func (m *PullBinlogReq) XXX_Size() int {
	return m.Size()
}
func (m *PullBinlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PullBinlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_PullBinlogReq proto.InternalMessageInfo

func (m *PullBinlogReq) GetClusterID() uint64 {
	if m != nil {
		return m.ClusterID
	}
	return 0
}

func (m *PullBinlogReq) GetStartFrom() Pos {
	if m != nil {
		return m.StartFrom
	}
	return Pos{}
}

type PullBinlogResp struct {
	// The binlog entity that send in a stream
	Entity               Entity   `protobuf:"bytes,1,opt,name=entity" json:"entity"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullBinlogResp) Reset()         { *m = PullBinlogResp{} }
func (m *PullBinlogResp) String() string { return proto.CompactTextString(m) }
func (*PullBinlogResp) ProtoMessage()    {}
func (*PullBinlogResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_pump_8952f341c4b311c3, []int{3}
}
func (m *PullBinlogResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PullBinlogResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PullBinlogResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PullBinlogResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullBinlogResp.Merge(dst, src)
}
func (m *PullBinlogResp) XXX_Size() int {
	return m.Size()
}
func (m *PullBinlogResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PullBinlogResp.DiscardUnknown(m)
}

var xxx_messageInfo_PullBinlogResp proto.InternalMessageInfo

func (m *PullBinlogResp) GetEntity() Entity {
	if m != nil {
		return m.Entity
	}
	return Entity{}
}

// Binlogs are stored in a number of sequential files in a directory.
// The Pos describes the position of a binlog.
type Pos struct {
	// The suffix of binlog file, like .000001 .000002
	Suffix uint64 `protobuf:"varint,1,opt,name=suffix,proto3" json:"suffix,omitempty"`
	// The binlog offset in a file.
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pos) Reset()         { *m = Pos{} }
func (m *Pos) String() string { return proto.CompactTextString(m) }
func (*Pos) ProtoMessage()    {}
func (*Pos) Descriptor() ([]byte, []int) {
	return fileDescriptor_pump_8952f341c4b311c3, []int{4}
}
func (m *Pos) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pos.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Pos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pos.Merge(dst, src)
}
func (m *Pos) XXX_Size() int {
	return m.Size()
}
func (m *Pos) XXX_DiscardUnknown() {
	xxx_messageInfo_Pos.DiscardUnknown(m)
}

var xxx_messageInfo_Pos proto.InternalMessageInfo

func (m *Pos) GetSuffix() uint64 {
	if m != nil {
		return m.Suffix
	}
	return 0
}

func (m *Pos) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type Entity struct {
	// The position of the binlog entity.
	Pos Pos `protobuf:"bytes,1,opt,name=pos" json:"pos"`
	// The payload of binlog entity.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// checksum of binlog payload.
	Checksum             []byte   `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_pump_8952f341c4b311c3, []int{5}
}
func (m *Entity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(dst, src)
}
func (m *Entity) XXX_Size() int {
	return m.Size()
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetPos() Pos {
	if m != nil {
		return m.Pos
	}
	return Pos{}
}

func (m *Entity) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Entity) GetChecksum() []byte {
	if m != nil {
		return m.Checksum
	}
	return nil
}

func init() {
	proto.RegisterType((*WriteBinlogReq)(nil), "binlog.WriteBinlogReq")
	proto.RegisterType((*WriteBinlogResp)(nil), "binlog.WriteBinlogResp")
	proto.RegisterType((*PullBinlogReq)(nil), "binlog.PullBinlogReq")
	proto.RegisterType((*PullBinlogResp)(nil), "binlog.PullBinlogResp")
	proto.RegisterType((*Pos)(nil), "binlog.Pos")
	proto.RegisterType((*Entity)(nil), "binlog.Entity")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PumpClient is the client API for Pump service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PumpClient interface {
	// Writes a binlog to the local file on the pump machine.
	// A response with an empty errmsg is returned if the binlog is written successfully.
	WriteBinlog(ctx context.Context, in *WriteBinlogReq, opts ...grpc.CallOption) (*WriteBinlogResp, error)
	// Sends binlog stream from a given location.
	PullBinlogs(ctx context.Context, in *PullBinlogReq, opts ...grpc.CallOption) (Pump_PullBinlogsClient, error)
}

type pumpClient struct {
	cc *grpc.ClientConn
}

func NewPumpClient(cc *grpc.ClientConn) PumpClient {
	return &pumpClient{cc}
}

func (c *pumpClient) WriteBinlog(ctx context.Context, in *WriteBinlogReq, opts ...grpc.CallOption) (*WriteBinlogResp, error) {
	out := new(WriteBinlogResp)
	err := c.cc.Invoke(ctx, "/binlog.Pump/WriteBinlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pumpClient) PullBinlogs(ctx context.Context, in *PullBinlogReq, opts ...grpc.CallOption) (Pump_PullBinlogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Pump_serviceDesc.Streams[0], "/binlog.Pump/PullBinlogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &pumpPullBinlogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Pump_PullBinlogsClient interface {
	Recv() (*PullBinlogResp, error)
	grpc.ClientStream
}

type pumpPullBinlogsClient struct {
	grpc.ClientStream
}

func (x *pumpPullBinlogsClient) Recv() (*PullBinlogResp, error) {
	m := new(PullBinlogResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PumpServer is the server API for Pump service.
type PumpServer interface {
	// Writes a binlog to the local file on the pump machine.
	// A response with an empty errmsg is returned if the binlog is written successfully.
	WriteBinlog(context.Context, *WriteBinlogReq) (*WriteBinlogResp, error)
	// Sends binlog stream from a given location.
	PullBinlogs(*PullBinlogReq, Pump_PullBinlogsServer) error
}

func RegisterPumpServer(s *grpc.Server, srv PumpServer) {
	s.RegisterService(&_Pump_serviceDesc, srv)
}

func _Pump_WriteBinlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteBinlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PumpServer).WriteBinlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binlog.Pump/WriteBinlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PumpServer).WriteBinlog(ctx, req.(*WriteBinlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pump_PullBinlogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullBinlogReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PumpServer).PullBinlogs(m, &pumpPullBinlogsServer{stream})
}

type Pump_PullBinlogsServer interface {
	Send(*PullBinlogResp) error
	grpc.ServerStream
}

type pumpPullBinlogsServer struct {
	grpc.ServerStream
}

func (x *pumpPullBinlogsServer) Send(m *PullBinlogResp) error {
	return x.ServerStream.SendMsg(m)
}

var _Pump_serviceDesc = grpc.ServiceDesc{
	ServiceName: "binlog.Pump",
	HandlerType: (*PumpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteBinlog",
			Handler:    _Pump_WriteBinlog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullBinlogs",
			Handler:       _Pump_PullBinlogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pump.proto",
}

func (m *WriteBinlogReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WriteBinlogReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ClusterID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPump(dAtA, i, uint64(m.ClusterID))
	}
	if len(m.Payload) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPump(dAtA, i, uint64(len(m.Payload)))
		i += copy(dAtA[i:], m.Payload)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *WriteBinlogResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WriteBinlogResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Errmsg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPump(dAtA, i, uint64(len(m.Errmsg)))
		i += copy(dAtA[i:], m.Errmsg)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PullBinlogReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PullBinlogReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ClusterID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPump(dAtA, i, uint64(m.ClusterID))
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintPump(dAtA, i, uint64(m.StartFrom.Size()))
	n1, err := m.StartFrom.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PullBinlogResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PullBinlogResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintPump(dAtA, i, uint64(m.Entity.Size()))
	n2, err := m.Entity.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Pos) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pos) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Suffix != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPump(dAtA, i, uint64(m.Suffix))
	}
	if m.Offset != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPump(dAtA, i, uint64(m.Offset))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Entity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entity) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintPump(dAtA, i, uint64(m.Pos.Size()))
	n3, err := m.Pos.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.Payload) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPump(dAtA, i, uint64(len(m.Payload)))
		i += copy(dAtA[i:], m.Payload)
	}
	if len(m.Checksum) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPump(dAtA, i, uint64(len(m.Checksum)))
		i += copy(dAtA[i:], m.Checksum)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPump(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *WriteBinlogReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClusterID != 0 {
		n += 1 + sovPump(uint64(m.ClusterID))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovPump(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WriteBinlogResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Errmsg)
	if l > 0 {
		n += 1 + l + sovPump(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PullBinlogReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClusterID != 0 {
		n += 1 + sovPump(uint64(m.ClusterID))
	}
	l = m.StartFrom.Size()
	n += 1 + l + sovPump(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PullBinlogResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Entity.Size()
	n += 1 + l + sovPump(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Pos) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Suffix != 0 {
		n += 1 + sovPump(uint64(m.Suffix))
	}
	if m.Offset != 0 {
		n += 1 + sovPump(uint64(m.Offset))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Entity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Pos.Size()
	n += 1 + l + sovPump(uint64(l))
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovPump(uint64(l))
	}
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovPump(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPump(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPump(x uint64) (n int) {
	return sovPump(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WriteBinlogReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPump
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WriteBinlogReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WriteBinlogReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			m.ClusterID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPump
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClusterID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPump
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPump
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPump(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPump
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WriteBinlogResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPump
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WriteBinlogResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WriteBinlogResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Errmsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPump
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPump
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Errmsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPump(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPump
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PullBinlogReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPump
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PullBinlogReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PullBinlogReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			m.ClusterID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPump
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClusterID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartFrom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPump
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPump
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StartFrom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPump(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPump
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PullBinlogResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPump
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PullBinlogResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PullBinlogResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPump
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPump
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Entity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPump(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPump
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Pos) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPump
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Pos: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pos: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suffix", wireType)
			}
			m.Suffix = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPump
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Suffix |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPump
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPump(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPump
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Entity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPump
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Entity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPump
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPump
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pos.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPump
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPump
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPump
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPump
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = append(m.Checksum[:0], dAtA[iNdEx:postIndex]...)
			if m.Checksum == nil {
				m.Checksum = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPump(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPump
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPump(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPump
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
					return 0, ErrIntOverflowPump
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPump
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPump
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPump
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPump(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPump = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPump   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pump.proto", fileDescriptor_pump_8952f341c4b311c3) }

var fileDescriptor_pump_8952f341c4b311c3 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0xed, 0x7c, 0x29, 0xf9, 0xec, 0x8d, 0x56, 0x19, 0xb4, 0x86, 0x20, 0xb1, 0x8c, 0x9b, 0x0a,
	0xd2, 0x4a, 0xc5, 0xad, 0x48, 0x51, 0xd1, 0x5d, 0xc8, 0xc6, 0x9d, 0xd0, 0xc6, 0x49, 0x0c, 0x26,
	0x9d, 0x71, 0x66, 0x02, 0xf6, 0x15, 0x7c, 0x02, 0x1f, 0xa9, 0x4b, 0x9f, 0x40, 0xa4, 0xbe, 0x88,
	0x64, 0x3a, 0xfd, 0x83, 0x0a, 0xee, 0x72, 0xce, 0xbd, 0xf7, 0x9c, 0x93, 0x7b, 0x07, 0x80, 0x17,
	0x39, 0x6f, 0x73, 0xc1, 0x14, 0xc3, 0xf6, 0x20, 0x1d, 0x66, 0x2c, 0xf1, 0x76, 0x13, 0x96, 0x30,
	0x4d, 0x75, 0xca, 0xaf, 0x69, 0x95, 0xdc, 0x42, 0xfd, 0x5e, 0xa4, 0x8a, 0xf6, 0x74, 0x53, 0x48,
	0x5f, 0xf0, 0x01, 0xd4, 0xa2, 0xac, 0x90, 0x8a, 0x8a, 0xbb, 0x2b, 0x17, 0x35, 0x51, 0xab, 0x1a,
	0x2e, 0x08, 0xec, 0xc2, 0x7f, 0xde, 0x1f, 0x65, 0xac, 0xff, 0xe8, 0xfe, 0x6b, 0xa2, 0xd6, 0x66,
	0x38, 0x83, 0xe4, 0x18, 0xb6, 0x57, 0x94, 0x24, 0xc7, 0x0d, 0xb0, 0xa9, 0x10, 0xb9, 0x4c, 0xb4,
	0x4e, 0x2d, 0x34, 0x88, 0x3c, 0xc0, 0x56, 0x50, 0x64, 0xd9, 0x5f, 0x3d, 0x3b, 0x50, 0x93, 0xaa,
	0x2f, 0xd4, 0x8d, 0x60, 0xb9, 0x76, 0x75, 0xba, 0x4e, 0x7b, 0xfa, 0x57, 0xed, 0x80, 0xc9, 0x5e,
	0x75, 0xfc, 0x79, 0x58, 0x09, 0x17, 0x3d, 0xe4, 0x02, 0xea, 0xcb, 0xfa, 0x92, 0xe3, 0x13, 0xb0,
	0xe9, 0x50, 0xa5, 0x6a, 0xa4, 0xd5, 0x9d, 0x6e, 0x7d, 0x36, 0x7f, 0xad, 0x59, 0x23, 0x61, 0x7a,
	0xc8, 0x39, 0x58, 0x01, 0x93, 0x65, 0x7c, 0x59, 0xc4, 0x71, 0xfa, 0x6a, 0x22, 0x19, 0x54, 0xf2,
	0x2c, 0x8e, 0x25, 0x55, 0x3a, 0x8c, 0x15, 0x1a, 0x44, 0x22, 0xb0, 0xa7, 0x72, 0xf8, 0x08, 0x2c,
	0xce, 0xa4, 0xf1, 0x5a, 0x93, 0xb5, 0xac, 0xfe, 0xbe, 0x4a, 0xec, 0xc1, 0x46, 0xf4, 0x44, 0xa3,
	0x67, 0x59, 0xe4, 0xae, 0xa5, 0x4b, 0x73, 0xdc, 0x7d, 0x43, 0x50, 0x0d, 0x8a, 0x9c, 0xe3, 0x4b,
	0x70, 0x96, 0xf6, 0x8d, 0x1b, 0x33, 0x97, 0xd5, 0x73, 0x7a, 0xfb, 0x6b, 0x79, 0xc9, 0x49, 0xa5,
	0x54, 0x58, 0xac, 0x49, 0xe2, 0xbd, 0x79, 0xce, 0xe5, 0xdb, 0x78, 0x8d, 0x75, 0x74, 0x39, 0x7f,
	0x8a, 0x7a, 0x3b, 0xe3, 0x89, 0x8f, 0x3e, 0x26, 0x3e, 0xfa, 0x9a, 0xf8, 0xe8, 0xfd, 0xdb, 0xaf,
	0x0c, 0x6c, 0xfd, 0xac, 0xce, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x2c, 0x86, 0x78, 0x82,
	0x02, 0x00, 0x00,
}
