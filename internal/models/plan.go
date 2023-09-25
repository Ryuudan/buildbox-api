package models

type CreatePlanRequest struct {
	Name        string  `json:"name" validate:"required,min=3"`
	Description string  `json:"description" validate:"required,min=10"`
	Price       float64 `json:"price,omitempty" validate:"omitempty,gte=0"`
}
