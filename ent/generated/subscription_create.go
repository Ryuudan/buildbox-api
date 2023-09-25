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
	"github.com/Pyakz/buildbox-api/ent/generated/plan"
	"github.com/Pyakz/buildbox-api/ent/generated/subscription"
	"github.com/google/uuid"
)

// SubscriptionCreate is the builder for creating a Subscription entity.
type SubscriptionCreate struct {
	config
	mutation *SubscriptionMutation
	hooks    []Hook
}

// SetAccountID sets the "account_id" field.
func (sc *SubscriptionCreate) SetAccountID(i int) *SubscriptionCreate {
	sc.mutation.SetAccountID(i)
	return sc
}

// SetPlanID sets the "plan_id" field.
func (sc *SubscriptionCreate) SetPlanID(i int) *SubscriptionCreate {
	sc.mutation.SetPlanID(i)
	return sc
}

// SetStartDate sets the "start_date" field.
func (sc *SubscriptionCreate) SetStartDate(t time.Time) *SubscriptionCreate {
	sc.mutation.SetStartDate(t)
	return sc
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (sc *SubscriptionCreate) SetNillableStartDate(t *time.Time) *SubscriptionCreate {
	if t != nil {
		sc.SetStartDate(*t)
	}
	return sc
}

// SetEndDate sets the "end_date" field.
func (sc *SubscriptionCreate) SetEndDate(t time.Time) *SubscriptionCreate {
	sc.mutation.SetEndDate(t)
	return sc
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (sc *SubscriptionCreate) SetNillableEndDate(t *time.Time) *SubscriptionCreate {
	if t != nil {
		sc.SetEndDate(*t)
	}
	return sc
}

// SetStatus sets the "status" field.
func (sc *SubscriptionCreate) SetStatus(s subscription.Status) *SubscriptionCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *SubscriptionCreate) SetNillableStatus(s *subscription.Status) *SubscriptionCreate {
	if s != nil {
		sc.SetStatus(*s)
	}
	return sc
}

// SetDiscount sets the "discount" field.
func (sc *SubscriptionCreate) SetDiscount(f float64) *SubscriptionCreate {
	sc.mutation.SetDiscount(f)
	return sc
}

// SetNillableDiscount sets the "discount" field if the given value is not nil.
func (sc *SubscriptionCreate) SetNillableDiscount(f *float64) *SubscriptionCreate {
	if f != nil {
		sc.SetDiscount(*f)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SubscriptionCreate) SetUpdatedAt(t time.Time) *SubscriptionCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SubscriptionCreate) SetNillableUpdatedAt(t *time.Time) *SubscriptionCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SubscriptionCreate) SetCreatedAt(t time.Time) *SubscriptionCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SubscriptionCreate) SetNillableCreatedAt(t *time.Time) *SubscriptionCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUUID sets the "uuid" field.
func (sc *SubscriptionCreate) SetUUID(u uuid.UUID) *SubscriptionCreate {
	sc.mutation.SetUUID(u)
	return sc
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (sc *SubscriptionCreate) SetNillableUUID(u *uuid.UUID) *SubscriptionCreate {
	if u != nil {
		sc.SetUUID(*u)
	}
	return sc
}

// SetAccount sets the "account" edge to the Account entity.
func (sc *SubscriptionCreate) SetAccount(a *Account) *SubscriptionCreate {
	return sc.SetAccountID(a.ID)
}

// SetPlan sets the "plan" edge to the Plan entity.
func (sc *SubscriptionCreate) SetPlan(p *Plan) *SubscriptionCreate {
	return sc.SetPlanID(p.ID)
}

// Mutation returns the SubscriptionMutation object of the builder.
func (sc *SubscriptionCreate) Mutation() *SubscriptionMutation {
	return sc.mutation
}

// Save creates the Subscription in the database.
func (sc *SubscriptionCreate) Save(ctx context.Context) (*Subscription, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubscriptionCreate) SaveX(ctx context.Context) *Subscription {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubscriptionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubscriptionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SubscriptionCreate) defaults() {
	if _, ok := sc.mutation.StartDate(); !ok {
		v := subscription.DefaultStartDate()
		sc.mutation.SetStartDate(v)
	}
	if _, ok := sc.mutation.Status(); !ok {
		v := subscription.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := subscription.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := subscription.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UUID(); !ok {
		v := subscription.DefaultUUID()
		sc.mutation.SetUUID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubscriptionCreate) check() error {
	if _, ok := sc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`generated: missing required field "Subscription.account_id"`)}
	}
	if _, ok := sc.mutation.PlanID(); !ok {
		return &ValidationError{Name: "plan_id", err: errors.New(`generated: missing required field "Subscription.plan_id"`)}
	}
	if _, ok := sc.mutation.StartDate(); !ok {
		return &ValidationError{Name: "start_date", err: errors.New(`generated: missing required field "Subscription.start_date"`)}
	}
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`generated: missing required field "Subscription.status"`)}
	}
	if v, ok := sc.mutation.Status(); ok {
		if err := subscription.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "Subscription.status": %w`, err)}
		}
	}
	if _, ok := sc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`generated: missing required field "Subscription.uuid"`)}
	}
	if _, ok := sc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`generated: missing required edge "Subscription.account"`)}
	}
	if _, ok := sc.mutation.PlanID(); !ok {
		return &ValidationError{Name: "plan", err: errors.New(`generated: missing required edge "Subscription.plan"`)}
	}
	return nil
}

func (sc *SubscriptionCreate) sqlSave(ctx context.Context) (*Subscription, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubscriptionCreate) createSpec() (*Subscription, *sqlgraph.CreateSpec) {
	var (
		_node = &Subscription{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(subscription.Table, sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.StartDate(); ok {
		_spec.SetField(subscription.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := sc.mutation.EndDate(); ok {
		_spec.SetField(subscription.FieldEndDate, field.TypeTime, value)
		_node.EndDate = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(subscription.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.Discount(); ok {
		_spec.SetField(subscription.FieldDiscount, field.TypeFloat64, value)
		_node.Discount = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(subscription.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(subscription.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UUID(); ok {
		_spec.SetField(subscription.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if nodes := sc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.AccountTable,
			Columns: []string{subscription.AccountColumn},
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
	if nodes := sc.mutation.PlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.PlanTable,
			Columns: []string{subscription.PlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(plan.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PlanID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubscriptionCreateBulk is the builder for creating many Subscription entities in bulk.
type SubscriptionCreateBulk struct {
	config
	err      error
	builders []*SubscriptionCreate
}

// Save creates the Subscription entities in the database.
func (scb *SubscriptionCreateBulk) Save(ctx context.Context) ([]*Subscription, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subscription, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubscriptionMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubscriptionCreateBulk) SaveX(ctx context.Context) []*Subscription {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubscriptionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubscriptionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
