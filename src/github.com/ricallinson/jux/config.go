package jux

import (
	"github.com/ricallinson/forgery"
	"io/ioutil"
	"launchpad.net/goyaml"
)

type AppCfg struct {
	App struct {
		// Page level configuration keys.
		Page struct {
			BaseUrl     string
			Name        string
			Description string
			Lang        string
			Direction   string
		}

		// Application level configuration keys.
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

// Registers a new component with the application.
func (this *AppCfg) RegisterComponent(name string, fn func(*f.Request, *f.Response, func())) {
	this.Components[name] = fn
}

// Returns a copy of the matched layout or an empty map.
func (this *AppCfg) GetLayout(name string) map[string][]string {
	layout := map[string][]string{}
	if _, ok := this.App.Layouts[name]; ok {
		for position, _ := range this.App.Layouts[name] {
			layout[position] = this.App.Layouts[name][position]
		}
	}
	return layout
}

// Populates the defaults for the application configuration.
func (this *AppCfg) init() {

	// Defaults for the page level configuration.
	this.App.Page.BaseUrl = "/"
	this.App.Page.Name = "Jux"
	this.App.Page.Description = ""
	this.App.Page.Lang = "en"
	this.App.Page.Direction = "ltr"

	// Defaults for the application level configuration.
	this.App.Defaults.Debug = false
	this.App.Defaults.Env = "development"
	this.App.Defaults.Theme = "public_theme" // this is a component
	this.App.Defaults.Component = "article"
	this.App.Defaults.ComponentView = "main"
	this.App.Defaults.AdminTheme = "admin_theme" // this is a component
	this.App.Defaults.AdminComponent = "dashboard"
	this.App.Defaults.AdminComponentView = "main"

	// Instantiate the Map of layouts.
	this.App.Layouts = map[string]map[string][]string{}

	// Create the default "public" layout.
	this.App.Layouts["public"] = map[string][]string{
		"position-03": {"link_menu", "article_menu"}, // Menu and Login.
	}

	// Create the default "admin" layout.
	this.App.Layouts["admin"] = map[string][]string{
		"position-01": {"f"}, // Sample error.
	}

	// Instantiate the Map of all available components.
	this.Components = map[string]func(*f.Request, *f.Response, func()){}

	// Register any default components.
	registerDefaultComponents(this)
}
