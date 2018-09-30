package catalog

import (
	"ArticleApp/models"
	"fmt"
	"strings"
	"time"
)

//Repository exposes methods for articles
type Repository interface {
	GetByID(id string) (*models.Article, error)
	GetAll() ([]*models.Article, error)
	CreateArticle(article *models.ArticleRequest) error
	GetTagSummaryByDateAndName(date time.Time, name string) (*models.TagSummary, error)
	Reset()
}

//ArticleCatalog construct for manipulating articles and tags
type ArticleCatalog struct {
	repository Repository
}

//GetByID gets articles by id passed as argument
func (ac *ArticleCatalog) GetByID(id string) (*models.Article, error) {
	return ac.repository.GetByID(id)
}

//GetAll gets all articles
func (ac *ArticleCatalog) GetAll() ([]*models.Article, error) {
	return ac.repository.GetAll()
}

//CreateArticle creates an article
func (ac *ArticleCatalog) CreateArticle(articleReq *models.ArticleRequest) error {
	if articleReq == nil {
		return fmt.Errorf("Article is empty")
	}

	validateErr := articleReq.Validate()
	if validateErr != nil {
		return validateErr
	}

	return ac.repository.CreateArticle(articleReq)
}

//GetTagSummaryByDateAndName  gets a in memory tag ac.repository stored tag, by name and date
func (ac *ArticleCatalog) GetTagSummaryByDateAndName(date time.Time, tagName string) (*models.TagSummary, error) {
	if strings.TrimSpace(tagName) == "" {
		return nil, fmt.Errorf("tag name is empty")
	}

	if date.IsZero() {
		return nil, fmt.Errorf("Empty date")
	}

	return ac.repository.GetTagSummaryByDateAndName(date, tagName)
}
