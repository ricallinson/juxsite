package article

import (
	"appengine"
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
	res.Locals["title"] = "All Articles"
	articles := ListArticles(req)
	res.Render("article/list.html", map[string][]*Article{
		"articles": articles,
	})
}

// Shows a list of articles for the given category.
func read(req *f.Request, res *f.Response, next func()) {
	article := &Article{Id: req.Query["id"]}
	article.Read(req)

	context := appengine.NewContext(req.Request.Request)
    context.Debugf("Article: %v", article)

	res.Render("article/read.html", map[string][]*Article{
		"articles": []*Article{article},
	})
}

// Shows a list of articles for the given category.
func listJson(req *f.Request, res *f.Response, next func()) {
	res.Json(loadArticles("data/articles"))
}
