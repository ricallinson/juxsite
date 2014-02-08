package main

import (
	"github.com/ricallinson/jux"
	"github.com/ricallinson/jux/components/not_found"
	"github.com/ricallinson/jux/components/public_theme"
	"github.com/ricallinson/jux_article"
	"github.com/ricallinson/jux_link_menu"
)

func init() {

	// Create a Jux Configuration.
	cfg := &jux.Config{}

	// Load the localconfiguration file.
	cfg.Load("site.yaml")

	// Register components.
	cfg.RegisterComponent("notfound", not_found.Handler)
	cfg.RegisterComponent("article", jux_article.Handler)
	cfg.RegisterComponent("article_menu", jux_article.Menu)
	cfg.RegisterComponent("link_menu", jux_link_menu.Handler)
	cfg.RegisterComponent("public_theme", public_theme.Template)
	cfg.RegisterComponent("public_404", public_theme.FourOFour)

	jux.Start(cfg)
}
