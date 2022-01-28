package valuesort

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/database"
)

// A struct to model the object
type ValueSortCard struct {
	Body    string `json:"body"`
	Details string `json:"details"`
}

type ValueSortColumn struct {
	Title string          `json:"title"`
	Cards []ValueSortCard `json:"cards"`
	Order int             `json:"order"`
}

type ValueSortBoard struct {
	Name    string            `json:"name"`
	Columns []ValueSortColumn `json:"columns"`
}

// An interface to refresent the Model (for mocking in test)
type ValueSortBoardDataAccessLayer interface {
	Create(boardName string) (err error)
	Get(boardName string) (board ValueSortBoard, err error)
	Upsert(board ValueSortBoard) (err error)
}

// The Model with Database Implementation
type ValueSortBoardModel struct {
	DB *pgx.Conn
}

// Fetch and Assemble Board
func (model *ValueSortBoardModel) Get(boardName string) (board ValueSortBoard, err error) {
	stmt := `
		SELECT board_name, card_body, card_details, column_name
		FROM value_sort_cards
		WHERE board_name = $1;
	`

	rows, err := model.DB.Query(context.Background(), stmt, boardName)
	if err != nil {
		return ValueSortBoard{}, err
	}

	colMap := map[string]ValueSortColumn{}
	columnTitles := []string{
		"Unsorted",
		"Not Important",
		"Somewhat Important",
		"Important",
		"Very Important",
		"Most Important",
	}
	for idx, colTitle := range columnTitles {
		colMap[colTitle] = ValueSortColumn{
			Title: colTitle,
			Cards: []ValueSortCard{},
			Order: idx,
		}
	}

	for rows.Next() {
		var columnName string
		var card ValueSortCard
		err = rows.Scan(&board.Name, &card.Body, &card.Details, &columnName)
		if err != nil {
			return ValueSortBoard{}, nil
		}

		if col, ok := colMap[columnName]; ok {
			col.Cards = append(col.Cards, card)
			colMap[columnName] = col
		}
	}

	for _, column := range colMap {
		board.Columns = append(
			board.Columns,
			column,
		)
	}

	return board, err
}

// Create Board w/ Default Cards
func (model *ValueSortBoardModel) Create(boardName string) (err error) {
	var initialData []ValueSortColumn
	err = json.Unmarshal([]byte(InitialData), &initialData)
	if err != nil {
		return err
	}

	for _, col := range initialData {
		for _, card := range col.Cards {
			stmt := `
				INSERT INTO value_sort_cards (board_name, card_body, card_details, column_name) 
				VALUES ($1, $2, $3, $4)
				RETURNING board_name
			`

			var result string
			err := model.DB.QueryRow(
				context.Background(),
				stmt,
				boardName, card.Body, card.Details, col.Title,
			).Scan(&result)

			if err != nil {
				return database.TranslateError(err)
			}
		}
	}

	return err
}

func (model *ValueSortBoardModel) Upsert(board ValueSortBoard) (err error) {
	for _, col := range board.Columns {
		for _, card := range col.Cards {
			// TODO: can accidentally create new tables
			stmt := `
				INSERT INTO value_sort_cards (board_name, card_body, card_details, column_name) 
				VALUES ($1, $2, $3, $4)
				ON CONFLICT (board_name, card_body) 
				DO 
					UPDATE SET column_name = $4
				RETURNING board_name
			`

			var result string
			err := model.DB.QueryRow(
				context.Background(),
				stmt,
				board.Name, card.Body, card.Details, col.Title,
			).Scan(&result)

			if err != nil {
				fmt.Println(err)
				return database.TranslateError(err)
			}
		}
	}
	return nil
}
