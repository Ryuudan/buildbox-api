package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ServiceProvider holds the schema definition for the ServiceProvider entity.
type ServiceProvider struct {
	ent.Schema
}

// Fields of the ServiceProvider.
func (ServiceProvider) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id"),
		field.Int("created_by"),
		field.String("name").
			NotEmpty().
			StructTag(`json:"name" validate:"required,min=1"`),
		field.String("email").
			Unique().
			StructTag(`json:"email" validate:"required,email"`),
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
		field.String("phone_number").
			Unique().
			StructTag(`json:"phone_number" validate:"required,e164"`),
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
func (ServiceProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("service_provider_projects", ProjectServiceProvider.Type),
		edge.From("account", Account.Type).
			Field("account_id").
			Ref("service_providers").
			Required().
			Unique(),
		edge.From("user", User.Type).
			Field("created_by").
			Ref("service_providers").
			Required().
			Unique(),
	}
}
