package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// TODO: add status and complete this schema
// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id"),
		field.Int("created_by"),
		field.Int("project_id").StructTag(`validate:"required"`),
		field.String("title").
			StructTag(`json:"title" validate:"required,min=1,max=100"`).
			NotEmpty(),
		field.String("description").
			StructTag(`json:"description" validate:"required,min=1,max=100"`).
			NotEmpty(),
		field.Time("updated_at").
			Default(time.Now),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Bool("deleted").
			StructTag(`json:"deleted"`).
			Optional().
			Default(false),
		field.UUID("uuid", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Field("account_id").
			Ref("tasks").
			Required().
			Unique(),
		edge.From("user", User.Type).
			Field("created_by").
			Ref("tasks").
			Required().
			Unique(),
		edge.From("project", Project.Type).
			Field("project_id").
			Ref("tasks").
			Required().
			Unique(),
	}
}
