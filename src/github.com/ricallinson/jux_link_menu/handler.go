package jux_link_menu

import (
	"github.com/ricallinson/forgery"
	"github.com/ricallinson/jux"
	"strings"
)

func Handler(req *f.Request, res *f.Response, next func()) {

	cfg := Config{}
	jux.GetSite(req).GetConfig("jux_link_menu", &cfg)

	parts := []string{jux.GetPosition(req), req.Params[jux.JuxComp], req.Params[jux.JuxView], "*"}

	for i := 4; i > 0; i-- {
		key := strings.Join(parts[i:4], ":")
		if menu, ok := cfg[key]; ok {
			res.Render("jux_link_menu/main.html", menu)
			return
		}
	}

	res.End("")
}
