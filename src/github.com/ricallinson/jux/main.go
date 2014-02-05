package jux

import (
	"fmt"
	"github.com/ricallinson/fmustache"
	"github.com/ricallinson/forgery"
	"net/http"
	"time"
)

// Register all default components.
func registerDefaultComponents(cfg *AppCfg) {

	cfg.RegisterComponent("a", func(req *f.Request, res *f.Response, next func()) {
		res.Send("Nav Bar")
	})

	cfg.RegisterComponent("article", func(req *f.Request, res *f.Response, next func()) {
		res.Send(req.Params["juxcomp"] + "/" + req.Params["juxview"])
	})

	cfg.RegisterComponent("c", func(req *f.Request, res *f.Response, next func()) {
		res.End("Footer")
	})
}

func Start(cfg *AppCfg) {

	app := f.CreateServer()

	app.Use(f.ResponseTime())
	app.Use(f.Favicon())
	app.Use(f.Static())

	app.Engine(".html", fmustache.Make())

	registerDefaultComponents(cfg)

	// Set template locals.

	app.Locals["baseUrl"] = cfg.Site.BaseUrl
	app.Locals["siteName"] = cfg.Site.Name
	app.Locals["SiteDescription"] = cfg.Site.Description
	app.Locals["lang"] = cfg.Site.Lang
	app.Locals["direction"] = cfg.Site.Direction
	app.Locals["year"] = fmt.Sprint(time.Now().Year())

	// Create standard routes.

	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxcomp"] = cfg.Defaults.Component
		req.Params["juxview"] = cfg.Defaults.ComponentView
		Render(req, res, next, cfg)
	})

	app.Get("/:juxcomp", func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxview"] = cfg.Defaults.ComponentView
		Render(req, res, next, cfg)
	})

	app.Get("/:juxcomp/:juxview", func(req *f.Request, res *f.Response, next func()) {
		Render(req, res, next, cfg)
	})

	http.Handle("/", app)
}
