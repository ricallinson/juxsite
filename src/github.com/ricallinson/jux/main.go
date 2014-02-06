package jux

import (
	"fmt"
	"github.com/ricallinson/fmustache"
	"github.com/ricallinson/forgery"
	"github.com/ricallinson/jux/secure"
	"net/http"
	"time"
)

func Start(cfg *AppCfg) {

	app := f.CreateServer()

	app.Use(f.ResponseTime())
	app.Use(f.Favicon())
	app.Use(f.Static())

	app.Engine(".html", fmustache.Make())

	registerDefaultComponents(cfg)

	// Set template locals.
	app.Env = cfg.Defaults.Env
	app.Locals["baseUrl"] = cfg.Page.BaseUrl
	app.Locals["siteName"] = cfg.Page.Name
	app.Locals["pageTitle"] = cfg.Page.Name
	app.Locals["description"] = cfg.Page.Description
	app.Locals["lang"] = cfg.Page.Lang
	app.Locals["direction"] = cfg.Page.Direction
	app.Locals["year"] = fmt.Sprint(time.Now().Year())

	// Create standard routes.

	app.Get("/admin", secure.AdminAuth, func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxcomp"] = cfg.Defaults.Component
		req.Params["juxview"] = cfg.Defaults.ComponentView
		Render(req, res, next, cfg, app)
	})

	app.Get("/admin/:juxcomp", secure.AdminAuth, func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxview"] = cfg.Defaults.ComponentView
		Render(req, res, next, cfg, app)
	})

	app.Get("/admin/:juxcomp/:juxview", secure.AdminAuth, func(req *f.Request, res *f.Response, next func()) {
		Render(req, res, next, cfg, app)
	})

	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxcomp"] = cfg.Defaults.Component
		req.Params["juxview"] = cfg.Defaults.ComponentView
		Render(req, res, next, cfg, app)
	})

	app.Get("/:juxcomp", func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxview"] = cfg.Defaults.ComponentView
		Render(req, res, next, cfg, app)
	})

	app.Get("/:juxcomp/:juxview", func(req *f.Request, res *f.Response, next func()) {
		Render(req, res, next, cfg, app)
	})

	http.Handle("/", app)
}
