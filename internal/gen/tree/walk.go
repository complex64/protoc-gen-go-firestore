package tree

import (
	"sort"
)

type Parent[T any] struct {
	Parent *Parent[T]
	Name   string
	Value  T
}

type WalkFunc[T any] func(parent *Parent[T], name string, value T)

func (t *Tree[T]) Walk(f WalkFunc[T]) {
	t.root.walk(nil, f)
}

func (n *node[T]) walk(parent *Parent[T], f WalkFunc[T]) {
	if n == nil {
		return
	}
	sort.SliceStable(n.collections, func(i, j int) bool {
		return n.collections[i].name < n.collections[j].name
	})
	for _, c := range n.collections {
		c.walk(parent, f)
	}
}

func (coll *collection[T]) walk(parent *Parent[T], f WalkFunc[T]) {
	newParent := &Parent[T]{
		Parent: parent,
		Name:   coll.name,
		Value:  coll.value,
	}
	f(parent, coll.name, coll.value)
	coll.children.walk(newParent, f)
}
