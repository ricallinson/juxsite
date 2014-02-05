package main

import (
	"github.com/ricallinson/fmustache"
	"github.com/ricallinson/forgery"
	"net/http"
)

func init() {

	app := f.CreateServer()

	app.Use(f.ResponseTime())
	app.Use(f.Favicon())
	app.Use(f.Static())

	app.Engine(".html", fmustache.Make())

	app.Locals["title"] = "Jux - the Content Managment System"

	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		res.Render("index.html")
	})

	http.Handle("/", app)
}
