// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 2.0.1
// 	protoc          (unknown)

package converter

import (
	firestore "cloud.google.com/go/firestore"
)

type Firestore struct {
	client *firestore.Client
}

func WithFirestore(client *firestore.Client) *Firestore {
	return &Firestore{client: client}
}
