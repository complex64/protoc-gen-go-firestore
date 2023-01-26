package gen

import (
	"strings"

	orderedmap "github.com/wk8/go-ordered-map/v2"
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
			}
			if doc != nil {
				doc.Collections.Set(coll.Segment, coll)
			}
			if i == 0 {
				root.Collection = coll
			}
			if i+1 == len(elems) {
				coll.Message = m
			}
		} else {
			doc = &Document{
				Collections: orderedmap.New[string, *Collection](),
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

func (left *Collection) Merge(right *Collection) *Collection {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.Segment != right.Segment {
		panic("BUG")
	}
	if left.Title != right.Title {
		panic("BUG")
	}

	merged := new(Collection)
	merged.Segment = left.Segment
	merged.Title = left.Title

	if left.Message == nil {
		merged.Message = right.Message
	} else {
		merged.Message = left.Message
	}

	if left.Document == nil && right.Document != nil {
		merged.Document = right.Document
	}
	if right.Document == nil && left.Document != nil {
		merged.Document = left.Document
	}

	merged.Document = left.Document.Merge(right.Document)
	merged.Document.Parent = merged
	return merged
}

type Document struct {
	Collections *orderedmap.OrderedMap[string, *Collection]
	Parent      *Collection
}

func (left *Document) Merge(right *Document) *Document {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	combined := orderedmap.New[string, *Collection]()

	for pair := left.Collections.Oldest(); pair != nil; pair = pair.Next() {
		if existing, ok := combined.Get(pair.Key); !ok {
			combined.Set(pair.Key, pair.Value)
		} else {
			combined.Set(pair.Key, pair.Value.Merge(existing))
		}
	}

	for pair := right.Collections.Oldest(); pair != nil; pair = pair.Next() {
		if existing, ok := combined.Get(pair.Key); !ok {
			combined.Set(pair.Key, pair.Value)
		} else {
			combined.Set(pair.Key, pair.Value.Merge(existing))
		}
	}

	merged := &Document{Collections: combined}

	if left.Parent == nil {
		merged.Parent = right.Parent
	} else {
		merged.Parent = left.Parent
	}

	return merged
}
