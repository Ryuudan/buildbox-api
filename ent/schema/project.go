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

func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id"),
		field.Int("created_by"),
		field.Int("client_id").
			StructTag(`json:"client_id"`).
			Optional().
			Nillable(),
		field.Int("manager_id").
			StructTag(`json:"manager_id"`).
			Optional().
			Nillable(),
		field.String("name").
			StructTag(`json:"name" validate:"required,min=1,max=100"`).
			NotEmpty(),
		field.Enum("status").
			Default("planning").
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
			StructTag(`json:"status" validate:"omitempty,oneof=planning in_progress on_hold completed cancelled delayed under_review pending_approval in_testing emergency on_schedule behind_schedule in_review archived in_negotiation"`).
			Optional().
			Nillable(),
		field.String("location").
			StructTag(`json:"location"`).
			Optional().
			Nillable(),
		field.Float("budget").
			StructTag(`json:"budget" validate:"omitempty,gte=0"`).
			Optional().
			Min(0).
			Nillable(),
		field.String("description").
			StructTag(`json:"description" validate:"omitempty,min=1,max=300"`).
			Optional().
			Nillable(),
		field.String("notes").
			StructTag(`json:"notes" validate:"omitempty,min=1,max=200"`).
			Optional().
			Nillable(),
		field.Time("start_date").
			Default(time.Now).
			StructTag(`json:"start_date" validate:"omitempty,ltfield=EndDate"`).
			Optional().
			Nillable(),
		field.Time("end_date").
			StructTag(`json:"end_date"`).
			Optional().
			Nillable(),
		field.UUID("uuid", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.Bool("deleted").
			StructTag(`json:"deleted"`).
			Optional().
			Default(false),
		field.Time("updated_at").
			Default(time.Now),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
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
