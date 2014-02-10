package datastore

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"errors"
	"reflect"
	"time"
)

type Entry struct {

	// The path used for the Item.
	Id string

	// The User ID.
	Uid string

	// Version as a time stamp
	Created int64

	// Version as a time stamp
	Updated int64
}

type DataStore struct {
	// Appengine Context
	Context appengine.Context
	Uid     string
}

func New(context appengine.Context) *DataStore {
	ds := &DataStore{}
	ds.Context = context
	ds.Uid = user.Current(context).ID
	return ds
}

func (this *DataStore) List(i interface{}, offset int, limit int, list interface{}) error {

	// Grab a handle to the value.
	value := reflect.ValueOf(i).Elem()

	// Create a query for the given interface.
	query := datastore.NewQuery(reflect.TypeOf(i).Elem().Name()).
		Filter("Category =", value.FieldByName("Category").String()).
		Order("Created").
		Offset(offset).
		Limit(limit)

	// Run the query.
	_, err := query.GetAll(this.Context, list)
	return err
}

func (this *DataStore) Validate(i interface{}) error {
	entry := reflect.ValueOf(i).Elem()
	if entry.FieldByName("Id").String() == "" {
		return errors.New("Id is a required field.")
	}
	if entry.FieldByName("Uid").String() == "" {
		return errors.New("Uid is a required field.")
	}
	return nil
}

// Create a Key from the interface and id.
func (this *DataStore) CreateKey(i interface{}) *datastore.Key {
	// Get the Entry.Id from the interface.
	id := reflect.ValueOf(i).Elem().FieldByName("Id").String()
	return datastore.NewKey(this.Context, reflect.TypeOf(i).Elem().Name(), id, 0, nil)
}

// Create an Entry in the datastore.
func (this *DataStore) Create(i interface{}) error {
	if reflect.ValueOf(i).Elem().FieldByName("Uid").String() == "" {
		reflect.ValueOf(i).Elem().FieldByName("Uid").SetString(this.Uid)
	}
	if err := this.Validate(i); err != nil {
		return err
	}
	reflect.ValueOf(i).Elem().FieldByName("Created").SetInt(time.Now().Unix())
	_, err := datastore.Put(this.Context, this.CreateKey(i), i)
	return err
}

// Read an Entry from the datastore.
func (this *DataStore) Read(i interface{}) error {
	err := datastore.Get(this.Context, this.CreateKey(i), i)
	if err != nil {
		return errors.New("Entry not found.")
	}
	return nil
}

// Update an Entry in the datastore.
func (this *DataStore) Update(i interface{}) error {
	if err := this.Validate(i); err != nil {
		return err
	}
	reflect.ValueOf(i).Elem().FieldByName("Updated").SetInt(time.Now().Unix())
	return this.Create(i)
}

// Delete an Entry from the datastore.
func (this *DataStore) Delete(i interface{}) error {
	return datastore.Delete(this.Context, this.CreateKey(i))
}
