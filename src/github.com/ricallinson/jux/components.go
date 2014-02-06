package jux

import (
	"github.com/ricallinson/jux/components"
)

// Register all default components.
func registerDefaultComponents(cfg *AppCfg) {
	cfg.RegisterComponent("error", components.ErrorHandler)
	cfg.RegisterComponent("default", components.DefaultHandler)
	cfg.RegisterComponent("article", components.ArticleHandler)
	cfg.RegisterComponent("a", components.AHandler)
	cfg.RegisterComponent("b", components.BHandler)
	cfg.RegisterComponent("c", components.CHandler)
}
