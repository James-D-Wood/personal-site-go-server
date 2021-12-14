package articles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jdwoo/personal-site-go-server/internal/webserverutils"
)

func GetArticlesHandler(model ArticleDataAccessLayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := model.All()
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "problem fetching articles", http.StatusInternalServerError)
			return
		}
		jbytes, err := json.Marshal(articles)
		if err != nil {
			http.Error(w, "internal error building response", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(jbytes)
	}
}

func GetArticleHandler(model ArticleDataAccessLayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		articleURI := vars["articleURI"]
		if articleURI == "" {
			http.Error(w, "request missing article ID", http.StatusNotFound)
			return
		}
		article, err := model.Get(articleURI)
		if err != nil {
			if err.Error() == "no rows in result set" {
				http.Error(w, "article not found", http.StatusNotFound)
			} else {
				http.Error(w, "problem fetching article", http.StatusInternalServerError)
			}
			return
		}
		jbytes, err := json.Marshal(article)
		if err != nil {
			http.Error(w, "internal error building response", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(jbytes)
	}
}

func CreateArticleHandler(model ArticleDataAccessLayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a Article
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&a)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		errs := model.Validate(a)
		if len(errs) > 0 {
			errMsgs := []string{}
			for _, err := range errs {
				errMsgs = append(errMsgs, err.Error())
			}
			msg := webserverutils.NewRequestError(strings.Join(errMsgs, ", "))
			http.Error(w, msg.Error(), http.StatusUnprocessableEntity)
			return
		}

		savedArticle, err := model.Save(a)
		if err != nil {
			if strings.Contains(err.Error(), "Invalid Request Body:") {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		bytes, _ := json.Marshal(savedArticle)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(bytes)
	}
}

func UpdateArticleHandler(model ArticleDataAccessLayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		articleURI := vars["articleURI"]
		if articleURI == "" {
			http.Error(w, "request missing article ID", http.StatusNotFound)
			return
		}

		var a Article
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&a)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		errs := model.Validate(a)
		if len(errs) > 0 {
			errMsgs := []string{}
			for _, err := range errs {
				errMsgs = append(errMsgs, err.Error())
			}
			msg := webserverutils.NewRequestError(strings.Join(errMsgs, ", "))
			http.Error(w, msg.Error(), http.StatusUnprocessableEntity)
			return
		}

		updatedArticle, err := model.Update(articleURI, a)
		if err != nil {
			if strings.Contains(err.Error(), "Invalid Request Body:") {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			} else if err.Error() == "no rows in result set" {
				http.Error(w, "article not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		bytes, _ := json.Marshal(updatedArticle)
		w.Header().Add("Content-Type", "application/json")
		w.Write(bytes)
	}
}
