package articles

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/database"
	"github.com/jdwoo/personal-site-go-server/internal/webserverutils"
)

// A struct to model the object
type Article struct {
	ID          int        `json:"id"`
	URI         string     `json:"uri"`
	Title       string     `json:"title"`
	Summary     string     `json:"summary"`
	Body        string     `json:"body"`
	DateCreated time.Time  `json:"dateCreated"`
	DateUpdated *time.Time `json:"dateUpdated"`
}

// An interface to refresent the Model (for mocking in test)
type ArticleDataAccessLayer interface {
	All() ([]Article, error)
	Get(uri string) (result Article, err error)
	Save(a Article) (result Article, err error)
	Update(uri string, a Article) (result Article, err error)
	Validate(a Article) (errs []error)
}

// The Model with Database Implementation
type ArticleModel struct {
	DB *pgx.Conn
}

func (model *ArticleModel) All() (articles []Article, err error) {

	stmt := `
		SELECT id, uri, title, summary, body_md, dt_created, dt_updated
		FROM articles
		ORDER BY dt_created DESC;
	`
	rows, err := model.DB.Query(context.Background(), stmt)
	if err != nil {
		return articles, err
	}

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID, &article.URI, &article.Title, &article.Summary,
			&article.Body, &article.DateCreated, &article.DateUpdated)
		articles = append(articles, article)
	}
	return articles, err
}

func (model *ArticleModel) Get(uri string) (article Article, err error) {
	stmt := `
		SELECT id, uri, title, summary, body_md, dt_created, dt_updated
		FROM articles
		WHERE uri = $1;
	`
	err = model.DB.QueryRow(context.Background(), stmt, uri).
		Scan(&article.ID, &article.URI, &article.Title, &article.Summary,
			&article.Body, &article.DateCreated, &article.DateUpdated)
	return article, err
}

func (model *ArticleModel) Update(uri string, a Article) (result Article, err error) {
	if uri != a.URI {
		return Article{}, webserverutils.NewRequestError("URI in path does not match URI in body.")
	}

	todayDate := time.Now()
	stmt := `
		UPDATE articles
		SET title=$1, summary=$2, body_md=$3, dt_updated=$4
		WHERE uri=$5
		RETURNING id, title, summary, body_md, dt_created, dt_updated
	`
	err = model.DB.QueryRow(
		context.Background(),
		stmt,
		a.Title, a.Summary, a.Body, todayDate, uri,
	).Scan(&result.ID, &result.Title, &result.Summary, &result.Body, &result.DateCreated, &result.DateUpdated)
	return result, err
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

func (model *ArticleModel) Save(a Article) (newArticle Article, err error) {
	newArticle = a

	stmt := `
		INSERT INTO articles (title, uri, summary, body_md, dt_created) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, dt_created
	`
	todayDate := time.Now()

	err = model.DB.QueryRow(
		context.Background(),
		stmt,
		a.Title, a.URI, a.Summary, a.Body, todayDate,
	).Scan(&newArticle.ID, &newArticle.DateCreated)

	if err != nil {
		err = database.TranslateError(err)
	}

	return newArticle, err
}
