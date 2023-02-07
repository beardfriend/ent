// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/attachedfile"
	"entgo.io/ent/entc/integration/edgeschema/ent/file"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/process"
	"entgo.io/ent/schema/field"
)

// ProcessQuery is the builder for querying Process entities.
type ProcessQuery struct {
	config
	ctx               *QueryContext
	order             []OrderFunc
	inters            []Interceptor
	predicates        []predicate.Process
	withFiles         *FileQuery
	withAttachedFiles *AttachedFileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProcessQuery builder.
func (pq *ProcessQuery) Where(ps ...predicate.Process) *ProcessQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *ProcessQuery) Limit(limit int) *ProcessQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *ProcessQuery) Offset(offset int) *ProcessQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ProcessQuery) Unique(unique bool) *ProcessQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *ProcessQuery) Order(o ...OrderFunc) *ProcessQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryFiles chains the current query on the "files" edge.
func (pq *ProcessQuery) QueryFiles() *FileQuery {
	query := (&FileClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(process.Table, process.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, process.FilesTable, process.FilesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAttachedFiles chains the current query on the "attached_files" edge.
func (pq *ProcessQuery) QueryAttachedFiles() *AttachedFileQuery {
	query := (&AttachedFileClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(process.Table, process.FieldID, selector),
			sqlgraph.To(attachedfile.Table, attachedfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, process.AttachedFilesTable, process.AttachedFilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Process entity from the query.
// Returns a *NotFoundError when no Process was found.
func (pq *ProcessQuery) First(ctx context.Context) (*Process, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{process.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProcessQuery) FirstX(ctx context.Context) *Process {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Process ID from the query.
// Returns a *NotFoundError when no Process ID was found.
func (pq *ProcessQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{process.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ProcessQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Process entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Process entity is found.
// Returns a *NotFoundError when no Process entities are found.
func (pq *ProcessQuery) Only(ctx context.Context) (*Process, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{process.Label}
	default:
		return nil, &NotSingularError{process.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProcessQuery) OnlyX(ctx context.Context) *Process {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Process ID in the query.
// Returns a *NotSingularError when more than one Process ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ProcessQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{process.Label}
	default:
		err = &NotSingularError{process.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ProcessQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Processes.
func (pq *ProcessQuery) All(ctx context.Context) ([]*Process, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Process, *ProcessQuery]()
	return withInterceptors[[]*Process](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProcessQuery) AllX(ctx context.Context) []*Process {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Process IDs.
func (pq *ProcessQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(process.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProcessQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProcessQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*ProcessQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProcessQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProcessQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProcessQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProcessQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProcessQuery) Clone() *ProcessQuery {
	if pq == nil {
		return nil
	}
	return &ProcessQuery{
		config:            pq.config,
		ctx:               pq.ctx.Clone(),
		order:             append([]OrderFunc{}, pq.order...),
		inters:            append([]Interceptor{}, pq.inters...),
		predicates:        append([]predicate.Process{}, pq.predicates...),
		withFiles:         pq.withFiles.Clone(),
		withAttachedFiles: pq.withAttachedFiles.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProcessQuery) WithFiles(opts ...func(*FileQuery)) *ProcessQuery {
	query := (&FileClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withFiles = query
	return pq
}

// WithAttachedFiles tells the query-builder to eager-load the nodes that are connected to
// the "attached_files" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProcessQuery) WithAttachedFiles(opts ...func(*AttachedFileQuery)) *ProcessQuery {
	query := (&AttachedFileClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withAttachedFiles = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (pq *ProcessQuery) GroupBy(field string, fields ...string) *ProcessGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProcessGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = process.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (pq *ProcessQuery) Select(fields ...string) *ProcessSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &ProcessSelect{ProcessQuery: pq}
	sbuild.label = process.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProcessSelect configured with the given aggregations.
func (pq *ProcessQuery) Aggregate(fns ...AggregateFunc) *ProcessSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *ProcessQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !process.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *ProcessQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Process, error) {
	var (
		nodes       = []*Process{}
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withFiles != nil,
			pq.withAttachedFiles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Process).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Process{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withFiles; query != nil {
		if err := pq.loadFiles(ctx, query, nodes,
			func(n *Process) { n.Edges.Files = []*File{} },
			func(n *Process, e *File) { n.Edges.Files = append(n.Edges.Files, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withAttachedFiles; query != nil {
		if err := pq.loadAttachedFiles(ctx, query, nodes,
			func(n *Process) { n.Edges.AttachedFiles = []*AttachedFile{} },
			func(n *Process, e *AttachedFile) { n.Edges.AttachedFiles = append(n.Edges.AttachedFiles, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *ProcessQuery) loadFiles(ctx context.Context, query *FileQuery, nodes []*Process, init func(*Process), assign func(*Process, *File)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Process)
	nids := make(map[int]map[*Process]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(process.FilesTable)
		s.Join(joinT).On(s.C(file.FieldID), joinT.C(process.FilesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(process.FilesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(process.FilesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Process]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*File](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "files" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *ProcessQuery) loadAttachedFiles(ctx context.Context, query *AttachedFileQuery, nodes []*Process, init func(*Process), assign func(*Process, *AttachedFile)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Process)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.AttachedFile(func(s *sql.Selector) {
		s.Where(sql.InValues(process.AttachedFilesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProcID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "proc_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *ProcessQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProcessQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(process.Table, process.Columns, sqlgraph.NewFieldSpec(process.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, process.FieldID)
		for i := range fields {
			if fields[i] != process.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ProcessQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(process.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = process.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProcessGroupBy is the group-by builder for Process entities.
type ProcessGroupBy struct {
	selector
	build *ProcessQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProcessGroupBy) Aggregate(fns ...AggregateFunc) *ProcessGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *ProcessGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProcessQuery, *ProcessGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *ProcessGroupBy) sqlScan(ctx context.Context, root *ProcessQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProcessSelect is the builder for selecting fields of Process entities.
type ProcessSelect struct {
	*ProcessQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ProcessSelect) Aggregate(fns ...AggregateFunc) *ProcessSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ProcessSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProcessQuery, *ProcessSelect](ctx, ps.ProcessQuery, ps, ps.inters, v)
}

func (ps *ProcessSelect) sqlScan(ctx context.Context, root *ProcessQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
