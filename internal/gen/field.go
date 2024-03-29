package gen

import (
	"github.com/complex64/protoc-gen-go-firestore/firestorepb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func NewField(msg *Message, proto *protogen.Field) (*Field, error) {
	field := &Field{
		msg:   msg,
		proto: proto,
	}
	if err := field.init(); err != nil {
		return nil, err
	}
	return field, nil
}

type Field struct {
	msg   *Message
	proto *protogen.Field
	opts  *firestorepb.FieldOptions
}

func (f *Field) init() error {
	f.initOpts()
	return nil
}

func (f *Field) initOpts() {
	descOpts := f.proto.Desc.Options()
	opts, ok := proto.GetExtension(descOpts, firestorepb.E_Field).(*firestorepb.FieldOptions)
	if ok && opts != nil {
		f.opts = opts
	} else {
		f.opts = &firestorepb.FieldOptions{}
	}
}

func (f *Field) Name() string {
	return f.proto.GoName
}

func (f *Field) Annotate(symbol string, loc protogen.Location) {
	f.msg.Annotate(symbol, loc)
}

func (f *Field) P(v ...interface{}) {
	f.msg.P(v...)
}
