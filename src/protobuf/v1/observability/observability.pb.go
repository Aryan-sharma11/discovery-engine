// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: v1/observability/observability.proto

package observability

import (
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodName       string `protobuf:"bytes,1,opt,name=PodName,proto3" json:"PodName,omitempty"`
	NameSpace     string `protobuf:"bytes,2,opt,name=NameSpace,proto3" json:"NameSpace,omitempty"`
	ClusterName   string `protobuf:"bytes,3,opt,name=ClusterName,proto3" json:"ClusterName,omitempty"`
	Label         string `protobuf:"bytes,4,opt,name=Label,proto3" json:"Label,omitempty"`
	ContainerName string `protobuf:"bytes,5,opt,name=ContainerName,proto3" json:"ContainerName,omitempty"`
	Type          string `protobuf:"bytes,6,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observability_observability_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observability_observability_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_v1_observability_observability_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *Request) GetNameSpace() string {
	if x != nil {
		return x.NameSpace
	}
	return ""
}

func (x *Request) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *Request) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Request) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *Request) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodName           string                    `protobuf:"bytes,1,opt,name=PodName,proto3" json:"PodName,omitempty"`
	ClusterName       string                    `protobuf:"bytes,2,opt,name=ClusterName,proto3" json:"ClusterName,omitempty"`
	Namespace         string                    `protobuf:"bytes,3,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	Label             string                    `protobuf:"bytes,4,opt,name=Label,proto3" json:"Label,omitempty"`
	ContainerName     string                    `protobuf:"bytes,5,opt,name=ContainerName,proto3" json:"ContainerName,omitempty"`
	ProcessData       []*SysProcFileSummaryData `protobuf:"bytes,6,rep,name=ProcessData,proto3" json:"ProcessData,omitempty"`
	FileData          []*SysProcFileSummaryData `protobuf:"bytes,7,rep,name=FileData,proto3" json:"FileData,omitempty"`
	InNwData          []*SysNwSummaryData       `protobuf:"bytes,8,rep,name=InNwData,proto3" json:"InNwData,omitempty"`
	OutNwData         []*SysNwSummaryData       `protobuf:"bytes,9,rep,name=OutNwData,proto3" json:"OutNwData,omitempty"`
	IngressEgressData []*CiliumSummData         `protobuf:"bytes,10,rep,name=IngressEgressData,proto3" json:"IngressEgressData,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observability_observability_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observability_observability_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_v1_observability_observability_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *Response) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *Response) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Response) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Response) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *Response) GetProcessData() []*SysProcFileSummaryData {
	if x != nil {
		return x.ProcessData
	}
	return nil
}

func (x *Response) GetFileData() []*SysProcFileSummaryData {
	if x != nil {
		return x.FileData
	}
	return nil
}

func (x *Response) GetInNwData() []*SysNwSummaryData {
	if x != nil {
		return x.InNwData
	}
	return nil
}

func (x *Response) GetOutNwData() []*SysNwSummaryData {
	if x != nil {
		return x.OutNwData
	}
	return nil
}

func (x *Response) GetIngressEgressData() []*CiliumSummData {
	if x != nil {
		return x.IngressEgressData
	}
	return nil
}

type SysProcFileSummaryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcName       string `protobuf:"bytes,1,opt,name=ProcName,proto3" json:"ProcName,omitempty"`
	ParentProcName string `protobuf:"bytes,2,opt,name=ParentProcName,proto3" json:"ParentProcName,omitempty"`
	Count          string `protobuf:"bytes,3,opt,name=Count,proto3" json:"Count,omitempty"`
	UpdatedTime    string `protobuf:"bytes,4,opt,name=UpdatedTime,proto3" json:"UpdatedTime,omitempty"`
	Status         string `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *SysProcFileSummaryData) Reset() {
	*x = SysProcFileSummaryData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observability_observability_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysProcFileSummaryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysProcFileSummaryData) ProtoMessage() {}

func (x *SysProcFileSummaryData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observability_observability_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysProcFileSummaryData.ProtoReflect.Descriptor instead.
func (*SysProcFileSummaryData) Descriptor() ([]byte, []int) {
	return file_v1_observability_observability_proto_rawDescGZIP(), []int{2}
}

func (x *SysProcFileSummaryData) GetProcName() string {
	if x != nil {
		return x.ProcName
	}
	return ""
}

func (x *SysProcFileSummaryData) GetParentProcName() string {
	if x != nil {
		return x.ParentProcName
	}
	return ""
}

func (x *SysProcFileSummaryData) GetCount() string {
	if x != nil {
		return x.Count
	}
	return ""
}

func (x *SysProcFileSummaryData) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *SysProcFileSummaryData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type SysNwSummaryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol  string `protobuf:"bytes,1,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	Command   string `protobuf:"bytes,2,opt,name=Command,proto3" json:"Command,omitempty"`
	IP        string `protobuf:"bytes,3,opt,name=IP,proto3" json:"IP,omitempty"`
	Port      string `protobuf:"bytes,4,opt,name=Port,proto3" json:"Port,omitempty"`
	Labels    string `protobuf:"bytes,5,opt,name=Labels,proto3" json:"Labels,omitempty"`
	Namespace string `protobuf:"bytes,6,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
}

func (x *SysNwSummaryData) Reset() {
	*x = SysNwSummaryData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observability_observability_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysNwSummaryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysNwSummaryData) ProtoMessage() {}

func (x *SysNwSummaryData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observability_observability_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysNwSummaryData.ProtoReflect.Descriptor instead.
func (*SysNwSummaryData) Descriptor() ([]byte, []int) {
	return file_v1_observability_observability_proto_rawDescGZIP(), []int{3}
}

func (x *SysNwSummaryData) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *SysNwSummaryData) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *SysNwSummaryData) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *SysNwSummaryData) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *SysNwSummaryData) GetLabels() string {
	if x != nil {
		return x.Labels
	}
	return ""
}

func (x *SysNwSummaryData) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type CiliumSummData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcDestPod  string `protobuf:"bytes,1,opt,name=SrcDestPod,proto3" json:"SrcDestPod,omitempty"`
	Protocol    string `protobuf:"bytes,2,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	Port        string `protobuf:"bytes,3,opt,name=Port,proto3" json:"Port,omitempty"`
	Count       string `protobuf:"bytes,4,opt,name=Count,proto3" json:"Count,omitempty"`
	UpdatedTime string `protobuf:"bytes,5,opt,name=UpdatedTime,proto3" json:"UpdatedTime,omitempty"`
	Status      string `protobuf:"bytes,6,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *CiliumSummData) Reset() {
	*x = CiliumSummData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observability_observability_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CiliumSummData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CiliumSummData) ProtoMessage() {}

func (x *CiliumSummData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observability_observability_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CiliumSummData.ProtoReflect.Descriptor instead.
func (*CiliumSummData) Descriptor() ([]byte, []int) {
	return file_v1_observability_observability_proto_rawDescGZIP(), []int{4}
}

func (x *CiliumSummData) GetSrcDestPod() string {
	if x != nil {
		return x.SrcDestPod
	}
	return ""
}

func (x *CiliumSummData) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *CiliumSummData) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *CiliumSummData) GetCount() string {
	if x != nil {
		return x.Count
	}
	return ""
}

func (x *CiliumSummData) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *CiliumSummData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type PodNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodName []string `protobuf:"bytes,1,rep,name=PodName,proto3" json:"PodName,omitempty"`
}

func (x *PodNameResponse) Reset() {
	*x = PodNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observability_observability_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodNameResponse) ProtoMessage() {}

func (x *PodNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observability_observability_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodNameResponse.ProtoReflect.Descriptor instead.
func (*PodNameResponse) Descriptor() ([]byte, []int) {
	return file_v1_observability_observability_proto_rawDescGZIP(), []int{5}
}

func (x *PodNameResponse) GetPodName() []string {
	if x != nil {
		return x.PodName
	}
	return nil
}

var File_v1_observability_observability_proto protoreflect.FileDescriptor

var file_v1_observability_observability_proto_rawDesc = []byte{
	0x0a, 0x24, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x76, 0x31, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0xb3, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x84,
	0x04, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6f,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x79, 0x73, 0x50, 0x72, 0x6f,
	0x63, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x44, 0x0a,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2e, 0x53, 0x79, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x08, 0x49, 0x6e, 0x4e, 0x77, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x79, 0x73, 0x4e, 0x77, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x49, 0x6e, 0x4e, 0x77, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x09, 0x4f, 0x75, 0x74, 0x4e, 0x77, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x79, 0x73, 0x4e, 0x77, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x4f, 0x75, 0x74, 0x4e,
	0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4e, 0x0a, 0x11, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x53, 0x75, 0x6d, 0x6d, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x11, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0xac, 0x01, 0x0a, 0x16, 0x53, 0x79, 0x73, 0x50, 0x72, 0x6f,
	0x63, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x4e, 0x77, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x0e, 0x43, 0x69,
	0x6c, 0x69, 0x75, 0x6d, 0x53, 0x75, 0x6d, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a,
	0x53, 0x72, 0x63, 0x44, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x53, 0x72, 0x63, 0x44, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2b, 0x0a, 0x0f,
	0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x9e, 0x01, 0x0a, 0x0d, 0x4f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x07, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x76,
	0x31, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6f, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x75, 0x6b, 0x6e, 0x6f,
	0x78, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x2d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2d, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_observability_observability_proto_rawDescOnce sync.Once
	file_v1_observability_observability_proto_rawDescData = file_v1_observability_observability_proto_rawDesc
)

func file_v1_observability_observability_proto_rawDescGZIP() []byte {
	file_v1_observability_observability_proto_rawDescOnce.Do(func() {
		file_v1_observability_observability_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_observability_observability_proto_rawDescData)
	})
	return file_v1_observability_observability_proto_rawDescData
}

var file_v1_observability_observability_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_observability_observability_proto_goTypes = []interface{}{
	(*Request)(nil),                // 0: v1.observability.Request
	(*Response)(nil),               // 1: v1.observability.Response
	(*SysProcFileSummaryData)(nil), // 2: v1.observability.SysProcFileSummaryData
	(*SysNwSummaryData)(nil),       // 3: v1.observability.SysNwSummaryData
	(*CiliumSummData)(nil),         // 4: v1.observability.CiliumSummData
	(*PodNameResponse)(nil),        // 5: v1.observability.PodNameResponse
}
var file_v1_observability_observability_proto_depIdxs = []int32{
	2, // 0: v1.observability.Response.ProcessData:type_name -> v1.observability.SysProcFileSummaryData
	2, // 1: v1.observability.Response.FileData:type_name -> v1.observability.SysProcFileSummaryData
	3, // 2: v1.observability.Response.InNwData:type_name -> v1.observability.SysNwSummaryData
	3, // 3: v1.observability.Response.OutNwData:type_name -> v1.observability.SysNwSummaryData
	4, // 4: v1.observability.Response.IngressEgressData:type_name -> v1.observability.CiliumSummData
	0, // 5: v1.observability.Observability.Summary:input_type -> v1.observability.Request
	0, // 6: v1.observability.Observability.GetPodNames:input_type -> v1.observability.Request
	1, // 7: v1.observability.Observability.Summary:output_type -> v1.observability.Response
	5, // 8: v1.observability.Observability.GetPodNames:output_type -> v1.observability.PodNameResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_observability_observability_proto_init() }
func file_v1_observability_observability_proto_init() {
	if File_v1_observability_observability_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_observability_observability_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_v1_observability_observability_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_v1_observability_observability_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysProcFileSummaryData); i {
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
		file_v1_observability_observability_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysNwSummaryData); i {
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
		file_v1_observability_observability_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CiliumSummData); i {
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
		file_v1_observability_observability_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodNameResponse); i {
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
			RawDescriptor: file_v1_observability_observability_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_observability_observability_proto_goTypes,
		DependencyIndexes: file_v1_observability_observability_proto_depIdxs,
		MessageInfos:      file_v1_observability_observability_proto_msgTypes,
	}.Build()
	File_v1_observability_observability_proto = out.File
	file_v1_observability_observability_proto_rawDesc = nil
	file_v1_observability_observability_proto_goTypes = nil
	file_v1_observability_observability_proto_depIdxs = nil
}
