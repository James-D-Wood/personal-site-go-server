package articles

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/database"
)

// A struct to model the object
type Article struct {
	ID      int    `json:"id"`
	URI     string `json:"uri"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Body    string `json:"body"`
}

// An interface to refresent the Model (for mocking in test)
type ArticleDataAccessLayer interface {
	Validate(a Article) (errs []error)
	Save(a Article) (newID int, err error)
}

// The Model with Database Implementation
type ArticleModel struct {
	DB *pgx.Conn
}

func (model *ArticleModel) Validate(a Article) (errs []error) {
	errs = []error{}
	if a.Body == "" {
		errs = append(errs, errors.New("missing article body"))
	}
	if a.Summary == "" {
		errs = append(errs, errors.New("missing article summary"))
	}
	if a.Title == "" {
		errs = append(errs, errors.New("missing article title"))
	}
	if a.URI == "" {
		errs = append(errs, errors.New("missing article uri"))
	}
	return errs
}

func (model *ArticleModel) Save(a Article) (newID int, err error) {
	stmt := `
		INSERT INTO articles (title, uri, summary, body_md) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err = model.DB.QueryRow(
		context.Background(),
		stmt,
		a.Title, a.URI, a.Summary, a.Body,
	).Scan(&newID)

	if err != nil {
		err = database.TranslateError(err)
	}

	return newID, err
}
