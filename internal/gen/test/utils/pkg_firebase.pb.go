// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 3.1.1
// 	protoc          (unknown)

package utils

import (
	firestore "cloud.google.com/go/firestore"
	context "context"
	errors "errors"
	iterator "google.golang.org/api/iterator"
)

type Firestore struct {
	client *firestore.Client
}

func WithFirestore(client *firestore.Client) *Firestore {
	return &Firestore{client: client}
}

func (fs *Firestore) Accounts() *FirestoreAccountsCollectionRef {
	return &FirestoreAccountsCollectionRef{
		coll: fs.client.Collection(FirestoreCollectionAccounts),
	}
}
func (ref *FirestoreAccountsCollectionRef) Limit(n int) *FirestoreAccountsQuery {
	return &FirestoreAccountsQuery{
		query: ref.coll.Limit(n),
	}
}

func (q *FirestoreAccountsQuery) Limit(n int) *FirestoreAccountsQuery {
	return &FirestoreAccountsQuery{
		query: q.query.Limit(n),
	}
}

func (ref *FirestoreAccountsCollectionRef) OrderBy(path string, dir firestore.Direction) *FirestoreAccountsQuery {
	return &FirestoreAccountsQuery{
		query: ref.coll.OrderBy(path, dir),
	}
}

func (q *FirestoreAccountsQuery) OrderBy(path string, dir firestore.Direction) *FirestoreAccountsQuery {
	return &FirestoreAccountsQuery{
		query: q.query.OrderBy(path, dir),
	}
}

func (ref *FirestoreAccountsCollectionRef) Where(path, op string, value interface{}) *FirestoreAccountsQuery {
	return &FirestoreAccountsQuery{
		query: ref.coll.Where(path, op, value),
	}
}

func (q *FirestoreAccountsQuery) Where(path, op string, value interface{}) *FirestoreAccountsQuery {
	return &FirestoreAccountsQuery{
		query: q.query.Where(path, op, value),
	}
}

// FirestoreAccountsCollectionRef holds a reference to the Firestore collection `accounts`.
type FirestoreAccountsCollectionRef struct {
	coll *firestore.CollectionRef
}

func (ref *FirestoreAccountsCollectionRef) Doc(id string) *FirestoreAccountsDocumentRef {
	return &FirestoreAccountsDocumentRef{
		doc: ref.coll.Doc(id),
	}
}

func (ref *FirestoreAccountsDocumentRef) Delete(ctx context.Context, preconds ...firestore.Precondition) (*firestore.WriteResult, error) {
	return ref.doc.Delete(ctx, preconds...)
}

func (ref *FirestoreAccountsDocumentRef) Get(ctx context.Context) (*Account, error) {
	snapshot, err := ref.doc.Get(ctx)
	if err != nil {
		return nil, err
	}
	m := snapshot.Data()
	if proto, err := FirestoreMapToAccount(m); err != nil {
		return nil, err
	} else {
		return proto, nil
	}
}

func (ref *FirestoreAccountsDocumentRef) Ref() *firestore.DocumentRef {
	return ref.doc
}

func (ref *FirestoreAccountsDocumentRef) Set(ctx context.Context, msg *Account, opts ...firestore.SetOption) error {
	fs, err := msg.ToFirestoreMap()
	if err != nil {
		return err
	}
	if _, err := ref.doc.Set(ctx, fs, opts...); err != nil {
		return err
	}
	return nil
}

// FirestoreAccountsDocumentRef holds a reference to a Firestore document in collection `accounts`.
type FirestoreAccountsDocumentRef struct {
	doc *firestore.DocumentRef
}

func (i *FirestoreAccountsIterator) GetAll() ([]*Account, error) {
	snaps, err := i.iter.GetAll()
	if err != nil {
		return nil, err
	}
	protos := make([]*Account, len(snaps))
	for j, snapshot := range snaps {
		m := snapshot.Data()
		if proto, err := FirestoreMapToAccount(m); err != nil {
			return nil, err
		} else {
			protos[j] = proto
		}
	}
	return protos, nil
}

func (i *FirestoreAccountsIterator) GetAllAsSnapshots() ([]*firestore.DocumentSnapshot, error) {
	return i.iter.GetAll()
}

func (i *FirestoreAccountsIterator) Next() (*Account, error) {
	snapshot, err := i.iter.Next()
	if err != nil {
		return nil, err
	}
	m := snapshot.Data()
	if p, err := FirestoreMapToAccount(m); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

func (i *FirestoreAccountsIterator) NextAsSnapshot() (*firestore.DocumentSnapshot, error) {
	return i.iter.Next()
}

func (i *FirestoreAccountsIterator) Stop() {
	i.iter.Stop()
}

type FirestoreAccountsIterator struct {
	iter *firestore.DocumentIterator
}

func (q *FirestoreAccountsQuery) First(ctx context.Context) (*Account, error) {
	iter := q.query.Limit(1).Documents(ctx)
	defer iter.Stop()
	snapshot, err := iter.Next()
	if err != nil {
		if errors.Is(err, iterator.Done) {
			return nil, nil
		}
		return nil, err
	}
	m := snapshot.Data()
	if proto, err := FirestoreMapToAccount(m); err != nil {
		return nil, err
	} else {
		return proto, nil
	}
}

type FirestoreAccountsQuery struct {
	query firestore.Query
}

func (q *FirestoreAccountsQuery) Documents(ctx context.Context) *FirestoreAccountsIterator {
	return &FirestoreAccountsIterator{
		iter: q.query.Documents(ctx),
	}
}

func (q *FirestoreAccountsQuery) Value() firestore.Query {
	return q.query
}

func (ref *FirestoreAccountsDocumentRef) Users() *FirestoreUsersCollectionRef {
	return &FirestoreUsersCollectionRef{
		coll: ref.doc.Collection(FirestoreCollectionUsers),
	}
}
func (ref *FirestoreUsersCollectionRef) Limit(n int) *FirestoreUsersQuery {
	return &FirestoreUsersQuery{
		query: ref.coll.Limit(n),
	}
}

func (q *FirestoreUsersQuery) Limit(n int) *FirestoreUsersQuery {
	return &FirestoreUsersQuery{
		query: q.query.Limit(n),
	}
}

func (ref *FirestoreUsersCollectionRef) OrderBy(path string, dir firestore.Direction) *FirestoreUsersQuery {
	return &FirestoreUsersQuery{
		query: ref.coll.OrderBy(path, dir),
	}
}

func (q *FirestoreUsersQuery) OrderBy(path string, dir firestore.Direction) *FirestoreUsersQuery {
	return &FirestoreUsersQuery{
		query: q.query.OrderBy(path, dir),
	}
}

func (ref *FirestoreUsersCollectionRef) Where(path, op string, value interface{}) *FirestoreUsersQuery {
	return &FirestoreUsersQuery{
		query: ref.coll.Where(path, op, value),
	}
}

func (q *FirestoreUsersQuery) Where(path, op string, value interface{}) *FirestoreUsersQuery {
	return &FirestoreUsersQuery{
		query: q.query.Where(path, op, value),
	}
}

// FirestoreUsersCollectionRef holds a reference to the Firestore collection `users`.
type FirestoreUsersCollectionRef struct {
	coll *firestore.CollectionRef
}

func (ref *FirestoreUsersCollectionRef) Doc(id string) *FirestoreUsersDocumentRef {
	return &FirestoreUsersDocumentRef{
		doc: ref.coll.Doc(id),
	}
}

func (ref *FirestoreUsersDocumentRef) Delete(ctx context.Context, preconds ...firestore.Precondition) (*firestore.WriteResult, error) {
	return ref.doc.Delete(ctx, preconds...)
}

func (ref *FirestoreUsersDocumentRef) Get(ctx context.Context) (*User, error) {
	snapshot, err := ref.doc.Get(ctx)
	if err != nil {
		return nil, err
	}
	m := snapshot.Data()
	if proto, err := FirestoreMapToUser(m); err != nil {
		return nil, err
	} else {
		return proto, nil
	}
}

func (ref *FirestoreUsersDocumentRef) Ref() *firestore.DocumentRef {
	return ref.doc
}

func (ref *FirestoreUsersDocumentRef) Set(ctx context.Context, msg *User, opts ...firestore.SetOption) error {
	fs, err := msg.ToFirestoreMap()
	if err != nil {
		return err
	}
	if _, err := ref.doc.Set(ctx, fs, opts...); err != nil {
		return err
	}
	return nil
}

// FirestoreUsersDocumentRef holds a reference to a Firestore document in collection `users`.
type FirestoreUsersDocumentRef struct {
	doc *firestore.DocumentRef
}

func (i *FirestoreUsersIterator) GetAll() ([]*User, error) {
	snaps, err := i.iter.GetAll()
	if err != nil {
		return nil, err
	}
	protos := make([]*User, len(snaps))
	for j, snapshot := range snaps {
		m := snapshot.Data()
		if proto, err := FirestoreMapToUser(m); err != nil {
			return nil, err
		} else {
			protos[j] = proto
		}
	}
	return protos, nil
}

func (i *FirestoreUsersIterator) GetAllAsSnapshots() ([]*firestore.DocumentSnapshot, error) {
	return i.iter.GetAll()
}

func (i *FirestoreUsersIterator) Next() (*User, error) {
	snapshot, err := i.iter.Next()
	if err != nil {
		return nil, err
	}
	m := snapshot.Data()
	if p, err := FirestoreMapToUser(m); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

func (i *FirestoreUsersIterator) NextAsSnapshot() (*firestore.DocumentSnapshot, error) {
	return i.iter.Next()
}

func (i *FirestoreUsersIterator) Stop() {
	i.iter.Stop()
}

type FirestoreUsersIterator struct {
	iter *firestore.DocumentIterator
}

func (q *FirestoreUsersQuery) First(ctx context.Context) (*User, error) {
	iter := q.query.Limit(1).Documents(ctx)
	defer iter.Stop()
	snapshot, err := iter.Next()
	if err != nil {
		if errors.Is(err, iterator.Done) {
			return nil, nil
		}
		return nil, err
	}
	m := snapshot.Data()
	if proto, err := FirestoreMapToUser(m); err != nil {
		return nil, err
	} else {
		return proto, nil
	}
}

type FirestoreUsersQuery struct {
	query firestore.Query
}

func (q *FirestoreUsersQuery) Documents(ctx context.Context) *FirestoreUsersIterator {
	return &FirestoreUsersIterator{
		iter: q.query.Documents(ctx),
	}
}

func (q *FirestoreUsersQuery) Value() firestore.Query {
	return q.query
}

func (ref *FirestoreUsersDocumentRef) Sessions() *FirestoreSessionsCollectionRef {
	return &FirestoreSessionsCollectionRef{
		coll: ref.doc.Collection(FirestoreCollectionSessions),
	}
}
func (ref *FirestoreSessionsCollectionRef) Limit(n int) *FirestoreSessionsQuery {
	return &FirestoreSessionsQuery{
		query: ref.coll.Limit(n),
	}
}

func (q *FirestoreSessionsQuery) Limit(n int) *FirestoreSessionsQuery {
	return &FirestoreSessionsQuery{
		query: q.query.Limit(n),
	}
}

func (ref *FirestoreSessionsCollectionRef) OrderBy(path string, dir firestore.Direction) *FirestoreSessionsQuery {
	return &FirestoreSessionsQuery{
		query: ref.coll.OrderBy(path, dir),
	}
}

func (q *FirestoreSessionsQuery) OrderBy(path string, dir firestore.Direction) *FirestoreSessionsQuery {
	return &FirestoreSessionsQuery{
		query: q.query.OrderBy(path, dir),
	}
}

func (ref *FirestoreSessionsCollectionRef) Where(path, op string, value interface{}) *FirestoreSessionsQuery {
	return &FirestoreSessionsQuery{
		query: ref.coll.Where(path, op, value),
	}
}

func (q *FirestoreSessionsQuery) Where(path, op string, value interface{}) *FirestoreSessionsQuery {
	return &FirestoreSessionsQuery{
		query: q.query.Where(path, op, value),
	}
}

// FirestoreSessionsCollectionRef holds a reference to the Firestore collection `sessions`.
type FirestoreSessionsCollectionRef struct {
	coll *firestore.CollectionRef
}

func (ref *FirestoreSessionsCollectionRef) Doc(id string) *FirestoreSessionsDocumentRef {
	return &FirestoreSessionsDocumentRef{
		doc: ref.coll.Doc(id),
	}
}

func (ref *FirestoreSessionsDocumentRef) Delete(ctx context.Context, preconds ...firestore.Precondition) (*firestore.WriteResult, error) {
	return ref.doc.Delete(ctx, preconds...)
}

func (ref *FirestoreSessionsDocumentRef) Get(ctx context.Context) (*Session, error) {
	snapshot, err := ref.doc.Get(ctx)
	if err != nil {
		return nil, err
	}
	m := snapshot.Data()
	if proto, err := FirestoreMapToSession(m); err != nil {
		return nil, err
	} else {
		return proto, nil
	}
}

func (ref *FirestoreSessionsDocumentRef) Ref() *firestore.DocumentRef {
	return ref.doc
}

func (ref *FirestoreSessionsDocumentRef) Set(ctx context.Context, msg *Session, opts ...firestore.SetOption) error {
	fs, err := msg.ToFirestoreMap()
	if err != nil {
		return err
	}
	if _, err := ref.doc.Set(ctx, fs, opts...); err != nil {
		return err
	}
	return nil
}

// FirestoreSessionsDocumentRef holds a reference to a Firestore document in collection `sessions`.
type FirestoreSessionsDocumentRef struct {
	doc *firestore.DocumentRef
}

func (i *FirestoreSessionsIterator) GetAll() ([]*Session, error) {
	snaps, err := i.iter.GetAll()
	if err != nil {
		return nil, err
	}
	protos := make([]*Session, len(snaps))
	for j, snapshot := range snaps {
		m := snapshot.Data()
		if proto, err := FirestoreMapToSession(m); err != nil {
			return nil, err
		} else {
			protos[j] = proto
		}
	}
	return protos, nil
}

func (i *FirestoreSessionsIterator) GetAllAsSnapshots() ([]*firestore.DocumentSnapshot, error) {
	return i.iter.GetAll()
}

func (i *FirestoreSessionsIterator) Next() (*Session, error) {
	snapshot, err := i.iter.Next()
	if err != nil {
		return nil, err
	}
	m := snapshot.Data()
	if p, err := FirestoreMapToSession(m); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

func (i *FirestoreSessionsIterator) NextAsSnapshot() (*firestore.DocumentSnapshot, error) {
	return i.iter.Next()
}

func (i *FirestoreSessionsIterator) Stop() {
	i.iter.Stop()
}

type FirestoreSessionsIterator struct {
	iter *firestore.DocumentIterator
}

func (q *FirestoreSessionsQuery) First(ctx context.Context) (*Session, error) {
	iter := q.query.Limit(1).Documents(ctx)
	defer iter.Stop()
	snapshot, err := iter.Next()
	if err != nil {
		if errors.Is(err, iterator.Done) {
			return nil, nil
		}
		return nil, err
	}
	m := snapshot.Data()
	if proto, err := FirestoreMapToSession(m); err != nil {
		return nil, err
	} else {
		return proto, nil
	}
}

type FirestoreSessionsQuery struct {
	query firestore.Query
}

func (q *FirestoreSessionsQuery) Documents(ctx context.Context) *FirestoreSessionsIterator {
	return &FirestoreSessionsIterator{
		iter: q.query.Documents(ctx),
	}
}

func (q *FirestoreSessionsQuery) Value() firestore.Query {
	return q.query
}

func (fs *Firestore) Installations() *FirestoreInstallationsCollectionRef {
	return &FirestoreInstallationsCollectionRef{
		coll: fs.client.Collection("installations"),
	}
}
func (ref *FirestoreInstallationsCollectionRef) Limit(n int) *FirestoreInstallationsQuery {
	return &FirestoreInstallationsQuery{
		query: ref.coll.Limit(n),
	}
}

func (q *FirestoreInstallationsQuery) Limit(n int) *FirestoreInstallationsQuery {
	return &FirestoreInstallationsQuery{
		query: q.query.Limit(n),
	}
}

func (ref *FirestoreInstallationsCollectionRef) OrderBy(path string, dir firestore.Direction) *FirestoreInstallationsQuery {
	return &FirestoreInstallationsQuery{
		query: ref.coll.OrderBy(path, dir),
	}
}

func (q *FirestoreInstallationsQuery) OrderBy(path string, dir firestore.Direction) *FirestoreInstallationsQuery {
	return &FirestoreInstallationsQuery{
		query: q.query.OrderBy(path, dir),
	}
}

func (ref *FirestoreInstallationsCollectionRef) Where(path, op string, value interface{}) *FirestoreInstallationsQuery {
	return &FirestoreInstallationsQuery{
		query: ref.coll.Where(path, op, value),
	}
}

func (q *FirestoreInstallationsQuery) Where(path, op string, value interface{}) *FirestoreInstallationsQuery {
	return &FirestoreInstallationsQuery{
		query: q.query.Where(path, op, value),
	}
}

// FirestoreInstallationsCollectionRef holds a reference to the Firestore collection `installations`.
type FirestoreInstallationsCollectionRef struct {
	coll *firestore.CollectionRef
}

func (ref *FirestoreInstallationsCollectionRef) Doc(id string) *FirestoreInstallationsDocumentRef {
	return &FirestoreInstallationsDocumentRef{
		doc: ref.coll.Doc(id),
	}
}

// FirestoreInstallationsDocumentRef holds a reference to a Firestore document in collection `installations`.
type FirestoreInstallationsDocumentRef struct {
	doc *firestore.DocumentRef
}

func (i *FirestoreInstallationsIterator) Stop() {
	i.iter.Stop()
}

type FirestoreInstallationsIterator struct {
	iter *firestore.DocumentIterator
}

type FirestoreInstallationsQuery struct {
	query firestore.Query
}

func (q *FirestoreInstallationsQuery) Documents(ctx context.Context) *FirestoreInstallationsIterator {
	return &FirestoreInstallationsIterator{
		iter: q.query.Documents(ctx),
	}
}

func (q *FirestoreInstallationsQuery) Value() firestore.Query {
	return q.query
}

func (ref *FirestoreInstallationsDocumentRef) Repositories() *FirestoreRepositoriesCollectionRef {
	return &FirestoreRepositoriesCollectionRef{
		coll: ref.doc.Collection("repositories"),
	}
}
func (ref *FirestoreRepositoriesCollectionRef) Limit(n int) *FirestoreRepositoriesQuery {
	return &FirestoreRepositoriesQuery{
		query: ref.coll.Limit(n),
	}
}

func (q *FirestoreRepositoriesQuery) Limit(n int) *FirestoreRepositoriesQuery {
	return &FirestoreRepositoriesQuery{
		query: q.query.Limit(n),
	}
}

func (ref *FirestoreRepositoriesCollectionRef) OrderBy(path string, dir firestore.Direction) *FirestoreRepositoriesQuery {
	return &FirestoreRepositoriesQuery{
		query: ref.coll.OrderBy(path, dir),
	}
}

func (q *FirestoreRepositoriesQuery) OrderBy(path string, dir firestore.Direction) *FirestoreRepositoriesQuery {
	return &FirestoreRepositoriesQuery{
		query: q.query.OrderBy(path, dir),
	}
}

func (ref *FirestoreRepositoriesCollectionRef) Where(path, op string, value interface{}) *FirestoreRepositoriesQuery {
	return &FirestoreRepositoriesQuery{
		query: ref.coll.Where(path, op, value),
	}
}

func (q *FirestoreRepositoriesQuery) Where(path, op string, value interface{}) *FirestoreRepositoriesQuery {
	return &FirestoreRepositoriesQuery{
		query: q.query.Where(path, op, value),
	}
}

// FirestoreRepositoriesCollectionRef holds a reference to the Firestore collection `repositories`.
type FirestoreRepositoriesCollectionRef struct {
	coll *firestore.CollectionRef
}

func (ref *FirestoreRepositoriesCollectionRef) Doc(id string) *FirestoreRepositoriesDocumentRef {
	return &FirestoreRepositoriesDocumentRef{
		doc: ref.coll.Doc(id),
	}
}

// FirestoreRepositoriesDocumentRef holds a reference to a Firestore document in collection `repositories`.
type FirestoreRepositoriesDocumentRef struct {
	doc *firestore.DocumentRef
}

func (i *FirestoreRepositoriesIterator) Stop() {
	i.iter.Stop()
}

type FirestoreRepositoriesIterator struct {
	iter *firestore.DocumentIterator
}

type FirestoreRepositoriesQuery struct {
	query firestore.Query
}

func (q *FirestoreRepositoriesQuery) Documents(ctx context.Context) *FirestoreRepositoriesIterator {
	return &FirestoreRepositoriesIterator{
		iter: q.query.Documents(ctx),
	}
}

func (q *FirestoreRepositoriesQuery) Value() firestore.Query {
	return q.query
}

func (ref *FirestoreRepositoriesDocumentRef) Manifests() *FirestoreManifestsCollectionRef {
	return &FirestoreManifestsCollectionRef{
		coll: ref.doc.Collection(FirestoreCollectionManifests),
	}
}
func (ref *FirestoreManifestsCollectionRef) Limit(n int) *FirestoreManifestsQuery {
	return &FirestoreManifestsQuery{
		query: ref.coll.Limit(n),
	}
}

func (q *FirestoreManifestsQuery) Limit(n int) *FirestoreManifestsQuery {
	return &FirestoreManifestsQuery{
		query: q.query.Limit(n),
	}
}

func (ref *FirestoreManifestsCollectionRef) OrderBy(path string, dir firestore.Direction) *FirestoreManifestsQuery {
	return &FirestoreManifestsQuery{
		query: ref.coll.OrderBy(path, dir),
	}
}

func (q *FirestoreManifestsQuery) OrderBy(path string, dir firestore.Direction) *FirestoreManifestsQuery {
	return &FirestoreManifestsQuery{
		query: q.query.OrderBy(path, dir),
	}
}

func (ref *FirestoreManifestsCollectionRef) Where(path, op string, value interface{}) *FirestoreManifestsQuery {
	return &FirestoreManifestsQuery{
		query: ref.coll.Where(path, op, value),
	}
}

func (q *FirestoreManifestsQuery) Where(path, op string, value interface{}) *FirestoreManifestsQuery {
	return &FirestoreManifestsQuery{
		query: q.query.Where(path, op, value),
	}
}

// FirestoreManifestsCollectionRef holds a reference to the Firestore collection `manifests`.
type FirestoreManifestsCollectionRef struct {
	coll *firestore.CollectionRef
}

func (ref *FirestoreManifestsCollectionRef) Doc(id string) *FirestoreManifestsDocumentRef {
	return &FirestoreManifestsDocumentRef{
		doc: ref.coll.Doc(id),
	}
}

func (ref *FirestoreManifestsDocumentRef) Delete(ctx context.Context, preconds ...firestore.Precondition) (*firestore.WriteResult, error) {
	return ref.doc.Delete(ctx, preconds...)
}

func (ref *FirestoreManifestsDocumentRef) Get(ctx context.Context) (*Manifest, error) {
	snapshot, err := ref.doc.Get(ctx)
	if err != nil {
		return nil, err
	}
	m := snapshot.Data()
	if proto, err := FirestoreMapToManifest(m); err != nil {
		return nil, err
	} else {
		return proto, nil
	}
}

func (ref *FirestoreManifestsDocumentRef) Ref() *firestore.DocumentRef {
	return ref.doc
}

func (ref *FirestoreManifestsDocumentRef) Set(ctx context.Context, msg *Manifest, opts ...firestore.SetOption) error {
	fs, err := msg.ToFirestoreMap()
	if err != nil {
		return err
	}
	if _, err := ref.doc.Set(ctx, fs, opts...); err != nil {
		return err
	}
	return nil
}

// FirestoreManifestsDocumentRef holds a reference to a Firestore document in collection `manifests`.
type FirestoreManifestsDocumentRef struct {
	doc *firestore.DocumentRef
}

func (i *FirestoreManifestsIterator) GetAll() ([]*Manifest, error) {
	snaps, err := i.iter.GetAll()
	if err != nil {
		return nil, err
	}
	protos := make([]*Manifest, len(snaps))
	for j, snapshot := range snaps {
		m := snapshot.Data()
		if proto, err := FirestoreMapToManifest(m); err != nil {
			return nil, err
		} else {
			protos[j] = proto
		}
	}
	return protos, nil
}

func (i *FirestoreManifestsIterator) GetAllAsSnapshots() ([]*firestore.DocumentSnapshot, error) {
	return i.iter.GetAll()
}

func (i *FirestoreManifestsIterator) Next() (*Manifest, error) {
	snapshot, err := i.iter.Next()
	if err != nil {
		return nil, err
	}
	m := snapshot.Data()
	if p, err := FirestoreMapToManifest(m); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

func (i *FirestoreManifestsIterator) NextAsSnapshot() (*firestore.DocumentSnapshot, error) {
	return i.iter.Next()
}

func (i *FirestoreManifestsIterator) Stop() {
	i.iter.Stop()
}

type FirestoreManifestsIterator struct {
	iter *firestore.DocumentIterator
}

func (q *FirestoreManifestsQuery) First(ctx context.Context) (*Manifest, error) {
	iter := q.query.Limit(1).Documents(ctx)
	defer iter.Stop()
	snapshot, err := iter.Next()
	if err != nil {
		if errors.Is(err, iterator.Done) {
			return nil, nil
		}
		return nil, err
	}
	m := snapshot.Data()
	if proto, err := FirestoreMapToManifest(m); err != nil {
		return nil, err
	} else {
		return proto, nil
	}
}

type FirestoreManifestsQuery struct {
	query firestore.Query
}

func (q *FirestoreManifestsQuery) Documents(ctx context.Context) *FirestoreManifestsIterator {
	return &FirestoreManifestsIterator{
		iter: q.query.Documents(ctx),
	}
}

func (q *FirestoreManifestsQuery) Value() firestore.Query {
	return q.query
}

func (ref *FirestoreManifestsDocumentRef) Actions() *FirestoreActionsCollectionRef {
	return &FirestoreActionsCollectionRef{
		coll: ref.doc.Collection(FirestoreCollectionActions),
	}
}
func (ref *FirestoreActionsCollectionRef) Limit(n int) *FirestoreActionsQuery {
	return &FirestoreActionsQuery{
		query: ref.coll.Limit(n),
	}
}

func (q *FirestoreActionsQuery) Limit(n int) *FirestoreActionsQuery {
	return &FirestoreActionsQuery{
		query: q.query.Limit(n),
	}
}

func (ref *FirestoreActionsCollectionRef) OrderBy(path string, dir firestore.Direction) *FirestoreActionsQuery {
	return &FirestoreActionsQuery{
		query: ref.coll.OrderBy(path, dir),
	}
}

func (q *FirestoreActionsQuery) OrderBy(path string, dir firestore.Direction) *FirestoreActionsQuery {
	return &FirestoreActionsQuery{
		query: q.query.OrderBy(path, dir),
	}
}

func (ref *FirestoreActionsCollectionRef) Where(path, op string, value interface{}) *FirestoreActionsQuery {
	return &FirestoreActionsQuery{
		query: ref.coll.Where(path, op, value),
	}
}

func (q *FirestoreActionsQuery) Where(path, op string, value interface{}) *FirestoreActionsQuery {
	return &FirestoreActionsQuery{
		query: q.query.Where(path, op, value),
	}
}

// FirestoreActionsCollectionRef holds a reference to the Firestore collection `actions`.
type FirestoreActionsCollectionRef struct {
	coll *firestore.CollectionRef
}

func (ref *FirestoreActionsCollectionRef) Doc(id string) *FirestoreActionsDocumentRef {
	return &FirestoreActionsDocumentRef{
		doc: ref.coll.Doc(id),
	}
}

func (ref *FirestoreActionsDocumentRef) Delete(ctx context.Context, preconds ...firestore.Precondition) (*firestore.WriteResult, error) {
	return ref.doc.Delete(ctx, preconds...)
}

func (ref *FirestoreActionsDocumentRef) Get(ctx context.Context) (*Action, error) {
	snapshot, err := ref.doc.Get(ctx)
	if err != nil {
		return nil, err
	}
	m := snapshot.Data()
	if proto, err := FirestoreMapToAction(m); err != nil {
		return nil, err
	} else {
		return proto, nil
	}
}

func (ref *FirestoreActionsDocumentRef) Ref() *firestore.DocumentRef {
	return ref.doc
}

func (ref *FirestoreActionsDocumentRef) Set(ctx context.Context, msg *Action, opts ...firestore.SetOption) error {
	fs, err := msg.ToFirestoreMap()
	if err != nil {
		return err
	}
	if _, err := ref.doc.Set(ctx, fs, opts...); err != nil {
		return err
	}
	return nil
}

// FirestoreActionsDocumentRef holds a reference to a Firestore document in collection `actions`.
type FirestoreActionsDocumentRef struct {
	doc *firestore.DocumentRef
}

func (i *FirestoreActionsIterator) GetAll() ([]*Action, error) {
	snaps, err := i.iter.GetAll()
	if err != nil {
		return nil, err
	}
	protos := make([]*Action, len(snaps))
	for j, snapshot := range snaps {
		m := snapshot.Data()
		if proto, err := FirestoreMapToAction(m); err != nil {
			return nil, err
		} else {
			protos[j] = proto
		}
	}
	return protos, nil
}

func (i *FirestoreActionsIterator) GetAllAsSnapshots() ([]*firestore.DocumentSnapshot, error) {
	return i.iter.GetAll()
}

func (i *FirestoreActionsIterator) Next() (*Action, error) {
	snapshot, err := i.iter.Next()
	if err != nil {
		return nil, err
	}
	m := snapshot.Data()
	if p, err := FirestoreMapToAction(m); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

func (i *FirestoreActionsIterator) NextAsSnapshot() (*firestore.DocumentSnapshot, error) {
	return i.iter.Next()
}

func (i *FirestoreActionsIterator) Stop() {
	i.iter.Stop()
}

type FirestoreActionsIterator struct {
	iter *firestore.DocumentIterator
}

func (q *FirestoreActionsQuery) First(ctx context.Context) (*Action, error) {
	iter := q.query.Limit(1).Documents(ctx)
	defer iter.Stop()
	snapshot, err := iter.Next()
	if err != nil {
		if errors.Is(err, iterator.Done) {
			return nil, nil
		}
		return nil, err
	}
	m := snapshot.Data()
	if proto, err := FirestoreMapToAction(m); err != nil {
		return nil, err
	} else {
		return proto, nil
	}
}

type FirestoreActionsQuery struct {
	query firestore.Query
}

func (q *FirestoreActionsQuery) Documents(ctx context.Context) *FirestoreActionsIterator {
	return &FirestoreActionsIterator{
		iter: q.query.Documents(ctx),
	}
}

func (q *FirestoreActionsQuery) Value() firestore.Query {
	return q.query
}
