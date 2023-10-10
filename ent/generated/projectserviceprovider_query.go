// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/project"
	"github.com/Pyakz/buildbox-api/ent/generated/projectserviceprovider"
	"github.com/Pyakz/buildbox-api/ent/generated/serviceprovider"
)

// ProjectServiceProviderQuery is the builder for querying ProjectServiceProvider entities.
type ProjectServiceProviderQuery struct {
	config
	ctx                 *QueryContext
	order               []projectserviceprovider.OrderOption
	inters              []Interceptor
	predicates          []predicate.ProjectServiceProvider
	withProject         *ProjectQuery
	withServiceProvider *ServiceProviderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProjectServiceProviderQuery builder.
func (pspq *ProjectServiceProviderQuery) Where(ps ...predicate.ProjectServiceProvider) *ProjectServiceProviderQuery {
	pspq.predicates = append(pspq.predicates, ps...)
	return pspq
}

// Limit the number of records to be returned by this query.
func (pspq *ProjectServiceProviderQuery) Limit(limit int) *ProjectServiceProviderQuery {
	pspq.ctx.Limit = &limit
	return pspq
}

// Offset to start from.
func (pspq *ProjectServiceProviderQuery) Offset(offset int) *ProjectServiceProviderQuery {
	pspq.ctx.Offset = &offset
	return pspq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pspq *ProjectServiceProviderQuery) Unique(unique bool) *ProjectServiceProviderQuery {
	pspq.ctx.Unique = &unique
	return pspq
}

// Order specifies how the records should be ordered.
func (pspq *ProjectServiceProviderQuery) Order(o ...projectserviceprovider.OrderOption) *ProjectServiceProviderQuery {
	pspq.order = append(pspq.order, o...)
	return pspq
}

// QueryProject chains the current query on the "project" edge.
func (pspq *ProjectServiceProviderQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: pspq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pspq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pspq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projectserviceprovider.Table, projectserviceprovider.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projectserviceprovider.ProjectTable, projectserviceprovider.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(pspq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryServiceProvider chains the current query on the "service_provider" edge.
func (pspq *ProjectServiceProviderQuery) QueryServiceProvider() *ServiceProviderQuery {
	query := (&ServiceProviderClient{config: pspq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pspq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pspq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projectserviceprovider.Table, projectserviceprovider.FieldID, selector),
			sqlgraph.To(serviceprovider.Table, serviceprovider.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projectserviceprovider.ServiceProviderTable, projectserviceprovider.ServiceProviderColumn),
		)
		fromU = sqlgraph.SetNeighbors(pspq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProjectServiceProvider entity from the query.
// Returns a *NotFoundError when no ProjectServiceProvider was found.
func (pspq *ProjectServiceProviderQuery) First(ctx context.Context) (*ProjectServiceProvider, error) {
	nodes, err := pspq.Limit(1).All(setContextOp(ctx, pspq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{projectserviceprovider.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pspq *ProjectServiceProviderQuery) FirstX(ctx context.Context) *ProjectServiceProvider {
	node, err := pspq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProjectServiceProvider ID from the query.
// Returns a *NotFoundError when no ProjectServiceProvider ID was found.
func (pspq *ProjectServiceProviderQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pspq.Limit(1).IDs(setContextOp(ctx, pspq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{projectserviceprovider.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pspq *ProjectServiceProviderQuery) FirstIDX(ctx context.Context) int {
	id, err := pspq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProjectServiceProvider entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProjectServiceProvider entity is found.
// Returns a *NotFoundError when no ProjectServiceProvider entities are found.
func (pspq *ProjectServiceProviderQuery) Only(ctx context.Context) (*ProjectServiceProvider, error) {
	nodes, err := pspq.Limit(2).All(setContextOp(ctx, pspq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{projectserviceprovider.Label}
	default:
		return nil, &NotSingularError{projectserviceprovider.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pspq *ProjectServiceProviderQuery) OnlyX(ctx context.Context) *ProjectServiceProvider {
	node, err := pspq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProjectServiceProvider ID in the query.
// Returns a *NotSingularError when more than one ProjectServiceProvider ID is found.
// Returns a *NotFoundError when no entities are found.
func (pspq *ProjectServiceProviderQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pspq.Limit(2).IDs(setContextOp(ctx, pspq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{projectserviceprovider.Label}
	default:
		err = &NotSingularError{projectserviceprovider.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pspq *ProjectServiceProviderQuery) OnlyIDX(ctx context.Context) int {
	id, err := pspq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProjectServiceProviders.
func (pspq *ProjectServiceProviderQuery) All(ctx context.Context) ([]*ProjectServiceProvider, error) {
	ctx = setContextOp(ctx, pspq.ctx, "All")
	if err := pspq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProjectServiceProvider, *ProjectServiceProviderQuery]()
	return withInterceptors[[]*ProjectServiceProvider](ctx, pspq, qr, pspq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pspq *ProjectServiceProviderQuery) AllX(ctx context.Context) []*ProjectServiceProvider {
	nodes, err := pspq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProjectServiceProvider IDs.
func (pspq *ProjectServiceProviderQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pspq.ctx.Unique == nil && pspq.path != nil {
		pspq.Unique(true)
	}
	ctx = setContextOp(ctx, pspq.ctx, "IDs")
	if err = pspq.Select(projectserviceprovider.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pspq *ProjectServiceProviderQuery) IDsX(ctx context.Context) []int {
	ids, err := pspq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pspq *ProjectServiceProviderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pspq.ctx, "Count")
	if err := pspq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pspq, querierCount[*ProjectServiceProviderQuery](), pspq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pspq *ProjectServiceProviderQuery) CountX(ctx context.Context) int {
	count, err := pspq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pspq *ProjectServiceProviderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pspq.ctx, "Exist")
	switch _, err := pspq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pspq *ProjectServiceProviderQuery) ExistX(ctx context.Context) bool {
	exist, err := pspq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProjectServiceProviderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pspq *ProjectServiceProviderQuery) Clone() *ProjectServiceProviderQuery {
	if pspq == nil {
		return nil
	}
	return &ProjectServiceProviderQuery{
		config:              pspq.config,
		ctx:                 pspq.ctx.Clone(),
		order:               append([]projectserviceprovider.OrderOption{}, pspq.order...),
		inters:              append([]Interceptor{}, pspq.inters...),
		predicates:          append([]predicate.ProjectServiceProvider{}, pspq.predicates...),
		withProject:         pspq.withProject.Clone(),
		withServiceProvider: pspq.withServiceProvider.Clone(),
		// clone intermediate query.
		sql:  pspq.sql.Clone(),
		path: pspq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (pspq *ProjectServiceProviderQuery) WithProject(opts ...func(*ProjectQuery)) *ProjectServiceProviderQuery {
	query := (&ProjectClient{config: pspq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pspq.withProject = query
	return pspq
}

// WithServiceProvider tells the query-builder to eager-load the nodes that are connected to
// the "service_provider" edge. The optional arguments are used to configure the query builder of the edge.
func (pspq *ProjectServiceProviderQuery) WithServiceProvider(opts ...func(*ServiceProviderQuery)) *ProjectServiceProviderQuery {
	query := (&ServiceProviderClient{config: pspq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pspq.withServiceProvider = query
	return pspq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int `json:"created_by,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProjectServiceProvider.Query().
//		GroupBy(projectserviceprovider.FieldCreatedBy).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (pspq *ProjectServiceProviderQuery) GroupBy(field string, fields ...string) *ProjectServiceProviderGroupBy {
	pspq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProjectServiceProviderGroupBy{build: pspq}
	grbuild.flds = &pspq.ctx.Fields
	grbuild.label = projectserviceprovider.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int `json:"created_by,omitempty"`
//	}
//
//	client.ProjectServiceProvider.Query().
//		Select(projectserviceprovider.FieldCreatedBy).
//		Scan(ctx, &v)
func (pspq *ProjectServiceProviderQuery) Select(fields ...string) *ProjectServiceProviderSelect {
	pspq.ctx.Fields = append(pspq.ctx.Fields, fields...)
	sbuild := &ProjectServiceProviderSelect{ProjectServiceProviderQuery: pspq}
	sbuild.label = projectserviceprovider.Label
	sbuild.flds, sbuild.scan = &pspq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProjectServiceProviderSelect configured with the given aggregations.
func (pspq *ProjectServiceProviderQuery) Aggregate(fns ...AggregateFunc) *ProjectServiceProviderSelect {
	return pspq.Select().Aggregate(fns...)
}

func (pspq *ProjectServiceProviderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pspq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pspq); err != nil {
				return err
			}
		}
	}
	for _, f := range pspq.ctx.Fields {
		if !projectserviceprovider.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if pspq.path != nil {
		prev, err := pspq.path(ctx)
		if err != nil {
			return err
		}
		pspq.sql = prev
	}
	return nil
}

func (pspq *ProjectServiceProviderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProjectServiceProvider, error) {
	var (
		nodes       = []*ProjectServiceProvider{}
		_spec       = pspq.querySpec()
		loadedTypes = [2]bool{
			pspq.withProject != nil,
			pspq.withServiceProvider != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProjectServiceProvider).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProjectServiceProvider{config: pspq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pspq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pspq.withProject; query != nil {
		if err := pspq.loadProject(ctx, query, nodes, nil,
			func(n *ProjectServiceProvider, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	if query := pspq.withServiceProvider; query != nil {
		if err := pspq.loadServiceProvider(ctx, query, nodes, nil,
			func(n *ProjectServiceProvider, e *ServiceProvider) { n.Edges.ServiceProvider = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pspq *ProjectServiceProviderQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*ProjectServiceProvider, init func(*ProjectServiceProvider), assign func(*ProjectServiceProvider, *Project)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProjectServiceProvider)
	for i := range nodes {
		fk := nodes[i].ProjectProjectID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(project.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "project_project_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pspq *ProjectServiceProviderQuery) loadServiceProvider(ctx context.Context, query *ServiceProviderQuery, nodes []*ProjectServiceProvider, init func(*ProjectServiceProvider), assign func(*ProjectServiceProvider, *ServiceProvider)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProjectServiceProvider)
	for i := range nodes {
		fk := nodes[i].ProjectServiceProviderID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(serviceprovider.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "project_service_provider_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pspq *ProjectServiceProviderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pspq.querySpec()
	_spec.Node.Columns = pspq.ctx.Fields
	if len(pspq.ctx.Fields) > 0 {
		_spec.Unique = pspq.ctx.Unique != nil && *pspq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pspq.driver, _spec)
}

func (pspq *ProjectServiceProviderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(projectserviceprovider.Table, projectserviceprovider.Columns, sqlgraph.NewFieldSpec(projectserviceprovider.FieldID, field.TypeInt))
	_spec.From = pspq.sql
	if unique := pspq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pspq.path != nil {
		_spec.Unique = true
	}
	if fields := pspq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectserviceprovider.FieldID)
		for i := range fields {
			if fields[i] != projectserviceprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pspq.withProject != nil {
			_spec.Node.AddColumnOnce(projectserviceprovider.FieldProjectProjectID)
		}
		if pspq.withServiceProvider != nil {
			_spec.Node.AddColumnOnce(projectserviceprovider.FieldProjectServiceProviderID)
		}
	}
	if ps := pspq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pspq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pspq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pspq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pspq *ProjectServiceProviderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pspq.driver.Dialect())
	t1 := builder.Table(projectserviceprovider.Table)
	columns := pspq.ctx.Fields
	if len(columns) == 0 {
		columns = projectserviceprovider.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pspq.sql != nil {
		selector = pspq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pspq.ctx.Unique != nil && *pspq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pspq.predicates {
		p(selector)
	}
	for _, p := range pspq.order {
		p(selector)
	}
	if offset := pspq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pspq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProjectServiceProviderGroupBy is the group-by builder for ProjectServiceProvider entities.
type ProjectServiceProviderGroupBy struct {
	selector
	build *ProjectServiceProviderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pspgb *ProjectServiceProviderGroupBy) Aggregate(fns ...AggregateFunc) *ProjectServiceProviderGroupBy {
	pspgb.fns = append(pspgb.fns, fns...)
	return pspgb
}

// Scan applies the selector query and scans the result into the given value.
func (pspgb *ProjectServiceProviderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pspgb.build.ctx, "GroupBy")
	if err := pspgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProjectServiceProviderQuery, *ProjectServiceProviderGroupBy](ctx, pspgb.build, pspgb, pspgb.build.inters, v)
}

func (pspgb *ProjectServiceProviderGroupBy) sqlScan(ctx context.Context, root *ProjectServiceProviderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pspgb.fns))
	for _, fn := range pspgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pspgb.flds)+len(pspgb.fns))
		for _, f := range *pspgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pspgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pspgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProjectServiceProviderSelect is the builder for selecting fields of ProjectServiceProvider entities.
type ProjectServiceProviderSelect struct {
	*ProjectServiceProviderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (psps *ProjectServiceProviderSelect) Aggregate(fns ...AggregateFunc) *ProjectServiceProviderSelect {
	psps.fns = append(psps.fns, fns...)
	return psps
}

// Scan applies the selector query and scans the result into the given value.
func (psps *ProjectServiceProviderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, psps.ctx, "Select")
	if err := psps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProjectServiceProviderQuery, *ProjectServiceProviderSelect](ctx, psps.ProjectServiceProviderQuery, psps, psps.inters, v)
}

func (psps *ProjectServiceProviderSelect) sqlScan(ctx context.Context, root *ProjectServiceProviderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(psps.fns))
	for _, fn := range psps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*psps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
