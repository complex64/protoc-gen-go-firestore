package gen

func (m *Message) genConverterMethods() {
	m.genToProtoMethod()
	m.genToFirestoreMethod()
}
