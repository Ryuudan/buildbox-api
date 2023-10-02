package models

type OrderFields struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type Filters struct {
	Order []OrderFields `json:"order"`
}
