package services

import (
	"cmp"
	"slices"
	"strings"

	"github.com/eduahcb/world-ranks/pkg/api"
)

type SortType string
type StatusType string

const (
	SortByPopulation  SortType   = "population"
	SortByName        SortType   = "name"
	SortByArea        SortType   = "area"
	StatusIndependent StatusType = "independent"
	StatusMemberUN    StatusType = "member_un"
)

type QueryParams struct {
	Sort   SortType
	Status StatusType
	Region []string
	Search string
}

func GetAllCountriesService(queries QueryParams) ([]api.Country, error) {
	sort := queries.Sort
	status := queries.Status
	regions := queries.Region

	if sort == "" {
		sort = SortByPopulation
	}

	if status == "" {
		status = StatusIndependent
	}

	normalizedRegions := make([]string, len(regions))

	for i, r := range regions {
		normalizedRegions[i] = strings.ToLower(r)
	}

	countries, err := api.GetCountries()

	if err != nil {
		return nil, err
	}

	countries = slices.DeleteFunc(countries, func(c api.Country) bool {
		if status == StatusIndependent && !c.Independent {
			return true
		}

		if status == StatusMemberUN && !c.UnMember {
			return true
		}

		if len(normalizedRegions) > 0 {
			regionLower := strings.ToLower(c.Region)

			if !slices.Contains(normalizedRegions, regionLower) {
				return true
			}
		}

		if queries.Search != "" {
			if !strings.Contains(strings.ToLower(c.Name.Common), strings.ToLower(queries.Search)) {
				return true
			}
		}

		return false
	})

	slices.SortFunc(countries, func(a, b api.Country) int {
		switch sort {
		case SortByPopulation:
			return cmp.Compare(b.Population, a.Population)
		case SortByName:
			return cmp.Compare(a.Name.Common, b.Name.Common)
		case SortByArea:
			return cmp.Compare(b.Area, a.Area)
		default:
			return cmp.Compare(b.Population, a.Population)
		}
	})

	return countries, nil
}
