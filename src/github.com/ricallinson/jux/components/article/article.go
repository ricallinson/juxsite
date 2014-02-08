package article

import (
	// "appengine"
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
	// c := appengine.NewContext(req.Request.Request)
	ds := CreateDataStore("data")
	return ds.LoadTable("articles", category, from, to)
}

func (this *Article) Create(req *f.Request) error {
	// c := appengine.NewContext(req.Request.Request)
	return nil
}

func (this *Article) Read(req *f.Request) error {
	// c := appengine.NewContext(req.Request.Request)
	ds := CreateDataStore("data")
	if ds.LoadItem("articles", this) == false {
		return errors.New("Article not found.")
	}
	return nil
}

func (this *Article) Update(req *f.Request) error {
	// c := appengine.NewContext(req.Request.Request)
	return nil
}

func (this *Article) Delete(req *f.Request) error {
	// c := appengine.NewContext(req.Request.Request)
	return nil
}
