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
	"github.com/Pyakz/buildbox-api/ent/generated/projectserviceprovider"
	"github.com/Pyakz/buildbox-api/ent/generated/serviceprovider"
	"github.com/Pyakz/buildbox-api/ent/generated/user"
	"github.com/google/uuid"
)

// ServiceProviderCreate is the builder for creating a ServiceProvider entity.
type ServiceProviderCreate struct {
	config
	mutation *ServiceProviderMutation
	hooks    []Hook
}

// SetAccountID sets the "account_id" field.
func (spc *ServiceProviderCreate) SetAccountID(i int) *ServiceProviderCreate {
	spc.mutation.SetAccountID(i)
	return spc
}

// SetCreatedBy sets the "created_by" field.
func (spc *ServiceProviderCreate) SetCreatedBy(i int) *ServiceProviderCreate {
	spc.mutation.SetCreatedBy(i)
	return spc
}

// SetName sets the "name" field.
func (spc *ServiceProviderCreate) SetName(s string) *ServiceProviderCreate {
	spc.mutation.SetName(s)
	return spc
}

// SetEmail sets the "email" field.
func (spc *ServiceProviderCreate) SetEmail(s string) *ServiceProviderCreate {
	spc.mutation.SetEmail(s)
	return spc
}

// SetDescription sets the "description" field.
func (spc *ServiceProviderCreate) SetDescription(s string) *ServiceProviderCreate {
	spc.mutation.SetDescription(s)
	return spc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (spc *ServiceProviderCreate) SetNillableDescription(s *string) *ServiceProviderCreate {
	if s != nil {
		spc.SetDescription(*s)
	}
	return spc
}

// SetStatus sets the "status" field.
func (spc *ServiceProviderCreate) SetStatus(s serviceprovider.Status) *ServiceProviderCreate {
	spc.mutation.SetStatus(s)
	return spc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (spc *ServiceProviderCreate) SetNillableStatus(s *serviceprovider.Status) *ServiceProviderCreate {
	if s != nil {
		spc.SetStatus(*s)
	}
	return spc
}

// SetPhoneNumber sets the "phone_number" field.
func (spc *ServiceProviderCreate) SetPhoneNumber(s string) *ServiceProviderCreate {
	spc.mutation.SetPhoneNumber(s)
	return spc
}

// SetUpdatedAt sets the "updated_at" field.
func (spc *ServiceProviderCreate) SetUpdatedAt(t time.Time) *ServiceProviderCreate {
	spc.mutation.SetUpdatedAt(t)
	return spc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (spc *ServiceProviderCreate) SetNillableUpdatedAt(t *time.Time) *ServiceProviderCreate {
	if t != nil {
		spc.SetUpdatedAt(*t)
	}
	return spc
}

// SetCreatedAt sets the "created_at" field.
func (spc *ServiceProviderCreate) SetCreatedAt(t time.Time) *ServiceProviderCreate {
	spc.mutation.SetCreatedAt(t)
	return spc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (spc *ServiceProviderCreate) SetNillableCreatedAt(t *time.Time) *ServiceProviderCreate {
	if t != nil {
		spc.SetCreatedAt(*t)
	}
	return spc
}

// SetUUID sets the "uuid" field.
func (spc *ServiceProviderCreate) SetUUID(u uuid.UUID) *ServiceProviderCreate {
	spc.mutation.SetUUID(u)
	return spc
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (spc *ServiceProviderCreate) SetNillableUUID(u *uuid.UUID) *ServiceProviderCreate {
	if u != nil {
		spc.SetUUID(*u)
	}
	return spc
}

// AddServiceProviderProjectIDs adds the "service_provider_projects" edge to the ProjectServiceProvider entity by IDs.
func (spc *ServiceProviderCreate) AddServiceProviderProjectIDs(ids ...int) *ServiceProviderCreate {
	spc.mutation.AddServiceProviderProjectIDs(ids...)
	return spc
}

// AddServiceProviderProjects adds the "service_provider_projects" edges to the ProjectServiceProvider entity.
func (spc *ServiceProviderCreate) AddServiceProviderProjects(p ...*ProjectServiceProvider) *ServiceProviderCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return spc.AddServiceProviderProjectIDs(ids...)
}

// SetAccount sets the "account" edge to the Account entity.
func (spc *ServiceProviderCreate) SetAccount(a *Account) *ServiceProviderCreate {
	return spc.SetAccountID(a.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (spc *ServiceProviderCreate) SetUserID(id int) *ServiceProviderCreate {
	spc.mutation.SetUserID(id)
	return spc
}

// SetUser sets the "user" edge to the User entity.
func (spc *ServiceProviderCreate) SetUser(u *User) *ServiceProviderCreate {
	return spc.SetUserID(u.ID)
}

// Mutation returns the ServiceProviderMutation object of the builder.
func (spc *ServiceProviderCreate) Mutation() *ServiceProviderMutation {
	return spc.mutation
}

// Save creates the ServiceProvider in the database.
func (spc *ServiceProviderCreate) Save(ctx context.Context) (*ServiceProvider, error) {
	spc.defaults()
	return withHooks(ctx, spc.sqlSave, spc.mutation, spc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (spc *ServiceProviderCreate) SaveX(ctx context.Context) *ServiceProvider {
	v, err := spc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spc *ServiceProviderCreate) Exec(ctx context.Context) error {
	_, err := spc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spc *ServiceProviderCreate) ExecX(ctx context.Context) {
	if err := spc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spc *ServiceProviderCreate) defaults() {
	if _, ok := spc.mutation.Status(); !ok {
		v := serviceprovider.DefaultStatus
		spc.mutation.SetStatus(v)
	}
	if _, ok := spc.mutation.UpdatedAt(); !ok {
		v := serviceprovider.DefaultUpdatedAt()
		spc.mutation.SetUpdatedAt(v)
	}
	if _, ok := spc.mutation.CreatedAt(); !ok {
		v := serviceprovider.DefaultCreatedAt()
		spc.mutation.SetCreatedAt(v)
	}
	if _, ok := spc.mutation.UUID(); !ok {
		v := serviceprovider.DefaultUUID()
		spc.mutation.SetUUID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spc *ServiceProviderCreate) check() error {
	if _, ok := spc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`generated: missing required field "ServiceProvider.account_id"`)}
	}
	if _, ok := spc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`generated: missing required field "ServiceProvider.created_by"`)}
	}
	if _, ok := spc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "ServiceProvider.name"`)}
	}
	if v, ok := spc.mutation.Name(); ok {
		if err := serviceprovider.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "ServiceProvider.name": %w`, err)}
		}
	}
	if _, ok := spc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`generated: missing required field "ServiceProvider.email"`)}
	}
	if v, ok := spc.mutation.Status(); ok {
		if err := serviceprovider.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "ServiceProvider.status": %w`, err)}
		}
	}
	if _, ok := spc.mutation.PhoneNumber(); !ok {
		return &ValidationError{Name: "phone_number", err: errors.New(`generated: missing required field "ServiceProvider.phone_number"`)}
	}
	if _, ok := spc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`generated: missing required field "ServiceProvider.uuid"`)}
	}
	if _, ok := spc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`generated: missing required edge "ServiceProvider.account"`)}
	}
	if _, ok := spc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`generated: missing required edge "ServiceProvider.user"`)}
	}
	return nil
}

func (spc *ServiceProviderCreate) sqlSave(ctx context.Context) (*ServiceProvider, error) {
	if err := spc.check(); err != nil {
		return nil, err
	}
	_node, _spec := spc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	spc.mutation.id = &_node.ID
	spc.mutation.done = true
	return _node, nil
}

func (spc *ServiceProviderCreate) createSpec() (*ServiceProvider, *sqlgraph.CreateSpec) {
	var (
		_node = &ServiceProvider{config: spc.config}
		_spec = sqlgraph.NewCreateSpec(serviceprovider.Table, sqlgraph.NewFieldSpec(serviceprovider.FieldID, field.TypeInt))
	)
	if value, ok := spc.mutation.Name(); ok {
		_spec.SetField(serviceprovider.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := spc.mutation.Email(); ok {
		_spec.SetField(serviceprovider.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := spc.mutation.Description(); ok {
		_spec.SetField(serviceprovider.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if value, ok := spc.mutation.Status(); ok {
		_spec.SetField(serviceprovider.FieldStatus, field.TypeEnum, value)
		_node.Status = &value
	}
	if value, ok := spc.mutation.PhoneNumber(); ok {
		_spec.SetField(serviceprovider.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if value, ok := spc.mutation.UpdatedAt(); ok {
		_spec.SetField(serviceprovider.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := spc.mutation.CreatedAt(); ok {
		_spec.SetField(serviceprovider.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := spc.mutation.UUID(); ok {
		_spec.SetField(serviceprovider.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if nodes := spc.mutation.ServiceProviderProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceprovider.ServiceProviderProjectsTable,
			Columns: []string{serviceprovider.ServiceProviderProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectserviceprovider.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := spc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceprovider.AccountTable,
			Columns: []string{serviceprovider.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AccountID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := spc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceprovider.UserTable,
			Columns: []string{serviceprovider.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CreatedBy = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ServiceProviderCreateBulk is the builder for creating many ServiceProvider entities in bulk.
type ServiceProviderCreateBulk struct {
	config
	err      error
	builders []*ServiceProviderCreate
}

// Save creates the ServiceProvider entities in the database.
func (spcb *ServiceProviderCreateBulk) Save(ctx context.Context) ([]*ServiceProvider, error) {
	if spcb.err != nil {
		return nil, spcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(spcb.builders))
	nodes := make([]*ServiceProvider, len(spcb.builders))
	mutators := make([]Mutator, len(spcb.builders))
	for i := range spcb.builders {
		func(i int, root context.Context) {
			builder := spcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServiceProviderMutation)
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
					_, err = mutators[i+1].Mutate(root, spcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, spcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spcb *ServiceProviderCreateBulk) SaveX(ctx context.Context) []*ServiceProvider {
	v, err := spcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spcb *ServiceProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := spcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spcb *ServiceProviderCreateBulk) ExecX(ctx context.Context) {
	if err := spcb.Exec(ctx); err != nil {
		panic(err)
	}
}
