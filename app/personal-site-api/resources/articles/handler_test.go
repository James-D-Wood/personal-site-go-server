package articles

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockArticleModel struct {
	validationErrors []error
	saveError        error
}

func (model MockArticleModel) Validate(a Article) []error  { return model.validationErrors }
func (model MockArticleModel) Save(a Article) (int, error) { return 1, model.saveError }

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
