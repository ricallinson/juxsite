package assets

import (
	"strings"
)

type Assets struct {
	store map[string]string
}

func Make(store map[string]string) *Assets {
	return &Assets{store}
}

func (this *Assets) AddCss(uri string, loc ...string) {
	location := "top"
	if len(loc) == 1 {
		location = loc[0]
	}
	this.AddAsset(uri, "css", location)
}

func (this *Assets) AddJs(uri string, loc ...string) {
	location := "bottom"
	if len(loc) == 1 {
		location = loc[0]
	}
	this.AddAsset(uri, "js", location)
}

func (this *Assets) AddBlob(blob string, loc ...string) {
	location := "bottom"
	if len(loc) == 1 {
		location = loc[0]
	}
	this.AddAsset(blob, "", location)
}

func (this *Assets) AddAsset(asset, ext, location string) {
	if len(ext) > 0 {
		asset += ","
	}
	this.storeAsset("blob-"+location+"-"+ext, asset)
}

func (this *Assets) Render() {
	this.renderCss()
	this.renderJs()
}

func (this *Assets) storeAsset(key, asset string) {
	if _, ok := this.store[key]; ok {
		this.store[key] += asset
	} else {
		this.store[key] = asset
	}
}

func (this *Assets) renderCss() {
	for _, uri := range strings.Split(this.store["blob-top-css"], ",") {
		if len(uri) > 0 {
			link := "<link rel=\"stylesheet\" href=\"" + uri + "\" type=\"text/css\" />"
			this.storeAsset("blob-top", link)
		}
	}
}

func (this *Assets) renderJs() {
	for _, uri := range strings.Split(this.store["blob-bottom-js"], ",") {
		if len(uri) > 0 {
			link := ""
			this.storeAsset("blob-bottom", link)
		}
	}
}
