// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: customobject/customobject.proto

package customobject

import (
	_ "github.com/complex64/protoc-gen-go-firestore/firestorepb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type City struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField            string                   `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	RepeatedStringField    []string                 `protobuf:"bytes,2,rep,name=repeated_string_field,json=repeatedStringField,proto3" json:"repeated_string_field,omitempty"`
	BoolField              bool                     `protobuf:"varint,3,opt,name=bool_field,json=boolField,proto3" json:"bool_field,omitempty"`
	RepeatedBoolField      []bool                   `protobuf:"varint,4,rep,packed,name=repeated_bool_field,json=repeatedBoolField,proto3" json:"repeated_bool_field,omitempty"`
	BytesField             []byte                   `protobuf:"bytes,5,opt,name=bytes_field,json=bytesField,proto3" json:"bytes_field,omitempty"`
	RepeatedBytesField     [][]byte                 `protobuf:"bytes,6,rep,name=repeated_bytes_field,json=repeatedBytesField,proto3" json:"repeated_bytes_field,omitempty"`
	Int32Field             int32                    `protobuf:"varint,7,opt,name=int32_field,json=int32Field,proto3" json:"int32_field,omitempty"`
	RepeatedInt32Field     []int32                  `protobuf:"varint,8,rep,packed,name=repeated_int32_field,json=repeatedInt32Field,proto3" json:"repeated_int32_field,omitempty"`
	Int64Field             int64                    `protobuf:"varint,9,opt,name=int64_field,json=int64Field,proto3" json:"int64_field,omitempty"`
	RepeatedInt64Field     []int64                  `protobuf:"varint,10,rep,packed,name=repeated_int64_field,json=repeatedInt64Field,proto3" json:"repeated_int64_field,omitempty"`
	Uint32Field            uint32                   `protobuf:"varint,11,opt,name=uint32_field,json=uint32Field,proto3" json:"uint32_field,omitempty"`
	RepeatedUint32Field    []uint32                 `protobuf:"varint,12,rep,packed,name=repeated_uint32_field,json=repeatedUint32Field,proto3" json:"repeated_uint32_field,omitempty"`
	Uint64Field            uint64                   `protobuf:"varint,13,opt,name=uint64_field,json=uint64Field,proto3" json:"uint64_field,omitempty"`
	RepeatedUint64Field    []uint64                 `protobuf:"varint,14,rep,packed,name=repeated_uint64_field,json=repeatedUint64Field,proto3" json:"repeated_uint64_field,omitempty"`
	Sint32Field            int32                    `protobuf:"zigzag32,15,opt,name=sint32_field,json=sint32Field,proto3" json:"sint32_field,omitempty"`
	RepeatedSint32Field    []int32                  `protobuf:"zigzag32,16,rep,packed,name=repeated_sint32_field,json=repeatedSint32Field,proto3" json:"repeated_sint32_field,omitempty"`
	Sint64Field            int64                    `protobuf:"zigzag64,17,opt,name=sint64_field,json=sint64Field,proto3" json:"sint64_field,omitempty"`
	RepeatedSint64Field    []int64                  `protobuf:"zigzag64,18,rep,packed,name=repeated_sint64_field,json=repeatedSint64Field,proto3" json:"repeated_sint64_field,omitempty"`
	FloatField             float32                  `protobuf:"fixed32,19,opt,name=float_field,json=floatField,proto3" json:"float_field,omitempty"`
	RepeatedFloatField     []float32                `protobuf:"fixed32,20,rep,packed,name=repeated_float_field,json=repeatedFloatField,proto3" json:"repeated_float_field,omitempty"`
	DoubleField            float64                  `protobuf:"fixed64,21,opt,name=double_field,json=doubleField,proto3" json:"double_field,omitempty"`
	RepeatedDoubleField    []float64                `protobuf:"fixed64,22,rep,packed,name=repeated_double_field,json=repeatedDoubleField,proto3" json:"repeated_double_field,omitempty"`
	TimestampField         *timestamppb.Timestamp   `protobuf:"bytes,23,opt,name=timestamp_field,json=timestampField,proto3" json:"timestamp_field,omitempty"`
	RepeatedTimestampField []*timestamppb.Timestamp `protobuf:"bytes,24,rep,name=repeated_timestamp_field,json=repeatedTimestampField,proto3" json:"repeated_timestamp_field,omitempty"`
	NamedStringField       string                   `protobuf:"bytes,25,opt,name=named_string_field,json=namedStringField,proto3" json:"named_string_field,omitempty"`
	IgnoredStringField     string                   `protobuf:"bytes,26,opt,name=ignored_string_field,json=ignoredStringField,proto3" json:"ignored_string_field,omitempty"`
	ServerTimestampField   *timestamppb.Timestamp   `protobuf:"bytes,27,opt,name=server_timestamp_field,json=serverTimestampField,proto3" json:"server_timestamp_field,omitempty"`
}

func (x *City) Reset() {
	*x = City{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customobject_customobject_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *City) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*City) ProtoMessage() {}

func (x *City) ProtoReflect() protoreflect.Message {
	mi := &file_customobject_customobject_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use City.ProtoReflect.Descriptor instead.
func (*City) Descriptor() ([]byte, []int) {
	return file_customobject_customobject_proto_rawDescGZIP(), []int{0}
}

func (x *City) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *City) GetRepeatedStringField() []string {
	if x != nil {
		return x.RepeatedStringField
	}
	return nil
}

func (x *City) GetBoolField() bool {
	if x != nil {
		return x.BoolField
	}
	return false
}

func (x *City) GetRepeatedBoolField() []bool {
	if x != nil {
		return x.RepeatedBoolField
	}
	return nil
}

func (x *City) GetBytesField() []byte {
	if x != nil {
		return x.BytesField
	}
	return nil
}

func (x *City) GetRepeatedBytesField() [][]byte {
	if x != nil {
		return x.RepeatedBytesField
	}
	return nil
}

func (x *City) GetInt32Field() int32 {
	if x != nil {
		return x.Int32Field
	}
	return 0
}

func (x *City) GetRepeatedInt32Field() []int32 {
	if x != nil {
		return x.RepeatedInt32Field
	}
	return nil
}

func (x *City) GetInt64Field() int64 {
	if x != nil {
		return x.Int64Field
	}
	return 0
}

func (x *City) GetRepeatedInt64Field() []int64 {
	if x != nil {
		return x.RepeatedInt64Field
	}
	return nil
}

func (x *City) GetUint32Field() uint32 {
	if x != nil {
		return x.Uint32Field
	}
	return 0
}

func (x *City) GetRepeatedUint32Field() []uint32 {
	if x != nil {
		return x.RepeatedUint32Field
	}
	return nil
}

func (x *City) GetUint64Field() uint64 {
	if x != nil {
		return x.Uint64Field
	}
	return 0
}

func (x *City) GetRepeatedUint64Field() []uint64 {
	if x != nil {
		return x.RepeatedUint64Field
	}
	return nil
}

func (x *City) GetSint32Field() int32 {
	if x != nil {
		return x.Sint32Field
	}
	return 0
}

func (x *City) GetRepeatedSint32Field() []int32 {
	if x != nil {
		return x.RepeatedSint32Field
	}
	return nil
}

func (x *City) GetSint64Field() int64 {
	if x != nil {
		return x.Sint64Field
	}
	return 0
}

func (x *City) GetRepeatedSint64Field() []int64 {
	if x != nil {
		return x.RepeatedSint64Field
	}
	return nil
}

func (x *City) GetFloatField() float32 {
	if x != nil {
		return x.FloatField
	}
	return 0
}

func (x *City) GetRepeatedFloatField() []float32 {
	if x != nil {
		return x.RepeatedFloatField
	}
	return nil
}

func (x *City) GetDoubleField() float64 {
	if x != nil {
		return x.DoubleField
	}
	return 0
}

func (x *City) GetRepeatedDoubleField() []float64 {
	if x != nil {
		return x.RepeatedDoubleField
	}
	return nil
}

func (x *City) GetTimestampField() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampField
	}
	return nil
}

func (x *City) GetRepeatedTimestampField() []*timestamppb.Timestamp {
	if x != nil {
		return x.RepeatedTimestampField
	}
	return nil
}

func (x *City) GetNamedStringField() string {
	if x != nil {
		return x.NamedStringField
	}
	return ""
}

func (x *City) GetIgnoredStringField() string {
	if x != nil {
		return x.IgnoredStringField
	}
	return ""
}

func (x *City) GetServerTimestampField() *timestamppb.Timestamp {
	if x != nil {
		return x.ServerTimestampField
	}
	return nil
}

var File_customobject_customobject_proto protoreflect.FileDescriptor

var file_customobject_customobject_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a,
	0x17, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x0a, 0x0a, 0x04, 0x43, 0x69,
	0x74, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f,
	0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62,
	0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x08, 0x52, 0x11, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x12, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x30, 0x0a, 0x14,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x08, 0x20, 0x03, 0x28, 0x05, 0x52, 0x12, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x30, 0x0a, 0x14, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x03, 0x52, 0x12, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x13, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x55, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x04, 0x52, 0x13, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x10, 0x20, 0x03, 0x28,
	0x11, 0x52, 0x13, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0b, 0x73, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x12, 0x20, 0x03, 0x28, 0x12, 0x52, 0x13, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x30,
	0x0a, 0x14, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x14, 0x20, 0x03, 0x28, 0x02, 0x52, 0x12, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x16, 0x20, 0x03,
	0x28, 0x01, 0x52, 0x13, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x43, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x54, 0x0a, 0x18,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x16, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x3f, 0x0a, 0x12, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11,
	0xd2, 0xc1, 0x18, 0x0d, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x10, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x38, 0x0a, 0x14, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xd2, 0xc1, 0x18, 0x02, 0x10, 0x01, 0x52, 0x12, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x58, 0x0a,
	0x16, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xd2, 0xc1, 0x18, 0x02, 0x18,
	0x01, 0x52, 0x14, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x06, 0xca, 0xc1, 0x18, 0x02, 0x08, 0x01, 0x42,
	0x5c, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x36, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2d, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customobject_customobject_proto_rawDescOnce sync.Once
	file_customobject_customobject_proto_rawDescData = file_customobject_customobject_proto_rawDesc
)

func file_customobject_customobject_proto_rawDescGZIP() []byte {
	file_customobject_customobject_proto_rawDescOnce.Do(func() {
		file_customobject_customobject_proto_rawDescData = protoimpl.X.CompressGZIP(file_customobject_customobject_proto_rawDescData)
	})
	return file_customobject_customobject_proto_rawDescData
}

var file_customobject_customobject_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_customobject_customobject_proto_goTypes = []interface{}{
	(*City)(nil),                  // 0: customobject.City
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_customobject_customobject_proto_depIdxs = []int32{
	1, // 0: customobject.City.timestamp_field:type_name -> google.protobuf.Timestamp
	1, // 1: customobject.City.repeated_timestamp_field:type_name -> google.protobuf.Timestamp
	1, // 2: customobject.City.server_timestamp_field:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_customobject_customobject_proto_init() }
func file_customobject_customobject_proto_init() {
	if File_customobject_customobject_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customobject_customobject_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*City); i {
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
			RawDescriptor: file_customobject_customobject_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customobject_customobject_proto_goTypes,
		DependencyIndexes: file_customobject_customobject_proto_depIdxs,
		MessageInfos:      file_customobject_customobject_proto_msgTypes,
	}.Build()
	File_customobject_customobject_proto = out.File
	file_customobject_customobject_proto_rawDesc = nil
	file_customobject_customobject_proto_goTypes = nil
	file_customobject_customobject_proto_depIdxs = nil
}
