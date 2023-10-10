// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldMiddleName holds the string denoting the middle_name field in the database.
	FieldMiddleName = "middle_name"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// EdgeTasks holds the string denoting the tasks edge name in mutations.
	EdgeTasks = "tasks"
	// EdgeMilestones holds the string denoting the milestones edge name in mutations.
	EdgeMilestones = "milestones"
	// EdgeIssues holds the string denoting the issues edge name in mutations.
	EdgeIssues = "issues"
	// EdgeServiceProviders holds the string denoting the service_providers edge name in mutations.
	EdgeServiceProviders = "service_providers"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "users"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_id"
	// TasksTable is the table that holds the tasks relation/edge.
	TasksTable = "tasks"
	// TasksInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TasksInverseTable = "tasks"
	// TasksColumn is the table column denoting the tasks relation/edge.
	TasksColumn = "created_by"
	// MilestonesTable is the table that holds the milestones relation/edge.
	MilestonesTable = "milestones"
	// MilestonesInverseTable is the table name for the Milestone entity.
	// It exists in this package in order to avoid circular dependency with the "milestone" package.
	MilestonesInverseTable = "milestones"
	// MilestonesColumn is the table column denoting the milestones relation/edge.
	MilestonesColumn = "created_by"
	// IssuesTable is the table that holds the issues relation/edge.
	IssuesTable = "issues"
	// IssuesInverseTable is the table name for the Issue entity.
	// It exists in this package in order to avoid circular dependency with the "issue" package.
	IssuesInverseTable = "issues"
	// IssuesColumn is the table column denoting the issues relation/edge.
	IssuesColumn = "created_by"
	// ServiceProvidersTable is the table that holds the service_providers relation/edge.
	ServiceProvidersTable = "service_providers"
	// ServiceProvidersInverseTable is the table name for the ServiceProvider entity.
	// It exists in this package in order to avoid circular dependency with the "serviceprovider" package.
	ServiceProvidersInverseTable = "service_providers"
	// ServiceProvidersColumn is the table column denoting the service_providers relation/edge.
	ServiceProvidersColumn = "created_by"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAccountID,
	FieldFirstName,
	FieldLastName,
	FieldMiddleName,
	FieldBirthday,
	FieldEmail,
	FieldPhoneNumber,
	FieldPassword,
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

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAccountID orders the results by the account_id field.
func ByAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountID, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByMiddleName orders the results by the middle_name field.
func ByMiddleName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMiddleName, opts...).ToFunc()
}

// ByBirthday orders the results by the birthday field.
func ByBirthday(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthday, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhoneNumber orders the results by the phone_number field.
func ByPhoneNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumber, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
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

// ByAccountField orders the results by account field.
func ByAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountStep(), sql.OrderByField(field, opts...))
	}
}

// ByTasksCount orders the results by tasks count.
func ByTasksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTasksStep(), opts...)
	}
}

// ByTasks orders the results by tasks terms.
func ByTasks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTasksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMilestonesCount orders the results by milestones count.
func ByMilestonesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMilestonesStep(), opts...)
	}
}

// ByMilestones orders the results by milestones terms.
func ByMilestones(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMilestonesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByIssuesCount orders the results by issues count.
func ByIssuesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIssuesStep(), opts...)
	}
}

// ByIssues orders the results by issues terms.
func ByIssues(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIssuesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByServiceProvidersCount orders the results by service_providers count.
func ByServiceProvidersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newServiceProvidersStep(), opts...)
	}
}

// ByServiceProviders orders the results by service_providers terms.
func ByServiceProviders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServiceProvidersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
	)
}
func newTasksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TasksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TasksTable, TasksColumn),
	)
}
func newMilestonesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MilestonesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MilestonesTable, MilestonesColumn),
	)
}
func newIssuesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IssuesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IssuesTable, IssuesColumn),
	)
}
func newServiceProvidersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServiceProvidersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ServiceProvidersTable, ServiceProvidersColumn),
	)
}
