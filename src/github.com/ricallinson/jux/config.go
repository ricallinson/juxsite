package jux

import (
	"github.com/ricallinson/forgery"
	"path"
)

type Config struct {
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

// Load the application configuration from the given YAML file.
func (this *Config) Load() /*(error)*/ {
	filename := path.Join("config", "site.yaml")
	FromYamlFile(filename, &this.App)
}

// Registers a new component with the application.
func (this *Config) RegisterComponent(name string, fn func(*f.Request, *f.Response, func())) {
	if this.Components == nil {
		// Instantiate the Map of all available components.
		this.Components = map[string]func(*f.Request, *f.Response, func()){}
	}
	this.Components[name] = fn
}

// Read the config file for the given component name.
func (this *Config) Get(name string, i interface{}) {
	filepath := path.Join("config", name+".yaml")
	FromYamlFile(filepath, i)
}

// Returns a copy of the matched component or the Not_Found Component.
func (this *Config) GetComponent(name string) func(*f.Request, *f.Response, func()) {
	if component, ok := this.Components[name]; ok {
		return component
	}
	return this.Components["not_found"]
}

// Returns a copy of the matched layout or an empty map.
func (this *Config) GetLayout(name string) map[string][]string {
	layout := map[string][]string{}
	if _, ok := this.App.Layouts[name]; ok {
		for position, _ := range this.App.Layouts[name] {
			layout[position] = this.App.Layouts[name][position]
		}
	}
	return layout
}

// Return the configuration as a YAML string.
func (this *Config) String() string {
	return string(ToYaml(this.App))
}
