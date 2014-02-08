package assets

import (
	"strings"
)

// The Assets handler.
type Assets struct {
	store map[string]string
}

// Returns a new assets handler using the given "store".
func New(store map[string]string) *Assets {
	return &Assets{store}
}

// Add a CSS URI to optional an location.
// Default location is "top".
func (this *Assets) AddCss(uri string, loc ...string) {
	location := "top"
	if len(loc) == 1 {
		location = loc[0]
	}
	this.AddAsset(uri, "css", location)
}

// Add a JS URI to optional an location.
// Default location is "bottom".
func (this *Assets) AddJs(uri string, loc ...string) {
	location := "bottom"
	if len(loc) == 1 {
		location = loc[0]
	}
	this.AddAsset(uri, "js", location)
}

// Add a BLOB string to optional an location.
// Default location is "bottom".
func (this *Assets) AddBlob(blob string, loc ...string) {
	location := "bottom"
	if len(loc) == 1 {
		location = loc[0]
	}
	this.AddAsset(blob, "", location)
}

// Add string asset with a fixed extension to the given location.
// Default location is "bottom".
func (this *Assets) AddAsset(asset, ext, location string) {
	if len(ext) > 0 {
		asset += ","
	}
	this.storeAsset("blob-"+location+"-"+ext, asset)
}

// Render the predefined assets to keys in the predefined "store".
func (this *Assets) Render() {
	this.renderCss()
	this.renderJs()
}

// Appends the given assets to the given key in the predefined "store".
func (this *Assets) storeAsset(key, asset string) {
	if _, ok := this.store[key]; ok {
		this.store[key] += asset
	} else {
		this.store[key] = asset
	}
}

// Renders the stored CSS to the key "blob-top" in the predefined "store".
func (this *Assets) renderCss() {
	for _, uri := range strings.Split(this.store["blob-top-css"], ",") {
		if len(uri) > 0 {
			link := "<link rel=\"stylesheet\" href=\"" + uri + "\" type=\"text/css\" />"
			this.storeAsset("blob-top", link)
		}
	}
}

// Renders the stored JS to the key "blob-bottom" in the predefined "store".
func (this *Assets) renderJs() {
	for _, uri := range strings.Split(this.store["blob-bottom-js"], ",") {
		if len(uri) > 0 {
			link := ""
			this.storeAsset("blob-bottom", link)
		}
	}
}
