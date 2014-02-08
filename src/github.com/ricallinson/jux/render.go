package jux

import (
	"github.com/ricallinson/fcomposite"
	"github.com/ricallinson/forgery"
	"strconv"
)

// Query parameters that effect the final rendered output.
//
// juxmode - Either "public" or "admin".
// juxcomp - The component to use in layout position "maincontent".
// juxview - The view to use with the layout position "maincontent".
// juxskip - Skip rendering the layout and render only the requested "juxcomp" & "juxview".
//
// Render the requested layout wrapped using a theme.
func Render(req *f.Request, res *f.Response, next func()) {

	// Get the Config from the req.Map.
	site := req.Map["juxsite"].(*Site)

	// If "juxskip" is set then just reander the requested component.
	if _, ok := req.Query["juxskip"]; ok {
		site.GetComponent(req.Params["juxcomp"])(req, res, next)
		return
	}

	// Based on "juxmode" we need to pick a layout.
	layout := site.GetLayout(req.Params["juxmode"])
	layout["maincontent"] = []string{req.Params["juxcomp"]}
	composite := fcomposite.Map{}
	count := 0

	// Walk over the layout and fill the fcomposite.Map{}.
	for pos, names := range layout {
		for _, name := range names {
			composite[pos+strconv.Itoa(count)] = site.GetComponent(name)
			count++
		}
	}

	// Dispatch the fcomposite.Map{}.
	data := composite.Dispatch(req, res, next)

	// Collapse the dispatched positions back into their layout positions.
	for key, val := range data {
		position := key[:11] // position-01(00)
		if _, ok := res.Locals[position]; ok {
			res.Locals[position] += val // append
		} else {
			res.Locals[position] = val // create
		}
	}

	// Based on "juxmode" we need to pick a theme.
	theme := site.Config.Defaults.Theme
	if req.Params["juxmode"] == "admin" {
		theme = site.Config.Defaults.AdminTheme
	}

	// Render the final component.
	site.Components[theme](req, res, next)
}
