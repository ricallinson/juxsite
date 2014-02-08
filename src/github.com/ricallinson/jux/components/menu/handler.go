package menu

import (
	"github.com/ricallinson/forgery"
)

func Handler(req *f.Request, res *f.Response, next func()) {

	type Link struct {
		Uri  string
		Text string
	}

	res.Render("a/main.html",
		map[string]string{
			"title": "Main Menu",
		},
		map[string][]Link{
			"links": []Link{
				Link{
					Uri:  "/",
					Text: "Home",
				},
				Link{
					Uri:  "/?category=alice",
					Text: "Alice's Adventures in Wonderland",
				},
			},
		})
}
