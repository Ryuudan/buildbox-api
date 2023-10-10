package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ProjectServiceProvider holds the schema definition for the ProjectServiceProvider entity.
type ProjectServiceProvider struct {
	ent.Schema
}

// Fields of the ProjectServiceProvider.
func (ProjectServiceProvider) Fields() []ent.Field {
	return []ent.Field{
		field.Int("created_by"),
		field.Int("project_service_provider_id").
			// StructTag(`json:"project_service_provider_id"`).
			Optional().
			StorageKey("service_provider_id"),
		field.Int("project_project_id").
			// StructTag(`json:"project_project_id"`).
			Optional().StorageKey("project_id"),
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

// Edges of the ProjectServiceProvider.
func (ProjectServiceProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Field("project_project_id").
			Ref("project_service_providers").
			Unique(),
		edge.From("service_provider", ServiceProvider.Type).
			Field("project_service_provider_id").
			Ref("service_provider_projects").
			Unique(),
	}
}
