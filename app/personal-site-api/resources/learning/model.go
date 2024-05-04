package learning

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/database"
)

type Reference struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Url    string `json:"url"`
}

type Lesson struct {
	ID             int         `json:"id"`
	Topic          string      `json:"topic" validate:"required"`
	Tags           []string    `json:"tags"`
	References     []Reference `json:"references"`
	Takeaways      []string    `json:"takeaways"` // what was learned
	Questions      []string    `json:"questions"` // what I still want to know or reflect on
	Exercises      []string    `json:"exercises"` // how I should refresh myself in the future
	CreatedAt      time.Time   `json:"createdAt"`
	LastModifiedAt *time.Time  `json:"lastModifiedAt"`
}

type LessonDataAccessLayer interface {
	Create(l Lesson) (err error)
	// Get(lessonID string) (lesson Lesson, err error)
	// Where(query string, tags []string) (lessons []Lesson, err error)
}

type LessonModel struct {
	DB *pgx.Conn
}

func (model *LessonModel) Create(l Lesson) (err error) {
	// UNIQUE constraint on topic
	// tags table
	// references table

	stmt := `
		INSERT INTO lessons (topic, takeaways, questions, exercises, dt_created) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var id int
	err = model.DB.QueryRow(
		context.Background(),
		stmt,
		l.Topic, l.Takeaways, l.Questions, l.Exercises, l.CreatedAt,
	).Scan(&id)

	if err != nil {
		return database.TranslateError(err)
	}

	for _, tag := range l.Tags {
		var t string
		stmt = `
			INSERT INTO lesson_tags (lesson_id, tag)
			VALUES ($1, $2)
			RETURNING tag
		`
		err = model.DB.QueryRow(
			context.Background(),
			stmt,
			id, tag,
		).Scan(&t)

		if err != nil {
			return database.TranslateError(err)
		}
	}

	for _, reference := range l.References {
		var t string
		stmt = `
			INSERT INTO lesson_references (lesson_id, title, author, url)
			VALUES ($1, $2, $3, $4)
			RETURNING title
		`
		err = model.DB.QueryRow(
			context.Background(),
			stmt,
			id, reference.Title, reference.Author, reference.Url,
		).Scan(&t)

		if err != nil {
			return database.TranslateError(err)
		}
	}

	return err
}
