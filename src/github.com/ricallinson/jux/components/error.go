package components

import (
	"github.com/ricallinson/forgery"
)

func ErrorHandler(req *f.Request, res *f.Response, next func()) {
	res.Render("error/main.html", map[string]string{
		"error": "Component not found.",
	})
}
