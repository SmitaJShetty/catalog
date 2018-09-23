package repo

import (
	"github.com/pborman/uuid"
)

//RedisArticleRepo article repo implemented in redis
type RedisArticleRepo struct {
}

//GetByID gets articles by id passed as argument
func (r *RedisArticleRepo) GetByID(id uuid.UUID) (*Article, error) {
	//validate by id
	if id == nil {
		return nil, nil
	}

	return nil, nil
}

//GetAll gets all articles
func (r *RedisArticleRepo) GetAll() ([]*Article, error) {
	return nil, nil
}

//CreateArticle creates an article
func (r *RedisArticleRepo) CreateArticle(article *ArticleRequest) error {
	return nil
}
