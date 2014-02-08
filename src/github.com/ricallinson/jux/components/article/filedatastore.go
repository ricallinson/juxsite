package article

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

type DataStore struct {
	Root string
}

func CreateDataStore(dirname string) *DataStore {
	ds := &DataStore{}
	ds.Root = dirname
	return ds
}

func (this *DataStore) LoadTable(table string, category string, from int, to int) ([]*Article, int) {
	articles := []*Article{}
	dirname := path.Join(this.Root, table)
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
			if this.LoadItem(table, article) {
				if category == "" || article.Category == category {
					articles = append(articles, article)
				}
			}
		}
	}
	return articles, count
}

func (this *DataStore) LoadItem(table string, article *Article) bool {
	filename := path.Join(this.Root, table, article.Id)
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
	return true
}
