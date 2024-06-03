package dto

type ProductCreateRequest struct {
	Name  string `json:"name"  validate:"required,max=255"`
	Price int    `json:"price" validate:"required"`
}
