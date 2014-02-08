package components

import (
	"github.com/ricallinson/forgery"
)

func BHandler(req *f.Request, res *f.Response, next func()) {
	res.Render("b/main.html", map[string]string{
		"title": "Login Form",
		"msg":   "B",
	})
}

func CHandler(req *f.Request, res *f.Response, next func()) {
	res.Render("c/main.html", map[string]string{
		"msg": "C",
	})
}
