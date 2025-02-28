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
	"github.com/Pyakz/buildbox-api/ent/generated/issue"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/Pyakz/buildbox-api/ent/generated/project"
	"github.com/Pyakz/buildbox-api/ent/generated/task"
	"github.com/Pyakz/buildbox-api/ent/generated/user"
)

// IssueQuery is the builder for querying Issue entities.
type IssueQuery struct {
	config
	ctx         *QueryContext
	order       []issue.OrderOption
	inters      []Interceptor
	predicates  []predicate.Issue
	withTasks   *TaskQuery
	withAccount *AccountQuery
	withUser    *UserQuery
	withProject *ProjectQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IssueQuery builder.
func (iq *IssueQuery) Where(ps ...predicate.Issue) *IssueQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *IssueQuery) Limit(limit int) *IssueQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *IssueQuery) Offset(offset int) *IssueQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *IssueQuery) Unique(unique bool) *IssueQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *IssueQuery) Order(o ...issue.OrderOption) *IssueQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryTasks chains the current query on the "tasks" edge.
func (iq *IssueQuery) QueryTasks() *TaskQuery {
	query := (&TaskClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(issue.Table, issue.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, issue.TasksTable, issue.TasksColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAccount chains the current query on the "account" edge.
func (iq *IssueQuery) QueryAccount() *AccountQuery {
	query := (&AccountClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(issue.Table, issue.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, issue.AccountTable, issue.AccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (iq *IssueQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(issue.Table, issue.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, issue.UserTable, issue.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProject chains the current query on the "project" edge.
func (iq *IssueQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(issue.Table, issue.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, issue.ProjectTable, issue.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Issue entity from the query.
// Returns a *NotFoundError when no Issue was found.
func (iq *IssueQuery) First(ctx context.Context) (*Issue, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{issue.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *IssueQuery) FirstX(ctx context.Context) *Issue {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Issue ID from the query.
// Returns a *NotFoundError when no Issue ID was found.
func (iq *IssueQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{issue.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *IssueQuery) FirstIDX(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Issue entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Issue entity is found.
// Returns a *NotFoundError when no Issue entities are found.
func (iq *IssueQuery) Only(ctx context.Context) (*Issue, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{issue.Label}
	default:
		return nil, &NotSingularError{issue.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *IssueQuery) OnlyX(ctx context.Context) *Issue {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Issue ID in the query.
// Returns a *NotSingularError when more than one Issue ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *IssueQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{issue.Label}
	default:
		err = &NotSingularError{issue.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *IssueQuery) OnlyIDX(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Issues.
func (iq *IssueQuery) All(ctx context.Context) ([]*Issue, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Issue, *IssueQuery]()
	return withInterceptors[[]*Issue](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *IssueQuery) AllX(ctx context.Context) []*Issue {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Issue IDs.
func (iq *IssueQuery) IDs(ctx context.Context) (ids []int, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(issue.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *IssueQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *IssueQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*IssueQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *IssueQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *IssueQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *IssueQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IssueQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *IssueQuery) Clone() *IssueQuery {
	if iq == nil {
		return nil
	}
	return &IssueQuery{
		config:      iq.config,
		ctx:         iq.ctx.Clone(),
		order:       append([]issue.OrderOption{}, iq.order...),
		inters:      append([]Interceptor{}, iq.inters...),
		predicates:  append([]predicate.Issue{}, iq.predicates...),
		withTasks:   iq.withTasks.Clone(),
		withAccount: iq.withAccount.Clone(),
		withUser:    iq.withUser.Clone(),
		withProject: iq.withProject.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithTasks tells the query-builder to eager-load the nodes that are connected to
// the "tasks" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IssueQuery) WithTasks(opts ...func(*TaskQuery)) *IssueQuery {
	query := (&TaskClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withTasks = query
	return iq
}

// WithAccount tells the query-builder to eager-load the nodes that are connected to
// the "account" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IssueQuery) WithAccount(opts ...func(*AccountQuery)) *IssueQuery {
	query := (&AccountClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withAccount = query
	return iq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IssueQuery) WithUser(opts ...func(*UserQuery)) *IssueQuery {
	query := (&UserClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withUser = query
	return iq
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IssueQuery) WithProject(opts ...func(*ProjectQuery)) *IssueQuery {
	query := (&ProjectClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withProject = query
	return iq
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
//	client.Issue.Query().
//		GroupBy(issue.FieldAccountID).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (iq *IssueQuery) GroupBy(field string, fields ...string) *IssueGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IssueGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = issue.Label
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
//	client.Issue.Query().
//		Select(issue.FieldAccountID).
//		Scan(ctx, &v)
func (iq *IssueQuery) Select(fields ...string) *IssueSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &IssueSelect{IssueQuery: iq}
	sbuild.label = issue.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IssueSelect configured with the given aggregations.
func (iq *IssueQuery) Aggregate(fns ...AggregateFunc) *IssueSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *IssueQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !issue.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *IssueQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Issue, error) {
	var (
		nodes       = []*Issue{}
		_spec       = iq.querySpec()
		loadedTypes = [4]bool{
			iq.withTasks != nil,
			iq.withAccount != nil,
			iq.withUser != nil,
			iq.withProject != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Issue).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Issue{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withTasks; query != nil {
		if err := iq.loadTasks(ctx, query, nodes,
			func(n *Issue) { n.Edges.Tasks = []*Task{} },
			func(n *Issue, e *Task) { n.Edges.Tasks = append(n.Edges.Tasks, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withAccount; query != nil {
		if err := iq.loadAccount(ctx, query, nodes, nil,
			func(n *Issue, e *Account) { n.Edges.Account = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withUser; query != nil {
		if err := iq.loadUser(ctx, query, nodes, nil,
			func(n *Issue, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withProject; query != nil {
		if err := iq.loadProject(ctx, query, nodes, nil,
			func(n *Issue, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *IssueQuery) loadTasks(ctx context.Context, query *TaskQuery, nodes []*Issue, init func(*Issue), assign func(*Issue, *Task)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Issue)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(task.FieldTaskIssueID)
	}
	query.Where(predicate.Task(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(issue.TasksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TaskIssueID
		if fk == nil {
			return fmt.Errorf(`foreign-key "task_issue_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "task_issue_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (iq *IssueQuery) loadAccount(ctx context.Context, query *AccountQuery, nodes []*Issue, init func(*Issue), assign func(*Issue, *Account)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Issue)
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
func (iq *IssueQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Issue, init func(*Issue), assign func(*Issue, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Issue)
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
func (iq *IssueQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*Issue, init func(*Issue), assign func(*Issue, *Project)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Issue)
	for i := range nodes {
		fk := nodes[i].ProjectID
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
			return fmt.Errorf(`unexpected foreign-key "project_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (iq *IssueQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *IssueQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(issue.Table, issue.Columns, sqlgraph.NewFieldSpec(issue.FieldID, field.TypeInt))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, issue.FieldID)
		for i := range fields {
			if fields[i] != issue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if iq.withAccount != nil {
			_spec.Node.AddColumnOnce(issue.FieldAccountID)
		}
		if iq.withUser != nil {
			_spec.Node.AddColumnOnce(issue.FieldCreatedBy)
		}
		if iq.withProject != nil {
			_spec.Node.AddColumnOnce(issue.FieldProjectID)
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *IssueQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(issue.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = issue.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IssueGroupBy is the group-by builder for Issue entities.
type IssueGroupBy struct {
	selector
	build *IssueQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *IssueGroupBy) Aggregate(fns ...AggregateFunc) *IssueGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *IssueGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IssueQuery, *IssueGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *IssueGroupBy) sqlScan(ctx context.Context, root *IssueQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IssueSelect is the builder for selecting fields of Issue entities.
type IssueSelect struct {
	*IssueQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *IssueSelect) Aggregate(fns ...AggregateFunc) *IssueSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *IssueSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IssueQuery, *IssueSelect](ctx, is.IssueQuery, is, is.inters, v)
}

func (is *IssueSelect) sqlScan(ctx context.Context, root *IssueQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
