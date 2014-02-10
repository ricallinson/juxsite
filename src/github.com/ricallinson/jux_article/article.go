package jux_article

import (
	"bytes"
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
