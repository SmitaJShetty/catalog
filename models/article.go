package models

//Article - construct for a basic article
type Article struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
	Body  string   `json:"body"`
}

//NewArticle returns a new article with a generated id
func NewArticle(req *ArticleRequest) *Article {
	return &Article{
		ID:    req.ID,
		Title: req.Title,
		Date:  req.Date,
		Tags:  req.Tags,
		Body:  req.Body,
	}
}
