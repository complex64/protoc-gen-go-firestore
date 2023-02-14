package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTree_Walk(t *testing.T) {
	tree := new(Tree[string])
	require.NoError(t, tree.Add("installations/{owner}/repositories/{repository}/manifests/{manifest}", "manifests"))
	require.NoError(t, tree.Add("installations/{owner}/repositories/{repository}/manifests/{manifest}/actions/{action}", "actions"))
	require.NoError(t, tree.Add("installations/{owner}/tags/{tag}", "tags"))

	tree.Walk(func(parent *Parent[string], name string, value string) {
		fmt.Printf("parent = %+v\n", parent)
		fmt.Printf("name = %+v\n", name)
		fmt.Printf("value = %+v\n\n", value)
	})
}
