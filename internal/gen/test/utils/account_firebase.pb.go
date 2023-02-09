// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 1.2.0
// 	protoc          (unknown)
// source: utils/account.proto

package utils

import (
	_ "github.com/complex64/protoc-gen-go-firestore/firestorepb"
)

// FirestoreCollectionAccounts is the Firestore collection name for documents of type utils.Account.
const FirestoreCollectionAccounts = "accounts"

// FirestoreAccount is the Firestore Custom Object for utils.Account.
type FirestoreAccount struct {
	Name string `firestore:"name,omitempty"`
}

// ToProto converts this FirestoreAccount to its protobuf representation.
func (m *FirestoreAccount) ToProto() (*Account, error) {
	x := new(Account)
	x.Name = m.Name
	return x, nil
}

// ToFirestore returns the Firestore Custom Object for Account.
func (x *Account) ToFirestore() (*FirestoreAccount, error) {
	m := new(FirestoreAccount)
	m.Name = x.Name
	return m, nil
}