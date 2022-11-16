package gen

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FieldType struct {
	field *Field
	Go    protogen.GoIdent

	Firestore protogen.GoIdent
	Pointer   bool
	Enum      bool
	Custom    bool
	External  bool
}

func NewFieldType(field *Field) (*FieldType, error) {
	types := &FieldType{
		field: field,
		Go:    field.proto.GoIdent,
	}
	if err := types.init(); err != nil {
		return nil, err
	}
	return types, nil
}

func (t *FieldType) init() error {
	// if t.JSON {
	// 	t.Firestore.GoName = "[]byte"
	// 	t.Pointer = false
	// 	return nil
	// }

	switch t.field.proto.Desc.Kind() {
	case protoreflect.BoolKind:
		t.Firestore.GoName = "bool"
	case protoreflect.EnumKind:
		t.Firestore.GoName = "int32"
		t.Enum = true
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		t.Firestore.GoName = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		t.Firestore.GoName = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		t.Firestore.GoName = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		t.Firestore.GoName = "uint64"
	case protoreflect.FloatKind:
		t.Firestore.GoName = "float32"
	case protoreflect.DoubleKind:
		t.Firestore.GoName = "float64"
	case protoreflect.StringKind:
		t.Firestore.GoName = "string"
	case protoreflect.BytesKind:
		t.Firestore.GoName = "[]byte"
		t.Pointer = false

	case protoreflect.MessageKind, protoreflect.GroupKind:
		if t.isTimestamp() {
			t.Firestore.GoName = "Time"
			t.Firestore.GoImportPath = "time"
			return nil
		}
		nested := t.field.proto.Message
		t.Go = nested.GoIdent

		filePkg := t.field.msg.file.proto.GoImportPath
		fieldPkg := nested.GoIdent.GoImportPath
		t.External = filePkg != fieldPkg
	}

	if unmapped := t.Firestore.GoName == ""; unmapped {
		t.Custom = true
	}

	switch {
	case t.Custom && t.External:
		panic(fmt.Sprintf("TODO: External custom types: %+v", t.Go))
		// t.Firestore.GoName = t.alias()

	case t.Custom && !t.External:
		panic(fmt.Sprintf("TODO: Internal custom types: %+v", t.Go))
		// t.Firestore.GoName = t.Go.GoName
	}
	return nil
}

func (t *FieldType) String() string {
	if t.Firestore.GoImportPath != "" {
		id := t.field.msg.file.out.QualifiedGoIdent(t.Firestore)
		if t.Pointer {
			id = "*" + id
		}
		return id
	}
	if t.Pointer {
		return "*" + t.Firestore.GoName
	}
	return t.Firestore.GoName
}

func (t *FieldType) isTimestamp() bool {
	if t.field.proto.Message == nil {
		return false
	}
	var (
		path = t.field.proto.Message.GoIdent.GoImportPath
		name = t.field.proto.Message.GoIdent.GoName
	)
	return path == "google.golang.org/protobuf/types/known/timestamppb" &&
		name == "Timestamp"
}
