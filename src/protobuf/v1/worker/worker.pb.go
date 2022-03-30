// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: v1/worker/worker.proto

package worker

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

type WorkerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policytype  string `protobuf:"bytes,1,opt,name=policytype,proto3" json:"policytype,omitempty"`
	Req         string `protobuf:"bytes,2,opt,name=req,proto3" json:"req,omitempty"`
	Logfile     string `protobuf:"bytes,3,opt,name=logfile,proto3" json:"logfile,omitempty"`
	Namespace   string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Clustername string `protobuf:"bytes,5,opt,name=clustername,proto3" json:"clustername,omitempty"`
	Labels      string `protobuf:"bytes,6,opt,name=labels,proto3" json:"labels,omitempty"`
}

func (x *WorkerRequest) Reset() {
	*x = WorkerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_worker_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerRequest) ProtoMessage() {}

func (x *WorkerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_worker_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerRequest.ProtoReflect.Descriptor instead.
func (*WorkerRequest) Descriptor() ([]byte, []int) {
	return file_v1_worker_worker_proto_rawDescGZIP(), []int{0}
}

func (x *WorkerRequest) GetPolicytype() string {
	if x != nil {
		return x.Policytype
	}
	return ""
}

func (x *WorkerRequest) GetReq() string {
	if x != nil {
		return x.Req
	}
	return ""
}

func (x *WorkerRequest) GetLogfile() string {
	if x != nil {
		return x.Logfile
	}
	return ""
}

func (x *WorkerRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *WorkerRequest) GetClustername() string {
	if x != nil {
		return x.Clustername
	}
	return ""
}

func (x *WorkerRequest) GetLabels() string {
	if x != nil {
		return x.Labels
	}
	return ""
}

type WorkerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res string `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *WorkerResponse) Reset() {
	*x = WorkerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_worker_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerResponse) ProtoMessage() {}

func (x *WorkerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_worker_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerResponse.ProtoReflect.Descriptor instead.
func (*WorkerResponse) Descriptor() ([]byte, []int) {
	return file_v1_worker_worker_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerResponse) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

var File_v1_worker_worker_proto protoreflect.FileDescriptor

var file_v1_worker_worker_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x22, 0xb3, 0x01, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x22, 0x0a, 0x0e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x73, 0x32, 0x8b, 0x02,
	0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x76, 0x31,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x75, 0x6b, 0x6e,
	0x6f, 0x78, 0x2f, 0x6b, 0x6e, 0x6f, 0x78, 0x41, 0x75, 0x74, 0x6f, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_worker_worker_proto_rawDescOnce sync.Once
	file_v1_worker_worker_proto_rawDescData = file_v1_worker_worker_proto_rawDesc
)

func file_v1_worker_worker_proto_rawDescGZIP() []byte {
	file_v1_worker_worker_proto_rawDescOnce.Do(func() {
		file_v1_worker_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_worker_worker_proto_rawDescData)
	})
	return file_v1_worker_worker_proto_rawDescData
}

var file_v1_worker_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_worker_worker_proto_goTypes = []interface{}{
	(*WorkerRequest)(nil),  // 0: v1.worker.WorkerRequest
	(*WorkerResponse)(nil), // 1: v1.worker.WorkerResponse
}
var file_v1_worker_worker_proto_depIdxs = []int32{
	0, // 0: v1.worker.Worker.GetWorkerStatus:input_type -> v1.worker.WorkerRequest
	0, // 1: v1.worker.Worker.Start:input_type -> v1.worker.WorkerRequest
	0, // 2: v1.worker.Worker.Stop:input_type -> v1.worker.WorkerRequest
	0, // 3: v1.worker.Worker.Convert:input_type -> v1.worker.WorkerRequest
	1, // 4: v1.worker.Worker.GetWorkerStatus:output_type -> v1.worker.WorkerResponse
	1, // 5: v1.worker.Worker.Start:output_type -> v1.worker.WorkerResponse
	1, // 6: v1.worker.Worker.Stop:output_type -> v1.worker.WorkerResponse
	1, // 7: v1.worker.Worker.Convert:output_type -> v1.worker.WorkerResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_worker_worker_proto_init() }
func file_v1_worker_worker_proto_init() {
	if File_v1_worker_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_worker_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerRequest); i {
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
		file_v1_worker_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerResponse); i {
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
			RawDescriptor: file_v1_worker_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_worker_worker_proto_goTypes,
		DependencyIndexes: file_v1_worker_worker_proto_depIdxs,
		MessageInfos:      file_v1_worker_worker_proto_msgTypes,
	}.Build()
	File_v1_worker_worker_proto = out.File
	file_v1_worker_worker_proto_rawDesc = nil
	file_v1_worker_worker_proto_goTypes = nil
	file_v1_worker_worker_proto_depIdxs = nil
}
