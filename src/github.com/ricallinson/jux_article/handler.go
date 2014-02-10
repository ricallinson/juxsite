package jux_article

import (
	"github.com/ricallinson/forgery"
	"github.com/ricallinson/jux"
	"github.com/ricallinson/jux/helpers/assets"
	"github.com/ricallinson/jux/helpers/datastore"
	"strconv"
	"strings"
)

// Route the request to the correct handler function.
func Handler(req *f.Request, res *f.Response, next func()) {
	switch req.Params["juxview"] {
	case "main":
		list(req, res, next)
	case "read":
		read(req, res, next)
	case "listjson":
		listJson(req, res, next)
	}
}

// Shows a list of articles for the given category.
func list(req *f.Request, res *f.Response, next func()) {

	// Should this be in a config?
	asset := assets.New(res.Locals)
	asset.AddCss("/jux_article/css/screen.css")

	// First fetch all the params needed.
	batch := 5
	start, _ := strconv.Atoi(req.Query["start"])
	category := strings.ToLower(req.Query["category"])

	// If "category" is empty check the config.
	if len(category) == 0 {
		category = "general"
	}

	// Create a Query.
	query := &Article{
		Category: category,
	}
	articles := []*Article{}

	// Grab the datastore.
	ds := datastore.New(jux.GetNewContext(req))

	// Get the list of articles matching the request.
	ds.List(query, start, batch, &articles)

	// Calculate the previous/next links.
	less := start - batch
	more := start + batch

	// Render the Summary as HTML.
	for _, article := range articles {
		article.InflateSummary()
	}
	res.Render("jux_article/list.html", map[string][]*Article{
		"Articles": articles,
	}, map[string]string{
		"Title":    "All Articles",
		"Less":     strconv.Itoa(less),
		"More":     strconv.Itoa(more),
		"Category": category,
	}, map[string]bool{
		"Show_less": less >= 0,
		"Show_more": batch == len(articles),
	})
}

// Shows a single article for the given id.
func read(req *f.Request, res *f.Response, next func()) {

	// Create a Query.
	article := &Article{}
	article.Id = req.Query["id"]

	// Grab the datastore.
	ds := datastore.New(jux.GetNewContext(req))
	if err := ds.Read(article); err != nil {
		res.Render("notfound/main.html", map[string]string{
			"error": "Article not found.",
		})
		return
	}

	// Inflate the Article.
	article.InflateBody()

	// Render the Article.
	res.Locals["pageTitle"] = article.Title
	res.Render("jux_article/read.html", map[string][]*Article{
		"Articles": []*Article{article},
	})
}

// Shows a JSON Object that is a list of articles for the given category.
func listJson(req *f.Request, res *f.Response, next func()) {

	// Used to prime the datastore.
	LoadJsonArticles(req, "data/articles")

	// Process query params.
	start, _ := strconv.Atoi(req.Query["start"])
	category := strings.ToLower(req.Query["category"])

	// Create a Query.
	query := &Article{
		Category: category,
	}
	articles := []*Article{}

	// Grab the datastore.
	ds := datastore.New(jux.GetNewContext(req))
	if err := ds.List(query, start, -1, &articles); err != nil {
		panic(err.Error())
	}

	// Inflate the Article.
	for _, article := range articles {
		article.InflateSummary()
		article.InflateBody()
	}

	// Render as Json.
	res.Json(articles)
}

// Shows a menu of articles for the given category.
func Menu(req *f.Request, res *f.Response, next func()) {

	if req.Params["juxview"] != "article" && req.Params["juxview"] != "read" {
		return
	}

	// Create a Query.
	query := &Article{
		Category: req.Query["category"],
	}
	articles := []*Article{}

	// Grab the datastore.
	ds := datastore.New(jux.GetNewContext(req))
	ds.List(query, 0, -1, &articles)

	// Render.
	res.Render("jux_article/menu.html", map[string][]*Article{
		"links": articles,
	}, map[string]string{
		"title": "Category Articles",
	})
}
