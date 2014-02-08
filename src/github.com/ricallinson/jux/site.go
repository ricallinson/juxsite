package jux

import (
	"github.com/ricallinson/forgery"
	"path"
)

type Site struct {
	// Map of all available components.
	Components map[string]func(*f.Request, *f.Response, func())

	// Map of all the loaded YAML files.
	Configs map[string][]byte

	// The Site configuration.
	Config *Config
}

// Public method to start the site listening for requests.
func (this *Site) Listen() {
	if this.Config == nil {
		this.Config = &Config{}
	}
	this.GetConfig("site", &this.Config)
	listen(this)
}

// Registers a new component with the application.
func (this *Site) RegisterComponent(name string, fn func(*f.Request, *f.Response, func())) {
	if this.Components == nil {
		// Instantiate the Map of all available components.
		this.Components = map[string]func(*f.Request, *f.Response, func()){}
	}
	this.Components[name] = fn
}

// Maps the config file for the given component name into the given instance.
func (this *Site) GetConfig(name string, i interface{}) {
	// Map the YAML to the given interface.
	FromYaml(this.GetConfigYaml(name), i)
}

// Reads the config file for the given component name.
func (this *Site) GetConfigYaml(name string) []byte {
	// Instantiate the Map of all loaded configs.
	if this.Configs == nil {
		this.Configs = map[string][]byte{}
	}
	// If the file has not been read then cache it.
	if _, ok := this.Configs[name]; ok == false {
		filepath := path.Join("config", name+".yaml")
		this.Configs[name] = FromFile(filepath)
	}
	return this.Configs[name]
}

// Returns a copy of the matched component or the "not_found" Component.
func (this *Site) GetComponent(name string) func(*f.Request, *f.Response, func()) {
	if component, ok := this.Components[name]; ok {
		return component
	}
	return this.Components["not_found"]
}

// Returns a copy of the matched layout or an empty map.
func (this *Site) GetLayout(name string) map[string][]string {
	layout := map[string][]string{}
	if _, ok := this.Config.Layouts[name]; ok {
		for position, _ := range this.Config.Layouts[name] {
			layout[position] = this.Config.Layouts[name][position]
		}
	}
	return layout
}
