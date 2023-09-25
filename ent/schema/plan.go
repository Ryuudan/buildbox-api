package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Plan holds the schema definition for the Plan entity.
type Plan struct {
	ent.Schema
}

// Fields of the Plan.
func (Plan) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			StructTag(`json:"name" validate:"required,min=3"`).
			Unique(),
		field.Text("description").
			StructTag(`json:"description" validate:"required,min=10"`).
			Optional(),
		field.Float("price").
			Nillable().
			StructTag(`json:"price,omitempty" validate:"omitempty,gte=0"`).
			Default(0),
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

// Edges of the Plan.
func (Plan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subscriptions", Subscription.Type),
	}
}
