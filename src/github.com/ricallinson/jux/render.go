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

func Render(req *f.Request, res *f.Response, next func(), cfg *AppCfg, app *f.Server) {

	layout := cfg.Layouts["default"]
	layout["maincontent"] = []string{req.Params["juxcomp"]}
	composite := fcomposite.Map{}
	count := 0

	// Walk over the layout and fill the fcomposite.Map{}.
	for pos, names := range layout {
		for _, name := range names {
			if component, ok := cfg.Components[name]; ok {
				composite[pos+fmt.Sprint(count)] = component
			} else {
				composite[pos+fmt.Sprint(count)] = cfg.Components["error"]
			}
			count++
		}
	}

	// Dispatch the fcomposite.Map{}.
	data := composite.Dispatch(req, res, next)
	build := map[string]string{}

	// Collapse the dispatched positions back into their layout positions.
	for key, val := range data {
		position := key[:11] // position-01(00)
		if _, ok := build[position]; ok {
			build[position] += val // append
		} else {
			build[position] = val // create
		}
	}

	if cfg.Site.Debug {
		build["debug"] = debug(app.Locals, res.Locals, build)
	}

	// Render the final build map into the template.
	res.Render("themes/"+cfg.Defaults.Theme+"/index.html", build)
}
