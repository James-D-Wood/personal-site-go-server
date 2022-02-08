package learning

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jdwoo/personal-site-go-server/internal/webserverutils"
	"gopkg.in/go-playground/validator.v9"
)

func CreateLessonHandler(model LessonDataAccessLayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var l Lesson
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&l)
		if err != nil {
			webserverutils.RespondWithJsonError(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		// validator
		validate := validator.New()
		err = validate.Struct(l)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				fmt.Println(err)
			}

			errs := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errs = append(errs, fmt.Sprintf("%s - %s", e.StructField(), e.Tag()))
			}

			webserverutils.RespondWithJsonError(w, strings.Join(errs, ", "), http.StatusUnprocessableEntity)
			return
		}

		l.CreatedAt = time.Now().UTC()

		err = model.Create(l)
		if err != nil {
			webserverutils.RespondWithJsonError(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		bytes := []byte("created")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(bytes)
	}
}
