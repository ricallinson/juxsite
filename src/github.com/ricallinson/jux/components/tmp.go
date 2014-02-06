package components

import (
	"github.com/ricallinson/forgery"
)

func ArticleHandler(req *f.Request, res *f.Response, next func()) {
	res.Locals["title"] = "Article"
	res.Render("article/main.html", map[string]string{
		"msg": req.Params["juxcomp"] + "/" + req.Params["juxview"],
	})
}

func AHandler(req *f.Request, res *f.Response, next func()) {
	type Link struct {
		Uri  string
		Text string
	}

	res.Render("a/main.html",
		map[string]string{
			"title": "Main Menu",
		},
		map[string][]Link{
			"links": []Link{
				Link{
					Uri:  "/",
					Text: "Home",
				},
			},
		})
}

func BHandler(req *f.Request, res *f.Response, next func()) {
	res.Render("b/main.html", map[string]string{
		"title": "Login Form",
		"msg":   "B",
	})
}

func CHandler(req *f.Request, res *f.Response, next func()) {
	res.Render("c/main.html", map[string]string{
		"msg": "C",
	})
}
