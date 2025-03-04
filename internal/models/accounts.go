package models

import "github.com/Pyakz/buildbox-api/ent/generated/subscription"

type RegisterAccountStruct struct {
	FirstName    string                    `json:"first_name" validate:"required,min=1,max=200"`
	LastName     string                    `json:"last_name" validate:"required,min=1,max=200"`
	MiddleName   string                    `json:"middle_name" validate:"required,min=1,max=200"`
	Email        string                    `json:"email" validate:"required,email"`
	Password     string                    `json:"password" validate:"required,min=3"`
	Name         string                    `json:"name" validate:"required,min=1,max=200"` // company name
	Birthday     string                    `json:"birthday" validate:"omitempty,datetime"`
	PhoneNumber  string                    `json:"phone_number" validate:"required,e164"`
	PlanID       int                       `json:"plan_id" validate:"required"`
	BillingCycle subscription.BillingCycle `json:"billing_cycle" validate:"required,oneof=monthly yearly"`
}
