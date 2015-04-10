// Code generated by protoc-gen-go.
// source: service.auth/proto/auth.proto
// DO NOT EDIT!

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	service.auth/proto/auth.proto

It has these top-level messages:
	Req
	Customer
*/
package auth

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Req struct {
	AuthToken string `protobuf:"bytes,1,opt" json:"AuthToken,omitempty"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}

type Customer struct {
	ID        int32  `protobuf:"varint,1,opt" json:"ID,omitempty"`
	Email     string `protobuf:"bytes,2,opt" json:"Email,omitempty"`
	Name      string `protobuf:"bytes,3,opt" json:"Name,omitempty"`
	AuthToken string `protobuf:"bytes,4,opt" json:"AuthToken,omitempty"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}

func init() {
}

// Client API for Auth service

type AuthClient interface {
	GetCustomer(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Customer, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GetCustomer(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := grpc.Invoke(ctx, "/auth.Auth/GetCustomer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	GetCustomer(context.Context, *Req) (*Customer, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_GetCustomer_Handler(srv interface{}, ctx context.Context, buf []byte) (interface{}, error) {
	in := new(Req)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthServer).GetCustomer(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomer",
			Handler:    _Auth_GetCustomer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
