package jux

import(
	"github.com/ricallinson/fcomposite"
	"github.com/ricallinson/forgery"
)

func Render(req *f.Request, res *f.Response, next func(), cfg *AppCfg) {

	composite := fcomposite.Map{
        "header": cfg.Components["a"],
        "body": cfg.Components["b"],
        "footer": cfg.Components["c"],
    }

    data := composite.Dispatch(req, res, next)

	res.Render("themes/" + cfg.Page.Theme + "/index.html", data)
}