// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/privval/service.proto

package privval

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("tendermint/privval/service.proto", fileDescriptor_7afe74f9f46d3dc9) }

var fileDescriptor_7afe74f9f46d3dc9 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0x13, 0x0b, 0xd1, 0x60, 0x21, 0x0b, 0x36, 0x57, 0x2c, 0x16, 0x2a, 0x28, 0x24, 0x0b,
	0x5a, 0x5b, 0x68, 0x23, 0x62, 0x13, 0x14, 0x4e, 0x11, 0x2c, 0x36, 0xc9, 0x70, 0x37, 0x90, 0xdb,
	0x59, 0x77, 0x27, 0x81, 0x7b, 0x0b, 0x9f, 0xc6, 0x67, 0xb0, 0xbc, 0xd2, 0x52, 0x92, 0x17, 0x11,
	0xcd, 0x85, 0x43, 0x3c, 0x63, 0x3b, 0xf3, 0x7d, 0xff, 0x5f, 0xfc, 0xd1, 0x3e, 0x83, 0x29, 0xc0,
	0xcd, 0xd0, 0xb0, 0xb2, 0x0e, 0xeb, 0x5a, 0x97, 0xca, 0x83, 0xab, 0x31, 0x87, 0xc4, 0x3a, 0x62,
	0x12, 0x7b, 0x1e, 0x70, 0x05, 0x25, 0x4b, 0x68, 0x24, 0xd7, 0x88, 0x3c, 0xb7, 0xe0, 0x3b, 0xed,
	0xf4, 0x75, 0x23, 0xda, 0x4d, 0x1d, 0xd6, 0x63, 0x5d, 0x62, 0xa1, 0x99, 0xdc, 0x45, 0x7a, 0x2d,
	0x1e, 0xa2, 0xed, 0x2b, 0xe0, 0xb4, 0xca, 0x6e, 0x60, 0x2e, 0x0e, 0x92, 0xb5, 0xc9, 0x49, 0xf7,
	0xbe, 0x85, 0xe7, 0x0a, 0x3c, 0x8f, 0x0e, 0xff, 0xa1, 0xbc, 0x25, 0xe3, 0x41, 0x3c, 0x45, 0x5b,
	0x77, 0x38, 0x31, 0x63, 0x62, 0x10, 0x47, 0x7f, 0x28, 0x3d, 0xd0, 0x47, 0x1f, 0x0f, 0x70, 0x50,
	0x74, 0xe4, 0x32, 0x1e, 0xa3, 0x9d, 0xaf, 0x6b, 0xea, 0xc8, 0x92, 0xd7, 0xa5, 0x38, 0x19, 0x50,
	0x7b, 0xa8, 0xaf, 0x89, 0x07, 0x6b, 0x56, 0x74, 0x57, 0x75, 0x79, 0xff, 0xd6, 0xc8, 0x70, 0xd1,
	0xc8, 0xf0, 0xa3, 0x91, 0xe1, 0x4b, 0x2b, 0x83, 0x45, 0x2b, 0x83, 0xf7, 0x56, 0x06, 0x8f, 0xe7,
	0x13, 0xe4, 0x69, 0x95, 0x25, 0x39, 0xcd, 0x94, 0x76, 0x18, 0x6b, 0x93, 0x4f, 0xc9, 0x29, 0x0f,
	0x18, 0xff, 0x18, 0x83, 0x98, 0xd4, 0xef, 0x75, 0xb2, 0xcd, 0xef, 0xcf, 0xd9, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xcf, 0x10, 0xb9, 0xf1, 0xf3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PrivValidatorAPIClient is the client API for PrivValidatorAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrivValidatorAPIClient interface {
	GetPubKey(ctx context.Context, in *PubKeyRequest, opts ...grpc.CallOption) (*PubKeyResponse, error)
	SignVote(ctx context.Context, in *SignVoteRequest, opts ...grpc.CallOption) (*SignedVoteResponse, error)
	SignProposal(ctx context.Context, in *SignProposalRequest, opts ...grpc.CallOption) (*SignedProposalResponse, error)
}

type privValidatorAPIClient struct {
	cc *grpc.ClientConn
}

func NewPrivValidatorAPIClient(cc *grpc.ClientConn) PrivValidatorAPIClient {
	return &privValidatorAPIClient{cc}
}

func (c *privValidatorAPIClient) GetPubKey(ctx context.Context, in *PubKeyRequest, opts ...grpc.CallOption) (*PubKeyResponse, error) {
	out := new(PubKeyResponse)
	err := c.cc.Invoke(ctx, "/seitendermint.privval.PrivValidatorAPI/GetPubKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privValidatorAPIClient) SignVote(ctx context.Context, in *SignVoteRequest, opts ...grpc.CallOption) (*SignedVoteResponse, error) {
	out := new(SignedVoteResponse)
	err := c.cc.Invoke(ctx, "/seitendermint.privval.PrivValidatorAPI/SignVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privValidatorAPIClient) SignProposal(ctx context.Context, in *SignProposalRequest, opts ...grpc.CallOption) (*SignedProposalResponse, error) {
	out := new(SignedProposalResponse)
	err := c.cc.Invoke(ctx, "/seitendermint.privval.PrivValidatorAPI/SignProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivValidatorAPIServer is the server API for PrivValidatorAPI service.
type PrivValidatorAPIServer interface {
	GetPubKey(context.Context, *PubKeyRequest) (*PubKeyResponse, error)
	SignVote(context.Context, *SignVoteRequest) (*SignedVoteResponse, error)
	SignProposal(context.Context, *SignProposalRequest) (*SignedProposalResponse, error)
}

// UnimplementedPrivValidatorAPIServer can be embedded to have forward compatible implementations.
type UnimplementedPrivValidatorAPIServer struct {
}

func (*UnimplementedPrivValidatorAPIServer) GetPubKey(ctx context.Context, req *PubKeyRequest) (*PubKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPubKey not implemented")
}
func (*UnimplementedPrivValidatorAPIServer) SignVote(ctx context.Context, req *SignVoteRequest) (*SignedVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignVote not implemented")
}
func (*UnimplementedPrivValidatorAPIServer) SignProposal(ctx context.Context, req *SignProposalRequest) (*SignedProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignProposal not implemented")
}

func RegisterPrivValidatorAPIServer(s *grpc.Server, srv PrivValidatorAPIServer) {
	s.RegisterService(&_PrivValidatorAPI_serviceDesc, srv)
}

func _PrivValidatorAPI_GetPubKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivValidatorAPIServer).GetPubKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seitendermint.privval.PrivValidatorAPI/GetPubKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivValidatorAPIServer).GetPubKey(ctx, req.(*PubKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivValidatorAPI_SignVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivValidatorAPIServer).SignVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seitendermint.privval.PrivValidatorAPI/SignVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivValidatorAPIServer).SignVote(ctx, req.(*SignVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivValidatorAPI_SignProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivValidatorAPIServer).SignProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seitendermint.privval.PrivValidatorAPI/SignProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivValidatorAPIServer).SignProposal(ctx, req.(*SignProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrivValidatorAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "seitendermint.privval.PrivValidatorAPI",
	HandlerType: (*PrivValidatorAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPubKey",
			Handler:    _PrivValidatorAPI_GetPubKey_Handler,
		},
		{
			MethodName: "SignVote",
			Handler:    _PrivValidatorAPI_SignVote_Handler,
		},
		{
			MethodName: "SignProposal",
			Handler:    _PrivValidatorAPI_SignProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tendermint/privval/service.proto",
}
