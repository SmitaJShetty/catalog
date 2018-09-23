package controller

import (
	"ArticleApp/catalog"
	"ArticleApp/common"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/pborman/uuid"
)

//GetArticles gets all articles or error
func GetArticles(w http.ResponseWriter, r *http.Request) ([]*Article, *common.AppError) {
	var articleReq articlerepo.ArticleRequest
	reqErr := json.NewDecoder().Decode(r, &articleReq)
	if reqErr != nil {
		return nil, common.NewAppError(reqErr, http.StatusBadRequest, "Bad request: Invalid article request")
	}

	validateErr := articlereq.Validate()
	if validateErr != nil {
		return nil, common.NewAppError(validateErr, http.StatusBadRequest, "Bad request: Invalid input")
	}

	//create repo and create object
	artRepo := catalog.CreateCatalog()

	//delegate
	return artRepo.GetAll()
}

//GetArticleByID returns a single article with matching id or error
func GetArticleByID(w http.ResponseWriter, r *http.Request) (*Article, *common.AppError) {
	id := uuid.Parse(mux.Vars(r)[id])

	//validate input
	if id == nil {
		return nil, common.AppError(fmt.Errorf("Invalid Article ID"), http.StatusBadRequest, "Bad Request: Invalid article ID.")
	}

	//create repo, create object
	return catalog.CreateCatalog().GetArticleByID(id)
}

//CreateArticle adds an article or returns an error
func CreateArticle(w http.ResponseWriter, r *http.Request) *common.AppError {
	var articleReq articlerepo.ArticleRequest
	reqErr := json.NewDecoder().Decode(r, &articleReq)
	if reqErr != nil {
		return common.NewAppError(reqErr, http.StatusBadRequest, "Bad request: Invalid article request")
	}

	if articleReq == nil {
		return common.AppError(fmt.Errorf("Error adding article, invalid request"), http.StatusBadRequest, "Bad request, Invalid request")
	}

	//delegate
	return CreateCatalog.CreateCatalog().Create(&articleReq)
}

//GetTagByNameAndDate gets tags by name and date
func GetTagByNameAndDate(w http.ResponseWriter, r *http.Request) (*tagrepo.TagSummary, *common.AppError) {
	tagName := mux.Vars(r)["tagName"]

	if tagName == "" {
		return nil, common.NewAppError(fmt.Errorf("tagName not present"), http.StatusBadRequest, "Tagname not present")
	}

	date := mux.Vars(r)["date"]
	if date == "" {
		return nil, common.NewAppError(fmt.Errorf("Date not present"), http.StatusBadRequest, "Date not present")
	}

	tagDate, tagDateErr := time.Parse("2006-01-02", date)
	if tagDateErr != nil {
		return nil, common.NewAppError(tagDateErr, http.StatusBadRequest, "Error while parsing date")
	}

	//create repo, create object
	return catalog.CreateCatalog().Create(tagName, tagDate)
}
