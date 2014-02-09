package datastore

import (
	"appengine"
	"appengine/datastore"
	"errors"
	"reflect"
	// "sort"
	// "time"
)

type Entry struct {

	// The path used for the Item.
	Id string

	// The User ID.
	Uid string

	// Version as a time stamp
	Created string

	// Version as a time stamp
	Updated string
}

type DataStore struct {
	// Appengine Context
	Context appengine.Context
}

func New(context appengine.Context) *DataStore {
	ds := &DataStore{}
	ds.Context = context
	return ds
}

// Create a Key from the interface and id.
func (this *DataStore) CreateKey(id string, i interface{}) *datastore.Key {
	return datastore.NewKey(this.Context, reflect.TypeOf(i).Elem().Name(), id, 0, nil)
}

// Create an Entry in the datastore.
func (this *DataStore) Create(id string, i interface{}) error {
	_, err := datastore.Put(this.Context, this.CreateKey(id, i), i)
	return err
}

// Read an Entry from the datastore.
func (this *DataStore) Read(id string, i interface{}) error {
	err := datastore.Get(this.Context, this.CreateKey(id, i), i)
	if err != nil {
		return errors.New("Entry not found.")
	}
	return nil
}

// Update an Entry in the datastore.
func (this *DataStore) Update(id string, i interface{}) error {
	return this.Create(id, i)
}

// Delete an Entry from the datastore.
func (this *DataStore) Delete(id string, i interface{}) error {
	return datastore.Delete(this.Context, this.CreateKey(id, i))
}
