package model

type Exchange struct {
	Id   int    `json:"id"`
	From string `json:"from" validate:"required"`
	To   string `json:"to" validate:"required"`
}

type ExhangeRate struct {
	Id   int     `json:"id"`
	Date string  `json:"date"`
	From string  `json:"from" validate:"required"`
	To   string  `json:"to" validate:"required"`
	Rate float64 `json:"rate"`
}
