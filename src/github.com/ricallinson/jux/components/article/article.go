package article

import (
	// "appengine"
	"github.com/ricallinson/forgery"
)

type Article struct {
	Id       string
	Title    string
	Category string
	Summary  string
	Text     string
}

func ListArticles(req *f.Request) []*Article {
	// c := appengine.NewContext(req.Request.Request)
	return loadArticles("data/articles")
}

func (this *Article) Create(req *f.Request) error {
	// c := appengine.NewContext(req.Request.Request)
	return nil
}

func (this *Article) Read(req *f.Request) error {
	// c := appengine.NewContext(req.Request.Request)
	loadArticle("data/articles", this)
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
