// Code generated by ent, DO NOT EDIT.

package entgenerated

import (
	"api/ent/entgenerated/emailcredential"
	"api/ent/entgenerated/loginsession"
	"api/ent/entgenerated/predicate"
	"api/ent/entgenerated/superuserprofile"
	"api/ent/entgenerated/user"
	"api/ent/entgenerated/userpublicprofile"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	ctx                    *QueryContext
	order                  []user.OrderOption
	inters                 []Interceptor
	predicates             []predicate.User
	withSuperuserProfile   *SuperuserProfileQuery
	withEmailCredential    *EmailCredentialQuery
	withLoginSessions      *LoginSessionQuery
	withPublicProfile      *UserPublicProfileQuery
	modifiers              []func(*sql.Selector)
	loadTotal              []func(context.Context, []*User) error
	withNamedLoginSessions map[string]*LoginSessionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserQuery builder.
func (uq *UserQuery) Where(ps ...predicate.User) *UserQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UserQuery) Limit(limit int) *UserQuery {
	uq.ctx.Limit = &limit
	return uq
}

// Offset to start from.
func (uq *UserQuery) Offset(offset int) *UserQuery {
	uq.ctx.Offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UserQuery) Unique(unique bool) *UserQuery {
	uq.ctx.Unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UserQuery) Order(o ...user.OrderOption) *UserQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QuerySuperuserProfile chains the current query on the "superuser_profile" edge.
func (uq *UserQuery) QuerySuperuserProfile() *SuperuserProfileQuery {
	query := (&SuperuserProfileClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(superuserprofile.Table, superuserprofile.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.SuperuserProfileTable, user.SuperuserProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEmailCredential chains the current query on the "email_credential" edge.
func (uq *UserQuery) QueryEmailCredential() *EmailCredentialQuery {
	query := (&EmailCredentialClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(emailcredential.Table, emailcredential.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.EmailCredentialTable, user.EmailCredentialColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLoginSessions chains the current query on the "login_sessions" edge.
func (uq *UserQuery) QueryLoginSessions() *LoginSessionQuery {
	query := (&LoginSessionClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(loginsession.Table, loginsession.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.LoginSessionsTable, user.LoginSessionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPublicProfile chains the current query on the "public_profile" edge.
func (uq *UserQuery) QueryPublicProfile() *UserPublicProfileQuery {
	query := (&UserPublicProfileClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(userpublicprofile.Table, userpublicprofile.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.PublicProfileTable, user.PublicProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first User entity from the query.
// Returns a *NotFoundError when no User was found.
func (uq *UserQuery) First(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(1).All(setContextOp(ctx, uq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserQuery) FirstX(ctx context.Context) *User {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User ID from the query.
// Returns a *NotFoundError when no User ID was found.
func (uq *UserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(setContextOp(ctx, uq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UserQuery) FirstIDX(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User entity is found.
// Returns a *NotFoundError when no User entities are found.
func (uq *UserQuery) Only(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(2).All(setContextOp(ctx, uq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserQuery) OnlyX(ctx context.Context) *User {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User ID in the query.
// Returns a *NotSingularError when more than one User ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(setContextOp(ctx, uq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UserQuery) OnlyIDX(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (uq *UserQuery) All(ctx context.Context) ([]*User, error) {
	ctx = setContextOp(ctx, uq.ctx, "All")
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*User, *UserQuery]()
	return withInterceptors[[]*User](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserQuery) AllX(ctx context.Context) []*User {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User IDs.
func (uq *UserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if uq.ctx.Unique == nil && uq.path != nil {
		uq.Unique(true)
	}
	ctx = setContextOp(ctx, uq.ctx, "IDs")
	if err = uq.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uq.ctx, "Count")
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UserQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uq.ctx, "Exist")
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entgenerated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserQuery) Clone() *UserQuery {
	if uq == nil {
		return nil
	}
	return &UserQuery{
		config:               uq.config,
		ctx:                  uq.ctx.Clone(),
		order:                append([]user.OrderOption{}, uq.order...),
		inters:               append([]Interceptor{}, uq.inters...),
		predicates:           append([]predicate.User{}, uq.predicates...),
		withSuperuserProfile: uq.withSuperuserProfile.Clone(),
		withEmailCredential:  uq.withEmailCredential.Clone(),
		withLoginSessions:    uq.withLoginSessions.Clone(),
		withPublicProfile:    uq.withPublicProfile.Clone(),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// WithSuperuserProfile tells the query-builder to eager-load the nodes that are connected to
// the "superuser_profile" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithSuperuserProfile(opts ...func(*SuperuserProfileQuery)) *UserQuery {
	query := (&SuperuserProfileClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withSuperuserProfile = query
	return uq
}

// WithEmailCredential tells the query-builder to eager-load the nodes that are connected to
// the "email_credential" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithEmailCredential(opts ...func(*EmailCredentialQuery)) *UserQuery {
	query := (&EmailCredentialClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withEmailCredential = query
	return uq
}

// WithLoginSessions tells the query-builder to eager-load the nodes that are connected to
// the "login_sessions" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithLoginSessions(opts ...func(*LoginSessionQuery)) *UserQuery {
	query := (&LoginSessionClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withLoginSessions = query
	return uq
}

// WithPublicProfile tells the query-builder to eager-load the nodes that are connected to
// the "public_profile" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithPublicProfile(opts ...func(*UserPublicProfileQuery)) *UserQuery {
	query := (&UserPublicProfileClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withPublicProfile = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (uq *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	uq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserGroupBy{build: uq}
	grbuild.flds = &uq.ctx.Fields
	grbuild.label = user.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (uq *UserQuery) Select(fields ...string) *UserSelect {
	uq.ctx.Fields = append(uq.ctx.Fields, fields...)
	sbuild := &UserSelect{UserQuery: uq}
	sbuild.label = user.Label
	sbuild.flds, sbuild.scan = &uq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSelect configured with the given aggregations.
func (uq *UserQuery) Aggregate(fns ...AggregateFunc) *UserSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("entgenerated: uninitialized interceptor (forgotten import entgenerated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.ctx.Fields {
		if !user.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entgenerated: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	if user.Policy == nil {
		return errors.New("entgenerated: uninitialized user.Policy (forgotten import entgenerated/runtime?)")
	}
	if err := user.Policy.EvalQuery(ctx, uq); err != nil {
		return err
	}
	return nil
}

func (uq *UserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*User, error) {
	var (
		nodes       = []*User{}
		_spec       = uq.querySpec()
		loadedTypes = [4]bool{
			uq.withSuperuserProfile != nil,
			uq.withEmailCredential != nil,
			uq.withLoginSessions != nil,
			uq.withPublicProfile != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*User).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &User{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(uq.modifiers) > 0 {
		_spec.Modifiers = uq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withSuperuserProfile; query != nil {
		if err := uq.loadSuperuserProfile(ctx, query, nodes, nil,
			func(n *User, e *SuperuserProfile) { n.Edges.SuperuserProfile = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withEmailCredential; query != nil {
		if err := uq.loadEmailCredential(ctx, query, nodes, nil,
			func(n *User, e *EmailCredential) { n.Edges.EmailCredential = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withLoginSessions; query != nil {
		if err := uq.loadLoginSessions(ctx, query, nodes,
			func(n *User) { n.Edges.LoginSessions = []*LoginSession{} },
			func(n *User, e *LoginSession) { n.Edges.LoginSessions = append(n.Edges.LoginSessions, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withPublicProfile; query != nil {
		if err := uq.loadPublicProfile(ctx, query, nodes, nil,
			func(n *User, e *UserPublicProfile) { n.Edges.PublicProfile = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range uq.withNamedLoginSessions {
		if err := uq.loadLoginSessions(ctx, query, nodes,
			func(n *User) { n.appendNamedLoginSessions(name) },
			func(n *User, e *LoginSession) { n.appendNamedLoginSessions(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range uq.loadTotal {
		if err := uq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UserQuery) loadSuperuserProfile(ctx context.Context, query *SuperuserProfileQuery, nodes []*User, init func(*User), assign func(*User, *SuperuserProfile)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(superuserprofile.FieldOwnerID)
	}
	query.Where(predicate.SuperuserProfile(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.SuperuserProfileColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OwnerID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "owner_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadEmailCredential(ctx context.Context, query *EmailCredentialQuery, nodes []*User, init func(*User), assign func(*User, *EmailCredential)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(emailcredential.FieldOwnerID)
	}
	query.Where(predicate.EmailCredential(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.EmailCredentialColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OwnerID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "owner_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadLoginSessions(ctx context.Context, query *LoginSessionQuery, nodes []*User, init func(*User), assign func(*User, *LoginSession)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(loginsession.FieldOwnerID)
	}
	query.Where(predicate.LoginSession(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.LoginSessionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OwnerID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "owner_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadPublicProfile(ctx context.Context, query *UserPublicProfileQuery, nodes []*User, init func(*User), assign func(*User, *UserPublicProfile)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(userpublicprofile.FieldOwnerID)
	}
	query.Where(predicate.UserPublicProfile(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.PublicProfileColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OwnerID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "owner_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (uq *UserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	if len(uq.modifiers) > 0 {
		_spec.Modifiers = uq.modifiers
	}
	_spec.Node.Columns = uq.ctx.Fields
	if len(uq.ctx.Fields) > 0 {
		_spec.Unique = uq.ctx.Unique != nil && *uq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	_spec.From = uq.sql
	if unique := uq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uq.path != nil {
		_spec.Unique = true
	}
	if fields := uq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for i := range fields {
			if fields[i] != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(user.Table)
	columns := uq.ctx.Fields
	if len(columns) == 0 {
		columns = user.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.ctx.Unique != nil && *uq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedLoginSessions tells the query-builder to eager-load the nodes that are connected to the "login_sessions"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNamedLoginSessions(name string, opts ...func(*LoginSessionQuery)) *UserQuery {
	query := (&LoginSessionClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uq.withNamedLoginSessions == nil {
		uq.withNamedLoginSessions = make(map[string]*LoginSessionQuery)
	}
	uq.withNamedLoginSessions[name] = query
	return uq
}

// UserGroupBy is the group-by builder for User entities.
type UserGroupBy struct {
	selector
	build *UserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, "GroupBy")
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UserGroupBy) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserSelect is the builder for selecting fields of User entities.
type UserSelect struct {
	*UserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UserSelect) Aggregate(fns ...AggregateFunc) *UserSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, "Select")
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserSelect](ctx, us.UserQuery, us, us.inters, v)
}

func (us *UserSelect) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
