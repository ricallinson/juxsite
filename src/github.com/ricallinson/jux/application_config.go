package jux

import(
	"github.com/ricallinson/forgery"
)

type Components map[string]func(*f.Request, *f.Response, func())

type AppCfg struct {

	Page struct {
		Title string
		Theme string
		Lang string
	}

	DefaultComponent string
	DefaultComponentView string

	// Map of all available components.
	Components Components
}

func (this *AppCfg) Load(file string) {

	this.Page.Title = "Jux"
	this.Page.Theme = "default"
	this.Page.Lang = "en-gb"
	this.Page.Direction = "ltr"

	this.DefaultComponent = "article"
	this.DefaultComponentView = "main"

	// Instantiate the Map of all available components.
	this.Components = Components{}
}

func (this *AppCfg) Save(file string) {

}

func (this *AppCfg) RegisterComponent(name string, fn func(*f.Request, *f.Response, func())) {
	this.Components[name] = fn
}
