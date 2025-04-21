package models

type Todo struct {
	ID       string `json:"id"`
	Titulo   string `json:"título"`
	Compelto bool   `json:"compelto"`
}
