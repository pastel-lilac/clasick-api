// Code generated by protoc-gen-go. DO NOT EDIT.
// source: music/v1alpha/music.proto

package v1alpha

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	v1alpha "github.com/pastel-lilac/clasick-api/pkg/core/apis/common/v1alpha"
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

type Music struct {
	Id                   uint32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	JacketPhoto          string               `protobuf:"bytes,3,opt,name=jacket_photo,json=jacketPhoto,proto3" json:"jacket_photo,omitempty"`
	DataSource           string               `protobuf:"bytes,4,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	ReleaseDate          *timestamp.Timestamp `protobuf:"bytes,5,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Music) Reset()         { *m = Music{} }
func (m *Music) String() string { return proto.CompactTextString(m) }
func (*Music) ProtoMessage()    {}
func (*Music) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30112ff8dc588d2, []int{0}
}

func (m *Music) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Music.Unmarshal(m, b)
}
func (m *Music) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Music.Marshal(b, m, deterministic)
}
func (m *Music) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Music.Merge(m, src)
}
func (m *Music) XXX_Size() int {
	return xxx_messageInfo_Music.Size(m)
}
func (m *Music) XXX_DiscardUnknown() {
	xxx_messageInfo_Music.DiscardUnknown(m)
}

var xxx_messageInfo_Music proto.InternalMessageInfo

func (m *Music) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Music) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Music) GetJacketPhoto() string {
	if m != nil {
		return m.JacketPhoto
	}
	return ""
}

func (m *Music) GetDataSource() string {
	if m != nil {
		return m.DataSource
	}
	return ""
}

func (m *Music) GetReleaseDate() *timestamp.Timestamp {
	if m != nil {
		return m.ReleaseDate
	}
	return nil
}

type MusicResponse struct {
	Musics               []*Music `protobuf:"bytes,1,rep,name=musics,proto3" json:"musics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MusicResponse) Reset()         { *m = MusicResponse{} }
func (m *MusicResponse) String() string { return proto.CompactTextString(m) }
func (*MusicResponse) ProtoMessage()    {}
func (*MusicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30112ff8dc588d2, []int{1}
}

func (m *MusicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MusicResponse.Unmarshal(m, b)
}
func (m *MusicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MusicResponse.Marshal(b, m, deterministic)
}
func (m *MusicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MusicResponse.Merge(m, src)
}
func (m *MusicResponse) XXX_Size() int {
	return xxx_messageInfo_MusicResponse.Size(m)
}
func (m *MusicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MusicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MusicResponse proto.InternalMessageInfo

func (m *MusicResponse) GetMusics() []*Music {
	if m != nil {
		return m.Musics
	}
	return nil
}

// --- [GET] /musics/:artist_id --- //
type GetArtistMusicsRequest struct {
	Field                *v1alpha.MustField `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	ArtistId             uint32             `protobuf:"varint,2,opt,name=artist_id,json=artistId,proto3" json:"artist_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetArtistMusicsRequest) Reset()         { *m = GetArtistMusicsRequest{} }
func (m *GetArtistMusicsRequest) String() string { return proto.CompactTextString(m) }
func (*GetArtistMusicsRequest) ProtoMessage()    {}
func (*GetArtistMusicsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30112ff8dc588d2, []int{2}
}

func (m *GetArtistMusicsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtistMusicsRequest.Unmarshal(m, b)
}
func (m *GetArtistMusicsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtistMusicsRequest.Marshal(b, m, deterministic)
}
func (m *GetArtistMusicsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtistMusicsRequest.Merge(m, src)
}
func (m *GetArtistMusicsRequest) XXX_Size() int {
	return xxx_messageInfo_GetArtistMusicsRequest.Size(m)
}
func (m *GetArtistMusicsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtistMusicsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtistMusicsRequest proto.InternalMessageInfo

func (m *GetArtistMusicsRequest) GetField() *v1alpha.MustField {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *GetArtistMusicsRequest) GetArtistId() uint32 {
	if m != nil {
		return m.ArtistId
	}
	return 0
}

// --- [GET] /musics/:album_id --- //
type GetAlbumMusicsRequest struct {
	Field                *v1alpha.MustField `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	AlbumId              uint32             `protobuf:"varint,2,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetAlbumMusicsRequest) Reset()         { *m = GetAlbumMusicsRequest{} }
func (m *GetAlbumMusicsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAlbumMusicsRequest) ProtoMessage()    {}
func (*GetAlbumMusicsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30112ff8dc588d2, []int{3}
}

func (m *GetAlbumMusicsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAlbumMusicsRequest.Unmarshal(m, b)
}
func (m *GetAlbumMusicsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAlbumMusicsRequest.Marshal(b, m, deterministic)
}
func (m *GetAlbumMusicsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAlbumMusicsRequest.Merge(m, src)
}
func (m *GetAlbumMusicsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAlbumMusicsRequest.Size(m)
}
func (m *GetAlbumMusicsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAlbumMusicsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAlbumMusicsRequest proto.InternalMessageInfo

func (m *GetAlbumMusicsRequest) GetField() *v1alpha.MustField {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *GetAlbumMusicsRequest) GetAlbumId() uint32 {
	if m != nil {
		return m.AlbumId
	}
	return 0
}

// --- [GET] /musics/:playlist_id --- //
type GetPlaylistMusicsRequest struct {
	Field                *v1alpha.MustField `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	PlaylistId           uint32             `protobuf:"varint,2,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetPlaylistMusicsRequest) Reset()         { *m = GetPlaylistMusicsRequest{} }
func (m *GetPlaylistMusicsRequest) String() string { return proto.CompactTextString(m) }
func (*GetPlaylistMusicsRequest) ProtoMessage()    {}
func (*GetPlaylistMusicsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30112ff8dc588d2, []int{4}
}

func (m *GetPlaylistMusicsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlaylistMusicsRequest.Unmarshal(m, b)
}
func (m *GetPlaylistMusicsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlaylistMusicsRequest.Marshal(b, m, deterministic)
}
func (m *GetPlaylistMusicsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlaylistMusicsRequest.Merge(m, src)
}
func (m *GetPlaylistMusicsRequest) XXX_Size() int {
	return xxx_messageInfo_GetPlaylistMusicsRequest.Size(m)
}
func (m *GetPlaylistMusicsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlaylistMusicsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlaylistMusicsRequest proto.InternalMessageInfo

func (m *GetPlaylistMusicsRequest) GetField() *v1alpha.MustField {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *GetPlaylistMusicsRequest) GetPlaylistId() uint32 {
	if m != nil {
		return m.PlaylistId
	}
	return 0
}

func init() {
	proto.RegisterType((*Music)(nil), "v1alpha.Music")
	proto.RegisterType((*MusicResponse)(nil), "v1alpha.MusicResponse")
	proto.RegisterType((*GetArtistMusicsRequest)(nil), "v1alpha.GetArtistMusicsRequest")
	proto.RegisterType((*GetAlbumMusicsRequest)(nil), "v1alpha.GetAlbumMusicsRequest")
	proto.RegisterType((*GetPlaylistMusicsRequest)(nil), "v1alpha.GetPlaylistMusicsRequest")
}

func init() { proto.RegisterFile("music/v1alpha/music.proto", fileDescriptor_c30112ff8dc588d2) }

var fileDescriptor_c30112ff8dc588d2 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x41, 0x6f, 0xd3, 0x3e,
	0x18, 0xc6, 0xff, 0xe9, 0xd6, 0xad, 0x7b, 0xb3, 0xf6, 0x2f, 0x2c, 0x31, 0x65, 0x9d, 0x44, 0xbb,
	0x1c, 0x50, 0x2e, 0x8b, 0x45, 0x39, 0x70, 0x42, 0x68, 0x80, 0x98, 0x8a, 0x40, 0xaa, 0x02, 0x27,
	0x84, 0x14, 0xb9, 0xce, 0xbb, 0xd6, 0xd4, 0xa9, 0x4d, 0xec, 0x20, 0xf1, 0xa5, 0xf8, 0x78, 0x9c,
	0x51, 0xec, 0x74, 0x5b, 0xb7, 0xc1, 0x01, 0x6e, 0x79, 0x9f, 0xf7, 0xc9, 0xcf, 0x79, 0x9e, 0x24,
	0x70, 0x5c, 0xd6, 0x46, 0x70, 0xfa, 0xed, 0x09, 0x93, 0x7a, 0xc9, 0xa8, 0x9b, 0x52, 0x5d, 0x29,
	0xab, 0xc8, 0x7e, 0x2b, 0x0e, 0x47, 0x0b, 0xa5, 0x16, 0x12, 0xa9, 0x93, 0xe7, 0xf5, 0x25, 0xb5,
	0xa2, 0x44, 0x63, 0x59, 0xa9, 0xbd, 0x73, 0x78, 0xc2, 0x55, 0x59, 0xaa, 0xf5, 0x15, 0xc5, 0x8f,
	0x7e, 0x19, 0xff, 0x08, 0xa0, 0xfb, 0xbe, 0xc1, 0x92, 0x01, 0x74, 0x44, 0x11, 0x05, 0xe3, 0x20,
	0xe9, 0x67, 0x1d, 0x51, 0x10, 0x02, 0xbb, 0x6b, 0x56, 0x62, 0xd4, 0x19, 0x07, 0xc9, 0x41, 0xe6,
	0xae, 0xc9, 0x29, 0x1c, 0x7e, 0x61, 0x7c, 0x85, 0x36, 0xd7, 0x4b, 0x65, 0x55, 0xb4, 0xe3, 0x76,
	0xa1, 0xd7, 0x66, 0x8d, 0x44, 0x46, 0x10, 0x16, 0xcc, 0xb2, 0xdc, 0xa8, 0xba, 0xe2, 0x18, 0xed,
	0x3a, 0x07, 0x34, 0xd2, 0x07, 0xa7, 0x90, 0xe7, 0x70, 0x58, 0xa1, 0x44, 0x66, 0x30, 0x2f, 0x98,
	0xc5, 0xa8, 0x3b, 0x0e, 0x92, 0x70, 0x32, 0x4c, 0x7d, 0x8c, 0x74, 0x13, 0x23, 0xfd, 0xb8, 0x89,
	0x91, 0x85, 0xad, 0xff, 0x35, 0xb3, 0x18, 0x3f, 0x83, 0xbe, 0x7b, 0xde, 0x0c, 0x8d, 0x56, 0x6b,
	0x83, 0xe4, 0x31, 0xec, 0xb9, 0x5e, 0x4c, 0x14, 0x8c, 0x77, 0x92, 0x70, 0x32, 0x48, 0xdb, 0xa0,
	0xa9, 0xf7, 0xb5, 0xdb, 0x38, 0x87, 0xa3, 0x0b, 0xb4, 0xe7, 0x95, 0x15, 0xc6, 0xba, 0x8d, 0xc9,
	0xf0, 0x6b, 0x8d, 0xc6, 0x92, 0x04, 0xba, 0x97, 0x02, 0xa5, 0x0f, 0x1f, 0x4e, 0xc8, 0x4d, 0x80,
	0x7d, 0xd3, 0x6c, 0x32, 0x6f, 0x20, 0x27, 0x70, 0xc0, 0x1c, 0x20, 0x17, 0x85, 0x2b, 0xa6, 0x9f,
	0xf5, 0xbc, 0x30, 0x2d, 0xe2, 0xcf, 0xf0, 0xb0, 0x39, 0x40, 0xce, 0xeb, 0xf2, 0x6f, 0xf9, 0xc7,
	0xd0, 0x63, 0xcd, 0xfd, 0xd7, 0xf8, 0x7d, 0x37, 0x4f, 0x8b, 0x18, 0x21, 0xba, 0x40, 0x3b, 0x93,
	0xec, 0xbb, 0xfc, 0x87, 0x00, 0x23, 0x08, 0x75, 0x8b, 0xb8, 0x3e, 0x03, 0x36, 0xd2, 0xb4, 0x98,
	0xfc, 0x0c, 0xa0, 0xe7, 0x7b, 0x9b, 0xbd, 0x22, 0xef, 0xe0, 0xff, 0x5b, 0x95, 0x91, 0xd1, 0x15,
	0xfb, 0xfe, 0x32, 0x87, 0x47, 0xb7, 0xea, 0x6f, 0x5f, 0x53, 0xfc, 0x1f, 0x79, 0x0b, 0x83, 0xed,
	0x7e, 0xc8, 0xa3, 0x2d, 0xd8, 0x9d, 0xe2, 0xfe, 0xc0, 0x9a, 0xc1, 0x83, 0x3b, 0x6d, 0x90, 0xd3,
	0x9b, 0xb8, 0x7b, 0x9b, 0xfa, 0x3d, 0xf1, 0xe5, 0xf9, 0xa7, 0x17, 0x0b, 0x61, 0x97, 0xf5, 0x3c,
	0xe5, 0xaa, 0xa4, 0x9a, 0x19, 0x8b, 0xf2, 0x4c, 0x0a, 0xc9, 0x38, 0xe5, 0x92, 0x19, 0xc1, 0x57,
	0x67, 0x4c, 0x0b, 0xaa, 0x57, 0x0b, 0xca, 0x55, 0x85, 0x94, 0x69, 0x61, 0xe8, 0xd6, 0xef, 0x39,
	0xdf, 0x73, 0xdf, 0xee, 0xd3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x56, 0xab, 0x09, 0xc3, 0xb6,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MusicRPCClient is the client API for MusicRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MusicRPCClient interface {
	GetArtistMusics(ctx context.Context, in *GetArtistMusicsRequest, opts ...grpc.CallOption) (*MusicResponse, error)
	GetAlbumMusics(ctx context.Context, in *GetAlbumMusicsRequest, opts ...grpc.CallOption) (*MusicResponse, error)
	GetPlaylistMusics(ctx context.Context, in *GetPlaylistMusicsRequest, opts ...grpc.CallOption) (*MusicResponse, error)
}

type musicRPCClient struct {
	cc *grpc.ClientConn
}

func NewMusicRPCClient(cc *grpc.ClientConn) MusicRPCClient {
	return &musicRPCClient{cc}
}

func (c *musicRPCClient) GetArtistMusics(ctx context.Context, in *GetArtistMusicsRequest, opts ...grpc.CallOption) (*MusicResponse, error) {
	out := new(MusicResponse)
	err := c.cc.Invoke(ctx, "/v1alpha.MusicRPC/GetArtistMusics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicRPCClient) GetAlbumMusics(ctx context.Context, in *GetAlbumMusicsRequest, opts ...grpc.CallOption) (*MusicResponse, error) {
	out := new(MusicResponse)
	err := c.cc.Invoke(ctx, "/v1alpha.MusicRPC/GetAlbumMusics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicRPCClient) GetPlaylistMusics(ctx context.Context, in *GetPlaylistMusicsRequest, opts ...grpc.CallOption) (*MusicResponse, error) {
	out := new(MusicResponse)
	err := c.cc.Invoke(ctx, "/v1alpha.MusicRPC/GetPlaylistMusics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MusicRPCServer is the server API for MusicRPC service.
type MusicRPCServer interface {
	GetArtistMusics(context.Context, *GetArtistMusicsRequest) (*MusicResponse, error)
	GetAlbumMusics(context.Context, *GetAlbumMusicsRequest) (*MusicResponse, error)
	GetPlaylistMusics(context.Context, *GetPlaylistMusicsRequest) (*MusicResponse, error)
}

// UnimplementedMusicRPCServer can be embedded to have forward compatible implementations.
type UnimplementedMusicRPCServer struct {
}

func (*UnimplementedMusicRPCServer) GetArtistMusics(ctx context.Context, req *GetArtistMusicsRequest) (*MusicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtistMusics not implemented")
}
func (*UnimplementedMusicRPCServer) GetAlbumMusics(ctx context.Context, req *GetAlbumMusicsRequest) (*MusicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlbumMusics not implemented")
}
func (*UnimplementedMusicRPCServer) GetPlaylistMusics(ctx context.Context, req *GetPlaylistMusicsRequest) (*MusicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylistMusics not implemented")
}

func RegisterMusicRPCServer(s *grpc.Server, srv MusicRPCServer) {
	s.RegisterService(&_MusicRPC_serviceDesc, srv)
}

func _MusicRPC_GetArtistMusics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtistMusicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicRPCServer).GetArtistMusics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha.MusicRPC/GetArtistMusics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicRPCServer).GetArtistMusics(ctx, req.(*GetArtistMusicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicRPC_GetAlbumMusics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlbumMusicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicRPCServer).GetAlbumMusics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha.MusicRPC/GetAlbumMusics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicRPCServer).GetAlbumMusics(ctx, req.(*GetAlbumMusicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicRPC_GetPlaylistMusics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaylistMusicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicRPCServer).GetPlaylistMusics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha.MusicRPC/GetPlaylistMusics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicRPCServer).GetPlaylistMusics(ctx, req.(*GetPlaylistMusicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MusicRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha.MusicRPC",
	HandlerType: (*MusicRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArtistMusics",
			Handler:    _MusicRPC_GetArtistMusics_Handler,
		},
		{
			MethodName: "GetAlbumMusics",
			Handler:    _MusicRPC_GetAlbumMusics_Handler,
		},
		{
			MethodName: "GetPlaylistMusics",
			Handler:    _MusicRPC_GetPlaylistMusics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "music/v1alpha/music.proto",
}
