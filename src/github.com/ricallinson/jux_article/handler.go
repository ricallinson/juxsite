package jux_article

import (
	"github.com/ricallinson/forgery"
	"github.com/russross/blackfriday"
	"strconv"
	"strings"
)

// Route the request to the correct handler function.
func Handler(req *f.Request, res *f.Response, next func()) {
	switch req.Params["juxview"] {
	case "main":
		list(req, res, next)
	case "read":
		read(req, res, next)
	case "listjson":
		listJson(req, res, next)
	}
}

// Shows a list of articles for the given category.
func list(req *f.Request, res *f.Response, next func()) {
	batch := 5
	start, _ := strconv.Atoi(req.Query["start"])
	category := strings.ToLower(req.Query["category"])
	articles, count := ListArticles(req, category, start, start+batch)
	less := start - batch
	more := start + batch
	// Render the Summary as HTML.
	for _, article := range articles {
		if line := strings.Index(article.Text, "\r\n"); line > 0 {
			article.Summary = string(blackfriday.MarkdownBasic([]byte(article.Text[:line])))
		}
	}
	res.Render("article/list.html", map[string][]*Article{
		"articles": articles,
	}, map[string]string{
		"title": "All Articles",
		"less":  strconv.Itoa(less),
		"more":  strconv.Itoa(more),
	}, map[string]bool{
		"show_less": less >= 0,
		"show_more": more < count,
	})
}

// Shows the single article for the given id.
func read(req *f.Request, res *f.Response, next func()) {
	article := &Article{Id: req.Query["id"]}
	if article.Id == "" || article.Read(req) != nil {
		res.Render("notfound/main.html", map[string]string{
			"error": "Article not found.",
		})
		return
	}
	res.Locals["pageTitle"] = article.Title
	// Render the text as HTML.
	article.Text = string(blackfriday.MarkdownBasic([]byte(article.Text)))
	res.Render("article/read.html", map[string][]*Article{
		"articles": []*Article{article},
	})
}

// Shows a list of articles for the given category.
func listJson(req *f.Request, res *f.Response, next func()) {
	batch := 5
	start, _ := strconv.Atoi(req.Query["start"])
	category := strings.ToLower(req.Query["category"])
	articles, _ := ListArticles(req, category, start, start+batch)
	res.Json(articles)
}
