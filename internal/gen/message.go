package gen

import (
	"fmt"

	"github.com/complex64/protoc-gen-go-firestore/firestorepb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func NewMessage(file *File, proto *protogen.Message) (*Message, error) {
	msg := &Message{
		file:  file,
		proto: proto,
	}
	if err := msg.init(); err != nil {
		return nil, err
	}
	return msg, nil
}

type Message struct {
	file  *File
	proto *protogen.Message

	opts   *firestorepb.MessageOptions
	fields []*Field
}

func (m *Message) init() error {
	m.initOpts()
	if err := m.initFields(); err != nil {
		return err
	}
	return nil
}

func (m *Message) initFields() error {
	if !m.enabled() {
		return nil
	}

	for _, f := range m.proto.Fields {
		if err := m.initField(f); err != nil {
			return err
		}
	}
	return nil
}

func (m *Message) initField(proto *protogen.Field) error {
	field, err := NewField(m, proto)
	if err != nil {
		return err
	}
	m.fields = append(m.fields, field)
	return nil
}

// initOpts reads the protoc-gen-go-firestore options set for this message.
// Example: message MyMessage { option (gorm.message).model = true; }
func (m *Message) initOpts() {
	descOpts := m.proto.Desc.Options()
	opts, ok := proto.GetExtension(descOpts, firestorepb.E_Message).(*firestorepb.MessageOptions)
	if ok && opts != nil {
		m.opts = opts
	} else {
		m.opts = &firestorepb.MessageOptions{}
	}
}

// Gen generates GORM models and supporting APIs.
func (m *Message) Gen() {
	if !m.enabled() {
		return
	}
	m.genCollectionNameConstant()
	m.genCustomObjectStructType()
	m.genConverterMethods()
}

func (m *Message) genCollectionNameConstant() {
	m.Annotate(m.CollectionConstantName(), m.proto.Location)
	m.P(m.leadingConstComment(), "const ", m.CollectionConstantName(), " = \"TODO\"")
	m.P()
}

func (m *Message) leadingConstComment() protogen.Comments {
	return appendDeprecationNotice(
		Comment(" %s is the Firestore collection name for documents of type %s.%s.",
			m.CollectionConstantName(),
			m.file.proto.GoPackageName,
			m.proto.GoIdent.GoName),
		m.deprecated(),
	)
}

func (m *Message) genCustomObjectStructType() {
	m.Annotate(m.CustomObjectName(), m.proto.Location) // Message/document type declaration.
	m.P(m.leadingStructComment(), "type ", m.CustomObjectName(), " struct {")
	m.genFields()
	m.P("}")
	m.P()
}

func (m *Message) leadingStructComment() protogen.Comments {
	return appendDeprecationNotice(
		Comment(" %s is the Firestore Custom Object for %s.%s.",
			m.CustomObjectName(),
			m.file.proto.GoPackageName,
			m.proto.GoIdent.GoName),
		m.deprecated(),
	)
}

func (m *Message) deprecated() bool {
	return m.proto.Desc.Options().(*descriptorpb.MessageOptions).GetDeprecated()
}

func (m *Message) genFields() {
	for _, field := range m.fields {
		field.Gen()
	}
}

func (m *Message) ProtoName() string {
	return m.proto.GoIdent.GoName
}

func (m *Message) CustomObjectName() string {
	return fmt.Sprintf("Firestore%s", m.ProtoName())
}

func (m *Message) CollectionConstantName() string {
	return fmt.Sprintf("Firestore%sCollection", m.ProtoName())
}

func (m *Message) enabled() bool {
	return m.opts.Enabled || m.file.Enabled()
}

func (m *Message) Annotate(symbol string, loc protogen.Location) {
	m.file.Annotate(symbol, loc)
}

func (m *Message) P(v ...interface{}) {
	m.file.P(v...)
}
