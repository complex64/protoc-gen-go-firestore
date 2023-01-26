// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 1.0.2
// 	protoc          (unknown)

package utils

import (
	firestore "cloud.google.com/go/firestore"
	context "context"
)

type FS_utils struct {
	client *firestore.Client
}

func Firestore(client *firestore.Client) *FS_utils {
	return &FS_utils{
		client: client,
	}
}

func (x *FS_utils) Accounts() *FS_utils_Accounts {
	return &FS_utils_Accounts{
		c: x.client.Collection("accounts"),
	}
}

type FS_utils_Accounts struct {
	c *firestore.CollectionRef
}

type FS_utils_Accounts_Iter struct {
	i *firestore.DocumentIterator
}

type FS_utils_Accounts_Query struct {
	q firestore.Query
}

func (x *FS_utils_Accounts_Query) Documents(ctx context.Context) *FS_utils_Accounts_Iter {
	return &FS_utils_Accounts_Iter{
		i: x.q.Documents(ctx),
	}
}

func (x *FS_utils_Accounts_Query) Value() firestore.Query {
	return x.q
}

func (x *FS_utils_Accounts) Where(path, op string, value interface{}) *FS_utils_Accounts_Query {
	return &FS_utils_Accounts_Query{
		q: x.c.Where(path, op, value),
	}
}

func (x *FS_utils_Accounts_Query) Where(path, op string, value interface{}) *FS_utils_Accounts_Query {
	return &FS_utils_Accounts_Query{
		q: x.q.Where(path, op, value),
	}
}

func (x *FS_utils_Accounts) OrderBy(path string, dir firestore.Direction) *FS_utils_Accounts_Query {
	return &FS_utils_Accounts_Query{
		q: x.c.OrderBy(path, dir),
	}
}

func (x *FS_utils_Accounts_Query) OrderBy(path string, dir firestore.Direction) *FS_utils_Accounts_Query {
	return &FS_utils_Accounts_Query{
		q: x.q.OrderBy(path, dir),
	}
}

func (x *FS_utils_Accounts) Limit(n int) *FS_utils_Accounts_Query {
	return &FS_utils_Accounts_Query{
		q: x.c.Limit(n),
	}
}

func (x *FS_utils_Accounts_Query) Limit(n int) *FS_utils_Accounts_Query {
	return &FS_utils_Accounts_Query{
		q: x.q.Limit(n),
	}
}

func (x *FS_utils_Accounts_Iter) Stop() {
	x.i.Stop()
}

func (x *FS_utils_Accounts) Doc(id string) *FS_utils_Accounts_Doc {
	return &FS_utils_Accounts_Doc{
		d: x.c.Doc(id),
	}
}

type FS_utils_Accounts_Doc struct {
	d *firestore.DocumentRef
}

func (x *FS_utils_Accounts_Doc) Users() *FS_utils_Accounts_Users {
	return &FS_utils_Accounts_Users{
		c: x.d.Collection("users"),
	}
}

type FS_utils_Accounts_Users struct {
	c *firestore.CollectionRef
}

type FS_utils_Accounts_Users_Iter struct {
	i *firestore.DocumentIterator
}

type FS_utils_Accounts_Users_Query struct {
	q firestore.Query
}

func (x *FS_utils_Accounts_Users_Query) Documents(ctx context.Context) *FS_utils_Accounts_Users_Iter {
	return &FS_utils_Accounts_Users_Iter{
		i: x.q.Documents(ctx),
	}
}

func (x *FS_utils_Accounts_Users_Query) Value() firestore.Query {
	return x.q
}

func (x *FS_utils_Accounts_Users) Where(path, op string, value interface{}) *FS_utils_Accounts_Users_Query {
	return &FS_utils_Accounts_Users_Query{
		q: x.c.Where(path, op, value),
	}
}

func (x *FS_utils_Accounts_Users_Query) Where(path, op string, value interface{}) *FS_utils_Accounts_Users_Query {
	return &FS_utils_Accounts_Users_Query{
		q: x.q.Where(path, op, value),
	}
}

func (x *FS_utils_Accounts_Users) OrderBy(path string, dir firestore.Direction) *FS_utils_Accounts_Users_Query {
	return &FS_utils_Accounts_Users_Query{
		q: x.c.OrderBy(path, dir),
	}
}

func (x *FS_utils_Accounts_Users_Query) OrderBy(path string, dir firestore.Direction) *FS_utils_Accounts_Users_Query {
	return &FS_utils_Accounts_Users_Query{
		q: x.q.OrderBy(path, dir),
	}
}

func (x *FS_utils_Accounts_Users) Limit(n int) *FS_utils_Accounts_Users_Query {
	return &FS_utils_Accounts_Users_Query{
		q: x.c.Limit(n),
	}
}

func (x *FS_utils_Accounts_Users_Query) Limit(n int) *FS_utils_Accounts_Users_Query {
	return &FS_utils_Accounts_Users_Query{
		q: x.q.Limit(n),
	}
}

func (x *FS_utils_Accounts_Users_Iter) Stop() {
	x.i.Stop()
}

func (x *FS_utils_Accounts_Users) Doc(id string) *FS_utils_Accounts_Users_Doc {
	return &FS_utils_Accounts_Users_Doc{
		d: x.c.Doc(id),
	}
}

type FS_utils_Accounts_Users_Doc struct {
	d *firestore.DocumentRef
}

func (x *FS_utils_Accounts_Users_Doc) Sessions() *FS_utils_Accounts_Users_Sessions {
	return &FS_utils_Accounts_Users_Sessions{
		c: x.d.Collection(FirestoreCollectionSessions),
	}
}

type FS_utils_Accounts_Users_Sessions struct {
	c *firestore.CollectionRef
}

type FS_utils_Accounts_Users_Sessions_Iter struct {
	i *firestore.DocumentIterator
}

type FS_utils_Accounts_Users_Sessions_Query struct {
	q firestore.Query
}

func (x *FS_utils_Accounts_Users_Sessions_Query) Documents(ctx context.Context) *FS_utils_Accounts_Users_Sessions_Iter {
	return &FS_utils_Accounts_Users_Sessions_Iter{
		i: x.q.Documents(ctx),
	}
}

func (x *FS_utils_Accounts_Users_Sessions_Query) Value() firestore.Query {
	return x.q
}

func (x *FS_utils_Accounts_Users_Sessions) Where(path, op string, value interface{}) *FS_utils_Accounts_Users_Sessions_Query {
	return &FS_utils_Accounts_Users_Sessions_Query{
		q: x.c.Where(path, op, value),
	}
}

func (x *FS_utils_Accounts_Users_Sessions_Query) Where(path, op string, value interface{}) *FS_utils_Accounts_Users_Sessions_Query {
	return &FS_utils_Accounts_Users_Sessions_Query{
		q: x.q.Where(path, op, value),
	}
}

func (x *FS_utils_Accounts_Users_Sessions) OrderBy(path string, dir firestore.Direction) *FS_utils_Accounts_Users_Sessions_Query {
	return &FS_utils_Accounts_Users_Sessions_Query{
		q: x.c.OrderBy(path, dir),
	}
}

func (x *FS_utils_Accounts_Users_Sessions_Query) OrderBy(path string, dir firestore.Direction) *FS_utils_Accounts_Users_Sessions_Query {
	return &FS_utils_Accounts_Users_Sessions_Query{
		q: x.q.OrderBy(path, dir),
	}
}

func (x *FS_utils_Accounts_Users_Sessions) Limit(n int) *FS_utils_Accounts_Users_Sessions_Query {
	return &FS_utils_Accounts_Users_Sessions_Query{
		q: x.c.Limit(n),
	}
}

func (x *FS_utils_Accounts_Users_Sessions_Query) Limit(n int) *FS_utils_Accounts_Users_Sessions_Query {
	return &FS_utils_Accounts_Users_Sessions_Query{
		q: x.q.Limit(n),
	}
}

func (x *FS_utils_Accounts_Users_Sessions_Query) First(ctx context.Context) (*Session, error) {
	iter := x.q.Limit(1).Documents(ctx)
	defer iter.Stop()
	snap, err := iter.Next()
	if err != nil {
		return nil, err
	}
	o := new(FirestoreSession)
	if err := snap.DataTo(o); err != nil {
		return nil, err
	}
	if p, err := o.ToProto(); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

func (x *FS_utils_Accounts_Users_Sessions_Iter) GetAll() ([]*Session, error) {
	snaps, err := x.i.GetAll()
	if err != nil {
		return nil, err
	}
	protos := make([]*Session, len(snaps))
	for i, snap := range snaps {
		o := new(FirestoreSession)
		if err := snap.DataTo(o); err != nil {
			return nil, err
		}
		if p, err := o.ToProto(); err != nil {
			return nil, err
		} else {
			protos[i] = p
		}
	}
	return protos, nil
}

func (x *FS_utils_Accounts_Users_Sessions_Iter) GetAllAsSnapshots() ([]*firestore.DocumentSnapshot, error) {
	return x.i.GetAll()
}

func (x *FS_utils_Accounts_Users_Sessions_Iter) Next() (*Session, error) {
	snap, err := x.i.Next()
	if err != nil {
		return nil, err
	}
	o := new(FirestoreSession)
	if err := snap.DataTo(o); err != nil {
		return nil, err
	}
	if p, err := o.ToProto(); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

func (x *FS_utils_Accounts_Users_Sessions_Iter) NextAsSnapshot() (*firestore.DocumentSnapshot, error) {
	return x.i.Next()
}

func (x *FS_utils_Accounts_Users_Sessions_Iter) Stop() {
	x.i.Stop()
}

func (x *FS_utils_Accounts_Users_Sessions) Doc(id string) *FS_utils_Accounts_Users_Sessions_Doc {
	return &FS_utils_Accounts_Users_Sessions_Doc{
		d: x.c.Doc(id),
	}
}

type FS_utils_Accounts_Users_Sessions_Doc struct {
	d *firestore.DocumentRef
}

func (x *FS_utils_Accounts_Users_Sessions_Doc) Set(ctx context.Context, m *Session) error {
	fs, err := m.ToFirestore()
	if err != nil {
		return err
	}
	if _, err := x.d.Set(ctx, fs); err != nil {
		return err
	}
	return nil
}

func (x *FS_utils_Accounts_Users_Sessions_Doc) Delete(ctx context.Context, preconds ...firestore.Precondition) (*firestore.WriteResult, error) {
	return x.d.Delete(ctx, preconds...)
}

func (x *FS_utils_Accounts_Users_Sessions_Doc) Ref() *firestore.DocumentRef {
	return x.d
}
