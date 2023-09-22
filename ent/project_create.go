// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Pyakz/buildbox-api/ent/project"
	"github.com/google/uuid"
)

// ProjectCreate is the builder for creating a Project entity.
type ProjectCreate struct {
	config
	mutation *ProjectMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (pc *ProjectCreate) SetUUID(u uuid.UUID) *ProjectCreate {
	pc.mutation.SetUUID(u)
	return pc
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUUID(u *uuid.UUID) *ProjectCreate {
	if u != nil {
		pc.SetUUID(*u)
	}
	return pc
}

// SetAccountID sets the "account_id" field.
func (pc *ProjectCreate) SetAccountID(s string) *ProjectCreate {
	pc.mutation.SetAccountID(s)
	return pc
}

// SetClientID sets the "client_id" field.
func (pc *ProjectCreate) SetClientID(s string) *ProjectCreate {
	pc.mutation.SetClientID(s)
	return pc
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableClientID(s *string) *ProjectCreate {
	if s != nil {
		pc.SetClientID(*s)
	}
	return pc
}

// SetManagerID sets the "manager_id" field.
func (pc *ProjectCreate) SetManagerID(s string) *ProjectCreate {
	pc.mutation.SetManagerID(s)
	return pc
}

// SetNillableManagerID sets the "manager_id" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableManagerID(s *string) *ProjectCreate {
	if s != nil {
		pc.SetManagerID(*s)
	}
	return pc
}

// SetCreatedBy sets the "created_by" field.
func (pc *ProjectCreate) SetCreatedBy(s string) *ProjectCreate {
	pc.mutation.SetCreatedBy(s)
	return pc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableCreatedBy(s *string) *ProjectCreate {
	if s != nil {
		pc.SetCreatedBy(*s)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *ProjectCreate) SetName(s string) *ProjectCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProjectCreate) SetDescription(s string) *ProjectCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableDescription(s *string) *ProjectCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetNotes sets the "notes" field.
func (pc *ProjectCreate) SetNotes(s string) *ProjectCreate {
	pc.mutation.SetNotes(s)
	return pc
}

// SetNillableNotes sets the "notes" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableNotes(s *string) *ProjectCreate {
	if s != nil {
		pc.SetNotes(*s)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *ProjectCreate) SetStatus(pr project.Status) *ProjectCreate {
	pc.mutation.SetStatus(pr)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableStatus(pr *project.Status) *ProjectCreate {
	if pr != nil {
		pc.SetStatus(*pr)
	}
	return pc
}

// SetLocation sets the "location" field.
func (pc *ProjectCreate) SetLocation(s string) *ProjectCreate {
	pc.mutation.SetLocation(s)
	return pc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableLocation(s *string) *ProjectCreate {
	if s != nil {
		pc.SetLocation(*s)
	}
	return pc
}

// SetBudget sets the "budget" field.
func (pc *ProjectCreate) SetBudget(f float64) *ProjectCreate {
	pc.mutation.SetBudget(f)
	return pc
}

// SetNillableBudget sets the "budget" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableBudget(f *float64) *ProjectCreate {
	if f != nil {
		pc.SetBudget(*f)
	}
	return pc
}

// SetDeleted sets the "deleted" field.
func (pc *ProjectCreate) SetDeleted(b bool) *ProjectCreate {
	pc.mutation.SetDeleted(b)
	return pc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableDeleted(b *bool) *ProjectCreate {
	if b != nil {
		pc.SetDeleted(*b)
	}
	return pc
}

// SetStartDate sets the "start_date" field.
func (pc *ProjectCreate) SetStartDate(t time.Time) *ProjectCreate {
	pc.mutation.SetStartDate(t)
	return pc
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableStartDate(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetStartDate(*t)
	}
	return pc
}

// SetEndDate sets the "end_date" field.
func (pc *ProjectCreate) SetEndDate(t time.Time) *ProjectCreate {
	pc.mutation.SetEndDate(t)
	return pc
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableEndDate(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetEndDate(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProjectCreate) SetUpdatedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUpdatedAt(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProjectCreate) SetCreatedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableCreatedAt(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// Mutation returns the ProjectMutation object of the builder.
func (pc *ProjectCreate) Mutation() *ProjectMutation {
	return pc.mutation
}

// Save creates the Project in the database.
func (pc *ProjectCreate) Save(ctx context.Context) (*Project, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProjectCreate) SaveX(ctx context.Context) *Project {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProjectCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProjectCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProjectCreate) defaults() {
	if _, ok := pc.mutation.UUID(); !ok {
		v := project.DefaultUUID()
		pc.mutation.SetUUID(v)
	}
	if _, ok := pc.mutation.Status(); !ok {
		v := project.DefaultStatus
		pc.mutation.SetStatus(v)
	}
	if _, ok := pc.mutation.Deleted(); !ok {
		v := project.DefaultDeleted
		pc.mutation.SetDeleted(v)
	}
	if _, ok := pc.mutation.StartDate(); !ok {
		v := project.DefaultStartDate()
		pc.mutation.SetStartDate(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := project.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := project.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProjectCreate) check() error {
	if _, ok := pc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "Project.uuid"`)}
	}
	if _, ok := pc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`ent: missing required field "Project.account_id"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Project.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Project.name": %w`, err)}
		}
	}
	if v, ok := pc.mutation.Status(); ok {
		if err := project.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Project.status": %w`, err)}
		}
	}
	if v, ok := pc.mutation.Budget(); ok {
		if err := project.BudgetValidator(v); err != nil {
			return &ValidationError{Name: "budget", err: fmt.Errorf(`ent: validator failed for field "Project.budget": %w`, err)}
		}
	}
	return nil
}

func (pc *ProjectCreate) sqlSave(ctx context.Context) (*Project, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProjectCreate) createSpec() (*Project, *sqlgraph.CreateSpec) {
	var (
		_node = &Project{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(project.Table, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.UUID(); ok {
		_spec.SetField(project.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if value, ok := pc.mutation.AccountID(); ok {
		_spec.SetField(project.FieldAccountID, field.TypeString, value)
		_node.AccountID = value
	}
	if value, ok := pc.mutation.ClientID(); ok {
		_spec.SetField(project.FieldClientID, field.TypeString, value)
		_node.ClientID = &value
	}
	if value, ok := pc.mutation.ManagerID(); ok {
		_spec.SetField(project.FieldManagerID, field.TypeString, value)
		_node.ManagerID = &value
	}
	if value, ok := pc.mutation.CreatedBy(); ok {
		_spec.SetField(project.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if value, ok := pc.mutation.Notes(); ok {
		_spec.SetField(project.FieldNotes, field.TypeString, value)
		_node.Notes = &value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(project.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := pc.mutation.Location(); ok {
		_spec.SetField(project.FieldLocation, field.TypeString, value)
		_node.Location = &value
	}
	if value, ok := pc.mutation.Budget(); ok {
		_spec.SetField(project.FieldBudget, field.TypeFloat64, value)
		_node.Budget = value
	}
	if value, ok := pc.mutation.Deleted(); ok {
		_spec.SetField(project.FieldDeleted, field.TypeBool, value)
		_node.Deleted = &value
	}
	if value, ok := pc.mutation.StartDate(); ok {
		_spec.SetField(project.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := pc.mutation.EndDate(); ok {
		_spec.SetField(project.FieldEndDate, field.TypeTime, value)
		_node.EndDate = &value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(project.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(project.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// ProjectCreateBulk is the builder for creating many Project entities in bulk.
type ProjectCreateBulk struct {
	config
	err      error
	builders []*ProjectCreate
}

// Save creates the Project entities in the database.
func (pcb *ProjectCreateBulk) Save(ctx context.Context) ([]*Project, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Project, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProjectCreateBulk) SaveX(ctx context.Context) []*Project {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProjectCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProjectCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
