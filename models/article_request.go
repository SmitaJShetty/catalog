package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//ArticleRequest request structure for article
type ArticleRequest struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
	Body  string   `json:"body"`
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

	_, parseErr := time.Parse("2006-01-02", ar.Date)
	if parseErr != nil {
		return parseErr
	}

	id, numParseErr := strconv.Atoi(ar.ID)
	if numParseErr != nil {
		return nil
	}

	if id <= 0 {
		return fmt.Errorf("invalid id")
	}

	return nil
}
