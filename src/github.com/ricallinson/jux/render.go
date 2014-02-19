package jux

import (
	"github.com/ricallinson/fcomposite"
	"github.com/ricallinson/forgery"
	"strconv"
)

const (
	PositionSize int = 11
	JuxSite          = "juxsite"
	JuxSkip          = "juxskip"
	JuxMode          = "juxmode"
	JuxComp          = "juxcomp"
	JuxView          = "juxview"
)

// Given a map select the layout to use and return an fcomposite.Map{}.
func selectLayout(params map[string]string, site *Site) fcomposite.Map {

	// Based on "juxmode" we need to pick a layout.
	layout := site.GetLayout(params[JuxMode])
	// If we were given a component then replace "maincontent" with it.
	if len(params[JuxComp]) > 0 {
		layout["maincontent"] = []string{params[JuxComp]}
	}
	composite := fcomposite.Map{}
	count := 0

	// Walk over the layout and fill the fcomposite.Map{}.
	for pos, names := range layout {
		for _, name := range names {
			composite[pos+strconv.Itoa(count)] = site.GetComponent(name)
			count++
		}
	}
	return composite
}

// Given a Config object select the theme to use and return it as a string.
func selectTheme(mode string, config *Config) string {

	// Based on "juxmode" we need to pick a theme.
	if mode == "admin" {
		return config.Defaults.AdminTheme
	}

	return config.Defaults.PublicTheme
}

// Collapse the items positions back into their layout positions.
func collapseLayout(in map[string]string, out map[string]string) {

	for key, val := range in {
		position := key[:PositionSize] // position-01(00)
		if _, ok := out[position]; ok {
			out[position] += val // append
		} else {
			out[position] = val // create
		}
	}
}

// Query parameters that effect the final rendered output.
//
// juxmode - Either "public" or "admin".
// juxcomp - The component to use in layout position "maincontent".
// juxview - The view to use with the layout position "maincontent".
// juxskip - Skip rendering the layout and render only the requested "juxcomp" & "juxview".
//
// Render the requested layout wrapped using a theme.
func render(req *f.Request, res *f.Response, next func()) {

	// Get the Config from the req.Map.
	site := req.Map[JuxSite].(*Site)

	// If "juxskip" is set then just reander the requested component.
	if _, ok := req.Query[JuxSkip]; ok {
		site.GetComponent(req.Params[JuxComp])(req, res, next)
		return
	}

	// Select a layout to use for this request.
	composite := selectLayout(req.Params, site)

	// Dispatch the fcomposite.Map{}.
	data := composite.Dispatch(req, res, next)

	// Collapse the dispathed layout in the Locals map.
	collapseLayout(data, res.Locals)

	// Select the theme to use.
	theme := selectTheme(req.Params[JuxMode], site.Config)

	// Render the final component.
	site.Components[theme](req, res, next)
}
