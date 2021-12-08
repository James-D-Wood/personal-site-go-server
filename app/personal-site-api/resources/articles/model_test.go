package articles

import "testing"

func TestValidateSuccess(t *testing.T) {
	model := ArticleModel{}
	a := Article{
		URI:     "some-uri",
		Title:   "Some title",
		Summary: "some-ary",
		Body:    "some body",
	}
	errs := model.Validate(a)
	if len(errs) > 0 {
		t.Errorf("Expected no errors during validation but received %d", len(errs))
	}
}

func TestValidateWithErrors(t *testing.T) {
	model := ArticleModel{}
	a := Article{
		URI:     "some-uri",
		Title:   "Some title",
		Summary: "",
		Body:    "",
	}
	expectedErrCount := 2
	errs := model.Validate(a)
	if len(errs) != expectedErrCount {
		t.Errorf("Expected %d errors during validation but received %d", expectedErrCount, len(errs))
	}
}
