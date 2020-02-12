package handlers

import (
	"net/http"
	"strconv"
	"time"

	"squidward.confs.tech/conference"
)

func GetConferencesHandler(conferencesService *conference.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		output := conferencesService

		if country := params.Get("country"); country != "" {
			output = output.FilterByCountry(country)
		}
		if city := params.Get("city"); city != "" {
			output = output.FilterByCity(city)
		}
		if category := params.Get("category"); category != "" {
			output = output.FilterByCategory(category)
		}

		startParameter := params.Get("start")
		endParameter := params.Get("end")
		thresholdParameter := params.Get("threshold")
		if startParameter != "" && endParameter != "" {
			threshold := 1
			if thresholdParameter != "" {
				threshold, _ = strconv.Atoi(thresholdParameter)
			}

			start, _ := time.Parse("2006-01-02", startParameter)
			end, _ := time.Parse("2006-01-02", endParameter)
			output = output.FilterByDateRange(start, end, threshold)
		}

		sendCollection(w, output.Entries)
	}
}
