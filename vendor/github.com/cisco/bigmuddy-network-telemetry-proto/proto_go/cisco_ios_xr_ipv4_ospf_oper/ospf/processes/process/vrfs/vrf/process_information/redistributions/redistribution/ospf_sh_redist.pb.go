// Code generated by protoc-gen-go.
// source: ospf_sh_redist.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_process_information_redistributions_redistribution is a generated protocol buffer package.

It is generated from these files:
	ospf_sh_redist.proto

It has these top-level messages:
	OspfShRedist_KEYS
	OspfShRedist
	StringType
	OspfShRedistProto
*/
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_process_information_redistributions_redistribution

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

// OSPF Redistribution Information
type OspfShRedist_KEYS struct {
	ProcessName   string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	VrfName       string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	ProtocolName  string `protobuf:"bytes,3,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	ProcessName_1 string `protobuf:"bytes,4,opt,name=process_name_1,json=processName1" json:"process_name_1,omitempty"`
}

func (m *OspfShRedist_KEYS) Reset()                    { *m = OspfShRedist_KEYS{} }
func (m *OspfShRedist_KEYS) String() string            { return proto.CompactTextString(m) }
func (*OspfShRedist_KEYS) ProtoMessage()               {}
func (*OspfShRedist_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OspfShRedist_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShRedist_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShRedist_KEYS) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *OspfShRedist_KEYS) GetProcessName_1() string {
	if m != nil {
		return m.ProcessName_1
	}
	return ""
}

type OspfShRedist struct {
	// Protocol Information
	RedistributionProtocol *OspfShRedistProto `protobuf:"bytes,50,opt,name=redistribution_protocol,json=redistributionProtocol" json:"redistribution_protocol,omitempty"`
	// If true, Metric configured
	MetricFlag bool `protobuf:"varint,51,opt,name=metric_flag,json=metricFlag" json:"metric_flag,omitempty"`
	// Route metric
	Metric uint32 `protobuf:"varint,52,opt,name=metric" json:"metric,omitempty"`
	// Whether to use CIDR
	Classless bool `protobuf:"varint,53,opt,name=classless" json:"classless,omitempty"`
	// To NSSA areas only
	NssaOnly bool `protobuf:"varint,54,opt,name=nssa_only,json=nssaOnly" json:"nssa_only,omitempty"`
}

func (m *OspfShRedist) Reset()                    { *m = OspfShRedist{} }
func (m *OspfShRedist) String() string            { return proto.CompactTextString(m) }
func (*OspfShRedist) ProtoMessage()               {}
func (*OspfShRedist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OspfShRedist) GetRedistributionProtocol() *OspfShRedistProto {
	if m != nil {
		return m.RedistributionProtocol
	}
	return nil
}

func (m *OspfShRedist) GetMetricFlag() bool {
	if m != nil {
		return m.MetricFlag
	}
	return false
}

func (m *OspfShRedist) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *OspfShRedist) GetClassless() bool {
	if m != nil {
		return m.Classless
	}
	return false
}

func (m *OspfShRedist) GetNssaOnly() bool {
	if m != nil {
		return m.NssaOnly
	}
	return false
}

type StringType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StringType) Reset()                    { *m = StringType{} }
func (m *StringType) String() string            { return proto.CompactTextString(m) }
func (*StringType) ProtoMessage()               {}
func (*StringType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StringType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Redistributed protocol instance
type OspfShRedistProto struct {
	ProtocolType string `protobuf:"bytes,1,opt,name=protocol_type,json=protocolType" json:"protocol_type,omitempty"`
	// ISIS instance identifier
	IsisInstanceId *StringType `protobuf:"bytes,2,opt,name=isis_instance_id,json=isisInstanceId" json:"isis_instance_id,omitempty"`
	// OSPF process identifier
	OspfProcessId *StringType `protobuf:"bytes,3,opt,name=ospf_process_id,json=ospfProcessId" json:"ospf_process_id,omitempty"`
	// BGP Autonomous System Number
	BgpAsNumber *StringType `protobuf:"bytes,4,opt,name=bgp_as_number,json=bgpAsNumber" json:"bgp_as_number,omitempty"`
	// Autonomous System Number
	EigrpAsNumber *StringType `protobuf:"bytes,5,opt,name=eigrp_as_number,json=eigrpAsNumber" json:"eigrp_as_number,omitempty"`
	// Connected instance name
	ConnectedInstance *StringType `protobuf:"bytes,6,opt,name=connected_instance,json=connectedInstance" json:"connected_instance,omitempty"`
}

func (m *OspfShRedistProto) Reset()                    { *m = OspfShRedistProto{} }
func (m *OspfShRedistProto) String() string            { return proto.CompactTextString(m) }
func (*OspfShRedistProto) ProtoMessage()               {}
func (*OspfShRedistProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OspfShRedistProto) GetProtocolType() string {
	if m != nil {
		return m.ProtocolType
	}
	return ""
}

func (m *OspfShRedistProto) GetIsisInstanceId() *StringType {
	if m != nil {
		return m.IsisInstanceId
	}
	return nil
}

func (m *OspfShRedistProto) GetOspfProcessId() *StringType {
	if m != nil {
		return m.OspfProcessId
	}
	return nil
}

func (m *OspfShRedistProto) GetBgpAsNumber() *StringType {
	if m != nil {
		return m.BgpAsNumber
	}
	return nil
}

func (m *OspfShRedistProto) GetEigrpAsNumber() *StringType {
	if m != nil {
		return m.EigrpAsNumber
	}
	return nil
}

func (m *OspfShRedistProto) GetConnectedInstance() *StringType {
	if m != nil {
		return m.ConnectedInstance
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShRedist_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.process_information.redistributions.redistribution.ospf_sh_redist_KEYS")
	proto.RegisterType((*OspfShRedist)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.process_information.redistributions.redistribution.ospf_sh_redist")
	proto.RegisterType((*StringType)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.process_information.redistributions.redistribution.string_type")
	proto.RegisterType((*OspfShRedistProto)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.process_information.redistributions.redistribution.ospf_sh_redist_proto")
}

func init() { proto.RegisterFile("ospf_sh_redist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0xe5, 0xc1, 0x4a, 0xfb, 0x64, 0x2d, 0x60, 0xa6, 0x11, 0x04, 0x12, 0xa5, 0xe3, 0xd0,
	0x53, 0xa4, 0x75, 0x83, 0x3b, 0x07, 0x90, 0x2a, 0xa4, 0x31, 0x15, 0x2e, 0x9c, 0x2c, 0x27, 0x71,
	0x82, 0xa5, 0xc4, 0x8e, 0xfc, 0xa4, 0x11, 0x15, 0x57, 0x3e, 0x07, 0x1c, 0x10, 0x17, 0x3e, 0x04,
	0x37, 0xbe, 0x16, 0x42, 0x76, 0x5e, 0x68, 0xfa, 0x01, 0xd6, 0x5d, 0xa2, 0x3c, 0x2f, 0xfe, 0xfb,
	0xf7, 0xbc, 0x18, 0x8e, 0x35, 0x16, 0x09, 0xc3, 0x4f, 0xcc, 0x88, 0x58, 0x62, 0x19, 0x14, 0x46,
	0x97, 0x9a, 0xaa, 0x48, 0x62, 0xa4, 0x99, 0xd4, 0xc8, 0x3e, 0x1b, 0x26, 0x8b, 0xea, 0x82, 0xb9,
	0x3c, 0x5d, 0x08, 0x13, 0xd8, 0x3f, 0x9b, 0x17, 0x09, 0x44, 0x81, 0xed, 0x5f, 0x50, 0x99, 0xc4,
	0x7d, 0x5a, 0x07, 0x93, 0x2a, 0xd1, 0x26, 0xe7, 0xa5, 0xd4, 0x2a, 0xa8, 0xd5, 0x8d, 0x0c, 0xd7,
	0xd6, 0xc4, 0x1d, 0x7b, 0xf6, 0x8d, 0xc0, 0x83, 0x3e, 0x08, 0x7b, 0xfb, 0xfa, 0xe3, 0x7b, 0xfa,
	0x0c, 0x8e, 0x5a, 0x35, 0xc5, 0x73, 0xe1, 0x93, 0x29, 0x99, 0x8f, 0x56, 0x5e, 0xe3, 0xbb, 0xe4,
	0xb9, 0xa0, 0x8f, 0x60, 0x58, 0x99, 0xa4, 0x0e, 0x1f, 0xb8, 0xf0, 0x9d, 0xca, 0x24, 0x2e, 0x74,
	0x0a, 0x63, 0x57, 0x4e, 0xa4, 0xb3, 0x3a, 0x7e, 0xcb, 0xc5, 0x8f, 0x5a, 0xa7, 0x4b, 0x7a, 0x0e,
	0x93, 0xed, 0x2b, 0xd8, 0x99, 0x7f, 0xbb, 0xcb, 0x6a, 0x2f, 0x39, 0x9b, 0xfd, 0x39, 0x80, 0x49,
	0x1f, 0x90, 0xfe, 0x26, 0xf0, 0xb0, 0x5f, 0x06, 0x6b, 0x85, 0xfd, 0xc5, 0x94, 0xcc, 0xbd, 0xc5,
	0x57, 0x12, 0x5c, 0x6f, 0x1f, 0x83, 0x9d, 0x1e, 0x3a, 0x9e, 0xd5, 0x49, 0x3f, 0xe9, 0xaa, 0x81,
	0xa4, 0x4f, 0xc1, 0xcb, 0x45, 0x69, 0x64, 0xc4, 0x92, 0x8c, 0xa7, 0xfe, 0xf9, 0x94, 0xcc, 0x87,
	0x2b, 0xa8, 0x5d, 0x6f, 0x32, 0x9e, 0xd2, 0x13, 0x18, 0xd4, 0x96, 0x7f, 0x31, 0x25, 0xf3, 0xf1,
	0xaa, 0xb1, 0xe8, 0x13, 0x18, 0x45, 0x19, 0x47, 0xcc, 0x04, 0xa2, 0xff, 0xc2, 0x1d, 0xfb, 0xef,
	0xa0, 0x8f, 0x61, 0xa4, 0x10, 0x39, 0xd3, 0x2a, 0xdb, 0xf8, 0x2f, 0x5d, 0x74, 0x68, 0x1d, 0xef,
	0x54, 0xb6, 0x99, 0x9d, 0x82, 0x67, 0x49, 0x54, 0xca, 0xca, 0x4d, 0x21, 0xe8, 0x31, 0x1c, 0x56,
	0x3c, 0x5b, 0xb7, 0x83, 0xad, 0x8d, 0xd9, 0xdf, 0xc1, 0xee, 0x5a, 0xd6, 0x95, 0xf4, 0x06, 0x6a,
	0xcf, 0x37, 0xc7, 0xba, 0x81, 0x7e, 0xb0, 0x9a, 0x3f, 0x09, 0xdc, 0x93, 0x28, 0x6d, 0xdf, 0xb0,
	0xe4, 0x2a, 0x12, 0x4c, 0xc6, 0x6e, 0x33, 0xbc, 0xc5, 0x97, 0xeb, 0x9e, 0xc7, 0x56, 0xad, 0xab,
	0x89, 0x85, 0x5a, 0x36, 0x4c, 0xcb, 0x98, 0xfe, 0x20, 0x70, 0xd7, 0x5d, 0xde, 0xe9, 0xc6, 0x6e,
	0x41, 0xf7, 0x8c, 0x39, 0xb6, 0xfa, 0x57, 0xb5, 0xc6, 0x32, 0xa6, 0xdf, 0x09, 0x8c, 0xc3, 0xb4,
	0x60, 0x1c, 0x99, 0x5a, 0xe7, 0xa1, 0x30, 0xee, 0x79, 0xec, 0x99, 0xd1, 0x0b, 0xd3, 0xe2, 0x15,
	0x5e, 0x3a, 0x1e, 0xd7, 0x47, 0x21, 0x53, 0xb3, 0xcd, 0x78, 0x78, 0x03, 0xfa, 0xe8, 0x98, 0x3a,
	0xca, 0x5f, 0x04, 0x68, 0xa4, 0x95, 0x12, 0x51, 0x29, 0xe2, 0x6e, 0x35, 0xfd, 0xc1, 0xfe, 0x41,
	0xef, 0x77, 0x58, 0xed, 0x72, 0x86, 0x03, 0xf7, 0xa0, 0xce, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x60, 0x4f, 0x74, 0x24, 0x1d, 0x06, 0x00, 0x00,
}
