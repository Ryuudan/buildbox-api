package projects

import "time"

type Project struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	AccountID string `json:"accountId,omitempty" bson:"accountId,omitempty"`
	ClientID  string `json:"clientId,omitempty" bson:"clientId,omitempty"`
	ManagerID string `json:"managerId,omitempty" bson:"managerId,omitempty"`
	CreatedBy string `json:"createdBy,omitempty" bson:"createdBy,omitempty"`

	Name        string  `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=1,max=100"`
	Description string  `json:"description,omitempty" bson:"description,omitempty" validate:"omitempty,min=1,max=200"`
	Notes       string  `json:"notes,omitempty" bson:"notes,omitempty" validate:"omitempty,min=1,max=200"`
	Status      string  `json:"status,omitempty" bson:"status,omitempty" validate:"omitempty,oneof=planning in_progress on_hold completed cancelled delayed under_review pending_approval in_testing emergency on_schedule behind_schedule in_review archived in_negotiation"`
	Location    string  `json:"location,omitempty" bson:"location,omitempty"`
	Budget      float64 `json:"budget,omitempty" bson:"budget,omitempty"`
	Deleted     bool    `json:"deleted,omitempty" bson:"deleted,omitempty"`

	StartDate *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   *time.Time `json:"endDate,omitempty" bson:"endDate,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

type CreateProject struct {
	ID          string     `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string     `json:"name,omitempty" validate:"required,min=1,max=100"`
	Description string     `json:"description,omitempty" validate:"omitempty,min=1,max=200"`
	Notes       string     `json:"notes,omitempty" validate:"omitempty,max=200"`
	Status      string     `json:"status,omitempty" validate:"omitempty,oneof=planning in_progress on_hold completed cancelled delayed under_review pending_approval in_testing emergency on_schedule behind_schedule in_review archived in_negotiation"`
	ManagerID   string     `json:"managerId,omitempty" bson:"managerId,omitempty"`
	Location    string     `json:"location,omitempty" bson:"location,omitempty"`
	Budget      float64    `json:"budget,omitempty" bson:"budget,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty" bson:"endDate,omitempty"`
}
