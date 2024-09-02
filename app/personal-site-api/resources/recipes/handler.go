package recipes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRecipesHandler(model RecipeDataAccessLayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipes, err := model.All()
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "problem fetching recipes", http.StatusInternalServerError)
			return
		}
		jbytes, err := json.Marshal(recipes)
		if err != nil {
			http.Error(w, "internal error building response", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(jbytes)
	}
}
