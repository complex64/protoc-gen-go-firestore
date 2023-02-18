// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 2.0.1
// 	protoc          (unknown)

package unlinked

import (
	firestore "cloud.google.com/go/firestore"
	context "context"
	errors "errors"
	iterator "google.golang.org/api/iterator"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type Firestore struct {
	client *firestore.Client
}

func WithFirestore(client *firestore.Client) *Firestore {
	return &Firestore{client: client}
}

func (fs *Firestore) Parents() *FirestoreParentsCollectionRef {
	return &FirestoreParentsCollectionRef{
		coll: fs.client.Collection("parents"),
	}
}
func (ref *FirestoreParentsCollectionRef) Limit(n int) *FirestoreParentsQuery {
	return &FirestoreParentsQuery{
		query: ref.coll.Limit(n),
	}
}

func (q *FirestoreParentsQuery) Limit(n int) *FirestoreParentsQuery {
	return &FirestoreParentsQuery{
		query: q.query.Limit(n),
	}
}

func (ref *FirestoreParentsCollectionRef) OrderBy(path string, dir firestore.Direction) *FirestoreParentsQuery {
	return &FirestoreParentsQuery{
		query: ref.coll.OrderBy(path, dir),
	}
}

func (q *FirestoreParentsQuery) OrderBy(path string, dir firestore.Direction) *FirestoreParentsQuery {
	return &FirestoreParentsQuery{
		query: q.query.OrderBy(path, dir),
	}
}

func (ref *FirestoreParentsCollectionRef) Where(path, op string, value interface{}) *FirestoreParentsQuery {
	return &FirestoreParentsQuery{
		query: ref.coll.Where(path, op, value),
	}
}

func (q *FirestoreParentsQuery) Where(path, op string, value interface{}) *FirestoreParentsQuery {
	return &FirestoreParentsQuery{
		query: q.query.Where(path, op, value),
	}
}

// FirestoreParentsCollectionRef holds a reference to the Firestore collection `parents`.
type FirestoreParentsCollectionRef struct {
	coll *firestore.CollectionRef
}

func (ref *FirestoreParentsCollectionRef) Doc(id string) *FirestoreParentsDocumentRef {
	return &FirestoreParentsDocumentRef{
		doc: ref.coll.Doc(id),
	}
}

// FirestoreParentsDocumentRef holds a reference to a Firestore document in collection `parents`.
type FirestoreParentsDocumentRef struct {
	doc *firestore.DocumentRef
}

func (i *FirestoreParentsIterator) Stop() {
	i.iter.Stop()
}

type FirestoreParentsIterator struct {
	iter *firestore.DocumentIterator
}

type FirestoreParentsQuery struct {
	query firestore.Query
}

func (q *FirestoreParentsQuery) Documents(ctx context.Context) *FirestoreParentsIterator {
	return &FirestoreParentsIterator{
		iter: q.query.Documents(ctx),
	}
}

func (q *FirestoreParentsQuery) Value() firestore.Query {
	return q.query
}

func (ref *FirestoreParentsDocumentRef) Subparents() *FirestoreSubparentsCollectionRef {
	return &FirestoreSubparentsCollectionRef{
		coll: ref.doc.Collection("subparents"),
	}
}
func (ref *FirestoreSubparentsCollectionRef) Limit(n int) *FirestoreSubparentsQuery {
	return &FirestoreSubparentsQuery{
		query: ref.coll.Limit(n),
	}
}

func (q *FirestoreSubparentsQuery) Limit(n int) *FirestoreSubparentsQuery {
	return &FirestoreSubparentsQuery{
		query: q.query.Limit(n),
	}
}

func (ref *FirestoreSubparentsCollectionRef) OrderBy(path string, dir firestore.Direction) *FirestoreSubparentsQuery {
	return &FirestoreSubparentsQuery{
		query: ref.coll.OrderBy(path, dir),
	}
}

func (q *FirestoreSubparentsQuery) OrderBy(path string, dir firestore.Direction) *FirestoreSubparentsQuery {
	return &FirestoreSubparentsQuery{
		query: q.query.OrderBy(path, dir),
	}
}

func (ref *FirestoreSubparentsCollectionRef) Where(path, op string, value interface{}) *FirestoreSubparentsQuery {
	return &FirestoreSubparentsQuery{
		query: ref.coll.Where(path, op, value),
	}
}

func (q *FirestoreSubparentsQuery) Where(path, op string, value interface{}) *FirestoreSubparentsQuery {
	return &FirestoreSubparentsQuery{
		query: q.query.Where(path, op, value),
	}
}

// FirestoreSubparentsCollectionRef holds a reference to the Firestore collection `subparents`.
type FirestoreSubparentsCollectionRef struct {
	coll *firestore.CollectionRef
}

func (ref *FirestoreSubparentsCollectionRef) Doc(id string) *FirestoreSubparentsDocumentRef {
	return &FirestoreSubparentsDocumentRef{
		doc: ref.coll.Doc(id),
	}
}

// FirestoreSubparentsDocumentRef holds a reference to a Firestore document in collection `subparents`.
type FirestoreSubparentsDocumentRef struct {
	doc *firestore.DocumentRef
}

func (i *FirestoreSubparentsIterator) Stop() {
	i.iter.Stop()
}

type FirestoreSubparentsIterator struct {
	iter *firestore.DocumentIterator
}

type FirestoreSubparentsQuery struct {
	query firestore.Query
}

func (q *FirestoreSubparentsQuery) Documents(ctx context.Context) *FirestoreSubparentsIterator {
	return &FirestoreSubparentsIterator{
		iter: q.query.Documents(ctx),
	}
}

func (q *FirestoreSubparentsQuery) Value() firestore.Query {
	return q.query
}

func (ref *FirestoreSubparentsDocumentRef) Items() *FirestoreItemsCollectionRef {
	return &FirestoreItemsCollectionRef{
		coll: ref.doc.Collection(FirestoreCollectionItems),
	}
}
func (ref *FirestoreItemsCollectionRef) Create(ctx context.Context, p *Item) (*firestore.WriteResult, error) {
	fs, err := p.ToFirestore()
	if err != nil {
		return nil, err
	}
	id := fs.Name
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "empty id")
	}
	res, err := ref.coll.Doc(id).Create(ctx, fs)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ref *FirestoreItemsCollectionRef) Limit(n int) *FirestoreItemsQuery {
	return &FirestoreItemsQuery{
		query: ref.coll.Limit(n),
	}
}

func (q *FirestoreItemsQuery) Limit(n int) *FirestoreItemsQuery {
	return &FirestoreItemsQuery{
		query: q.query.Limit(n),
	}
}

func (ref *FirestoreItemsCollectionRef) OrderBy(path string, dir firestore.Direction) *FirestoreItemsQuery {
	return &FirestoreItemsQuery{
		query: ref.coll.OrderBy(path, dir),
	}
}

func (q *FirestoreItemsQuery) OrderBy(path string, dir firestore.Direction) *FirestoreItemsQuery {
	return &FirestoreItemsQuery{
		query: q.query.OrderBy(path, dir),
	}
}

func (ref *FirestoreItemsCollectionRef) Where(path, op string, value interface{}) *FirestoreItemsQuery {
	return &FirestoreItemsQuery{
		query: ref.coll.Where(path, op, value),
	}
}

func (q *FirestoreItemsQuery) Where(path, op string, value interface{}) *FirestoreItemsQuery {
	return &FirestoreItemsQuery{
		query: q.query.Where(path, op, value),
	}
}

// FirestoreItemsCollectionRef holds a reference to the Firestore collection `items`.
type FirestoreItemsCollectionRef struct {
	coll *firestore.CollectionRef
}

func (ref *FirestoreItemsCollectionRef) Doc(id string) *FirestoreItemsDocumentRef {
	return &FirestoreItemsDocumentRef{
		doc: ref.coll.Doc(id),
	}
}

func (ref *FirestoreItemsDocumentRef) Delete(ctx context.Context, preconds ...firestore.Precondition) (*firestore.WriteResult, error) {
	return ref.doc.Delete(ctx, preconds...)
}

func (ref *FirestoreItemsDocumentRef) Get(ctx context.Context) (*Item, error) {
	snapshot, err := ref.doc.Get(ctx)
	if err != nil {
		return nil, err
	}
	obj := new(FirestoreItem)
	if err := snapshot.DataTo(obj); err != nil {
		return nil, err
	}
	if proto, err := obj.ToProto(); err != nil {
		return nil, err
	} else {
		return proto, nil
	}
}

func (ref *FirestoreItemsDocumentRef) Ref() *firestore.DocumentRef {
	return ref.doc
}

func (ref *FirestoreItemsDocumentRef) Set(ctx context.Context, msg *Item, opts ...firestore.SetOption) error {
	fs, err := msg.ToFirestore()
	if err != nil {
		return err
	}
	if _, err := ref.doc.Set(ctx, fs, opts...); err != nil {
		return err
	}
	return nil
}

// FirestoreItemsDocumentRef holds a reference to a Firestore document in collection `items`.
type FirestoreItemsDocumentRef struct {
	doc *firestore.DocumentRef
}

func (i *FirestoreItemsIterator) GetAll() ([]*Item, error) {
	snaps, err := i.iter.GetAll()
	if err != nil {
		return nil, err
	}
	protos := make([]*Item, len(snaps))
	for j, snapshot := range snaps {
		o := new(FirestoreItem)
		if err := snapshot.DataTo(o); err != nil {
			return nil, err
		}
		if p, err := o.ToProto(); err != nil {
			return nil, err
		} else {
			protos[j] = p
		}
	}
	return protos, nil
}

func (i *FirestoreItemsIterator) GetAllAsSnapshots() ([]*firestore.DocumentSnapshot, error) {
	return i.iter.GetAll()
}

func (i *FirestoreItemsIterator) Next() (*Item, error) {
	snapshot, err := i.iter.Next()
	if err != nil {
		return nil, err
	}
	obj := new(FirestoreItem)
	if err := snapshot.DataTo(obj); err != nil {
		return nil, err
	}
	if p, err := obj.ToProto(); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

func (i *FirestoreItemsIterator) NextAsSnapshot() (*firestore.DocumentSnapshot, error) {
	return i.iter.Next()
}

func (i *FirestoreItemsIterator) Stop() {
	i.iter.Stop()
}

type FirestoreItemsIterator struct {
	iter *firestore.DocumentIterator
}

func (q *FirestoreItemsQuery) First(ctx context.Context) (*Item, error) {
	iter := q.query.Limit(1).Documents(ctx)
	defer iter.Stop()
	snapshot, err := iter.Next()
	if err != nil {
		if errors.Is(err, iterator.Done) {
			return nil, nil
		}
		return nil, err
	}
	obj := new(FirestoreItem)
	if err := snapshot.DataTo(obj); err != nil {
		return nil, err
	}
	if proto, err := obj.ToProto(); err != nil {
		return nil, err
	} else {
		return proto, nil
	}
}

type FirestoreItemsQuery struct {
	query firestore.Query
}

func (q *FirestoreItemsQuery) Documents(ctx context.Context) *FirestoreItemsIterator {
	return &FirestoreItemsIterator{
		iter: q.query.Documents(ctx),
	}
}

func (q *FirestoreItemsQuery) Value() firestore.Query {
	return q.query
}
