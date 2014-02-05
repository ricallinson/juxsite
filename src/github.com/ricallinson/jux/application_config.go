package jux

type AppCfg struct {

	PageTitle string
	PageTheme string

	DefaultComponent string
	DefaultComponentView string
}

func (this *AppCfg) Load(file string) {

	this.PageTitle = "Jux - the Content Managment System"
	this.PageTheme = "default"

	this.DefaultComponent = "article"
	this.DefaultComponentView = "main"
}