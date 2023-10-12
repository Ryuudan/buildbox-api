package models

import "github.com/Pyakz/buildbox-api/ent/generated/team"

type UpdateTeamPayload struct {
	Name        *string      `json:"name" validate:"omitempty,min=1"`
	Description *string      `json:"description" validate:"omitempty,max=300"`
	Status      *team.Status `json:"status" validate:"omitempty,oneof=active inactive"`
}
