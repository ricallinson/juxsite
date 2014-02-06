package jux

import (
	"github.com/ricallinson/forgery"
	"github.com/ricallinson/jux/assets"
)

// Register all default components.
func registerDefaultComponents(cfg *AppCfg) {

	cfg.RegisterComponent("default", func(req *f.Request, res *f.Response, next func()) {
		a := assets.Make(res.Locals)
		a.AddCss("default/bootstrap/css/bootstrap.css")
		a.AddCss("default/css/screen.css")
		a.Render()
		res.Render("default/main.html")
	})

	cfg.RegisterComponent("a", func(req *f.Request, res *f.Response, next func()) {

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
	})

	cfg.RegisterComponent("b", func(req *f.Request, res *f.Response, next func()) {
		res.Render("b/main.html", map[string]string{
			"title": "Login Form",
			"msg":   "B",
		})
	})

	cfg.RegisterComponent("c", func(req *f.Request, res *f.Response, next func()) {
		res.Render("c/main.html", map[string]string{
			"msg": "C",
		})
	})

	cfg.RegisterComponent("article", func(req *f.Request, res *f.Response, next func()) {
		res.Locals["title"] = "Article"
		res.Render("article/main.html", map[string]string{
			"msg": req.Params["juxcomp"] + "/" + req.Params["juxview"],
		})
	})

	cfg.RegisterComponent("error", func(req *f.Request, res *f.Response, next func()) {
		res.Render("error/main.html", map[string]string{
			"error": "Component not found.",
		})
	})
}
