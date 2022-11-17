package gen

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FieldType struct {
	field     *Field
	Go        protogen.GoIdent
	Firestore protogen.GoIdent

	IsPointer  bool // TODO: unused?
	IsEnum     bool
	IsCustom   bool
	IsExternal bool
	IsList     bool
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
	if t.field.proto.Desc.IsMap() {
		panic(fmt.Sprintf("TODO: Map type: %+v", t.Go))
	}

	t.IsList = t.field.proto.Desc.IsList()

	switch t.field.proto.Desc.Kind() {
	case protoreflect.BoolKind:
		t.Firestore.GoName = "bool"
	case protoreflect.EnumKind:
		t.Firestore.GoName = "int32"
		t.IsEnum = true
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		t.Firestore.GoName = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		t.Firestore.GoName = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		t.Firestore.GoName = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		t.Firestore.GoName = "int64"
	case protoreflect.FloatKind:
		t.Firestore.GoName = "float32"
	case protoreflect.DoubleKind:
		t.Firestore.GoName = "float64"
	case protoreflect.StringKind:
		t.Firestore.GoName = "string"
	case protoreflect.BytesKind:
		t.Firestore.GoName = "[]byte"

	case protoreflect.MessageKind, protoreflect.GroupKind:
		if t.isTimestamp() {
			t.Firestore.GoName += "Time"
			t.Firestore.GoImportPath = "time"
			return nil
		}
		nested := t.field.proto.Message
		t.Go = nested.GoIdent

		filePkg := t.field.msg.file.proto.GoImportPath
		fieldPkg := nested.GoIdent.GoImportPath
		t.IsExternal = filePkg != fieldPkg
	}

	if unmapped := t.Firestore.GoName == ""; unmapped {
		t.IsCustom = true
	}

	switch {
	case t.IsCustom && t.IsExternal:
		panic(fmt.Sprintf("TODO: External custom types: %+v", t.Go))

	case t.IsCustom && !t.IsExternal:
		panic(fmt.Sprintf("TODO: Internal custom types: %+v", t.Go))
	}

	return nil
}

func (t *FieldType) String() string {
	if t.Firestore.GoImportPath != "" {
		id := t.field.msg.file.out.QualifiedGoIdent(t.Firestore)
		if t.IsList {
			id = "[]" + id
		}
		if t.IsPointer {
			id = "*" + id
		}
		return id
	}
	if t.IsPointer {
		return "*" + t.Firestore.GoName
	}
	id := t.Firestore.GoName
	if t.IsList {
		id = "[]" + id
	}
	return id
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
