package gen

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func (m *Message) genConvertToFirestoreFields() {
	for _, field := range m.fields {
		field.genConvertFieldToFirestore()
	}
}

func (f *Field) genConvertFieldToFirestore() {
	switch {
	case f.fieldType.IsEnum:
		f.genEnumToFirestore()
	case f.fieldType.isTimestamp():
		f.genConvertTimestampToFirestore()
	case f.fieldType.IsPointer:
		f.genConvertPointerToFirestore()
	default:
		f.genConvertDefaultToFirebase()
	}
}

func (f *Field) genConvertDefaultToFirebase() {
	f.P("m.", f.Name(), " = x.", f.Name())
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

func (f *Field) genConvertPointerToFirestore() {
	if f.fieldType.IsCustom && !f.fieldType.IsExternal {
		f.P("if c, err := x.", f.Name(), ".ToFirestore(); err != nil {")
		f.P("return nil, err")
		f.P("} else {")
		f.P("m.", f.Name(), " = c")
		f.P("}")
	} else {
		f.P("m.", f.Name(), " = *x.", f.Name())
	}
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
