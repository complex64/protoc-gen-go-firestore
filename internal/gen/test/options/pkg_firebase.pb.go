// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 1.2.0
// 	protoc          (unknown)

package options

import (
	firestore "cloud.google.com/go/firestore"
)

type FS_options struct {
	client *firestore.Client
}

func Firestore(client *firestore.Client) *FS_options {
	return &FS_options{
		client: client,
	}
}
