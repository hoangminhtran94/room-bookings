package models

import "github.com/hoangminhtran94/room-bookings/internal/form"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flask     string
	Warning   string
	Error     string
	Form *form.Form
}
