// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Pyakz/buildbox-api/ent/generated/account"
	"github.com/Pyakz/buildbox-api/ent/generated/issue"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/project"
	"github.com/Pyakz/buildbox-api/ent/generated/task"
	"github.com/Pyakz/buildbox-api/ent/generated/user"
)

// IssueUpdate is the builder for updating Issue entities.
type IssueUpdate struct {
	config
	hooks    []Hook
	mutation *IssueMutation
}

// Where appends a list predicates to the IssueUpdate builder.
func (iu *IssueUpdate) Where(ps ...predicate.Issue) *IssueUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetAccountID sets the "account_id" field.
func (iu *IssueUpdate) SetAccountID(i int) *IssueUpdate {
	iu.mutation.SetAccountID(i)
	return iu
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (iu *IssueUpdate) SetNillableAccountID(i *int) *IssueUpdate {
	if i != nil {
		iu.SetAccountID(*i)
	}
	return iu
}

// SetCreatedBy sets the "created_by" field.
func (iu *IssueUpdate) SetCreatedBy(i int) *IssueUpdate {
	iu.mutation.SetCreatedBy(i)
	return iu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (iu *IssueUpdate) SetNillableCreatedBy(i *int) *IssueUpdate {
	if i != nil {
		iu.SetCreatedBy(*i)
	}
	return iu
}

// SetProjectID sets the "project_id" field.
func (iu *IssueUpdate) SetProjectID(i int) *IssueUpdate {
	iu.mutation.SetProjectID(i)
	return iu
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (iu *IssueUpdate) SetNillableProjectID(i *int) *IssueUpdate {
	if i != nil {
		iu.SetProjectID(*i)
	}
	return iu
}

// SetTitle sets the "title" field.
func (iu *IssueUpdate) SetTitle(s string) *IssueUpdate {
	iu.mutation.SetTitle(s)
	return iu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (iu *IssueUpdate) SetNillableTitle(s *string) *IssueUpdate {
	if s != nil {
		iu.SetTitle(*s)
	}
	return iu
}

// SetDescription sets the "description" field.
func (iu *IssueUpdate) SetDescription(s string) *IssueUpdate {
	iu.mutation.SetDescription(s)
	return iu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iu *IssueUpdate) SetNillableDescription(s *string) *IssueUpdate {
	if s != nil {
		iu.SetDescription(*s)
	}
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *IssueUpdate) SetUpdatedAt(t time.Time) *IssueUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (iu *IssueUpdate) SetNillableUpdatedAt(t *time.Time) *IssueUpdate {
	if t != nil {
		iu.SetUpdatedAt(*t)
	}
	return iu
}

// SetDeleted sets the "deleted" field.
func (iu *IssueUpdate) SetDeleted(b bool) *IssueUpdate {
	iu.mutation.SetDeleted(b)
	return iu
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (iu *IssueUpdate) SetNillableDeleted(b *bool) *IssueUpdate {
	if b != nil {
		iu.SetDeleted(*b)
	}
	return iu
}

// ClearDeleted clears the value of the "deleted" field.
func (iu *IssueUpdate) ClearDeleted() *IssueUpdate {
	iu.mutation.ClearDeleted()
	return iu
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (iu *IssueUpdate) AddTaskIDs(ids ...int) *IssueUpdate {
	iu.mutation.AddTaskIDs(ids...)
	return iu
}

// AddTasks adds the "tasks" edges to the Task entity.
func (iu *IssueUpdate) AddTasks(t ...*Task) *IssueUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return iu.AddTaskIDs(ids...)
}

// SetAccount sets the "account" edge to the Account entity.
func (iu *IssueUpdate) SetAccount(a *Account) *IssueUpdate {
	return iu.SetAccountID(a.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (iu *IssueUpdate) SetUserID(id int) *IssueUpdate {
	iu.mutation.SetUserID(id)
	return iu
}

// SetUser sets the "user" edge to the User entity.
func (iu *IssueUpdate) SetUser(u *User) *IssueUpdate {
	return iu.SetUserID(u.ID)
}

// SetProject sets the "project" edge to the Project entity.
func (iu *IssueUpdate) SetProject(p *Project) *IssueUpdate {
	return iu.SetProjectID(p.ID)
}

// Mutation returns the IssueMutation object of the builder.
func (iu *IssueUpdate) Mutation() *IssueMutation {
	return iu.mutation
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (iu *IssueUpdate) ClearTasks() *IssueUpdate {
	iu.mutation.ClearTasks()
	return iu
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (iu *IssueUpdate) RemoveTaskIDs(ids ...int) *IssueUpdate {
	iu.mutation.RemoveTaskIDs(ids...)
	return iu
}

// RemoveTasks removes "tasks" edges to Task entities.
func (iu *IssueUpdate) RemoveTasks(t ...*Task) *IssueUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return iu.RemoveTaskIDs(ids...)
}

// ClearAccount clears the "account" edge to the Account entity.
func (iu *IssueUpdate) ClearAccount() *IssueUpdate {
	iu.mutation.ClearAccount()
	return iu
}

// ClearUser clears the "user" edge to the User entity.
func (iu *IssueUpdate) ClearUser() *IssueUpdate {
	iu.mutation.ClearUser()
	return iu
}

// ClearProject clears the "project" edge to the Project entity.
func (iu *IssueUpdate) ClearProject() *IssueUpdate {
	iu.mutation.ClearProject()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IssueUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IssueUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IssueUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IssueUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *IssueUpdate) check() error {
	if v, ok := iu.mutation.Title(); ok {
		if err := issue.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`generated: validator failed for field "Issue.title": %w`, err)}
		}
	}
	if v, ok := iu.mutation.Description(); ok {
		if err := issue.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`generated: validator failed for field "Issue.description": %w`, err)}
		}
	}
	if _, ok := iu.mutation.AccountID(); iu.mutation.AccountCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Issue.account"`)
	}
	if _, ok := iu.mutation.UserID(); iu.mutation.UserCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Issue.user"`)
	}
	if _, ok := iu.mutation.ProjectID(); iu.mutation.ProjectCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Issue.project"`)
	}
	return nil
}

func (iu *IssueUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(issue.Table, issue.Columns, sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Title(); ok {
		_spec.SetField(issue.FieldTitle, field.TypeString, value)
	}
	if value, ok := iu.mutation.Description(); ok {
		_spec.SetField(issue.FieldDescription, field.TypeString, value)
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(issue.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iu.mutation.Deleted(); ok {
		_spec.SetField(issue.FieldDeleted, field.TypeBool, value)
	}
	if iu.mutation.DeletedCleared() {
		_spec.ClearField(issue.FieldDeleted, field.TypeBool)
	}
	if iu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   issue.TasksTable,
			Columns: []string{issue.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !iu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   issue.TasksTable,
			Columns: []string{issue.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   issue.TasksTable,
			Columns: []string{issue.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issue.AccountTable,
			Columns: []string{issue.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issue.AccountTable,
			Columns: []string{issue.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issue.UserTable,
			Columns: []string{issue.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issue.UserTable,
			Columns: []string{issue.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issue.ProjectTable,
			Columns: []string{issue.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issue.ProjectTable,
			Columns: []string{issue.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{issue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// IssueUpdateOne is the builder for updating a single Issue entity.
type IssueUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IssueMutation
}

// SetAccountID sets the "account_id" field.
func (iuo *IssueUpdateOne) SetAccountID(i int) *IssueUpdateOne {
	iuo.mutation.SetAccountID(i)
	return iuo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (iuo *IssueUpdateOne) SetNillableAccountID(i *int) *IssueUpdateOne {
	if i != nil {
		iuo.SetAccountID(*i)
	}
	return iuo
}

// SetCreatedBy sets the "created_by" field.
func (iuo *IssueUpdateOne) SetCreatedBy(i int) *IssueUpdateOne {
	iuo.mutation.SetCreatedBy(i)
	return iuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (iuo *IssueUpdateOne) SetNillableCreatedBy(i *int) *IssueUpdateOne {
	if i != nil {
		iuo.SetCreatedBy(*i)
	}
	return iuo
}

// SetProjectID sets the "project_id" field.
func (iuo *IssueUpdateOne) SetProjectID(i int) *IssueUpdateOne {
	iuo.mutation.SetProjectID(i)
	return iuo
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (iuo *IssueUpdateOne) SetNillableProjectID(i *int) *IssueUpdateOne {
	if i != nil {
		iuo.SetProjectID(*i)
	}
	return iuo
}

// SetTitle sets the "title" field.
func (iuo *IssueUpdateOne) SetTitle(s string) *IssueUpdateOne {
	iuo.mutation.SetTitle(s)
	return iuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (iuo *IssueUpdateOne) SetNillableTitle(s *string) *IssueUpdateOne {
	if s != nil {
		iuo.SetTitle(*s)
	}
	return iuo
}

// SetDescription sets the "description" field.
func (iuo *IssueUpdateOne) SetDescription(s string) *IssueUpdateOne {
	iuo.mutation.SetDescription(s)
	return iuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iuo *IssueUpdateOne) SetNillableDescription(s *string) *IssueUpdateOne {
	if s != nil {
		iuo.SetDescription(*s)
	}
	return iuo
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *IssueUpdateOne) SetUpdatedAt(t time.Time) *IssueUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (iuo *IssueUpdateOne) SetNillableUpdatedAt(t *time.Time) *IssueUpdateOne {
	if t != nil {
		iuo.SetUpdatedAt(*t)
	}
	return iuo
}

// SetDeleted sets the "deleted" field.
func (iuo *IssueUpdateOne) SetDeleted(b bool) *IssueUpdateOne {
	iuo.mutation.SetDeleted(b)
	return iuo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (iuo *IssueUpdateOne) SetNillableDeleted(b *bool) *IssueUpdateOne {
	if b != nil {
		iuo.SetDeleted(*b)
	}
	return iuo
}

// ClearDeleted clears the value of the "deleted" field.
func (iuo *IssueUpdateOne) ClearDeleted() *IssueUpdateOne {
	iuo.mutation.ClearDeleted()
	return iuo
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (iuo *IssueUpdateOne) AddTaskIDs(ids ...int) *IssueUpdateOne {
	iuo.mutation.AddTaskIDs(ids...)
	return iuo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (iuo *IssueUpdateOne) AddTasks(t ...*Task) *IssueUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return iuo.AddTaskIDs(ids...)
}

// SetAccount sets the "account" edge to the Account entity.
func (iuo *IssueUpdateOne) SetAccount(a *Account) *IssueUpdateOne {
	return iuo.SetAccountID(a.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (iuo *IssueUpdateOne) SetUserID(id int) *IssueUpdateOne {
	iuo.mutation.SetUserID(id)
	return iuo
}

// SetUser sets the "user" edge to the User entity.
func (iuo *IssueUpdateOne) SetUser(u *User) *IssueUpdateOne {
	return iuo.SetUserID(u.ID)
}

// SetProject sets the "project" edge to the Project entity.
func (iuo *IssueUpdateOne) SetProject(p *Project) *IssueUpdateOne {
	return iuo.SetProjectID(p.ID)
}

// Mutation returns the IssueMutation object of the builder.
func (iuo *IssueUpdateOne) Mutation() *IssueMutation {
	return iuo.mutation
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (iuo *IssueUpdateOne) ClearTasks() *IssueUpdateOne {
	iuo.mutation.ClearTasks()
	return iuo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (iuo *IssueUpdateOne) RemoveTaskIDs(ids ...int) *IssueUpdateOne {
	iuo.mutation.RemoveTaskIDs(ids...)
	return iuo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (iuo *IssueUpdateOne) RemoveTasks(t ...*Task) *IssueUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return iuo.RemoveTaskIDs(ids...)
}

// ClearAccount clears the "account" edge to the Account entity.
func (iuo *IssueUpdateOne) ClearAccount() *IssueUpdateOne {
	iuo.mutation.ClearAccount()
	return iuo
}

// ClearUser clears the "user" edge to the User entity.
func (iuo *IssueUpdateOne) ClearUser() *IssueUpdateOne {
	iuo.mutation.ClearUser()
	return iuo
}

// ClearProject clears the "project" edge to the Project entity.
func (iuo *IssueUpdateOne) ClearProject() *IssueUpdateOne {
	iuo.mutation.ClearProject()
	return iuo
}

// Where appends a list predicates to the IssueUpdate builder.
func (iuo *IssueUpdateOne) Where(ps ...predicate.Issue) *IssueUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IssueUpdateOne) Select(field string, fields ...string) *IssueUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Issue entity.
func (iuo *IssueUpdateOne) Save(ctx context.Context) (*Issue, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IssueUpdateOne) SaveX(ctx context.Context) *Issue {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IssueUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IssueUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *IssueUpdateOne) check() error {
	if v, ok := iuo.mutation.Title(); ok {
		if err := issue.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`generated: validator failed for field "Issue.title": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.Description(); ok {
		if err := issue.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`generated: validator failed for field "Issue.description": %w`, err)}
		}
	}
	if _, ok := iuo.mutation.AccountID(); iuo.mutation.AccountCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Issue.account"`)
	}
	if _, ok := iuo.mutation.UserID(); iuo.mutation.UserCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Issue.user"`)
	}
	if _, ok := iuo.mutation.ProjectID(); iuo.mutation.ProjectCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Issue.project"`)
	}
	return nil
}

func (iuo *IssueUpdateOne) sqlSave(ctx context.Context) (_node *Issue, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(issue.Table, issue.Columns, sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Issue.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, issue.FieldID)
		for _, f := range fields {
			if !issue.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != issue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Title(); ok {
		_spec.SetField(issue.FieldTitle, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Description(); ok {
		_spec.SetField(issue.FieldDescription, field.TypeString, value)
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(issue.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.Deleted(); ok {
		_spec.SetField(issue.FieldDeleted, field.TypeBool, value)
	}
	if iuo.mutation.DeletedCleared() {
		_spec.ClearField(issue.FieldDeleted, field.TypeBool)
	}
	if iuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   issue.TasksTable,
			Columns: []string{issue.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !iuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   issue.TasksTable,
			Columns: []string{issue.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   issue.TasksTable,
			Columns: []string{issue.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issue.AccountTable,
			Columns: []string{issue.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issue.AccountTable,
			Columns: []string{issue.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issue.UserTable,
			Columns: []string{issue.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issue.UserTable,
			Columns: []string{issue.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issue.ProjectTable,
			Columns: []string{issue.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issue.ProjectTable,
			Columns: []string{issue.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Issue{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{issue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
