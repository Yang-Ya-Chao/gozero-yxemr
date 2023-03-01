// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: reg.proto

package reg

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

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ilx   int64    `protobuf:"varint,1,opt,name=Ilx,proto3" json:"Ilx,omitempty"`
	Ibrlx int64    `protobuf:"varint,2,opt,name=Ibrlx,proto3" json:"Ibrlx,omitempty"`
	Cbrh  string   `protobuf:"bytes,3,opt,name=Cbrh,proto3" json:"Cbrh,omitempty"`
	Csqdh string   `protobuf:"bytes,4,opt,name=Csqdh,proto3" json:"Csqdh,omitempty"`
	Cztbm []string `protobuf:"bytes,5,rep,name=Cztbm,proto3" json:"Cztbm,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_reg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_reg_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetIlx() int64 {
	if x != nil {
		return x.Ilx
	}
	return 0
}

func (x *Req) GetIbrlx() int64 {
	if x != nil {
		return x.Ibrlx
	}
	return 0
}

func (x *Req) GetCbrh() string {
	if x != nil {
		return x.Cbrh
	}
	return ""
}

func (x *Req) GetCsqdh() string {
	if x != nil {
		return x.Csqdh
	}
	return ""
}

func (x *Req) GetCztbm() []string {
	if x != nil {
		return x.Cztbm
	}
	return nil
}

type Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Resp) Reset() {
	*x = Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resp) ProtoMessage() {}

func (x *Resp) ProtoReflect() protoreflect.Message {
	mi := &file_reg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resp.ProtoReflect.Descriptor instead.
func (*Resp) Descriptor() ([]byte, []int) {
	return file_reg_proto_rawDescGZIP(), []int{1}
}

func (x *Resp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Resp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Resp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_reg_proto protoreflect.FileDescriptor

var file_reg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x65, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x65, 0x67,
	0x22, 0x6d, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x6c, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x49, 0x6c, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x62, 0x72,
	0x6c, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x49, 0x62, 0x72, 0x6c, 0x78, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x62, 0x72, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43,
	0x62, 0x72, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x73, 0x71, 0x64, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x43, 0x73, 0x71, 0x64, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x7a, 0x74,
	0x62, 0x6d, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x43, 0x7a, 0x74, 0x62, 0x6d, 0x22,
	0x40, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x32, 0x22, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x02, 0x44, 0x6f,
	0x12, 0x08, 0x2e, 0x72, 0x65, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x72, 0x65, 0x67,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x65, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reg_proto_rawDescOnce sync.Once
	file_reg_proto_rawDescData = file_reg_proto_rawDesc
)

func file_reg_proto_rawDescGZIP() []byte {
	file_reg_proto_rawDescOnce.Do(func() {
		file_reg_proto_rawDescData = protoimpl.X.CompressGZIP(file_reg_proto_rawDescData)
	})
	return file_reg_proto_rawDescData
}

var file_reg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_reg_proto_goTypes = []interface{}{
	(*Req)(nil),  // 0: reg.Req
	(*Resp)(nil), // 1: reg.Resp
}
var file_reg_proto_depIdxs = []int32{
	0, // 0: reg.reger.Do:input_type -> reg.Req
	1, // 1: reg.reger.Do:output_type -> reg.Resp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reg_proto_init() }
func file_reg_proto_init() {
	if File_reg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
		file_reg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resp); i {
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
			RawDescriptor: file_reg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reg_proto_goTypes,
		DependencyIndexes: file_reg_proto_depIdxs,
		MessageInfos:      file_reg_proto_msgTypes,
	}.Build()
	File_reg_proto = out.File
	file_reg_proto_rawDesc = nil
	file_reg_proto_goTypes = nil
	file_reg_proto_depIdxs = nil
}
