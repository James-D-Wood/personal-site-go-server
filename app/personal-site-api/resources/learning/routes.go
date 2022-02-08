package learning

import (
	"github.com/gorilla/mux"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/middleware"
)

func InitializeRoutes(router *mux.Router, model LessonDataAccessLayer) {
	router.HandleFunc("", middleware.AuthMiddleware(CreateLessonHandler(model))).Methods("POST")
}
