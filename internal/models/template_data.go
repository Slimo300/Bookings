package models

import "github.com/Slimo300/Bookings/internal/forms"

type TemplateData struct {
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      forms.Form
}
