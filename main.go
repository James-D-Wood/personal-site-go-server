package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/cfg"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/database"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/middleware"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/resources/articles"
	valuesort "github.com/jdwoo/personal-site-go-server/app/personal-site-api/resources/value_sort"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "We're up and running, captain!")
}

//go:generate sh -c "printf %s $(git rev-parse HEAD) > .build/commit.txt"
//go:generate sh -c "date -u -I seconds | tr -d '[:space:]' > .build/build-time.txt"
var (
	//go:embed .build/commit.txt
	Commit string
	//go:embed .build/build-time.txt
	BuildTime string
)

type MetaResponse struct {
	CommitSHA string `json:"sha"`
	BuildTime string `json:"build_time"`
	GitHubURL string `json:"github_url"`
}

func metaHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(MetaResponse{Commit, BuildTime, "https://github.com/James-D-Wood/personal-site-go-server"})
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func initializeRoutes(db *pgx.Conn) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/meta", metaHandler)
	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(middleware.CorsMiddleware)
	articles.InitializeRoutes(apiV1.PathPrefix("/articles").Subrouter(), &articles.ArticleModel{DB: db})
	valuesort.InitializeRoutes(apiV1.PathPrefix("/value-sort").Subrouter(), &valuesort.ValueSortBoardModel{DB: db})
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
