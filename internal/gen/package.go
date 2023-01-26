package gen

import (
	"fmt"
	"path"

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
	name := string(file.proto.GoPackageName)
	if _, ok := p.list[name]; !ok {
		p.list[name] = &Package{
			plugin:      p.plugin,
			name:        name,
			importPath:  file.proto.GoImportPath,
			collections: map[string]*Collection{},
		}
		p.list[name].initOut(file)
	}
	pkg := p.list[name]

	for msgName, msg := range file.msgs {
		_ = msgName
		pkg.CollectPath(msg.path)
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
	plugin      *protogen.Plugin
	name        string
	dirPrefix   string
	importPath  protogen.GoImportPath
	out         *protogen.GeneratedFile
	collections map[string]*Collection
}

func (p *Package) CollectPath(pa *Path) {
	if pa == nil || pa.Collection == nil {
		return
	}
	if existing, ok := p.collections[pa.Collection.Segment]; ok {
		p.collections[pa.Collection.Segment] = existing.Merge(pa.Collection)
	} else {
		p.collections[pa.Collection.Segment] = pa.Collection
	}
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
	p.genCollectionChainMethods(nil, p.collections)
}

func (p *Package) genFirestoreTypeAndMethod() {
	clientType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Client",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	firestoreType := p.packageFirestoreType()
	p.P(Comment(""),
		"type ", firestoreType, " struct {")
	p.P("client *", clientType)
	p.P("}")
	p.P()

	p.P(Comment(""),
		"func Firestore(client *", clientType, ") *", firestoreType, " {")
	p.P("return &", firestoreType, "{")
	p.P("client: client", ",")
	p.P("}")
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

func (p *Package) packageFirestoreType() string {
	return "FS_" + p.name
}

func (p *Package) genCollectionChainMethods(
	parent *Collection,
	collections map[string]*Collection,
) {
	for _, coll := range collections {
		p.genCollectionChainMethod(parent, coll)
	}
}

func (p *Package) genCollectionChainMethod(
	parent *Collection,
	c *Collection,
) {
	if c == nil {
		return
	}
	p.genCollectionMethod(c)
	p.genCollectionType(c)
	p.genCollectionTypeIterator(c)
	p.genCollectionTypeQuery(c)
	p.genCollectionMethodQueryValue(c)
	p.genCollectionMethodWhere(c)
	p.genCollectionMethodOrderBy(c)
	p.genCollectionMethodLimit(c)
	p.genCollectionMethodFirst(c)
	p.genCollectionMethodIterGetAll(c)
	p.genCollectionMethodIterNext(c)
	p.genCollectionMethodIterStop(c)
	p.genDocumentMethod(c)
	p.genDocumentType(c)
	p.genDocumentMethodGet(c)
	p.genDocumentMethodSet(c)
	p.genDocumentMethodDelete(c)
	p.genDocumentMethodRef(c)

	if c.Document != nil {
		p.genDocumentChainMethod(c, c.Document)
	}
}

func (p *Package) genDocumentType(c *Collection) {
	docType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "DocumentRef",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"type ", c.NestedDocumentTypeName(p.packageFirestoreType()), " struct {")
	p.P("d *", docType)
	p.P("}")
	p.P()
}

func (p *Package) genDocumentMethod(c *Collection) {
	p.P(Comment(""),
		"func (x *", c.TypeName(p.packageFirestoreType()), ") ",
		"Doc(id string)",
		"*", c.NestedDocumentTypeName(p.packageFirestoreType()),
		" {")
	p.P("return &", c.NestedDocumentTypeName(p.packageFirestoreType()), " {")
	p.P("d: x.c.Doc(id),")
	p.P("}")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethod(c *Collection) {
	segment := "\"" + c.Segment + "\""
	if c.Message != nil {
		segment = c.Message.CollectionConstantName()
	}

	p.P(Comment(""),
		"func (x *", c.ParentDocumentTypeName(p.packageFirestoreType()), ") ", c.Title, "()",
		"*", c.TypeName(p.packageFirestoreType()),
		" {")
	p.P("return &", c.TypeName(p.packageFirestoreType()), "{")

	if c.Parent == nil {
		p.P("c: x.client.Collection(", segment, "),")
	} else {
		p.P("c: x.d.Collection(", segment, "),")
	}

	p.P("}")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionType(c *Collection) {
	collType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "CollectionRef",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"type ", c.TypeName(p.packageFirestoreType()), " struct {")
	p.P("c *", collType)
	p.P("}")
	p.P()
}

func (p *Package) genCollectionTypeIterator(c *Collection) {
	iterType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "DocumentIterator",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"type ", c.TypeNameIter(p.packageFirestoreType()), " struct {")
	p.P("i *", iterType)
	p.P("}")
	p.P()
}

func (p *Package) genCollectionTypeQuery(c *Collection) {
	qType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Query",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"type ", c.TypeNameQuery(p.packageFirestoreType()), " struct {")
	p.P("q ", qType)
	p.P("}")
	p.P()

	ctxType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})

	p.P(Comment(""),
		"func (x *", c.TypeNameQuery(p.packageFirestoreType()), ") Documents(",
		"ctx ", ctxType,
		")",
		"*", c.TypeNameIter(p.packageFirestoreType()),
		" {")
	p.P("return &", c.TypeNameIter(p.packageFirestoreType()), "{")
	p.P("i: x.q.Documents(ctx),")
	p.P("}")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethodQueryValue(c *Collection) {
	qType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Query",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"func (x *", c.TypeNameQuery(p.packageFirestoreType()), ") Value()",
		qType,
		" {")
	p.P("return x.q")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethodWhere(c *Collection) {
	p.P(Comment(""),
		"func (x *", c.TypeName(p.packageFirestoreType()), ") Where(",
		"path, op string, value interface{}",
		")",
		"*", c.TypeNameQuery(p.packageFirestoreType()),
		" {")
	p.P("return &", c.TypeNameQuery(p.packageFirestoreType()), "{")
	p.P("q: x.c.Where(path, op, value),")
	p.P("}")
	p.P("}")
	p.P()

	p.P(Comment(""),
		"func (x *", c.TypeNameQuery(p.packageFirestoreType()), ") Where(",
		"path, op string, value interface{}",
		")",
		"*", c.TypeNameQuery(p.packageFirestoreType()),
		" {")
	p.P("return &", c.TypeNameQuery(p.packageFirestoreType()), "{")
	p.P("q: x.q.Where(path, op, value),")
	p.P("}")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethodOrderBy(c *Collection) {
	dirType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Direction",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"func (x *", c.TypeName(p.packageFirestoreType()), ") OrderBy(",
		"path string, dir ", dirType,
		")",
		"*", c.TypeNameQuery(p.packageFirestoreType()),
		" {")
	p.P("return &", c.TypeNameQuery(p.packageFirestoreType()), "{")
	p.P("q: x.c.OrderBy(path, dir),")
	p.P("}")
	p.P("}")
	p.P()

	p.P(Comment(""),
		"func (x *", c.TypeNameQuery(p.packageFirestoreType()), ") OrderBy(",
		"path string, dir ", dirType,
		")",
		"*", c.TypeNameQuery(p.packageFirestoreType()),
		" {")
	p.P("return &", c.TypeNameQuery(p.packageFirestoreType()), "{")
	p.P("q: x.q.OrderBy(path, dir),")
	p.P("}")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethodLimit(c *Collection) {
	p.P(Comment(""),
		"func (x *", c.TypeName(p.packageFirestoreType()), ") Limit(",
		"n int",
		")",
		"*", c.TypeNameQuery(p.packageFirestoreType()),
		" {")
	p.P("return &", c.TypeNameQuery(p.packageFirestoreType()), "{")
	p.P("q: x.c.Limit(n),")
	p.P("}")
	p.P("}")
	p.P()

	p.P(Comment(""),
		"func (x *", c.TypeNameQuery(p.packageFirestoreType()), ") Limit(",
		"n int",
		")",
		"*", c.TypeNameQuery(p.packageFirestoreType()),
		" {")
	p.P("return &", c.TypeNameQuery(p.packageFirestoreType()), "{")
	p.P("q: x.q.Limit(n),")
	p.P("}")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethodIterGetAll(c *Collection) {
	if c.Message == nil {
		return
	}

	snapType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "DocumentSnapshot",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"func (x *", c.TypeNameIter(p.packageFirestoreType()), ") GetAll() (",
		"[]*", c.Message.ProtoName(), ", ",
		"error",
		") {")

	p.P("snaps, err := x.i.GetAll()")
	p.P("if err != nil {")
	p.P("return nil, err")
	p.P("}")

	p.P("protos := make([]*", c.Message.ProtoName(), ", len(snaps))")

	p.P("for i, snap := range snaps {")
	{
		p.P("o := new(", c.Message.CustomObjectName(), ")")
		p.P("if err := snap.DataTo(o); err != nil {")
		p.P("return nil, err")
		p.P("}")

		p.P("if p, err := o.ToProto(); err != nil {")
		p.P("return nil, err")
		p.P("} else {")
		p.P("protos[i] = p")
		p.P("}")
	}
	p.P("}")

	p.P("return protos, nil")
	p.P("}") // func
	p.P()

	p.P(Comment(""),
		"func (x *", c.TypeNameIter(p.packageFirestoreType()), ") GetAllAsSnapshots() (",
		"[]*", snapType, ", ",
		"error",
		") {")
	p.P("return x.i.GetAll()")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethodIterNext(c *Collection) {
	if c.Message == nil {
		return
	}

	snapType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "DocumentSnapshot",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"func (x *", c.TypeNameIter(p.packageFirestoreType()), ") Next() (",
		"*", c.Message.ProtoName(), ", ",
		"error",
		") {")
	p.P("snap, err := x.i.Next()")
	p.P("if err != nil {")
	p.P("return nil, err")
	p.P("}")

	p.P("o := new(", c.Message.CustomObjectName(), ")")
	p.P("if err := snap.DataTo(o); err != nil {")
	p.P("return nil, err")
	p.P("}")

	p.P("if p, err := o.ToProto(); err != nil {")
	p.P("return nil, err")
	p.P("} else {")
	p.P("return p, nil")
	p.P("}")

	p.P("}") // func
	p.P()

	p.P(Comment(""),
		"func (x *", c.TypeNameIter(p.packageFirestoreType()), ") NextAsSnapshot() (",
		"*", snapType, ", ",
		"error",
		") {")
	p.P("return x.i.Next()")
	p.P("}")
	p.P()

}

func (p *Package) genCollectionMethodIterStop(c *Collection) {
	p.P(Comment(""),
		"func (x *", c.TypeNameIter(p.packageFirestoreType()), ") Stop() {")
	p.P("x.i.Stop()")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethodFirst(c *Collection) {
	if c.Message == nil {
		return
	}

	ctxType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})

	p.P(Comment(""),
		"func (x *", c.TypeNameQuery(p.packageFirestoreType()), ") First(ctx ", ctxType, ") (",
		"*", c.Message.ProtoName(), ", ",
		"error",
		") {")

	p.P("iter := x.q.Limit(1).Documents(ctx)")
	p.P("defer iter.Stop()")

	p.P("snap, err := iter.Next()")
	p.P("if err != nil {")
	p.P("return nil, err")
	p.P("}")

	p.P("o := new(", c.Message.CustomObjectName(), ")")
	p.P("if err := snap.DataTo(o); err != nil {")
	p.P("return nil, err")
	p.P("}")

	p.P("if p, err := o.ToProto(); err != nil {")
	p.P("return nil, err")
	p.P("} else {")
	p.P("return p, nil")
	p.P("}")

	p.P("}") // func
	p.P()
}

func (p *Package) genDocumentChainMethod(parent *Collection, doc *Document) {
	if len(doc.Collections) > 0 {
		p.genCollectionChainMethods(parent, doc.Collections)
	}
}

func (d *Document) TypeName(prefix string) string {
	if d.Parent != nil {
		return d.Parent.TypeName(prefix) + "_Doc"
	}
	return prefix + "_Doc"
}

func (c *Collection) NestedDocumentTypeName(prefix string) string {
	return c.TypeName(prefix) + "_Doc"
}

func (c *Collection) TypeName(prefix string) string {
	t := ""
	var cur = c
	for {
		if cur == nil {
			if t == "" {
				return prefix
			}
			return prefix + "_" + t
		}

		if t == "" {
			t = cur.Title
		} else {
			t = cur.Title + "_" + t
		}

		if cur.Parent != nil && cur.Parent.Parent != nil {
			cur = cur.Parent.Parent
		} else {
			cur = nil
		}
	}
}

func (c *Collection) TypeNameQuery(prefix string) string {
	return c.TypeName(prefix) + "_Query"
}

func (c *Collection) TypeNameIter(prefix string) string {
	return c.TypeName(prefix) + "_Iter"
}

func (c *Collection) ParentDocumentTypeName(prefix string) string {
	if c.Parent == nil {
		return prefix
	}
	return c.Parent.TypeName(prefix)
}

func (p *Package) genDocumentMethodGet(c *Collection) {
	if c.Message == nil {
		return
	}

	ctxType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})

	// refType := p.out.QualifiedGoIdent(protogen.GoIdent{
	// 	GoName:       "DocumentRef",
	// 	GoImportPath: "cloud.google.com/go/firestore",
	// })

	p.P(Comment(""),
		"func (x *", c.NestedDocumentTypeName(p.packageFirestoreType()), ") ",
		"Get(", "ctx ", ctxType, ") ",
		// "*", c.NestedDocumentTypeName(p.packageFirestoreType()),
		"(",
		"*", c.Message.ProtoName(), ", error",
		")",
		" {")

	{
		p.P("snap, err := x.d.Get(ctx)")
		p.P("if err != nil {")
		p.P("return nil, err")
		p.P("}")

		p.P("o := new(", c.Message.CustomObjectName(), ")")
		p.P("if err := snap.DataTo(o); err != nil {")
		p.P("return nil, err")
		p.P("}")

		p.P("if p, err := o.ToProto(); err != nil {")
		p.P("return nil, err")
		p.P("} else {")
		p.P("return p, nil")
		p.P("}")
	}

	p.P("}")
	p.P()
}

func (p *Package) genDocumentMethodSet(c *Collection) {
	if c.Message == nil {
		return
	}

	ctxType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})

	p.P(Comment(""),
		"func (x *", c.NestedDocumentTypeName(p.packageFirestoreType()), ") ",
		"Set(",
		"ctx ", ctxType, ", ",
		"m *", c.Message.ProtoName(),
		") ",
		// "*", c.NestedDocumentTypeName(p.packageFirestoreType()),
		"error ",
		" {")

	p.P("fs, err := m.ToFirestore()")
	p.P("if err  != nil {")
	p.P("return err")
	p.P("}")

	p.P("if _, err := x.d.Set(ctx, ", "fs); err != nil {")
	p.P("return err")
	p.P("}")

	p.P("return nil")
	p.P("}")
	p.P()
}

func (p *Package) genDocumentMethodDelete(c *Collection) {
	if c.Message == nil {
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
		"func (x *", c.NestedDocumentTypeName(p.packageFirestoreType()), ") ",
		"Delete(",
		"ctx ", ctxType, ", ",
		"preconds ...", preconType,
		") ",
		"(*", resultType, ", error) ",
		" {")

	p.P("return x.d.Delete(ctx, preconds...)")
	p.P("}")
	p.P()
}

func (p *Package) genDocumentMethodRef(c *Collection) {
	if c.Message == nil {
		return
	}

	refType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "DocumentRef",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"func (x *", c.NestedDocumentTypeName(p.packageFirestoreType()), ") ",
		"Ref() ",
		"*", refType,
		" {")
	p.P("return x.d")
	p.P("}")
	p.P()
}
