package jux_article

import (
	"encoding/json"
	"github.com/ricallinson/forgery"
	"io/ioutil"
	"path"
	"sort"
)

type FileDataStore struct {
	// The directory path to the file store.
	Root string
	// Tmp hack for caching all articles.
	ArticlesMap map[string]*Article
}

type Articles []*Article

func (a Articles) Len() int           { return len(a) }
func (a Articles) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Articles) Less(i, j int) bool { return a[i].Id < a[j].Id }

// Returns a new FileDataStore or an instance already attached to the request.
func GetFileDataStore(req *f.Request, dirname string) *FileDataStore {
	key := "filedatastore-" + dirname
	if ds, ok := req.Map[key]; ok {
		return ds.(*FileDataStore)
	}
	ds := &FileDataStore{}
	ds.Root = dirname
	req.Map[key] = ds
	return ds
}

func (this *FileDataStore) fillCache(table string) {
	this.ArticlesMap = map[string]*Article{}
	dirname := path.Join(this.Root, table)
	list, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	for _, file := range list {
		if file.IsDir() != true {
			article := &Article{Id: file.Name()}
			if this.readFile(table, article) {
				this.ArticlesMap[article.Id] = article
			}
		}
	}
}

func (this *FileDataStore) readFile(table string, article *Article) bool {
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

func (this *FileDataStore) LoadTable(table string, category string, from int, to int) (Articles, int) {
	if this.ArticlesMap == nil {
		this.fillCache(table)
	}
	articles := []*Article{}
	for _, article := range this.ArticlesMap {
		if category == "" || article.Category == category {
			articles = append(articles, article)
		}
	}
	if from < 0 {
		from = 0
	}
	count := len(articles)
	if to > count {
		to = count
	}
	sort.Sort(Articles(articles))
	return articles[from:to], count
}

func (this *FileDataStore) LoadItem(table string, id string) (*Article, bool) {
	if this.ArticlesMap == nil {
		this.fillCache(table)
	}
	if article, ok := this.ArticlesMap[id]; ok {
		return article, true
	}
	return nil, false
}
