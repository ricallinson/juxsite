package jux_article

import (
	"errors"
	"github.com/ricallinson/forgery"
)

// An Article.
type Article struct {
	Id       string
	Title    string
	Category string
	Summary  string
	Text     string
}

// Returns a list of all items in the given "table".
func ListArticles(req *f.Request, category string, from int, to int) ([]*Article, int) {
	ds := GetFileDataStore(req, "data")
	return ds.GetTable("articles", category, from, to)
}

// Store a this Article in persistent storage.
func (this *Article) Create(req *f.Request) error {
	panic("Not implemented.")
	return nil
}

// Read a this Article from persistent storage.
func (this *Article) Read(req *f.Request) error {
	ds := GetFileDataStore(req, "data")
	if article, ok := ds.GetItem("articles", this.Id); ok {
		// This is a hack, need to work out why "this = article" fails.
		this.Title = article.Title
		this.Category = article.Category
		this.Summary = article.Summary
		this.Text = article.Text
		return nil
	}
	return errors.New("Article not found.")
}

// Update a this Article in persistent storage.
func (this *Article) Update(req *f.Request) error {
	panic("Not implemented.")
	return nil
}

// Remove a this Article from persistent storage.
func (this *Article) Delete(req *f.Request) error {
	panic("Not implemented.")
	return nil
}
