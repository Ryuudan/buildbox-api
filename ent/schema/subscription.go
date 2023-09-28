package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id").
			StructTag(`json:"account_id"`),
		field.Int("plan_id").
			StructTag(`json:"plan_id"`),
		field.Time("start_date").
			Default(time.Now),
		field.Time("end_date").
			Optional().
			StructTag(`json:"end_date,omitempty"`),
		field.Enum("status").
			Default("active").
			Values("active", "canceled", "expired").
			StructTag(`json:"status" validate:"omitempty,oneof=active canceled expired"`),
		field.Enum("billing_cycle").
			Default("monthly").
			Values("monthly", "yearly").
			StructTag(`json:"billing_cycle" validate:"omitempty,oneof=monthly yearly"`),
		field.Float("discount").
			Optional().
			Comment("Discount applied to the subscription, this percent.").
			StructTag(`json:"discount" validate:"gte=0"`),
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

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Field("account_id").
			Ref("subscriptions").
			Required().
			Unique(),
		edge.From("plan", Plan.Type).
			Field("plan_id").
			Ref("subscriptions").
			Required().
			Unique(),
	}
}
