package jux

import (
	"github.com/ricallinson/forgery"
	"io/ioutil"
	"launchpad.net/goyaml"
)

type AppCfg struct {
	App struct {
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

		// Map of all usable layouts.
		Layouts map[string]map[string][]string
	}

	// Map of all available components.
	Components map[string]func(*f.Request, *f.Response, func())
}

func (this *AppCfg) Load(file string) /*(error)*/ {
	// Prime this AppCfg instance.
	this.init()
	// Read the source file.
	data, err1 := ioutil.ReadFile(file)
	if err1 != nil {
		panic(err1)
		return // error
	}
	// Unmarshal the source into this AppCfg instance.
	// This overrides any default settings from the init() call.
	err2 := goyaml.Unmarshal([]byte(data), &this.App)
	if err2 != nil {
		panic(err2)
		return // error
	}
}

func (this *AppCfg) String() string {
	data, err1 := goyaml.Marshal(this.App)
	if err1 != nil {
		panic(err1)
	}
	return string(data)
}

func (this *AppCfg) RegisterComponent(name string, fn func(*f.Request, *f.Response, func())) {
	this.Components[name] = fn
}

func (this *AppCfg) init() {

	this.App.Page.BaseUrl = "/"
	this.App.Page.Name = "Jux"
	this.App.Page.Description = ""
	this.App.Page.Lang = "en"
	this.App.Page.Direction = "ltr"

	this.App.Defaults.Debug = false
	this.App.Defaults.Env = "development"
	this.App.Defaults.Theme = "default"
	this.App.Defaults.Component = "article"
	this.App.Defaults.ComponentView = "main"

	// Instantiate the Map of layouts.
	this.App.Layouts = map[string]map[string][]string{}

	// Create the default layout.
	this.App.Layouts["default"] = map[string][]string{
		// "position-01": {"a"}, // Bread crumbs.
		"position-03": {"a", "b"}, // Menu and Login.
		"position-04": {"f"},      // Sample error.
	}

	// Instantiate the Map of all available components.
	this.Components = map[string]func(*f.Request, *f.Response, func()){}

	// Register any default components.
	registerDefaultComponents(this)
}