package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the ServiceProvider.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id"),
		field.Int("created_by"),
		field.String("name").
			NotEmpty().
			StructTag(`json:"name" validate:"required,min=1"`).
			Optional().
			Nillable(),
		field.String("description").
			StructTag(`json:"description" validate:"omitempty,min=1,max=300"`).
			Optional().
			Nillable(),
		field.Enum("status").
			Default("active").
			Values(
				"active",
				"inactive",
			).
			StructTag(`json:"status" validate:"omitempty,oneof=active inactive"`).
			Optional().
			Nillable(),
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

// Edges of the ServiceProvider.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Field("account_id").
			Ref("teams").
			Required().
			Unique(),
	}
}
