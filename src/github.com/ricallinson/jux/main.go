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

	// Set template locals.
	app.Env = cfg.App.Defaults.Env
	app.Locals["baseUrl"] = cfg.App.Page.BaseUrl
	app.Locals["siteName"] = cfg.App.Page.Name
	app.Locals["siteDescription"] = cfg.App.Page.Description
	app.Locals["pageTitle"] = cfg.App.Page.Name
	app.Locals["pageDescription"] = cfg.App.Page.Description
	app.Locals["lang"] = cfg.App.Page.Lang
	app.Locals["direction"] = cfg.App.Page.Direction
	app.Locals["year"] = fmt.Sprint(time.Now().Year())

	// Create standard routes.

	app.Get("/site.yaml", secure.AdminAuth, func(req *f.Request, res *f.Response, next func()) {
		res.ContentType("text/plain")
		res.Send(cfg.String())
	})

	app.Get("/admin", secure.AdminAuth, func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxmode"] = "admin"
		req.Params["juxcomp"] = cfg.App.Defaults.AdminComponent
		req.Params["juxview"] = cfg.App.Defaults.AdminComponentView
		Render(req, res, next, cfg, app)
	})

	app.Get("/admin/:juxcomp", secure.AdminAuth, func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxmode"] = "admin"
		req.Params["juxview"] = cfg.App.Defaults.AdminComponentView
		Render(req, res, next, cfg, app)
	})

	app.Get("/admin/:juxcomp/:juxview", secure.AdminAuth, func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxmode"] = "admin"
		Render(req, res, next, cfg, app)
	})

	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxmode"] = "public"
		req.Params["juxcomp"] = cfg.App.Defaults.Component
		req.Params["juxview"] = cfg.App.Defaults.ComponentView
		Render(req, res, next, cfg, app)
	})

	app.Get("/:juxcomp", func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxmode"] = "public"
		req.Params["juxview"] = cfg.App.Defaults.ComponentView
		Render(req, res, next, cfg, app)
	})

	app.Get("/:juxcomp/:juxview", func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxmode"] = "public"
		Render(req, res, next, cfg, app)
	})

	app.Get("*", func(req *f.Request, res *f.Response, next func()) {
		req.Query["juxskip"] = "1"
		req.Params["juxcomp"] = "404"
		Render(req, res, next, cfg, app)
	})

	http.Handle("/", app)
}
