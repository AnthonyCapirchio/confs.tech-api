package conference

import (
	"strings"
	"time"
)

type Store struct {
	Entries    []Conference
	Cities     []string
	Countries  []string
	Categories []string
}

type FilterFunc func(entry Conference, index int) bool

func (s *Store) Push(c Conference) {
	s.Entries = append(s.Entries, c)

	if !stringExists(s.Cities, c.City) {
		s.Cities = append(s.Cities, c.City)
	}

	if !stringExists(s.Countries, c.Country) {
		s.Countries = append(s.Countries, c.Country)
	}

	if !stringExists(s.Categories, c.Category) {
		s.Categories = append(s.Categories, c.Category)
	}
}

func stringExists(slice []string, match string) bool {
	for _, entry := range slice {
		if entry == match {
			return true
		}
	}
	return false
}

func (s *Store) Filter(filterFn FilterFunc) *Store {
	output := Store{
		Entries: []Conference{},
	}

	for index, entry := range s.Entries {
		if filterFn(entry, index) {
			output.Entries = append(output.Entries, entry)
		}
	}

	return &output
}

func (s *Store) FilterByCountry(country string) *Store {
	return s.Filter(func(entry Conference, index int) bool {
		return strings.ToUpper(entry.Country) == strings.ToUpper(country)
	})
}

func (s *Store) FilterByCity(city string) *Store {
	return s.Filter(func(entry Conference, index int) bool {
		return strings.ToUpper(entry.City) == strings.ToUpper(city)
	})
}

func (s *Store) FilterByCategory(category string) *Store {
	return s.Filter(func(entry Conference, index int) bool {
		return strings.ToUpper(entry.Category) == strings.ToUpper(category)
	})
}

func (s *Store) FilterByName(name string, exact bool) *Store {
	return s.Filter(func(entry Conference, index int) bool {
		if exact {
			return strings.ToUpper(entry.Name) == strings.ToUpper(name)
		} else {
			return strings.Contains(strings.ToUpper(entry.Name), strings.ToUpper(name))
		}
	})
}

func (s *Store) FilterByDateRange(start, end time.Time, threshold int) *Store {
	startLimit := 0 - (threshold * 24)
	endLimit := 0 + (threshold * 24)

	return s.Filter(func(entry Conference, index int) bool {
		startDelta := entry.StartDate.Sub(start).Hours()
		endDelta := entry.EndDate.Sub(end).Hours()

		return int(startDelta) >= startLimit && int(endDelta) <= endLimit
	})
}
