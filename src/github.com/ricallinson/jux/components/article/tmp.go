package article

import (
	"encoding/json"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"path"
	"strings"
)

func loadArticles(dirname string, category string, from int, to int) ([]*Article, int) {
	articles := []*Article{}
	list, err := ioutil.ReadDir(dirname)
	if err != nil {
		// panic(err)
		return articles, 0
	}
	count := len(list)
	if from < 0 {
		from = 0
	}
	if to > count {
		to = count
	}
	for i := from; i < to; i++ {
		file := list[i]
		if file.IsDir() != true {
			article := &Article{Id: file.Name()}
			if loadArticle(dirname, article) {
				if category == "" || article.Category == category {
					articles = append(articles, article)
				}
			}
		}
	}
	return articles, count
}

func loadArticle(dirname string, article *Article) bool {
	filename := path.Join(dirname, article.Id)
	j, err1 := ioutil.ReadFile(filename)
	if err1 != nil {
		// panic(err1)
		return false
	}
	err2 := json.Unmarshal(j, &article)
	if err2 != nil {
		// panic(err2)
		return false
	}
	if line := strings.Index(article.Text, "\r\n"); line > 0 {
		article.Summary = string(blackfriday.MarkdownBasic([]byte(article.Text[:line])))
	}
	article.Text = string(blackfriday.MarkdownBasic([]byte(article.Text)))
	return true
}
