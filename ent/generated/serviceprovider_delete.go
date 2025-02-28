// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/serviceprovider"
)

// ServiceProviderDelete is the builder for deleting a ServiceProvider entity.
type ServiceProviderDelete struct {
	config
	hooks    []Hook
	mutation *ServiceProviderMutation
}

// Where appends a list predicates to the ServiceProviderDelete builder.
func (spd *ServiceProviderDelete) Where(ps ...predicate.ServiceProvider) *ServiceProviderDelete {
	spd.mutation.Where(ps...)
	return spd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (spd *ServiceProviderDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, spd.sqlExec, spd.mutation, spd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (spd *ServiceProviderDelete) ExecX(ctx context.Context) int {
	n, err := spd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (spd *ServiceProviderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(serviceprovider.Table, sqlgraph.NewFieldSpec(serviceprovider.FieldID, field.TypeInt))
	if ps := spd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, spd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	spd.mutation.done = true
	return affected, err
}

// ServiceProviderDeleteOne is the builder for deleting a single ServiceProvider entity.
type ServiceProviderDeleteOne struct {
	spd *ServiceProviderDelete
}

// Where appends a list predicates to the ServiceProviderDelete builder.
func (spdo *ServiceProviderDeleteOne) Where(ps ...predicate.ServiceProvider) *ServiceProviderDeleteOne {
	spdo.spd.mutation.Where(ps...)
	return spdo
}

// Exec executes the deletion query.
func (spdo *ServiceProviderDeleteOne) Exec(ctx context.Context) error {
	n, err := spdo.spd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{serviceprovider.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (spdo *ServiceProviderDeleteOne) ExecX(ctx context.Context) {
	if err := spdo.Exec(ctx); err != nil {
		panic(err)
	}
}
