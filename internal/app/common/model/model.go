package model

type Measure struct {
	SensorID    string  `json:"sensorId"`
	AirportID   string  `json:"airportId"`
	MeasureType string  `json:"measureType"`
	Value       float64 `json:"_value"`
	Timestamp   string  `json:"timestamp"`
}
