// Code generated by protoc-gen-go. DO NOT EDIT.
// source: appconfig.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type InstanceOpsCmd struct {
	Counter uint32 `protobuf:"varint,2,opt,name=counter" json:"counter,omitempty"`
	OpsTime string `protobuf:"bytes,4,opt,name=opsTime" json:"opsTime,omitempty"`
}

func (m *InstanceOpsCmd) Reset()                    { *m = InstanceOpsCmd{} }
func (m *InstanceOpsCmd) String() string            { return proto.CompactTextString(m) }
func (*InstanceOpsCmd) ProtoMessage()               {}
func (*InstanceOpsCmd) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *InstanceOpsCmd) GetCounter() uint32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *InstanceOpsCmd) GetOpsTime() string {
	if m != nil {
		return m.OpsTime
	}
	return ""
}

type AppInstanceConfig struct {
	Uuidandversion *UUIDandVersion `protobuf:"bytes,1,opt,name=uuidandversion" json:"uuidandversion,omitempty"`
	Displayname    string          `protobuf:"bytes,2,opt,name=displayname" json:"displayname,omitempty"`
	Fixedresources *VmConfig       `protobuf:"bytes,3,opt,name=fixedresources" json:"fixedresources,omitempty"`
	Drives         []*Drive        `protobuf:"bytes,4,rep,name=drives" json:"drives,omitempty"`
	Activate       bool            `protobuf:"varint,5,opt,name=activate" json:"activate,omitempty"`
	// NetworkAdapter are virtual adapters assigned to the application
	// Physical adapters such as eth1 are part of Adapter
	Interfaces []*NetworkAdapter `protobuf:"bytes,6,rep,name=interfaces" json:"interfaces,omitempty"`
	Adapters   []*Adapter        `protobuf:"bytes,7,rep,name=adapters" json:"adapters,omitempty"`
	Restart    *InstanceOpsCmd   `protobuf:"bytes,9,opt,name=restart" json:"restart,omitempty"`
	// The device behavior for a purge command is to restart the domU.
	// with the disks/drives recreated from the downloaded images
	// (whether preserve is set or not).
	//    if the manifest is changed with purge option, new manifest will
	//    be used. Device doesn't know what has changed, it will get the
	//    changed config.
	//
	//    if disks section have changed will be purged automatically.
	//    phase 1: we would purge all disks irrespective preserve flag
	Purge *InstanceOpsCmd `protobuf:"bytes,10,opt,name=purge" json:"purge,omitempty"`
}

func (m *AppInstanceConfig) Reset()                    { *m = AppInstanceConfig{} }
func (m *AppInstanceConfig) String() string            { return proto.CompactTextString(m) }
func (*AppInstanceConfig) ProtoMessage()               {}
func (*AppInstanceConfig) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *AppInstanceConfig) GetUuidandversion() *UUIDandVersion {
	if m != nil {
		return m.Uuidandversion
	}
	return nil
}

func (m *AppInstanceConfig) GetDisplayname() string {
	if m != nil {
		return m.Displayname
	}
	return ""
}

func (m *AppInstanceConfig) GetFixedresources() *VmConfig {
	if m != nil {
		return m.Fixedresources
	}
	return nil
}

func (m *AppInstanceConfig) GetDrives() []*Drive {
	if m != nil {
		return m.Drives
	}
	return nil
}

func (m *AppInstanceConfig) GetActivate() bool {
	if m != nil {
		return m.Activate
	}
	return false
}

func (m *AppInstanceConfig) GetInterfaces() []*NetworkAdapter {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *AppInstanceConfig) GetAdapters() []*Adapter {
	if m != nil {
		return m.Adapters
	}
	return nil
}

func (m *AppInstanceConfig) GetRestart() *InstanceOpsCmd {
	if m != nil {
		return m.Restart
	}
	return nil
}

func (m *AppInstanceConfig) GetPurge() *InstanceOpsCmd {
	if m != nil {
		return m.Purge
	}
	return nil
}

func init() {
	proto.RegisterType((*InstanceOpsCmd)(nil), "InstanceOpsCmd")
	proto.RegisterType((*AppInstanceConfig)(nil), "AppInstanceConfig")
}

func init() { proto.RegisterFile("appconfig.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcb, 0x6f, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0xa4, 0x4d, 0x36, 0x13, 0x35, 0x11, 0x3e, 0x59, 0x95, 0x80, 0x55, 0x05, 0xd2,
	0x72, 0x71, 0x44, 0x39, 0x70, 0xa5, 0x34, 0x97, 0x5e, 0x40, 0xb2, 0x68, 0x0f, 0xdc, 0x5c, 0x7b,
	0xb2, 0x58, 0xd4, 0x0f, 0xf9, 0xb1, 0x40, 0xff, 0x00, 0xfe, 0x6e, 0xb4, 0xaf, 0x2a, 0x44, 0x1c,
	0xbf, 0xef, 0xfb, 0xcd, 0x8c, 0x35, 0x1e, 0xd8, 0x08, 0xef, 0xa5, 0xb3, 0x7b, 0xdd, 0x30, 0x1f,
	0x5c, 0x72, 0xe7, 0x1b, 0x85, 0xad, 0x74, 0xc6, 0x38, 0x3b, 0x1a, 0x67, 0x31, 0xb9, 0x20, 0x1a,
	0x1c, 0x65, 0xd9, 0x9a, 0x89, 0xb4, 0x98, 0x0e, 0x4b, 0x2f, 0x76, 0xb0, 0xbe, 0xb1, 0x31, 0x09,
	0x2b, 0xf1, 0x8b, 0x8f, 0xd7, 0x46, 0x11, 0x0a, 0x0b, 0xe9, 0xb2, 0x4d, 0x18, 0xe8, 0xb3, 0xaa,
	0xa8, 0xcf, 0xf8, 0x24, 0xbb, 0xc4, 0xf9, 0xf8, 0x55, 0x1b, 0xa4, 0x27, 0x55, 0x51, 0x2f, 0xf9,
	0x24, 0x2f, 0xfe, 0xcc, 0xe0, 0xf9, 0x95, 0xf7, 0x53, 0xa7, 0xeb, 0x7e, 0x02, 0xf9, 0x00, 0xeb,
	0x9c, 0xb5, 0x12, 0x56, 0xb5, 0x18, 0xa2, 0x76, 0x96, 0x16, 0x55, 0x51, 0xaf, 0x2e, 0x37, 0xec,
	0xf6, 0xf6, 0x66, 0x27, 0xac, 0xba, 0x1b, 0x6c, 0x7e, 0x84, 0x91, 0x0a, 0x56, 0x4a, 0x47, 0xff,
	0x20, 0x7e, 0x5b, 0x61, 0xb0, 0x7f, 0xc6, 0x92, 0x1f, 0x5a, 0xe4, 0x1d, 0xac, 0xf7, 0xfa, 0x17,
	0xaa, 0x80, 0xd1, 0xe5, 0x20, 0x31, 0xd2, 0x59, 0xdf, 0x7a, 0xc9, 0xee, 0xcc, 0x30, 0x9d, 0x1f,
	0x01, 0xe4, 0x25, 0xcc, 0x55, 0xd0, 0x2d, 0x46, 0x7a, 0x52, 0xcd, 0xea, 0xd5, 0xe5, 0x9c, 0xed,
	0x3a, 0xc9, 0x47, 0x97, 0x9c, 0x43, 0x29, 0x64, 0xd2, 0xad, 0x48, 0x48, 0x4f, 0xab, 0xa2, 0x2e,
	0xf9, 0x93, 0x26, 0x5b, 0x00, 0xdd, 0xad, 0x60, 0x2f, 0xba, 0x51, 0xf3, 0xbe, 0x7e, 0xc3, 0x3e,
	0x63, 0xfa, 0xe9, 0xc2, 0x8f, 0x2b, 0x25, 0x7c, 0xc2, 0xc0, 0x0f, 0x10, 0xf2, 0x1a, 0x4a, 0x31,
	0xd8, 0x91, 0x2e, 0x7a, 0xbc, 0x64, 0x13, 0xf7, 0x94, 0x90, 0xb7, 0xb0, 0x08, 0x18, 0x93, 0x08,
	0x89, 0x2e, 0xc7, 0xcd, 0xfc, 0xfb, 0x19, 0x7c, 0xca, 0xc9, 0x1b, 0x38, 0xf5, 0x39, 0x34, 0x48,
	0xe1, 0xff, 0xe0, 0x90, 0x7e, 0xfa, 0x08, 0xaf, 0xa4, 0x33, 0xec, 0x11, 0x15, 0x2a, 0xc1, 0xe4,
	0x83, 0xcb, 0x8a, 0xe5, 0x88, 0xa1, 0xd5, 0x72, 0x3c, 0x86, 0x6f, 0x2f, 0x1a, 0x9d, 0xbe, 0xe7,
	0x7b, 0x26, 0x9d, 0xd9, 0x0e, 0xdc, 0x56, 0x78, 0xbd, 0x7d, 0x1c, 0xce, 0xe2, 0x7e, 0xde, 0x53,
	0xef, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xfb, 0x3f, 0xb7, 0x65, 0x02, 0x00, 0x00,
}
