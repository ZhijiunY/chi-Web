package models

import "github.com/ZhijiunY/chi-web/internal/forms"

type PageData struct {
	StrMap          map[string]string
	IntMap          map[string]int
	FltMap          map[string]float32
	DataMap         map[string]interface{}
	CSRFToken       string
	Waring          string
	Error           string
	Form            *forms.Form
	Data            map[string]interface{}
	IsAuthenticated int
}
