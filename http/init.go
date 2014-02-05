package main

import (
	"github.com/ricallinson/jux"
)

func init() {
	cfg := &jux.AppCfg{}
	cfg.Load("site.yaml")
	jux.Start(cfg)
}
