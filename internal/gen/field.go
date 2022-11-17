package gen

import (
	"fmt"
	"strconv"
	"strings"

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

	opts      *firestorepb.FieldOptions
	fieldType *FieldType
}

func (f *Field) init() error {
	f.initOpts()
	if err := f.initTypes(); err != nil {
		return err
	}
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

func (f *Field) initTypes() error {
	fieldType, err := NewFieldType(f)
	if err != nil {
		return err
	}
	f.fieldType = fieldType
	return nil
}

func (f *Field) Gen() {
	name := f.proto.GoName
	f.Annotate(f.msg.CustomObjectName()+"."+name, f.proto.Location)
	f.P(name, " ", f.fieldType.String(), f.tags())
}

func (f *Field) tags() string {
	if f.opts == nil {
		return ""
	}
	vals := f.tagVals()
	if len(vals) == 0 {
		return ""
	}

	joined := strings.Join(vals, ";")
	quoted := strconv.Quote(joined)
	escaped := strings.Replace(quoted, "`", `\x60`, -1)
	return fmt.Sprintf("`firestore:%s`", escaped)
}

func (f *Field) tagVals() (values []string) {
	if f.opts.Ignore {
		return []string{"-"}
	}
	if f.opts.ServerTimestamp {
		return []string{"serverTimestamp"}
	}
	values = append(values, f.FirestoreFieldName()+",omitempty")
	return
}

func (f *Field) FirestoreFieldName() string {
	if f.opts.Name != "" {
		return f.opts.Name
	}
	return f.proto.Desc.JSONName()
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
