// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Pyakz/buildbox-api/ent/generated/account"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/role"
	"github.com/Pyakz/buildbox-api/ent/generated/user"
)

// RoleQuery is the builder for querying Role entities.
type RoleQuery struct {
	config
	ctx         *QueryContext
	order       []role.OrderOption
	inters      []Interceptor
	predicates  []predicate.Role
	withAccount *AccountQuery
	withUser    *UserQuery
	withUsers   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoleQuery builder.
func (rq *RoleQuery) Where(ps ...predicate.Role) *RoleQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RoleQuery) Limit(limit int) *RoleQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RoleQuery) Offset(offset int) *RoleQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RoleQuery) Unique(unique bool) *RoleQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RoleQuery) Order(o ...role.OrderOption) *RoleQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryAccount chains the current query on the "account" edge.
func (rq *RoleQuery) QueryAccount() *AccountQuery {
	query := (&AccountClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, role.AccountTable, role.AccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (rq *RoleQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, role.UserTable, role.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the "users" edge.
func (rq *RoleQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, role.UsersTable, role.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Role entity from the query.
// Returns a *NotFoundError when no Role was found.
func (rq *RoleQuery) First(ctx context.Context) (*Role, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{role.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RoleQuery) FirstX(ctx context.Context) *Role {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Role ID from the query.
// Returns a *NotFoundError when no Role ID was found.
func (rq *RoleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{role.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RoleQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Role entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Role entity is found.
// Returns a *NotFoundError when no Role entities are found.
func (rq *RoleQuery) Only(ctx context.Context) (*Role, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{role.Label}
	default:
		return nil, &NotSingularError{role.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RoleQuery) OnlyX(ctx context.Context) *Role {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Role ID in the query.
// Returns a *NotSingularError when more than one Role ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RoleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{role.Label}
	default:
		err = &NotSingularError{role.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RoleQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Roles.
func (rq *RoleQuery) All(ctx context.Context) ([]*Role, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Role, *RoleQuery]()
	return withInterceptors[[]*Role](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RoleQuery) AllX(ctx context.Context) []*Role {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Role IDs.
func (rq *RoleQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(role.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RoleQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RoleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RoleQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RoleQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RoleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RoleQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RoleQuery) Clone() *RoleQuery {
	if rq == nil {
		return nil
	}
	return &RoleQuery{
		config:      rq.config,
		ctx:         rq.ctx.Clone(),
		order:       append([]role.OrderOption{}, rq.order...),
		inters:      append([]Interceptor{}, rq.inters...),
		predicates:  append([]predicate.Role{}, rq.predicates...),
		withAccount: rq.withAccount.Clone(),
		withUser:    rq.withUser.Clone(),
		withUsers:   rq.withUsers.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithAccount tells the query-builder to eager-load the nodes that are connected to
// the "account" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithAccount(opts ...func(*AccountQuery)) *RoleQuery {
	query := (&AccountClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withAccount = query
	return rq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithUser(opts ...func(*UserQuery)) *RoleQuery {
	query := (&UserClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withUser = query
	return rq
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithUsers(opts ...func(*UserQuery)) *RoleQuery {
	query := (&UserClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withUsers = query
	return rq
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
//	client.Role.Query().
//		GroupBy(role.FieldAccountID).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (rq *RoleQuery) GroupBy(field string, fields ...string) *RoleGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RoleGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = role.Label
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
//	client.Role.Query().
//		Select(role.FieldAccountID).
//		Scan(ctx, &v)
func (rq *RoleQuery) Select(fields ...string) *RoleSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RoleSelect{RoleQuery: rq}
	sbuild.label = role.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RoleSelect configured with the given aggregations.
func (rq *RoleQuery) Aggregate(fns ...AggregateFunc) *RoleSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RoleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !role.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RoleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Role, error) {
	var (
		nodes       = []*Role{}
		_spec       = rq.querySpec()
		loadedTypes = [3]bool{
			rq.withAccount != nil,
			rq.withUser != nil,
			rq.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Role).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Role{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withAccount; query != nil {
		if err := rq.loadAccount(ctx, query, nodes, nil,
			func(n *Role, e *Account) { n.Edges.Account = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withUser; query != nil {
		if err := rq.loadUser(ctx, query, nodes, nil,
			func(n *Role, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withUsers; query != nil {
		if err := rq.loadUsers(ctx, query, nodes,
			func(n *Role) { n.Edges.Users = []*User{} },
			func(n *Role, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RoleQuery) loadAccount(ctx context.Context, query *AccountQuery, nodes []*Role, init func(*Role), assign func(*Role, *Account)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Role)
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
func (rq *RoleQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Role, init func(*Role), assign func(*Role, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Role)
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
func (rq *RoleQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Role, init func(*Role), assign func(*Role, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Role)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(user.FieldRoleID)
	}
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(role.UsersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RoleID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "role_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rq *RoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for i := range fields {
			if fields[i] != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if rq.withAccount != nil {
			_spec.Node.AddColumnOnce(role.FieldAccountID)
		}
		if rq.withUser != nil {
			_spec.Node.AddColumnOnce(role.FieldCreatedBy)
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(role.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = role.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoleGroupBy is the group-by builder for Role entities.
type RoleGroupBy struct {
	selector
	build *RoleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RoleGroupBy) Aggregate(fns ...AggregateFunc) *RoleGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RoleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoleQuery, *RoleGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RoleGroupBy) sqlScan(ctx context.Context, root *RoleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RoleSelect is the builder for selecting fields of Role entities.
type RoleSelect struct {
	*RoleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RoleSelect) Aggregate(fns ...AggregateFunc) *RoleSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RoleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoleQuery, *RoleSelect](ctx, rs.RoleQuery, rs, rs.inters, v)
}

func (rs *RoleSelect) sqlScan(ctx context.Context, root *RoleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
