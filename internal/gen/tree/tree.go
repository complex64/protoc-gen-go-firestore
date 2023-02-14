package tree

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidFormat = errors.New("expected (<collection>/<id>)+")
)

type Tree[T any] struct {
	root *node[T]
}

type collection[T any] struct {
	node     *node[T]
	children *node[T]
	name     string
	value    T
}

type node[T any] struct {
	collections []*collection[T]
	parent      *node[T]
}

func (coll *collection[T]) add(elems []string, value T) {
	if len(elems) == 0 {
		coll.value = value
		return
	}
	if coll.children == nil {
		coll.children = &node[T]{parent: coll.node}
	}
	coll.children.add(elems, value)
}

func (t *Tree[T]) Add(path string, value T) error {
	path = strings.TrimPrefix(path, "/")
	split := strings.Split(path, "/")
	if len(split) < 2 || len(split)%2 == 1 {
		return fmt.Errorf(
			"expected (<collection>/<id>)[/<collection>/<id>]*, "+
				"but have `%s`", path)
	}
	t.add(split, value)
	return nil
}

func (t *Tree[T]) add(elems []string, value T) {
	if t.root == nil {
		t.root = &node[T]{}
	}
	t.root.add(elems, value)
}

func (n *node[T]) add(elems []string, value T) {
	name := elems[0]

	for _, coll := range n.collections {
		if coll.name == name {
			coll.add(elems[2:], value)
			return
		}
	}

	c := &collection[T]{
		node: n,
		name: name,
	}
	n.collections = append(n.collections, c)
	c.add(elems[2:], value)
}
