// Code generated by protoc-gen-go-firestore. DO NOT EDIT.
// versions:
// 	protoc-gen-go-firestore 1.0.1
// 	protoc          (unknown)

package utils

import (
	firestore "cloud.google.com/go/firestore"
)

type FS_utils struct {
	client *firestore.Client
}

func Firestore(client *firestore.Client) *FS_utils {
	return &FS_utils{
		client: client,
	}
}

func (f *FS_utils) Accounts() *FS_utils_Accounts {
	return &FS_utils_Accounts{
		c: f.client.Collection("accounts"),
	}
}

type FS_utils_Accounts struct {
	c *firestore.CollectionRef
}

func (f *FS_utils_Accounts) Doc(id string) *FS_utils_Accounts_Doc {
	return &FS_utils_Accounts_Doc{
		d: f.c.Doc(id),
	}
}

type FS_utils_Accounts_Doc struct {
	d *firestore.DocumentRef
}

func (f *FS_utils_Accounts_Doc) Users() *FS_utils_Accounts_Users {
	return &FS_utils_Accounts_Users{
		c: f.d.Collection("users"),
	}
}

type FS_utils_Accounts_Users struct {
	c *firestore.CollectionRef
}

func (f *FS_utils_Accounts_Users) Doc(id string) *FS_utils_Accounts_Users_Doc {
	return &FS_utils_Accounts_Users_Doc{
		d: f.c.Doc(id),
	}
}

type FS_utils_Accounts_Users_Doc struct {
	d *firestore.DocumentRef
}

func (f *FS_utils_Accounts_Users_Doc) Sessions() *FS_utils_Accounts_Users_Sessions {
	return &FS_utils_Accounts_Users_Sessions{
		c: f.d.Collection("sessions"),
	}
}

type FS_utils_Accounts_Users_Sessions struct {
	c *firestore.CollectionRef
}

func (f *FS_utils_Accounts_Users_Sessions) Doc(id string) *FS_utils_Accounts_Users_Sessions_Doc {
	return &FS_utils_Accounts_Users_Sessions_Doc{
		d: f.c.Doc(id),
	}
}

type FS_utils_Accounts_Users_Sessions_Doc struct {
	d *firestore.DocumentRef
}
