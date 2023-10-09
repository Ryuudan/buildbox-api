// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Pyakz/buildbox-api/ent/generated/plan"
	"github.com/google/uuid"
)

// Plan is the model entity for the Plan schema.
type Plan struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name" validate:"required,min=3"`
	// Description holds the value of the "description" field.
	Description string `json:"description" validate:"required,min=10"`
	// Price holds the value of the "price" field.
	Price *float64 `json:"price" validate:"required,gte=0"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlanQuery when eager-loading is set.
	Edges        PlanEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PlanEdges holds the relations/edges for other nodes in the graph.
type PlanEdges struct {
	// Subscriptions holds the value of the subscriptions edge.
	Subscriptions []*Subscription `json:"subscriptions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SubscriptionsOrErr returns the Subscriptions value or an error if the edge
// was not loaded in eager-loading.
func (e PlanEdges) SubscriptionsOrErr() ([]*Subscription, error) {
	if e.loadedTypes[0] {
		return e.Subscriptions, nil
	}
	return nil, &NotLoadedError{edge: "subscriptions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Plan) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case plan.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case plan.FieldID:
			values[i] = new(sql.NullInt64)
		case plan.FieldName, plan.FieldDescription:
			values[i] = new(sql.NullString)
		case plan.FieldUpdatedAt, plan.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case plan.FieldUUID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Plan fields.
func (pl *Plan) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case plan.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pl.ID = int(value.Int64)
		case plan.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case plan.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pl.Description = value.String
			}
		case plan.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				pl.Price = new(float64)
				*pl.Price = value.Float64
			}
		case plan.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pl.UpdatedAt = value.Time
			}
		case plan.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pl.CreatedAt = value.Time
			}
		case plan.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				pl.UUID = *value
			}
		default:
			pl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Plan.
// This includes values selected through modifiers, order, etc.
func (pl *Plan) Value(name string) (ent.Value, error) {
	return pl.selectValues.Get(name)
}

// QuerySubscriptions queries the "subscriptions" edge of the Plan entity.
func (pl *Plan) QuerySubscriptions() *SubscriptionQuery {
	return NewPlanClient(pl.config).QuerySubscriptions(pl)
}

// Update returns a builder for updating this Plan.
// Note that you need to call Plan.Unwrap() before calling this method if this Plan
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Plan) Update() *PlanUpdateOne {
	return NewPlanClient(pl.config).UpdateOne(pl)
}

// Unwrap unwraps the Plan entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Plan) Unwrap() *Plan {
	_tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("generated: Plan is not a transactional entity")
	}
	pl.config.driver = _tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Plan) String() string {
	var builder strings.Builder
	builder.WriteString("Plan(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pl.ID))
	builder.WriteString("name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pl.Description)
	builder.WriteString(", ")
	if v := pl.Price; v != nil {
		builder.WriteString("price=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pl.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("uuid=")
	builder.WriteString(fmt.Sprintf("%v", pl.UUID))
	builder.WriteByte(')')
	return builder.String()
}

// Plans is a parsable slice of Plan.
type Plans []*Plan
