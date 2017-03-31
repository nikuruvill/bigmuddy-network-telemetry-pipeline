// Code generated by protoc-gen-go.
// source: ospf_sh_stats_aggt.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_ospf_summary is a generated protocol buffer package.

It is generated from these files:
	ospf_sh_stats_aggt.proto

It has these top-level messages:
	OspfShStatsAggt_KEYS
	OspfShStatsAggt
*/
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_ospf_summary

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

type OspfShStatsAggt_KEYS struct {
	ProcessName string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	VrfName     string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
}

func (m *OspfShStatsAggt_KEYS) Reset()                    { *m = OspfShStatsAggt_KEYS{} }
func (m *OspfShStatsAggt_KEYS) String() string            { return proto.CompactTextString(m) }
func (*OspfShStatsAggt_KEYS) ProtoMessage()               {}
func (*OspfShStatsAggt_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OspfShStatsAggt_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShStatsAggt_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type OspfShStatsAggt struct {
	SaNumNbrs            uint32 `protobuf:"varint,50,opt,name=sa_num_nbrs,json=saNumNbrs" json:"sa_num_nbrs,omitempty"`
	SaNumNbrsUp          uint32 `protobuf:"varint,51,opt,name=sa_num_nbrs_up,json=saNumNbrsUp" json:"sa_num_nbrs_up,omitempty"`
	SaNumIntf            uint32 `protobuf:"varint,52,opt,name=sa_num_intf,json=saNumIntf" json:"sa_num_intf,omitempty"`
	SaNumIntfUp          uint32 `protobuf:"varint,53,opt,name=sa_num_intf_up,json=saNumIntfUp" json:"sa_num_intf_up,omitempty"`
	SaNumVintfUp         uint32 `protobuf:"varint,54,opt,name=sa_num_vintf_up,json=saNumVintfUp" json:"sa_num_vintf_up,omitempty"`
	SaNumSlintfUp        uint32 `protobuf:"varint,55,opt,name=sa_num_slintf_up,json=saNumSlintfUp" json:"sa_num_slintf_up,omitempty"`
	SaNumAreas           uint32 `protobuf:"varint,56,opt,name=sa_num_areas,json=saNumAreas" json:"sa_num_areas,omitempty"`
	SaLsaCntTypeRtr      uint32 `protobuf:"varint,57,opt,name=sa_lsa_cnt_type_rtr,json=saLsaCntTypeRtr" json:"sa_lsa_cnt_type_rtr,omitempty"`
	SaLsaCntTypeNet      uint32 `protobuf:"varint,58,opt,name=sa_lsa_cnt_type_net,json=saLsaCntTypeNet" json:"sa_lsa_cnt_type_net,omitempty"`
	SaLsaCntTypeSumNet   uint32 `protobuf:"varint,59,opt,name=sa_lsa_cnt_type_sum_net,json=saLsaCntTypeSumNet" json:"sa_lsa_cnt_type_sum_net,omitempty"`
	SaLsaCntTypeSumAsb   uint32 `protobuf:"varint,60,opt,name=sa_lsa_cnt_type_sum_asb,json=saLsaCntTypeSumAsb" json:"sa_lsa_cnt_type_sum_asb,omitempty"`
	SaLsaCntTypeAse      uint32 `protobuf:"varint,61,opt,name=sa_lsa_cnt_type_ase,json=saLsaCntTypeAse" json:"sa_lsa_cnt_type_ase,omitempty"`
	SaLsaCntTypeMospf    uint32 `protobuf:"varint,62,opt,name=sa_lsa_cnt_type_mospf,json=saLsaCntTypeMospf" json:"sa_lsa_cnt_type_mospf,omitempty"`
	SaLsaCntType_7Ase    uint32 `protobuf:"varint,63,opt,name=sa_lsa_cnt_type_7ase,json=saLsaCntType7ase" json:"sa_lsa_cnt_type_7ase,omitempty"`
	SaLsaCntType_8Ignore uint32 `protobuf:"varint,64,opt,name=sa_lsa_cnt_type_8_ignore,json=saLsaCntType8Ignore" json:"sa_lsa_cnt_type_8_ignore,omitempty"`
	SaLsaCntTypeOpqLink  uint32 `protobuf:"varint,65,opt,name=sa_lsa_cnt_type_opq_link,json=saLsaCntTypeOpqLink" json:"sa_lsa_cnt_type_opq_link,omitempty"`
	SaLsaCntTypeOpqArea  uint32 `protobuf:"varint,66,opt,name=sa_lsa_cnt_type_opq_area,json=saLsaCntTypeOpqArea" json:"sa_lsa_cnt_type_opq_area,omitempty"`
	SaLsaCntTypeOpqAs    uint32 `protobuf:"varint,67,opt,name=sa_lsa_cnt_type_opq_as,json=saLsaCntTypeOpqAs" json:"sa_lsa_cnt_type_opq_as,omitempty"`
}

func (m *OspfShStatsAggt) Reset()                    { *m = OspfShStatsAggt{} }
func (m *OspfShStatsAggt) String() string            { return proto.CompactTextString(m) }
func (*OspfShStatsAggt) ProtoMessage()               {}
func (*OspfShStatsAggt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OspfShStatsAggt) GetSaNumNbrs() uint32 {
	if m != nil {
		return m.SaNumNbrs
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaNumNbrsUp() uint32 {
	if m != nil {
		return m.SaNumNbrsUp
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaNumIntf() uint32 {
	if m != nil {
		return m.SaNumIntf
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaNumIntfUp() uint32 {
	if m != nil {
		return m.SaNumIntfUp
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaNumVintfUp() uint32 {
	if m != nil {
		return m.SaNumVintfUp
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaNumSlintfUp() uint32 {
	if m != nil {
		return m.SaNumSlintfUp
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaNumAreas() uint32 {
	if m != nil {
		return m.SaNumAreas
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeRtr() uint32 {
	if m != nil {
		return m.SaLsaCntTypeRtr
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeNet() uint32 {
	if m != nil {
		return m.SaLsaCntTypeNet
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeSumNet() uint32 {
	if m != nil {
		return m.SaLsaCntTypeSumNet
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeSumAsb() uint32 {
	if m != nil {
		return m.SaLsaCntTypeSumAsb
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeAse() uint32 {
	if m != nil {
		return m.SaLsaCntTypeAse
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeMospf() uint32 {
	if m != nil {
		return m.SaLsaCntTypeMospf
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntType_7Ase() uint32 {
	if m != nil {
		return m.SaLsaCntType_7Ase
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntType_8Ignore() uint32 {
	if m != nil {
		return m.SaLsaCntType_8Ignore
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeOpqLink() uint32 {
	if m != nil {
		return m.SaLsaCntTypeOpqLink
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeOpqArea() uint32 {
	if m != nil {
		return m.SaLsaCntTypeOpqArea
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeOpqAs() uint32 {
	if m != nil {
		return m.SaLsaCntTypeOpqAs
	}
	return 0
}

func init() {
	proto.RegisterType((*OspfShStatsAggt_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.ospf_summary.ospf_sh_stats_aggt_KEYS")
	proto.RegisterType((*OspfShStatsAggt)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.ospf_summary.ospf_sh_stats_aggt")
}

func init() { proto.RegisterFile("ospf_sh_stats_aggt.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x86, 0x55, 0x2e, 0x06, 0x73, 0x3b, 0x36, 0x3c, 0x60, 0xe6, 0x06, 0x95, 0x21, 0xc4, 0x2e,
	0x50, 0x04, 0x74, 0x63, 0xe5, 0x9b, 0x32, 0x21, 0x31, 0x31, 0x8a, 0xd4, 0x32, 0x10, 0x57, 0x47,
	0x4e, 0xe5, 0x94, 0x68, 0x8d, 0xe3, 0xf9, 0x38, 0x11, 0xfd, 0x73, 0xfc, 0x36, 0xe4, 0xd3, 0x14,
	0x59, 0x49, 0xe0, 0xa6, 0xb2, 0xce, 0xfb, 0x3c, 0xaf, 0x2d, 0xbb, 0x61, 0x22, 0x47, 0x93, 0x00,
	0xfe, 0x04, 0x74, 0xd2, 0x21, 0xc8, 0xf9, 0xdc, 0x45, 0xc6, 0xe6, 0x2e, 0xe7, 0x1f, 0x67, 0x29,
	0xce, 0x72, 0x48, 0x73, 0x84, 0x5f, 0x16, 0x52, 0x53, 0x1e, 0x02, 0xb1, 0xb9, 0x51, 0x36, 0xf2,
	0x2b, 0xcf, 0xcd, 0x14, 0xa2, 0xc2, 0xf5, 0x2a, 0x2a, 0x6d, 0x42, 0x3f, 0xd1, 0xaa, 0xb5, 0xc8,
	0x32, 0x69, 0x97, 0xfb, 0xdf, 0xd9, 0x5e, 0x73, 0x17, 0xf8, 0xf4, 0xe1, 0xc7, 0x94, 0xdf, 0x63,
	0xbd, 0xca, 0x05, 0x2d, 0x33, 0x25, 0x3a, 0xfd, 0xce, 0xc1, 0xe6, 0xa4, 0x5b, 0xcd, 0xc6, 0x32,
	0x53, 0xfc, 0x0e, 0xbb, 0x56, 0xda, 0x64, 0x15, 0x5f, 0xa1, 0xf8, 0x6a, 0x69, 0x13, 0x1f, 0xed,
	0xff, 0xde, 0x60, 0xbc, 0xd9, 0xcc, 0xef, 0xb2, 0x2e, 0x4a, 0xd0, 0x45, 0x06, 0x3a, 0xb6, 0x28,
	0x9e, 0xf6, 0x3b, 0x07, 0x5b, 0x93, 0x4d, 0x94, 0xe3, 0x22, 0x1b, 0xc7, 0x16, 0xf9, 0x7d, 0x76,
	0x3d, 0xc8, 0xa1, 0x30, 0x62, 0x40, 0x48, 0xf7, 0x2f, 0x72, 0x6e, 0x82, 0x92, 0x54, 0xbb, 0x44,
	0x1c, 0x06, 0x25, 0xa7, 0xda, 0x25, 0x41, 0x89, 0xcf, 0x7d, 0xc9, 0x51, 0x50, 0xe2, 0x91, 0x73,
	0xc3, 0x1f, 0xb0, 0xed, 0x0a, 0x2a, 0xd7, 0xd4, 0x33, 0xa2, 0x7a, 0x44, 0x7d, 0x4b, 0x57, 0xd8,
	0x43, 0xb6, 0x53, 0x61, 0xb8, 0x58, 0x73, 0xc7, 0xc4, 0x6d, 0x11, 0x37, 0x5d, 0x54, 0x60, 0x9f,
	0xf5, 0x2a, 0x50, 0x5a, 0x25, 0x51, 0x0c, 0x09, 0x62, 0x04, 0x8d, 0xfc, 0x84, 0x3f, 0x62, 0xbb,
	0x28, 0x61, 0x81, 0x12, 0x66, 0xda, 0x81, 0x5b, 0x1a, 0x05, 0xd6, 0x59, 0xf1, 0x9c, 0xc0, 0x6d,
	0x94, 0x67, 0x28, 0x4f, 0xb4, 0xfb, 0xba, 0x34, 0x6a, 0xe2, 0x6c, 0x1b, 0xad, 0x95, 0x13, 0x2f,
	0x9a, 0xf4, 0x58, 0x39, 0x3e, 0x60, 0x7b, 0x75, 0x1a, 0xfd, 0x25, 0x2a, 0x27, 0x5e, 0x92, 0xc1,
	0x43, 0x63, 0x5a, 0x64, 0xff, 0x91, 0x24, 0xc6, 0xe2, 0x55, 0xab, 0x34, 0xc2, 0xb8, 0xed, 0x5c,
	0x12, 0x95, 0x78, 0xdd, 0x3c, 0xd7, 0x08, 0x15, 0x7f, 0xcc, 0x6e, 0xd5, 0xe9, 0xcc, 0xff, 0x2d,
	0xc4, 0x1b, 0xe2, 0x6f, 0x84, 0xfc, 0x67, 0x1f, 0xf0, 0x88, 0xdd, 0xac, 0x1b, 0xc7, 0x7e, 0x83,
	0xb7, 0x24, 0xec, 0x84, 0x82, 0x9f, 0xf3, 0x23, 0x26, 0xea, 0xfc, 0x10, 0xd2, 0xb9, 0xce, 0xad,
	0x12, 0xef, 0xc8, 0xd9, 0x0d, 0x9d, 0xe1, 0x29, 0x45, 0x6d, 0x5a, 0x6e, 0x2e, 0x61, 0x91, 0xea,
	0x0b, 0x31, 0x6a, 0x6a, 0x5f, 0xcc, 0xe5, 0x59, 0xaa, 0x2f, 0xfe, 0xa5, 0xf9, 0x27, 0x17, 0xef,
	0x5b, 0x35, 0xff, 0xf6, 0xfc, 0x09, 0xbb, 0xdd, 0xaa, 0xa1, 0x38, 0x69, 0xde, 0x83, 0x97, 0x30,
	0xde, 0xa0, 0x4f, 0x7d, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x33, 0xaa, 0xa2, 0x06, 0x04,
	0x00, 0x00,
}
