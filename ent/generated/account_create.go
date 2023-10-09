// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Pyakz/buildbox-api/ent/generated/account"
	"github.com/Pyakz/buildbox-api/ent/generated/issue"
	"github.com/Pyakz/buildbox-api/ent/generated/milestone"
	"github.com/Pyakz/buildbox-api/ent/generated/project"
	"github.com/Pyakz/buildbox-api/ent/generated/role"
	"github.com/Pyakz/buildbox-api/ent/generated/subscription"
	"github.com/Pyakz/buildbox-api/ent/generated/task"
	"github.com/Pyakz/buildbox-api/ent/generated/user"
	"github.com/google/uuid"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AccountCreate) SetName(s string) *AccountCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetEmail sets the "email" field.
func (ac *AccountCreate) SetEmail(s string) *AccountCreate {
	ac.mutation.SetEmail(s)
	return ac
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (ac *AccountCreate) SetNillableEmail(s *string) *AccountCreate {
	if s != nil {
		ac.SetEmail(*s)
	}
	return ac
}

// SetPhoneNumber sets the "phone_number" field.
func (ac *AccountCreate) SetPhoneNumber(s string) *AccountCreate {
	ac.mutation.SetPhoneNumber(s)
	return ac
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (ac *AccountCreate) SetNillablePhoneNumber(s *string) *AccountCreate {
	if s != nil {
		ac.SetPhoneNumber(*s)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AccountCreate) SetUpdatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUpdatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AccountCreate) SetCreatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableCreatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUUID sets the "uuid" field.
func (ac *AccountCreate) SetUUID(u uuid.UUID) *AccountCreate {
	ac.mutation.SetUUID(u)
	return ac
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUUID(u *uuid.UUID) *AccountCreate {
	if u != nil {
		ac.SetUUID(*u)
	}
	return ac
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (ac *AccountCreate) AddUserIDs(ids ...int) *AccountCreate {
	ac.mutation.AddUserIDs(ids...)
	return ac
}

// AddUsers adds the "users" edges to the User entity.
func (ac *AccountCreate) AddUsers(u ...*User) *AccountCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ac.AddUserIDs(ids...)
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (ac *AccountCreate) AddProjectIDs(ids ...int) *AccountCreate {
	ac.mutation.AddProjectIDs(ids...)
	return ac
}

// AddProjects adds the "projects" edges to the Project entity.
func (ac *AccountCreate) AddProjects(p ...*Project) *AccountCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddProjectIDs(ids...)
}

// AddSubscriptionIDs adds the "subscriptions" edge to the Subscription entity by IDs.
func (ac *AccountCreate) AddSubscriptionIDs(ids ...int) *AccountCreate {
	ac.mutation.AddSubscriptionIDs(ids...)
	return ac
}

// AddSubscriptions adds the "subscriptions" edges to the Subscription entity.
func (ac *AccountCreate) AddSubscriptions(s ...*Subscription) *AccountCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ac.AddSubscriptionIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (ac *AccountCreate) AddRoleIDs(ids ...int) *AccountCreate {
	ac.mutation.AddRoleIDs(ids...)
	return ac
}

// AddRoles adds the "roles" edges to the Role entity.
func (ac *AccountCreate) AddRoles(r ...*Role) *AccountCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ac.AddRoleIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (ac *AccountCreate) AddTaskIDs(ids ...int) *AccountCreate {
	ac.mutation.AddTaskIDs(ids...)
	return ac
}

// AddTasks adds the "tasks" edges to the Task entity.
func (ac *AccountCreate) AddTasks(t ...*Task) *AccountCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTaskIDs(ids...)
}

// AddMilestoneIDs adds the "milestones" edge to the Milestone entity by IDs.
func (ac *AccountCreate) AddMilestoneIDs(ids ...int) *AccountCreate {
	ac.mutation.AddMilestoneIDs(ids...)
	return ac
}

// AddMilestones adds the "milestones" edges to the Milestone entity.
func (ac *AccountCreate) AddMilestones(m ...*Milestone) *AccountCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ac.AddMilestoneIDs(ids...)
}

// AddIssueIDs adds the "issues" edge to the Issue entity by IDs.
func (ac *AccountCreate) AddIssueIDs(ids ...int) *AccountCreate {
	ac.mutation.AddIssueIDs(ids...)
	return ac
}

// AddIssues adds the "issues" edges to the Issue entity.
func (ac *AccountCreate) AddIssues(i ...*Issue) *AccountCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ac.AddIssueIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AccountCreate) defaults() {
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := account.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := account.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UUID(); !ok {
		v := account.DefaultUUID()
		ac.mutation.SetUUID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "Account.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := account.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Account.name": %w`, err)}
		}
	}
	if v, ok := ac.mutation.Email(); ok {
		if err := account.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`generated: validator failed for field "Account.email": %w`, err)}
		}
	}
	if v, ok := ac.mutation.PhoneNumber(); ok {
		if err := account.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`generated: validator failed for field "Account.phone_number": %w`, err)}
		}
	}
	if _, ok := ac.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`generated: missing required field "Account.uuid"`)}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(account.Table, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(account.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Email(); ok {
		_spec.SetField(account.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := ac.mutation.PhoneNumber(); ok {
		_spec.SetField(account.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UUID(); ok {
		_spec.SetField(account.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if nodes := ac.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.UsersTable,
			Columns: []string{account.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.ProjectsTable,
			Columns: []string{account.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.SubscriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.SubscriptionsTable,
			Columns: []string{account.SubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.RolesTable,
			Columns: []string{account.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TasksTable,
			Columns: []string{account.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.MilestonesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.MilestonesTable,
			Columns: []string{account.MilestonesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(milestone.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.IssuesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.IssuesTable,
			Columns: []string{account.IssuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	err      error
	builders []*AccountCreate
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
