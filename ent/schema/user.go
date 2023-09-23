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
		field.UUID("uuid", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.String("firstName"),
		field.String("middleName"),
		field.String("lastName"),
		field.Int("age"),

		field.String("email").
			Optional(),
		field.String("password").
			Optional(),
		field.Time("updated_at").
			Optional().
			Default(time.Now),
		field.Time("created_at").
			Immutable().
			Optional().
			Default(time.Now),
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
	}
}
