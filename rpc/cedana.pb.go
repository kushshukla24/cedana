// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/cedana.proto

package cedana_orch

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Encapsulates data required for the server to make checkpointing decisions
type ClientState struct {
	Timestamp            int64        `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ClientInfo           *ClientInfo  `protobuf:"bytes,2,opt,name=client_info,json=clientInfo,proto3" json:"client_info,omitempty"`
	ProcessInfo          *ProcessInfo `protobuf:"bytes,3,opt,name=process_info,json=processInfo,proto3" json:"process_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f75f48e5338b9f, []int{0}
}

func (m *ClientState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientState.Unmarshal(m, b)
}
func (m *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientState.Marshal(b, m, deterministic)
}
func (m *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(m, src)
}
func (m *ClientState) XXX_Size() int {
	return xxx_messageInfo_ClientState.Size(m)
}
func (m *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(m)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

func (m *ClientState) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ClientState) GetClientInfo() *ClientInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

func (m *ClientState) GetProcessInfo() *ProcessInfo {
	if m != nil {
		return m.ProcessInfo
	}
	return nil
}

type ClientInfo struct {
	RemainingMemory      int32           `protobuf:"varint,1,opt,name=remaining_memory,json=remainingMemory,proto3" json:"remaining_memory,omitempty"`
	Os                   string          `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	Platform             string          `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	Uptime               uint32          `protobuf:"varint,4,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Network              *ConnectionStat `protobuf:"bytes,7,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ClientInfo) Reset()         { *m = ClientInfo{} }
func (m *ClientInfo) String() string { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()    {}
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f75f48e5338b9f, []int{1}
}

func (m *ClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInfo.Unmarshal(m, b)
}
func (m *ClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInfo.Marshal(b, m, deterministic)
}
func (m *ClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInfo.Merge(m, src)
}
func (m *ClientInfo) XXX_Size() int {
	return xxx_messageInfo_ClientInfo.Size(m)
}
func (m *ClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInfo proto.InternalMessageInfo

func (m *ClientInfo) GetRemainingMemory() int32 {
	if m != nil {
		return m.RemainingMemory
	}
	return 0
}

func (m *ClientInfo) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *ClientInfo) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *ClientInfo) GetUptime() uint32 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *ClientInfo) GetNetwork() *ConnectionStat {
	if m != nil {
		return m.Network
	}
	return nil
}

type DockerInfo struct {
	CgroupMem            uint32   `protobuf:"varint,1,opt,name=cgroup_mem,json=cgroupMem,proto3" json:"cgroup_mem,omitempty"`
	CgroupCpu            uint32   `protobuf:"varint,2,opt,name=cgroup_cpu,json=cgroupCpu,proto3" json:"cgroup_cpu,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DockerInfo) Reset()         { *m = DockerInfo{} }
func (m *DockerInfo) String() string { return proto.CompactTextString(m) }
func (*DockerInfo) ProtoMessage()    {}
func (*DockerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f75f48e5338b9f, []int{2}
}

func (m *DockerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DockerInfo.Unmarshal(m, b)
}
func (m *DockerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DockerInfo.Marshal(b, m, deterministic)
}
func (m *DockerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DockerInfo.Merge(m, src)
}
func (m *DockerInfo) XXX_Size() int {
	return xxx_messageInfo_DockerInfo.Size(m)
}
func (m *DockerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DockerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DockerInfo proto.InternalMessageInfo

func (m *DockerInfo) GetCgroupMem() uint32 {
	if m != nil {
		return m.CgroupMem
	}
	return 0
}

func (m *DockerInfo) GetCgroupCpu() uint32 {
	if m != nil {
		return m.CgroupCpu
	}
	return 0
}

type NetworkAddr struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkAddr) Reset()         { *m = NetworkAddr{} }
func (m *NetworkAddr) String() string { return proto.CompactTextString(m) }
func (*NetworkAddr) ProtoMessage()    {}
func (*NetworkAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f75f48e5338b9f, []int{3}
}

func (m *NetworkAddr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkAddr.Unmarshal(m, b)
}
func (m *NetworkAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkAddr.Marshal(b, m, deterministic)
}
func (m *NetworkAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkAddr.Merge(m, src)
}
func (m *NetworkAddr) XXX_Size() int {
	return xxx_messageInfo_NetworkAddr.Size(m)
}
func (m *NetworkAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkAddr.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkAddr proto.InternalMessageInfo

func (m *NetworkAddr) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *NetworkAddr) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// lifted from gopsutil
type ConnectionStat struct {
	Fd                   uint32       `protobuf:"varint,1,opt,name=fd,proto3" json:"fd,omitempty"`
	Family               uint32       `protobuf:"varint,2,opt,name=family,proto3" json:"family,omitempty"`
	Type                 uint32       `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Laddr                *NetworkAddr `protobuf:"bytes,4,opt,name=laddr,proto3" json:"laddr,omitempty"`
	Raddr                *NetworkAddr `protobuf:"bytes,5,opt,name=raddr,proto3" json:"raddr,omitempty"`
	Status               string       `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Uids                 []int32      `protobuf:"varint,7,rep,packed,name=uids,proto3" json:"uids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConnectionStat) Reset()         { *m = ConnectionStat{} }
func (m *ConnectionStat) String() string { return proto.CompactTextString(m) }
func (*ConnectionStat) ProtoMessage()    {}
func (*ConnectionStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f75f48e5338b9f, []int{4}
}

func (m *ConnectionStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionStat.Unmarshal(m, b)
}
func (m *ConnectionStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionStat.Marshal(b, m, deterministic)
}
func (m *ConnectionStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionStat.Merge(m, src)
}
func (m *ConnectionStat) XXX_Size() int {
	return xxx_messageInfo_ConnectionStat.Size(m)
}
func (m *ConnectionStat) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionStat.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionStat proto.InternalMessageInfo

func (m *ConnectionStat) GetFd() uint32 {
	if m != nil {
		return m.Fd
	}
	return 0
}

func (m *ConnectionStat) GetFamily() uint32 {
	if m != nil {
		return m.Family
	}
	return 0
}

func (m *ConnectionStat) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ConnectionStat) GetLaddr() *NetworkAddr {
	if m != nil {
		return m.Laddr
	}
	return nil
}

func (m *ConnectionStat) GetRaddr() *NetworkAddr {
	if m != nil {
		return m.Raddr
	}
	return nil
}

func (m *ConnectionStat) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ConnectionStat) GetUids() []int32 {
	if m != nil {
		return m.Uids
	}
	return nil
}

type ProcessInfo struct {
	ProcessPid           uint32   `protobuf:"varint,1,opt,name=process_pid,json=processPid,proto3" json:"process_pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessInfo) Reset()         { *m = ProcessInfo{} }
func (m *ProcessInfo) String() string { return proto.CompactTextString(m) }
func (*ProcessInfo) ProtoMessage()    {}
func (*ProcessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f75f48e5338b9f, []int{5}
}

func (m *ProcessInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessInfo.Unmarshal(m, b)
}
func (m *ProcessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessInfo.Marshal(b, m, deterministic)
}
func (m *ProcessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessInfo.Merge(m, src)
}
func (m *ProcessInfo) XXX_Size() int {
	return xxx_messageInfo_ProcessInfo.Size(m)
}
func (m *ProcessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessInfo proto.InternalMessageInfo

func (m *ProcessInfo) GetProcessPid() uint32 {
	if m != nil {
		return m.ProcessPid
	}
	return 0
}

type ConfigClient struct {
	DumpFrequency        uint32   `protobuf:"varint,1,opt,name=dump_frequency,json=dumpFrequency,proto3" json:"dump_frequency,omitempty"`
	OrchestratorInfo     string   `protobuf:"bytes,2,opt,name=orchestrator_info,json=orchestratorInfo,proto3" json:"orchestrator_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigClient) Reset()         { *m = ConfigClient{} }
func (m *ConfigClient) String() string { return proto.CompactTextString(m) }
func (*ConfigClient) ProtoMessage()    {}
func (*ConfigClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f75f48e5338b9f, []int{6}
}

func (m *ConfigClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigClient.Unmarshal(m, b)
}
func (m *ConfigClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigClient.Marshal(b, m, deterministic)
}
func (m *ConfigClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigClient.Merge(m, src)
}
func (m *ConfigClient) XXX_Size() int {
	return xxx_messageInfo_ConfigClient.Size(m)
}
func (m *ConfigClient) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigClient.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigClient proto.InternalMessageInfo

func (m *ConfigClient) GetDumpFrequency() uint32 {
	if m != nil {
		return m.DumpFrequency
	}
	return 0
}

func (m *ConfigClient) GetOrchestratorInfo() string {
	if m != nil {
		return m.OrchestratorInfo
	}
	return ""
}

type ClientStateAck struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Acked                bool     `protobuf:"varint,2,opt,name=acked,proto3" json:"acked,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStateAck) Reset()         { *m = ClientStateAck{} }
func (m *ClientStateAck) String() string { return proto.CompactTextString(m) }
func (*ClientStateAck) ProtoMessage()    {}
func (*ClientStateAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f75f48e5338b9f, []int{7}
}

func (m *ClientStateAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStateAck.Unmarshal(m, b)
}
func (m *ClientStateAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStateAck.Marshal(b, m, deterministic)
}
func (m *ClientStateAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStateAck.Merge(m, src)
}
func (m *ClientStateAck) XXX_Size() int {
	return xxx_messageInfo_ClientStateAck.Size(m)
}
func (m *ClientStateAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStateAck.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStateAck proto.InternalMessageInfo

func (m *ClientStateAck) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ClientStateAck) GetAcked() bool {
	if m != nil {
		return m.Acked
	}
	return false
}

type ClientCommand struct {
	Checkpoint           bool                `protobuf:"varint,1,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
	Restore              bool                `protobuf:"varint,2,opt,name=restore,proto3" json:"restore,omitempty"`
	Transfer             *TransferCheckpoint `protobuf:"bytes,3,opt,name=transfer,proto3" json:"transfer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ClientCommand) Reset()         { *m = ClientCommand{} }
func (m *ClientCommand) String() string { return proto.CompactTextString(m) }
func (*ClientCommand) ProtoMessage()    {}
func (*ClientCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f75f48e5338b9f, []int{8}
}

func (m *ClientCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientCommand.Unmarshal(m, b)
}
func (m *ClientCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientCommand.Marshal(b, m, deterministic)
}
func (m *ClientCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientCommand.Merge(m, src)
}
func (m *ClientCommand) XXX_Size() int {
	return xxx_messageInfo_ClientCommand.Size(m)
}
func (m *ClientCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientCommand.DiscardUnknown(m)
}

var xxx_messageInfo_ClientCommand proto.InternalMessageInfo

func (m *ClientCommand) GetCheckpoint() bool {
	if m != nil {
		return m.Checkpoint
	}
	return false
}

func (m *ClientCommand) GetRestore() bool {
	if m != nil {
		return m.Restore
	}
	return false
}

func (m *ClientCommand) GetTransfer() *TransferCheckpoint {
	if m != nil {
		return m.Transfer
	}
	return nil
}

type TransferCheckpoint struct {
	Transfer             bool     `protobuf:"varint,1,opt,name=transfer,proto3" json:"transfer,omitempty"`
	Location             string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferCheckpoint) Reset()         { *m = TransferCheckpoint{} }
func (m *TransferCheckpoint) String() string { return proto.CompactTextString(m) }
func (*TransferCheckpoint) ProtoMessage()    {}
func (*TransferCheckpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f75f48e5338b9f, []int{9}
}

func (m *TransferCheckpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferCheckpoint.Unmarshal(m, b)
}
func (m *TransferCheckpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferCheckpoint.Marshal(b, m, deterministic)
}
func (m *TransferCheckpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferCheckpoint.Merge(m, src)
}
func (m *TransferCheckpoint) XXX_Size() int {
	return xxx_messageInfo_TransferCheckpoint.Size(m)
}
func (m *TransferCheckpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferCheckpoint.DiscardUnknown(m)
}

var xxx_messageInfo_TransferCheckpoint proto.InternalMessageInfo

func (m *TransferCheckpoint) GetTransfer() bool {
	if m != nil {
		return m.Transfer
	}
	return false
}

func (m *TransferCheckpoint) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientState)(nil), "cedana.ClientState")
	proto.RegisterType((*ClientInfo)(nil), "cedana.ClientInfo")
	proto.RegisterType((*DockerInfo)(nil), "cedana.DockerInfo")
	proto.RegisterType((*NetworkAddr)(nil), "cedana.NetworkAddr")
	proto.RegisterType((*ConnectionStat)(nil), "cedana.ConnectionStat")
	proto.RegisterType((*ProcessInfo)(nil), "cedana.ProcessInfo")
	proto.RegisterType((*ConfigClient)(nil), "cedana.ConfigClient")
	proto.RegisterType((*ClientStateAck)(nil), "cedana.ClientStateAck")
	proto.RegisterType((*ClientCommand)(nil), "cedana.ClientCommand")
	proto.RegisterType((*TransferCheckpoint)(nil), "cedana.TransferCheckpoint")
}

func init() { proto.RegisterFile("rpc/cedana.proto", fileDescriptor_05f75f48e5338b9f) }

var fileDescriptor_05f75f48e5338b9f = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdd, 0x6a, 0x1b, 0x3d,
	0x10, 0xcd, 0x3a, 0xf1, 0xdf, 0x38, 0x76, 0xf2, 0x29, 0xf9, 0xc2, 0x12, 0xfa, 0x63, 0x16, 0x0a,
	0x0e, 0xa5, 0x6e, 0x9a, 0x40, 0xee, 0x5a, 0x48, 0x1c, 0x02, 0x2d, 0x4d, 0x09, 0x6a, 0xaf, 0x7a,
	0x63, 0x14, 0xad, 0xd6, 0x11, 0xde, 0x95, 0x54, 0xad, 0x96, 0xe2, 0xcb, 0xbe, 0x45, 0x5f, 0xa2,
	0x4f, 0x52, 0xe8, 0x33, 0x15, 0x69, 0xb5, 0xeb, 0x0d, 0x0d, 0xed, 0x9d, 0xe6, 0xcc, 0xcc, 0xd1,
	0x99, 0xd1, 0x8c, 0x60, 0x57, 0x2b, 0xfa, 0x92, 0xb2, 0x98, 0x08, 0x32, 0x55, 0x5a, 0x1a, 0x89,
	0x3a, 0xa5, 0x15, 0x7d, 0x0f, 0x60, 0x30, 0x4b, 0x39, 0x13, 0xe6, 0xa3, 0x21, 0x86, 0xa1, 0x47,
	0xd0, 0x37, 0x3c, 0x63, 0xb9, 0x21, 0x99, 0x0a, 0x83, 0x71, 0x30, 0xd9, 0xc4, 0x6b, 0x00, 0x9d,
	0xc2, 0x80, 0xba, 0xe0, 0x39, 0x17, 0x89, 0x0c, 0x5b, 0xe3, 0x60, 0x32, 0x38, 0x41, 0x53, 0xcf,
	0x5c, 0xf2, 0xbc, 0x15, 0x89, 0xc4, 0x40, 0xeb, 0x33, 0x3a, 0x83, 0x6d, 0xa5, 0x25, 0x65, 0x79,
	0x5e, 0x66, 0x6d, 0xba, 0xac, 0xbd, 0x2a, 0xeb, 0xa6, 0xf4, 0xb9, 0xb4, 0x81, 0x5a, 0x1b, 0xd1,
	0x8f, 0x00, 0x60, 0x4d, 0x89, 0x8e, 0x60, 0x57, 0xb3, 0x8c, 0x70, 0xc1, 0xc5, 0x62, 0x9e, 0xb1,
	0x4c, 0xea, 0x95, 0x13, 0xd8, 0xc6, 0x3b, 0x35, 0x7e, 0xed, 0x60, 0x34, 0x82, 0x96, 0xcc, 0x9d,
	0xba, 0x3e, 0x6e, 0xc9, 0x1c, 0x1d, 0x42, 0x4f, 0xa5, 0xc4, 0x24, 0x52, 0x67, 0xee, 0xf6, 0x3e,
	0xae, 0x6d, 0x74, 0x00, 0x9d, 0x42, 0xd9, 0x0a, 0xc3, 0xad, 0x71, 0x30, 0x19, 0x62, 0x6f, 0xa1,
	0x63, 0xe8, 0x0a, 0x66, 0xbe, 0x4a, 0xbd, 0x0c, 0xbb, 0x4e, 0xf0, 0x41, 0x5d, 0xa6, 0x14, 0x82,
	0x51, 0xc3, 0xa5, 0xb0, 0x2d, 0xc3, 0x55, 0x58, 0xf4, 0x0e, 0xe0, 0x52, 0xd2, 0x25, 0xd3, 0x4e,
	0xee, 0x63, 0x00, 0xba, 0xd0, 0xb2, 0x50, 0x56, 0xab, 0x13, 0x3a, 0xc4, 0xfd, 0x12, 0xb9, 0x66,
	0x59, 0xc3, 0x4d, 0x55, 0xe1, 0xa4, 0xd6, 0xee, 0x99, 0x2a, 0xa2, 0x57, 0x30, 0xf8, 0x50, 0xd2,
	0x9e, 0xc7, 0xb1, 0xb6, 0x05, 0xf1, 0xf2, 0x39, 0xfa, 0xb8, 0xc5, 0x15, 0x42, 0xb0, 0xa5, 0xa4,
	0x36, 0x3e, 0xcf, 0x9d, 0xa3, 0x5f, 0x01, 0x8c, 0xee, 0x4b, 0xb3, 0x69, 0x49, 0xec, 0xef, 0x6e,
	0x25, 0xb1, 0xad, 0x35, 0x21, 0x19, 0x4f, 0x57, 0x3e, 0xd1, 0x5b, 0x96, 0xce, 0xac, 0x14, 0x73,
	0xbd, 0x19, 0x62, 0x77, 0x46, 0x47, 0xd0, 0x4e, 0x49, 0x1c, 0x6b, 0xd7, 0x96, 0xc6, 0x73, 0x35,
	0x64, 0xe1, 0x32, 0xc2, 0x86, 0x6a, 0x17, 0xda, 0xfe, 0x4b, 0xa8, 0x8b, 0xb0, 0x0a, 0x72, 0x43,
	0x4c, 0x91, 0x87, 0x1d, 0x57, 0x8c, 0xb7, 0xac, 0x82, 0x82, 0xc7, 0x79, 0xd8, 0x1d, 0x6f, 0x4e,
	0xda, 0xd8, 0x9d, 0xa3, 0x29, 0x0c, 0x1a, 0xb3, 0x81, 0x9e, 0x42, 0x35, 0x1d, 0x73, 0xc5, 0xab,
	0xaa, 0xc0, 0x43, 0x37, 0x3c, 0x8e, 0x6e, 0x61, 0x7b, 0x26, 0x45, 0xc2, 0x17, 0xe5, 0xd0, 0xa0,
	0x67, 0x30, 0x8a, 0x8b, 0x4c, 0xcd, 0x13, 0xcd, 0xbe, 0x14, 0x4c, 0xd0, 0x95, 0xcf, 0x19, 0x5a,
	0xf4, 0xaa, 0x02, 0xd1, 0x73, 0xf8, 0x4f, 0x6a, 0x7a, 0xc7, 0x72, 0xa3, 0x89, 0x91, 0x7a, 0x3d,
	0xd9, 0x7d, 0xbc, 0xdb, 0x74, 0xb8, 0x99, 0xbc, 0x84, 0x51, 0x63, 0x5b, 0xce, 0xe9, 0xf2, 0x1f,
	0x0b, 0xb3, 0x0f, 0x6d, 0x42, 0x97, 0x2c, 0x76, 0x84, 0x3d, 0x5c, 0x1a, 0xd1, 0xb7, 0x00, 0x86,
	0x25, 0xcd, 0x4c, 0x66, 0x19, 0x11, 0x31, 0x7a, 0x02, 0x40, 0xef, 0x18, 0x5d, 0x2a, 0xc9, 0x85,
	0x71, 0x34, 0x3d, 0xdc, 0x40, 0x50, 0x08, 0x5d, 0xcd, 0x72, 0x23, 0x35, 0xf3, 0x4c, 0x95, 0x89,
	0xce, 0xa0, 0x67, 0x34, 0x11, 0x79, 0xc2, 0xb4, 0xdf, 0xac, 0xc3, 0xaa, 0xff, 0x9f, 0x3c, 0x3e,
	0xab, 0x79, 0x70, 0x1d, 0x1b, 0xbd, 0x07, 0xf4, 0xa7, 0xdf, 0x6e, 0x4a, 0xcd, 0x56, 0xaa, 0xa8,
	0x6d, 0xeb, 0x4b, 0x25, 0x25, 0x76, 0xba, 0x7c, 0x7f, 0x6a, 0xfb, 0xe4, 0x67, 0x00, 0x9d, 0x99,
	0xbb, 0x15, 0xbd, 0x86, 0x11, 0x66, 0x0b, 0x9e, 0x1b, 0xa6, 0xfd, 0x43, 0xec, 0xdd, 0xff, 0x20,
	0x5c, 0xeb, 0x0e, 0xf7, 0x1b, 0xeb, 0x54, 0xbf, 0x59, 0xb4, 0x81, 0xde, 0xc0, 0x00, 0x33, 0x2a,
	0x75, 0x5c, 0xfe, 0x47, 0x0f, 0xe6, 0x1e, 0x3c, 0x00, 0x9e, 0xd3, 0x65, 0xb4, 0x31, 0x09, 0xd0,
	0x05, 0x8c, 0x6e, 0x64, 0x9a, 0x5e, 0x49, 0x5d, 0xf5, 0xf6, 0x41, 0x8a, 0xff, 0xef, 0x83, 0x3e,
	0xd6, 0x32, 0x1c, 0x07, 0x17, 0x3b, 0x9f, 0x87, 0x53, 0xff, 0x5d, 0xbe, 0xb0, 0x23, 0x70, 0xdb,
	0x71, 0x9f, 0xe6, 0xe9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0xd8, 0xee, 0x0a, 0x48, 0x05,
	0x00, 0x00,
}
