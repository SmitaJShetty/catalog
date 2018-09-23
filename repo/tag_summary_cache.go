package repo

import (
	"ArticleApp/models"
	"fmt"
	"strings"
)

//TagSummaryCache for easy access to Tag Summary
type TagSummaryCache struct {
	summary map[string]*map[string]*models.TagSummary
}

var tagSummaryMapInstance *TagSummaryCache
var initialized uint32

//GetTagSummaryInstance returns new tag summary
func GetTagSummaryInstance() *TagSummaryCache {
	if atomic.LoadUInt32(initialized) == 1 {
		return tagSummaryMapInstance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		tagSummaryMapInstance = &TagSummaryCache{}
		tagSummaryMapInstance.summary = make(map[string]map[Tag]TagSummary)
		atomic.StoreUInt32(&initialized, 1)
	}

	return tagSummaryMapInstance
}

//updateTagSummaryCache creates a tagsummary and updates cache
func (r *TagSummaryCache) updateTagSummaryCache(article *models.Article) error {
	if strings.Trim(tagName) == "" {
		return fmt.Errorf("Empty tagName")
	}

	if article == nil {
		return fmt.Errorf("Invalid article")
	}

	articleDate := article.Date.Format("2006-01-02")
	sumry := GetTagSummaryInstance()
	if dateMapValue, dateMapValue := sumry[articleDate]; !dateMapValue { //tag summary exists
		tags := article.Tags
		for _, tag := range tags {
			m := make(map[string]*models.TagSummary)

			newTagSummary, createTagSummaryErr := createNewTagSummary(&m, article)
			if createTagSummaryErr != nil {
				fmt.Printf(createTagSummaryErr.Error())
				continue
			}

			m[tag] = newTagSummary
		}
		sumry[articleDate] = &m

	} else { //tag summary map exists for article's article date
		tags := article.Tags
		for _, tag := range tags {
			var updateErr error
			if existingTagSummary, tagMapFound := dateMapValue[tag]; tagMapFound {
				existingTagSummary, updateErr = updateExistingTagSummary(tag, existingTagSummary, article)
				if updateErr != nil {
					fmt.Printf(updateErr.Error())
					continue
				}

				dateMapValue[tag] = existingTagSummary
			} else {
				newTagSummary, createTagSummaryErr := createNewTagSummary(dateMapValue, article)
				if createTagSummaryErr != nil {
					fmt.Printf(createTagSummaryErr.Error())
					continue
				}

				dateMapValue[tag] = newTagSummary
			}
		}
		sumry[articleDate] = dateMapValue
	}
	return nil
}

func (r *TagSummaryCache) updateExistingTagSummary(tagName string, existingTagSummary *models.TagSummary, article *models.Article) (*models.TagSummary, error) {
	var differentArticle bool
	differentArticle = false

	if !common.IfExists(article.ID, existingTagSummary.ArticleIDs) {
		existingTagSummary.ArticleIDs = append(existingTagSummary.ArticleIDs, article.ID)
		differentArticle = true
	}

	for _, articleTag := range article.Tags {
		if common.IfExists(articleTag, existingTagSummary.RelatedTags) {
			existingTagSummary.RelatedTags = append(existingTagSummary.RelatedTags, articleTag)
			differentArticle = true
		}
	}

	if differentArticle {
		existingTagSummary.Count++
	}

	return existingTagSummary, nil
}

func (r *TagSummaryCache) createNewTagSummary(article *models.Article) (*models.TagSummary, error) {
	if article == nil {
		return fmt.Errorf("Empty article")
	}

	var articleIDs []strings
	var relatedTags []strings

	articleIDs = append(articleIDs, article.ID)
	relatedTags = append(relatedTags, article.Tags...)
	return models.NewTagSummary(articleIDs, relatedTags), nil
}
