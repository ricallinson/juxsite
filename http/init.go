package main

import (
	"github.com/ricallinson/jux"
	"github.com/ricallinson/jux/components/not_found"
	"github.com/ricallinson/jux/components/public_theme"
	"github.com/ricallinson/jux_article"
	"github.com/ricallinson/jux_link_menu"
)

func init() {
	// Make a new site.
	site := jux.New()
	// Register components.
	site.RegisterComponent("not_found", not_found.Handler)
	site.RegisterComponent("article", jux_article.Handler)
	site.RegisterComponent("article_menu", jux_article.Menu)
	site.RegisterComponent("link_menu", jux_link_menu.Handler)
	site.RegisterComponent("public_theme", public_theme.Template)
	site.RegisterComponent("public_404", public_theme.FourOFour)
	// Start the application.
	site.Listen()
}
