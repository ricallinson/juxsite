package publictheme

import (
	"github.com/ricallinson/forgery"
	"github.com/ricallinson/jux/assets"
)

func Template(req *f.Request, res *f.Response, next func()) {
	asset := assets.New(res.Locals)
	asset.AddCss("/publictheme/bootstrap/css/bootstrap.min.css")
	asset.AddCss("/publictheme/css/screen.css")
	asset.Render()
	res.Render("publictheme/main.html")
}

func FourOFour(req *f.Request, res *f.Response, next func()) {
	asset := assets.New(res.Locals)
	asset.AddCss("/publictheme/bootstrap/css/bootstrap.min.css")
	asset.AddCss("/publictheme/css/screen.css")
	asset.Render()
	res.Render("publictheme/404.html")
}
