package repo

import (
	"ArticleApp/models"
	"time"
)

//RedisArticleRepo article repo implemented in redis
type RedisArticleRepo struct {
}

//GetByID gets articles by id passed as argument
func (r *RedisArticleRepo) GetByID(id string) (*models.Article, error) {
	return nil, nil
}

//GetAll gets all articles
func (r *RedisArticleRepo) GetAll() ([]*models.Article, error) {
	return nil, nil
}

//CreateArticle creates an article
func (r *RedisArticleRepo) CreateArticle(article *models.ArticleRequest) error {
	return nil
}

//GetTagSummaryByDateAndName  gets a in memory tag repo stored tag, by name and date
func (r *RedisArticleRepo) GetTagSummaryByDateAndName(articleDate time.Time, tagName string) (*models.TagSummary, error) {
	return nil, nil
}
