package main

import (
	"github.com/ricallinson/jux"
	"github.com/ricallinson/jux/components/not_found"
	"github.com/ricallinson/jux/components/public_theme"
	"github.com/ricallinson/jux_article"
)

func init() {
	// Make a new site.
	site := jux.New()
	// Register components.
	site.RegisterComponent("not_found", not_found.Handler)
	site.RegisterComponent("article", jux_article.Handler)
	site.RegisterComponent("article_menu_articles", jux_article.HandlerMenuArticles)
	site.RegisterComponent("article_menu_categories", jux_article.HandlerMenuCategories)
	site.RegisterComponent("public_theme", public_theme.Handler)
	site.RegisterComponent("public_404", public_theme.HandlerFourOFour)
	// Start the application.
	site.Listen()
}
