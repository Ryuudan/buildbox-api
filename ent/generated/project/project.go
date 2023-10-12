// Code generated by ent, DO NOT EDIT.

package project

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldClientID holds the string denoting the client_id field in the database.
	FieldClientID = "client_id"
	// FieldManagerID holds the string denoting the manager_id field in the database.
	FieldManagerID = "manager_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldBudget holds the string denoting the budget field in the database.
	FieldBudget = "budget"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldNotes holds the string denoting the notes field in the database.
	FieldNotes = "notes"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldDeleted holds the string denoting the deleted field in the database.
	FieldDeleted = "deleted"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeTasks holds the string denoting the tasks edge name in mutations.
	EdgeTasks = "tasks"
	// EdgeMilestones holds the string denoting the milestones edge name in mutations.
	EdgeMilestones = "milestones"
	// EdgeIssues holds the string denoting the issues edge name in mutations.
	EdgeIssues = "issues"
	// EdgeTeams holds the string denoting the teams edge name in mutations.
	EdgeTeams = "teams"
	// Table holds the table name of the project in the database.
	Table = "projects"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "projects"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "projects"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "created_by"
	// TasksTable is the table that holds the tasks relation/edge.
	TasksTable = "tasks"
	// TasksInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TasksInverseTable = "tasks"
	// TasksColumn is the table column denoting the tasks relation/edge.
	TasksColumn = "project_id"
	// MilestonesTable is the table that holds the milestones relation/edge.
	MilestonesTable = "milestones"
	// MilestonesInverseTable is the table name for the Milestone entity.
	// It exists in this package in order to avoid circular dependency with the "milestone" package.
	MilestonesInverseTable = "milestones"
	// MilestonesColumn is the table column denoting the milestones relation/edge.
	MilestonesColumn = "project_id"
	// IssuesTable is the table that holds the issues relation/edge.
	IssuesTable = "issues"
	// IssuesInverseTable is the table name for the Issue entity.
	// It exists in this package in order to avoid circular dependency with the "issue" package.
	IssuesInverseTable = "issues"
	// IssuesColumn is the table column denoting the issues relation/edge.
	IssuesColumn = "project_id"
	// TeamsTable is the table that holds the teams relation/edge.
	TeamsTable = "teams"
	// TeamsInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamsInverseTable = "teams"
	// TeamsColumn is the table column denoting the teams relation/edge.
	TeamsColumn = "project_teams"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldAccountID,
	FieldCreatedBy,
	FieldClientID,
	FieldManagerID,
	FieldName,
	FieldStatus,
	FieldLocation,
	FieldBudget,
	FieldDescription,
	FieldNotes,
	FieldStartDate,
	FieldEndDate,
	FieldUUID,
	FieldDeleted,
	FieldUpdatedAt,
	FieldCreatedAt,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// BudgetValidator is a validator for the "budget" field. It is called by the builders before save.
	BudgetValidator func(float64) error
	// DefaultStartDate holds the default value on creation for the "start_date" field.
	DefaultStartDate func() time.Time
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
	// DefaultDeleted holds the default value on creation for the "deleted" field.
	DefaultDeleted bool
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// StatusPlanning is the default value of the Status enum.
const DefaultStatus = StatusPlanning

// Status values.
const (
	StatusPlanning        Status = "planning"
	StatusInProgress      Status = "in_progress"
	StatusOnHold          Status = "on_hold"
	StatusCompleted       Status = "completed"
	StatusCancelled       Status = "cancelled"
	StatusDelayed         Status = "delayed"
	StatusUnderReview     Status = "under_review"
	StatusPendingApproval Status = "pending_approval"
	StatusInTesting       Status = "in_testing"
	StatusEmergency       Status = "emergency"
	StatusOnSchedule      Status = "on_schedule"
	StatusBehindSchedule  Status = "behind_schedule"
	StatusInReview        Status = "in_review"
	StatusArchived        Status = "archived"
	StatusInNegotiation   Status = "in_negotiation"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPlanning, StatusInProgress, StatusOnHold, StatusCompleted, StatusCancelled, StatusDelayed, StatusUnderReview, StatusPendingApproval, StatusInTesting, StatusEmergency, StatusOnSchedule, StatusBehindSchedule, StatusInReview, StatusArchived, StatusInNegotiation:
		return nil
	default:
		return fmt.Errorf("project: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Project queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAccountID orders the results by the account_id field.
func ByAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountID, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByClientID orders the results by the client_id field.
func ByClientID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientID, opts...).ToFunc()
}

// ByManagerID orders the results by the manager_id field.
func ByManagerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldManagerID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByBudget orders the results by the budget field.
func ByBudget(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBudget, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByNotes orders the results by the notes field.
func ByNotes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotes, opts...).ToFunc()
}

// ByStartDate orders the results by the start_date field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the end_date field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByDeleted orders the results by the deleted field.
func ByDeleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleted, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByAccountField orders the results by account field.
func ByAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
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

// ByTeamsCount orders the results by teams count.
func ByTeamsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTeamsStep(), opts...)
	}
}

// ByTeams orders the results by teams terms.
func ByTeams(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
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
func newTeamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TeamsTable, TeamsColumn),
	)
}
