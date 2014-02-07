package article

import (
	"github.com/ricallinson/forgery"
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
	articles := ListArticles(req, 0, 100)
	res.Render("article/list.html", map[string][]*Article{
		"articles": articles,
	}, map[string]string{
		"title": "All Articles",
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
	res.Json(ListArticles(req, 0, 100))
}
