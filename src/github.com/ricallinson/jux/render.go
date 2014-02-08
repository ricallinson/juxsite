package jux

import (
	"encoding/json"
	"fmt"
	"github.com/ricallinson/fcomposite"
	"github.com/ricallinson/forgery"
)

func debug(maps ...map[string]string) string {
	for i, _ := range maps {
		for k, v := range maps[i] {
			maps[0][k] = v
		}
	}
	debug, _ := json.MarshalIndent(maps[0], "", "    ")
	return string(debug)
}

func Render(req *f.Request, res *f.Response, next func()) {

	// Get the AppCfg from the req.Map.
	cfg := req.Map["cfg"].(*AppCfg)

	// If "juxskip" is set then just reander the requested component.
	if _, ok := req.Query["juxskip"]; ok {
		if component, ok := cfg.Components[req.Params["juxcomp"]]; ok {
			component(req, res, next)
		} else {
			cfg.Components["notfound"](req, res, next)
		}
		return
	}

	// Based on "juxmode" we need to pick a layout.
	layout := cfg.GetLayout(req.Params["juxmode"])
	layout["maincontent"] = []string{req.Params["juxcomp"]}
	composite := fcomposite.Map{}
	count := 0

	// Walk over the layout and fill the fcomposite.Map{}.
	for pos, names := range layout {
		for _, name := range names {
			if component, ok := cfg.Components[name]; ok {
				composite[pos+fmt.Sprint(count)] = component
			} else {
				composite[pos+fmt.Sprint(count)] = cfg.Components["notfound"]
			}
			count++
		}
	}

	// Dispatch the fcomposite.Map{}.
	data := composite.Dispatch(req, res, next)

	// Collapse the dispatched positions back into their layout positions.
	for key, val := range data {
		position := key[:11] // position-01(00)
		if _, ok := res.Locals[position]; ok {
			res.Locals[position] += val // append
		} else {
			res.Locals[position] = val // create
		}
	}

	// Based on "juxmode" we need to pick a theme.
	theme := cfg.App.Defaults.Theme
	if req.Params["juxmode"] == "admin" {
		theme = cfg.App.Defaults.AdminTheme
	}
	// Render the final component.
	cfg.Components[theme](req, res, next)
}
