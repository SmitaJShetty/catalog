package repo

import (
	"ArticleApp/models"
	"fmt"
	"strings"
	"sync"
	"time"
)

//InMemoryArticleRepo article repo implemented in in-memory
type InMemoryArticleRepo struct {
	//dataStore ArticleMap
}

//GetByID gets articles by id passed as argument
func (r *InMemoryArticleRepo) GetByID(id string) (*models.Article, error) {
	store := GetArticleMapInstance()

	value := store.ArticleMap[id]
	return value, nil
}

//GetAll gets all articles
func (r *InMemoryArticleRepo) GetAll() ([]*models.Article, error) {
	return r.getAllInMap()
}

func (r *InMemoryArticleRepo) getAllInMap() ([]*models.Article, error) {
	var articles []*models.Article
	store := GetArticleMapInstance()

	for _, value := range store.ArticleMap {
		articles = append(articles, value)
	}

	return articles, nil
}

//CreateArticle creates an article
func (r *InMemoryArticleRepo) CreateArticle(articleReq *models.ArticleRequest) error {
	if articleReq == nil {
		return fmt.Errorf("Invalid article")
	}

	validateErr := articleReq.Validate()
	if validateErr != nil {
		return validateErr
	}

	newArticle := models.NewArticle(articleReq)

	store := GetArticleMapInstance()

	(store.ArticleMap)[newArticle.ID] = newArticle

	fmt.Println("article map items:", store.ArticleMap)

	tagSummaryInstance := GetTagSummaryInstance()
	tagSummaryInstance.updateTagSummaries(newArticle)
	return nil
}

//GetTagSummaryByDateAndName  gets a in memory tag repo stored tag, by name and date
func (r *InMemoryArticleRepo) GetTagSummaryByDateAndName(articleDate time.Time, tagName string) (*models.TagSummary, error) {
	if strings.Trim(tagName, " ") == "" {
		return nil, fmt.Errorf("Empty tagName")
	}

	var mu sync.RWMutex

	date := articleDate.Format("2006-01-02")
	tagSummaryMapInstance := GetTagSummaryInstance()
	fmt.Println("tagSummaryMapInstance", tagSummaryMapInstance.summary[date])

	mu.RLock()
	defer mu.RUnlock()
	summaryByDate := tagSummaryMapInstance.summary[date]
	if summaryByDate == nil {
		return nil, nil
	}

	return (*summaryByDate)[tagName], nil
}
