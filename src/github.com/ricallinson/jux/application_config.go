package jux

import (
	"github.com/ricallinson/forgery"
)

// A Map of string IDs to forgery functions.
type Components map[string]func(*f.Request, *f.Response, func())

type AppCfg struct {
	// Page configuration.
	Page struct {
		BaseUrl     string
		Name        string
		Description string
		Lang        string
		Direction   string
	}

	// Default values at the application level.
	Defaults struct {
		Env           string
		Debug         bool
		Theme         string
		Component     string
		ComponentView string
	}

	// Map of all available components.
	Components Components

	// Map of all usable layouts.
	Layouts map[string]map[string][]string
}

func (this *AppCfg) Load(file string) {

	this.Page.BaseUrl = "/"
	this.Page.Name = "Jux"
	this.Page.Description = ""
	this.Page.Lang = "en"
	this.Page.Direction = "ltr"

	this.Defaults.Debug = false
	this.Defaults.Env = "development"
	this.Defaults.Theme = "default"
	this.Defaults.Component = "article"
	this.Defaults.ComponentView = "main"

	// Instantiate the Map of layouts.
	this.Layouts = map[string]map[string][]string{}

	// Create the default layout.
	this.Layouts["default"] = map[string][]string{
		// "position-01": {"a"}, // Bread crumbs.
		"position-03": {"a", "b"}, // Menu and Login.
		"position-04": {"f"},      // Sample error.
	}

	// Instantiate the Map of all available components.
	this.Components = Components{}
}

func (this *AppCfg) Save(file string) {

}

func (this *AppCfg) RegisterComponent(name string, fn func(*f.Request, *f.Response, func())) {
	this.Components[name] = fn
}
