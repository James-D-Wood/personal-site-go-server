package valuesort

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router, model ValueSortBoardDataAccessLayer) {
	router.HandleFunc("/boards", CreateBoardHandler(model)).Methods("POST")
	router.HandleFunc("/boards/{boardName}", GetBoardHandler(model)).Methods("GET")
	//router.HandleFunc("/boards/{boardName}/cards/{cardName}", UpdateBoardCardHandler(model)).Methods("PUT")
}
