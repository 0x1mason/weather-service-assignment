package handler

type PointsResponseProperties struct {
	Forecast string `json:"forecast"`
}

type PointsResponse struct {
	Properties PointsResponseProperties `json:"properties"`
}

type ForecastPropertiesPeriod struct {
	Temperature   int    `json:"temperature"`
	ShortForecast string `json:"shortForecast"`
}

type ForecastProperties struct {
	Periods []ForecastPropertiesPeriod `json:"periods"`
}

type ForecastResponse struct {
	Properties ForecastProperties `json:"properties"`
}

type Response struct {
	Characterization string `json:"characterization"`
	ShortForecast    string `json:"shortForecast"`
}
