// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routing/v1alpha1/http_fault.proto

package v1alpha1 // import "istio.io/api/routing/v1alpha1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HTTPFaultInjection can be used to specify one or more faults to inject
// while forwarding http requests to the destination specified in the route
// rule.  Fault specification is part of a route rule. Faults include
// aborting the Http request from downstream service, and/or delaying
// proxying of requests. A fault rule MUST HAVE delay or abort or both.
//
// *Note:* Delay and abort faults are independent of one another, even if
// both are specified simultaneously.
type HTTPFaultInjection struct {
	// Delay requests before forwarding, emulating various failures such as
	// network issues, overloaded upstream service, etc.
	Delay *HTTPFaultInjection_Delay `protobuf:"bytes,1,opt,name=delay" json:"delay,omitempty"`
	// Abort Http request attempts and return error codes back to downstream
	// service, giving the impression that the upstream service is faulty.
	Abort                *HTTPFaultInjection_Abort `protobuf:"bytes,2,opt,name=abort" json:"abort,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *HTTPFaultInjection) Reset()         { *m = HTTPFaultInjection{} }
func (m *HTTPFaultInjection) String() string { return proto.CompactTextString(m) }
func (*HTTPFaultInjection) ProtoMessage()    {}
func (*HTTPFaultInjection) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_fault_60ea0c0365563a95, []int{0}
}
func (m *HTTPFaultInjection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPFaultInjection.Unmarshal(m, b)
}
func (m *HTTPFaultInjection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPFaultInjection.Marshal(b, m, deterministic)
}
func (dst *HTTPFaultInjection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFaultInjection.Merge(dst, src)
}
func (m *HTTPFaultInjection) XXX_Size() int {
	return xxx_messageInfo_HTTPFaultInjection.Size(m)
}
func (m *HTTPFaultInjection) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFaultInjection.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFaultInjection proto.InternalMessageInfo

func (m *HTTPFaultInjection) GetDelay() *HTTPFaultInjection_Delay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFaultInjection) GetAbort() *HTTPFaultInjection_Abort {
	if m != nil {
		return m.Abort
	}
	return nil
}

// Delay specification is used to inject latency into the request
// forwarding path. The following example will introduce a 5 second delay
// in 10% of the requests to the "v1" version of the "reviews"
// service.
//
//     metadata:
//       name: my-rule
//     spec:
//       destination:
//         name: reviews
//       route:
//       - labels:
//           version: v1
//       httpFault:
//         delay:
//           percent: 10
//           fixedDelay: 5s
//
// The _fixedDelay_ field is used to indicate the amount of delay in
// seconds. An optional _percent_ field, a value between 0 and 100, can
// be used to only delay a certain percentage of requests. If left
// unspecified, all request will be delayed.
type HTTPFaultInjection_Delay struct {
	// percentage of requests on which the delay will be injected (0-100)
	Percent float32 `protobuf:"fixed32,1,opt,name=percent" json:"percent,omitempty"`
	// Types that are valid to be assigned to HttpDelayType:
	//	*HTTPFaultInjection_Delay_FixedDelay
	//	*HTTPFaultInjection_Delay_ExponentialDelay
	HttpDelayType isHTTPFaultInjection_Delay_HttpDelayType `protobuf_oneof:"http_delay_type"`
	// (-- Specify delay duration as part of Http request.
	// TODO: The semantics and syntax of the headers is undefined. --)
	OverrideHeaderName   string   `protobuf:"bytes,4,opt,name=override_header_name,json=overrideHeaderName" json:"override_header_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HTTPFaultInjection_Delay) Reset()         { *m = HTTPFaultInjection_Delay{} }
func (m *HTTPFaultInjection_Delay) String() string { return proto.CompactTextString(m) }
func (*HTTPFaultInjection_Delay) ProtoMessage()    {}
func (*HTTPFaultInjection_Delay) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_fault_60ea0c0365563a95, []int{0, 0}
}
func (m *HTTPFaultInjection_Delay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPFaultInjection_Delay.Unmarshal(m, b)
}
func (m *HTTPFaultInjection_Delay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPFaultInjection_Delay.Marshal(b, m, deterministic)
}
func (dst *HTTPFaultInjection_Delay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFaultInjection_Delay.Merge(dst, src)
}
func (m *HTTPFaultInjection_Delay) XXX_Size() int {
	return xxx_messageInfo_HTTPFaultInjection_Delay.Size(m)
}
func (m *HTTPFaultInjection_Delay) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFaultInjection_Delay.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFaultInjection_Delay proto.InternalMessageInfo

type isHTTPFaultInjection_Delay_HttpDelayType interface {
	isHTTPFaultInjection_Delay_HttpDelayType()
}

type HTTPFaultInjection_Delay_FixedDelay struct {
	FixedDelay *duration.Duration `protobuf:"bytes,2,opt,name=fixed_delay,json=fixedDelay,oneof"`
}
type HTTPFaultInjection_Delay_ExponentialDelay struct {
	ExponentialDelay *duration.Duration `protobuf:"bytes,3,opt,name=exponential_delay,json=exponentialDelay,oneof"`
}

func (*HTTPFaultInjection_Delay_FixedDelay) isHTTPFaultInjection_Delay_HttpDelayType()       {}
func (*HTTPFaultInjection_Delay_ExponentialDelay) isHTTPFaultInjection_Delay_HttpDelayType() {}

func (m *HTTPFaultInjection_Delay) GetHttpDelayType() isHTTPFaultInjection_Delay_HttpDelayType {
	if m != nil {
		return m.HttpDelayType
	}
	return nil
}

func (m *HTTPFaultInjection_Delay) GetPercent() float32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *HTTPFaultInjection_Delay) GetFixedDelay() *duration.Duration {
	if x, ok := m.GetHttpDelayType().(*HTTPFaultInjection_Delay_FixedDelay); ok {
		return x.FixedDelay
	}
	return nil
}

func (m *HTTPFaultInjection_Delay) GetExponentialDelay() *duration.Duration {
	if x, ok := m.GetHttpDelayType().(*HTTPFaultInjection_Delay_ExponentialDelay); ok {
		return x.ExponentialDelay
	}
	return nil
}

func (m *HTTPFaultInjection_Delay) GetOverrideHeaderName() string {
	if m != nil {
		return m.OverrideHeaderName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HTTPFaultInjection_Delay) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HTTPFaultInjection_Delay_OneofMarshaler, _HTTPFaultInjection_Delay_OneofUnmarshaler, _HTTPFaultInjection_Delay_OneofSizer, []interface{}{
		(*HTTPFaultInjection_Delay_FixedDelay)(nil),
		(*HTTPFaultInjection_Delay_ExponentialDelay)(nil),
	}
}

func _HTTPFaultInjection_Delay_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HTTPFaultInjection_Delay)
	// http_delay_type
	switch x := m.HttpDelayType.(type) {
	case *HTTPFaultInjection_Delay_FixedDelay:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FixedDelay); err != nil {
			return err
		}
	case *HTTPFaultInjection_Delay_ExponentialDelay:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExponentialDelay); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HTTPFaultInjection_Delay.HttpDelayType has unexpected type %T", x)
	}
	return nil
}

func _HTTPFaultInjection_Delay_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HTTPFaultInjection_Delay)
	switch tag {
	case 2: // http_delay_type.fixed_delay
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(duration.Duration)
		err := b.DecodeMessage(msg)
		m.HttpDelayType = &HTTPFaultInjection_Delay_FixedDelay{msg}
		return true, err
	case 3: // http_delay_type.exponential_delay
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(duration.Duration)
		err := b.DecodeMessage(msg)
		m.HttpDelayType = &HTTPFaultInjection_Delay_ExponentialDelay{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HTTPFaultInjection_Delay_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HTTPFaultInjection_Delay)
	// http_delay_type
	switch x := m.HttpDelayType.(type) {
	case *HTTPFaultInjection_Delay_FixedDelay:
		s := proto.Size(x.FixedDelay)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HTTPFaultInjection_Delay_ExponentialDelay:
		s := proto.Size(x.ExponentialDelay)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Abort specification is used to prematurely abort a request with a
// pre-specified error code. The following example will return an HTTP
// 400 error code for 10% of the requests to the "ratings" service "v1".
//
//     metadata:
//       name: my-rule
//     spec:
//       destination:
//         name: reviews
//       route:
//       - labels:
//           version: v1
//       httpFault:
//         abort:
//           percent: 10
//           httpStatus: 400
//
// The _httpStatus_ field is used to indicate the HTTP status code to
// return to the caller. The optional _percent_ field, a value between 0
// and 100, is used to only abort a certain percentage of requests. If
// not specified, all requests are aborted.
type HTTPFaultInjection_Abort struct {
	// percentage of requests to be aborted with the error code provided (0-100).
	Percent float32 `protobuf:"fixed32,1,opt,name=percent" json:"percent,omitempty"`
	// Types that are valid to be assigned to ErrorType:
	//	*HTTPFaultInjection_Abort_GrpcStatus
	//	*HTTPFaultInjection_Abort_Http2Error
	//	*HTTPFaultInjection_Abort_HttpStatus
	ErrorType isHTTPFaultInjection_Abort_ErrorType `protobuf_oneof:"error_type"`
	// (-- Specify abort code as part of Http request.
	// TODO: The semantics and syntax of the headers is undefined. --)
	OverrideHeaderName   string   `protobuf:"bytes,5,opt,name=override_header_name,json=overrideHeaderName" json:"override_header_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HTTPFaultInjection_Abort) Reset()         { *m = HTTPFaultInjection_Abort{} }
func (m *HTTPFaultInjection_Abort) String() string { return proto.CompactTextString(m) }
func (*HTTPFaultInjection_Abort) ProtoMessage()    {}
func (*HTTPFaultInjection_Abort) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_fault_60ea0c0365563a95, []int{0, 1}
}
func (m *HTTPFaultInjection_Abort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPFaultInjection_Abort.Unmarshal(m, b)
}
func (m *HTTPFaultInjection_Abort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPFaultInjection_Abort.Marshal(b, m, deterministic)
}
func (dst *HTTPFaultInjection_Abort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFaultInjection_Abort.Merge(dst, src)
}
func (m *HTTPFaultInjection_Abort) XXX_Size() int {
	return xxx_messageInfo_HTTPFaultInjection_Abort.Size(m)
}
func (m *HTTPFaultInjection_Abort) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFaultInjection_Abort.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFaultInjection_Abort proto.InternalMessageInfo

type isHTTPFaultInjection_Abort_ErrorType interface {
	isHTTPFaultInjection_Abort_ErrorType()
}

type HTTPFaultInjection_Abort_GrpcStatus struct {
	GrpcStatus string `protobuf:"bytes,2,opt,name=grpc_status,json=grpcStatus,oneof"`
}
type HTTPFaultInjection_Abort_Http2Error struct {
	Http2Error string `protobuf:"bytes,3,opt,name=http2_error,json=http2Error,oneof"`
}
type HTTPFaultInjection_Abort_HttpStatus struct {
	HttpStatus int32 `protobuf:"varint,4,opt,name=http_status,json=httpStatus,oneof"`
}

func (*HTTPFaultInjection_Abort_GrpcStatus) isHTTPFaultInjection_Abort_ErrorType() {}
func (*HTTPFaultInjection_Abort_Http2Error) isHTTPFaultInjection_Abort_ErrorType() {}
func (*HTTPFaultInjection_Abort_HttpStatus) isHTTPFaultInjection_Abort_ErrorType() {}

func (m *HTTPFaultInjection_Abort) GetErrorType() isHTTPFaultInjection_Abort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *HTTPFaultInjection_Abort) GetPercent() float32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *HTTPFaultInjection_Abort) GetGrpcStatus() string {
	if x, ok := m.GetErrorType().(*HTTPFaultInjection_Abort_GrpcStatus); ok {
		return x.GrpcStatus
	}
	return ""
}

func (m *HTTPFaultInjection_Abort) GetHttp2Error() string {
	if x, ok := m.GetErrorType().(*HTTPFaultInjection_Abort_Http2Error); ok {
		return x.Http2Error
	}
	return ""
}

func (m *HTTPFaultInjection_Abort) GetHttpStatus() int32 {
	if x, ok := m.GetErrorType().(*HTTPFaultInjection_Abort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (m *HTTPFaultInjection_Abort) GetOverrideHeaderName() string {
	if m != nil {
		return m.OverrideHeaderName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HTTPFaultInjection_Abort) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HTTPFaultInjection_Abort_OneofMarshaler, _HTTPFaultInjection_Abort_OneofUnmarshaler, _HTTPFaultInjection_Abort_OneofSizer, []interface{}{
		(*HTTPFaultInjection_Abort_GrpcStatus)(nil),
		(*HTTPFaultInjection_Abort_Http2Error)(nil),
		(*HTTPFaultInjection_Abort_HttpStatus)(nil),
	}
}

func _HTTPFaultInjection_Abort_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HTTPFaultInjection_Abort)
	// error_type
	switch x := m.ErrorType.(type) {
	case *HTTPFaultInjection_Abort_GrpcStatus:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.GrpcStatus)
	case *HTTPFaultInjection_Abort_Http2Error:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Http2Error)
	case *HTTPFaultInjection_Abort_HttpStatus:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.HttpStatus))
	case nil:
	default:
		return fmt.Errorf("HTTPFaultInjection_Abort.ErrorType has unexpected type %T", x)
	}
	return nil
}

func _HTTPFaultInjection_Abort_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HTTPFaultInjection_Abort)
	switch tag {
	case 2: // error_type.grpc_status
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ErrorType = &HTTPFaultInjection_Abort_GrpcStatus{x}
		return true, err
	case 3: // error_type.http2_error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ErrorType = &HTTPFaultInjection_Abort_Http2Error{x}
		return true, err
	case 4: // error_type.http_status
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ErrorType = &HTTPFaultInjection_Abort_HttpStatus{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _HTTPFaultInjection_Abort_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HTTPFaultInjection_Abort)
	// error_type
	switch x := m.ErrorType.(type) {
	case *HTTPFaultInjection_Abort_GrpcStatus:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.GrpcStatus)))
		n += len(x.GrpcStatus)
	case *HTTPFaultInjection_Abort_Http2Error:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Http2Error)))
		n += len(x.Http2Error)
	case *HTTPFaultInjection_Abort_HttpStatus:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.HttpStatus))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*HTTPFaultInjection)(nil), "istio.routing.v1alpha1.HTTPFaultInjection")
	proto.RegisterType((*HTTPFaultInjection_Delay)(nil), "istio.routing.v1alpha1.HTTPFaultInjection.Delay")
	proto.RegisterType((*HTTPFaultInjection_Abort)(nil), "istio.routing.v1alpha1.HTTPFaultInjection.Abort")
}

func init() {
	proto.RegisterFile("routing/v1alpha1/http_fault.proto", fileDescriptor_http_fault_60ea0c0365563a95)
}

var fileDescriptor_http_fault_60ea0c0365563a95 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x6f, 0xb9, 0xf4, 0xde, 0x30, 0xdc, 0xe4, 0x4a, 0x63, 0x4c, 0x25, 0x51, 0xc1, 0x15,
	0xab, 0x29, 0xe0, 0xd6, 0x8d, 0x04, 0x49, 0xdd, 0x18, 0x53, 0x59, 0xb9, 0x69, 0x06, 0x7a, 0x28,
	0x63, 0x4a, 0x67, 0x32, 0x4c, 0x09, 0xbc, 0xa2, 0x8f, 0xe1, 0x03, 0xf8, 0x0c, 0x66, 0xce, 0xb4,
	0x89, 0x51, 0xd1, 0xb8, 0xec, 0x39, 0xff, 0xff, 0x9d, 0xf3, 0x9f, 0x29, 0xe9, 0x2a, 0x51, 0x68,
	0x9e, 0xa7, 0xc1, 0x66, 0xc0, 0x32, 0xb9, 0x64, 0x83, 0x60, 0xa9, 0xb5, 0x8c, 0x17, 0xac, 0xc8,
	0x34, 0x95, 0x4a, 0x68, 0xe1, 0x1d, 0xf1, 0xb5, 0xe6, 0x82, 0x96, 0x42, 0x5a, 0x09, 0xdb, 0xa7,
	0xa9, 0x10, 0x69, 0x06, 0x01, 0xaa, 0x66, 0xc5, 0x22, 0x48, 0x0a, 0xc5, 0x34, 0x17, 0xb9, 0xf5,
	0x9d, 0x3f, 0xd7, 0x89, 0x17, 0x4e, 0xa7, 0x77, 0x13, 0xc3, 0xba, 0xc9, 0x1f, 0x61, 0x6e, 0x9a,
	0xde, 0x84, 0xb8, 0x09, 0x64, 0x6c, 0xe7, 0x3b, 0x1d, 0xa7, 0xd7, 0x1c, 0xf6, 0xe9, 0xe7, 0x78,
	0xfa, 0xd1, 0x4a, 0xc7, 0xc6, 0x17, 0x59, 0xbb, 0xe1, 0xb0, 0x99, 0x50, 0xda, 0xaf, 0xfd, 0x98,
	0x73, 0x65, 0x7c, 0x91, 0xb5, 0xb7, 0x5f, 0x1c, 0xe2, 0x22, 0xd8, 0xf3, 0xc9, 0x5f, 0x09, 0x6a,
	0x0e, 0xb9, 0xc6, 0xdd, 0x6a, 0x51, 0xf5, 0xe9, 0x5d, 0x92, 0xe6, 0x82, 0x6f, 0x21, 0x89, 0xed,
	0xe6, 0x76, 0xe2, 0x31, 0xb5, 0x07, 0xa0, 0xd5, 0x01, 0xe8, 0xb8, 0x3c, 0x40, 0xf8, 0x2b, 0x22,
	0xa8, 0xb7, 0xdc, 0x90, 0xb4, 0x60, 0x2b, 0x45, 0x0e, 0xb9, 0xe6, 0x2c, 0x2b, 0x19, 0xbf, 0xbf,
	0x67, 0x1c, 0xbc, 0x71, 0x59, 0x52, 0x9f, 0x1c, 0x8a, 0x0d, 0x28, 0xc5, 0x13, 0x88, 0x97, 0xc0,
	0x12, 0x50, 0x71, 0xce, 0x56, 0xe0, 0xd7, 0x3b, 0x4e, 0xaf, 0x11, 0x79, 0x55, 0x2f, 0xc4, 0xd6,
	0x2d, 0x5b, 0xc1, 0xa8, 0x45, 0xfe, 0xe3, 0x83, 0xe2, 0xd0, 0x58, 0xef, 0x24, 0xb4, 0x9f, 0x1c,
	0xe2, 0xe2, 0x05, 0xbe, 0x08, 0xdc, 0x25, 0xcd, 0x54, 0xc9, 0x79, 0xbc, 0xd6, 0x4c, 0x17, 0x6b,
	0x0c, 0xdc, 0x30, 0xa9, 0x4c, 0xf1, 0x1e, 0x6b, 0x46, 0x62, 0xc8, 0xc3, 0x18, 0x94, 0x12, 0x0a,
	0xf3, 0xa0, 0x04, 0x8b, 0xd7, 0xa6, 0x56, 0x49, 0x2a, 0x8a, 0xd9, 0xd2, 0xad, 0x24, 0x25, 0x65,
	0x5f, 0x22, 0x77, 0x6f, 0xa2, 0x7f, 0x84, 0xe0, 0x44, 0x0c, 0x33, 0x3a, 0x7b, 0x38, 0xb1, 0xef,
	0xce, 0x45, 0xc0, 0x24, 0x0f, 0xde, 0xff, 0xce, 0xb3, 0x3f, 0x78, 0xd9, 0x8b, 0xd7, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x5c, 0x51, 0x51, 0x18, 0xe9, 0x02, 0x00, 0x00,
}