// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 1.2.0
// 	protoc          (unknown)

package customobject

import (
	firestore "cloud.google.com/go/firestore"
)

type FS_customobject struct {
	client *firestore.Client
}

func Firestore(client *firestore.Client) *FS_customobject {
	return &FS_customobject{
		client: client,
	}
}