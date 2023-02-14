# protoc-gen-go-firestore

[![Tests](https://github.com/complex64/protoc-gen-go-firestore/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/complex64/protoc-gen-go-firestore/actions/workflows/tests.yml) [![Linters](https://github.com/complex64/protoc-gen-go-firestore/actions/workflows/linters.yml/badge.svg?branch=main)](https://github.com/complex64/protoc-gen-go-firestore/actions/workflows/linters.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/complex64/protoc-gen-go-firestore)](https://goreportcard.com/report/github.com/complex64/protoc-gen-go-firestore) [![Maintainability](https://api.codeclimate.com/v1/badges/69739915a43041e34892/maintainability)](https://codeclimate.com/github/complex64/protoc-gen-go-firestore/maintainability) [![Go Reference](https://pkg.go.dev/badge/github.com/complex64/protoc-gen-go-firestore.svg)](https://pkg.go.dev/github.com/complex64/protoc-gen-go-firestore)

Exploration into generating Firestore bindings for Go from your .proto files.

## Example

We generate a convenient API to read/write your protos from/to Firestore:

```go
package main

import (
	"context"

	"cloud.google.com/go/firestore"
	servicev1 "github.com/myorg/apis-go/pkg/my/service/v1"
)

func main() {
	servicev1.Firestore(client()).
		Accounts().
		Doc("myid").
		Set(context.Background(), &servicev1.Account{
			Name: "myaccount",
		})
}

func client() *firestore.Client {
	c, err := firestore.NewClient(context.Background(), "project")
	if err != nil {
		panic(err)
	}
	return c
}

```

Given:

```protobuf
syntax = "proto3";
package my.service.v1;

import "firestore/options.proto";

option go_package = "github.com/myorg/apis-go/pkg/my/service/v1;servicev1";

message Account {
  option (firestore.message).collection = "accounts/{id}";
  string name = 1;
}

message User {
  option (firestore.message).collection = "accounts/{id}/users/{id}";
  string name = 1;
}
```
