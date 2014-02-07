package article

import (
	"encoding/json"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"path"
	"strings"
)

func loadArticles(dirname string) []*Article {
	articles := []*Article{}
	list, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
		return articles
	}
	for _, file := range list {
		if file.IsDir() != true {
			article := &Article{Id: file.Name()}
			if loadArticle(dirname, article) {
				articles = append(articles, article)
			}
		}
	}
	return articles
}

func loadArticle(dirname string, article *Article) (bool) {
	filename := path.Join(dirname, article.Id)
	j, err1 := ioutil.ReadFile(filename)
	if err1 != nil {
		panic(err1)
		return false
	}
	err2 := json.Unmarshal(j, &article)
	if err2 != nil {
		panic(err2)
		return false
	}
	if line := strings.Index(article.Text, "\r\n"); line > 0 {
		article.Summary = string(blackfriday.MarkdownBasic([]byte(article.Text[:line])))
	}
	article.Text = string(blackfriday.MarkdownBasic([]byte(article.Text)))
	return true
}
