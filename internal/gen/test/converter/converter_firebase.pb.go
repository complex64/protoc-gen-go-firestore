// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 3.1.1
// 	protoc          (unknown)
// source: converter/converter.proto

package converter

import (
	json "encoding/json"
	_ "github.com/complex64/protoc-gen-go-firestore/firestorepb"
	protojson "google.golang.org/protobuf/encoding/protojson"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

func FirestoreMapToCity(m map[string]any) (*City, error) {
	bs, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	msg := new(City)
	if err := protojson.Unmarshal(bs, msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (x *City) ToFirestoreMap() (map[string]any, error) {
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
