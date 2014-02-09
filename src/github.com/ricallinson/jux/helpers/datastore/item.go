package datastore

import (
	"appengine"
	"appengine/datastore"
	"errors"
	"sort"
	"time"
)

type Item struct {

	// The path used for the Item.
	Id string

	// The User ID.
	Uid string

	// Version as a time stamp
	Version string

	// Appengine Context
	context appengine.Context

	// Key for the datastore
	key *datastore.Key
}

/*
   List all the Items.
*/
func List(context appengine.Context) (list []string) {
	var items []Item
	q := datastore.NewQuery("Item") //.Filter("Uid =", uid)
	_, err := q.GetAll(context, &items)
	if err != nil || len(items) == 0 {
		return
	}
	for _, item := range items {
		list = append(list, item.Id)
	}
	sort.Strings(list)
	return
}

/*
   Returns the datastore Key for this Item.
*/
func (this *Item) Key(context appengine.Context) *datastore.Key {
	if this.key != nil {
		return this.key
	}
	var items []Item
	q := datastore.NewQuery("Item").Filter("Id =", this.Id).Filter("Uid =", this.Uid).Limit(1)
	keys, err := q.GetAll(context, &items)
	if err != nil || len(items) == 0 {
		return &datastore.Key{}
	}
	this.key = keys[0]
	return this.key
}

/*
   Create the Item in the datastore.
*/
func (this *Item) Create(context appengine.Context) error {
	this.Version = time.Now().UTC().Format(time.RFC1123Z)
	key, err := datastore.Put(context, datastore.NewIncompleteKey(context, "Item", nil), this)
	this.key = key
	return err
}

/*
   Read this Item from the datastore.
*/
func (this *Item) Read(context appengine.Context) error {
	err := datastore.Get(context, this.Key(context), this)
	if err != nil {
		return errors.New("Item not found")
	}
	return nil
}

/*
   Update this Item in the datastore.
*/
func (this *Item) Update(context appengine.Context) error {
	this.Version = time.Now().UTC().Format(time.RFC1123Z)
	_, err := datastore.Put(context, this.Key(context), this)
	return err
}

/*
   Delete this Item from the datastore.
*/
func (this *Item) Delete(context appengine.Context) error {
	return datastore.Delete(context, this.Key(context))
}
