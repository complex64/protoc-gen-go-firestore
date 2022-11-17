package gen

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (m *Message) genToProtoMethod() {
	m.P(Comment(" ToProto converts this %s to its protobuf representation.", m.CustomObjectName()),
		"func (m *", m.CustomObjectName(), ") ToProto() (*", m.proto.GoIdent.GoName, ", error) {")
	m.P("x := new(", m.proto.GoIdent.GoName, ")")
	m.genToProtoFields()
	m.P("return x, nil")
	m.P("}") // func
	m.P()
}

func (m *Message) genToProtoFields() {
	for _, field := range m.fields {
		field.genConvertFieldToProto()
	}
}

func (f *Field) genConvertFieldToProto() {
	switch {
	case f.fieldType.IsEnum:
		f.genEnumToProto()
	case f.fieldType.isTimestamp():
		f.genConvertTimeToProto()
	case f.fieldType.IsPointer:
		f.genConvertPointerToProto()
	default:
		f.genConvertDefaultToProto()
	}
}

func (f *Field) genConvertPointerToProto() {
	if f.fieldType.IsCustom && !f.fieldType.IsExternal {
		f.P("if m.", f.Name(), " != nil {")
		f.P("if c, err := m.", f.Name(), ".ToProto(); err != nil {")
		f.P("return nil, err")
		f.P("} else {")
		f.P("x.", f.Name(), " = c")
		f.P("}")
		f.P("}")
	} else {
		f.P("x.", f.Name(), " = *m.", f.Name())
	}
}

func (f *Field) genConvertDefaultToProto() {
	f.P("x.", f.Name(), " = m.", f.Name())
}

func (f *Field) genConvertTimeToProto() {
	tsType := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Timestamp",
		GoImportPath: "google.golang.org/protobuf/types/known/timestamppb",
	})
	newTimestamp := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "New",
		GoImportPath: "google.golang.org/protobuf/types/known/timestamppb",
	})
	if f.fieldType.IsList {
		f.P("{")
		f.P("l := len(m.", f.Name(), ")")
		f.P("x.", f.Name(), " = make([]*", tsType, ", l)")
		f.P("for i:=0; i < l; i++ {")
		f.P("if m.", f.Name(), "[i] != (time.Time{}) {")
		f.P("x.", f.Name(), "[i] = ", newTimestamp, "(m.", f.Name(), "[i])")
		f.P("}") // if
		f.P("}") // for
		f.P("}")
	} else {
		f.P("if m.", f.Name(), " != (time.Time{}) {")
		f.P("x.", f.Name(), " = ", newTimestamp, "(m.", f.Name(), ")")
		f.P("}")
	}
}

func (f *Field) genEnumToProto() {
	typename := f.proto.Desc.Enum().Name()
	parent := f.proto.Desc.Enum().Parent()
	if x, ok := parent.(protoreflect.MessageDescriptor); ok {
		typename = x.Name() + "_" + typename
	}
	f.P("x.", f.Name(), " = ", typename, "(m.", f.Name(), ")")
}
