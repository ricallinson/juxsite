package jux_article

import (
	ds "appengine/datastore"
	"bytes"
	"github.com/ricallinson/forgery"
	"github.com/ricallinson/jux"
	"github.com/ricallinson/jux/helpers/datastore"
	"github.com/russross/blackfriday"
)

// An Article.
type Article struct {
	datastore.Entry
	Title    string
	Category string
	Text     []byte
	Summary  string `datastore:"-"` // This value is NOT added to the datastore.
	Body     string `datastore:"-"` // This value is NOT added to the datastore.
}

// Fill in Summary and Body values from the Text byte slice.
func (this *Article) InflateSummary() {
	line := bytes.Index(this.Text, []byte("\r\n"))
	if line == -1 {
		line = len(this.Text)
	}
	this.Summary = string(blackfriday.MarkdownBasic(this.Text[:line]))
}

// Fill in Summary and Body values from the Text byte slice.
func (this *Article) InflateBody() {
	this.Body = string(blackfriday.MarkdownBasic(this.Text))
}

// Return a list of articles matching the given category.
func GetArticles(req *f.Request, category string, offset int, limit int) ([]*Article, error) {

	articles := []*Article{}

	// Create a query for the given interface.
	query := ds.NewQuery("Article").
		Filter("Category =", category).
		Order("Created").
		Offset(offset).
		Limit(limit)

	_, err := query.GetAll(jux.GetNewContext(req), &articles)

	return articles, err
}
