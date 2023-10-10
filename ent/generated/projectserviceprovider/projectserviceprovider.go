// Code generated by ent, DO NOT EDIT.

package projectserviceprovider

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the projectserviceprovider type in the database.
	Label = "project_service_provider"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldProjectServiceProviderID holds the string denoting the project_service_provider_id field in the database.
	FieldProjectServiceProviderID = "service_provider_id"
	// FieldProjectProjectID holds the string denoting the project_project_id field in the database.
	FieldProjectProjectID = "project_id"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeServiceProvider holds the string denoting the service_provider edge name in mutations.
	EdgeServiceProvider = "service_provider"
	// Table holds the table name of the projectserviceprovider in the database.
	Table = "project_service_providers"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "project_service_providers"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_id"
	// ServiceProviderTable is the table that holds the service_provider relation/edge.
	ServiceProviderTable = "project_service_providers"
	// ServiceProviderInverseTable is the table name for the ServiceProvider entity.
	// It exists in this package in order to avoid circular dependency with the "serviceprovider" package.
	ServiceProviderInverseTable = "service_providers"
	// ServiceProviderColumn is the table column denoting the service_provider relation/edge.
	ServiceProviderColumn = "service_provider_id"
)

// Columns holds all SQL columns for projectserviceprovider fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldProjectServiceProviderID,
	FieldProjectProjectID,
	FieldUpdatedAt,
	FieldCreatedAt,
	FieldUUID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
)

// OrderOption defines the ordering options for the ProjectServiceProvider queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByProjectServiceProviderID orders the results by the project_service_provider_id field.
func ByProjectServiceProviderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectServiceProviderID, opts...).ToFunc()
}

// ByProjectProjectID orders the results by the project_project_id field.
func ByProjectProjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectProjectID, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}

// ByServiceProviderField orders the results by service_provider field.
func ByServiceProviderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServiceProviderStep(), sql.OrderByField(field, opts...))
	}
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
	)
}
func newServiceProviderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServiceProviderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ServiceProviderTable, ServiceProviderColumn),
	)
}
