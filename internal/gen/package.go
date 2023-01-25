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
	collection *Collection,
) {
	if collection == nil {
		return
	}
	p.genCollectionMethod(collection)
	p.genCollectionType(collection)
	p.genDocumentMethod(collection)
	p.genDocumentType(collection)

	if collection.Document != nil {
		p.genDocumentChainMethod(collection, collection.Document)
	}
}

func (p *Package) genDocumentType(collection *Collection) {
	docType := p.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "DocumentRef",
		GoImportPath: "cloud.google.com/go/firestore",
	})

	p.P(Comment(""),
		"type ", collection.NestedDocumentTypeName(p.packageFirestoreType()), " struct {")
	p.P("document *", docType)
	p.P("}")
	p.P()
}

func (p *Package) genDocumentMethod(collection *Collection) {
	p.P(Comment(""),
		"func (f *", collection.TypeName(p.packageFirestoreType()), ") ",
		"Doc(id string)",
		"*", collection.NestedDocumentTypeName(p.packageFirestoreType()),
		" {")
	p.P("return &", collection.NestedDocumentTypeName(p.packageFirestoreType()), " {")
	p.P("document: f.collection.Doc(id),")
	p.P("}")
	p.P("}")
	p.P()
}

func (p *Package) genCollectionMethod(c *Collection) {
	p.P(Comment(""),
		"func (f *", c.ParentDocumentTypeName(p.packageFirestoreType()), ") ", c.Title, "()",
		"*", c.TypeName(p.packageFirestoreType()),
		" {")
	p.P("return &", c.TypeName(p.packageFirestoreType()), "{")

	if c.Parent == nil {
		p.P("collection: f.client.Collection(\"", c.Segment, "\"),")
	} else {
		p.P("collection: f.document.Collection(\"", c.Segment, "\"),")
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
	p.P("collection *", collType)
	p.P("}")
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

func (c *Collection) ParentDocumentTypeName(prefix string) string {
	if c.Parent == nil {
		return prefix
	}
	return c.Parent.TypeName(prefix)
}
