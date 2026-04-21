package controllers

import (
	"net/http"
	"text/template"

	"github.com/eduahcb/world-ranks/pkg/api"
)

type DetailsData struct {
	Country api.Country
}

func DetailsHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")

		country, err := api.GetCountryByCode(id)

		if err != nil {
			println("Error fetching country details:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := DetailsData{
			Country: country,
		}

		err = tmpl.ExecuteTemplate(w, "details.html", data)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
