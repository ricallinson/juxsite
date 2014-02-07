package article

import (
	"encoding/json"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"path"
)

type article struct {
	Title    string
	Category string
	Text     string
}

func loadArticles(dirname string) []article {
	articles := []article{}
	list, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	for _, file := range list {
		if file.IsDir() != true {
			filename := path.Join(dirname, file.Name())
			if article, ok := loadFile(filename); ok {
				article.Text = string(blackfriday.MarkdownBasic([]byte(article.Text)))
				articles = append(articles, article)
			}
		}
	}
	return articles
}

func loadFile(filename string) (article, bool) {
	j, err1 := ioutil.ReadFile(filename)
	if err1 != nil {
		panic(err1)
	}
	a := article{}
	err2 := json.Unmarshal(j, &a)
	if err2 != nil {
		panic(err2)
	}
	return a, true
}
