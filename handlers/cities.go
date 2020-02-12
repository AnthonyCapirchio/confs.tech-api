package handlers

import (
	"net/http"

	"squidward.confs.tech/conference"
)

func GetCitiesHandler(conferencesService *conference.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendCollection(w, conferencesService.Cities)
	}
}

func GetCountriesHandler(conferencesService *conference.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendCollection(w, conferencesService.Countries)
	}
}

func GetCategoriesHandler(conferencesService *conference.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendCollection(w, conferencesService.Categories)
	}
}
