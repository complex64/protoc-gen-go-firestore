package gen

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func (m *Message) genToProtoMethod() {
	var (
		marshal   = m.file.out.QualifiedGoIdent(protogen.GoIdent{GoName: "Marshal", GoImportPath: "encoding/json"})
		unmarshal = m.file.out.QualifiedGoIdent(protogen.GoIdent{GoName: "Unmarshal", GoImportPath: "google.golang.org/protobuf/encoding/protojson"})
	)
	m.P("func FirestoreMapTo", m.proto.GoIdent.GoName, "(m map[string]any) (*", m.proto.GoIdent.GoName, ", error) {")
	{

		m.P("bs, err := ", marshal, "(m)")
		m.P("if err != nil {")
		m.P("return nil, err")
		m.P("}")

		m.P("msg := new(", m.proto.GoIdent.GoName, ")")
		m.P("if err := ", unmarshal, "(bs, msg); err != nil {")
		m.P("return nil, err")
		m.P("}")

		m.P("return msg, nil")
	}
	m.P("}")
	m.P()
}
