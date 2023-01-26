# About

The plugin's implementation.

Having `internal` in its path prevents imports from outside the project.

Code may move into a public package at the root of the repository if there is a use case.

## Example

Given:

```protobuf
syntax = "proto3";
package my.service.v1;
import "firestore/options.proto";
option go_package = "github.com/myorg/apis-go/pkg/my/service/v1;servicev1";

message Account {
  option (firestore.message).collection = "accounts";
  string name = 1;
}

message User {
  option (firestore.message).collection = "accounts/{id}/users";
  string name = 1;
}
```

We generate a convenient API to read/write your protos from/to Firestore:

```go
package main

import (
	"cloud.google.com/go/firestore"
	servicev1 "github.com/myorg/apis-go/pkg/my/service/v1"
)

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "project")
	if err != nil {
		panic(err)
	}

	account := &servicev1.Account{
		Name: "myaccount",
	}

	err = servicev1.Firestore(client).
		Accounts().
		Doc("myid").
		Set(ctx, account)

	if err != nil {
		panic(err)
	}
}

```
