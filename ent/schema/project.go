package schema

import (
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
		field.Int("account_id"),
		field.Int("created_by"),
		field.Int("client_id").
			Optional().
			StructTag(`json:"client_id"`).
			Nillable(),
		field.Int("manager_id").
			Optional().
			StructTag(`json:"manager_id"`).
			Nillable(),
		field.String("name").
			NotEmpty().
			StructTag(`json:"name" validate:"required,min=1,max=100"`),
		field.String("description").
			Optional().
			Nillable().
			StructTag(`json:"description" validate:"omitempty,min=1,max=300"`),
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
			Nillable(),
		field.String("location").
			Optional().
			StructTag(`json:"location"`).
			Nillable(),
		field.Float("budget").
			Optional().
			Min(0).
			StructTag(`json:"budget" validate:"omitempty,gte=0"`).
			Nillable(),
		field.Bool("deleted").
			Optional().
			StructTag(`json:"deleted"`).
			Default(false),
		field.Time("start_date").
			Optional().
			Nillable().
			StructTag(`json:"start_date" validate:"omitempty,ltfield=EndDate"`).
			Default(time.Now),
		field.Time("end_date").
			Optional().
			Nillable().
			StructTag(`json:"end_date"`),
		field.Time("updated_at").
			Optional().
			Default(time.Now),
		field.Time("created_at").
			Immutable().
			Optional().
			Default(time.Now),
		field.UUID("uuid", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
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
