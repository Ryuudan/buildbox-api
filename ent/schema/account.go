package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("The name of the account, or company.").
			StructTag(`json:"name" validate:"required,min=1"`),
		field.String("email").
			Optional().
			Unique().
			NotEmpty().
			StructTag(`json:"email" validate:"required,email"`),
		field.String("phone_number").
			Optional().
			NotEmpty().
			StructTag(`json:"phone_number" validate:"required,phone"`),
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

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("projects", Project.Type),
		edge.To("subscriptions", Subscription.Type),
	}
}
