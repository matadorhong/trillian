// Code generated by protoc-gen-go.
// source: quotapb.proto
// DO NOT EDIT!

/*
Package quotapb is a generated protocol buffer package.

Package quotapb contains definitions for quota API protos and RPC service.
TODO(codingllama): Add a brief explanation of how the quota system works and
example configurations.

It is generated from these files:
	quotapb.proto

It has these top-level messages:
	Config
	SequencingBasedStrategy
	TimeBasedStrategy
	CreateConfigRequest
	DeleteConfigRequest
	GetConfigRequest
	ListConfigRequest
	ListConfigResponse
	UpdateConfigRequest
*/
package quotapb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf2 "google.golang.org/genproto/protobuf/field_mask"

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

// Possible states of a quota configuration.
type Config_State int32

const (
	// Unknown quota state. Invalid.
	Config_UNKNOWN_CONFIG_STATE Config_State = 0
	// Quota is enabled.
	Config_ENABLED Config_State = 1
	// Quota is disabled (considered infinite).
	Config_DISABLED Config_State = 2
)

var Config_State_name = map[int32]string{
	0: "UNKNOWN_CONFIG_STATE",
	1: "ENABLED",
	2: "DISABLED",
}
var Config_State_value = map[string]int32{
	"UNKNOWN_CONFIG_STATE": 0,
	"ENABLED":              1,
	"DISABLED":             2,
}

func (x Config_State) String() string {
	return proto.EnumName(Config_State_name, int32(x))
}
func (Config_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Possible views for ListConfig.
type ListConfigRequest_ListView int32

const (
	// Only the Config name gets returned.
	ListConfigRequest_BASIC ListConfigRequest_ListView = 0
	// Complete Config.
	ListConfigRequest_FULL ListConfigRequest_ListView = 1
)

var ListConfigRequest_ListView_name = map[int32]string{
	0: "BASIC",
	1: "FULL",
}
var ListConfigRequest_ListView_value = map[string]int32{
	"BASIC": 0,
	"FULL":  1,
}

func (x ListConfigRequest_ListView) String() string {
	return proto.EnumName(ListConfigRequest_ListView_name, int32(x))
}
func (ListConfigRequest_ListView) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

// Configuration of a quota.
type Config struct {
	// Name of the config, eg, “quotas/trees/1234/read/config”.
	// Readonly.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// State of the config.
	State Config_State `protobuf:"varint,2,opt,name=state,enum=quotapb.Config_State" json:"state,omitempty"`
	// Max number of tokens available for the config.
	MaxTokens int64 `protobuf:"varint,3,opt,name=max_tokens,json=maxTokens" json:"max_tokens,omitempty"`
	// Replenishment strategy used by the config.
	//
	// Types that are valid to be assigned to ReplenishmentStrategy:
	//	*Config_SequencingBased
	//	*Config_TimeBased
	ReplenishmentStrategy isConfig_ReplenishmentStrategy `protobuf_oneof:"replenishment_strategy"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isConfig_ReplenishmentStrategy interface {
	isConfig_ReplenishmentStrategy()
}

type Config_SequencingBased struct {
	SequencingBased *SequencingBasedStrategy `protobuf:"bytes,4,opt,name=sequencing_based,json=sequencingBased,oneof"`
}
type Config_TimeBased struct {
	TimeBased *TimeBasedStrategy `protobuf:"bytes,5,opt,name=time_based,json=timeBased,oneof"`
}

func (*Config_SequencingBased) isConfig_ReplenishmentStrategy() {}
func (*Config_TimeBased) isConfig_ReplenishmentStrategy()       {}

func (m *Config) GetReplenishmentStrategy() isConfig_ReplenishmentStrategy {
	if m != nil {
		return m.ReplenishmentStrategy
	}
	return nil
}

func (m *Config) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Config) GetState() Config_State {
	if m != nil {
		return m.State
	}
	return Config_UNKNOWN_CONFIG_STATE
}

func (m *Config) GetMaxTokens() int64 {
	if m != nil {
		return m.MaxTokens
	}
	return 0
}

func (m *Config) GetSequencingBased() *SequencingBasedStrategy {
	if x, ok := m.GetReplenishmentStrategy().(*Config_SequencingBased); ok {
		return x.SequencingBased
	}
	return nil
}

func (m *Config) GetTimeBased() *TimeBasedStrategy {
	if x, ok := m.GetReplenishmentStrategy().(*Config_TimeBased); ok {
		return x.TimeBased
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Config) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Config_OneofMarshaler, _Config_OneofUnmarshaler, _Config_OneofSizer, []interface{}{
		(*Config_SequencingBased)(nil),
		(*Config_TimeBased)(nil),
	}
}

func _Config_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Config)
	// replenishment_strategy
	switch x := m.ReplenishmentStrategy.(type) {
	case *Config_SequencingBased:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SequencingBased); err != nil {
			return err
		}
	case *Config_TimeBased:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimeBased); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Config.ReplenishmentStrategy has unexpected type %T", x)
	}
	return nil
}

func _Config_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Config)
	switch tag {
	case 4: // replenishment_strategy.sequencing_based
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SequencingBasedStrategy)
		err := b.DecodeMessage(msg)
		m.ReplenishmentStrategy = &Config_SequencingBased{msg}
		return true, err
	case 5: // replenishment_strategy.time_based
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TimeBasedStrategy)
		err := b.DecodeMessage(msg)
		m.ReplenishmentStrategy = &Config_TimeBased{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Config_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Config)
	// replenishment_strategy
	switch x := m.ReplenishmentStrategy.(type) {
	case *Config_SequencingBased:
		s := proto.Size(x.SequencingBased)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Config_TimeBased:
		s := proto.Size(x.TimeBased)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Sequencing-based replenishment strategy settings.
// Only global and tree quotas may use sequencing-based replenishment.
type SequencingBasedStrategy struct {
}

func (m *SequencingBasedStrategy) Reset()                    { *m = SequencingBasedStrategy{} }
func (m *SequencingBasedStrategy) String() string            { return proto.CompactTextString(m) }
func (*SequencingBasedStrategy) ProtoMessage()               {}
func (*SequencingBasedStrategy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Time-based replenishment strategy settings.
type TimeBasedStrategy struct {
	// Number of tokens to replenish at every replenish_interval_seconds.
	TokensToReplenish int64 `protobuf:"varint,1,opt,name=tokens_to_replenish,json=tokensToReplenish" json:"tokens_to_replenish,omitempty"`
	// Interval at which tokens_to_replenish get replenished.
	ReplenishIntervalSeconds int64 `protobuf:"varint,2,opt,name=replenish_interval_seconds,json=replenishIntervalSeconds" json:"replenish_interval_seconds,omitempty"`
}

func (m *TimeBasedStrategy) Reset()                    { *m = TimeBasedStrategy{} }
func (m *TimeBasedStrategy) String() string            { return proto.CompactTextString(m) }
func (*TimeBasedStrategy) ProtoMessage()               {}
func (*TimeBasedStrategy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TimeBasedStrategy) GetTokensToReplenish() int64 {
	if m != nil {
		return m.TokensToReplenish
	}
	return 0
}

func (m *TimeBasedStrategy) GetReplenishIntervalSeconds() int64 {
	if m != nil {
		return m.ReplenishIntervalSeconds
	}
	return 0
}

// CreateConfig request.
type CreateConfigRequest struct {
	// Name of the config to create.
	// For example, "quotas/global/read/config" (global/read quota) or
	// "quotas/trees/1234/write/config" (write quota for tree 1234).
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Config to be created.
	Config *Config `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
}

func (m *CreateConfigRequest) Reset()                    { *m = CreateConfigRequest{} }
func (m *CreateConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateConfigRequest) ProtoMessage()               {}
func (*CreateConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateConfigRequest) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

// DeleteConfig request.
type DeleteConfigRequest struct {
	// Name of the config to delete.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteConfigRequest) Reset()                    { *m = DeleteConfigRequest{} }
func (m *DeleteConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteConfigRequest) ProtoMessage()               {}
func (*DeleteConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// GetConfig request.
type GetConfigRequest struct {
	// Name of the config to retrieve.
	// For example, "quotas/global/read/config".
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetConfigRequest) Reset()                    { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()               {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// ListConfig request.
type ListConfigRequest struct {
	// Name of the config to retrieve. For example, "quotas/global/read/config".
	// If empty, all configs are listed.
	// Name components may be substituted by "-" to search for all variations of
	// that component. For example:
	// - "quotas/global/-/config" (both read and write global quotas)
	// - "quotas/trees/-/-/config" (all tree quotas)
	Name []string `protobuf:"bytes,1,rep,name=name" json:"name,omitempty"`
	// View specifies how much data to return.
	View ListConfigRequest_ListView `protobuf:"varint,2,opt,name=view,enum=quotapb.ListConfigRequest_ListView" json:"view,omitempty"`
}

func (m *ListConfigRequest) Reset()                    { *m = ListConfigRequest{} }
func (m *ListConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ListConfigRequest) ProtoMessage()               {}
func (*ListConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListConfigRequest) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ListConfigRequest) GetView() ListConfigRequest_ListView {
	if m != nil {
		return m.View
	}
	return ListConfigRequest_BASIC
}

// ListConfig response.
type ListConfigResponse struct {
	// Configs matching the request filter.
	Config []*Config `protobuf:"bytes,1,rep,name=config" json:"config,omitempty"`
}

func (m *ListConfigResponse) Reset()                    { *m = ListConfigResponse{} }
func (m *ListConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ListConfigResponse) ProtoMessage()               {}
func (*ListConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListConfigResponse) GetConfig() []*Config {
	if m != nil {
		return m.Config
	}
	return nil
}

// UpdateConfig request.
type UpdateConfigRequest struct {
	// Name of the config to update.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Config to update. Only the fields specified by update_mask need to be
	// filled.
	Config *Config `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
	// Fields modified by the update request.
	// For example: "state" or "max_tokens".
	UpdateMask *google_protobuf2.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdateConfigRequest) Reset()                    { *m = UpdateConfigRequest{} }
func (m *UpdateConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateConfigRequest) ProtoMessage()               {}
func (*UpdateConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateConfigRequest) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *UpdateConfigRequest) GetUpdateMask() *google_protobuf2.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "quotapb.Config")
	proto.RegisterType((*SequencingBasedStrategy)(nil), "quotapb.SequencingBasedStrategy")
	proto.RegisterType((*TimeBasedStrategy)(nil), "quotapb.TimeBasedStrategy")
	proto.RegisterType((*CreateConfigRequest)(nil), "quotapb.CreateConfigRequest")
	proto.RegisterType((*DeleteConfigRequest)(nil), "quotapb.DeleteConfigRequest")
	proto.RegisterType((*GetConfigRequest)(nil), "quotapb.GetConfigRequest")
	proto.RegisterType((*ListConfigRequest)(nil), "quotapb.ListConfigRequest")
	proto.RegisterType((*ListConfigResponse)(nil), "quotapb.ListConfigResponse")
	proto.RegisterType((*UpdateConfigRequest)(nil), "quotapb.UpdateConfigRequest")
	proto.RegisterEnum("quotapb.Config_State", Config_State_name, Config_State_value)
	proto.RegisterEnum("quotapb.ListConfigRequest_ListView", ListConfigRequest_ListView_name, ListConfigRequest_ListView_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Quota service

type QuotaClient interface {
	// Creates a new quota.
	CreateConfig(ctx context.Context, in *CreateConfigRequest, opts ...grpc.CallOption) (*Config, error)
	// Deletes an existing quota. Non-existing quotas are considered infinite by
	// the quota system.
	DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Retrieves a quota by name.
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*Config, error)
	// Lists quotas according to the specified criteria.
	ListConfig(ctx context.Context, in *ListConfigRequest, opts ...grpc.CallOption) (*ListConfigResponse, error)
	// Updates a quota.
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*Config, error)
}

type quotaClient struct {
	cc *grpc.ClientConn
}

func NewQuotaClient(cc *grpc.ClientConn) QuotaClient {
	return &quotaClient{cc}
}

func (c *quotaClient) CreateConfig(ctx context.Context, in *CreateConfigRequest, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := grpc.Invoke(ctx, "/quotapb.Quota/CreateConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaClient) DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/quotapb.Quota/DeleteConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := grpc.Invoke(ctx, "/quotapb.Quota/GetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaClient) ListConfig(ctx context.Context, in *ListConfigRequest, opts ...grpc.CallOption) (*ListConfigResponse, error) {
	out := new(ListConfigResponse)
	err := grpc.Invoke(ctx, "/quotapb.Quota/ListConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := grpc.Invoke(ctx, "/quotapb.Quota/UpdateConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Quota service

type QuotaServer interface {
	// Creates a new quota.
	CreateConfig(context.Context, *CreateConfigRequest) (*Config, error)
	// Deletes an existing quota. Non-existing quotas are considered infinite by
	// the quota system.
	DeleteConfig(context.Context, *DeleteConfigRequest) (*google_protobuf1.Empty, error)
	// Retrieves a quota by name.
	GetConfig(context.Context, *GetConfigRequest) (*Config, error)
	// Lists quotas according to the specified criteria.
	ListConfig(context.Context, *ListConfigRequest) (*ListConfigResponse, error)
	// Updates a quota.
	UpdateConfig(context.Context, *UpdateConfigRequest) (*Config, error)
}

func RegisterQuotaServer(s *grpc.Server, srv QuotaServer) {
	s.RegisterService(&_Quota_serviceDesc, srv)
}

func _Quota_CreateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServer).CreateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotapb.Quota/CreateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServer).CreateConfig(ctx, req.(*CreateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quota_DeleteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServer).DeleteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotapb.Quota/DeleteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServer).DeleteConfig(ctx, req.(*DeleteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quota_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotapb.Quota/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quota_ListConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServer).ListConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotapb.Quota/ListConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServer).ListConfig(ctx, req.(*ListConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quota_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotapb.Quota/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServer).UpdateConfig(ctx, req.(*UpdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Quota_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quotapb.Quota",
	HandlerType: (*QuotaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConfig",
			Handler:    _Quota_CreateConfig_Handler,
		},
		{
			MethodName: "DeleteConfig",
			Handler:    _Quota_DeleteConfig_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _Quota_GetConfig_Handler,
		},
		{
			MethodName: "ListConfig",
			Handler:    _Quota_ListConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _Quota_UpdateConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quotapb.proto",
}

func init() { proto.RegisterFile("quotapb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5d, 0x4f, 0x13, 0x4d,
	0x18, 0x65, 0xfb, 0x01, 0xf4, 0x29, 0xef, 0x4b, 0x99, 0xbe, 0x2f, 0x94, 0x05, 0x63, 0x33, 0x46,
	0xad, 0x25, 0x69, 0x43, 0xbd, 0x30, 0x11, 0xb8, 0xa0, 0xa5, 0x60, 0x63, 0x29, 0x71, 0x5b, 0xf4,
	0xce, 0xcd, 0x94, 0x0e, 0x75, 0xa5, 0x3b, 0xbb, 0x74, 0xa6, 0x7c, 0xc4, 0x78, 0xa1, 0x7f, 0xc0,
	0x44, 0xff, 0x96, 0x77, 0xfe, 0x05, 0x7f, 0x88, 0xd9, 0x99, 0xed, 0xb2, 0xf6, 0x43, 0xb8, 0xf0,
	0x6e, 0x67, 0xce, 0x99, 0x73, 0x9e, 0x79, 0x9e, 0x33, 0x0b, 0xff, 0x9c, 0x0f, 0x1c, 0x41, 0xdc,
	0x76, 0xc1, 0xed, 0x3b, 0xc2, 0x41, 0x73, 0xfe, 0x52, 0x5f, 0xef, 0x3a, 0x4e, 0xb7, 0x47, 0x8b,
	0xc4, 0xb5, 0x8a, 0x84, 0x31, 0x47, 0x10, 0x61, 0x39, 0x8c, 0x2b, 0x9a, 0xbe, 0xe6, 0xa3, 0x72,
	0xd5, 0x1e, 0x9c, 0x16, 0xa9, 0xed, 0x8a, 0x6b, 0x1f, 0xcc, 0x8e, 0x82, 0xa7, 0x16, 0xed, 0x75,
	0x4c, 0x9b, 0xf0, 0x33, 0xc5, 0xc0, 0xdf, 0x23, 0x30, 0x5b, 0x71, 0xd8, 0xa9, 0xd5, 0x45, 0x08,
	0x62, 0x8c, 0xd8, 0x34, 0xa3, 0x65, 0xb5, 0x5c, 0xc2, 0x90, 0xdf, 0x68, 0x03, 0xe2, 0x5c, 0x10,
	0x41, 0x33, 0x91, 0xac, 0x96, 0xfb, 0xb7, 0xf4, 0x7f, 0x61, 0x58, 0xa3, 0x3a, 0x53, 0x68, 0x7a,
	0xa0, 0xa1, 0x38, 0xe8, 0x1e, 0x80, 0x4d, 0xae, 0x4c, 0xe1, 0x9c, 0x51, 0xc6, 0x33, 0xd1, 0xac,
	0x96, 0x8b, 0x1a, 0x09, 0x9b, 0x5c, 0xb5, 0xe4, 0x06, 0x3a, 0x84, 0x14, 0xa7, 0xe7, 0x03, 0xca,
	0x4e, 0x2c, 0xd6, 0x35, 0xdb, 0x84, 0xd3, 0x4e, 0x26, 0x96, 0xd5, 0x72, 0xc9, 0x52, 0x36, 0x90,
	0x6d, 0x06, 0x84, 0xb2, 0x87, 0x37, 0x45, 0x9f, 0x08, 0xda, 0xbd, 0x7e, 0x31, 0x63, 0x2c, 0xf2,
	0xdf, 0x21, 0xb4, 0x05, 0x20, 0x2c, 0x9b, 0xfa, 0x42, 0x71, 0x29, 0xa4, 0x07, 0x42, 0x2d, 0xcb,
	0xa6, 0xa3, 0x12, 0x09, 0x31, 0xdc, 0xc4, 0xdb, 0x10, 0x97, 0xa5, 0xa3, 0x0c, 0xfc, 0x77, 0xdc,
	0x78, 0xd9, 0x38, 0x7a, 0xd3, 0x30, 0x2b, 0x47, 0x8d, 0xfd, 0xda, 0x81, 0xd9, 0x6c, 0xed, 0xb6,
	0xaa, 0xa9, 0x19, 0x94, 0x84, 0xb9, 0x6a, 0x63, 0xb7, 0x5c, 0xaf, 0xee, 0xa5, 0x34, 0xb4, 0x00,
	0xf3, 0x7b, 0xb5, 0xa6, 0x5a, 0x45, 0xca, 0x19, 0x58, 0xee, 0x53, 0xb7, 0x47, 0x99, 0xc5, 0xdf,
	0xd9, 0x94, 0x09, 0x93, 0xfb, 0x26, 0x78, 0x15, 0x56, 0xa6, 0x5c, 0x01, 0x7f, 0xd2, 0x60, 0x69,
	0xac, 0x2a, 0x54, 0x80, 0xb4, 0xea, 0x97, 0x29, 0x1c, 0x33, 0x10, 0x95, 0x33, 0x88, 0x1a, 0x4b,
	0x0a, 0x6a, 0x39, 0xc6, 0x10, 0x40, 0xdb, 0xa0, 0x07, 0x2c, 0xd3, 0x62, 0x82, 0xf6, 0x2f, 0x48,
	0xcf, 0xe4, 0xf4, 0xc4, 0x61, 0x1d, 0x2e, 0xa7, 0x14, 0x35, 0x32, 0x01, 0xa3, 0xe6, 0x13, 0x9a,
	0x0a, 0xc7, 0x06, 0xa4, 0x2b, 0x7d, 0x4a, 0x04, 0x55, 0xe3, 0x33, 0xbc, 0x52, 0xb9, 0x98, 0x38,
	0xf9, 0xc7, 0x30, 0x7b, 0x22, 0x49, 0x52, 0x34, 0x59, 0x5a, 0x1c, 0x19, 0xbd, 0xe1, 0xc3, 0xf8,
	0x09, 0xa4, 0xf7, 0x68, 0x8f, 0xde, 0x41, 0x13, 0x3f, 0x82, 0xd4, 0x01, 0x15, 0xb7, 0xf3, 0xbc,
	0x56, 0xd5, 0x2d, 0x3e, 0x95, 0x19, 0x0d, 0xaa, 0x7c, 0x06, 0xb1, 0x0b, 0x8b, 0x5e, 0xfa, 0xf1,
	0x7c, 0x10, 0xd4, 0x38, 0x76, 0x5a, 0xee, 0xbc, 0xb6, 0xe8, 0xa5, 0x21, 0x0f, 0xe0, 0xfb, 0x30,
	0x3f, 0xdc, 0x41, 0x09, 0x88, 0x97, 0x77, 0x9b, 0xb5, 0x4a, 0x6a, 0x06, 0xcd, 0x43, 0x6c, 0xff,
	0xb8, 0x5e, 0x4f, 0x69, 0x78, 0x07, 0x50, 0x58, 0x84, 0xbb, 0x0e, 0xe3, 0xe1, 0xae, 0x78, 0x55,
	0xfc, 0xa1, 0x2b, 0x5f, 0x34, 0x48, 0x1f, 0xbb, 0x9d, 0xbf, 0xda, 0x6a, 0xb4, 0x05, 0xc9, 0x81,
	0xd4, 0x94, 0x2f, 0x58, 0xbe, 0x30, 0x2f, 0xf3, 0xea, 0x91, 0x17, 0x86, 0x8f, 0xbc, 0xb0, 0xef,
	0x3d, 0xf2, 0x43, 0xc2, 0xcf, 0x0c, 0x50, 0x74, 0xef, 0xbb, 0xf4, 0x35, 0x06, 0xf1, 0x57, 0x9e,
	0x2e, 0x7a, 0x0f, 0x0b, 0xe1, 0x14, 0xa0, 0xf5, 0x1b, 0xbf, 0xf1, 0x70, 0xe8, 0xa3, 0xd5, 0xe0,
	0x8d, 0xcf, 0x3f, 0x7e, 0x7e, 0x8b, 0x3c, 0xc4, 0xd9, 0xe2, 0xc5, 0x66, 0x9b, 0x0a, 0xb2, 0x59,
	0xfc, 0xe0, 0x5d, 0x63, 0x47, 0xd2, 0x78, 0x31, 0x9f, 0x2f, 0xaa, 0x72, 0x3f, 0x3e, 0xd7, 0xf2,
	0x88, 0xc1, 0x42, 0x38, 0x1d, 0x21, 0xaf, 0x09, 0xa1, 0xd1, 0x97, 0xc7, 0xee, 0x52, 0xf5, 0xfe,
	0x66, 0x38, 0x27, 0x2d, 0x71, 0xfe, 0x56, 0x4b, 0x44, 0x20, 0x11, 0x44, 0x0c, 0xad, 0x06, 0x66,
	0xa3, 0xb1, 0x1b, 0xbf, 0x95, 0x6f, 0x81, 0x6e, 0xb7, 0x78, 0x0b, 0x70, 0x93, 0x0c, 0xa4, 0x4f,
	0xcf, 0x9c, 0xbe, 0x36, 0x11, 0x53, 0x51, 0xc2, 0x2b, 0xd2, 0x70, 0x09, 0x2d, 0x06, 0x86, 0xca,
	0xca, 0x1b, 0x4f, 0x38, 0x39, 0xa1, 0x96, 0x4d, 0x08, 0xd4, 0xd4, 0xf1, 0x94, 0xee, 0x32, 0x9e,
	0xf6, 0xac, 0x6c, 0xf4, 0xd3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xe5, 0xd5, 0x36, 0x7c,
	0x06, 0x00, 0x00,
}
