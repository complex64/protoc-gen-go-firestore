// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 2.1.1
// 	protoc          (unknown)
// source: utils/user.proto

package utils

import (
	json "encoding/json"
	_ "github.com/complex64/protoc-gen-go-firestore/firestorepb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// FirestoreCollectionUsers is the Firestore collection name for documents of type utils.User.
const FirestoreCollectionUsers = "users"

// FirestoreUser is the Firestore Custom Object for utils.User.
type FirestoreUser struct {
	Name string `firestore:"name,omitempty"`
}

func FirestoreMapToUser(m map[string]any) (*User, error) {
	bs, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	msg := new(User)
	if err := protojson.Unmarshal(bs, msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (x *User) ToFirestoreMap() (map[string]any, error) {
	bs, err := protojson.Marshal(x)
	if err != nil {
		return nil, err
	}
	m := map[string]any{}
	if err := json.Unmarshal(bs, &m); err != nil {
		return nil, err
	}
	return m, nil
}
