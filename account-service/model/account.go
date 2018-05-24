package model

type Account struct {
	Id string `json:"id"` //The case of the first letter denotes scoping (Upper-case == public, lower-case package-scoped)
	Name string  `json:"name"`
}