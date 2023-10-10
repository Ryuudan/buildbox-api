package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id"),
		field.String("first_name").
			StructTag(`json:"first_name" validate:"required,min=1"`),
		field.String("last_name").StructTag(`json:"last_name" validate:"required,min=1"`),
		field.String("middle_name").
			Optional().
			Nillable().
			StructTag(`json:"middle_name" validate:"omitempty,min=1"`),
		field.Time("birthday").
			StructTag(`json:"birthday" validate:"required"`).
			Nillable().
			Optional(),
		field.String("email").
			Unique().
			StructTag(`json:"email" validate:"required,email"`),
		field.String("phone_number").
			Unique().
			Optional().
			Nillable().
			StructTag(`json:"phone_number" validate:"e164"`),
		field.String("password").
			StructTag(`json:"password" validate:"required,min=3"`),
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

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Field("account_id").
			Ref("users").
			Required().
			Unique(),

		// There are users who can create these following items
		// This is for tracking who created or invited who
		// Example:
		// "Mark Vergel" created a new project
		// "Mark Vergel" created a new task
		// "Mark Vergel" invited a new service provider
		// "Mark Vergel" added a new milestone
		// "Mark Vergel" created an issue
		// "Mark Vergel" assigned "Service Provider 1" to "Project 1"
		// "Mark Vergel" added a new role
		edge.To("tasks", Task.Type),
		edge.To("milestones", Milestone.Type),
		edge.To("issues", Issue.Type),
		edge.To("service_providers", ServiceProvider.Type),
		edge.To("projects", Project.Type),
		// edge.To("project_service_providers", ProjectServiceProvider.Type),
		edge.To("roles", Role.Type),
	}
}
