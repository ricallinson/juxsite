package public_theme

import (
	"github.com/ricallinson/forgery"
	"github.com/ricallinson/jux/helpers/assets"
)

// Page theme.
func Handler(req *f.Request, res *f.Response, next func()) {
	asset := assets.New(res.Locals)
	asset.AddCss("/public_theme/bootstrap/css/bootstrap.min.css")
	asset.AddCss("/public_theme/css/screen.css")
	asset.Render()
	res.Render("public_theme/main.html")
}

// Fall back 404 page.
func HandlerFourOFour(req *f.Request, res *f.Response, next func()) {
	asset := assets.New(res.Locals)
	asset.AddCss("/public_theme/bootstrap/css/bootstrap.min.css")
	asset.AddCss("/public_theme/css/screen.css")
	asset.Render()
	res.Render("public_theme/404.html")
}
