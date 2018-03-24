// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: weatherman_svc.proto

/*
Package weatherman is a generated protocol buffer package.

It is generated from these files:
	weatherman_svc.proto

It has these top-level messages:
	HealthCheckResponse
*/
package weatherman

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type HealthCheckResponse struct {
	// Whether the service is overall healthy.
	Healthy bool `protobuf:"varint,1,opt,name=healthy,proto3" json:"healthy,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptorWeathermanSvc, []int{0} }

func (m *HealthCheckResponse) GetHealthy() bool {
	if m != nil {
		return m.Healthy
	}
	return false
}

func init() {
	proto.RegisterType((*HealthCheckResponse)(nil), "weatherman.HealthCheckResponse")
	golang_proto.RegisterType((*HealthCheckResponse)(nil), "weatherman.HealthCheckResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Weatherman service

type WeathermanClient interface {
	// Get service health
	Healthcheck(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type weathermanClient struct {
	cc *grpc.ClientConn
}

func NewWeathermanClient(cc *grpc.ClientConn) WeathermanClient {
	return &weathermanClient{cc}
}

func (c *weathermanClient) Healthcheck(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/weatherman.Weatherman/Healthcheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Weatherman service

type WeathermanServer interface {
	// Get service health
	Healthcheck(context.Context, *google_protobuf.Empty) (*HealthCheckResponse, error)
}

func RegisterWeathermanServer(s *grpc.Server, srv WeathermanServer) {
	s.RegisterService(&_Weatherman_serviceDesc, srv)
}

func _Weatherman_Healthcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeathermanServer).Healthcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weatherman.Weatherman/Healthcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeathermanServer).Healthcheck(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Weatherman_serviceDesc = grpc.ServiceDesc{
	ServiceName: "weatherman.Weatherman",
	HandlerType: (*WeathermanServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthcheck",
			Handler:    _Weatherman_Healthcheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weatherman_svc.proto",
}

func (m *HealthCheckResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthCheckResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Healthy {
		dAtA[i] = 0x8
		i++
		if m.Healthy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintWeathermanSvc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HealthCheckResponse) Size() (n int) {
	var l int
	_ = l
	if m.Healthy {
		n += 2
	}
	return n
}

func sovWeathermanSvc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWeathermanSvc(x uint64) (n int) {
	return sovWeathermanSvc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HealthCheckResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWeathermanSvc
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
			return fmt.Errorf("proto: HealthCheckResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheckResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Healthy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWeathermanSvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Healthy = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipWeathermanSvc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWeathermanSvc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWeathermanSvc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWeathermanSvc
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
					return 0, ErrIntOverflowWeathermanSvc
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
					return 0, ErrIntOverflowWeathermanSvc
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
				return 0, ErrInvalidLengthWeathermanSvc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWeathermanSvc
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
				next, err := skipWeathermanSvc(dAtA[start:])
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
	ErrInvalidLengthWeathermanSvc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWeathermanSvc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("weatherman_svc.proto", fileDescriptorWeathermanSvc) }
func init() { golang_proto.RegisterFile("weatherman_svc.proto", fileDescriptorWeathermanSvc) }

var fileDescriptorWeathermanSvc = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x89, 0x83, 0x4a, 0xdc, 0xea, 0x29, 0x47, 0x95, 0x7a, 0xdc, 0x24, 0x42, 0xf3, 0x50,
	0x37, 0x71, 0x52, 0x04, 0x27, 0x87, 0x5b, 0x04, 0x17, 0x49, 0x4b, 0x4c, 0x6a, 0xdb, 0xbc, 0xd0,
	0xa4, 0x95, 0xae, 0x7e, 0x05, 0xbf, 0x90, 0xa3, 0xa3, 0xe0, 0x17, 0x90, 0xea, 0x07, 0x91, 0x36,
	0x77, 0xd6, 0xc1, 0xed, 0xe5, 0xf7, 0x7f, 0x2f, 0xef, 0xf7, 0xe8, 0xe4, 0x49, 0x70, 0xa7, 0x44,
	0x55, 0x72, 0x7d, 0x6f, 0x9b, 0x94, 0x99, 0x0a, 0x1d, 0x06, 0x74, 0xa4, 0xe1, 0x9e, 0x44, 0x94,
	0x85, 0x80, 0x21, 0x49, 0xea, 0x07, 0x10, 0xa5, 0x71, 0xad, 0x6f, 0x0c, 0xf7, 0x97, 0x21, 0x37,
	0x19, 0x70, 0xad, 0xd1, 0x71, 0x97, 0xa1, 0xb6, 0xcb, 0x74, 0x22, 0x51, 0xe2, 0x50, 0x42, 0x5f,
	0x79, 0x3a, 0x07, 0xba, 0x7d, 0x2d, 0x78, 0xe1, 0xd4, 0xa5, 0x12, 0x69, 0xbe, 0x10, 0xd6, 0xa0,
	0xb6, 0x22, 0x98, 0xd2, 0x0d, 0x35, 0xe0, 0x76, 0x4a, 0x66, 0xe4, 0x70, 0x73, 0xb1, 0x7a, 0x9e,
	0x34, 0x94, 0xde, 0xfe, 0xfa, 0x04, 0x8a, 0x6e, 0xf9, 0xf1, 0xb4, 0x1f, 0x0f, 0x76, 0x99, 0x57,
	0x60, 0x2b, 0x3f, 0x76, 0xd5, 0xfb, 0x85, 0x07, 0x6c, 0xbc, 0x81, 0xfd, 0xb3, 0x6f, 0x3e, 0x7b,
	0xfe, 0xf8, 0x7e, 0x59, 0x0b, 0xe7, 0x3b, 0xd0, 0x1c, 0xc3, 0xd8, 0x0b, 0x7e, 0xeb, 0x19, 0x39,
	0xba, 0xb8, 0x79, 0xeb, 0x22, 0xf2, 0xde, 0x45, 0xe4, 0xb3, 0x8b, 0xc8, 0xeb, 0x57, 0x44, 0xee,
	0xce, 0x65, 0xe6, 0x54, 0x9d, 0xb0, 0x14, 0x4b, 0x78, 0xe4, 0x65, 0x66, 0x51, 0xab, 0x96, 0x3b,
	0x07, 0xb2, 0x32, 0x69, 0x5c, 0xd6, 0x85, 0xcb, 0x62, 0x93, 0xcb, 0x78, 0x70, 0xb1, 0x60, 0x72,
	0xf9, 0xe7, 0xe7, 0x64, 0x7d, 0xc0, 0xa7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x06, 0x5d,
	0x8f, 0x74, 0x01, 0x00, 0x00,
}
