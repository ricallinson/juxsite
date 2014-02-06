package secure

import (
	"appengine"
	"appengine/user"
	"github.com/ricallinson/forgery"
)

/*
   Protect a URL by forcing a login.
*/
func UserAuth(req *f.Request, res *f.Response, next func()) {
	context := appengine.NewContext(req.Request.Request)
	u := user.Current(context)
	if u == nil {
		url, _ := user.LoginURL(context, "/")
		res.Redirect(url)
		return
	}
	context.Debugf("User: %v", u.String())
	next()
}

/*
   Protect a URL by forcing a login.
*/
func AdminAuth(req *f.Request, res *f.Response, next func()) {
	context := appengine.NewContext(req.Request.Request)
	u := user.Current(context)
	if u == nil || user.IsAdmin(context) == false {
		url, _ := user.LoginURL(context, "/")
		res.Redirect(url)
		return
	}
	context.Debugf("User: %v", u.String())
	next()
}
