package jux_article

import (
	ds "appengine/datastore"
	"github.com/ricallinson/forgery"
	"github.com/ricallinson/jux"
	"github.com/ricallinson/jux/helpers/datastore"
)

type Category struct {
	datastore.Entry
	Category string
	Title    string
}

func GetCategories(req *f.Request) ([]*Category, error) {

	categories := []*Category{}

	// I would like to use datastore.DataStore.List() here.
	query := ds.NewQuery("Category")
	_, err := query.GetAll(jux.GetNewContext(req), &categories)

	return categories, err
}

func GetCategoriesMap(req *f.Request) (map[string]string, error) {

	categories := map[string]string{}

	list, err := GetCategories(req)

	for _, category := range list {
		categories[category.Id] = category.Title
	}

	return categories, err
}

func LookUpCategory(req *f.Request, category string, def string) string {

	categories, _ := GetCategoriesMap(req)

	if match, ok := categories[category]; ok {
		return match
	}

	return def
}
