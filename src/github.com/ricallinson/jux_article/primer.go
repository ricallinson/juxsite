package jux_article

import (
	"encoding/json"
	"github.com/ricallinson/forgery"
	"github.com/ricallinson/jux"
	"github.com/ricallinson/jux/helpers/datastore"
	"io/ioutil"
	"path"
	"time"
)

type JsonArticle struct {
	Id       string
	Title    string
	Category string
	Text     string
}

// Read all files in the given directory and store them in the DataStore.
func LoadJsonArticles(req *f.Request, dirname string) error {
	ds := datastore.New(jux.GetNewContext(req))
	list, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	for _, file := range list {
		if file.IsDir() != true {
			filepath := path.Join(dirname, file.Name())
			if source, ok := ReadJsonArticle(filepath); ok {
				article := &Article{}
				article.Id = file.Name()
				article.Title = source.Title
				article.Category = source.Category
				article.Text = []byte(source.Text)
				if err := ds.Create(article); err != nil {
					return err
				}
			}
		}
		time.Sleep(time.Second)
	}
	return nil
}

// Read the JSON file at given "filepath" into a &JsonArticle{}.
func ReadJsonArticle(filepath string) (*JsonArticle, bool) {
	article := &JsonArticle{}
	j, err1 := ioutil.ReadFile(filepath)
	if err1 != nil {
		panic(err1)
		return article, false
	}
	err2 := json.Unmarshal(j, &article)
	if err2 != nil {
		panic(err2)
		return article, false
	}
	return article, true
}

// Add the listed categories to the DataStore.
func LoadCategories(req *f.Request) error {

	// Grab the data store.
	ds := datastore.New(jux.GetNewContext(req))

	// Get the Category title.
	categories := map[string]string{
		"general": "General",
		"alice": "Alice's Adventures in Wonderland",
	}

	// Walk over the categories and add them to the store.
	for key, value := range categories {
		category := &Category{}
		category.Id = key
		category.Title = value
		if err := ds.Create(category); err != nil {
			return err
		}
	}

	return nil
}
