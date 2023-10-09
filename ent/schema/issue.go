package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Issue holds the schema definition for the Issue entity.
type Issue struct {
	ent.Schema
}

// Fields of the Issue.
func (Issue) Fields() []ent.Field {
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

// Edges of the Issue.
func (Issue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tasks", Task.Type),
		edge.From("account", Account.Type).
			Field("account_id").
			Ref("issues").
			Required().
			Unique(),
		edge.From("user", User.Type).
			Field("created_by").
			Ref("issues").
			Required().
			Unique(),
		edge.From("project", Project.Type).
			Field("project_id").
			Ref("issues").
			Required().
			Unique(),
	}
}
