package controllers

import (
	"html/template"
	"net/http"

	"github.com/grasieliw/GWStore/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}
