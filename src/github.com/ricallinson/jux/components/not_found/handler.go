package not_found

import (
	"github.com/ricallinson/forgery"
)

// Standard not found handler.
func Handler(req *f.Request, res *f.Response, next func()) {
	res.Render("not_found/main.html", map[string]string{
		"error": "Component not found.",
	})
}
