package jux_link_menu

import (
	"github.com/ricallinson/forgery"
	"github.com/ricallinson/jux"
)

func Handler(req *f.Request, res *f.Response, next func()) {
	// Create a Config instance.
	cfg := &Config{}
	// Load the YAML configuration into the Config instance.
	jux.GetConfig(req).Get("jux_link_menu", cfg)
	// Create a slice to hold the links.
	menu := Menu{}
	// Look for a match in the configuration we loaded.
	if match, ok := cfg.Match[req.Params["juxcomp"]+":"+req.Params["juxview"]]; ok {
		// If the "juxcomp" and "juxview" match use that.
		menu = match
	} else if match, ok := cfg.Match[req.Params["juxcomp"]]; ok {
		// If just the "juxcomp" matches use that.
		menu = match
	} else if match, ok := cfg.Match["*"]; ok {
		// Otherwise if there is a "*" use that.
		menu = match
	} else {
		// Otherwise there are no mathces so res.End() and return.
		res.End("")
		return
	}
	// res.Send(jux.GetConfig(req).ToYaml(menu.Links))
	// return
	// If we got here, render the links.
	res.Render("jux_link_menu/main.html",
		map[string]string{
			"title": menu.Title,
		},
		map[string][]Link{
			"links": menu.Links,
		})
}
