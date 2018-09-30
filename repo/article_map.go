package repo

import (
	"ArticleApp/models"
	"sync"
)

type ArticleMap struct {
	ArticleMap map[string]*models.Article
}

var articleMapInstance *ArticleMap
var mu sync.Mutex

//GetArticleMapInstance returns new tag summary
func GetArticleMapInstance() *ArticleMap {
	if articleMapInstance == nil {
		mu.Lock()
		defer mu.Unlock()

		articleMapInstance = &ArticleMap{}
		articleMapInstance.ArticleMap = make(map[string]*models.Article)
	}

	return articleMapInstance
}

func ResetArticleMap() {
	articleMapInstance = nil
}
