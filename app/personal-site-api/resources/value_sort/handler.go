package valuesort

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func CreateBoardHandler(model ValueSortBoardDataAccessLayer) http.HandlerFunc {
	type CreateBoardReqBody struct {
		BoardName string `json:"boardName"`
	}

	type CreateBoardResBody struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody CreateBoardReqBody
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&reqBody)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not process request body - %s", err.Error()), http.StatusUnprocessableEntity)
		}

		err = model.Create(reqBody.BoardName)
		if err != nil {
			if strings.Contains(err.Error(), "Invalid Request Body:") {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			} else {
				http.Error(w, "unable to create board", http.StatusInternalServerError)
			}
			return
		}

		w.Write([]byte("success"))
	}
}

func GetBoardHandler(model ValueSortBoardDataAccessLayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		boardName := vars["boardName"]
		board, err := model.Get(boardName)

		if err != nil {
			http.Error(w, "problem fetching value sort cards", http.StatusInternalServerError)
			return
		}

		if len(board.Columns) == 0 {
			http.Error(w, "board not found", http.StatusNotFound)
			return
		}

		jbytes, err := json.Marshal(board)
		if err != nil {
			http.Error(w, "internal error building response", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(jbytes)
	}
}

func UpdateBoardHandler(model ValueSortBoardDataAccessLayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var board ValueSortBoard
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&board)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not process request body - %s", err.Error()), http.StatusUnprocessableEntity)
		}

		err = model.Upsert(board)
		if err != nil {
			http.Error(w, "unable to update board", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("success"))
	}
}
