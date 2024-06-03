package dto

type ProductResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
