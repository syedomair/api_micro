// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	proto/common.proto

It has these top-level messages:
	Response
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Response struct {
	Result string            `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Data   map[string]string `protobuf:"bytes,2,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Response) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *Response) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Response)(nil), "common.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CommonService service

type CommonServiceClient interface {
}

type commonServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommonServiceClient(cc *grpc.ClientConn) CommonServiceClient {
	return &commonServiceClient{cc}
}

// Server API for CommonService service

type CommonServiceServer interface {
}

func RegisterCommonServiceServer(s *grpc.Server, srv CommonServiceServer) {
	s.RegisterService(&_CommonService_serviceDesc, srv)
}

var _CommonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "common.CommonService",
	HandlerType: (*CommonServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "proto/common.proto",
}

func init() { proto.RegisterFile("proto/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x03, 0x73, 0x84, 0xd8, 0x20, 0x3c, 0x29,
	0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc,
	0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x2a, 0xa5, 0x6e, 0x46, 0x2e, 0x8e, 0xa0, 0xd4,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0x4f, 0x48, 0x8f, 0x8b, 0x25, 0x25, 0xb1, 0x24,
	0x51, 0x82, 0x49, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x4a, 0x0f, 0x6a, 0x0f, 0x4c, 0x9f, 0x9e, 0x4b,
	0x62, 0x49, 0xa2, 0x6b, 0x5e, 0x49, 0x51, 0x65, 0x10, 0x58, 0x9d, 0x94, 0x39, 0x17, 0x27, 0x5c,
	0x48, 0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x12, 0x6a, 0x22, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a,
	0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0x04, 0x16, 0x83, 0x70, 0xac, 0x98, 0x2c, 0x18, 0x8d, 0xf8,
	0xb9, 0x78, 0x9d, 0xc1, 0x66, 0x07, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x26, 0xb1, 0x81, 0x5d,
	0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x52, 0xeb, 0x16, 0xe1, 0x00, 0x00, 0x00,
}
