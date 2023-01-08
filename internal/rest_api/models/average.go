package models

type Average struct {
	Average  float64 `json:"average"`
	DataType string  `json:"datatype"`
}

type Metrics struct {
	Date    string
	Value   string
	Airport string
	Sensor  string
}
