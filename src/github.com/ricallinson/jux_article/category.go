package jux_article

import (
    "github.com/ricallinson/forgery"
    "github.com/ricallinson/jux"
    "github.com/ricallinson/jux/helpers/datastore"
)

type Category struct {
    datastore.Entry
    Title string
}

func GetCategories(req *f.Request) []*Category {

    // Create a Query.
    query := &Category{}
    categories := []*Category{}

    // Grab the datastore.
    ds := datastore.New(jux.GetNewContext(req))
    ds.List(query, 0, -1, &categories)

    return categories
}

func GetCategoriesMap(req *f.Request) map[string]string {

    categories := map[string]string{}

    for _, category := range GetCategories(req) {
        categories[category.Id] = category.Title
    }

    return categories
}

func LookUpCategory(req *f.Request, category string, def string) string {

    categories := GetCategoriesMap(req)

    if match, ok := categories[category]; ok {
        return match
    }

    return def
}
