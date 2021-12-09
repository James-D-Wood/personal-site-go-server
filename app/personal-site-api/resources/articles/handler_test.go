package articles

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jdwoo/personal-site-go-server/internal/webserverutils"
)

type MockArticleModel struct {
	articles         []Article
	validationErrors []error
	fetchError       error
	updateError      error
	saveError        error
}

func (model MockArticleModel) All() ([]Article, error) { return model.articles, model.fetchError }
func (model MockArticleModel) Get(uri string) (Article, error) {
	for _, article := range model.articles {
		if article.URI == uri {
			return article, nil
		}
	}
	return Article{}, errors.New("no rows in result set")
}
func (model MockArticleModel) Update(uri string, a Article) (int, error) {
	for _, article := range model.articles {
		fmt.Println(article.URI)
		if article.URI == uri {
			return article.ID, model.updateError
		}
	}
	return 0, errors.New("no rows in result set")
}
func (model MockArticleModel) Validate(a Article) []error  { return model.validationErrors }
func (model MockArticleModel) Save(a Article) (int, error) { return 1, model.saveError }

func TestGetArticlesHandlerSuccess(t *testing.T) {
	model := MockArticleModel{
		articles: []Article{
			{
				ID:      1,
				URI:     "some-article-1",
				Title:   "Some Article: Part 1",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
			{
				ID:      2,
				URI:     "some-article-2",
				Title:   "Some Article: Part 2",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
		},
	}

	req, err := http.NewRequest("GET", "/api/v1/articles", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	GetArticlesHandler(model).ServeHTTP(rr, req)

	expectedCode := 200
	expectedRespLen := 2

	if rr.Code != expectedCode {
		t.Errorf("expected status code %d but received %d", expectedCode, rr.Result().StatusCode)
	}

	var respBody []Article
	err = json.Unmarshal(rr.Body.Bytes(), &respBody)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(respBody) != expectedRespLen {
		t.Errorf("expected response body ID %d but received %d", expectedRespLen, len(respBody))
	}
}

func TestGetArticlesHandlerFailure(t *testing.T) {
	model := MockArticleModel{
		articles: []Article{
			{
				ID:      1,
				URI:     "some-article-1",
				Title:   "Some Article: Part 1",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
			{
				ID:      2,
				URI:     "some-article-2",
				Title:   "Some Article: Part 2",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
		},
		fetchError: errors.New("unexpected error"),
	}

	req, err := http.NewRequest("GET", "/api/v1/articles", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	GetArticlesHandler(model).ServeHTTP(rr, req)

	expectedCode := 500

	if rr.Code != expectedCode {
		t.Errorf("expected status code %d but received %d", expectedCode, rr.Result().StatusCode)
	}
}

func TestGetArticleHandlerSuccess(t *testing.T) {
	model := MockArticleModel{
		articles: []Article{
			{
				ID:      1,
				URI:     "some-article-1",
				Title:   "Some Article: Part 1",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
			{
				ID:      2,
				URI:     "some-article-2",
				Title:   "Some Article: Part 2",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
		},
	}

	targetArticle := "some-article-1"

	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/article/%s", targetArticle), nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"articleURI": targetArticle})

	rr := httptest.NewRecorder()
	GetArticleHandler(model).ServeHTTP(rr, req)

	expectedCode := 200

	if rr.Code != expectedCode {
		t.Errorf("expected status code %d but received %d", expectedCode, rr.Result().StatusCode)
	}

	var respBody Article
	err = json.Unmarshal(rr.Body.Bytes(), &respBody)
	if err != nil {
		t.Errorf(err.Error())
	}
	if respBody.URI != targetArticle {
		t.Errorf("expected response article '%s' but received '%s'", targetArticle, respBody.URI)
	}
}

func TestGetArticleHandlerMissingArticle(t *testing.T) {
	model := MockArticleModel{
		articles: []Article{
			{
				ID:      1,
				URI:     "some-article-1",
				Title:   "Some Article: Part 1",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
			{
				ID:      2,
				URI:     "some-article-2",
				Title:   "Some Article: Part 2",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
		},
	}

	targetArticle := "some-article-3"

	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/article/%s", targetArticle), nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"articleURI": targetArticle})

	rr := httptest.NewRecorder()
	GetArticleHandler(model).ServeHTTP(rr, req)

	expectedCode := 404

	if rr.Code != expectedCode {
		t.Errorf("expected status code %d but received %d", expectedCode, rr.Result().StatusCode)
	}
}

func TestGetArticleHandlerDBError(t *testing.T) {
	model := MockArticleModel{
		articles: []Article{
			{
				ID:      1,
				URI:     "some-article-1",
				Title:   "Some Article: Part 1",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
			{
				ID:      2,
				URI:     "some-article-2",
				Title:   "Some Article: Part 2",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
		},
		fetchError: errors.New("unexpected error"),
	}

	targetArticle := "some-article-3"

	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/article/%s", targetArticle), nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"articleURI": targetArticle})

	rr := httptest.NewRecorder()
	GetArticleHandler(model).ServeHTTP(rr, req)

	expectedCode := 404

	if rr.Code != expectedCode {
		t.Errorf("expected status code %d but received %d", expectedCode, rr.Result().StatusCode)
	}
}
func TestCreateArticleHandlerSuccess(t *testing.T) {
	model := MockArticleModel{}
	handler := CreateArticleHandler(model)

	reqBody := `
		{
			"title": "Some article title",
			"uri": "some-article-title",
			"summary": "some summary",
			"body": "some thoughts"
		}
	`
	req, err := http.NewRequest("POST", "/api/v1/article", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != 201 {
		t.Errorf("expected status code %d but received %d", 201, rr.Result().StatusCode)
	}

	var respBody Article
	err = json.Unmarshal(rr.Body.Bytes(), &respBody)
	if err != nil {
		t.Errorf(err.Error())
	}
	if respBody.ID != 1 {
		t.Errorf("expected response body ID %d but received %d", 1, respBody.ID)
	}
}
func TestCreateArticleHandlerValidateError(t *testing.T) {

	model := MockArticleModel{
		validationErrors: []error{
			errors.New("test error 1"),
			errors.New("test error 2"),
		},
	}
	handler := CreateArticleHandler(model)

	reqBody := `
		{
			"title": "Some article title",
			"uri": "some-article-title",
			"summary": "some summary",
			"body": "some thoughts"
		}
	`
	req, err := http.NewRequest("POST", "/api/v1/article", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	expectedCode := 422
	expectedBody := "Invalid Request Body: test error 1, test error 2"
	if rr.Code != expectedCode {
		t.Errorf("expected status code %d but received %d", expectedCode, rr.Result().StatusCode)
	}

	respBody := strings.TrimSpace(rr.Body.String())
	if respBody != expectedBody { // TODO: remove contains for strict equality
		t.Errorf("expected response body '%s' but received '%s'", expectedBody, respBody)
	}
}

func TestCreateArticleHandlerSaveErrorDuplicateEntry(t *testing.T) {

	model := MockArticleModel{
		saveError: errors.New("Invalid Request Body: violates unique constraint"),
	}
	handler := CreateArticleHandler(model)

	reqBody := `
		{
			"title": "Some article title",
			"uri": "some-article-title",
			"summary": "some summary",
			"body": "some thoughts"
		}
	`
	req, err := http.NewRequest("POST", "/api/v1/article", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	expectedCode := 422
	expectedBody := model.saveError.Error()

	if rr.Code != expectedCode {
		t.Errorf("expected status code %d but received %d", expectedCode, rr.Result().StatusCode)
	}

	respBody := strings.TrimSpace(rr.Body.String())
	if respBody != expectedBody {
		t.Errorf("expected response body '%s' but received '%s'", expectedBody, respBody)
	}
}

func TestCreateArticleHandlerSaveError(t *testing.T) {

	model := MockArticleModel{
		saveError: errors.New("unexpected error"),
	}
	handler := CreateArticleHandler(model)

	reqBody := `
		{
			"title": "Some article title",
			"uri": "some-article-title",
			"summary": "some summary",
			"body": "some thoughts"
		}
	`
	req, err := http.NewRequest("POST", "/api/v1/article", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	expectedCode := 500
	expectedBody := model.saveError.Error()

	if rr.Code != expectedCode {
		t.Errorf("expected status code %d but received %d", expectedCode, rr.Result().StatusCode)
	}

	respBody := strings.TrimSpace(rr.Body.String())
	if respBody != expectedBody {
		t.Errorf("expected response body '%s' but received '%s'", expectedBody, respBody)
	}
}

func TestUpdateArticleHandlerSuccess(t *testing.T) {

	model := MockArticleModel{
		articles: []Article{
			{
				ID:      1,
				URI:     "some-article-1",
				Title:   "Some Article: Part 1",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
			{
				ID:      2,
				URI:     "some-article-2",
				Title:   "Some Article: Part 2",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
		},
	}
	handler := UpdateArticleHandler(model)

	targetURI := "some-article-2"
	reqBody := fmt.Sprintf(`
		{
			"title": "Some article title",
			"uri": "%s",
			"summary": "some summary",
			"body": "some thoughts"
		}
	`, targetURI)
	req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v1/article/%s", targetURI), bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"articleURI": targetURI})

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	expectedCode := 200

	if rr.Code != expectedCode {
		t.Errorf("expected status code %d but received %d", expectedCode, rr.Result().StatusCode)
	}
}

func TestUpdateArticleHandlerMismatchURI(t *testing.T) {

	expectedBody := "URI in URL var does not match URI in body"

	model := MockArticleModel{
		articles: []Article{
			{
				ID:      1,
				URI:     "some-article-1",
				Title:   "Some Article: Part 1",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
			{
				ID:      2,
				URI:     "some-article-2",
				Title:   "Some Article: Part 2",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
		},
		updateError: webserverutils.NewRequestError(expectedBody),
	}
	handler := UpdateArticleHandler(model)

	targetURI := "some-article-2"
	reqBody := `
		{
			"title": "Some article title",
			"uri": "some-other-uri",
			"summary": "some summary",
			"body": "some thoughts"
		}
	`
	req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v1/article/%s", targetURI), bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"articleURI": targetURI})

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	expectedCode := 422
	if rr.Code != expectedCode {
		t.Errorf("expected status code %d but received %d", expectedCode, rr.Result().StatusCode)
	}

	respBody := strings.TrimSpace(rr.Body.String())
	if !strings.Contains(respBody, expectedBody) {
		t.Errorf("expected response body '%s' but received '%s'", expectedBody, respBody)
	}
}

func TestUpdateArticleHandlerInvalidPayload(t *testing.T) {

	model := MockArticleModel{
		validationErrors: []error{
			errors.New("test error 1"),
			errors.New("test error 2"),
		},
	}
	handler := UpdateArticleHandler(model)

	targetURI := "some-article-2"
	reqBody := fmt.Sprintf(`
		{
			"title": "Some article title",
			"uri": "%s",
			"summary": "some summary",
			"body": "some thoughts"
		}
	`, targetURI)
	req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v1/article/%s", targetURI), bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"articleURI": targetURI})

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	expectedCode := 422
	expectedBody := "Invalid Request Body: test error 1, test error 2"
	if rr.Code != expectedCode {
		t.Errorf("expected status code %d but received %d", expectedCode, rr.Result().StatusCode)
	}

	respBody := strings.TrimSpace(rr.Body.String())
	if respBody != expectedBody { // TODO: remove contains for strict equality
		t.Errorf("expected response body '%s' but received '%s'", expectedBody, respBody)
	}
}
func TestUpdateArticleHandlerNotFound(t *testing.T) {

	model := MockArticleModel{
		articles: []Article{
			{
				ID:      1,
				URI:     "some-article-1",
				Title:   "Some Article: Part 1",
				Summary: "A Short Summary",
				Body:    "A Body",
			},
		},
	}
	handler := UpdateArticleHandler(model)

	targetURI := "some-article-2"
	reqBody := fmt.Sprintf(`
		{
			"title": "Some article title",
			"uri": "%s",
			"summary": "some summary",
			"body": "some thoughts"
		}
	`, targetURI)
	req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v1/article/%s", targetURI), bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"articleURI": targetURI})

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	expectedCode := 404
	expectedBody := "article not found"
	if rr.Code != expectedCode {
		t.Errorf("expected status code %d but received %d", expectedCode, rr.Result().StatusCode)
	}

	respBody := strings.TrimSpace(rr.Body.String())
	if respBody != expectedBody {
		t.Errorf("expected response body '%s' but received '%s'", expectedBody, respBody)
	}
}
