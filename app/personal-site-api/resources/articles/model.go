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
	All() ([]Article, error)
	Get(uri string) (Article, error)
	Validate(a Article) (errs []error)
	Save(a Article) (newID int, err error)
}

// The Model with Database Implementation
type ArticleModel struct {
	DB *pgx.Conn
}

func (model *ArticleModel) All() (articles []Article, err error) {
	stmt := `
		SELECT id, uri, title, summary, body_md
		FROM articles;
	`
	rows, err := model.DB.Query(context.Background(), stmt)
	if err != nil {
		return articles, err
	}

	for rows.Next() {
		var article Article
		rows.Scan(&article.ID, &article.URI, &article.Title, &article.Summary, &article.Body)
		articles = append(articles, article)
	}
	return articles, err
}

func (model *ArticleModel) Get(uri string) (article Article, err error) {
	stmt := `
		SELECT id, uri, title, summary, body_md
		FROM articles
		WHERE uri = $1;
	`
	err = model.DB.QueryRow(context.Background(), stmt, uri).Scan(&article.ID, &article.URI, &article.Title, &article.Summary, &article.Body)
	return article, err
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
