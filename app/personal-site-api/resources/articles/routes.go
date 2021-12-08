package articles

import (
	"github.com/gorilla/mux"
)

func InitializeRoutes(router *mux.Router, model ArticleDataAccessLayer) {
	router.HandleFunc("", CreateArticleHandler(model)).Methods("POST")
}
