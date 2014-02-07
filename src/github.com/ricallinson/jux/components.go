package jux

import (
	"github.com/ricallinson/jux/components"
	"github.com/ricallinson/jux/components/article"
	"github.com/ricallinson/jux/components/notfound"
)

// Register all default components.
func registerDefaultComponents(cfg *AppCfg) {
	cfg.RegisterComponent("error", notfound.Handler)
	cfg.RegisterComponent("default", components.DefaultHandler)
	cfg.RegisterComponent("article", article.Handler)
	cfg.RegisterComponent("a", components.AHandler)
	cfg.RegisterComponent("b", components.BHandler)
	cfg.RegisterComponent("c", components.CHandler)
}
