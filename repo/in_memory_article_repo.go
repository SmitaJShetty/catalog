package repo

import (
	"ArticleApp/models"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/pborman/uuid"
)

//InMemoryArticleRepo article repo implemented in in-memory
type InMemoryArticleRepo struct {
	ArticleMap map[int]models.Article
}

//GetByID gets articles by id passed as argument
func (r *InMemoryArticleRepo) GetByID(id uuid.UUID) (*models.Article, error) {
	value := r.ArticleMap[id]
	return &value, nil
}

//GetAll gets all articles
func (r *InMemoryArticleRepo) GetAll() ([]*models.Article, error) {
	return getAllInMap(), nil
}

func (r *InMemoryArticleRepo) getAllInMap() ([]models.Article, error) {
	var articles []models.Article
	count := len(r.ArticleMap)

	for _, value := range r.ArticleMap {
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

	//save article in repo
	r.ArticleMap[newArticle.ID] = newArticle

	//update tag summary
	r.updateTagSummaryCache(newArticle)
	return nil
}

//GetTagSummaryByDateAndName  gets a in memory tag repo stored tag, by name and date
func (r *InMemoryArticleRepo) GetTagSummaryByDateAndName(articleDate time.Time, tagName string) ([]*models.Tag, error) {
	if strings.Trim(tagName) == "" {
		return nil, fmt.Errorf()
	}

	var mu sync.Mutex

	date := articleDate.Format("2006-01-02")
	tagSummaryMapInstance := GetTagSummaryInstance()

	mu.RLock()
	defer mu.RUnlock()
	return tagSummaryMapInstance.summary[date][tagName], nil
}
