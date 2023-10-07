// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "phone_number", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "uuid", Type: field.TypeUUID},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
	}
	// PlansColumns holds the columns for the "plans" table.
	PlansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "price", Type: field.TypeFloat64, Default: 0},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "uuid", Type: field.TypeUUID},
	}
	// PlansTable holds the schema information for the "plans" table.
	PlansTable = &schema.Table{
		Name:       "plans",
		Columns:    PlansColumns,
		PrimaryKey: []*schema.Column{PlansColumns[0]},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "client_id", Type: field.TypeInt, Nullable: true},
		{Name: "manager_id", Type: field.TypeInt, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"planning", "in_progress", "on_hold", "completed", "cancelled", "delayed", "under_review", "pending_approval", "in_testing", "emergency", "on_schedule", "behind_schedule", "in_review", "archived", "in_negotiation"}, Default: "planning"},
		{Name: "location", Type: field.TypeString, Nullable: true},
		{Name: "budget", Type: field.TypeFloat64, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "notes", Type: field.TypeString, Nullable: true},
		{Name: "start_date", Type: field.TypeTime, Nullable: true},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "deleted", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "account_id", Type: field.TypeInt},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "projects_accounts_projects",
				Columns:    []*schema.Column{ProjectsColumns[16]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "account_id", Type: field.TypeInt},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "roles_accounts_roles",
				Columns:    []*schema.Column{RolesColumns[6]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "canceled", "expired"}, Default: "active"},
		{Name: "billing_cycle", Type: field.TypeEnum, Enums: []string{"monthly", "yearly"}, Default: "monthly"},
		{Name: "discount", Type: field.TypeFloat64, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "account_id", Type: field.TypeInt},
		{Name: "plan_id", Type: field.TypeInt},
	}
	// SubscriptionsTable holds the schema information for the "subscriptions" table.
	SubscriptionsTable = &schema.Table{
		Name:       "subscriptions",
		Columns:    SubscriptionsColumns,
		PrimaryKey: []*schema.Column{SubscriptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subscriptions_accounts_subscriptions",
				Columns:    []*schema.Column{SubscriptionsColumns[9]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "subscriptions_plans_subscriptions",
				Columns:    []*schema.Column{SubscriptionsColumns[10]},
				RefColumns: []*schema.Column{PlansColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "deleted", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "account_id", Type: field.TypeInt},
		{Name: "project_id", Type: field.TypeInt},
		{Name: "created_by", Type: field.TypeInt},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tasks_accounts_tasks",
				Columns:    []*schema.Column{TasksColumns[7]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "tasks_projects_tasks",
				Columns:    []*schema.Column{TasksColumns[8]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "tasks_users_tasks",
				Columns:    []*schema.Column{TasksColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "middle_name", Type: field.TypeString, Nullable: true},
		{Name: "birthday", Type: field.TypeTime, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "account_id", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_accounts_users",
				Columns:    []*schema.Column{UsersColumns[10]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		PlansTable,
		ProjectsTable,
		RolesTable,
		SubscriptionsTable,
		TasksTable,
		UsersTable,
	}
)

func init() {
	ProjectsTable.ForeignKeys[0].RefTable = AccountsTable
	RolesTable.ForeignKeys[0].RefTable = AccountsTable
	SubscriptionsTable.ForeignKeys[0].RefTable = AccountsTable
	SubscriptionsTable.ForeignKeys[1].RefTable = PlansTable
	TasksTable.ForeignKeys[0].RefTable = AccountsTable
	TasksTable.ForeignKeys[1].RefTable = ProjectsTable
	TasksTable.ForeignKeys[2].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = AccountsTable
}
