# About

This module declares and documents the options you can use in your `.proto` files.

# Usage

Include `options.proto` in your own `.proto` files and apply options to the file, your messages, or fields to control what `protoc-gen-go-firestore` generates.

```proto
import "firestore/options.proto";
```

Your editor and build setup needs to reference this file.

You have multiple options, depending what tools you use:

- Copy the file into your project, e.g. a `vendor` directory
- Use [Buf's Schema Registry (BSR)](https://docs.buf.build/bsr/introduction)

## Companion Go Module

Importing `options.proto` means your generated Go code imports the package `github.com/complex64/protoc-gen-go-firestore/firestorepb`. This package is distributed as [a small Go module](/firestorepb) with minimal dependencies.
