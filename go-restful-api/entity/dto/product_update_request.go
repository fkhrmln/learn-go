package dto

type ProductUpdateRequest struct {
	Id    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required,max=255"`
	Price int    `json:"price" validate:"required"`
}
