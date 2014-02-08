package jux_article

import (
	"errors"
	"github.com/ricallinson/forgery"
)

type Article struct {
	Id       string
	Title    string
	Category string
	Summary  string
	Text     string
}

func ListArticles(req *f.Request, category string, from int, to int) ([]*Article, int) {
	ds := GetFileDataStore(req, "data")
	return ds.LoadTable("articles", category, from, to)
}

func (this *Article) Create(req *f.Request) error {
	return nil
}

func (this *Article) Read(req *f.Request) error {
	ds := GetFileDataStore(req, "data")
	if article, ok := ds.LoadItem("articles", this.Id); ok {
		// This is a hack, need to work out why "this = article" fails.
		this.Title = article.Title
		this.Category = article.Category
		this.Summary = article.Summary
		this.Text = article.Text
		return nil
	}
	return errors.New("Article not found.")
}

func (this *Article) Update(req *f.Request) error {
	return nil
}

func (this *Article) Delete(req *f.Request) error {
	return nil
}
