package notfound

import (
	"github.com/ricallinson/forgery"
)

func Handler(req *f.Request, res *f.Response, next func()) {
	res.Render("error/main.html", map[string]string{
		"error": "Component not found.",
	})
}
