package jux

import (
	"github.com/ricallinson/fmustache"
	"github.com/ricallinson/forgery"
	"net/http"
)

// Register all default components.
func registerDefaultComponents(cfg *AppCfg) {

	cfg.RegisterComponent("a", func(req *f.Request, res *f.Response, next func()) {
	    res.Send("Header")
	})

	cfg.RegisterComponent("b", func(req *f.Request, res *f.Response, next func()) {
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

	// Set all template locals

	app.Locals["title"] = cfg.Page.Title

	// Create standard routes

	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxcomp"] = cfg.DefaultComponent
		req.Params["juxview"] = cfg.DefaultComponentView
		Render(req, res, next, cfg)
	})

	app.Get("/:juxcomp", func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxview"] = cfg.DefaultComponentView
		Render(req, res, next, cfg)
	})

	app.Get("/:juxcomp/:juxview", func(req *f.Request, res *f.Response, next func()) {
		Render(req, res, next, cfg)
	})

	http.Handle("/", app)
}
