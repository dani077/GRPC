// Code generated by protoc-gen-go. DO NOT EDIT.
// source: menu.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	menu.proto

It has these top-level messages:
	AddMenuReq
	ReadMenuByNamaMenuReq
	ReadMenuByNamaMenuResp
	ReadMenuResp
	UpdateMenurReq
	ReadMenuByKeteranganMenuReq
	ReadMenuByKeteranganMenuResp
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type AddMenuReq struct {
	Namamenu       string `protobuf:"bytes,1,opt,name=namamenu" json:"namamenu,omitempty"`
	Harga          int32  `protobuf:"varint,2,opt,name=harga" json:"harga,omitempty"`
	Idkategorimenu int32  `protobuf:"varint,3,opt,name=idkategorimenu" json:"idkategorimenu,omitempty"`
	Status         int32  `protobuf:"varint,4,opt,name=status" json:"status,omitempty"`
	Createdby      string `protobuf:"bytes,5,opt,name=createdby" json:"createdby,omitempty"`
}

func (m *AddMenuReq) Reset()                    { *m = AddMenuReq{} }
func (m *AddMenuReq) String() string            { return proto.CompactTextString(m) }
func (*AddMenuReq) ProtoMessage()               {}
func (*AddMenuReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddMenuReq) GetNamamenu() string {
	if m != nil {
		return m.Namamenu
	}
	return ""
}

func (m *AddMenuReq) GetHarga() int32 {
	if m != nil {
		return m.Harga
	}
	return 0
}

func (m *AddMenuReq) GetIdkategorimenu() int32 {
	if m != nil {
		return m.Idkategorimenu
	}
	return 0
}

func (m *AddMenuReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AddMenuReq) GetCreatedby() string {
	if m != nil {
		return m.Createdby
	}
	return ""
}

type ReadMenuByNamaMenuReq struct {
	Namamenu string `protobuf:"bytes,1,opt,name=namamenu" json:"namamenu,omitempty"`
}

func (m *ReadMenuByNamaMenuReq) Reset()                    { *m = ReadMenuByNamaMenuReq{} }
func (m *ReadMenuByNamaMenuReq) String() string            { return proto.CompactTextString(m) }
func (*ReadMenuByNamaMenuReq) ProtoMessage()               {}
func (*ReadMenuByNamaMenuReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadMenuByNamaMenuReq) GetNamamenu() string {
	if m != nil {
		return m.Namamenu
	}
	return ""
}

type ReadMenuByNamaMenuResp struct {
	Idmenu         int32  `protobuf:"varint,1,opt,name=idmenu" json:"idmenu,omitempty"`
	Namamenu       string `protobuf:"bytes,2,opt,name=namamenu" json:"namamenu,omitempty"`
	Harga          int32  `protobuf:"varint,3,opt,name=harga" json:"harga,omitempty"`
	Idkategorimenu int32  `protobuf:"varint,4,opt,name=idkategorimenu" json:"idkategorimenu,omitempty"`
	Status         int32  `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`
	Createdby      string `protobuf:"bytes,6,opt,name=createdby" json:"createdby,omitempty"`
	Keterangan     string `protobuf:"bytes,7,opt,name=keterangan" json:"keterangan,omitempty"`
	CreatedOn      string `protobuf:"bytes,8,opt,name=createdOn" json:"createdOn,omitempty"`
}

func (m *ReadMenuByNamaMenuResp) Reset()                    { *m = ReadMenuByNamaMenuResp{} }
func (m *ReadMenuByNamaMenuResp) String() string            { return proto.CompactTextString(m) }
func (*ReadMenuByNamaMenuResp) ProtoMessage()               {}
func (*ReadMenuByNamaMenuResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadMenuByNamaMenuResp) GetIdmenu() int32 {
	if m != nil {
		return m.Idmenu
	}
	return 0
}

func (m *ReadMenuByNamaMenuResp) GetNamamenu() string {
	if m != nil {
		return m.Namamenu
	}
	return ""
}

func (m *ReadMenuByNamaMenuResp) GetHarga() int32 {
	if m != nil {
		return m.Harga
	}
	return 0
}

func (m *ReadMenuByNamaMenuResp) GetIdkategorimenu() int32 {
	if m != nil {
		return m.Idkategorimenu
	}
	return 0
}

func (m *ReadMenuByNamaMenuResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReadMenuByNamaMenuResp) GetCreatedby() string {
	if m != nil {
		return m.Createdby
	}
	return ""
}

func (m *ReadMenuByNamaMenuResp) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

func (m *ReadMenuByNamaMenuResp) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

type ReadMenuResp struct {
	AllMenu []*ReadMenuByNamaMenuResp `protobuf:"bytes,1,rep,name=allMenu" json:"allMenu,omitempty"`
}

func (m *ReadMenuResp) Reset()                    { *m = ReadMenuResp{} }
func (m *ReadMenuResp) String() string            { return proto.CompactTextString(m) }
func (*ReadMenuResp) ProtoMessage()               {}
func (*ReadMenuResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadMenuResp) GetAllMenu() []*ReadMenuByNamaMenuResp {
	if m != nil {
		return m.AllMenu
	}
	return nil
}

type UpdateMenurReq struct {
	Idmenu         int32  `protobuf:"varint,1,opt,name=idmenu" json:"idmenu,omitempty"`
	Namamenu       string `protobuf:"bytes,2,opt,name=namamenu" json:"namamenu,omitempty"`
	Harga          int32  `protobuf:"varint,3,opt,name=harga" json:"harga,omitempty"`
	Idkategorimenu int32  `protobuf:"varint,4,opt,name=idkategorimenu" json:"idkategorimenu,omitempty"`
	Status         int32  `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`
	UpdateBy       string `protobuf:"bytes,6,opt,name=updateBy" json:"updateBy,omitempty"`
}

func (m *UpdateMenurReq) Reset()                    { *m = UpdateMenurReq{} }
func (m *UpdateMenurReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateMenurReq) ProtoMessage()               {}
func (*UpdateMenurReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateMenurReq) GetIdmenu() int32 {
	if m != nil {
		return m.Idmenu
	}
	return 0
}

func (m *UpdateMenurReq) GetNamamenu() string {
	if m != nil {
		return m.Namamenu
	}
	return ""
}

func (m *UpdateMenurReq) GetHarga() int32 {
	if m != nil {
		return m.Harga
	}
	return 0
}

func (m *UpdateMenurReq) GetIdkategorimenu() int32 {
	if m != nil {
		return m.Idkategorimenu
	}
	return 0
}

func (m *UpdateMenurReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateMenurReq) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

type ReadMenuByKeteranganMenuReq struct {
	Keterangan string `protobuf:"bytes,1,opt,name=keterangan" json:"keterangan,omitempty"`
}

func (m *ReadMenuByKeteranganMenuReq) Reset()                    { *m = ReadMenuByKeteranganMenuReq{} }
func (m *ReadMenuByKeteranganMenuReq) String() string            { return proto.CompactTextString(m) }
func (*ReadMenuByKeteranganMenuReq) ProtoMessage()               {}
func (*ReadMenuByKeteranganMenuReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReadMenuByKeteranganMenuReq) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

type ReadMenuByKeteranganMenuResp struct {
	AllKeteranganMenu []*ReadMenuByNamaMenuResp `protobuf:"bytes,1,rep,name=allKeteranganMenu" json:"allKeteranganMenu,omitempty"`
}

func (m *ReadMenuByKeteranganMenuResp) Reset()                    { *m = ReadMenuByKeteranganMenuResp{} }
func (m *ReadMenuByKeteranganMenuResp) String() string            { return proto.CompactTextString(m) }
func (*ReadMenuByKeteranganMenuResp) ProtoMessage()               {}
func (*ReadMenuByKeteranganMenuResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadMenuByKeteranganMenuResp) GetAllKeteranganMenu() []*ReadMenuByNamaMenuResp {
	if m != nil {
		return m.AllKeteranganMenu
	}
	return nil
}

func init() {
	proto.RegisterType((*AddMenuReq)(nil), "grpc.AddMenuReq")
	proto.RegisterType((*ReadMenuByNamaMenuReq)(nil), "grpc.ReadMenuByNamaMenuReq")
	proto.RegisterType((*ReadMenuByNamaMenuResp)(nil), "grpc.ReadMenuByNamaMenuResp")
	proto.RegisterType((*ReadMenuResp)(nil), "grpc.ReadMenuResp")
	proto.RegisterType((*UpdateMenurReq)(nil), "grpc.UpdateMenurReq")
	proto.RegisterType((*ReadMenuByKeteranganMenuReq)(nil), "grpc.ReadMenuByKeteranganMenuReq")
	proto.RegisterType((*ReadMenuByKeteranganMenuResp)(nil), "grpc.ReadMenuByKeteranganMenuResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for MenuService service

type MenuServiceClient interface {
	AddMenu(ctx context.Context, in *AddMenuReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadMenuByNamaMenu(ctx context.Context, in *ReadMenuByNamaMenuReq, opts ...grpc1.CallOption) (*ReadMenuByNamaMenuResp, error)
	ReadMenu(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadMenuResp, error)
	UpdateMenu(ctx context.Context, in *UpdateMenurReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadMenuByKeteranganMenu(ctx context.Context, in *ReadMenuByKeteranganMenuReq, opts ...grpc1.CallOption) (*ReadMenuByKeteranganMenuResp, error)
}

type menuServiceClient struct {
	cc *grpc1.ClientConn
}

func NewMenuServiceClient(cc *grpc1.ClientConn) MenuServiceClient {
	return &menuServiceClient{cc}
}

func (c *menuServiceClient) AddMenu(ctx context.Context, in *AddMenuReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.MenuService/AddMenu", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) ReadMenuByNamaMenu(ctx context.Context, in *ReadMenuByNamaMenuReq, opts ...grpc1.CallOption) (*ReadMenuByNamaMenuResp, error) {
	out := new(ReadMenuByNamaMenuResp)
	err := grpc1.Invoke(ctx, "/grpc.MenuService/ReadMenuByNamaMenu", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) ReadMenu(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadMenuResp, error) {
	out := new(ReadMenuResp)
	err := grpc1.Invoke(ctx, "/grpc.MenuService/ReadMenu", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) UpdateMenu(ctx context.Context, in *UpdateMenurReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.MenuService/UpdateMenu", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) ReadMenuByKeteranganMenu(ctx context.Context, in *ReadMenuByKeteranganMenuReq, opts ...grpc1.CallOption) (*ReadMenuByKeteranganMenuResp, error) {
	out := new(ReadMenuByKeteranganMenuResp)
	err := grpc1.Invoke(ctx, "/grpc.MenuService/ReadMenuByKeteranganMenu", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MenuService service

type MenuServiceServer interface {
	AddMenu(context.Context, *AddMenuReq) (*google_protobuf.Empty, error)
	ReadMenuByNamaMenu(context.Context, *ReadMenuByNamaMenuReq) (*ReadMenuByNamaMenuResp, error)
	ReadMenu(context.Context, *google_protobuf.Empty) (*ReadMenuResp, error)
	UpdateMenu(context.Context, *UpdateMenurReq) (*google_protobuf.Empty, error)
	ReadMenuByKeteranganMenu(context.Context, *ReadMenuByKeteranganMenuReq) (*ReadMenuByKeteranganMenuResp, error)
}

func RegisterMenuServiceServer(s *grpc1.Server, srv MenuServiceServer) {
	s.RegisterService(&_MenuService_serviceDesc, srv)
}

func _MenuService_AddMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).AddMenu(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MenuService/AddMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).AddMenu(ctx, req.(*AddMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_ReadMenuByNamaMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMenuByNamaMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).ReadMenuByNamaMenu(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MenuService/ReadMenuByNamaMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).ReadMenuByNamaMenu(ctx, req.(*ReadMenuByNamaMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_ReadMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).ReadMenu(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MenuService/ReadMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).ReadMenu(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_UpdateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMenurReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).UpdateMenu(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MenuService/UpdateMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).UpdateMenu(ctx, req.(*UpdateMenurReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_ReadMenuByKeteranganMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMenuByKeteranganMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).ReadMenuByKeteranganMenu(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MenuService/ReadMenuByKeteranganMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).ReadMenuByKeteranganMenu(ctx, req.(*ReadMenuByKeteranganMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MenuService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.MenuService",
	HandlerType: (*MenuServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddMenu",
			Handler:    _MenuService_AddMenu_Handler,
		},
		{
			MethodName: "ReadMenuByNamaMenu",
			Handler:    _MenuService_ReadMenuByNamaMenu_Handler,
		},
		{
			MethodName: "ReadMenu",
			Handler:    _MenuService_ReadMenu_Handler,
		},
		{
			MethodName: "UpdateMenu",
			Handler:    _MenuService_UpdateMenu_Handler,
		},
		{
			MethodName: "ReadMenuByKeteranganMenu",
			Handler:    _MenuService_ReadMenuByKeteranganMenu_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "menu.proto",
}

func init() { proto.RegisterFile("menu.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x93, 0x38, 0x71, 0xa7, 0xa8, 0x82, 0x51, 0x89, 0x2c, 0x27, 0x42, 0x65, 0x0f, 0x28,
	0x27, 0x57, 0x6a, 0x05, 0xe2, 0x00, 0x07, 0x2a, 0xc1, 0x01, 0x54, 0x2a, 0x8c, 0x78, 0x80, 0x49,
	0xbc, 0x18, 0x53, 0xff, 0x75, 0xbd, 0x46, 0xca, 0xbb, 0xf0, 0x1c, 0x3c, 0x18, 0x0f, 0x80, 0x90,
	0xd7, 0x7f, 0xb5, 0x53, 0x27, 0x39, 0xf6, 0xe6, 0xd9, 0x99, 0x6f, 0xe6, 0xfb, 0x3e, 0xcf, 0x00,
	0x84, 0x3c, 0xca, 0xec, 0x44, 0xc4, 0x32, 0xc6, 0x91, 0x27, 0x92, 0x95, 0x35, 0xf3, 0xe2, 0xd8,
	0x0b, 0xf8, 0x99, 0x7a, 0x5b, 0x66, 0xdf, 0xcf, 0x78, 0x98, 0xc8, 0x75, 0x51, 0xc2, 0x7e, 0x6b,
	0x00, 0xef, 0x5c, 0xf7, 0x8a, 0x47, 0x99, 0xc3, 0x6f, 0xd1, 0x02, 0x23, 0xa2, 0x90, 0xf2, 0x1e,
	0xa6, 0x76, 0xaa, 0x2d, 0x0e, 0x9d, 0x3a, 0xc6, 0x13, 0xd0, 0x7f, 0x90, 0xf0, 0xc8, 0x1c, 0x9c,
	0x6a, 0x0b, 0xdd, 0x29, 0x02, 0x7c, 0x01, 0xc7, 0xbe, 0x7b, 0x43, 0x92, 0x7b, 0xb1, 0xf0, 0x15,
	0x6e, 0xa8, 0xd2, 0x9d, 0x57, 0x9c, 0xc2, 0x38, 0x95, 0x24, 0xb3, 0xd4, 0x1c, 0xa9, 0x7c, 0x19,
	0xe1, 0x1c, 0x0e, 0x57, 0x82, 0x93, 0xe4, 0xee, 0x72, 0x6d, 0xea, 0x6a, 0x64, 0xf3, 0xc0, 0x2e,
	0xe0, 0xa9, 0xc3, 0x49, 0xd1, 0xbb, 0x5c, 0x7f, 0xa6, 0x90, 0xf6, 0x20, 0xca, 0xfe, 0x69, 0x30,
	0xbd, 0x0f, 0x95, 0x26, 0x39, 0x0b, 0xdf, 0xad, 0x41, 0xba, 0x53, 0x46, 0xad, 0x76, 0x83, 0x3e,
	0xdd, 0xc3, 0xed, 0xba, 0x47, 0x3b, 0x74, 0xeb, 0xfd, 0xba, 0xc7, 0x1d, 0xdd, 0xf8, 0x0c, 0xe0,
	0x86, 0x4b, 0x2e, 0x28, 0xf2, 0x28, 0x32, 0x27, 0x2a, 0x7d, 0xe7, 0xe5, 0x0e, 0xfa, 0x3a, 0x32,
	0x8d, 0x16, 0xfa, 0x3a, 0x62, 0x1f, 0xe0, 0x51, 0xa5, 0x5f, 0xa9, 0x7e, 0x05, 0x13, 0x0a, 0x82,
	0xab, 0x42, 0xf6, 0x70, 0x71, 0x74, 0x3e, 0xb7, 0xf3, 0xcd, 0xb0, 0xef, 0x37, 0xc9, 0xa9, 0x8a,
	0xd9, 0x1f, 0x0d, 0x8e, 0xbf, 0x25, 0x2e, 0x49, 0x9e, 0x87, 0x22, 0xf7, 0xfd, 0xe1, 0x18, 0x68,
	0x81, 0x91, 0x29, 0x6e, 0x97, 0x95, 0x7f, 0x75, 0xcc, 0xde, 0xc2, 0xac, 0xd1, 0xf6, 0xa9, 0xb6,
	0xad, 0x5a, 0x9e, 0xb6, 0xbb, 0x5a, 0xd7, 0x5d, 0xf6, 0x13, 0xe6, 0xfd, 0xf0, 0x34, 0xc1, 0x8f,
	0xf0, 0x84, 0x82, 0xa0, 0x9d, 0xd8, 0xcb, 0xd9, 0x4d, 0xd8, 0xf9, 0xdf, 0x01, 0x1c, 0xe5, 0x1f,
	0x5f, 0xb9, 0xf8, 0xe5, 0xaf, 0x38, 0xbe, 0x84, 0x49, 0x79, 0x8f, 0xf8, 0xb8, 0xe8, 0xd5, 0x9c,
	0xa7, 0x35, 0xb5, 0x8b, 0x5b, 0xb6, 0xab, 0x5b, 0xb6, 0xdf, 0xe7, 0xb7, 0xcc, 0x0e, 0xf0, 0x0b,
	0xe0, 0xe6, 0x4c, 0x9c, 0xf5, 0xb3, 0xb9, 0xb5, 0xb6, 0x52, 0x65, 0x07, 0xf8, 0x1a, 0x8c, 0x2a,
	0x87, 0x3d, 0x83, 0x2d, 0x6c, 0xf7, 0x28, 0x91, 0x6f, 0x00, 0x9a, 0xb5, 0xc1, 0x93, 0xa2, 0xa6,
	0xbd, 0x48, 0x5b, 0xa4, 0x10, 0x98, 0x7d, 0xee, 0xe3, 0xf3, 0x2e, 0xe7, 0x8d, 0x9f, 0x6b, 0xb1,
	0x5d, 0x25, 0x69, 0xb2, 0x1c, 0xab, 0xa1, 0x17, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x1f,
	0x7f, 0xeb, 0x2d, 0x05, 0x00, 0x00,
}
