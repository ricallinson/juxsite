package article

import (
	"github.com/ricallinson/forgery"
)

// Route the request to the correct handler function.
func MenuHandler(req *f.Request, res *f.Response, next func()) {
	if req.Params["juxview"] != "article" && req.Params["juxview"] != "read" {
		return
	}
	articles := ListArticles(req)
	res.Render("article/menu.html", map[string][]*Article{
		"links": articles,
	}, map[string]string{
		"title": "All Articles",
	})
}
