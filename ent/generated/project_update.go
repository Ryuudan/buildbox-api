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
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/project"
	"github.com/Pyakz/buildbox-api/ent/generated/task"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetAccountID sets the "account_id" field.
func (pu *ProjectUpdate) SetAccountID(i int) *ProjectUpdate {
	pu.mutation.SetAccountID(i)
	return pu
}

// SetCreatedBy sets the "created_by" field.
func (pu *ProjectUpdate) SetCreatedBy(i int) *ProjectUpdate {
	pu.mutation.ResetCreatedBy()
	pu.mutation.SetCreatedBy(i)
	return pu
}

// AddCreatedBy adds i to the "created_by" field.
func (pu *ProjectUpdate) AddCreatedBy(i int) *ProjectUpdate {
	pu.mutation.AddCreatedBy(i)
	return pu
}

// SetClientID sets the "client_id" field.
func (pu *ProjectUpdate) SetClientID(i int) *ProjectUpdate {
	pu.mutation.ResetClientID()
	pu.mutation.SetClientID(i)
	return pu
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableClientID(i *int) *ProjectUpdate {
	if i != nil {
		pu.SetClientID(*i)
	}
	return pu
}

// AddClientID adds i to the "client_id" field.
func (pu *ProjectUpdate) AddClientID(i int) *ProjectUpdate {
	pu.mutation.AddClientID(i)
	return pu
}

// ClearClientID clears the value of the "client_id" field.
func (pu *ProjectUpdate) ClearClientID() *ProjectUpdate {
	pu.mutation.ClearClientID()
	return pu
}

// SetManagerID sets the "manager_id" field.
func (pu *ProjectUpdate) SetManagerID(i int) *ProjectUpdate {
	pu.mutation.ResetManagerID()
	pu.mutation.SetManagerID(i)
	return pu
}

// SetNillableManagerID sets the "manager_id" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableManagerID(i *int) *ProjectUpdate {
	if i != nil {
		pu.SetManagerID(*i)
	}
	return pu
}

// AddManagerID adds i to the "manager_id" field.
func (pu *ProjectUpdate) AddManagerID(i int) *ProjectUpdate {
	pu.mutation.AddManagerID(i)
	return pu
}

// ClearManagerID clears the value of the "manager_id" field.
func (pu *ProjectUpdate) ClearManagerID() *ProjectUpdate {
	pu.mutation.ClearManagerID()
	return pu
}

// SetName sets the "name" field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetStatus sets the "status" field.
func (pu *ProjectUpdate) SetStatus(pr project.Status) *ProjectUpdate {
	pu.mutation.SetStatus(pr)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableStatus(pr *project.Status) *ProjectUpdate {
	if pr != nil {
		pu.SetStatus(*pr)
	}
	return pu
}

// ClearStatus clears the value of the "status" field.
func (pu *ProjectUpdate) ClearStatus() *ProjectUpdate {
	pu.mutation.ClearStatus()
	return pu
}

// SetLocation sets the "location" field.
func (pu *ProjectUpdate) SetLocation(s string) *ProjectUpdate {
	pu.mutation.SetLocation(s)
	return pu
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableLocation(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetLocation(*s)
	}
	return pu
}

// ClearLocation clears the value of the "location" field.
func (pu *ProjectUpdate) ClearLocation() *ProjectUpdate {
	pu.mutation.ClearLocation()
	return pu
}

// SetBudget sets the "budget" field.
func (pu *ProjectUpdate) SetBudget(f float64) *ProjectUpdate {
	pu.mutation.ResetBudget()
	pu.mutation.SetBudget(f)
	return pu
}

// SetNillableBudget sets the "budget" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableBudget(f *float64) *ProjectUpdate {
	if f != nil {
		pu.SetBudget(*f)
	}
	return pu
}

// AddBudget adds f to the "budget" field.
func (pu *ProjectUpdate) AddBudget(f float64) *ProjectUpdate {
	pu.mutation.AddBudget(f)
	return pu
}

// ClearBudget clears the value of the "budget" field.
func (pu *ProjectUpdate) ClearBudget() *ProjectUpdate {
	pu.mutation.ClearBudget()
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProjectUpdate) SetDescription(s string) *ProjectUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDescription(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *ProjectUpdate) ClearDescription() *ProjectUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetNotes sets the "notes" field.
func (pu *ProjectUpdate) SetNotes(s string) *ProjectUpdate {
	pu.mutation.SetNotes(s)
	return pu
}

// SetNillableNotes sets the "notes" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableNotes(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetNotes(*s)
	}
	return pu
}

// ClearNotes clears the value of the "notes" field.
func (pu *ProjectUpdate) ClearNotes() *ProjectUpdate {
	pu.mutation.ClearNotes()
	return pu
}

// SetStartDate sets the "start_date" field.
func (pu *ProjectUpdate) SetStartDate(t time.Time) *ProjectUpdate {
	pu.mutation.SetStartDate(t)
	return pu
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableStartDate(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetStartDate(*t)
	}
	return pu
}

// ClearStartDate clears the value of the "start_date" field.
func (pu *ProjectUpdate) ClearStartDate() *ProjectUpdate {
	pu.mutation.ClearStartDate()
	return pu
}

// SetEndDate sets the "end_date" field.
func (pu *ProjectUpdate) SetEndDate(t time.Time) *ProjectUpdate {
	pu.mutation.SetEndDate(t)
	return pu
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableEndDate(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetEndDate(*t)
	}
	return pu
}

// ClearEndDate clears the value of the "end_date" field.
func (pu *ProjectUpdate) ClearEndDate() *ProjectUpdate {
	pu.mutation.ClearEndDate()
	return pu
}

// SetDeleted sets the "deleted" field.
func (pu *ProjectUpdate) SetDeleted(b bool) *ProjectUpdate {
	pu.mutation.SetDeleted(b)
	return pu
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDeleted(b *bool) *ProjectUpdate {
	if b != nil {
		pu.SetDeleted(*b)
	}
	return pu
}

// ClearDeleted clears the value of the "deleted" field.
func (pu *ProjectUpdate) ClearDeleted() *ProjectUpdate {
	pu.mutation.ClearDeleted()
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProjectUpdate) SetUpdatedAt(t time.Time) *ProjectUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableUpdatedAt(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetUpdatedAt(*t)
	}
	return pu
}

// SetAccount sets the "account" edge to the Account entity.
func (pu *ProjectUpdate) SetAccount(a *Account) *ProjectUpdate {
	return pu.SetAccountID(a.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (pu *ProjectUpdate) AddTaskIDs(ids ...int) *ProjectUpdate {
	pu.mutation.AddTaskIDs(ids...)
	return pu
}

// AddTasks adds the "tasks" edges to the Task entity.
func (pu *ProjectUpdate) AddTasks(t ...*Task) *ProjectUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.AddTaskIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (pu *ProjectUpdate) ClearAccount() *ProjectUpdate {
	pu.mutation.ClearAccount()
	return pu
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (pu *ProjectUpdate) ClearTasks() *ProjectUpdate {
	pu.mutation.ClearTasks()
	return pu
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (pu *ProjectUpdate) RemoveTaskIDs(ids ...int) *ProjectUpdate {
	pu.mutation.RemoveTaskIDs(ids...)
	return pu
}

// RemoveTasks removes "tasks" edges to Task entities.
func (pu *ProjectUpdate) RemoveTasks(t ...*Task) *ProjectUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.RemoveTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProjectUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Project.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := project.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "Project.status": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Budget(); ok {
		if err := project.BudgetValidator(v); err != nil {
			return &ValidationError{Name: "budget", err: fmt.Errorf(`generated: validator failed for field "Project.budget": %w`, err)}
		}
	}
	if _, ok := pu.mutation.AccountID(); pu.mutation.AccountCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Project.account"`)
	}
	return nil
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.CreatedBy(); ok {
		_spec.SetField(project.FieldCreatedBy, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(project.FieldCreatedBy, field.TypeInt, value)
	}
	if value, ok := pu.mutation.ClientID(); ok {
		_spec.SetField(project.FieldClientID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedClientID(); ok {
		_spec.AddField(project.FieldClientID, field.TypeInt, value)
	}
	if pu.mutation.ClientIDCleared() {
		_spec.ClearField(project.FieldClientID, field.TypeInt)
	}
	if value, ok := pu.mutation.ManagerID(); ok {
		_spec.SetField(project.FieldManagerID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedManagerID(); ok {
		_spec.AddField(project.FieldManagerID, field.TypeInt, value)
	}
	if pu.mutation.ManagerIDCleared() {
		_spec.ClearField(project.FieldManagerID, field.TypeInt)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(project.FieldStatus, field.TypeEnum, value)
	}
	if pu.mutation.StatusCleared() {
		_spec.ClearField(project.FieldStatus, field.TypeEnum)
	}
	if value, ok := pu.mutation.Location(); ok {
		_spec.SetField(project.FieldLocation, field.TypeString, value)
	}
	if pu.mutation.LocationCleared() {
		_spec.ClearField(project.FieldLocation, field.TypeString)
	}
	if value, ok := pu.mutation.Budget(); ok {
		_spec.SetField(project.FieldBudget, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedBudget(); ok {
		_spec.AddField(project.FieldBudget, field.TypeFloat64, value)
	}
	if pu.mutation.BudgetCleared() {
		_spec.ClearField(project.FieldBudget, field.TypeFloat64)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(project.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.Notes(); ok {
		_spec.SetField(project.FieldNotes, field.TypeString, value)
	}
	if pu.mutation.NotesCleared() {
		_spec.ClearField(project.FieldNotes, field.TypeString)
	}
	if value, ok := pu.mutation.StartDate(); ok {
		_spec.SetField(project.FieldStartDate, field.TypeTime, value)
	}
	if pu.mutation.StartDateCleared() {
		_spec.ClearField(project.FieldStartDate, field.TypeTime)
	}
	if value, ok := pu.mutation.EndDate(); ok {
		_spec.SetField(project.FieldEndDate, field.TypeTime, value)
	}
	if pu.mutation.EndDateCleared() {
		_spec.ClearField(project.FieldEndDate, field.TypeTime)
	}
	if value, ok := pu.mutation.Deleted(); ok {
		_spec.SetField(project.FieldDeleted, field.TypeBool, value)
	}
	if pu.mutation.DeletedCleared() {
		_spec.ClearField(project.FieldDeleted, field.TypeBool)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(project.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.AccountTable,
			Columns: []string{project.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.AccountTable,
			Columns: []string{project.AccountColumn},
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
	if pu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TasksTable,
			Columns: []string{project.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !pu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TasksTable,
			Columns: []string{project.TasksColumn},
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
	if nodes := pu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TasksTable,
			Columns: []string{project.TasksColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMutation
}

// SetAccountID sets the "account_id" field.
func (puo *ProjectUpdateOne) SetAccountID(i int) *ProjectUpdateOne {
	puo.mutation.SetAccountID(i)
	return puo
}

// SetCreatedBy sets the "created_by" field.
func (puo *ProjectUpdateOne) SetCreatedBy(i int) *ProjectUpdateOne {
	puo.mutation.ResetCreatedBy()
	puo.mutation.SetCreatedBy(i)
	return puo
}

// AddCreatedBy adds i to the "created_by" field.
func (puo *ProjectUpdateOne) AddCreatedBy(i int) *ProjectUpdateOne {
	puo.mutation.AddCreatedBy(i)
	return puo
}

// SetClientID sets the "client_id" field.
func (puo *ProjectUpdateOne) SetClientID(i int) *ProjectUpdateOne {
	puo.mutation.ResetClientID()
	puo.mutation.SetClientID(i)
	return puo
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableClientID(i *int) *ProjectUpdateOne {
	if i != nil {
		puo.SetClientID(*i)
	}
	return puo
}

// AddClientID adds i to the "client_id" field.
func (puo *ProjectUpdateOne) AddClientID(i int) *ProjectUpdateOne {
	puo.mutation.AddClientID(i)
	return puo
}

// ClearClientID clears the value of the "client_id" field.
func (puo *ProjectUpdateOne) ClearClientID() *ProjectUpdateOne {
	puo.mutation.ClearClientID()
	return puo
}

// SetManagerID sets the "manager_id" field.
func (puo *ProjectUpdateOne) SetManagerID(i int) *ProjectUpdateOne {
	puo.mutation.ResetManagerID()
	puo.mutation.SetManagerID(i)
	return puo
}

// SetNillableManagerID sets the "manager_id" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableManagerID(i *int) *ProjectUpdateOne {
	if i != nil {
		puo.SetManagerID(*i)
	}
	return puo
}

// AddManagerID adds i to the "manager_id" field.
func (puo *ProjectUpdateOne) AddManagerID(i int) *ProjectUpdateOne {
	puo.mutation.AddManagerID(i)
	return puo
}

// ClearManagerID clears the value of the "manager_id" field.
func (puo *ProjectUpdateOne) ClearManagerID() *ProjectUpdateOne {
	puo.mutation.ClearManagerID()
	return puo
}

// SetName sets the "name" field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetStatus sets the "status" field.
func (puo *ProjectUpdateOne) SetStatus(pr project.Status) *ProjectUpdateOne {
	puo.mutation.SetStatus(pr)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableStatus(pr *project.Status) *ProjectUpdateOne {
	if pr != nil {
		puo.SetStatus(*pr)
	}
	return puo
}

// ClearStatus clears the value of the "status" field.
func (puo *ProjectUpdateOne) ClearStatus() *ProjectUpdateOne {
	puo.mutation.ClearStatus()
	return puo
}

// SetLocation sets the "location" field.
func (puo *ProjectUpdateOne) SetLocation(s string) *ProjectUpdateOne {
	puo.mutation.SetLocation(s)
	return puo
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableLocation(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetLocation(*s)
	}
	return puo
}

// ClearLocation clears the value of the "location" field.
func (puo *ProjectUpdateOne) ClearLocation() *ProjectUpdateOne {
	puo.mutation.ClearLocation()
	return puo
}

// SetBudget sets the "budget" field.
func (puo *ProjectUpdateOne) SetBudget(f float64) *ProjectUpdateOne {
	puo.mutation.ResetBudget()
	puo.mutation.SetBudget(f)
	return puo
}

// SetNillableBudget sets the "budget" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableBudget(f *float64) *ProjectUpdateOne {
	if f != nil {
		puo.SetBudget(*f)
	}
	return puo
}

// AddBudget adds f to the "budget" field.
func (puo *ProjectUpdateOne) AddBudget(f float64) *ProjectUpdateOne {
	puo.mutation.AddBudget(f)
	return puo
}

// ClearBudget clears the value of the "budget" field.
func (puo *ProjectUpdateOne) ClearBudget() *ProjectUpdateOne {
	puo.mutation.ClearBudget()
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProjectUpdateOne) SetDescription(s string) *ProjectUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDescription(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *ProjectUpdateOne) ClearDescription() *ProjectUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetNotes sets the "notes" field.
func (puo *ProjectUpdateOne) SetNotes(s string) *ProjectUpdateOne {
	puo.mutation.SetNotes(s)
	return puo
}

// SetNillableNotes sets the "notes" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableNotes(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetNotes(*s)
	}
	return puo
}

// ClearNotes clears the value of the "notes" field.
func (puo *ProjectUpdateOne) ClearNotes() *ProjectUpdateOne {
	puo.mutation.ClearNotes()
	return puo
}

// SetStartDate sets the "start_date" field.
func (puo *ProjectUpdateOne) SetStartDate(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetStartDate(t)
	return puo
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableStartDate(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetStartDate(*t)
	}
	return puo
}

// ClearStartDate clears the value of the "start_date" field.
func (puo *ProjectUpdateOne) ClearStartDate() *ProjectUpdateOne {
	puo.mutation.ClearStartDate()
	return puo
}

// SetEndDate sets the "end_date" field.
func (puo *ProjectUpdateOne) SetEndDate(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetEndDate(t)
	return puo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableEndDate(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetEndDate(*t)
	}
	return puo
}

// ClearEndDate clears the value of the "end_date" field.
func (puo *ProjectUpdateOne) ClearEndDate() *ProjectUpdateOne {
	puo.mutation.ClearEndDate()
	return puo
}

// SetDeleted sets the "deleted" field.
func (puo *ProjectUpdateOne) SetDeleted(b bool) *ProjectUpdateOne {
	puo.mutation.SetDeleted(b)
	return puo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDeleted(b *bool) *ProjectUpdateOne {
	if b != nil {
		puo.SetDeleted(*b)
	}
	return puo
}

// ClearDeleted clears the value of the "deleted" field.
func (puo *ProjectUpdateOne) ClearDeleted() *ProjectUpdateOne {
	puo.mutation.ClearDeleted()
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProjectUpdateOne) SetUpdatedAt(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableUpdatedAt(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetUpdatedAt(*t)
	}
	return puo
}

// SetAccount sets the "account" edge to the Account entity.
func (puo *ProjectUpdateOne) SetAccount(a *Account) *ProjectUpdateOne {
	return puo.SetAccountID(a.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (puo *ProjectUpdateOne) AddTaskIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.AddTaskIDs(ids...)
	return puo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (puo *ProjectUpdateOne) AddTasks(t ...*Task) *ProjectUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.AddTaskIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (puo *ProjectUpdateOne) ClearAccount() *ProjectUpdateOne {
	puo.mutation.ClearAccount()
	return puo
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (puo *ProjectUpdateOne) ClearTasks() *ProjectUpdateOne {
	puo.mutation.ClearTasks()
	return puo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (puo *ProjectUpdateOne) RemoveTaskIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.RemoveTaskIDs(ids...)
	return puo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (puo *ProjectUpdateOne) RemoveTasks(t ...*Task) *ProjectUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.RemoveTaskIDs(ids...)
}

// Where appends a list predicates to the ProjectUpdate builder.
func (puo *ProjectUpdateOne) Where(ps ...predicate.Project) *ProjectUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProjectUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Project.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := project.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "Project.status": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Budget(); ok {
		if err := project.BudgetValidator(v); err != nil {
			return &ValidationError{Name: "budget", err: fmt.Errorf(`generated: validator failed for field "Project.budget": %w`, err)}
		}
	}
	if _, ok := puo.mutation.AccountID(); puo.mutation.AccountCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Project.account"`)
	}
	return nil
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Project.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != project.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.CreatedBy(); ok {
		_spec.SetField(project.FieldCreatedBy, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(project.FieldCreatedBy, field.TypeInt, value)
	}
	if value, ok := puo.mutation.ClientID(); ok {
		_spec.SetField(project.FieldClientID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedClientID(); ok {
		_spec.AddField(project.FieldClientID, field.TypeInt, value)
	}
	if puo.mutation.ClientIDCleared() {
		_spec.ClearField(project.FieldClientID, field.TypeInt)
	}
	if value, ok := puo.mutation.ManagerID(); ok {
		_spec.SetField(project.FieldManagerID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedManagerID(); ok {
		_spec.AddField(project.FieldManagerID, field.TypeInt, value)
	}
	if puo.mutation.ManagerIDCleared() {
		_spec.ClearField(project.FieldManagerID, field.TypeInt)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(project.FieldStatus, field.TypeEnum, value)
	}
	if puo.mutation.StatusCleared() {
		_spec.ClearField(project.FieldStatus, field.TypeEnum)
	}
	if value, ok := puo.mutation.Location(); ok {
		_spec.SetField(project.FieldLocation, field.TypeString, value)
	}
	if puo.mutation.LocationCleared() {
		_spec.ClearField(project.FieldLocation, field.TypeString)
	}
	if value, ok := puo.mutation.Budget(); ok {
		_spec.SetField(project.FieldBudget, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedBudget(); ok {
		_spec.AddField(project.FieldBudget, field.TypeFloat64, value)
	}
	if puo.mutation.BudgetCleared() {
		_spec.ClearField(project.FieldBudget, field.TypeFloat64)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(project.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.Notes(); ok {
		_spec.SetField(project.FieldNotes, field.TypeString, value)
	}
	if puo.mutation.NotesCleared() {
		_spec.ClearField(project.FieldNotes, field.TypeString)
	}
	if value, ok := puo.mutation.StartDate(); ok {
		_spec.SetField(project.FieldStartDate, field.TypeTime, value)
	}
	if puo.mutation.StartDateCleared() {
		_spec.ClearField(project.FieldStartDate, field.TypeTime)
	}
	if value, ok := puo.mutation.EndDate(); ok {
		_spec.SetField(project.FieldEndDate, field.TypeTime, value)
	}
	if puo.mutation.EndDateCleared() {
		_spec.ClearField(project.FieldEndDate, field.TypeTime)
	}
	if value, ok := puo.mutation.Deleted(); ok {
		_spec.SetField(project.FieldDeleted, field.TypeBool, value)
	}
	if puo.mutation.DeletedCleared() {
		_spec.ClearField(project.FieldDeleted, field.TypeBool)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(project.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.AccountTable,
			Columns: []string{project.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.AccountTable,
			Columns: []string{project.AccountColumn},
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
	if puo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TasksTable,
			Columns: []string{project.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !puo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TasksTable,
			Columns: []string{project.TasksColumn},
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
	if nodes := puo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TasksTable,
			Columns: []string{project.TasksColumn},
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
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
