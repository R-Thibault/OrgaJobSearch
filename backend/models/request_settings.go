package models

type RequestSettings struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Order  string `json:"order"`
	Where  string `json:"where"`
}
