package jux

import(
	"github.com/ricallinson/fcomposite"
	"github.com/ricallinson/forgery"
)

func Render(req *f.Request, res *f.Response, next func(), cfg *AppCfg) {

	composite := fcomposite.Map{
        "header": func(req *f.Request, res *f.Response, next func()) {
            res.Send("Header")
        },
        "body": func(req *f.Request, res *f.Response, next func()) {
            res.Send(req.Params["juxcomp"] + "/" + req.Params["juxview"])
        },
        "footer": func(req *f.Request, res *f.Response, next func()) {
            res.End("Footer")
        },
    }

    data := composite.Dispatch(req, res, next)

	res.Render("themes/" + cfg.PageTheme + "/index.html", data)
}