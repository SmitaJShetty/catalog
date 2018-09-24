package models

//TagSummary summary of tag information
type TagSummary struct {
	Tag         string   `json:"tag"`
	Count       uint     `json:"count"`
	ArticleIDs  []string `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}

//NewTagSummary returns new tag summary
func NewTagSummary(tagName string, articleIDs []string) *TagSummary {
	return &TagSummary{
		Count:      1,
		ArticleIDs: articleIDs,
		Tag:        tagName,
	}
}
