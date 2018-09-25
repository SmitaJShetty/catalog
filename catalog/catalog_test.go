package catalog

import (
	"ArticleApp/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetAll_Success(t *testing.T) {
	catalog := CreateCatalog()

	articleReq := models.ArticleRequest{
		ID:    "1",
		Title: "The Shire",
		Tags:  []string{"Bilbo Baggins", "Samwise Gamgee"},
		Date:  "2008-09-08",
		Body:  "In the hole in the ground there lived a Hobbit. Not a nasty,dirty, wet hole...",
	}
	createErr := catalog.CreateArticle(&articleReq)
	assert.Nil(t, createErr)

	actual, err := catalog.GetByID("1")
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, actual.ID, "1")
	assert.Equal(t, actual.Date, "2008-09-08")
}

func Test_GetByID_Success(t *testing.T) {
	catalog := CreateCatalog()

	articleReq := models.ArticleRequest{
		ID:    "1",
		Title: "The Shire",
		Tags:  []string{"Bilbo Baggins", "Samwise Gamgee"},
		Date:  "2008-09-08",
		Body:  "In the hole in the ground there lived a Hobbit. Not a nasty,dirty, wet hole...",
	}
	createErr := catalog.CreateArticle(&articleReq)
	assert.Nil(t, createErr)

	articleReq = models.ArticleRequest{
		ID:    "2",
		Title: "One ring to rule them all",
		Tags:  []string{"Ranger", "Gandalf, the gray"},
		Date:  "2008-09-11",
		Body:  "When Mr. Bilbo Baggins of Bag End announced that he would shortly be celebrating his eleventy-first birthday...",
	}
	createErr = catalog.CreateArticle(&articleReq)
	assert.Nil(t, createErr)

	actual, err := catalog.GetByID("2")
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, actual.ID, "2")
	assert.Equal(t, actual.Date, "2008-09-11")
}

func Test_CreateArticle_Success(t *testing.T) {
	catalog := CreateCatalog()

	articleReq := models.ArticleRequest{
		ID:    "1",
		Title: "The Shire",
		Tags:  []string{"Bilbo Baggins", "Samwise Gamgee"},
		Date:  "2008-09-08",
		Body:  "In the hole in the ground there lived a Hobbit. Not a nasty,dirty, wet hole...",
	}
	createErr := catalog.CreateArticle(&articleReq)
	assert.Nil(t, createErr)
}

func Test_GetTagByDateAndName_Success(t *testing.T) {
	catalog := CreateCatalog()

	articleReq := models.ArticleRequest{
		ID:    "1",
		Title: "The Shire",
		Tags:  []string{"Bilbo Baggins", "Samwise Gamgee"},
		Date:  "2008-09-11",
		Body:  "In the hole in the ground there lived a Hobbit. Not a nasty,dirty, wet hole...",
	}
	createErr := catalog.CreateArticle(&articleReq)
	assert.Nil(t, createErr)

	articleReq = models.ArticleRequest{
		ID:    "2",
		Title: "One ring to rule them all",
		Tags:  []string{"Ranger", "Gandalf, the gray", "Bilbo Baggins"},
		Date:  "2008-09-11",
		Body:  "When Mr. Bilbo Baggins of Bag End announced that he would shortly be celebrating his eleventy-first birthday...",
	}
	createErr = catalog.CreateArticle(&articleReq)
	assert.Nil(t, createErr)

	date, tagDateErr := time.Parse("2006-01-02", "2008-09-11")
	assert.Nil(t, tagDateErr)

	actual, err := catalog.GetTagSummaryByDateAndName(date, "Bilbo Baggins")
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, actual.Tag, "Bilbo Baggins")
	assert.Equal(t, []string{"1", "2"}, actual.ArticleIDs)
	assert.Equal(t, []string{"Samwise Gamgee", "Ranger", "Gandalf, the gray"}, actual.RelatedTags)
}
