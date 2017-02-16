// Code generated by protoc-gen-go.
// source: proto/movie.proto
// DO NOT EDIT!

package renderdemo

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

type MakeMovieRequest struct {
	// GCS path to write output movie to.
	GcsOutput string `protobuf:"bytes,1,opt,name=gcs_output,json=gcsOutput" json:"gcs_output,omitempty"`
	// GCS paths to frames, in sequence.
	Frames []string `protobuf:"bytes,2,rep,name=frames" json:"frames,omitempty"`
}

func (m *MakeMovieRequest) Reset()                    { *m = MakeMovieRequest{} }
func (m *MakeMovieRequest) String() string            { return proto.CompactTextString(m) }
func (*MakeMovieRequest) ProtoMessage()               {}
func (*MakeMovieRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MakeMovieRequest) GetGcsOutput() string {
	if m != nil {
		return m.GcsOutput
	}
	return ""
}

func (m *MakeMovieRequest) GetFrames() []string {
	if m != nil {
		return m.Frames
	}
	return nil
}

type MakeMovieResponse struct {
}

func (m *MakeMovieResponse) Reset()                    { *m = MakeMovieResponse{} }
func (m *MakeMovieResponse) String() string            { return proto.CompactTextString(m) }
func (*MakeMovieResponse) ProtoMessage()               {}
func (*MakeMovieResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*MakeMovieRequest)(nil), "renderdemo.MakeMovieRequest")
	proto.RegisterType((*MakeMovieResponse)(nil), "renderdemo.MakeMovieResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Sequencer service

type SequencerClient interface {
	MakeMovie(ctx context.Context, in *MakeMovieRequest, opts ...grpc.CallOption) (*MakeMovieResponse, error)
}

type sequencerClient struct {
	cc *grpc.ClientConn
}

func NewSequencerClient(cc *grpc.ClientConn) SequencerClient {
	return &sequencerClient{cc}
}

func (c *sequencerClient) MakeMovie(ctx context.Context, in *MakeMovieRequest, opts ...grpc.CallOption) (*MakeMovieResponse, error) {
	out := new(MakeMovieResponse)
	err := grpc.Invoke(ctx, "/renderdemo.Sequencer/MakeMovie", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sequencer service

type SequencerServer interface {
	MakeMovie(context.Context, *MakeMovieRequest) (*MakeMovieResponse, error)
}

func RegisterSequencerServer(s *grpc.Server, srv SequencerServer) {
	s.RegisterService(&_Sequencer_serviceDesc, srv)
}

func _Sequencer_MakeMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequencerServer).MakeMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/renderdemo.Sequencer/MakeMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequencerServer).MakeMovie(ctx, req.(*MakeMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sequencer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "renderdemo.Sequencer",
	HandlerType: (*SequencerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeMovie",
			Handler:    _Sequencer_MakeMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/movie.proto",
}

func init() { proto.RegisterFile("proto/movie.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xcd, 0x2f, 0xcb, 0x4c, 0xd5, 0x03, 0xb3, 0x85, 0xb8, 0x8a, 0x52, 0xf3, 0x52,
	0x52, 0x8b, 0x52, 0x52, 0x73, 0xf3, 0x95, 0x3c, 0xb9, 0x04, 0x7c, 0x13, 0xb3, 0x53, 0x7d, 0x41,
	0xd2, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xb2, 0x5c, 0x5c, 0xe9, 0xc9, 0xc5, 0xf1,
	0xf9, 0xa5, 0x25, 0x05, 0xa5, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x9c, 0xe9, 0xc9,
	0xc5, 0xfe, 0x60, 0x01, 0x21, 0x31, 0x2e, 0xb6, 0xb4, 0xa2, 0xc4, 0xdc, 0xd4, 0x62, 0x09, 0x26,
	0x05, 0x66, 0x0d, 0xce, 0x20, 0x28, 0x4f, 0x49, 0x98, 0x4b, 0x10, 0xc9, 0xa8, 0xe2, 0x82, 0xfc,
	0xbc, 0xe2, 0x54, 0xa3, 0x50, 0x2e, 0xce, 0x60, 0x90, 0xb1, 0x79, 0xc9, 0xa9, 0x45, 0x42, 0x1e,
	0x5c, 0x9c, 0x70, 0x15, 0x42, 0x32, 0x7a, 0x08, 0x67, 0xe8, 0xa1, 0xbb, 0x41, 0x4a, 0x16, 0x87,
	0x2c, 0xc4, 0x58, 0x27, 0x9e, 0x28, 0x24, 0x4f, 0x24, 0xb1, 0x81, 0xfd, 0x65, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xb2, 0x20, 0x68, 0x66, 0xec, 0x00, 0x00, 0x00,
}
