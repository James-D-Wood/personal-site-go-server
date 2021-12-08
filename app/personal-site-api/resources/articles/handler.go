package articles

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jdwoo/personal-site-go-server/internal/webserverutils"
)

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

		a.ID, err = model.Save(a)
		if err != nil {
			if strings.Contains(err.Error(), "Invalid Request Body:") {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		bytes, _ := json.Marshal(a)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(bytes)
	}
}
