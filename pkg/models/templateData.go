package models

// TemplateData holds the data that gets sent to the template data struct
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
