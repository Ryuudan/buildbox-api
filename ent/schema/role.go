package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id"),
		field.Int("created_by"),
		field.String("name").
			StructTag(`json:"name" validate:"required,min=1,max=100"`).Optional().Nillable(),
		field.String("description").
			StructTag(`json:"description" validate:"required,min=1,max=100"`).Optional().Nillable(),
		field.Time("updated_at").
			Default(time.Now),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.UUID("uuid", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Field("account_id").
			Ref("roles").
			Required().
			Unique(),
		edge.From("user", User.Type).
			Field("created_by").
			Ref("roles").
			Required().
			Unique(),
	}
}
