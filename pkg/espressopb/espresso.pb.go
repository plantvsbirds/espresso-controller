// Code generated by protoc-gen-go. DO NOT EDIT.
// source: espresso.proto

package espressopb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TemperatureSample struct {
	Value                float32              `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	ObservedAt           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=observed_at,json=observedAt,proto3" json:"observed_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TemperatureSample) Reset()         { *m = TemperatureSample{} }
func (m *TemperatureSample) String() string { return proto.CompactTextString(m) }
func (*TemperatureSample) ProtoMessage()    {}
func (*TemperatureSample) Descriptor() ([]byte, []int) {
	return fileDescriptor_445399412d1702d2, []int{0}
}

func (m *TemperatureSample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemperatureSample.Unmarshal(m, b)
}
func (m *TemperatureSample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemperatureSample.Marshal(b, m, deterministic)
}
func (m *TemperatureSample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemperatureSample.Merge(m, src)
}
func (m *TemperatureSample) XXX_Size() int {
	return xxx_messageInfo_TemperatureSample.Size(m)
}
func (m *TemperatureSample) XXX_DiscardUnknown() {
	xxx_messageInfo_TemperatureSample.DiscardUnknown(m)
}

var xxx_messageInfo_TemperatureSample proto.InternalMessageInfo

func (m *TemperatureSample) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TemperatureSample) GetObservedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ObservedAt
	}
	return nil
}

type TemperatureHistory struct {
	Samples              []*TemperatureSample `protobuf:"bytes,1,rep,name=samples,proto3" json:"samples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TemperatureHistory) Reset()         { *m = TemperatureHistory{} }
func (m *TemperatureHistory) String() string { return proto.CompactTextString(m) }
func (*TemperatureHistory) ProtoMessage()    {}
func (*TemperatureHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_445399412d1702d2, []int{1}
}

func (m *TemperatureHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemperatureHistory.Unmarshal(m, b)
}
func (m *TemperatureHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemperatureHistory.Marshal(b, m, deterministic)
}
func (m *TemperatureHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemperatureHistory.Merge(m, src)
}
func (m *TemperatureHistory) XXX_Size() int {
	return xxx_messageInfo_TemperatureHistory.Size(m)
}
func (m *TemperatureHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_TemperatureHistory.DiscardUnknown(m)
}

var xxx_messageInfo_TemperatureHistory proto.InternalMessageInfo

func (m *TemperatureHistory) GetSamples() []*TemperatureSample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type TemperatureStreamRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemperatureStreamRequest) Reset()         { *m = TemperatureStreamRequest{} }
func (m *TemperatureStreamRequest) String() string { return proto.CompactTextString(m) }
func (*TemperatureStreamRequest) ProtoMessage()    {}
func (*TemperatureStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_445399412d1702d2, []int{2}
}

func (m *TemperatureStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemperatureStreamRequest.Unmarshal(m, b)
}
func (m *TemperatureStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemperatureStreamRequest.Marshal(b, m, deterministic)
}
func (m *TemperatureStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemperatureStreamRequest.Merge(m, src)
}
func (m *TemperatureStreamRequest) XXX_Size() int {
	return xxx_messageInfo_TemperatureStreamRequest.Size(m)
}
func (m *TemperatureStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TemperatureStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TemperatureStreamRequest proto.InternalMessageInfo

type TemperatureStreamResponse struct {
	// Types that are valid to be assigned to Data:
	//	*TemperatureStreamResponse_History
	//	*TemperatureStreamResponse_Sample
	Data                 isTemperatureStreamResponse_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *TemperatureStreamResponse) Reset()         { *m = TemperatureStreamResponse{} }
func (m *TemperatureStreamResponse) String() string { return proto.CompactTextString(m) }
func (*TemperatureStreamResponse) ProtoMessage()    {}
func (*TemperatureStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_445399412d1702d2, []int{3}
}

func (m *TemperatureStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemperatureStreamResponse.Unmarshal(m, b)
}
func (m *TemperatureStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemperatureStreamResponse.Marshal(b, m, deterministic)
}
func (m *TemperatureStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemperatureStreamResponse.Merge(m, src)
}
func (m *TemperatureStreamResponse) XXX_Size() int {
	return xxx_messageInfo_TemperatureStreamResponse.Size(m)
}
func (m *TemperatureStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TemperatureStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TemperatureStreamResponse proto.InternalMessageInfo

type isTemperatureStreamResponse_Data interface {
	isTemperatureStreamResponse_Data()
}

type TemperatureStreamResponse_History struct {
	History *TemperatureHistory `protobuf:"bytes,1,opt,name=history,proto3,oneof"`
}

type TemperatureStreamResponse_Sample struct {
	Sample *TemperatureSample `protobuf:"bytes,2,opt,name=sample,proto3,oneof"`
}

func (*TemperatureStreamResponse_History) isTemperatureStreamResponse_Data() {}

func (*TemperatureStreamResponse_Sample) isTemperatureStreamResponse_Data() {}

func (m *TemperatureStreamResponse) GetData() isTemperatureStreamResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TemperatureStreamResponse) GetHistory() *TemperatureHistory {
	if x, ok := m.GetData().(*TemperatureStreamResponse_History); ok {
		return x.History
	}
	return nil
}

func (m *TemperatureStreamResponse) GetSample() *TemperatureSample {
	if x, ok := m.GetData().(*TemperatureStreamResponse_Sample); ok {
		return x.Sample
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TemperatureStreamResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TemperatureStreamResponse_History)(nil),
		(*TemperatureStreamResponse_Sample)(nil),
	}
}

type GetConfigurationRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConfigurationRequest) Reset()         { *m = GetConfigurationRequest{} }
func (m *GetConfigurationRequest) String() string { return proto.CompactTextString(m) }
func (*GetConfigurationRequest) ProtoMessage()    {}
func (*GetConfigurationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_445399412d1702d2, []int{4}
}

func (m *GetConfigurationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigurationRequest.Unmarshal(m, b)
}
func (m *GetConfigurationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigurationRequest.Marshal(b, m, deterministic)
}
func (m *GetConfigurationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigurationRequest.Merge(m, src)
}
func (m *GetConfigurationRequest) XXX_Size() int {
	return xxx_messageInfo_GetConfigurationRequest.Size(m)
}
func (m *GetConfigurationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigurationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigurationRequest proto.InternalMessageInfo

type Configuration struct {
	Temperature          float32              `protobuf:"fixed32,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	P                    float32              `protobuf:"fixed32,2,opt,name=p,proto3" json:"p,omitempty"`
	I                    float32              `protobuf:"fixed32,3,opt,name=i,proto3" json:"i,omitempty"`
	D                    float32              `protobuf:"fixed32,4,opt,name=d,proto3" json:"d,omitempty"`
	SetAt                *timestamp.Timestamp `protobuf:"bytes,5,opt,name=set_at,json=setAt,proto3" json:"set_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_445399412d1702d2, []int{5}
}

func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetTemperature() float32 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

func (m *Configuration) GetP() float32 {
	if m != nil {
		return m.P
	}
	return 0
}

func (m *Configuration) GetI() float32 {
	if m != nil {
		return m.I
	}
	return 0
}

func (m *Configuration) GetD() float32 {
	if m != nil {
		return m.D
	}
	return 0
}

func (m *Configuration) GetSetAt() *timestamp.Timestamp {
	if m != nil {
		return m.SetAt
	}
	return nil
}

func init() {
	proto.RegisterType((*TemperatureSample)(nil), "espressopb.TemperatureSample")
	proto.RegisterType((*TemperatureHistory)(nil), "espressopb.TemperatureHistory")
	proto.RegisterType((*TemperatureStreamRequest)(nil), "espressopb.TemperatureStreamRequest")
	proto.RegisterType((*TemperatureStreamResponse)(nil), "espressopb.TemperatureStreamResponse")
	proto.RegisterType((*GetConfigurationRequest)(nil), "espressopb.GetConfigurationRequest")
	proto.RegisterType((*Configuration)(nil), "espressopb.Configuration")
}

func init() {
	proto.RegisterFile("espresso.proto", fileDescriptor_445399412d1702d2)
}

var fileDescriptor_445399412d1702d2 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xed, 0xba, 0x8d, 0x8b, 0xc6, 0x80, 0xda, 0x15, 0x12, 0x8e, 0x25, 0xc0, 0x32, 0x20, 0xe5,
	0xe4, 0x82, 0x39, 0x54, 0x82, 0x53, 0x8b, 0x10, 0xbe, 0x70, 0xd9, 0xf6, 0x8e, 0xd6, 0xf2, 0x24,
	0x58, 0xb2, 0xb3, 0xcb, 0xee, 0xb8, 0x12, 0xff, 0x80, 0x38, 0xf3, 0xb9, 0x28, 0x5e, 0x5b, 0x75,
	0x42, 0x9d, 0x1c, 0xdf, 0xce, 0xdb, 0x79, 0xef, 0xcd, 0x83, 0xa7, 0x68, 0xb5, 0x41, 0x6b, 0x55,
	0xaa, 0x8d, 0x22, 0xc5, 0x61, 0xc0, 0xba, 0x88, 0x5e, 0xad, 0x94, 0x5a, 0xd5, 0x78, 0xd1, 0x4d,
	0x8a, 0x76, 0x79, 0x41, 0x55, 0x83, 0x96, 0x64, 0xa3, 0x1d, 0x39, 0x59, 0xc2, 0xf9, 0x2d, 0x36,
	0x1a, 0x8d, 0xa4, 0xd6, 0xe0, 0x8d, 0x6c, 0x74, 0x8d, 0xfc, 0x19, 0xcc, 0xee, 0x64, 0xdd, 0x62,
	0xc8, 0x62, 0xb6, 0xf0, 0x84, 0x03, 0xfc, 0x13, 0x04, 0xaa, 0xb0, 0x68, 0xee, 0xb0, 0xfc, 0x2e,
	0x29, 0xf4, 0x62, 0xb6, 0x08, 0xb2, 0x28, 0x75, 0x0a, 0xe9, 0xa0, 0x90, 0xde, 0x0e, 0x0a, 0x02,
	0x06, 0xfa, 0x15, 0x25, 0xdf, 0x80, 0x8f, 0x74, 0xf2, 0xca, 0x92, 0x32, 0xbf, 0xf8, 0x25, 0x9c,
	0xda, 0x4e, 0xd2, 0x86, 0x2c, 0x3e, 0x5e, 0x04, 0xd9, 0x8b, 0xf4, 0xde, 0x7c, 0xfa, 0x9f, 0x31,
	0x31, 0xb0, 0x93, 0x08, 0xc2, 0xf1, 0x94, 0x0c, 0xca, 0x46, 0xe0, 0xcf, 0x16, 0x2d, 0x25, 0x7f,
	0x19, 0xcc, 0x1f, 0x18, 0x5a, 0xad, 0xd6, 0x16, 0xf9, 0x47, 0x38, 0xfd, 0xe1, 0xd4, 0xbb, 0x74,
	0x41, 0xf6, 0x72, 0x42, 0xb2, 0xf7, 0x98, 0x1f, 0x89, 0xe1, 0x03, 0xbf, 0x04, 0xdf, 0x19, 0xe8,
	0xc3, 0xef, 0x77, 0x9b, 0x1f, 0x89, 0x9e, 0x7e, 0xed, 0xc3, 0x49, 0x29, 0x49, 0x26, 0x73, 0x78,
	0xfe, 0x15, 0xe9, 0xb3, 0x5a, 0x2f, 0xab, 0x55, 0x6b, 0x24, 0x55, 0x6a, 0x3d, 0xb8, 0xfe, 0xc3,
	0xe0, 0xc9, 0xd6, 0x80, 0xc7, 0x10, 0xd0, 0xfd, 0xce, 0xbe, 0x8b, 0xf1, 0x13, 0x7f, 0x0c, 0x4c,
	0x77, 0x56, 0x3c, 0xc1, 0xf4, 0x06, 0x55, 0xe1, 0xb1, 0x43, 0xd5, 0x06, 0x95, 0xe1, 0x89, 0x43,
	0x25, 0x7f, 0x0f, 0xbe, 0x45, 0xda, 0xd4, 0x36, 0x3b, 0x58, 0xdb, 0xcc, 0x22, 0x5d, 0x51, 0xf6,
	0xdb, 0x83, 0x47, 0x5f, 0xfa, 0x78, 0xbc, 0x80, 0xf3, 0x6b, 0x55, 0xd5, 0x68, 0x46, 0x29, 0xf9,
	0x9b, 0xa9, 0xf8, 0xe3, 0x3a, 0xa2, 0xb7, 0x07, 0x58, 0xae, 0x97, 0x77, 0x8c, 0x0b, 0x38, 0xdb,
	0x3d, 0x0e, 0x7f, 0x3d, 0xfe, 0x3c, 0x71, 0xba, 0x68, 0x3e, 0x26, 0x6d, 0xff, 0xcf, 0xe1, 0xec,
	0x66, 0x77, 0xe7, 0x34, 0x7d, 0xcf, 0xa6, 0xc2, 0xef, 0x2e, 0xf5, 0xe1, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd9, 0x86, 0x01, 0x1c, 0x6e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EspressoClient is the client API for Espresso service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EspressoClient interface {
	BoilerTemperature(ctx context.Context, in *TemperatureStreamRequest, opts ...grpc.CallOption) (Espresso_BoilerTemperatureClient, error)
	GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*Configuration, error)
	SetConfiguration(ctx context.Context, in *Configuration, opts ...grpc.CallOption) (*Configuration, error)
}

type espressoClient struct {
	cc grpc.ClientConnInterface
}

func NewEspressoClient(cc grpc.ClientConnInterface) EspressoClient {
	return &espressoClient{cc}
}

func (c *espressoClient) BoilerTemperature(ctx context.Context, in *TemperatureStreamRequest, opts ...grpc.CallOption) (Espresso_BoilerTemperatureClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Espresso_serviceDesc.Streams[0], "/espressopb.Espresso/BoilerTemperature", opts...)
	if err != nil {
		return nil, err
	}
	x := &espressoBoilerTemperatureClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Espresso_BoilerTemperatureClient interface {
	Recv() (*TemperatureStreamResponse, error)
	grpc.ClientStream
}

type espressoBoilerTemperatureClient struct {
	grpc.ClientStream
}

func (x *espressoBoilerTemperatureClient) Recv() (*TemperatureStreamResponse, error) {
	m := new(TemperatureStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *espressoClient) GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*Configuration, error) {
	out := new(Configuration)
	err := c.cc.Invoke(ctx, "/espressopb.Espresso/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *espressoClient) SetConfiguration(ctx context.Context, in *Configuration, opts ...grpc.CallOption) (*Configuration, error) {
	out := new(Configuration)
	err := c.cc.Invoke(ctx, "/espressopb.Espresso/SetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EspressoServer is the server API for Espresso service.
type EspressoServer interface {
	BoilerTemperature(*TemperatureStreamRequest, Espresso_BoilerTemperatureServer) error
	GetConfiguration(context.Context, *GetConfigurationRequest) (*Configuration, error)
	SetConfiguration(context.Context, *Configuration) (*Configuration, error)
}

// UnimplementedEspressoServer can be embedded to have forward compatible implementations.
type UnimplementedEspressoServer struct {
}

func (*UnimplementedEspressoServer) BoilerTemperature(req *TemperatureStreamRequest, srv Espresso_BoilerTemperatureServer) error {
	return status.Errorf(codes.Unimplemented, "method BoilerTemperature not implemented")
}
func (*UnimplementedEspressoServer) GetConfiguration(ctx context.Context, req *GetConfigurationRequest) (*Configuration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (*UnimplementedEspressoServer) SetConfiguration(ctx context.Context, req *Configuration) (*Configuration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfiguration not implemented")
}

func RegisterEspressoServer(s *grpc.Server, srv EspressoServer) {
	s.RegisterService(&_Espresso_serviceDesc, srv)
}

func _Espresso_BoilerTemperature_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TemperatureStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EspressoServer).BoilerTemperature(m, &espressoBoilerTemperatureServer{stream})
}

type Espresso_BoilerTemperatureServer interface {
	Send(*TemperatureStreamResponse) error
	grpc.ServerStream
}

type espressoBoilerTemperatureServer struct {
	grpc.ServerStream
}

func (x *espressoBoilerTemperatureServer) Send(m *TemperatureStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Espresso_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EspressoServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/espressopb.Espresso/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EspressoServer).GetConfiguration(ctx, req.(*GetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Espresso_SetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Configuration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EspressoServer).SetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/espressopb.Espresso/SetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EspressoServer).SetConfiguration(ctx, req.(*Configuration))
	}
	return interceptor(ctx, in, info, handler)
}

var _Espresso_serviceDesc = grpc.ServiceDesc{
	ServiceName: "espressopb.Espresso",
	HandlerType: (*EspressoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfiguration",
			Handler:    _Espresso_GetConfiguration_Handler,
		},
		{
			MethodName: "SetConfiguration",
			Handler:    _Espresso_SetConfiguration_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BoilerTemperature",
			Handler:       _Espresso_BoilerTemperature_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "espresso.proto",
}
