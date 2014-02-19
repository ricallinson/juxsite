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
	default:
		list(req, res, next)
	case "read":
		read(req, res, next)
	case "listjson":
		list_json(req, res, next)
	case "primer":
		primer(req, res, next)
	}
}

// Used to prime the datastore.
func primer(req *f.Request, res *f.Response, next func()) {
	if err := LoadJsonArticles(req, "data/articles"); err != nil {
		res.End(err.Error())
		return
	}
	if err := LoadCategories(req); err != nil {
		res.End(err.Error())
		return
	}
	res.End("Primer complete.")
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

	articles, _ := GetArticles(req, category, start, batch)

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
		"Title":    LookUpCategory(req, category, "Articles"),
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
func list_json(req *f.Request, res *f.Response, next func()) {

	// Process query params.
	start, _ := strconv.Atoi(req.Query["start"])
	category := strings.ToLower(req.Query["category"])

	articles, err := GetArticles(req, category, start, -1)

	if err != nil {
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
func HandlerMenuCategories(req *f.Request, res *f.Response, next func()) {

	categories, _ := GetCategories(req)

	// Render.
	res.Render("jux_article/menu_categories.html", map[string][]*Category{
		"links": categories,
	}, map[string]string{
		"title": "Categories",
	})
}

// Shows a menu of articles for the given category.
func HandlerMenuArticles(req *f.Request, res *f.Response, next func()) {

	if req.Params["juxview"] != "article" && req.Params["juxview"] != "read" {
		return
	}

	articles, _ := GetArticles(req, req.Query["category"], 0, -1)

	// Render.
	res.Render("jux_article/menu_articles.html", map[string][]*Article{
		"links": articles,
	}, map[string]string{
		"title": "Category Articles",
	})
}
