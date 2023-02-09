package gen

func (m *Message) genUtilityMethods() {
	// m.genWithinCollectionMethod()
	// m.genToProtoMethod()
	// m.genToFirestoreMethod()
}

// func (m *Message) genWithinCollectionMethod() {
// 	// collT := m.file.out.QualifiedGoIdent(protogen.GoIdent{
// 	// 	GoName:       "CollectionRef",
// 	// 	GoImportPath: "cloud.google.com/go/firestore",
// 	// })
// 	//
// 	// m.P(
// 	// 	"func (m *", m.ProtoName(), ") Within(", collT, ") (*", m.proto.GoIdent.GoName, ", error) {")
// 	// m.P("x := new(", m.proto.GoIdent.GoName, ")")
// 	// m.genToProtoFields()
// 	// m.P("return x, nil")
// 	// m.P("}") // func
// 	// m.P()
// }
