package helpers

import (
	"strings"

	"github.com/eduahcb/world-ranks/pkg/api"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var printer = message.NewPrinter(language.BrazilianPortuguese)

func FormatNumber(num int) string {
	return printer.Sprintf("%d", num)
}

func FormatFloat(num float64) string {
	return printer.Sprintf("%.2f", num)
}

func FormatLaguages(languages map[string]string) string {
	var values []string

	for _, lang := range languages {
		values = append(values, lang)
	}

	return strings.Join(values, ", ")
}

func FormatCurrencies(currencies map[string]api.Currency) string {
	var values []string

	for _, currency := range currencies {
		values = append(values, currency.Name)
	}

	return strings.Join(values, ", ")
}
