// Code generated by protoc-gen-go. DO NOT EDIT.
// source: items.proto

/*
Package marketplace is a generated protocol buffer package.

It is generated from these files:
	items.proto

It has these top-level messages:
	ListItemsRequest
	ListItemsReply
	Item
*/
package marketplace

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type ListItemsRequest struct {
	PageNumber int32 `protobuf:"varint,1,opt,name=page_number,json=pageNumber" json:"page_number,omitempty"`
}

func (m *ListItemsRequest) Reset()                    { *m = ListItemsRequest{} }
func (m *ListItemsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListItemsRequest) ProtoMessage()               {}
func (*ListItemsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListItemsRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

type ListItemsReply struct {
	Items []*Item `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ListItemsReply) Reset()                    { *m = ListItemsReply{} }
func (m *ListItemsReply) String() string            { return proto.CompactTextString(m) }
func (*ListItemsReply) ProtoMessage()               {}
func (*ListItemsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListItemsReply) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type Item struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Price int32  `protobuf:"varint,3,opt,name=price" json:"price,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*ListItemsRequest)(nil), "marketplace.ListItemsRequest")
	proto.RegisterType((*ListItemsReply)(nil), "marketplace.ListItemsReply")
	proto.RegisterType((*Item)(nil), "marketplace.Item")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Items service

type ItemsClient interface {
	ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsReply, error)
}

type itemsClient struct {
	cc *grpc.ClientConn
}

func NewItemsClient(cc *grpc.ClientConn) ItemsClient {
	return &itemsClient{cc}
}

func (c *itemsClient) ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsReply, error) {
	out := new(ListItemsReply)
	err := grpc.Invoke(ctx, "/marketplace.Items/ListItems", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Items service

type ItemsServer interface {
	ListItems(context.Context, *ListItemsRequest) (*ListItemsReply, error)
}

func RegisterItemsServer(s *grpc.Server, srv ItemsServer) {
	s.RegisterService(&_Items_serviceDesc, srv)
}

func _Items_ListItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServer).ListItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketplace.Items/ListItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServer).ListItems(ctx, req.(*ListItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Items_serviceDesc = grpc.ServiceDesc{
	ServiceName: "marketplace.Items",
	HandlerType: (*ItemsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListItems",
			Handler:    _Items_ListItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "items.proto",
}

func init() { proto.RegisterFile("items.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x2c, 0x49, 0xcd,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0x29,
	0xc8, 0x49, 0x4c, 0x4e, 0x55, 0x32, 0xe6, 0x12, 0xf0, 0xc9, 0x2c, 0x2e, 0xf1, 0x04, 0xc9, 0x07,
	0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xc9, 0x73, 0x71, 0x17, 0x24, 0xa6, 0xa7, 0xc6, 0xe7,
	0x95, 0xe6, 0x26, 0xa5, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x71, 0x81, 0x84, 0xfc,
	0xc0, 0x22, 0x4a, 0x96, 0x5c, 0x7c, 0x48, 0x9a, 0x0a, 0x72, 0x2a, 0x85, 0xd4, 0xb9, 0x58, 0xc1,
	0x56, 0x48, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x09, 0xea, 0x21, 0xd9, 0xa1, 0x07, 0x52, 0x17,
	0x04, 0x91, 0x57, 0x72, 0xe0, 0x62, 0x01, 0x71, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0xc0, 0x46,
	0x73, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81,
	0x45, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0x82, 0xa2, 0xcc, 0xe4, 0x54, 0x09, 0x66, 0xb0, 0x0b,
	0x20, 0x1c, 0xa3, 0x20, 0x2e, 0x56, 0xb0, 0xc5, 0x42, 0x9e, 0x5c, 0x9c, 0x70, 0x57, 0x08, 0xc9,
	0xa2, 0xd8, 0x88, 0xee, 0x25, 0x29, 0x69, 0x5c, 0xd2, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x49, 0x6c,
	0xe0, 0x90, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xf3, 0xca, 0x3a, 0x28, 0x01, 0x00,
	0x00,
}
