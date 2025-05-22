package models

type Song struct {
	Name     string `json:"name"`
	Singer   string `json:"singer"`
	Code     string `json:"code"`
	Favorite bool   `json:"favorite"`
}
