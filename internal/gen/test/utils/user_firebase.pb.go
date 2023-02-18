// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 2.0.1
// 	protoc          (unknown)
// source: utils/user.proto

package utils

import (
	_ "github.com/complex64/protoc-gen-go-firestore/firestorepb"
)

// FirestoreCollectionUsers is the Firestore collection name for documents of type utils.User.
const FirestoreCollectionUsers = "users"

// FirestoreUser is the Firestore Custom Object for utils.User.
type FirestoreUser struct {
	Name string `firestore:"name,omitempty"`
}

// ToProto converts this FirestoreUser to its protobuf representation.
func (m *FirestoreUser) ToProto() (*User, error) {
	x := new(User)
	x.Name = m.Name
	return x, nil
}

// ToFirestore returns the Firestore Custom Object for User.
func (x *User) ToFirestore() (*FirestoreUser, error) {
	m := new(FirestoreUser)
	m.Name = x.Name
	return m, nil
}
