package jux

import (
	"github.com/ricallinson/jux/components/menu"
	"github.com/ricallinson/jux/components/notfound"
	"github.com/ricallinson/jux/components/publictheme"
	"github.com/ricallinson/jux_article"
)

// Register default components.
func registerDefaultComponents(cfg *AppCfg) {
	cfg.RegisterComponent("notfound", notfound.Handler)
	cfg.RegisterComponent("article", jux_article.Handler)
	cfg.RegisterComponent("article_menu", jux_article.Menu)
	cfg.RegisterComponent("link_menu", menu.Handler)
	cfg.RegisterComponent("public_theme", publictheme.Template)
	cfg.RegisterComponent("public_404", publictheme.FourOFour)
}
