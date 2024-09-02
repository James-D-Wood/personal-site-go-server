package recipes

import (
	"github.com/gorilla/mux"
)

func InitializeRoutes(router *mux.Router, model RecipeDataAccessLayer) {
	router.HandleFunc("", GetRecipesHandler(model)).Methods("GET")
}
