package article

import (
	"github.com/ricallinson/forgery"
)

// Route the request to the correct handler function.
func Menu(req *f.Request, res *f.Response, next func()) {
	if req.Params["juxview"] != "article" && req.Params["juxview"] != "read" {
		return
	}
	articles := ListArticles(req, 0, 100)
	res.Render("article/menu.html", map[string][]*Article{
		"links": articles,
	}, map[string]string{
		"title": "All Articles",
	})
}
