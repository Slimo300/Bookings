package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	form := New(url.Values{})

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	form := New(url.Values{})

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	form = New(url.Values{})
	form.Add("a", "a")
	form.Add("b", "b")
	form.Add("c", "c")

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("form shows invalid when should be valid")
	}
}

func TestHas(t *testing.T) {
	urlValues := make(url.Values)

	form := New(urlValues)

	if form.Has("key") {
		t.Error("got true when should be false")
	}

	form.Add("key", "value")

	if !form.Has("key") {
		t.Error("got false when should be true")
	}

}

func TestMinLength(t *testing.T) {
	urlValues := make(url.Values)

	form := New(urlValues)
	form.Add("a", "aa")

	if form.MinLength("a", 3) {
		t.Error("got true when should be false")
	}

	form.Del("a")
	form.Add("a", "aaa")

	if !form.MinLength("a", 3) {
		t.Error("got false when should be true")
	}

}

func TestIsEmail(t *testing.T) {
	urlValues := make(url.Values)

	form := New(urlValues)
	form.Add("a", "aa@aa")

	if form.IsEmail("a") {
		t.Error("got true when should be false")
	}

	form.Del("a")
	form.Add("a", "aa@aa.com")

	if !form.MinLength("a", 3) {
		t.Error("got false when should be true")
	}
}
