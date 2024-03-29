// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 3.1.1
// 	protoc          (unknown)
// source: utils/account.proto

package utils

import (
	json "encoding/json"
	_ "github.com/complex64/protoc-gen-go-firestore/firestorepb"
	protojson "google.golang.org/protobuf/encoding/protojson"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// FirestoreCollectionAccounts is the Firestore collection name for documents of type utils.Account.
const FirestoreCollectionAccounts = "accounts"

func FirestoreMapToAccount(m map[string]any) (*Account, error) {
	bs, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	msg := new(Account)
	if err := protojson.Unmarshal(bs, msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (x *Account) ToFirestoreMap() (map[string]any, error) {
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

// FirestoreCollectionManifests is the Firestore collection name for documents of type utils.Manifest.
const FirestoreCollectionManifests = "manifests"

func FirestoreMapToManifest(m map[string]any) (*Manifest, error) {
	bs, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	msg := new(Manifest)
	if err := protojson.Unmarshal(bs, msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (x *Manifest) ToFirestoreMap() (map[string]any, error) {
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

// FirestoreCollectionActions is the Firestore collection name for documents of type utils.Action.
const FirestoreCollectionActions = "actions"

func FirestoreMapToAction(m map[string]any) (*Action, error) {
	bs, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	msg := new(Action)
	if err := protojson.Unmarshal(bs, msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (x *Action) ToFirestoreMap() (map[string]any, error) {
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
