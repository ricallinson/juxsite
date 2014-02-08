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
			Env                string
			Debug              bool
			Theme              string
			Component          string
			ComponentView      string
			AdminTheme         string
			AdminComponent     string
			AdminComponentView string
		}

		// Map of all usable layouts.
		Layouts map[string]map[string][]string
	}

	// Map of all available components.
	Components map[string]func(*f.Request, *f.Response, func())
}

// Load the configuration from the give YAML file.
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

// Return the configuration as a YAML string.
func (this *AppCfg) String() string {
	data, err1 := goyaml.Marshal(this.App)
	if err1 != nil {
		panic(err1)
	}
	return string(data)
}

// Register a new component.
func (this *AppCfg) RegisterComponent(name string, fn func(*f.Request, *f.Response, func())) {
	this.Components[name] = fn
}

// returns a copy of the matched layout or an empty map.
func (this *AppCfg) GetLayout(name string) map[string][]string {
	layout := map[string][]string{}
	if _, ok := this.App.Layouts[name]; ok {
		for position, _ := range this.App.Layouts[name] {
			layout[position] = this.App.Layouts[name][position]
		}
	}
	return layout
}

// Populates the defaults for the configuration.
func (this *AppCfg) init() {

	this.App.Page.BaseUrl = "/"
	this.App.Page.Name = "Jux"
	this.App.Page.Description = ""
	this.App.Page.Lang = "en"
	this.App.Page.Direction = "ltr"

	this.App.Defaults.Debug = false
	this.App.Defaults.Env = "development"
	this.App.Defaults.Theme = "publictheme"
	this.App.Defaults.Component = "article"
	this.App.Defaults.ComponentView = "main"
	this.App.Defaults.AdminTheme = "admintheme"
	this.App.Defaults.AdminComponent = "dashboard"
	this.App.Defaults.AdminComponentView = "main"

	// Instantiate the Map of layouts.
	this.App.Layouts = map[string]map[string][]string{}

	// Create the default "public" layout.
	this.App.Layouts["public"] = map[string][]string{
		// "position-01": {""}, // Bread crumbs.
		"position-03": {"a", "article_menu"}, // Menu and Login.
		// "position-04": {"f"},                      // Sample error.
	}

	// Create the default "admin" layout.
	this.App.Layouts["admin"] = map[string][]string{
	// "position-01": {"f"}, // Sample error.
	}

	// Instantiate the Map of all available components.
	this.Components = map[string]func(*f.Request, *f.Response, func()){}

	// Register any default components.
	registerDefaultComponents(this)
}
