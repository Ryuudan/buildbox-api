package projects

import (
	"time"

	"github.com/Pyakz/buildbox-api/ent/project"
)

// for struct validation
type Project struct {
	ID          string         `json:"id,omitempty"`
	AccountID   string         `json:"accountId,omitempty"`
	ClientID    string         `json:"clientId,omitempty"`
	ManagerID   string         `json:"managerId,omitempty"`
	CreatedBy   string         `json:"createdBy,omitempty"`
	Name        string         `json:"name,omitempty" validate:"required,min=1,max=100"`
	Description string         `json:"description,omitempty" validate:"omitempty,min=1,max=200"`
	Notes       string         `json:"notes,omitempty" validate:"omitempty,min=1,max=200"`
	Status      project.Status `json:"status,omitempty" validate:"omitempty,oneof=planning in_progress on_hold completed cancelled delayed under_review pending_approval in_testing emergency on_schedule behind_schedule in_review archived in_negotiation"`
	Location    string         `json:"location,omitempty"`
	Budget      float64        `json:"budget,omitempty"`
	Deleted     bool           `json:"deleted,omitempty"`
	StartDate   *time.Time     `json:"startDate,omitempty"`
	EndDate     *time.Time     `json:"endDate,omitempty"`
	UpdatedAt   *time.Time     `json:"updatedAt,omitempty"`
	CreatedAt   *time.Time     `json:"createdAt,omitempty"`
}
