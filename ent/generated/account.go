// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Pyakz/buildbox-api/ent/generated/account"
	"github.com/google/uuid"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// The name of the account, or company.
	Name string `json:"name" validate:"required,min=1"`
	// Email holds the value of the "email" field.
	Email string `json:"email" validate:"required,email"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges        AccountEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// Subscriptions holds the value of the subscriptions edge.
	Subscriptions []*Subscription `json:"subscriptions,omitempty"`
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// Tasks holds the value of the tasks edge.
	Tasks []*Task `json:"tasks,omitempty"`
	// Milestones holds the value of the milestones edge.
	Milestones []*Milestone `json:"milestones,omitempty"`
	// Issues holds the value of the issues edge.
	Issues []*Issue `json:"issues,omitempty"`
	// ServiceProviders holds the value of the service_providers edge.
	ServiceProviders []*ServiceProvider `json:"service_providers,omitempty"`
	// Teams holds the value of the teams edge.
	Teams []*Team `json:"teams,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[1] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// SubscriptionsOrErr returns the Subscriptions value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) SubscriptionsOrErr() ([]*Subscription, error) {
	if e.loadedTypes[2] {
		return e.Subscriptions, nil
	}
	return nil, &NotLoadedError{edge: "subscriptions"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[3] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[4] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// MilestonesOrErr returns the Milestones value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) MilestonesOrErr() ([]*Milestone, error) {
	if e.loadedTypes[5] {
		return e.Milestones, nil
	}
	return nil, &NotLoadedError{edge: "milestones"}
}

// IssuesOrErr returns the Issues value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) IssuesOrErr() ([]*Issue, error) {
	if e.loadedTypes[6] {
		return e.Issues, nil
	}
	return nil, &NotLoadedError{edge: "issues"}
}

// ServiceProvidersOrErr returns the ServiceProviders value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) ServiceProvidersOrErr() ([]*ServiceProvider, error) {
	if e.loadedTypes[7] {
		return e.ServiceProviders, nil
	}
	return nil, &NotLoadedError{edge: "service_providers"}
}

// TeamsOrErr returns the Teams value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) TeamsOrErr() ([]*Team, error) {
	if e.loadedTypes[8] {
		return e.Teams, nil
	}
	return nil, &NotLoadedError{edge: "teams"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			values[i] = new(sql.NullInt64)
		case account.FieldName, account.FieldEmail, account.FieldPhoneNumber:
			values[i] = new(sql.NullString)
		case account.FieldUpdatedAt, account.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case account.FieldUUID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case account.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case account.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				a.Email = value.String
			}
		case account.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				a.PhoneNumber = value.String
			}
		case account.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case account.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case account.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				a.UUID = *value
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Account.
// This includes values selected through modifiers, order, etc.
func (a *Account) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Account entity.
func (a *Account) QueryUsers() *UserQuery {
	return NewAccountClient(a.config).QueryUsers(a)
}

// QueryProjects queries the "projects" edge of the Account entity.
func (a *Account) QueryProjects() *ProjectQuery {
	return NewAccountClient(a.config).QueryProjects(a)
}

// QuerySubscriptions queries the "subscriptions" edge of the Account entity.
func (a *Account) QuerySubscriptions() *SubscriptionQuery {
	return NewAccountClient(a.config).QuerySubscriptions(a)
}

// QueryRoles queries the "roles" edge of the Account entity.
func (a *Account) QueryRoles() *RoleQuery {
	return NewAccountClient(a.config).QueryRoles(a)
}

// QueryTasks queries the "tasks" edge of the Account entity.
func (a *Account) QueryTasks() *TaskQuery {
	return NewAccountClient(a.config).QueryTasks(a)
}

// QueryMilestones queries the "milestones" edge of the Account entity.
func (a *Account) QueryMilestones() *MilestoneQuery {
	return NewAccountClient(a.config).QueryMilestones(a)
}

// QueryIssues queries the "issues" edge of the Account entity.
func (a *Account) QueryIssues() *IssueQuery {
	return NewAccountClient(a.config).QueryIssues(a)
}

// QueryServiceProviders queries the "service_providers" edge of the Account entity.
func (a *Account) QueryServiceProviders() *ServiceProviderQuery {
	return NewAccountClient(a.config).QueryServiceProviders(a)
}

// QueryTeams queries the "teams" edge of the Account entity.
func (a *Account) QueryTeams() *TeamQuery {
	return NewAccountClient(a.config).QueryTeams(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return NewAccountClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("generated: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(a.Email)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(a.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("uuid=")
	builder.WriteString(fmt.Sprintf("%v", a.UUID))
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account
