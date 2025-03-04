// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Pyakz/buildbox-api/ent/generated/account"
	"github.com/Pyakz/buildbox-api/ent/generated/project"
	"github.com/Pyakz/buildbox-api/ent/generated/user"
	"github.com/google/uuid"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID int `json:"account_id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int `json:"created_by,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID *int `json:"client_id"`
	// ManagerID holds the value of the "manager_id" field.
	ManagerID *int `json:"manager_id"`
	// Name holds the value of the "name" field.
	Name string `json:"name" validate:"required,min=1,max=100"`
	// Status holds the value of the "status" field.
	Status *project.Status `json:"status" validate:"omitempty,oneof=planning in_progress on_hold completed cancelled delayed under_review pending_approval in_testing emergency on_schedule behind_schedule in_review archived in_negotiation"`
	// Location holds the value of the "location" field.
	Location *string `json:"location"`
	// Budget holds the value of the "budget" field.
	Budget *float64 `json:"budget" validate:"omitempty,gte=0"`
	// Description holds the value of the "description" field.
	Description *string `json:"description" validate:"omitempty,min=1,max=300"`
	// Notes holds the value of the "notes" field.
	Notes *string `json:"notes" validate:"omitempty,min=1,max=200"`
	// StartDate holds the value of the "start_date" field.
	StartDate *time.Time `json:"start_date" validate:"omitempty,ltfield=EndDate"`
	// EndDate holds the value of the "end_date" field.
	EndDate *time.Time `json:"end_date"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectQuery when eager-loading is set.
	Edges        ProjectEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProjectEdges holds the relations/edges for other nodes in the graph.
type ProjectEdges struct {
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Tasks holds the value of the tasks edge.
	Tasks []*Task `json:"tasks,omitempty"`
	// Milestones holds the value of the milestones edge.
	Milestones []*Milestone `json:"milestones,omitempty"`
	// Issues holds the value of the issues edge.
	Issues []*Issue `json:"issues,omitempty"`
	// Teams holds the value of the teams edge.
	Teams []*Team `json:"teams,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectEdges) AccountOrErr() (*Account, error) {
	if e.loadedTypes[0] {
		if e.Account == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Account, nil
	}
	return nil, &NotLoadedError{edge: "account"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[2] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// MilestonesOrErr returns the Milestones value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) MilestonesOrErr() ([]*Milestone, error) {
	if e.loadedTypes[3] {
		return e.Milestones, nil
	}
	return nil, &NotLoadedError{edge: "milestones"}
}

// IssuesOrErr returns the Issues value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) IssuesOrErr() ([]*Issue, error) {
	if e.loadedTypes[4] {
		return e.Issues, nil
	}
	return nil, &NotLoadedError{edge: "issues"}
}

// TeamsOrErr returns the Teams value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) TeamsOrErr() ([]*Team, error) {
	if e.loadedTypes[5] {
		return e.Teams, nil
	}
	return nil, &NotLoadedError{edge: "teams"}
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
		case project.FieldID, project.FieldAccountID, project.FieldCreatedBy, project.FieldClientID, project.FieldManagerID:
			values[i] = new(sql.NullInt64)
		case project.FieldName, project.FieldStatus, project.FieldLocation, project.FieldDescription, project.FieldNotes:
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
		case project.FieldAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				pr.AccountID = int(value.Int64)
			}
		case project.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pr.CreatedBy = int(value.Int64)
			}
		case project.FieldClientID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				pr.ClientID = new(int)
				*pr.ClientID = int(value.Int64)
			}
		case project.FieldManagerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field manager_id", values[i])
			} else if value.Valid {
				pr.ManagerID = new(int)
				*pr.ManagerID = int(value.Int64)
			}
		case project.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case project.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pr.Status = new(project.Status)
				*pr.Status = project.Status(value.String)
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
				pr.Budget = new(float64)
				*pr.Budget = value.Float64
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
		case project.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				pr.StartDate = new(time.Time)
				*pr.StartDate = value.Time
			}
		case project.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				pr.EndDate = new(time.Time)
				*pr.EndDate = value.Time
			}
		case project.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				pr.UUID = *value
			}
		case project.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				pr.Deleted = value.Bool
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

// QueryAccount queries the "account" edge of the Project entity.
func (pr *Project) QueryAccount() *AccountQuery {
	return NewProjectClient(pr.config).QueryAccount(pr)
}

// QueryUser queries the "user" edge of the Project entity.
func (pr *Project) QueryUser() *UserQuery {
	return NewProjectClient(pr.config).QueryUser(pr)
}

// QueryTasks queries the "tasks" edge of the Project entity.
func (pr *Project) QueryTasks() *TaskQuery {
	return NewProjectClient(pr.config).QueryTasks(pr)
}

// QueryMilestones queries the "milestones" edge of the Project entity.
func (pr *Project) QueryMilestones() *MilestoneQuery {
	return NewProjectClient(pr.config).QueryMilestones(pr)
}

// QueryIssues queries the "issues" edge of the Project entity.
func (pr *Project) QueryIssues() *IssueQuery {
	return NewProjectClient(pr.config).QueryIssues(pr)
}

// QueryTeams queries the "teams" edge of the Project entity.
func (pr *Project) QueryTeams() *TeamQuery {
	return NewProjectClient(pr.config).QueryTeams(pr)
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
		panic("generated: Project is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.AccountID))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", pr.CreatedBy))
	builder.WriteString(", ")
	if v := pr.ClientID; v != nil {
		builder.WriteString("client_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := pr.ManagerID; v != nil {
		builder.WriteString("manager_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	if v := pr.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := pr.Location; v != nil {
		builder.WriteString("location=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := pr.Budget; v != nil {
		builder.WriteString("budget=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
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
	if v := pr.StartDate; v != nil {
		builder.WriteString("start_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := pr.EndDate; v != nil {
		builder.WriteString("end_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("uuid=")
	builder.WriteString(fmt.Sprintf("%v", pr.UUID))
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", pr.Deleted))
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
