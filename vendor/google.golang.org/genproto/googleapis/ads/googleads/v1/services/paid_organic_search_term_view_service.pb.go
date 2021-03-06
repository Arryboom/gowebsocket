// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/paid_organic_search_term_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Request message for
// [PaidOrganicSearchTermViewService.GetPaidOrganicSearchTermView][google.ads.googleads.v1.services.PaidOrganicSearchTermViewService.GetPaidOrganicSearchTermView].
type GetPaidOrganicSearchTermViewRequest struct {
	// The resource name of the paid organic search term view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPaidOrganicSearchTermViewRequest) Reset()         { *m = GetPaidOrganicSearchTermViewRequest{} }
func (m *GetPaidOrganicSearchTermViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetPaidOrganicSearchTermViewRequest) ProtoMessage()    {}
func (*GetPaidOrganicSearchTermViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccacc683251a5992, []int{0}
}

func (m *GetPaidOrganicSearchTermViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPaidOrganicSearchTermViewRequest.Unmarshal(m, b)
}
func (m *GetPaidOrganicSearchTermViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPaidOrganicSearchTermViewRequest.Marshal(b, m, deterministic)
}
func (m *GetPaidOrganicSearchTermViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPaidOrganicSearchTermViewRequest.Merge(m, src)
}
func (m *GetPaidOrganicSearchTermViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetPaidOrganicSearchTermViewRequest.Size(m)
}
func (m *GetPaidOrganicSearchTermViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPaidOrganicSearchTermViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPaidOrganicSearchTermViewRequest proto.InternalMessageInfo

func (m *GetPaidOrganicSearchTermViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPaidOrganicSearchTermViewRequest)(nil), "google.ads.googleads.v1.services.GetPaidOrganicSearchTermViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/paid_organic_search_term_view_service.proto", fileDescriptor_ccacc683251a5992)
}

var fileDescriptor_ccacc683251a5992 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x86, 0x49, 0x04, 0xc1, 0xa0, 0x97, 0x9c, 0xa4, 0xf4, 0x10, 0xda, 0x0a, 0xd2, 0xc3, 0x0c,
	0xd1, 0xdb, 0x58, 0x85, 0x14, 0x4b, 0x45, 0x44, 0x4b, 0x2b, 0x39, 0x48, 0x20, 0x8c, 0xc9, 0x47,
	0x1c, 0x68, 0x32, 0x71, 0xbe, 0x34, 0x3d, 0x88, 0x17, 0xcf, 0xe2, 0xc5, 0x7f, 0xe0, 0xd1, 0x9f,
	0xb2, 0xd7, 0xfd, 0x0b, 0x7b, 0xd9, 0xfd, 0x15, 0x4b, 0x32, 0x9d, 0xc0, 0xc2, 0x66, 0xb3, 0xb7,
	0x97, 0x99, 0x97, 0xe7, 0x9d, 0xef, 0xfd, 0xc6, 0xf9, 0x90, 0x49, 0x99, 0xed, 0x81, 0xf2, 0x14,
	0xa9, 0x96, 0x8d, 0xaa, 0x7d, 0x8a, 0xa0, 0x6a, 0x91, 0x00, 0xd2, 0x92, 0x8b, 0x34, 0x96, 0x2a,
	0xe3, 0x85, 0x48, 0x62, 0x04, 0xae, 0x92, 0x6f, 0x71, 0x05, 0x2a, 0x8f, 0x6b, 0x01, 0xc7, 0xf8,
	0x64, 0x23, 0xa5, 0x92, 0x95, 0x74, 0x3d, 0x8d, 0x20, 0x3c, 0x45, 0xd2, 0xd1, 0x48, 0xed, 0x13,
	0x43, 0x1b, 0xad, 0xfa, 0xf2, 0x14, 0xa0, 0x3c, 0xa8, 0xc1, 0x40, 0x1d, 0x34, 0x1a, 0x1b, 0x4c,
	0x29, 0x28, 0x2f, 0x0a, 0x59, 0xf1, 0x4a, 0xc8, 0x02, 0xf5, 0xed, 0xe4, 0xbd, 0x33, 0x5d, 0x43,
	0xb5, 0xe1, 0x22, 0xfd, 0xa4, 0x31, 0xbb, 0x96, 0xf2, 0x19, 0x54, 0x1e, 0x0a, 0x38, 0x6e, 0xe1,
	0xfb, 0x01, 0xb0, 0x72, 0xa7, 0xce, 0x13, 0x93, 0x1a, 0x17, 0x3c, 0x87, 0xa7, 0x96, 0x67, 0x3d,
	0x7f, 0xb4, 0x7d, 0x6c, 0x0e, 0x3f, 0xf2, 0x1c, 0x5e, 0xfc, 0xb1, 0x1d, 0xaf, 0x97, 0xb4, 0xd3,
	0x63, 0xb9, 0x97, 0x96, 0x33, 0xbe, 0x2b, 0xd1, 0x5d, 0x91, 0xa1, 0x66, 0xc8, 0x3d, 0x5e, 0x3c,
	0x5a, 0xf4, 0x62, 0xba, 0xfa, 0x48, 0x2f, 0x64, 0xf2, 0xf6, 0xd7, 0xf9, 0xc5, 0x5f, 0xfb, 0x8d,
	0xbb, 0x68, 0xfa, 0xfe, 0x71, 0x63, 0xf4, 0xd7, 0xc9, 0x01, 0x2b, 0x99, 0x83, 0x42, 0x3a, 0x6f,
	0x17, 0x70, 0x2b, 0x01, 0xe9, 0xfc, 0xe7, 0xf2, 0xb7, 0xed, 0xcc, 0x12, 0x99, 0x0f, 0x0e, 0xb4,
	0x7c, 0x36, 0x54, 0xdb, 0xa6, 0x59, 0xd6, 0xc6, 0xfa, 0xf2, 0xee, 0x84, 0xca, 0xe4, 0x9e, 0x17,
	0x19, 0x91, 0x2a, 0xa3, 0x19, 0x14, 0xed, 0x2a, 0xcd, 0x1f, 0x29, 0x05, 0xf6, 0x7f, 0xd1, 0x57,
	0x46, 0xfc, 0xb3, 0x1f, 0xac, 0x83, 0xe0, 0xbf, 0xed, 0xad, 0x35, 0x30, 0x48, 0x91, 0x68, 0xd9,
	0xa8, 0xd0, 0x27, 0xa7, 0x60, 0x3c, 0x33, 0x96, 0x28, 0x48, 0x31, 0xea, 0x2c, 0x51, 0xe8, 0x47,
	0xc6, 0x72, 0x65, 0xcf, 0xf4, 0x39, 0x63, 0x41, 0x8a, 0x8c, 0x75, 0x26, 0xc6, 0x42, 0x9f, 0x31,
	0x63, 0xfb, 0xfa, 0xb0, 0x7d, 0xe7, 0xcb, 0xeb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0x18, 0xd6,
	0x92, 0x49, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PaidOrganicSearchTermViewServiceClient is the client API for PaidOrganicSearchTermViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaidOrganicSearchTermViewServiceClient interface {
	// Returns the requested paid organic search term view in full detail.
	GetPaidOrganicSearchTermView(ctx context.Context, in *GetPaidOrganicSearchTermViewRequest, opts ...grpc.CallOption) (*resources.PaidOrganicSearchTermView, error)
}

type paidOrganicSearchTermViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewPaidOrganicSearchTermViewServiceClient(cc *grpc.ClientConn) PaidOrganicSearchTermViewServiceClient {
	return &paidOrganicSearchTermViewServiceClient{cc}
}

func (c *paidOrganicSearchTermViewServiceClient) GetPaidOrganicSearchTermView(ctx context.Context, in *GetPaidOrganicSearchTermViewRequest, opts ...grpc.CallOption) (*resources.PaidOrganicSearchTermView, error) {
	out := new(resources.PaidOrganicSearchTermView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.PaidOrganicSearchTermViewService/GetPaidOrganicSearchTermView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaidOrganicSearchTermViewServiceServer is the server API for PaidOrganicSearchTermViewService service.
type PaidOrganicSearchTermViewServiceServer interface {
	// Returns the requested paid organic search term view in full detail.
	GetPaidOrganicSearchTermView(context.Context, *GetPaidOrganicSearchTermViewRequest) (*resources.PaidOrganicSearchTermView, error)
}

func RegisterPaidOrganicSearchTermViewServiceServer(s *grpc.Server, srv PaidOrganicSearchTermViewServiceServer) {
	s.RegisterService(&_PaidOrganicSearchTermViewService_serviceDesc, srv)
}

func _PaidOrganicSearchTermViewService_GetPaidOrganicSearchTermView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaidOrganicSearchTermViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaidOrganicSearchTermViewServiceServer).GetPaidOrganicSearchTermView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.PaidOrganicSearchTermViewService/GetPaidOrganicSearchTermView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaidOrganicSearchTermViewServiceServer).GetPaidOrganicSearchTermView(ctx, req.(*GetPaidOrganicSearchTermViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaidOrganicSearchTermViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.PaidOrganicSearchTermViewService",
	HandlerType: (*PaidOrganicSearchTermViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPaidOrganicSearchTermView",
			Handler:    _PaidOrganicSearchTermViewService_GetPaidOrganicSearchTermView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/paid_organic_search_term_view_service.proto",
}
