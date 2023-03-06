package forms

import "testing"

func TestAdd(t *testing.T) {
	err := make(errors)

	err.Add("key", "value")
}

func TestGet(t *testing.T) {
	err := make(errors)
	val := err.Get("key")

	if val != "" {
		t.Error("Returned value should be an empty string")
	}

	err.Add("key", "value")

	val = err.Get("key")
	if val != "value" {
		t.Error("Returned value should be \"value\"")
	}
}
