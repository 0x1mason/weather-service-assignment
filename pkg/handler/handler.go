package handler

import (
	"encoding/json"
	"net/http"
)

type RequestHandler struct {
	NoaaHost string
}

func (rh *RequestHandler) GetForecast(w http.ResponseWriter, req *http.Request) {
	coordinates := req.PathValue("coordinates")

	// https://api.weather.gov/points/{latitude},{longitude}
	res, err := http.Get(rh.NoaaHost + "/points/" + coordinates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.StatusCode != 200 {
		// assume it's a bad location for this exercise
		http.Error(w, "Invalid location", http.StatusBadRequest)
		return
	}

	var pointsResponse PointsResponse
	err = json.NewDecoder(res.Body).Decode(&pointsResponse)
	if err != nil {
		http.Error(w, "Invalid response data", http.StatusInternalServerError)
		return
	}

	res, err = http.Get(pointsResponse.Properties.Forecast)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.StatusCode != 200 {
		// assume it's a bad location for this exercise
		http.Error(w, "Invalid location", http.StatusBadRequest)
		return
	}

	var forecast ForecastResponse
	err = json.NewDecoder(res.Body).Decode(&forecast)
	if err != nil {
		http.Error(w, "Invalid response data", http.StatusInternalServerError)
		return
	}

	resp := Response{
		ShortForecast: forecast.Properties.Periods[0].ShortForecast,
	}

	if forecast.Properties.Periods[0].Temperature >= 80 {
		resp.Characterization = "hot"
	} else if forecast.Properties.Periods[0].Temperature <= 50 {
		resp.Characterization = "cold"
	} else {
		resp.Characterization = "moderate"
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		return
	}
}
