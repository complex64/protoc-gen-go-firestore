// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 1.2.0
// 	protoc          (unknown)

package converter

import (
	firestore "cloud.google.com/go/firestore"
)

type FS_converter struct {
	client *firestore.Client
}

func Firestore(client *firestore.Client) *FS_converter {
	return &FS_converter{
		client: client,
	}
}
