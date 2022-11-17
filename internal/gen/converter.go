package gen

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (m *Message) genConverterMethods() {
	// m.genToProtoMethod()
	// m.genToFirestoreMethod()
}

func (m *Message) genToProtoMethod() {
	m.P(Comment(" ToProto converts a %s to its protobuf representation.", m.CustomObjectName()),
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
		f.P("x.", f.Name(), " = *m.", f.Name())
	default:
		f.P("x.", f.Name(), " = m.", f.Name())
	}
}

func (f *Field) genConvertTimeToProto() {
	newTimestamp := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "New",
		GoImportPath: "google.golang.org/protobuf/types/known/timestamppb",
	})
	f.P("if m.", f.Name(), " != (time.Time{}) {")
	f.P("x.", f.Name(), " = ", newTimestamp, "(m.", f.Name(), ")")
	f.P("}")
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
	f.P("if t := x.", f.Name(), "; t != nil {")
	f.P("m.", f.Name(), " = t.AsTime()")
	f.P("}")
}
