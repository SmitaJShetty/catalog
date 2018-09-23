package catalog

import (
	"ArticleApp/models"
	"fmt"
	"time"

	"github.com/pborman/uuid"
)

//Repository exposes methods for articles
type Repository interface {
	GetByID(id uuid.UUID) (*models.Article, error)
	GetAll() ([]*models.Article, error)
	CreateArticle(article *models.ArticleRequest) error
	GetByDateAndName(date time.Time, name string) ([]*models.Tag, error)
}

//ArticleCatalog construct for manipulating articles and tags
type ArticleCatalog struct {
	repo Repository
}

//GetByID gets articles by id passed as argument
func (ac *ArticleCatalog) GetByID(id uuid.UUID) (*models.Article, error) {
	if id == nil {
		return nil, fmt.Errorf("Invalid guid")
	}
	return repo.GetByID(id)
}

//GetAll gets all articles
func (ac *ArticleCatalog) GetAll() ([]*models.Article, error) {
	return repo.GetAll()
}

//CreateArticle creates an article
func (ac *ArticleCatalog) CreateArticle(article *models.ArticleRequest) error {
	if article == nil {
		return fmt.Errorf("Article is empty")
	}

	return repo.CreateArticle(article)
}

//GetByDateAndName  gets a in memory tag repo stored tag, by name and date
func (ac *ArticleCatalog) GetByDateAndName(date time.Time, name string) ([]*models.Tag, error) {

}
