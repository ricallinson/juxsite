package jux_link_menu

type Link struct {
	Uri  string
	Text string
}

type Menu struct {
	Title string
	Links []Link
}

type Config struct {
	Match map[string]Menu
}
