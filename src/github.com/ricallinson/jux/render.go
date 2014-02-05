package jux

import (
	"github.com/ricallinson/fcomposite"
	"github.com/ricallinson/forgery"
)

func Render(req *f.Request, res *f.Response, next func(), cfg *AppCfg) {

	composite := fcomposite.Map{
		"position-1": cfg.Components["a"],
		// "position-2":  func(),
		"position-3": cfg.Components[req.Params["juxcomp"]],
		// "position-4":  func(),
		"position-5": cfg.Components["c"],
	}

	data := composite.Dispatch(req, res, next)

	res.Render("themes/"+cfg.Defaults.Theme+"/index.html", data)
}
