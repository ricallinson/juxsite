package components

import (
	"github.com/ricallinson/forgery"
	"github.com/ricallinson/jux/assets"
)

func DefaultHandler(req *f.Request, res *f.Response, next func()) {
	asset := assets.Make(res.Locals)
	asset.AddCss("default/bootstrap/css/bootstrap.css")
	asset.AddCss("default/css/screen.css")
	asset.Render()
	res.Render("default/main.html")
}
