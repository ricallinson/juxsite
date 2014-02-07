package article

import (
	"github.com/ricallinson/forgery"
)

// Route the request to the correct handler function.
func Handler(req *f.Request, res *f.Response, next func()) {
	switch req.Params["juxview"] {
	case "main":
		list(req, res, next)
	case "listjson":
		listJson(req, res, next)
	}
}

// Shows a list of articles for the given category.
func list(req *f.Request, res *f.Response, next func()) {
	res.Locals["title"] = "All Articles"
	articles := loadArticles("data/articles")
	res.Render("article/list.html", map[string][]article{
		"articles": articles,
	})
}

// Shows a list of articles for the given category.
func listJson(req *f.Request, res *f.Response, next func()) {
	res.Json(loadArticles("data/articles"))
}
