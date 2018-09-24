package controller

import (
	"ArticleApp/catalog"
	"ArticleApp/common"
	"ArticleApp/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//GetArticles gets all articles or error
func GetArticles(w http.ResponseWriter, r *http.Request) {
	//create repo and create object
	artRepo := catalog.CreateCatalog()
	articles, getErr := artRepo.GetAll()
	if getErr != nil {
		common.SendErrorResponse(w, r, common.NewAppError(getErr, http.StatusBadRequest, "Bad request: Error while processing output"))
	}

	fmt.Print(articles)

	result, resultErr := json.Marshal(articles)
	if resultErr != nil {
		common.SendErrorResponse(w, r, common.NewAppError(resultErr, http.StatusBadRequest, "Bad Request: Marshal error"))
	}

	common.SendResult(w, r, result)
}

//GetArticleByID returns a single article with matching id or error
func GetArticleByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	//create repo, create object
	article, getErr := catalog.CreateCatalog().GetByID(id)
	if getErr != nil {
		common.SendErrorResponse(w, r, common.NewAppError(getErr, http.StatusBadRequest, "Bad Request: Invalid article ID."))
	}

	result, resultErr := json.Marshal(article)
	if resultErr != nil {
		common.SendErrorResponse(w, r, common.NewAppError(resultErr, http.StatusBadRequest, "Bad Request: Marshal error"))
	}

	common.SendResult(w, r, result)
}

//CreateArticle adds an article or returns an error
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var articleReq models.ArticleRequest
	reqErr := json.NewDecoder(r.Body).Decode(&articleReq)
	if reqErr != nil {
		common.SendErrorResponse(w, r, common.NewAppError(reqErr, http.StatusBadRequest, "Bad request: Invalid article request"))
	}

	createErr := catalog.CreateCatalog().CreateArticle(&articleReq)
	if createErr != nil {
		common.SendErrorResponse(w, r, common.NewAppError(createErr, http.StatusBadRequest, "Bad request: Invalid article request"))
	}
}

//GetTagSummaryByDateAndName gets tags by name and date
func GetTagSummaryByDateAndName(w http.ResponseWriter, r *http.Request) {
	tagName := mux.Vars(r)["tagName"]

	if tagName == "" {
		common.SendErrorResponse(w, r, common.NewAppError(fmt.Errorf("tagName not present"), http.StatusBadRequest, "Tagname not present"))
	}

	date := mux.Vars(r)["date"]
	if date == "" {
		common.SendErrorResponse(w, r, common.NewAppError(fmt.Errorf("Date not present"), http.StatusBadRequest, "Date not present"))
	}

	tagDate, tagDateErr := time.Parse("2006-01-02", date)
	if tagDateErr != nil {
		common.SendErrorResponse(w, r, common.NewAppError(tagDateErr, http.StatusBadRequest, "Error while parsing date"))
	}

	tagSummary, tagSummaryErr := catalog.CreateCatalog().GetTagSummaryByDateAndName(tagDate, tagName)
	if tagSummaryErr != nil {
		common.SendErrorResponse(w, r, common.NewAppError(tagSummaryErr, http.StatusBadRequest, "Error while parsing date"))
	}

	result, resultErr := json.Marshal(tagSummary)
	if resultErr != nil {
		common.SendErrorResponse(w, r, common.NewAppError(resultErr, http.StatusBadRequest, "Bad Request: Marshal error"))
	}

	common.SendResult(w, r, result)
}
