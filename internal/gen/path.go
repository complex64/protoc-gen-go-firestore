package gen

import (
	"strings"
)

func (m *Message) parseCollection() (*Path, error) {
	elems := strings.Split(m.opts.Collection, "/")
	root := &Path{}

	var (
		coll *Collection
		doc  *Document
	)

	for i := 0; i < len(elems); i++ {
		if collection := i%2 == 0; collection {
			coll = &Collection{
				Segment: elems[i],
				Title:   title.String(elems[i]),
				Parent:  doc,
				Message: m,
			}
			if doc != nil {
				doc.Collections[coll.Segment] = coll
			}
			if i == 0 {
				root.Collection = coll
			}
		} else {
			doc = &Document{
				Collections: map[string]*Collection{},
				Parent:      coll,
			}
			coll.Document = doc
		}
	}

	return root, nil
}

type Path struct {
	Collection *Collection
}

type Collection struct {
	Segment  string
	Title    string
	Document *Document
	Parent   *Document
	Message  *Message
}

func (c *Collection) Merge(right *Collection) *Collection {
	if c == nil {
		return right
	}
	if right == nil {
		return c
	}
	if c.Document == nil && right.Document != nil {
		return right
	}
	if right.Document == nil && c.Document != nil {
		return c
	}
	c.Document = c.Document.Merge(right.Document)
	return c
}

type Document struct {
	Collections map[string]*Collection
	Parent      *Collection
}

func (d *Document) Merge(right *Document) *Document {
	if d == nil {
		return right
	} else if right == nil {
		return d
	}

	combined := map[string]*Collection{}

	for seg, coll := range d.Collections {
		if existing, ok := combined[seg]; !ok {
			combined[seg] = coll
		} else {
			combined[seg] = combined[seg].Merge(existing)
		}
	}

	for seg, coll := range right.Collections {
		if existing, ok := combined[seg]; !ok {
			combined[seg] = coll
		} else {
			combined[seg] = combined[seg].Merge(existing)
		}
	}

	return &Document{Collections: combined}
}
