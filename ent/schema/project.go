package schema

import (
	"errors"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		// Accont ID is not optional
		// this is for now, we will add authentication later
		field.Int("account_id"),
		field.UUID("uuid", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.String("client_id").
			Optional().
			Nillable(),
		field.String("manager_id").
			Optional().
			Nillable(),
		field.String("created_by").
			Optional(),
		field.String("name").
			NotEmpty().
			StructTag(`json:"name,omitempty" validate:"required,min=1,max=100"`),
		field.String("description").
			Optional().
			Nillable().
			StructTag(`json:"description,omitempty" validate:"omitempty,min=1,max=200"`),
		field.String("notes").
			Optional().
			Nillable().
			StructTag(`json:"notes,omitempty" validate:"omitempty,min=1,max=200"`),
		field.Enum("status").
			Optional().
			Values(
				"planning",
				"in_progress",
				"on_hold",
				"completed",
				"cancelled",
				"delayed",
				"under_review",
				"pending_approval",
				"in_testing",
				"emergency",
				"on_schedule",
				"behind_schedule",
				"in_review",
				"archived",
				"in_negotiation",
			).
			Default("planning").
			Nillable().
			StructTag(`json:"status,omitempty" validate:"omitempty,oneof=planning in_progress on_hold completed cancelled delayed under_review pending_approval in_testing emergency on_schedule behind_schedule in_review archived in_negotiation"`),
		field.String("location").
			Optional().
			Nillable(),
		field.Float("budget").
			Optional().
			Min(0).
			Positive().
			StructTag(`json:"budget,omitempty" validate:"omitempty,gte=0"`).
			Validate(func(f float64) error {
				if f < 0 {
					return errors.New("budget must be a non-negative value")
				}
				return nil
			}),
		field.Bool("deleted").
			Optional().
			Nillable().
			Default(false),
		field.Time("start_date").
			Optional().
			Default(time.Now).
			StructTag(`json:"start_date,omitempty" validate:"omitempty,ltefield=EndDate"`),
		field.Time("end_date").
			Optional().
			Nillable().
			StructTag(`json:"end_date,omitempty"`),
		field.Time("updated_at").
			Optional().
			Default(time.Now),
		field.Time("created_at").
			Immutable().
			Optional().
			Default(time.Now),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Field("account_id").
			Ref("projects").
			Required().
			Unique(),
	}
}
