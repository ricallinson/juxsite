package jux

import (
	"github.com/ricallinson/forgery"
)

// A Map of string IDs to forgery functions.
type Components map[string]func(*f.Request, *f.Response, func())

type AppCfg struct {
	// Page level config.
	Site struct {
		BaseUrl     string
		Name        string
		Description string
		Lang        string
		Direction   string
	}

	// Default values at the app level.
	Defaults struct {
		Theme         string
		Component     string
		ComponentView string
	}

	// Map of all available components.
	Components Components
}

func (this *AppCfg) Load(file string) {

	this.Site.BaseUrl = "/"
	this.Site.Name = "Jux"
	this.Site.Description = ""
	this.Site.Lang = "en-gb"
	this.Site.Direction = "ltr"

	this.Defaults.Theme = "default"
	this.Defaults.Component = "article"
	this.Defaults.ComponentView = "main"

	// Instantiate the Map of all available components.
	this.Components = Components{}
}

func (this *AppCfg) Save(file string) {

}

func (this *AppCfg) RegisterComponent(name string, fn func(*f.Request, *f.Response, func())) {
	this.Components[name] = fn
}
