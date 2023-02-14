package gen

import (
	"fmt"
	"path"

	"github.com/complex64/protoc-gen-go-firestore/internal/gen/tree"
	"github.com/complex64/protoc-gen-go-firestore/internal/version"
	"google.golang.org/protobuf/compiler/protogen"
)

type Packages struct {
	plugin *protogen.Plugin
	list   map[string]*Package
}

func NewPackages(plugin *protogen.Plugin) *Packages {
	return &Packages{
		plugin: plugin,
		list:   make(map[string]*Package),
	}
}

func (p *Packages) Collect(file *File) {
	pkgName := string(file.proto.GoPackageName)
	if _, ok := p.list[pkgName]; !ok {
		p.list[pkgName] = &Package{
			plugin:     p.plugin,
			name:       pkgName,
			importPath: file.proto.GoImportPath,
			tree:       new(tree.Tree[*Message]),
		}
		p.list[pkgName].initOut(file)
	}
	pkg := p.list[pkgName]

	for _, msg := range file.msgs {
		if msg.opts.Collection == "" {
			continue
		}
		if err := pkg.tree.Add(msg.opts.Collection, msg); err != nil {
			panic(err)
		}
	}
}

func (p *Packages) Gen() error {
	for _, pkg := range p.list {
		pkg.Gen()
		_ = pkg
	}
	return nil
}

type Package struct {
	plugin     *protogen.Plugin
	name       string
	importPath protogen.GoImportPath
	out        *protogen.GeneratedFile
	tree       *tree.Tree[*Message]
}

func (p *Package) initOut(file *File) {
	base := path.Dir(file.proto.GeneratedFilenamePrefix)
	// Use extension .pb.go, like `protoc-gen-go` does.
	filepath := path.Join(base, "pkg_firebase.pb.go")
	p.out = p.plugin.NewGeneratedFile(filepath, p.importPath)
}

func (p *Package) Gen() {
	p.genHeader()
	p.genPackage()
	p.genFirestoreTypeAndMethod()
	p.genCollectionChainMethods()
}

func (p *Package) genFirestoreTypeAndMethod() {
	clientType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Client",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"type Firestore struct {")
	p.P("client *", clientType)
	p.P("}")
	p.P()

	p.P(Comment(""),
		"func WithFirestore(client *", clientType, ") *Firestore {")
	p.P("return &Firestore {client: client}")
	p.P("}")
	p.P()
}

func (p *Package) genHeader() {
	p.P("// Code generated by protoc-gen-go-firestore. DO NOT EDIT.")
	p.P("// versions:")

	protocGenGormVersion := version.String()
	protocVersion := "(unknown)"

	if v := p.plugin.Request.GetCompilerVersion(); v != nil {
		protocVersion = fmt.Sprintf("v%v.%v.%v", v.GetMajor(), v.GetMinor(), v.GetPatch())
		if s := v.GetSuffix(); s != "" {
			protocVersion += "-" + s
		}
	}

	p.P("// \tprotoc-gen-go-firestore ", protocGenGormVersion)
	p.P("// \tprotoc          ", protocVersion)
	p.P()
}

func (p *Package) genPackage() {
	p.P("package ", p.name)
	p.P()
}

func (p *Package) P(v ...interface{}) { p.out.P(v...) }

func (p *Package) genCollectionChainMethods() {
	p.tree.Walk(func(parent *tree.Parent[*Message], collection string, msg *Message) {
		p.genCollectionMethod(parent, collection, msg)
		p.genCollectionMethodCreate(parent, collection, msg)
		p.genCollectionMethodLimit(parent, collection, msg)
		p.genCollectionMethodOrderBy(parent, collection, msg)
		p.genCollectionMethodWhere(parent, collection)
		p.genCollectionType(parent, collection)
		p.genDocumentMethod(parent, collection)
		p.genDocumentMethodDelete(parent, collection, msg)
		p.genDocumentMethodGet(parent, collection, msg)
		p.genDocumentMethodRef(parent, collection, msg)
		p.genDocumentMethodSet(parent, collection, msg)
		p.genDocumentType(parent, collection, msg)
		p.genIteratorGetAll(parent, collection, msg)
		p.genIteratorNext(parent, collection, msg)
		p.genIteratorStop(parent, collection)
		p.genIteratorType(parent, collection, msg)
		p.genQueryMethodFirst(parent, collection, msg)
		p.genQueryType(parent, collection, msg)
		p.genQueryValue(parent, collection)
	})
}

func (p *Package) documentTypeName(parent *tree.Parent[*Message], collection string) string {
	return p.typeName(parent, collection) + "DocumentRef"
}

func (p *Package) genDocumentType(parent *tree.Parent[*Message], coll string, m *Message) {
	docType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "DocumentRef",
		GoImportPath: "cloud.google.com/go/firestore",
	})
	typeName := p.documentTypeName(parent, coll)
	p.P(Comment("%s holds a reference to a Firestore document in collection `%s`.",
		typeName,
		coll,
	),
		"type ", typeName, " struct {")
	p.P("doc *", docType)
	p.P("}")
	p.P()
}

func (p *Package) genDocumentMethod(parent *tree.Parent[*Message], collection string) {
	collTypeName := p.collectionTypeName(parent, collection)
	docTypeName := p.documentTypeName(parent, collection)
	p.P(Comment(""),
		"func (ref *", collTypeName, ") ",
		"Doc(id string)",
		"*", docTypeName,
		" {")
	p.P("return &", docTypeName, " {")
	p.P("doc: ref.coll.Doc(id),")
	p.P("}")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethod(parent *tree.Parent[*Message], collection string, msg *Message) {
	collTypeName := p.collectionTypeName(parent, collection)
	if parent == nil {
		p.P(Comment(""),
			"func (fs *Firestore)", title.String(collection), "()",
			"*", collTypeName,
			" {")
		p.P("return &", collTypeName, "{")
		if msg != nil {
			p.P("coll: fs.client.Collection(", msg.CollectionConstantName(), "),")
		} else {
			p.P("coll: fs.client.Collection(\"", collection, "\"),")
		}
		p.P("}")
		p.P("}")
	} else {
		p.P(Comment(""),
			"func (ref *", p.documentTypeName(parent, collection), ")", title.String(collection), "()",
			"*", collTypeName,
			" {")
		p.P("return &", collTypeName, "{")
		if msg != nil {
			p.P("coll: ref.doc.Collection(", msg.CollectionConstantName(), "),")
		} else {
			p.P("coll: ref.doc.Collection(\"", collection, "\"),")
		}
		p.P("}")
		p.P("}")
	}
}

func (p *Package) collectionTypeName(parent *tree.Parent[*Message], collection string) string {
	return p.typeName(parent, collection) + "CollectionRef"
}

func (p *Package) queryTypeName(parent *tree.Parent[*Message], collection string) string {
	return p.typeName(parent, collection) + "Query"
}

func (p *Package) iterTypeName(parent *tree.Parent[*Message], collection string) string {
	return p.typeName(parent, collection) + "Iterator"
}

func (p *Package) typeName(_ *tree.Parent[*Message], collection string) string {
	return "Firestore" + title.String(collection)
}

func (p *Package) genCollectionType(parent *tree.Parent[*Message], coll string) {
	collType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "CollectionRef",
		GoImportPath: "cloud.google.com/go/firestore",
	})
	typeName := p.collectionTypeName(parent, coll)
	p.P(Comment("%s holds a reference to the Firestore collection `%s`.",
		typeName,
		coll,
	),
		"type ", typeName, " struct {")
	p.P("coll *", collType)
	p.P("}")
	p.P()
}

func (p *Package) genIteratorType(parent *tree.Parent[*Message], collection string, msg *Message) {
	iterType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "DocumentIterator",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	iterTypeName := p.iterTypeName(parent, collection)

	p.P(Comment(""),
		"type ", iterTypeName, " struct {")
	p.P("iter *", iterType)
	p.P("}")
	p.P()
}

func (p *Package) genQueryType(parent *tree.Parent[*Message], collection string, msg *Message) {
	qType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Query",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	queryTypeName := p.queryTypeName(parent, collection)
	iterTypeName := p.iterTypeName(parent, collection)

	p.P(Comment(""),
		"type ", queryTypeName, " struct {")
	p.P("query ", qType)
	p.P("}")
	p.P()

	ctxType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})

	p.P(Comment(""),
		"func (q *", queryTypeName, ") Documents(",
		"ctx ", ctxType,
		")",
		"*", iterTypeName,
		" {")
	p.P("return &", iterTypeName, "{")
	p.P("iter: q.query.Documents(ctx),")
	p.P("}")
	p.P("}")
	p.P()
}

func (p *Package) genQueryValue(parent *tree.Parent[*Message], collection string) {
	qType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Query",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"func (q *", p.queryTypeName(parent, collection), ") Value()",
		qType,
		" {")
	p.P("return q.query")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethodWhere(parent *tree.Parent[*Message], collection string) {
	typeName := p.collectionTypeName(parent, collection)
	queryTypeName := p.queryTypeName(parent, collection)

	p.P(Comment(""),
		"func (ref *", typeName, ") Where(",
		"path, op string, value interface{}",
		")",
		"*", queryTypeName,
		" {")
	p.P("return &", queryTypeName, "{")
	p.P("query: ref.coll.Where(path, op, value),")
	p.P("}")
	p.P("}")
	p.P()

	p.P(Comment(""),
		"func (q *", queryTypeName, ") Where(",
		"path, op string, value interface{}",
		")",
		"*", queryTypeName,
		" {")
	p.P("return &", queryTypeName, "{")
	p.P("query: q.query.Where(path, op, value),")
	p.P("}")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethodOrderBy(parent *tree.Parent[*Message], collection string, msg *Message) {
	dirType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Direction",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	typeName := p.collectionTypeName(parent, collection)
	queryTypeName := p.queryTypeName(parent, collection)

	p.P(Comment(""),
		"func (ref *", typeName, ") OrderBy(",
		"path string, dir ", dirType,
		")",
		"*", queryTypeName,
		" {")
	p.P("return &", queryTypeName, "{")
	p.P("query: ref.coll.OrderBy(path, dir),")
	p.P("}")
	p.P("}")
	p.P()

	p.P(Comment(""),
		"func (q *", queryTypeName, ") OrderBy(",
		"path string, dir ", dirType,
		")",
		"*", queryTypeName,
		" {")
	p.P("return &", queryTypeName, "{")
	p.P("query: q.query.OrderBy(path, dir),")
	p.P("}")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethodLimit(parent *tree.Parent[*Message], collection string, msg *Message) {
	typeName := p.collectionTypeName(parent, collection)
	queryTypeName := p.queryTypeName(parent, collection)

	p.P(Comment(""),
		"func (ref *", typeName, ") Limit(",
		"n int",
		") ",
		"*", queryTypeName,
		" {")
	p.P("return &", queryTypeName, "{")
	p.P("query: ref.coll.Limit(n),")
	p.P("}")
	p.P("}")
	p.P()

	p.P(Comment(""),
		"func (q *", queryTypeName, ") Limit(",
		"n int",
		") ",
		"*", queryTypeName,
		" {")
	p.P("return &", queryTypeName, "{")
	p.P("query: q.query.Limit(n),")
	p.P("}")
	p.P("}")
	p.P()
}

func (p *Package) genIteratorGetAll(parent *tree.Parent[*Message], collection string, msg *Message) {
	if msg == nil {
		return
	}

	snapType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "DocumentSnapshot",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"func (i *", p.iterTypeName(parent, collection), ") GetAll() (",
		"[]*", msg.ProtoName(), ", ",
		"error",
		") {")

	p.P("snaps, err := i.iter.GetAll()")
	p.P("if err != nil {")
	p.P("return nil, err")
	p.P("}")

	p.P("protos := make([]*", msg.ProtoName(), ", len(snaps))")

	p.P("for j, snapshot := range snaps {")
	{
		p.P("o := new(", msg.CustomObjectName(), ")")
		p.P("if err := snapshot.DataTo(o); err != nil {")
		p.P("return nil, err")
		p.P("}")

		p.P("if p, err := o.ToProto(); err != nil {")
		p.P("return nil, err")
		p.P("} else {")
		p.P("protos[j] = p")
		p.P("}")
	}
	p.P("}")

	p.P("return protos, nil")
	p.P("}") // func
	p.P()

	p.P(Comment(""),
		"func (i *", p.iterTypeName(parent, collection), ") GetAllAsSnapshots() (",
		"[]*", snapType, ", ",
		"error",
		") {")
	p.P("return i.iter.GetAll()")
	p.P("}")
	p.P()
}

func (p *Package) genIteratorNext(parent *tree.Parent[*Message], collection string, msg *Message) {
	if msg == nil {
		return
	}

	snapType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "DocumentSnapshot",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"func (i *", p.iterTypeName(parent, collection), ") Next() (",
		"*", msg.ProtoName(), ", ",
		"error",
		") {")
	p.P("snapshot, err := i.iter.Next()")
	p.P("if err != nil {")
	p.P("return nil, err")
	p.P("}")

	p.P("obj := new(", msg.CustomObjectName(), ")")
	p.P("if err := snapshot.DataTo(obj); err != nil {")
	p.P("return nil, err")
	p.P("}")

	p.P("if p, err := obj.ToProto(); err != nil {")
	p.P("return nil, err")
	p.P("} else {")
	p.P("return p, nil")
	p.P("}")

	p.P("}") // func
	p.P()

	p.P(Comment(""),
		"func (i *", p.iterTypeName(parent, collection), ") NextAsSnapshot() (",
		"*", snapType, ", ",
		"error",
		") {")
	p.P("return i.iter.Next()")
	p.P("}")
	p.P()
}

func (p *Package) genIteratorStop(parent *tree.Parent[*Message], collection string) {
	p.P(Comment(""),
		"func (i *", p.iterTypeName(parent, collection), ") Stop() {")
	p.P("i.iter.Stop()")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethodCreate(parent *tree.Parent[*Message], collection string, msg *Message) {
	if msg == nil || msg.idField == nil {
		return
	}

	ctxType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})

	resultType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "WriteResult",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	statusErrType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Error",
		GoImportPath: "google.golang.org/grpc/status",
	})

	codeType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "InvalidArgument",
		GoImportPath: "google.golang.org/grpc/codes",
	})

	collTypeName := p.collectionTypeName(parent, collection)

	p.P(Comment(""),
		"func (ref *", collTypeName, ") ",
		"Create(",
		"ctx ", ctxType, ", ",
		"p *", msg.ProtoName(),
		") (",
		"*", resultType, ", ",
		"error",
		") {")
	{
		p.P("fs, err := p.ToFirestore()")
		p.P("if err != nil {")
		p.P("return nil, err")
		p.P("}")

		p.P("id := fs.", msg.idField.Name())
		p.P("if id == \"\" {")
		p.P("return nil, ", statusErrType, "(", codeType, ", \"empty id\")")
		p.P("}")

		p.P("res, err := ref.coll.Doc(id).Create(ctx, fs)")
		p.P("if err != nil {")
		p.P("return nil, err")
		p.P("}")
		p.P("return res, nil")
	}
	p.P("}") // func
	p.P()
}

func (p *Package) genQueryMethodFirst(parent *tree.Parent[*Message], collection string, msg *Message) {
	if msg == nil {
		return
	}

	ctxType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})
	doneType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Done",
		GoImportPath: "google.golang.org/api/iterator",
	})
	errorsIsType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Is",
		GoImportPath: "errors",
	})

	queryTypeName := p.queryTypeName(parent, collection)

	p.P(Comment(""),
		"func (q *", queryTypeName, ") First(ctx ", ctxType, ") (",
		"*", msg.ProtoName(), ", ",
		"error",
		") {")

	p.P("iter := q.query.Limit(1).Documents(ctx)")
	p.P("defer iter.Stop()")

	p.P("snapshot, err := iter.Next()")
	p.P("if err != nil {")
	{
		p.P("if  ", errorsIsType, "(err, ", doneType, ") {")
		p.P("return nil, nil")
		p.P("}")
	}
	p.P("return nil, err")
	p.P("}")

	p.P("obj := new(", msg.CustomObjectName(), ")")
	p.P("if err := snapshot.DataTo(obj); err != nil {")
	p.P("return nil, err")
	p.P("}")

	p.P("if proto, err := obj.ToProto(); err != nil {")
	p.P("return nil, err")
	p.P("} else {")
	p.P("return proto, nil")
	p.P("}")

	p.P("}") // func
	p.P()
}

func (p *Package) genDocumentMethodGet(parent *tree.Parent[*Message], collection string, msg *Message) {
	if msg == nil {
		return
	}

	ctxType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})

	p.P(Comment(""),
		"func (ref *", p.documentTypeName(parent, collection), ") ",
		"Get(", "ctx ", ctxType, ") ",
		"(",
		"*", msg.ProtoName(), ", error",
		")",
		" {")

	{
		p.P("snapshot, err := ref.doc.Get(ctx)")
		p.P("if err != nil {")
		p.P("return nil, err")
		p.P("}")

		p.P("obj := new(", msg.CustomObjectName(), ")")
		p.P("if err := snapshot.DataTo(obj); err != nil {")
		p.P("return nil, err")
		p.P("}")

		p.P("if proto, err := obj.ToProto(); err != nil {")
		p.P("return nil, err")
		p.P("} else {")
		p.P("return proto, nil")
		p.P("}")
	}

	p.P("}")
	p.P()
}

func (p *Package) genDocumentMethodSet(parent *tree.Parent[*Message], collection string, msg *Message) {
	if msg == nil {
		return
	}

	ctxType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})

	p.P(Comment(""),
		"func (ref *", p.documentTypeName(parent, collection), ") ",
		"Set(",
		"ctx ", ctxType, ", ",
		"msg *", msg.ProtoName(),
		") ",
		"error ",
		" {")

	p.P("fs, err := msg.ToFirestore()")
	p.P("if err != nil {")
	p.P("return err")
	p.P("}")

	p.P("if _, err := ref.doc.Set(ctx, fs); err != nil {")
	p.P("return err")
	p.P("}")

	p.P("return nil")
	p.P("}")
	p.P()
}

func (p *Package) genDocumentMethodDelete(parent *tree.Parent[*Message], collection string, msg *Message) {
	if msg == nil {
		return
	}

	ctxType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})

	preconType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Precondition",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	resultType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "WriteResult",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"func (ref *", p.documentTypeName(parent, collection), ") ",
		"Delete(",
		"ctx ", ctxType, ", ",
		"preconds ...", preconType,
		") ",
		"(*", resultType, ", error) ",
		" {")

	p.P("return ref.doc.Delete(ctx, preconds...)")
	p.P("}")
	p.P()
}

func (p *Package) genDocumentMethodRef(parent *tree.Parent[*Message], collection string, msg *Message) {
	if msg == nil {
		return
	}

	refType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "DocumentRef",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"func (ref *", p.documentTypeName(parent, collection), ") ",
		"Ref() ",
		"*", refType,
		" {")
	p.P("return ref.doc")
	p.P("}")
	p.P()
}
