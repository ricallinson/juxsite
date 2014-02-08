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

// A slice of Articles.
type Articles []*Article

// Sort methdos for Articles.
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

// Read all files in the given "table" directory and store them in &FileDataStore{}.
func (this *FileDataStore) fillCache(table string) {
	this.ArticlesMap = map[string]*Article{}
	dirname := path.Join(this.Root, table)
	list, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	for _, file := range list {
		if file.IsDir() != true {
			filepath := path.Join(this.Root, table, file.Name())
			if article, ok := this.readFile(filepath); ok {
				article.Id = file.Name()
				this.ArticlesMap[article.Id] = article
			}
		}
	}
}

// Read the JSON file at given "filepath" into an &Article{}.
func (this *FileDataStore) readFile(filepath string) (*Article, bool) {
	article := &Article{}
	j, err1 := ioutil.ReadFile(filepath)
	if err1 != nil {
		// panic(err1)
		return article, false
	}
	err2 := json.Unmarshal(j, &article)
	if err2 != nil {
		// panic(err2)
		return article, false
	}
	return article, true
}

// Returns a list of all items in the given "table".
func (this *FileDataStore) GetTable(table string, category string, from int, to int) (Articles, int) {
	if this.ArticlesMap == nil {
		this.fillCache(table)
	}
	articles := []*Article{}
	for _, article := range this.ArticlesMap {
		if category == "" || article.Category == category {
			articles = append(articles, article)
		}
	}
	count := len(articles)
	if from < 0 {
		from = 0
	}
	if from > count {
		from = count - 1
	}
	if to > count {
		to = count
	}
	sort.Sort(Articles(articles))
	return articles[from:to], count
}

// Returns a single item identified by the given "id".
func (this *FileDataStore) GetItem(table string, id string) (*Article, bool) {
	if this.ArticlesMap == nil {
		this.fillCache(table)
	}
	if article, ok := this.ArticlesMap[id]; ok {
		return article, true
	}
	return nil, false
}
