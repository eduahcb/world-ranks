package controllers

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/eduahcb/world-ranks/internal/services"
	"github.com/eduahcb/world-ranks/pkg/api"
)

type Data struct {
	Countries       []api.Country
	SelectedRegions map[string]bool
	Status          string
	Sort            string
	Search          string
}

func HomeHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()

		sort := query.Get("sort")
		status := query.Get("status")
		region := query["region"]
		search := query.Get("search")

		if strings.TrimSpace(sort) == "" {
			sort = string(services.SortByPopulation)
		}

		if strings.TrimSpace(status) == "" {
			status = string(services.StatusIndependent)
		}

		selectedRegions := map[string]bool{
			"asia":      false,
			"europe":    false,
			"africa":    false,
			"americas":  false,
			"antarctic": false,
		}

		for _, region := range region {
			selectedRegions[region] = true
		}

		queries := services.QueryParams{
			Sort:   services.SortType(sort),
			Status: services.StatusType(status),
			Region: region,
			Search: search,
		}

		countries, err := services.GetAllCountriesService(queries)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := &Data{
			Countries:       countries,
			SelectedRegions: selectedRegions,
			Status:          status,
			Sort:            sort,
			Search:          search,
		}

		err = tmpl.ExecuteTemplate(w, "home.html", data)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
