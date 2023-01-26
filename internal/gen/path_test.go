package gen

import (
	"testing"

	"github.com/complex64/protoc-gen-go-firestore/firestorepb"
	"github.com/stretchr/testify/require"
)

func TestPath(t *testing.T) {
	// {
	// 	left := new(Path)
	// 	right := new(Path)
	// 	merged := left.Collection.Merge(right.Collection)
	// 	require.Nil(t, merged)
	// }
	// {
	// 	left := &Collection{}
	// 	var right *Collection = nil
	// 	require.Equal(t, left, left.Merge(right))
	// 	require.Equal(t, left, right.Merge(left))
	// }

	t.Run("simple path", func(t *testing.T) {
		msg := &Message{opts: &firestorepb.MessageOptions{Collection: "users"}}
		path, err := msg.parseCollection()
		require.NoError(t, err)
		require.NotNil(t, path)
		require.NotNil(t, path.Collection)
		require.Equal(t, "users", path.Collection.Segment)
		require.Equal(t, "Users", path.Collection.Title)
		require.Equal(t, msg, path.Collection.Message)
		require.Nil(t, path.Collection.Document)
		require.Nil(t, path.Collection.Parent)
	})

	t.Run("2-level path", func(t *testing.T) {
		msg := &Message{opts: &firestorepb.MessageOptions{Collection: "users/{id}/sessions"}}
		path, err := msg.parseCollection()
		require.NoError(t, err)
		require.NotNil(t, path)
		require.NotNil(t, path.Collection)
		require.Equal(t, "users", path.Collection.Segment)
		require.Equal(t, "Users", path.Collection.Title)
		require.Nil(t, path.Collection.Message)
		require.NotNil(t, path.Collection.Document)
		require.Nil(t, path.Collection.Parent)

		subcoll := path.Collection.Document.Collections
		require.Equal(t, 1, subcoll.Len())

		sess, ok := subcoll.Get("sessions")
		require.True(t, ok)
		require.Equal(t, "sessions", sess.Segment)
		require.Equal(t, "Sessions", sess.Title)
		require.Equal(t, msg, sess.Message)
		require.NotNil(t, sess.Parent)
		require.Nil(t, sess.Document)
	})

	t.Run("merge", func(t *testing.T) {
		users := &Message{opts: &firestorepb.MessageOptions{Collection: "users"}}
		up, err := users.parseCollection()
		require.NoError(t, err)
		require.NotNil(t, up)

		sessions := &Message{opts: &firestorepb.MessageOptions{Collection: "users/{id}/sessions"}}
		sp, err := sessions.parseCollection()
		require.NoError(t, err)
		require.NotNil(t, sp)

		leftMerged := up.Collection.Merge(sp.Collection)
		rightMerged := sp.Collection.Merge(up.Collection)

		require.Equal(t, users, leftMerged.Message)
		require.Equal(t, users, rightMerged.Message)

		require.Equal(t, "users", leftMerged.Segment)
		require.Equal(t, "users", rightMerged.Segment)

		require.Equal(t, "Users", leftMerged.Title)
		require.Equal(t, "Users", rightMerged.Title)

		require.NotNil(t, leftMerged.Document)
		require.NotNil(t, rightMerged.Document)

		leftSubColl := leftMerged.Document.Collections
		rightSubColl := rightMerged.Document.Collections

		require.Equal(t, 1, leftSubColl.Len())
		require.Equal(t, 1, rightSubColl.Len())
	})
}
