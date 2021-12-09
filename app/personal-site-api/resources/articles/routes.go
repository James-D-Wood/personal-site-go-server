package articles

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/middleware"
)

func NoContentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}
}

func InitializeRoutes(router *mux.Router, model ArticleDataAccessLayer) {
	// router.HandleFunc("", NoContentHandler()).Methods("OPTIONS")
	router.HandleFunc("", GetArticlesHandler(model)).Methods("GET")
	router.HandleFunc("", middleware.AuthMiddleware(CreateArticleHandler(model))).Methods("POST")
	// router.HandleFunc("/{articleURI}", NoContentHandler()).Methods("OPTIONS")
	router.HandleFunc("/{articleURI}", GetArticleHandler(model)).Methods("GET")
}
