package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
		field.String("account_id").
			Optional(),
		field.String("client_id").
			Optional().
			Nillable(),
		field.String("manager_id").
			Optional().
			Nillable(),
		field.String("created_by").
			Optional().
			Nillable(),
		field.String("name").
			NotEmpty(),
		field.String("description").
			Optional().
			Nillable(),
		field.String("notes").
			Optional().
			Nillable(),
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
			).Default("planning"),
		field.String("location").
			Optional().
			Nillable(),
		field.Float("budget").
			Optional().
			Nillable(),
		field.Bool("deleted").
			Optional().
			Nillable().
			Default(false),
		field.Time("start_date").
			Optional().
			Default(time.Now),
		field.Time("end_date").
			Optional().
			Nillable(),
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
	return nil
}
