package article

import (
	"github.com/ricallinson/forgery"
	"strconv"
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
	articles, count := ListArticles(req, start, start+batch)
	less := start - batch
	more := start + batch
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

// Shows a list of articles for the given category.
func read(req *f.Request, res *f.Response, next func()) {
	article := &Article{Id: req.Query["id"]}
	if article.Id == "" || article.Read(req) != nil {
		res.Render("notfound/main.html", map[string]string{
			"error": "Article not found.",
		})
		return
	}
	res.Locals["pageTitle"] = article.Title
	res.Render("article/read.html", map[string][]*Article{
		"articles": []*Article{article},
	})
}

// Shows a list of articles for the given category.
func listJson(req *f.Request, res *f.Response, next func()) {
	batch := 5
	start, _ := strconv.Atoi(req.Query["start"])
	articles, _ := ListArticles(req, start, start+batch)
	res.Json(articles)
}
