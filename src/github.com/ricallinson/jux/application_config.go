package jux

import(
	"github.com/ricallinson/forgery"
)

type Components map[string]func(*f.Request, *f.Response, func())

type AppCfg struct {

	PageTitle string
	PageTheme string

	DefaultComponent string
	DefaultComponentView string

	// Map of all available components.
	Components Components
}

func (this *AppCfg) Load(file string) {

	// Instantiate the Map of all available components.
	this.Components = Components{}

	this.PageTitle = "Jux - the Content Managment System"
	this.PageTheme = "default"

	this.DefaultComponent = "article"
	this.DefaultComponentView = "main"

	// Register all default components.

	this.RegisterComponent("a", func(req *f.Request, res *f.Response, next func()) {
	    res.Send("Header")
	})

	this.RegisterComponent("b", func(req *f.Request, res *f.Response, next func()) {
	    res.Send(req.Params["juxcomp"] + "/" + req.Params["juxview"])
	})

	this.RegisterComponent("c", func(req *f.Request, res *f.Response, next func()) {
	    res.End("Footer")
	})
}

func (this *AppCfg) RegisterComponent(name string, fn func(*f.Request, *f.Response, func())) {
	this.Components[name] = fn
}
