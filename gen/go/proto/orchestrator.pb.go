// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/orchestrator.proto

package pb

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orchestrator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orchestrator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_orchestrator_proto_rawDescGZIP(), []int{0}
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxThreads int64 `protobuf:"varint,1,opt,name=maxThreads,proto3" json:"maxThreads,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orchestrator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orchestrator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_proto_orchestrator_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterReq) GetMaxThreads() int64 {
	if x != nil {
		return x.MaxThreads
	}
	return 0
}

type RegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orchestrator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orchestrator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_proto_orchestrator_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ResultResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId   int64   `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Result   float64 `protobuf:"fixed64,2,opt,name=result,proto3" json:"result,omitempty"`
	Error    bool    `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
	SendTime int64   `protobuf:"varint,4,opt,name=sendTime,proto3" json:"sendTime,omitempty"`
}

func (x *ResultResp) Reset() {
	*x = ResultResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orchestrator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResp) ProtoMessage() {}

func (x *ResultResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orchestrator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResp.ProtoReflect.Descriptor instead.
func (*ResultResp) Descriptor() ([]byte, []int) {
	return file_proto_orchestrator_proto_rawDescGZIP(), []int{3}
}

func (x *ResultResp) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *ResultResp) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *ResultResp) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *ResultResp) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

var File_proto_orchestrator_proto protoreflect.FileDescriptor

var file_proto_orchestrator_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6f, 0x6e, 0x69, 0x61, 0x73, 0x31, 0x38, 0x39,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x73, 0x22, 0x1e, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x8c, 0x03, 0x0a, 0x0c, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x83, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x3a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x65, 0x6f, 0x6e, 0x69, 0x61, 0x73, 0x31, 0x38, 0x39, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x3b, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6f, 0x6e,
	0x69, 0x61, 0x73, 0x31, 0x38, 0x39, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x7d, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x39, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6f, 0x6e, 0x69, 0x61, 0x73, 0x31, 0x38, 0x39, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x1a, 0x33, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x65, 0x6f, 0x6e, 0x69, 0x61, 0x73, 0x31, 0x38, 0x39, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x28, 0x01, 0x30, 0x01, 0x12, 0x77, 0x0a, 0x0a, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x33, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6f, 0x6e, 0x69, 0x61, 0x73, 0x31, 0x38, 0x39,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x34,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6f, 0x6e, 0x69,
	0x61, 0x73, 0x31, 0x38, 0x39, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_orchestrator_proto_rawDescOnce sync.Once
	file_proto_orchestrator_proto_rawDescData = file_proto_orchestrator_proto_rawDesc
)

func file_proto_orchestrator_proto_rawDescGZIP() []byte {
	file_proto_orchestrator_proto_rawDescOnce.Do(func() {
		file_proto_orchestrator_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_orchestrator_proto_rawDescData)
	})
	return file_proto_orchestrator_proto_rawDescData
}

var file_proto_orchestrator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_orchestrator_proto_goTypes = []interface{}{
	(*Empty)(nil),        // 0: github.com.eonias189.calculationService.proto.Empty
	(*RegisterReq)(nil),  // 1: github.com.eonias189.calculationService.proto.RegisterReq
	(*RegisterResp)(nil), // 2: github.com.eonias189.calculationService.proto.RegisterResp
	(*ResultResp)(nil),   // 3: github.com.eonias189.calculationService.proto.ResultResp
	(*Task)(nil),         // 4: github.com.eonias189.calculationService.proto.Task
}
var file_proto_orchestrator_proto_depIdxs = []int32{
	1, // 0: github.com.eonias189.calculationService.proto.Orchestrator.Register:input_type -> github.com.eonias189.calculationService.proto.RegisterReq
	3, // 1: github.com.eonias189.calculationService.proto.Orchestrator.Connect:input_type -> github.com.eonias189.calculationService.proto.ResultResp
	4, // 2: github.com.eonias189.calculationService.proto.Orchestrator.Distribute:input_type -> github.com.eonias189.calculationService.proto.Task
	2, // 3: github.com.eonias189.calculationService.proto.Orchestrator.Register:output_type -> github.com.eonias189.calculationService.proto.RegisterResp
	4, // 4: github.com.eonias189.calculationService.proto.Orchestrator.Connect:output_type -> github.com.eonias189.calculationService.proto.Task
	0, // 5: github.com.eonias189.calculationService.proto.Orchestrator.Distribute:output_type -> github.com.eonias189.calculationService.proto.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_orchestrator_proto_init() }
func file_proto_orchestrator_proto_init() {
	if File_proto_orchestrator_proto != nil {
		return
	}
	file_proto_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_orchestrator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_orchestrator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_proto_orchestrator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResp); i {
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
		file_proto_orchestrator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultResp); i {
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
			RawDescriptor: file_proto_orchestrator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_orchestrator_proto_goTypes,
		DependencyIndexes: file_proto_orchestrator_proto_depIdxs,
		MessageInfos:      file_proto_orchestrator_proto_msgTypes,
	}.Build()
	File_proto_orchestrator_proto = out.File
	file_proto_orchestrator_proto_rawDesc = nil
	file_proto_orchestrator_proto_goTypes = nil
	file_proto_orchestrator_proto_depIdxs = nil
}
