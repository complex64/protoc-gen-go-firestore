package gen

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func (m *Message) genToFirestoreMethod() {
	var (
		marshal   = m.file.out.QualifiedGoIdent(protogen.GoIdent{GoName: "Marshal", GoImportPath: "google.golang.org/protobuf/encoding/protojson"})
		unmarshal = m.file.out.QualifiedGoIdent(protogen.GoIdent{GoName: "Unmarshal", GoImportPath: "encoding/json"})
	)
	m.P("func (x *", m.proto.GoIdent.GoName, ") ToFirestoreMap() (map[string]any, error) {")
	{

		m.P("bs, err := ", marshal, "(x)")
		m.P("if err != nil {")
		m.P("return nil, err")
		m.P("}")

		m.P("m := map[string]any{}")
		m.P("if err := ", unmarshal, "(bs, &m); err != nil {")
		m.P("return nil, err")
		m.P("}")

		m.P("return m, nil")
	}
	m.P("}")
	m.P()
}
