package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/cfg"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/database"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/middleware"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/resources/articles"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/resources/learning"
	valuesort "github.com/jdwoo/personal-site-go-server/app/personal-site-api/resources/value_sort"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "We're up and running, captain!")
}

func initializeRoutes(db *pgx.Conn) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.HandleFunc("/", rootHandler)
	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(middleware.CorsMiddleware)
	articles.InitializeRoutes(apiV1.PathPrefix("/articles").Subrouter(), &articles.ArticleModel{DB: db})
	valuesort.InitializeRoutes(apiV1.PathPrefix("/value-sort").Subrouter(), &valuesort.ValueSortBoardModel{DB: db})
	learning.InitializeRoutes(apiV1.PathPrefix("/lessons").Subrouter(), &learning.LessonModel{DB: db})
	return r
}

func main() {
	config := cfg.Load()

	db, err := database.InitalizeDatabase(config)
	if err != nil {
		panic(err)
	}
	defer database.TeardownDatabase(db)
	log.Fatal(http.ListenAndServe(":8080", initializeRoutes(db)))
}
