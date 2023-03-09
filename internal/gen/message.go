package gen

import (
	"fmt"
	"strings"

	"github.com/complex64/protoc-gen-go-firestore/firestorepb"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

	idField *Field
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

	if field.opts.Id {
		if m.idField != nil {
			panic("duplicate id field")
		}
		m.idField = field
	}

	m.fields = append(m.fields, field)
	return nil
}

// initOpts reads the protoc-gen-go-firestore options set for this message.
// Example: message MyMessage { option (firestore.message).model = true; }
func (m *Message) initOpts() {
	descOpts := m.proto.Desc.Options()
	opts, ok := proto.GetExtension(descOpts, firestorepb.E_Message).(*firestorepb.MessageOptions)
	if ok && opts != nil {
		m.opts = opts
	} else {
		m.opts = &firestorepb.MessageOptions{}
	}
}

func (m *Message) Gen() {
	log.Trace().
		Str("file", m.file.proto.Desc.Path()).
		Str("msg", m.ProtoName()).
		Msg("(m *Message).Gen()")

	if !m.enabled() {
		return
	}

	m.genCollectionNameConstant()
	m.genConverterMethods()
}

func (m *Message) genCollectionNameConstant() {
	if m.opts.Collection == "" {
		return
	}

	m.Annotate(m.CollectionConstantName(), m.proto.Location)
	m.P(m.leadingConstComment(),
		"const ", m.CollectionConstantName(), " = \"", m.CollectionName(), "\"")
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

func (m *Message) deprecated() bool {
	return m.proto.Desc.Options().(*descriptorpb.MessageOptions).GetDeprecated()
}

func (m *Message) ProtoName() string {
	return m.proto.GoIdent.GoName
}

func (m *Message) CollectionConstantName() string {
	return fmt.Sprintf("FirestoreCollection%s", m.CollectionNameTitle())
}

func (m *Message) enabled() bool {
	_, nested := m.file.nested[m.ProtoName()]
	return m.opts.Enabled || m.opts.Collection != "" || m.file.Enabled() || nested
}

func (m *Message) Annotate(symbol string, loc protogen.Location) {
	m.file.Annotate(symbol, loc)
}

func (m *Message) P(v ...interface{}) {
	m.file.P(v...)
}

var title cases.Caser

func init() {
	title = cases.Title(language.English)
}

func (m *Message) CollectionName() string {
	split := strings.Split(m.opts.Collection, "/")
	return split[len(split)-2]
}

func (m *Message) CollectionNameTitle() string {
	return title.String(m.CollectionName())
}
