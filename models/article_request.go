package models

import (
	"fmt"
	"strings"
	"time"
)

//ArticleRequest request structure for article
type ArticleRequest struct {
	ID    int       `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
	Tags  []string  `json:"tags"`
	Body  string    `json:"body"`
}

//Validate validates article request
func (ar *ArticleRequest) Validate() error {
	if strings.Trim(ar.Title, " ") == "" {
		return fmt.Errorf("Empty title of article")
	}

	if strings.Trim(ar.Body, " ") == "" {
		return fmt.Errorf("Empty article body")
	}

	if len(ar.Tags) == 0 {
		return fmt.Errorf("Article not tagged")
	}

	return nil
}
