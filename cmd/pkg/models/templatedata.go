package models

// TemplateData holds data for handlers
type TemplateData struct {
	StrMap    map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string //cross-site forgery token
	Flash     string
	Warning   string
	Error     string
}
