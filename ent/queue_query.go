// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lou-core/ent/city"
	"lou-core/ent/construction"
	"lou-core/ent/predicate"
	"lou-core/ent/queue"
	"lou-core/ent/user"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// QueueQuery is the builder for querying Queue entities.
type QueueQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Queue
	// eager-loading edges.
	withOwner        *UserQuery
	withCity         *CityQuery
	withConstruction *ConstructionQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (qq *QueueQuery) Where(ps ...predicate.Queue) *QueueQuery {
	qq.predicates = append(qq.predicates, ps...)
	return qq
}

// Limit adds a limit step to the query.
func (qq *QueueQuery) Limit(limit int) *QueueQuery {
	qq.limit = &limit
	return qq
}

// Offset adds an offset step to the query.
func (qq *QueueQuery) Offset(offset int) *QueueQuery {
	qq.offset = &offset
	return qq
}

// Order adds an order step to the query.
func (qq *QueueQuery) Order(o ...OrderFunc) *QueueQuery {
	qq.order = append(qq.order, o...)
	return qq
}

// QueryOwner chains the current query on the owner edge.
func (qq *QueueQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: qq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(queue.Table, queue.FieldID, qq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, queue.OwnerTable, queue.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(qq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCity chains the current query on the city edge.
func (qq *QueueQuery) QueryCity() *CityQuery {
	query := &CityQuery{config: qq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(queue.Table, queue.FieldID, qq.sqlQuery()),
			sqlgraph.To(city.Table, city.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, queue.CityTable, queue.CityColumn),
		)
		fromU = sqlgraph.SetNeighbors(qq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryConstruction chains the current query on the construction edge.
func (qq *QueueQuery) QueryConstruction() *ConstructionQuery {
	query := &ConstructionQuery{config: qq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(queue.Table, queue.FieldID, qq.sqlQuery()),
			sqlgraph.To(construction.Table, construction.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, queue.ConstructionTable, queue.ConstructionColumn),
		)
		fromU = sqlgraph.SetNeighbors(qq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Queue entity in the query. Returns *NotFoundError when no queue was found.
func (qq *QueueQuery) First(ctx context.Context) (*Queue, error) {
	qs, err := qq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(qs) == 0 {
		return nil, &NotFoundError{queue.Label}
	}
	return qs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qq *QueueQuery) FirstX(ctx context.Context) *Queue {
	q, err := qq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return q
}

// FirstID returns the first Queue id in the query. Returns *NotFoundError when no id was found.
func (qq *QueueQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{queue.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (qq *QueueQuery) FirstXID(ctx context.Context) int {
	id, err := qq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Queue entity in the query, returns an error if not exactly one entity was returned.
func (qq *QueueQuery) Only(ctx context.Context) (*Queue, error) {
	qs, err := qq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(qs) {
	case 1:
		return qs[0], nil
	case 0:
		return nil, &NotFoundError{queue.Label}
	default:
		return nil, &NotSingularError{queue.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qq *QueueQuery) OnlyX(ctx context.Context) *Queue {
	q, err := qq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return q
}

// OnlyID returns the only Queue id in the query, returns an error if not exactly one id was returned.
func (qq *QueueQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{queue.Label}
	default:
		err = &NotSingularError{queue.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (qq *QueueQuery) OnlyXID(ctx context.Context) int {
	id, err := qq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Queues.
func (qq *QueueQuery) All(ctx context.Context) ([]*Queue, error) {
	if err := qq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return qq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (qq *QueueQuery) AllX(ctx context.Context) []*Queue {
	qs, err := qq.All(ctx)
	if err != nil {
		panic(err)
	}
	return qs
}

// IDs executes the query and returns a list of Queue ids.
func (qq *QueueQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := qq.Select(queue.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qq *QueueQuery) IDsX(ctx context.Context) []int {
	ids, err := qq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qq *QueueQuery) Count(ctx context.Context) (int, error) {
	if err := qq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return qq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (qq *QueueQuery) CountX(ctx context.Context) int {
	count, err := qq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qq *QueueQuery) Exist(ctx context.Context) (bool, error) {
	if err := qq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return qq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (qq *QueueQuery) ExistX(ctx context.Context) bool {
	exist, err := qq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qq *QueueQuery) Clone() *QueueQuery {
	return &QueueQuery{
		config:     qq.config,
		limit:      qq.limit,
		offset:     qq.offset,
		order:      append([]OrderFunc{}, qq.order...),
		unique:     append([]string{}, qq.unique...),
		predicates: append([]predicate.Queue{}, qq.predicates...),
		// clone intermediate query.
		sql:  qq.sql.Clone(),
		path: qq.path,
	}
}

//  WithOwner tells the query-builder to eager-loads the nodes that are connected to
// the "owner" edge. The optional arguments used to configure the query builder of the edge.
func (qq *QueueQuery) WithOwner(opts ...func(*UserQuery)) *QueueQuery {
	query := &UserQuery{config: qq.config}
	for _, opt := range opts {
		opt(query)
	}
	qq.withOwner = query
	return qq
}

//  WithCity tells the query-builder to eager-loads the nodes that are connected to
// the "city" edge. The optional arguments used to configure the query builder of the edge.
func (qq *QueueQuery) WithCity(opts ...func(*CityQuery)) *QueueQuery {
	query := &CityQuery{config: qq.config}
	for _, opt := range opts {
		opt(query)
	}
	qq.withCity = query
	return qq
}

//  WithConstruction tells the query-builder to eager-loads the nodes that are connected to
// the "construction" edge. The optional arguments used to configure the query builder of the edge.
func (qq *QueueQuery) WithConstruction(opts ...func(*ConstructionQuery)) *QueueQuery {
	query := &ConstructionQuery{config: qq.config}
	for _, opt := range opts {
		opt(query)
	}
	qq.withConstruction = query
	return qq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Completion time.Time `json:"completion,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Queue.Query().
//		GroupBy(queue.FieldCompletion).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (qq *QueueQuery) GroupBy(field string, fields ...string) *QueueGroupBy {
	group := &QueueGroupBy{config: qq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return qq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Completion time.Time `json:"completion,omitempty"`
//	}
//
//	client.Queue.Query().
//		Select(queue.FieldCompletion).
//		Scan(ctx, &v)
//
func (qq *QueueQuery) Select(field string, fields ...string) *QueueSelect {
	selector := &QueueSelect{config: qq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return qq.sqlQuery(), nil
	}
	return selector
}

func (qq *QueueQuery) prepareQuery(ctx context.Context) error {
	if qq.path != nil {
		prev, err := qq.path(ctx)
		if err != nil {
			return err
		}
		qq.sql = prev
	}
	return nil
}

func (qq *QueueQuery) sqlAll(ctx context.Context) ([]*Queue, error) {
	var (
		nodes       = []*Queue{}
		withFKs     = qq.withFKs
		_spec       = qq.querySpec()
		loadedTypes = [3]bool{
			qq.withOwner != nil,
			qq.withCity != nil,
			qq.withConstruction != nil,
		}
	)
	if qq.withOwner != nil || qq.withCity != nil || qq.withConstruction != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, queue.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Queue{config: qq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, qq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := qq.withOwner; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Queue)
		for i := range nodes {
			if fk := nodes[i].user_queue; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_queue" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	if query := qq.withCity; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Queue)
		for i := range nodes {
			if fk := nodes[i].city_queue; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(city.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "city_queue" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.City = n
			}
		}
	}

	if query := qq.withConstruction; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Queue)
		for i := range nodes {
			if fk := nodes[i].construction_queue; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(construction.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "construction_queue" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Construction = n
			}
		}
	}

	return nodes, nil
}

func (qq *QueueQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qq.querySpec()
	return sqlgraph.CountNodes(ctx, qq.driver, _spec)
}

func (qq *QueueQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := qq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (qq *QueueQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   queue.Table,
			Columns: queue.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: queue.FieldID,
			},
		},
		From:   qq.sql,
		Unique: true,
	}
	if ps := qq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (qq *QueueQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(qq.driver.Dialect())
	t1 := builder.Table(queue.Table)
	selector := builder.Select(t1.Columns(queue.Columns...)...).From(t1)
	if qq.sql != nil {
		selector = qq.sql
		selector.Select(selector.Columns(queue.Columns...)...)
	}
	for _, p := range qq.predicates {
		p(selector)
	}
	for _, p := range qq.order {
		p(selector)
	}
	if offset := qq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// QueueGroupBy is the builder for group-by Queue entities.
type QueueGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qgb *QueueGroupBy) Aggregate(fns ...AggregateFunc) *QueueGroupBy {
	qgb.fns = append(qgb.fns, fns...)
	return qgb
}

// Scan applies the group-by query and scan the result into the given value.
func (qgb *QueueGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := qgb.path(ctx)
	if err != nil {
		return err
	}
	qgb.sql = query
	return qgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (qgb *QueueGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := qgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (qgb *QueueGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(qgb.fields) > 1 {
		return nil, errors.New("ent: QueueGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := qgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (qgb *QueueGroupBy) StringsX(ctx context.Context) []string {
	v, err := qgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (qgb *QueueGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(qgb.fields) > 1 {
		return nil, errors.New("ent: QueueGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := qgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (qgb *QueueGroupBy) IntsX(ctx context.Context) []int {
	v, err := qgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (qgb *QueueGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(qgb.fields) > 1 {
		return nil, errors.New("ent: QueueGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := qgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (qgb *QueueGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := qgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (qgb *QueueGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(qgb.fields) > 1 {
		return nil, errors.New("ent: QueueGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := qgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (qgb *QueueGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := qgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (qgb *QueueGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := qgb.sqlQuery().Query()
	if err := qgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (qgb *QueueGroupBy) sqlQuery() *sql.Selector {
	selector := qgb.sql
	columns := make([]string, 0, len(qgb.fields)+len(qgb.fns))
	columns = append(columns, qgb.fields...)
	for _, fn := range qgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(qgb.fields...)
}

// QueueSelect is the builder for select fields of Queue entities.
type QueueSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (qs *QueueSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := qs.path(ctx)
	if err != nil {
		return err
	}
	qs.sql = query
	return qs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (qs *QueueSelect) ScanX(ctx context.Context, v interface{}) {
	if err := qs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (qs *QueueSelect) Strings(ctx context.Context) ([]string, error) {
	if len(qs.fields) > 1 {
		return nil, errors.New("ent: QueueSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := qs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (qs *QueueSelect) StringsX(ctx context.Context) []string {
	v, err := qs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (qs *QueueSelect) Ints(ctx context.Context) ([]int, error) {
	if len(qs.fields) > 1 {
		return nil, errors.New("ent: QueueSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := qs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (qs *QueueSelect) IntsX(ctx context.Context) []int {
	v, err := qs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (qs *QueueSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(qs.fields) > 1 {
		return nil, errors.New("ent: QueueSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := qs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (qs *QueueSelect) Float64sX(ctx context.Context) []float64 {
	v, err := qs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (qs *QueueSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(qs.fields) > 1 {
		return nil, errors.New("ent: QueueSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := qs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (qs *QueueSelect) BoolsX(ctx context.Context) []bool {
	v, err := qs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (qs *QueueSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := qs.sqlQuery().Query()
	if err := qs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (qs *QueueSelect) sqlQuery() sql.Querier {
	selector := qs.sql
	selector.Select(selector.Columns(qs.fields...)...)
	return selector
}
