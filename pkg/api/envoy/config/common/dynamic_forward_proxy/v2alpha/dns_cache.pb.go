// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: envoy/config/common/dynamic_forward_proxy/v2alpha/dns_cache.proto

package v2alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v2 "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/api/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configuration for the dynamic forward proxy DNS cache. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
// [#next-free-field: 7]
type DnsCacheConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the cache. Multiple named caches allow independent dynamic forward proxy
	// configurations to operate within a single Envoy process using different configurations. All
	// configurations with the same name *must* otherwise have the same settings when referenced
	// from different configuration components. Configuration will fail to load if this is not
	// the case.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The DNS lookup family to use during resolution.
	//
	// [#comment:TODO(mattklein123): Figure out how to support IPv4/IPv6 "happy eyeballs" mode. The
	// way this might work is a new lookup family which returns both IPv4 and IPv6 addresses, and
	// then configures a host to have a primary and fall back address. With this, we could very
	// likely build a "happy eyeballs" connection pool which would race the primary / fall back
	// address and return the one that wins. This same method could potentially also be used for
	// QUIC to TCP fall back.]
	DnsLookupFamily v2.Cluster_DnsLookupFamily `protobuf:"varint,2,opt,name=dns_lookup_family,json=dnsLookupFamily,proto3,enum=envoy.api.v2.Cluster_DnsLookupFamily" json:"dns_lookup_family,omitempty"`
	// The DNS refresh rate for currently cached DNS hosts. If not specified defaults to 60s.
	//
	// .. note:
	//
	//  The returned DNS TTL is not currently used to alter the refresh rate. This feature will be
	//  added in a future change.
	//
	// .. note:
	//
	// The refresh rate is rounded to the closest millisecond, and must be at least 1ms.
	DnsRefreshRate *duration.Duration `protobuf:"bytes,3,opt,name=dns_refresh_rate,json=dnsRefreshRate,proto3" json:"dns_refresh_rate,omitempty"`
	// The TTL for hosts that are unused. Hosts that have not been used in the configured time
	// interval will be purged. If not specified defaults to 5m.
	//
	// .. note:
	//
	//   The TTL is only checked at the time of DNS refresh, as specified by *dns_refresh_rate*. This
	//   means that if the configured TTL is shorter than the refresh rate the host may not be removed
	//   immediately.
	//
	//  .. note:
	//
	//   The TTL has no relation to DNS TTL and is only used to control Envoy's resource usage.
	HostTtl *duration.Duration `protobuf:"bytes,4,opt,name=host_ttl,json=hostTtl,proto3" json:"host_ttl,omitempty"`
	// The maximum number of hosts that the cache will hold. If not specified defaults to 1024.
	//
	// .. note:
	//
	//   The implementation is approximate and enforced independently on each worker thread, thus
	//   it is possible for the maximum hosts in the cache to go slightly above the configured
	//   value depending on timing. This is similar to how other circuit breakers work.
	MaxHosts *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=max_hosts,json=maxHosts,proto3" json:"max_hosts,omitempty"`
	// If the DNS failure refresh rate is specified,
	// this is used as the cache's DNS refresh rate when DNS requests are failing. If this setting is
	// not specified, the failure refresh rate defaults to the dns_refresh_rate.
	DnsFailureRefreshRate *v2.Cluster_RefreshRate `protobuf:"bytes,6,opt,name=dns_failure_refresh_rate,json=dnsFailureRefreshRate,proto3" json:"dns_failure_refresh_rate,omitempty"`
}

func (x *DnsCacheConfig) Reset() {
	*x = DnsCacheConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsCacheConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsCacheConfig) ProtoMessage() {}

func (x *DnsCacheConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsCacheConfig.ProtoReflect.Descriptor instead.
func (*DnsCacheConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_rawDescGZIP(), []int{0}
}

func (x *DnsCacheConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DnsCacheConfig) GetDnsLookupFamily() v2.Cluster_DnsLookupFamily {
	if x != nil {
		return x.DnsLookupFamily
	}
	return v2.Cluster_DnsLookupFamily(0)
}

func (x *DnsCacheConfig) GetDnsRefreshRate() *duration.Duration {
	if x != nil {
		return x.DnsRefreshRate
	}
	return nil
}

func (x *DnsCacheConfig) GetHostTtl() *duration.Duration {
	if x != nil {
		return x.HostTtl
	}
	return nil
}

func (x *DnsCacheConfig) GetMaxHosts() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxHosts
	}
	return nil
}

func (x *DnsCacheConfig) GetDnsFailureRefreshRate() *v2.Cluster_RefreshRate {
	if x != nil {
		return x.DnsFailureRefreshRate
	}
	return nil
}

var File_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto protoreflect.FileDescriptor

var file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_rawDesc = []byte{
	0x0a, 0x41, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x03, 0x0a, 0x0e, 0x44,
	0x6e, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x5b, 0x0a, 0x11, 0x64, 0x6e,
	0x73, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x6e, 0x73,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x64, 0x6e, 0x73, 0x4c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x51, 0x0a, 0x10, 0x64, 0x6e, 0x73, 0x5f, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0xfa, 0x42,
	0x09, 0xaa, 0x01, 0x06, 0x32, 0x04, 0x10, 0xc0, 0x84, 0x3d, 0x52, 0x0e, 0x64, 0x6e, 0x73, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a,
	0x00, 0x52, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x54, 0x74, 0x6c, 0x12, 0x42, 0x0a, 0x09, 0x6d, 0x61,
	0x78, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x5a,
	0x0a, 0x18, 0x64, 0x6e, 0x73, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x15, 0x64, 0x6e, 0x73, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x42, 0xec, 0x01, 0x0a, 0x3f, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0d,
	0x44, 0x6e, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x32, 0x12,
	0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76,
	0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_rawDescOnce sync.Once
	file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_rawDescData = file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_rawDesc
)

func file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_rawDescGZIP() []byte {
	file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_rawDescOnce.Do(func() {
		file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_rawDescData)
	})
	return file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_rawDescData
}

var file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_goTypes = []interface{}{
	(*DnsCacheConfig)(nil),          // 0: envoy.config.common.dynamic_forward_proxy.v2alpha.DnsCacheConfig
	(v2.Cluster_DnsLookupFamily)(0), // 1: envoy.api.v2.Cluster.DnsLookupFamily
	(*duration.Duration)(nil),       // 2: google.protobuf.Duration
	(*wrappers.UInt32Value)(nil),    // 3: google.protobuf.UInt32Value
	(*v2.Cluster_RefreshRate)(nil),  // 4: envoy.api.v2.Cluster.RefreshRate
}
var file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_depIdxs = []int32{
	1, // 0: envoy.config.common.dynamic_forward_proxy.v2alpha.DnsCacheConfig.dns_lookup_family:type_name -> envoy.api.v2.Cluster.DnsLookupFamily
	2, // 1: envoy.config.common.dynamic_forward_proxy.v2alpha.DnsCacheConfig.dns_refresh_rate:type_name -> google.protobuf.Duration
	2, // 2: envoy.config.common.dynamic_forward_proxy.v2alpha.DnsCacheConfig.host_ttl:type_name -> google.protobuf.Duration
	3, // 3: envoy.config.common.dynamic_forward_proxy.v2alpha.DnsCacheConfig.max_hosts:type_name -> google.protobuf.UInt32Value
	4, // 4: envoy.config.common.dynamic_forward_proxy.v2alpha.DnsCacheConfig.dns_failure_refresh_rate:type_name -> envoy.api.v2.Cluster.RefreshRate
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_init() }
func file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_init() {
	if File_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsCacheConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_goTypes,
		DependencyIndexes: file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_depIdxs,
		MessageInfos:      file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_msgTypes,
	}.Build()
	File_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto = out.File
	file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_rawDesc = nil
	file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_goTypes = nil
	file_envoy_config_common_dynamic_forward_proxy_v2alpha_dns_cache_proto_depIdxs = nil
}
