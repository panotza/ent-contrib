// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/contrib/entproto/internal/entprototest/ent/validmessage"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ValidMessageQuery is the builder for querying ValidMessage entities.
type ValidMessageQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ValidMessage
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ValidMessageQuery builder.
func (vmq *ValidMessageQuery) Where(ps ...predicate.ValidMessage) *ValidMessageQuery {
	vmq.predicates = append(vmq.predicates, ps...)
	return vmq
}

// Limit adds a limit step to the query.
func (vmq *ValidMessageQuery) Limit(limit int) *ValidMessageQuery {
	vmq.limit = &limit
	return vmq
}

// Offset adds an offset step to the query.
func (vmq *ValidMessageQuery) Offset(offset int) *ValidMessageQuery {
	vmq.offset = &offset
	return vmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vmq *ValidMessageQuery) Unique(unique bool) *ValidMessageQuery {
	vmq.unique = &unique
	return vmq
}

// Order adds an order step to the query.
func (vmq *ValidMessageQuery) Order(o ...OrderFunc) *ValidMessageQuery {
	vmq.order = append(vmq.order, o...)
	return vmq
}

// First returns the first ValidMessage entity from the query.
// Returns a *NotFoundError when no ValidMessage was found.
func (vmq *ValidMessageQuery) First(ctx context.Context) (*ValidMessage, error) {
	nodes, err := vmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{validmessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vmq *ValidMessageQuery) FirstX(ctx context.Context) *ValidMessage {
	node, err := vmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ValidMessage ID from the query.
// Returns a *NotFoundError when no ValidMessage ID was found.
func (vmq *ValidMessageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{validmessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vmq *ValidMessageQuery) FirstIDX(ctx context.Context) int {
	id, err := vmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ValidMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ValidMessage entity is found.
// Returns a *NotFoundError when no ValidMessage entities are found.
func (vmq *ValidMessageQuery) Only(ctx context.Context) (*ValidMessage, error) {
	nodes, err := vmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{validmessage.Label}
	default:
		return nil, &NotSingularError{validmessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vmq *ValidMessageQuery) OnlyX(ctx context.Context) *ValidMessage {
	node, err := vmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ValidMessage ID in the query.
// Returns a *NotSingularError when more than one ValidMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (vmq *ValidMessageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{validmessage.Label}
	default:
		err = &NotSingularError{validmessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vmq *ValidMessageQuery) OnlyIDX(ctx context.Context) int {
	id, err := vmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ValidMessages.
func (vmq *ValidMessageQuery) All(ctx context.Context) ([]*ValidMessage, error) {
	if err := vmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vmq *ValidMessageQuery) AllX(ctx context.Context) []*ValidMessage {
	nodes, err := vmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ValidMessage IDs.
func (vmq *ValidMessageQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := vmq.Select(validmessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vmq *ValidMessageQuery) IDsX(ctx context.Context) []int {
	ids, err := vmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vmq *ValidMessageQuery) Count(ctx context.Context) (int, error) {
	if err := vmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vmq *ValidMessageQuery) CountX(ctx context.Context) int {
	count, err := vmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vmq *ValidMessageQuery) Exist(ctx context.Context) (bool, error) {
	if err := vmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vmq *ValidMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := vmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ValidMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vmq *ValidMessageQuery) Clone() *ValidMessageQuery {
	if vmq == nil {
		return nil
	}
	return &ValidMessageQuery{
		config:     vmq.config,
		limit:      vmq.limit,
		offset:     vmq.offset,
		order:      append([]OrderFunc{}, vmq.order...),
		predicates: append([]predicate.ValidMessage{}, vmq.predicates...),
		// clone intermediate query.
		sql:    vmq.sql.Clone(),
		path:   vmq.path,
		unique: vmq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ValidMessage.Query().
//		GroupBy(validmessage.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vmq *ValidMessageQuery) GroupBy(field string, fields ...string) *ValidMessageGroupBy {
	grbuild := &ValidMessageGroupBy{config: vmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vmq.sqlQuery(ctx), nil
	}
	grbuild.label = validmessage.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.ValidMessage.Query().
//		Select(validmessage.FieldName).
//		Scan(ctx, &v)
//
func (vmq *ValidMessageQuery) Select(fields ...string) *ValidMessageSelect {
	vmq.fields = append(vmq.fields, fields...)
	selbuild := &ValidMessageSelect{ValidMessageQuery: vmq}
	selbuild.label = validmessage.Label
	selbuild.flds, selbuild.scan = &vmq.fields, selbuild.Scan
	return selbuild
}

func (vmq *ValidMessageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vmq.fields {
		if !validmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vmq.path != nil {
		prev, err := vmq.path(ctx)
		if err != nil {
			return err
		}
		vmq.sql = prev
	}
	return nil
}

func (vmq *ValidMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ValidMessage, error) {
	var (
		nodes = []*ValidMessage{}
		_spec = vmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ValidMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ValidMessage{config: vmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vmq *ValidMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vmq.querySpec()
	_spec.Node.Columns = vmq.fields
	if len(vmq.fields) > 0 {
		_spec.Unique = vmq.unique != nil && *vmq.unique
	}
	return sqlgraph.CountNodes(ctx, vmq.driver, _spec)
}

func (vmq *ValidMessageQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (vmq *ValidMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   validmessage.Table,
			Columns: validmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: validmessage.FieldID,
			},
		},
		From:   vmq.sql,
		Unique: true,
	}
	if unique := vmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, validmessage.FieldID)
		for i := range fields {
			if fields[i] != validmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vmq *ValidMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vmq.driver.Dialect())
	t1 := builder.Table(validmessage.Table)
	columns := vmq.fields
	if len(columns) == 0 {
		columns = validmessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vmq.sql != nil {
		selector = vmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vmq.unique != nil && *vmq.unique {
		selector.Distinct()
	}
	for _, p := range vmq.predicates {
		p(selector)
	}
	for _, p := range vmq.order {
		p(selector)
	}
	if offset := vmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ValidMessageGroupBy is the group-by builder for ValidMessage entities.
type ValidMessageGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vmgb *ValidMessageGroupBy) Aggregate(fns ...AggregateFunc) *ValidMessageGroupBy {
	vmgb.fns = append(vmgb.fns, fns...)
	return vmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vmgb *ValidMessageGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vmgb.path(ctx)
	if err != nil {
		return err
	}
	vmgb.sql = query
	return vmgb.sqlScan(ctx, v)
}

func (vmgb *ValidMessageGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vmgb.fields {
		if !validmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vmgb *ValidMessageGroupBy) sqlQuery() *sql.Selector {
	selector := vmgb.sql.Select()
	aggregation := make([]string, 0, len(vmgb.fns))
	for _, fn := range vmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vmgb.fields)+len(vmgb.fns))
		for _, f := range vmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vmgb.fields...)...)
}

// ValidMessageSelect is the builder for selecting fields of ValidMessage entities.
type ValidMessageSelect struct {
	*ValidMessageQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vms *ValidMessageSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vms.prepareQuery(ctx); err != nil {
		return err
	}
	vms.sql = vms.ValidMessageQuery.sqlQuery(ctx)
	return vms.sqlScan(ctx, v)
}

func (vms *ValidMessageSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vms.sql.Query()
	if err := vms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
