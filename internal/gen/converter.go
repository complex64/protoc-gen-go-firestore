package gen

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (m *Message) genConverterMethods() {
	m.genToProtoMethod()
	m.genToFirestoreMethod()
}

func (m *Message) genToProtoMethod() {
	m.P(Comment(" ToProto converts this %s to its protobuf representation.", m.CustomObjectName()),
		"func (m *", m.CustomObjectName(), ") ToProto() (*", m.proto.GoIdent.GoName, ", error) {")
	m.P("x := new(", m.proto.GoIdent.GoName, ")")
	m.genModelToProtoFields()
	m.P("return x, nil")
	m.P("}") // func
	m.P()
}

func (m *Message) genModelToProtoFields() {
	for _, field := range m.fields {
		field.genConvertToProto()
	}
}

func (f *Field) genConvertToProto() {
	switch {
	case f.fieldType.IsEnum:
		f.genEnumToProto()
	case f.fieldType.isTimestamp():
		f.genConvertTimeToProto()
	case f.fieldType.IsPointer:
		f.genConvertPointerField()
	default:
		f.genConvertDefaultField()
	}
}

func (f *Field) genConvertPointerField() {
	f.P("x.", f.Name(), " = *m.", f.Name())
}

func (f *Field) genConvertDefaultField() {
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

func (m *Message) genToFirestoreMethod() {
	m.P(Comment(" ToFirestore returns the Firestore Custom Object for %s.", m.proto.GoIdent.GoName),
		"func (x *", m.proto.GoIdent.GoName, ") ToFirestore() (*", m.CustomObjectName(), ", error) {")
	m.P("m := new(", m.CustomObjectName(), ")")
	m.genConvertToFirestoreFields()
	m.P("return m, nil")
	m.P("}")
	m.P()
}

func (m *Message) genConvertToFirestoreFields() {
	for _, field := range m.fields {
		field.genConvertToFirestore()
	}
}

func (f *Field) genConvertToFirestore() {
	switch {
	case f.fieldType.IsEnum:
		f.genEnumToFirestore()
	case f.fieldType.isTimestamp():
		f.genConvertTimestampToFirestore()
	case f.fieldType.IsPointer:
		f.P("m.", f.Name(), " = *x.", f.Name())
	default:
		f.P("m.", f.Name(), " = x.", f.Name())
	}
}

func (f *Field) genEnumToFirestore() {
	f.P("m.", f.Name(), " = ", "int32(x.", f.Name(), ")")
}

func (f *Field) genConvertTimestampToFirestore() {
	tt := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Time",
		GoImportPath: "time",
	})
	if f.fieldType.IsList {
		f.P("{")
		f.P("l := len(x.", f.Name(), ")")
		f.P("m.", f.Name(), " = make([]", tt, ", l)")
		f.P("for i:=0; i < l; i++ {")
		f.P("if x.", f.Name(), "[i] != nil {")
		f.P("m.", f.Name(), "[i] = x.", f.Name(), "[i].AsTime()")
		f.P("}") // if
		f.P("}") // for
		f.P("}")
	} else {
		f.P("if t := x.", f.Name(), "; t != nil {")
		f.P("m.", f.Name(), " = t.AsTime()")
		f.P("}")
	}
}
