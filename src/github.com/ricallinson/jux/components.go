package jux

import (
	"github.com/ricallinson/forgery"
)

// Register all default components.
func registerDefaultComponents(cfg *AppCfg) {

	cfg.RegisterComponent("a", func(req *f.Request, res *f.Response, next func()) {
		res.Render("components/a/main.html", map[string]string{
			"msg": "A",
		})
		// res.End("A")
	})

	cfg.RegisterComponent("b", func(req *f.Request, res *f.Response, next func()) {
		res.Render("components/b/main.html", map[string]string{
			"msg": "B",
		})
		// res.End("B")
	})

	cfg.RegisterComponent("c", func(req *f.Request, res *f.Response, next func()) {
		res.Render("components/c/main.html", map[string]string{
			"msg": "C",
		})
		// res.End("C")
	})

	cfg.RegisterComponent("article", func(req *f.Request, res *f.Response, next func()) {
		res.Locals["title"] = "Article"
		res.Render("components/article/main.html", map[string]string{
			"msg": req.Params["juxcomp"] + "/" + req.Params["juxview"],
		})
	})

	cfg.RegisterComponent("error", func(req *f.Request, res *f.Response, next func()) {
		res.Render("components/error/main.html", map[string]string{
			"error": "Component not found.",
		})
	})
}
