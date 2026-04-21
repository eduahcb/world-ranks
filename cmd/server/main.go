package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/eduahcb/world-ranks/internal/controllers"
	"github.com/eduahcb/world-ranks/pkg/helpers"
)

func main() {
	tmpl := template.Must(
		template.New("").
			Funcs(template.FuncMap{
				"formatInt":        helpers.FormatNumber,
				"formatFloat":      helpers.FormatFloat,
				"join":             strings.Join,
				"formatLangs":      helpers.FormatLaguages,
				"formatCurrencies": helpers.FormatCurrencies,
			}).
			ParseGlob("internal/views/*.html"),
	)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", controllers.HomeHandler(tmpl))
	mux.HandleFunc("GET /details/{id}", controllers.DetailsHandler(tmpl))

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	}

	fs := http.FileServer(http.Dir("internal/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	fmt.Printf("server is running on port: %s \n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
