package tree

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name   string
		paths  []string
		values []string
		want   func() *Tree[string]
		error  error
	}{
		{
			name:   "empty",
			paths:  []string{""},
			values: []string{""},
			want:   func() *Tree[string] { return new(Tree[string]) },
			error:  errors.New("expected (<collection>/<id>)[/<collection>/<id>]*, but have ``"),
		},
		{
			name:   "no elements",
			paths:  []string{"/"},
			values: []string{""},
			want:   func() *Tree[string] { return new(Tree[string]) },
			error:  errors.New("expected (<collection>/<id>)[/<collection>/<id>]*, but have ``"),
		},
		{
			name:   "one element",
			paths:  []string{"collection"},
			values: []string{""},
			want:   func() *Tree[string] { return new(Tree[string]) },
			error:  errors.New("expected (<collection>/<id>)[/<collection>/<id>]*, but have `collection`"),
		},
		{
			name:   "simple collection",
			paths:  []string{"collection/{id}"},
			values: []string{"first"},
			want: func() *Tree[string] {
				coll := &collection[string]{
					name:  "collection",
					value: "first",
				}
				root := &node[string]{
					parent: nil,
					collections: []*collection[string]{
						coll,
					},
				}
				coll.node = root
				return &Tree[string]{root: root}
			},
			error: nil,
		},
		{
			name:   "single level nesting",
			paths:  []string{"foo/{id}/bar/{id}"},
			values: []string{"first"},
			want: func() *Tree[string] {
				bars := &collection[string]{
					name:  "bar",
					value: "first",
				}

				foos := &collection[string]{
					name: "foo",
					children: &node[string]{
						collections: []*collection[string]{
							bars,
						},
					},
				}

				root := &node[string]{
					parent: nil,
					collections: []*collection[string]{
						foos,
					},
				}

				foos.node = root // foo collection lives in the root node
				foos.children.parent = root
				bars.node = foos.children

				return &Tree[string]{root: root}
			},
			error: nil,
		},
		{
			name:   "double nesting",
			paths:  []string{"foo/{id}/bar/{id}/baz/{id}"},
			values: []string{"first"},
			want: func() *Tree[string] {
				bazs := &collection[string]{
					name:  "baz",
					value: "first",
				}

				bars := &collection[string]{
					name: "bar",
					children: &node[string]{
						collections: []*collection[string]{
							bazs,
						},
					},
				}

				foos := &collection[string]{
					name: "foo",
					children: &node[string]{
						collections: []*collection[string]{
							bars,
						},
					},
				}

				root := &node[string]{
					parent: nil,
					collections: []*collection[string]{
						foos,
					},
				}

				foos.node = root
				foos.children.parent = root

				bars.node = foos.children
				bars.children.parent = bars.node

				bazs.node = bars.children

				return &Tree[string]{root: root}
			},
			error: nil,
		},
		{
			name: "shorter path added last",
			paths: []string{
				"foo/{id}/bar/{id}",
				"foo/{id}",
			},
			values: []string{
				"long",
				"short",
			},
			want: func() *Tree[string] {
				bars := &collection[string]{
					name:  "bar",
					value: "long",
				}

				foos := &collection[string]{
					name:  "foo",
					value: "short",
					children: &node[string]{
						collections: []*collection[string]{
							bars,
						},
					},
				}

				root := &node[string]{
					parent: nil,
					collections: []*collection[string]{
						foos,
					},
				}

				foos.node = root
				foos.children.parent = root
				bars.node = foos.children

				return &Tree[string]{root: root}
			},
			error: nil,
		},
		{
			name: "shorter path added first",
			paths: []string{
				"foo/{id}",
				"foo/{id}/bar/{id}",
			},
			values: []string{
				"short",
				"long",
			},
			want: func() *Tree[string] {
				bars := &collection[string]{
					name:  "bar",
					value: "long",
				}

				foos := &collection[string]{
					name:  "foo",
					value: "short",
					children: &node[string]{
						collections: []*collection[string]{
							bars,
						},
					},
				}

				root := &node[string]{
					parent: nil,
					collections: []*collection[string]{
						foos,
					},
				}

				foos.node = root
				foos.children.parent = root
				bars.node = foos.children

				return &Tree[string]{root: root}
			},
			error: nil,
		},
		{
			name: "value gap",
			paths: []string{
				"foo/{id}/bar/{id}/baz/{id}",
				"foo/{id}",
			},
			values: []string{
				"last",
				"first",
			},
			want: func() *Tree[string] {
				bazs := &collection[string]{
					name:  "baz",
					value: "last",
				}

				bars := &collection[string]{
					name: "bar",
					children: &node[string]{
						collections: []*collection[string]{
							bazs,
						},
					},
				}

				foos := &collection[string]{
					name:  "foo",
					value: "first",
					children: &node[string]{
						collections: []*collection[string]{
							bars,
						},
					},
				}

				root := &node[string]{
					parent: nil,
					collections: []*collection[string]{
						foos,
					},
				}

				foos.node = root
				foos.children.parent = root
				bars.node = foos.children
				bars.children.parent = bars.node
				bazs.node = bars.children

				return &Tree[string]{root: root}
			},
			error: nil,
		},
	}

	for _, k := range cases {
		k := k
		t.Run(k.name, func(t *testing.T) {
			t.Parallel()
			tree := new(Tree[string])
			for i, p := range k.paths {
				err := tree.Add(p, k.values[i])
				require.Equal(t, k.error, err)
			}
			require.Equal(t, k.want(), tree)
		})
	}
}
