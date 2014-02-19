package jux_article

import (
	"github.com/ricallinson/forgery"
	// "github.com/ricallinson/jux"
	// "github.com/ricallinson/jux/helpers/datastore"
)

// Route the request to the correct handler function.
func HandlerAdmin(req *f.Request, res *f.Response, next func()) {
	switch req.Params["juxview"] {
	default:
		admin_list_articles(req, res, next)
	case "edit":
		admin_edit_article(req, res, next)
	case "read":
		admin_save_article(req, res, next)
	}
}

func admin_list_articles(req *f.Request, res *f.Response, next func()) {
	res.End("Admin List")
}

func admin_edit_article(req *f.Request, res *f.Response, next func()) {
	res.End("Admin Edit")
}

func admin_save_article(req *f.Request, res *f.Response, next func()) {
	res.End("Admin Save")
}
