// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Pyakz/buildbox-api/ent/generated/account"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/serviceprovider"
	"github.com/Pyakz/buildbox-api/ent/generated/user"
)

// ServiceProviderQuery is the builder for querying ServiceProvider entities.
type ServiceProviderQuery struct {
	config
	ctx         *QueryContext
	order       []serviceprovider.OrderOption
	inters      []Interceptor
	predicates  []predicate.ServiceProvider
	withAccount *AccountQuery
	withUser    *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ServiceProviderQuery builder.
func (spq *ServiceProviderQuery) Where(ps ...predicate.ServiceProvider) *ServiceProviderQuery {
	spq.predicates = append(spq.predicates, ps...)
	return spq
}

// Limit the number of records to be returned by this query.
func (spq *ServiceProviderQuery) Limit(limit int) *ServiceProviderQuery {
	spq.ctx.Limit = &limit
	return spq
}

// Offset to start from.
func (spq *ServiceProviderQuery) Offset(offset int) *ServiceProviderQuery {
	spq.ctx.Offset = &offset
	return spq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spq *ServiceProviderQuery) Unique(unique bool) *ServiceProviderQuery {
	spq.ctx.Unique = &unique
	return spq
}

// Order specifies how the records should be ordered.
func (spq *ServiceProviderQuery) Order(o ...serviceprovider.OrderOption) *ServiceProviderQuery {
	spq.order = append(spq.order, o...)
	return spq
}

// QueryAccount chains the current query on the "account" edge.
func (spq *ServiceProviderQuery) QueryAccount() *AccountQuery {
	query := (&AccountClient{config: spq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(serviceprovider.Table, serviceprovider.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, serviceprovider.AccountTable, serviceprovider.AccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (spq *ServiceProviderQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: spq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(serviceprovider.Table, serviceprovider.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, serviceprovider.UserTable, serviceprovider.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ServiceProvider entity from the query.
// Returns a *NotFoundError when no ServiceProvider was found.
func (spq *ServiceProviderQuery) First(ctx context.Context) (*ServiceProvider, error) {
	nodes, err := spq.Limit(1).All(setContextOp(ctx, spq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{serviceprovider.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spq *ServiceProviderQuery) FirstX(ctx context.Context) *ServiceProvider {
	node, err := spq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ServiceProvider ID from the query.
// Returns a *NotFoundError when no ServiceProvider ID was found.
func (spq *ServiceProviderQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = spq.Limit(1).IDs(setContextOp(ctx, spq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{serviceprovider.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spq *ServiceProviderQuery) FirstIDX(ctx context.Context) int {
	id, err := spq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ServiceProvider entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ServiceProvider entity is found.
// Returns a *NotFoundError when no ServiceProvider entities are found.
func (spq *ServiceProviderQuery) Only(ctx context.Context) (*ServiceProvider, error) {
	nodes, err := spq.Limit(2).All(setContextOp(ctx, spq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{serviceprovider.Label}
	default:
		return nil, &NotSingularError{serviceprovider.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spq *ServiceProviderQuery) OnlyX(ctx context.Context) *ServiceProvider {
	node, err := spq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ServiceProvider ID in the query.
// Returns a *NotSingularError when more than one ServiceProvider ID is found.
// Returns a *NotFoundError when no entities are found.
func (spq *ServiceProviderQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = spq.Limit(2).IDs(setContextOp(ctx, spq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{serviceprovider.Label}
	default:
		err = &NotSingularError{serviceprovider.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spq *ServiceProviderQuery) OnlyIDX(ctx context.Context) int {
	id, err := spq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ServiceProviders.
func (spq *ServiceProviderQuery) All(ctx context.Context) ([]*ServiceProvider, error) {
	ctx = setContextOp(ctx, spq.ctx, "All")
	if err := spq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ServiceProvider, *ServiceProviderQuery]()
	return withInterceptors[[]*ServiceProvider](ctx, spq, qr, spq.inters)
}

// AllX is like All, but panics if an error occurs.
func (spq *ServiceProviderQuery) AllX(ctx context.Context) []*ServiceProvider {
	nodes, err := spq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ServiceProvider IDs.
func (spq *ServiceProviderQuery) IDs(ctx context.Context) (ids []int, err error) {
	if spq.ctx.Unique == nil && spq.path != nil {
		spq.Unique(true)
	}
	ctx = setContextOp(ctx, spq.ctx, "IDs")
	if err = spq.Select(serviceprovider.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spq *ServiceProviderQuery) IDsX(ctx context.Context) []int {
	ids, err := spq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spq *ServiceProviderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, spq.ctx, "Count")
	if err := spq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, spq, querierCount[*ServiceProviderQuery](), spq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (spq *ServiceProviderQuery) CountX(ctx context.Context) int {
	count, err := spq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spq *ServiceProviderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, spq.ctx, "Exist")
	switch _, err := spq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (spq *ServiceProviderQuery) ExistX(ctx context.Context) bool {
	exist, err := spq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ServiceProviderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spq *ServiceProviderQuery) Clone() *ServiceProviderQuery {
	if spq == nil {
		return nil
	}
	return &ServiceProviderQuery{
		config:      spq.config,
		ctx:         spq.ctx.Clone(),
		order:       append([]serviceprovider.OrderOption{}, spq.order...),
		inters:      append([]Interceptor{}, spq.inters...),
		predicates:  append([]predicate.ServiceProvider{}, spq.predicates...),
		withAccount: spq.withAccount.Clone(),
		withUser:    spq.withUser.Clone(),
		// clone intermediate query.
		sql:  spq.sql.Clone(),
		path: spq.path,
	}
}

// WithAccount tells the query-builder to eager-load the nodes that are connected to
// the "account" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *ServiceProviderQuery) WithAccount(opts ...func(*AccountQuery)) *ServiceProviderQuery {
	query := (&AccountClient{config: spq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	spq.withAccount = query
	return spq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *ServiceProviderQuery) WithUser(opts ...func(*UserQuery)) *ServiceProviderQuery {
	query := (&UserClient{config: spq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	spq.withUser = query
	return spq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AccountID int `json:"account_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ServiceProvider.Query().
//		GroupBy(serviceprovider.FieldAccountID).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (spq *ServiceProviderQuery) GroupBy(field string, fields ...string) *ServiceProviderGroupBy {
	spq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ServiceProviderGroupBy{build: spq}
	grbuild.flds = &spq.ctx.Fields
	grbuild.label = serviceprovider.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AccountID int `json:"account_id,omitempty"`
//	}
//
//	client.ServiceProvider.Query().
//		Select(serviceprovider.FieldAccountID).
//		Scan(ctx, &v)
func (spq *ServiceProviderQuery) Select(fields ...string) *ServiceProviderSelect {
	spq.ctx.Fields = append(spq.ctx.Fields, fields...)
	sbuild := &ServiceProviderSelect{ServiceProviderQuery: spq}
	sbuild.label = serviceprovider.Label
	sbuild.flds, sbuild.scan = &spq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ServiceProviderSelect configured with the given aggregations.
func (spq *ServiceProviderQuery) Aggregate(fns ...AggregateFunc) *ServiceProviderSelect {
	return spq.Select().Aggregate(fns...)
}

func (spq *ServiceProviderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range spq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, spq); err != nil {
				return err
			}
		}
	}
	for _, f := range spq.ctx.Fields {
		if !serviceprovider.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if spq.path != nil {
		prev, err := spq.path(ctx)
		if err != nil {
			return err
		}
		spq.sql = prev
	}
	return nil
}

func (spq *ServiceProviderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ServiceProvider, error) {
	var (
		nodes       = []*ServiceProvider{}
		_spec       = spq.querySpec()
		loadedTypes = [2]bool{
			spq.withAccount != nil,
			spq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ServiceProvider).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ServiceProvider{config: spq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, spq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := spq.withAccount; query != nil {
		if err := spq.loadAccount(ctx, query, nodes, nil,
			func(n *ServiceProvider, e *Account) { n.Edges.Account = e }); err != nil {
			return nil, err
		}
	}
	if query := spq.withUser; query != nil {
		if err := spq.loadUser(ctx, query, nodes, nil,
			func(n *ServiceProvider, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (spq *ServiceProviderQuery) loadAccount(ctx context.Context, query *AccountQuery, nodes []*ServiceProvider, init func(*ServiceProvider), assign func(*ServiceProvider, *Account)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ServiceProvider)
	for i := range nodes {
		fk := nodes[i].AccountID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(account.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (spq *ServiceProviderQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*ServiceProvider, init func(*ServiceProvider), assign func(*ServiceProvider, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ServiceProvider)
	for i := range nodes {
		fk := nodes[i].CreatedBy
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "created_by" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (spq *ServiceProviderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := spq.querySpec()
	_spec.Node.Columns = spq.ctx.Fields
	if len(spq.ctx.Fields) > 0 {
		_spec.Unique = spq.ctx.Unique != nil && *spq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, spq.driver, _spec)
}

func (spq *ServiceProviderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(serviceprovider.Table, serviceprovider.Columns, sqlgraph.NewFieldSpec(serviceprovider.FieldID, field.TypeInt))
	_spec.From = spq.sql
	if unique := spq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if spq.path != nil {
		_spec.Unique = true
	}
	if fields := spq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, serviceprovider.FieldID)
		for i := range fields {
			if fields[i] != serviceprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if spq.withAccount != nil {
			_spec.Node.AddColumnOnce(serviceprovider.FieldAccountID)
		}
		if spq.withUser != nil {
			_spec.Node.AddColumnOnce(serviceprovider.FieldCreatedBy)
		}
	}
	if ps := spq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := spq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := spq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := spq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (spq *ServiceProviderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(spq.driver.Dialect())
	t1 := builder.Table(serviceprovider.Table)
	columns := spq.ctx.Fields
	if len(columns) == 0 {
		columns = serviceprovider.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if spq.sql != nil {
		selector = spq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if spq.ctx.Unique != nil && *spq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range spq.predicates {
		p(selector)
	}
	for _, p := range spq.order {
		p(selector)
	}
	if offset := spq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := spq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ServiceProviderGroupBy is the group-by builder for ServiceProvider entities.
type ServiceProviderGroupBy struct {
	selector
	build *ServiceProviderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spgb *ServiceProviderGroupBy) Aggregate(fns ...AggregateFunc) *ServiceProviderGroupBy {
	spgb.fns = append(spgb.fns, fns...)
	return spgb
}

// Scan applies the selector query and scans the result into the given value.
func (spgb *ServiceProviderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, spgb.build.ctx, "GroupBy")
	if err := spgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ServiceProviderQuery, *ServiceProviderGroupBy](ctx, spgb.build, spgb, spgb.build.inters, v)
}

func (spgb *ServiceProviderGroupBy) sqlScan(ctx context.Context, root *ServiceProviderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(spgb.fns))
	for _, fn := range spgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*spgb.flds)+len(spgb.fns))
		for _, f := range *spgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*spgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := spgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ServiceProviderSelect is the builder for selecting fields of ServiceProvider entities.
type ServiceProviderSelect struct {
	*ServiceProviderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sps *ServiceProviderSelect) Aggregate(fns ...AggregateFunc) *ServiceProviderSelect {
	sps.fns = append(sps.fns, fns...)
	return sps
}

// Scan applies the selector query and scans the result into the given value.
func (sps *ServiceProviderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sps.ctx, "Select")
	if err := sps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ServiceProviderQuery, *ServiceProviderSelect](ctx, sps.ServiceProviderQuery, sps, sps.inters, v)
}

func (sps *ServiceProviderSelect) sqlScan(ctx context.Context, root *ServiceProviderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sps.fns))
	for _, fn := range sps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
