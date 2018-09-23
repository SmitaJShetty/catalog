package router

import (
	"ArticleApp/controller"

	"github.com/gorilla/mux"
)

func addAppRoutes(r *mux.Router) {
	r.HandleFunc("/v1/articles/{id}", controller.GetArticleByID).Methods("GET")
	r.HandleFunc("/v1/articles", controller.CreateArticle).Methods("POST")
	r.HandleFunc("/v1/tags/{tagName}/{date}", controller.GetTagByNameAndDate).Methods("GET")
	r.HandleFunc("/v1/articles", controller.GetArticles).Methods("GET")
}
