// Code generated by protoc-gen-go.
// source: ipv4_rib_edm_proto.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_rip_non_as_information is a generated protocol buffer package.

It is generated from these files:
	ipv4_rib_edm_proto.proto

It has these top-level messages:
	Ipv4RibEdmProto_KEYS
	Ipv4RibEdmProto
*/
package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_rip_non_as_information

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Information of a rib protocol
type Ipv4RibEdmProto_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
}

func (m *Ipv4RibEdmProto_KEYS) Reset()                    { *m = Ipv4RibEdmProto_KEYS{} }
func (m *Ipv4RibEdmProto_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto_KEYS) ProtoMessage()               {}
func (*Ipv4RibEdmProto_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv4RibEdmProto_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

type Ipv4RibEdmProto struct {
	// Name
	ProtocolNames string `protobuf:"bytes,50,opt,name=protocol_names,json=protocolNames" json:"protocol_names,omitempty"`
	// Instance
	Instance string `protobuf:"bytes,51,opt,name=instance" json:"instance,omitempty"`
	// Proto version
	Version uint32 `protobuf:"varint,52,opt,name=version" json:"version,omitempty"`
	// Number of redist clients
	RedistributionClientCount uint32 `protobuf:"varint,53,opt,name=redistribution_client_count,json=redistributionClientCount" json:"redistribution_client_count,omitempty"`
	// Number of proto clients
	ProtocolClientsCount uint32 `protobuf:"varint,54,opt,name=protocol_clients_count,json=protocolClientsCount" json:"protocol_clients_count,omitempty"`
	// Number of routes (including active, backup and deleted), where, number of backup routes = RoutesCounts - ActiveRoutesCount - DeletedRoutesCount
	RoutesCounts uint32 `protobuf:"varint,55,opt,name=routes_counts,json=routesCounts" json:"routes_counts,omitempty"`
	// Number of active routes (not deleted)
	ActiveRoutesCount uint32 `protobuf:"varint,56,opt,name=active_routes_count,json=activeRoutesCount" json:"active_routes_count,omitempty"`
	// Number of deleted routes
	DeletedRoutesCount uint32 `protobuf:"varint,57,opt,name=deleted_routes_count,json=deletedRoutesCount" json:"deleted_routes_count,omitempty"`
	// Number of paths for all routes
	PathsCount uint32 `protobuf:"varint,58,opt,name=paths_count,json=pathsCount" json:"paths_count,omitempty"`
	// Memory for proto's routes and paths in bytes
	ProtocolRouteMemory uint32 `protobuf:"varint,59,opt,name=protocol_route_memory,json=protocolRouteMemory" json:"protocol_route_memory,omitempty"`
	// Number of backup routes
	BackupRoutesCount uint32 `protobuf:"varint,60,opt,name=backup_routes_count,json=backupRoutesCount" json:"backup_routes_count,omitempty"`
}

func (m *Ipv4RibEdmProto) Reset()                    { *m = Ipv4RibEdmProto{} }
func (m *Ipv4RibEdmProto) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto) ProtoMessage()               {}
func (*Ipv4RibEdmProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv4RibEdmProto) GetProtocolNames() string {
	if m != nil {
		return m.ProtocolNames
	}
	return ""
}

func (m *Ipv4RibEdmProto) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv4RibEdmProto) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetRedistributionClientCount() uint32 {
	if m != nil {
		return m.RedistributionClientCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetProtocolClientsCount() uint32 {
	if m != nil {
		return m.ProtocolClientsCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetRoutesCounts() uint32 {
	if m != nil {
		return m.RoutesCounts
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetActiveRoutesCount() uint32 {
	if m != nil {
		return m.ActiveRoutesCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetDeletedRoutesCount() uint32 {
	if m != nil {
		return m.DeletedRoutesCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetProtocolRouteMemory() uint32 {
	if m != nil {
		return m.ProtocolRouteMemory
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetBackupRoutesCount() uint32 {
	if m != nil {
		return m.BackupRoutesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.rip.non_as.information.ipv4_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv4RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.rip.non_as.information.ipv4_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv4_rib_edm_proto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x14, 0x35, 0x61, 0x20, 0x15, 0xb8, 0x85, 0x3a, 0x70, 0xa0, 0x2a, 0x42, 0xca,
	0xc9, 0x42, 0x6d, 0xf8, 0x8f, 0xb8, 0x54, 0x9c, 0x10, 0x1c, 0x02, 0x17, 0x4e, 0x96, 0xd7, 0xeb,
	0x15, 0x16, 0x59, 0x7b, 0xe5, 0xf1, 0xae, 0xe8, 0x63, 0xc0, 0xab, 0xf2, 0x02, 0xc8, 0xe3, 0xdd,
	0x90, 0x50, 0xf5, 0xe2, 0xc8, 0xf3, 0xfd, 0x7e, 0xce, 0x37, 0x09, 0x70, 0xdb, 0xf6, 0x4b, 0x19,
	0x6c, 0x29, 0x4d, 0xd5, 0xc8, 0x36, 0xf8, 0xe8, 0x05, 0x9d, 0xec, 0x57, 0xa1, 0x2d, 0x6a, 0x2f,
	0xad, 0x47, 0xf9, 0x33, 0x48, 0xdb, 0x12, 0x45, 0xb8, 0x6f, 0x4d, 0x10, 0xe9, 0x86, 0xb1, 0x2a,
	0x2f, 0x45, 0x1f, 0x6a, 0x4c, 0x87, 0x50, 0x35, 0x0a, 0x55, 0x0b, 0x4c, 0x9f, 0xa8, 0x6a, 0x31,
	0x38, 0xc1, 0x77, 0xd1, 0xc8, 0xa8, 0xca, 0xb5, 0x91, 0x4e, 0x35, 0x06, 0xaf, 0x0b, 0xf2, 0x37,
	0x6b, 0xbf, 0x16, 0xc1, 0xb6, 0xc2, 0x79, 0x27, 0x15, 0x0a, 0xeb, 0x6a, 0x1f, 0x1a, 0x15, 0xad,
	0x77, 0xa7, 0xbf, 0x0b, 0x38, 0xbe, 0x5a, 0x58, 0x7e, 0xfc, 0xf0, 0xed, 0x0b, 0x9b, 0xc3, 0xb4,
	0x0f, 0x35, 0xbd, 0xc3, 0x8b, 0x93, 0x62, 0x71, 0x6b, 0x35, 0xe9, 0x43, 0xfd, 0x59, 0x35, 0x86,
	0x1d, 0xc3, 0x44, 0x0d, 0xc9, 0x0d, 0x4a, 0xf6, 0x55, 0x0e, 0xe6, 0x30, 0xc5, 0x31, 0xd9, 0xcb,
	0x0e, 0x0e, 0xd1, 0x02, 0xee, 0xfe, 0x5f, 0x8f, 0xdf, 0x24, 0xe4, 0x80, 0xe6, 0x5f, 0xd3, 0x38,
	0x91, 0xa7, 0x7f, 0xf6, 0x80, 0x5d, 0x2d, 0xc5, 0x9e, 0xc2, 0xc1, 0xb8, 0x4e, 0xde, 0x9a, 0x9f,
	0x91, 0x3e, 0x1b, 0xa7, 0x49, 0x46, 0xf6, 0x10, 0xa6, 0xd6, 0x61, 0x54, 0x4e, 0x1b, 0x7e, 0x4e,
	0xc0, 0xe6, 0xce, 0x38, 0x4c, 0x7a, 0x13, 0xd0, 0x7a, 0xc7, 0x97, 0x27, 0xc5, 0x62, 0xb6, 0x1a,
	0xaf, 0xec, 0x3d, 0x3c, 0x0a, 0xa6, 0xb2, 0x18, 0x83, 0x2d, 0xbb, 0xf4, 0xd3, 0x48, 0xbd, 0xb6,
	0xc6, 0x45, 0xa9, 0x7d, 0xe7, 0x22, 0x7f, 0x4e, 0xf4, 0x7c, 0x17, 0xb9, 0x20, 0xe2, 0x22, 0x01,
	0x6c, 0x09, 0x0f, 0x36, 0xe5, 0xb2, 0x89, 0x83, 0xfa, 0x82, 0xd4, 0xa3, 0x31, 0xcd, 0x12, 0x66,
	0xeb, 0x09, 0xcc, 0x68, 0xf7, 0x81, 0x45, 0xfe, 0x92, 0xe0, 0x3b, 0x79, 0x48, 0x0c, 0x32, 0x01,
	0x87, 0x4a, 0x47, 0xdb, 0x1b, 0xb9, 0xcd, 0xf2, 0x57, 0x84, 0xde, 0xcb, 0xd1, 0xea, 0x9f, 0xc0,
	0x9e, 0xc1, 0x51, 0x65, 0xd6, 0x26, 0x9a, 0x6a, 0x57, 0x78, 0x4d, 0x02, 0x1b, 0xb2, 0x6d, 0xe3,
	0x31, 0xdc, 0x6e, 0x55, 0xfc, 0x3e, 0x82, 0x6f, 0x08, 0x04, 0x1a, 0x65, 0xe0, 0x0c, 0xee, 0x6f,
	0xb6, 0xcb, 0x7f, 0x62, 0x63, 0x1a, 0x1f, 0x2e, 0xf9, 0x5b, 0x42, 0x0f, 0xc7, 0x90, 0x1e, 0xfd,
	0x44, 0x51, 0xaa, 0x5d, 0x2a, 0xfd, 0xa3, 0x6b, 0x77, 0x5b, 0xbc, 0xcb, 0xb5, 0x73, 0xb4, 0x55,
	0xa2, 0xdc, 0xa7, 0x47, 0xce, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x65, 0x94, 0xaf, 0x4c, 0x41,
	0x03, 0x00, 0x00,
}
