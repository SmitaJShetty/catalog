package repo

import (
	"ArticleApp/common"
	"ArticleApp/models"
	"fmt"
	"sync"
)

//TagSummaries for easy access to Tag Summary
type TagSummaries struct {
	summary map[string]*map[string]*models.TagSummary
}

var tagSummaryMapInstance *TagSummaries
var mut sync.Mutex

//GetTagSummaryInstance returns new tag summary
func GetTagSummaryInstance() *TagSummaries {
	if tagSummaryMapInstance == nil {
		mut.Lock()
		defer mut.Unlock()

		tagSummaryMapInstance = &TagSummaries{}
		tagSummaryMapInstance.summary = make(map[string]*map[string]*models.TagSummary)
	}

	return tagSummaryMapInstance
}

//updateTagSummaries creates a tagsummary and updates cache
func (r *TagSummaries) updateTagSummaries(article *models.Article) error {
	if article == nil {
		return fmt.Errorf("Invalid article")
	}

	articleDate := article.Date
	summaryInstance := GetTagSummaryInstance()
	dateMapValue, dateMapValueFound := summaryInstance.summary[articleDate]

	if !dateMapValueFound {
		m := make(map[string]*models.TagSummary)

		for _, tag := range article.Tags {
			newTagSummary := models.NewTagSummary(tag, []string{article.ID})
			m[tag] = newTagSummary
		}

		summaryInstance.summary[articleDate] = &m
		dateMapValue = summaryInstance.summary[articleDate]
	}

	for _, tag := range article.Tags {
		tagSummary := (*dateMapValue)[tag]

		if tagSummary == nil {
			tagSummary = models.NewTagSummary(tag, []string{article.ID})
		}

		r.populateTagSummary(tagSummary, article)
		(*dateMapValue)[tag] = tagSummary
	}

	return nil
}

func (r *TagSummaries) populateTagSummary(tagSummary *models.TagSummary, article *models.Article) *models.TagSummary {
	differentArticle := false

	if !common.IfExists(article.ID, tagSummary.ArticleIDs) {
		tagSummary.ArticleIDs = append(tagSummary.ArticleIDs, article.ID)
		differentArticle = true
	}

	for _, articleTag := range article.Tags {
		if tagSummary.Tag != articleTag && !common.IfExists(articleTag, tagSummary.RelatedTags) {
			tagSummary.RelatedTags = append(tagSummary.RelatedTags, articleTag)
			differentArticle = true
		}
	}

	if differentArticle {
		tagSummary.Count++
	}

	return tagSummary
}

func (r *TagSummaries) updateExistingTagSummary(tagName string, tagSummary *models.TagSummary, article *models.Article) (*models.TagSummary, error) {
	var differentArticle bool
	differentArticle = false

	if !common.IfExists(article.ID, tagSummary.ArticleIDs) {
		tagSummary.ArticleIDs = append(tagSummary.ArticleIDs, article.ID)
		differentArticle = true
	}

	for _, articleTag := range article.Tags {
		if tagName != articleTag && !common.IfExists(articleTag, tagSummary.RelatedTags) {
			tagSummary.RelatedTags = append(tagSummary.RelatedTags, articleTag)
			differentArticle = true
		}
	}

	if differentArticle {
		tagSummary.Count++
	}

	return tagSummary, nil
}
