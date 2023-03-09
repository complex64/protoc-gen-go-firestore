// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 2.1.1
// 	protoc          (unknown)
// source: utils/session.proto

package utils

import (
	json "encoding/json"
	_ "github.com/complex64/protoc-gen-go-firestore/firestorepb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// FirestoreCollectionSessions is the Firestore collection name for documents of type utils.Session.
const FirestoreCollectionSessions = "sessions"

func FirestoreMapToSession(m map[string]any) (*Session, error) {
	bs, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	msg := new(Session)
	if err := protojson.Unmarshal(bs, msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (x *Session) ToFirestoreMap() (map[string]any, error) {
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
