package jux

import (
	"github.com/ricallinson/jux/components"
	"github.com/ricallinson/jux/components/article"
	"github.com/ricallinson/jux/components/notfound"
	"github.com/ricallinson/jux/components/publictheme"
)

// Register all default components.
func registerDefaultComponents(cfg *AppCfg) {
	cfg.RegisterComponent("notfound", notfound.Handler)
	cfg.RegisterComponent("publictheme", publictheme.Template)
	cfg.RegisterComponent("404", publictheme.FourOFour)
	cfg.RegisterComponent("article", article.Handler)
	cfg.RegisterComponent("article_menu", article.Menu)
	cfg.RegisterComponent("a", components.AHandler)
	cfg.RegisterComponent("b", components.BHandler)
	cfg.RegisterComponent("c", components.CHandler)
}
