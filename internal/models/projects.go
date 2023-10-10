package models

type AddServiceProviderToProjectRequest struct {
	ServiceProviderID int `json:"service_provider_id" validate:"required"`
}
