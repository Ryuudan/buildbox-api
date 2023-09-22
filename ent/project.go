// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Pyakz/buildbox-api/ent/project"
	"github.com/google/uuid"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID string `json:"account_id,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID *string `json:"client_id,omitempty"`
	// ManagerID holds the value of the "manager_id" field.
	ManagerID *string `json:"manager_id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty" validate:"required,min=1,max=100"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty" validate:"omitempty,min=1,max=200"`
	// Notes holds the value of the "notes" field.
	Notes *string `json:"notes,omitempty" validate:"omitempty,min=1,max=200"`
	// Status holds the value of the "status" field.
	Status project.Status `json:"status,omitempty" validate:"omitempty,oneof=planning in_progress on_hold completed cancelled delayed under_review pending_approval in_testing emergency on_schedule behind_schedule in_review archived in_negotiation"`
	// Location holds the value of the "location" field.
	Location *string `json:"location,omitempty"`
	// Budget holds the value of the "budget" field.
	Budget float64 `json:"budget,omitempty" validate:"omitempty,gte=0"`
	// Deleted holds the value of the "deleted" field.
	Deleted *bool `json:"deleted,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate time.Time `json:"start_date,omitempty" validate:"ltefield=EndDate"`
	// EndDate holds the value of the "end_date" field.
	EndDate *time.Time `json:"end_date,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case project.FieldDeleted:
			values[i] = new(sql.NullBool)
		case project.FieldBudget:
			values[i] = new(sql.NullFloat64)
		case project.FieldID:
			values[i] = new(sql.NullInt64)
		case project.FieldAccountID, project.FieldClientID, project.FieldManagerID, project.FieldCreatedBy, project.FieldName, project.FieldDescription, project.FieldNotes, project.FieldStatus, project.FieldLocation:
			values[i] = new(sql.NullString)
		case project.FieldStartDate, project.FieldEndDate, project.FieldUpdatedAt, project.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case project.FieldUUID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Project fields.
func (pr *Project) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case project.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				pr.UUID = *value
			}
		case project.FieldAccountID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				pr.AccountID = value.String
			}
		case project.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				pr.ClientID = new(string)
				*pr.ClientID = value.String
			}
		case project.FieldManagerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manager_id", values[i])
			} else if value.Valid {
				pr.ManagerID = new(string)
				*pr.ManagerID = value.String
			}
		case project.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pr.CreatedBy = value.String
			}
		case project.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case project.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = new(string)
				*pr.Description = value.String
			}
		case project.FieldNotes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notes", values[i])
			} else if value.Valid {
				pr.Notes = new(string)
				*pr.Notes = value.String
			}
		case project.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pr.Status = project.Status(value.String)
			}
		case project.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				pr.Location = new(string)
				*pr.Location = value.String
			}
		case project.FieldBudget:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field budget", values[i])
			} else if value.Valid {
				pr.Budget = value.Float64
			}
		case project.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				pr.Deleted = new(bool)
				*pr.Deleted = value.Bool
			}
		case project.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				pr.StartDate = value.Time
			}
		case project.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				pr.EndDate = new(time.Time)
				*pr.EndDate = value.Time
			}
		case project.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case project.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Project.
// This includes values selected through modifiers, order, etc.
func (pr *Project) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// Update returns a builder for updating this Project.
// Note that you need to call Project.Unwrap() before calling this method if this Project
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Project) Update() *ProjectUpdateOne {
	return NewProjectClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Project entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Project) Unwrap() *Project {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Project is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("uuid=")
	builder.WriteString(fmt.Sprintf("%v", pr.UUID))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(pr.AccountID)
	builder.WriteString(", ")
	if v := pr.ClientID; v != nil {
		builder.WriteString("client_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := pr.ManagerID; v != nil {
		builder.WriteString("manager_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(pr.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	if v := pr.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := pr.Notes; v != nil {
		builder.WriteString("notes=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pr.Status))
	builder.WriteString(", ")
	if v := pr.Location; v != nil {
		builder.WriteString("location=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("budget=")
	builder.WriteString(fmt.Sprintf("%v", pr.Budget))
	builder.WriteString(", ")
	if v := pr.Deleted; v != nil {
		builder.WriteString("deleted=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("start_date=")
	builder.WriteString(pr.StartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := pr.EndDate; v != nil {
		builder.WriteString("end_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Projects is a parsable slice of Project.
type Projects []*Project
