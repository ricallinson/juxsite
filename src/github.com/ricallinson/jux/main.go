package jux

import (
	"fmt"
	"github.com/ricallinson/fmustache"
	"github.com/ricallinson/forgery"
	"github.com/ricallinson/jux/helpers/secure"
	"net/http"
	"time"
)

// Maps the Jux site to the Forgery web framework.
func listen(site *Site) {

	// Create a Forgery Server.
	app := f.CreateServer()

	// Add Starkr middleware.
	app.Use(f.ResponseTime())
	app.Use(f.Favicon())
	app.Use(f.Static())

	// Set the render engine.
	app.Engine(".html", fmustache.Make())

	// Set the Stackr environment.
	app.Env = site.Config.Defaults.Env

	// Set template locals.
	app.Locals["baseUrl"] = site.Config.Page.BaseUrl
	app.Locals["siteName"] = site.Config.Page.Name
	app.Locals["siteDescription"] = site.Config.Page.Description
	app.Locals["pageTitle"] = site.Config.Page.Name
	app.Locals["pageDescription"] = site.Config.Page.Description
	app.Locals["lang"] = site.Config.Page.Lang
	app.Locals["direction"] = site.Config.Page.Direction
	app.Locals["year"] = fmt.Sprint(time.Now().Year())

	// Create standard routes.

	app.Get("/:name.yaml", secure.AdminAuth, func(req *f.Request, res *f.Response, next func()) {
		res.ContentType("text/plain")
		res.Send(site.GetConfigYaml(req.Params["name"]))
	})

	app.Get("/admin", secure.AdminAuth, func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxmode"] = "admin"
		req.Params["juxcomp"] = site.Config.Defaults.AdminComponent
		req.Params["juxview"] = site.Config.Defaults.AdminComponentView
		req.Map["juxsite"] = site
		Render(req, res, next)
	})

	app.Get("/admin/:juxcomp", secure.AdminAuth, func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxmode"] = "admin"
		req.Params["juxview"] = site.Config.Defaults.AdminComponentView
		req.Map["juxsite"] = site
		Render(req, res, next)
	})

	app.Get("/admin/:juxcomp/:juxview", secure.AdminAuth, func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxmode"] = "admin"
		req.Map["juxsite"] = site
		Render(req, res, next)
	})

	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxmode"] = "public"
		req.Params["juxcomp"] = site.Config.Defaults.Component
		req.Params["juxview"] = site.Config.Defaults.ComponentView
		req.Map["juxsite"] = site
		Render(req, res, next)
	})

	app.Get("/:juxcomp", func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxmode"] = "public"
		req.Params["juxview"] = site.Config.Defaults.ComponentView
		req.Map["juxsite"] = site
		Render(req, res, next)
	})

	app.Get("/:juxcomp/:juxview", func(req *f.Request, res *f.Response, next func()) {
		req.Params["juxmode"] = "public"
		req.Map["juxsite"] = site
		Render(req, res, next)
	})

	app.Get("*", func(req *f.Request, res *f.Response, next func()) {
		req.Query["juxskip"] = "1"
		req.Params["juxcomp"] = "public_404"
		req.Map["juxsite"] = site
		Render(req, res, next)
	})

	// Handle the application.
	http.Handle("/", app)
}
