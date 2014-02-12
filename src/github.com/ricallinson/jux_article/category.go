package jux_article

import (
    "appengine/datastore"
    "github.com/ricallinson/forgery"
    "github.com/ricallinson/jux"
    "github.com/ricallinson/jux/helpers/datastore"
)

type Category struct {
    datastore.Entry
    Title string
}

func GetCategories(req *f.Request) []*Category, err {

    // Create a Query.
    query := &Category{}
    categories := []*Category{}

    // ds := datastore.New(jux.GetNewContext(req))
    // ds.List(query, 0, -1, &categories)

    query := datastore.NewQuery("Category")
    _, err := query.GetAll(jux.GetNewContext(req), categories)

    return categories, err
}

func GetCategoriesMap(req *f.Request) map[string]string {

    categories := map[string]string{}

    list, _ := GetCategories(req)

    for _, category := range list {
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
