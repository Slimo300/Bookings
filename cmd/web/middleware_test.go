package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	csrfHandler := NoSurf(&myHandler{})

	switch csrfHandler.(type) {
	case http.Handler:

	default:
		t.Error("Type is not http.Handler")
	}
}

func TestSessionLoad(t *testing.T) {
	sessionHandler := SessionLoad(&myHandler{})

	switch sessionHandler.(type) {
	case http.Handler:

	default:
		t.Error("Type is not http.Handler")
	}
}
