// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/models.proto

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

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Expression string    `protobuf:"bytes,2,opt,name=expression,proto3" json:"expression,omitempty"`
	Timeouts   *Timeouts `protobuf:"bytes,3,opt,name=timeouts,proto3" json:"timeouts,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_models_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

func (x *Task) GetTimeouts() *Timeouts {
	if x != nil {
		return x.Timeouts
	}
	return nil
}

type Timeouts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Add uint64 `protobuf:"varint,1,opt,name=add,proto3" json:"add,omitempty"`
	Sub uint64 `protobuf:"varint,2,opt,name=sub,proto3" json:"sub,omitempty"`
	Mul uint64 `protobuf:"varint,3,opt,name=mul,proto3" json:"mul,omitempty"`
	Div uint64 `protobuf:"varint,4,opt,name=div,proto3" json:"div,omitempty"`
}

func (x *Timeouts) Reset() {
	*x = Timeouts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timeouts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timeouts) ProtoMessage() {}

func (x *Timeouts) ProtoReflect() protoreflect.Message {
	mi := &file_proto_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timeouts.ProtoReflect.Descriptor instead.
func (*Timeouts) Descriptor() ([]byte, []int) {
	return file_proto_models_proto_rawDescGZIP(), []int{1}
}

func (x *Timeouts) GetAdd() uint64 {
	if x != nil {
		return x.Add
	}
	return 0
}

func (x *Timeouts) GetSub() uint64 {
	if x != nil {
		return x.Sub
	}
	return 0
}

func (x *Timeouts) GetMul() uint64 {
	if x != nil {
		return x.Mul
	}
	return 0
}

func (x *Timeouts) GetDiv() uint64 {
	if x != nil {
		return x.Div
	}
	return 0
}

var File_proto_models_proto protoreflect.FileDescriptor

var file_proto_models_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x65, 0x6f, 0x6e, 0x69, 0x61, 0x73, 0x31, 0x38, 0x39, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6f, 0x6e, 0x69,
	0x61, 0x73, 0x31, 0x38, 0x39, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x73, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x73, 0x22, 0x52, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x61, 0x64, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x75,
	0x62, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x75, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x6d, 0x75, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x64, 0x69, 0x76, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_models_proto_rawDescOnce sync.Once
	file_proto_models_proto_rawDescData = file_proto_models_proto_rawDesc
)

func file_proto_models_proto_rawDescGZIP() []byte {
	file_proto_models_proto_rawDescOnce.Do(func() {
		file_proto_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_models_proto_rawDescData)
	})
	return file_proto_models_proto_rawDescData
}

var file_proto_models_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_models_proto_goTypes = []interface{}{
	(*Task)(nil),     // 0: github.com.eonias189.calculationService.proto.Task
	(*Timeouts)(nil), // 1: github.com.eonias189.calculationService.proto.Timeouts
}
var file_proto_models_proto_depIdxs = []int32{
	1, // 0: github.com.eonias189.calculationService.proto.Task.timeouts:type_name -> github.com.eonias189.calculationService.proto.Timeouts
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_models_proto_init() }
func file_proto_models_proto_init() {
	if File_proto_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_proto_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timeouts); i {
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
			RawDescriptor: file_proto_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_models_proto_goTypes,
		DependencyIndexes: file_proto_models_proto_depIdxs,
		MessageInfos:      file_proto_models_proto_msgTypes,
	}.Build()
	File_proto_models_proto = out.File
	file_proto_models_proto_rawDesc = nil
	file_proto_models_proto_goTypes = nil
	file_proto_models_proto_depIdxs = nil
}
