// Code generated by protoc-gen-go.
// source: snmp_bulkstats_b.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_snmp_agent_oper_snmp_information_bulk_stats_transfers_bulk_stats_transfer is a generated protocol buffer package.

It is generated from these files:
	snmp_bulkstats_b.proto

It has these top-level messages:
	SnmpBulkstatsB_KEYS
	SnmpBulkstatsB
*/
package cisco_ios_xr_snmp_agent_oper_snmp_information_bulk_stats_transfers_bulk_stats_transfer

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

// SNMP Bulkstats transfer Information
type SnmpBulkstatsB_KEYS struct {
	TransferName string `protobuf:"bytes,1,opt,name=transfer_name,json=transferName" json:"transfer_name,omitempty"`
}

func (m *SnmpBulkstatsB_KEYS) Reset()                    { *m = SnmpBulkstatsB_KEYS{} }
func (m *SnmpBulkstatsB_KEYS) String() string            { return proto.CompactTextString(m) }
func (*SnmpBulkstatsB_KEYS) ProtoMessage()               {}
func (*SnmpBulkstatsB_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SnmpBulkstatsB_KEYS) GetTransferName() string {
	if m != nil {
		return m.TransferName
	}
	return ""
}

type SnmpBulkstatsB struct {
	// Name of the bulkstats transfer session
	TransferName string `protobuf:"bytes,50,opt,name=transfer_name,json=transferName" json:"transfer_name,omitempty"`
	// Bulkstats transfer primary URL
	UrlPrimary string `protobuf:"bytes,51,opt,name=url_primary,json=urlPrimary" json:"url_primary,omitempty"`
	// Bulkstats transfer secondary URL
	UrlSecondary string `protobuf:"bytes,52,opt,name=url_secondary,json=urlSecondary" json:"url_secondary,omitempty"`
	// Bulkstats transfer retained file name
	RetainedFile string `protobuf:"bytes,53,opt,name=retained_file,json=retainedFile" json:"retained_file,omitempty"`
	// Bulkstats transfer retry time left in seconds
	TimeLeft uint32 `protobuf:"varint,54,opt,name=time_left,json=timeLeft" json:"time_left,omitempty"`
	// Bulkstats transfer retry attempt left
	RetryLeft uint32 `protobuf:"varint,55,opt,name=retry_left,json=retryLeft" json:"retry_left,omitempty"`
}

func (m *SnmpBulkstatsB) Reset()                    { *m = SnmpBulkstatsB{} }
func (m *SnmpBulkstatsB) String() string            { return proto.CompactTextString(m) }
func (*SnmpBulkstatsB) ProtoMessage()               {}
func (*SnmpBulkstatsB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SnmpBulkstatsB) GetTransferName() string {
	if m != nil {
		return m.TransferName
	}
	return ""
}

func (m *SnmpBulkstatsB) GetUrlPrimary() string {
	if m != nil {
		return m.UrlPrimary
	}
	return ""
}

func (m *SnmpBulkstatsB) GetUrlSecondary() string {
	if m != nil {
		return m.UrlSecondary
	}
	return ""
}

func (m *SnmpBulkstatsB) GetRetainedFile() string {
	if m != nil {
		return m.RetainedFile
	}
	return ""
}

func (m *SnmpBulkstatsB) GetTimeLeft() uint32 {
	if m != nil {
		return m.TimeLeft
	}
	return 0
}

func (m *SnmpBulkstatsB) GetRetryLeft() uint32 {
	if m != nil {
		return m.RetryLeft
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpBulkstatsB_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.bulk_stats_transfers.bulk_stats_transfer.snmp_bulkstats_b_KEYS")
	proto.RegisterType((*SnmpBulkstatsB)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.bulk_stats_transfers.bulk_stats_transfer.snmp_bulkstats_b")
}

func init() { proto.RegisterFile("snmp_bulkstats_b.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4d, 0x4b, 0x33, 0x31,
	0x14, 0x85, 0x99, 0xcd, 0xcb, 0x3b, 0xd1, 0x82, 0x04, 0x94, 0x01, 0x11, 0x4b, 0xdd, 0x74, 0x35,
	0x0b, 0xeb, 0xc7, 0xc6, 0xad, 0x6e, 0x14, 0x91, 0x16, 0x04, 0x57, 0x97, 0xcc, 0xf4, 0x46, 0x82,
	0xf9, 0x18, 0x6e, 0xee, 0x80, 0xfd, 0xb3, 0xfe, 0x16, 0x49, 0xc2, 0x6c, 0x5a, 0x97, 0x79, 0xce,
	0x73, 0x4e, 0xe0, 0x8a, 0xb3, 0xe8, 0xdd, 0x00, 0xdd, 0x68, 0xbf, 0x22, 0x2b, 0x8e, 0xd0, 0xb5,
	0x03, 0x05, 0x0e, 0xf2, 0xbd, 0x37, 0xb1, 0x0f, 0x60, 0x42, 0x84, 0x6f, 0x82, 0x2c, 0xa9, 0x4f,
	0xf4, 0x0c, 0x61, 0x40, 0x6a, 0xd3, 0xbb, 0x35, 0x5e, 0x07, 0x72, 0x8a, 0x4d, 0xf0, 0x6d, 0x1a,
	0x80, 0xb2, 0xc0, 0xa4, 0x7c, 0xd4, 0x48, 0xf1, 0x2f, 0xb8, 0x78, 0x10, 0xa7, 0xfb, 0x3f, 0xc2,
	0xf3, 0xe3, 0xc7, 0x46, 0x5e, 0x89, 0xd9, 0x24, 0x81, 0x57, 0x0e, 0x9b, 0x6a, 0x5e, 0x2d, 0xeb,
	0xf5, 0xf1, 0x04, 0x5f, 0x95, 0xc3, 0xc5, 0x4f, 0x25, 0x4e, 0xf6, 0xeb, 0x87, 0xcd, 0xeb, 0xc3,
	0xa6, 0xbc, 0x14, 0x47, 0x23, 0x59, 0x18, 0xc8, 0x38, 0x45, 0xbb, 0x66, 0x95, 0x15, 0x31, 0x92,
	0x7d, 0x2b, 0x24, 0xad, 0x24, 0x21, 0x62, 0x1f, 0xfc, 0x36, 0x29, 0x37, 0x65, 0x65, 0x24, 0xbb,
	0x99, 0x58, 0x92, 0x08, 0x59, 0x19, 0x8f, 0x5b, 0xd0, 0xc6, 0x62, 0x73, 0x5b, 0xa4, 0x09, 0x3e,
	0x19, 0x8b, 0xf2, 0x5c, 0xd4, 0x6c, 0x1c, 0x82, 0x45, 0xcd, 0xcd, 0xdd, 0xbc, 0x5a, 0xce, 0xd6,
	0xff, 0x13, 0x78, 0x41, 0xcd, 0xf2, 0x42, 0x08, 0x42, 0xa6, 0x5d, 0x49, 0xef, 0x73, 0x5a, 0x67,
	0x92, 0xe2, 0xee, 0x5f, 0xbe, 0xfe, 0xea, 0x37, 0x00, 0x00, 0xff, 0xff, 0x97, 0xf7, 0x48, 0xfe,
	0x97, 0x01, 0x00, 0x00,
}
