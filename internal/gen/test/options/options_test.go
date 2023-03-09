package options_test

import (
	// Assert compilation.

	"testing"

	"github.com/complex64/protoc-gen-go-firestore/firestorepb"
	"github.com/complex64/protoc-gen-go-firestore/v3/internal/gen/test/options"
	"github.com/complex64/protoc-gen-go-firestore/v3/internal/require"
)

// Test that all options in the .proto file are present as expected.

func TestFileOptions(t *testing.T) {
	msg := &options.Message{}
	defaults := &firestorepb.FileOptions{
		Enabled: false,
	}
	require.FileOptions(t, defaults, msg)
}

func TestMessageOptions(t *testing.T) {
	var (
		msg = &options.Message{}
	)
	t.Run("defaults", func(t *testing.T) {
		defaults := &firestorepb.MessageOptions{
			Enabled: false,
		}
		require.MessageOption(t, defaults, msg)
	})
}
