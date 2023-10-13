package models

type UpdatePlanRequest struct {
	Name        *string  `json:"name" validate:"omitempty,min=3"`
	Description *string  `json:"description" validate:"omitempty,min=10"`
	Price       *float64 `json:"price" validate:"omitempty,gte=0"`
}
