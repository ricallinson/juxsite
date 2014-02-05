package jux

import (
	"github.com/ricallinson/fmustache"
	"github.com/ricallinson/forgery"
	"net/http"
)

func Start() {

	app := f.CreateServer()

	app.Use(f.ResponseTime())
	app.Use(f.Favicon())
	app.Use(f.Static())

	app.Engine(".html", fmustache.Make())

	cfg := &AppCfg{}
	cfg.Load("site.yaml")

	app.Locals["title"] = cfg.PageTitle

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
