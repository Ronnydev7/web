// Code generated by ent, DO NOT EDIT.

package entgenerated

import (
	"api/ent/entgenerated/loginsession"
	"api/ent/entgenerated/predicate"
	"api/ent/entgenerated/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LoginSessionUpdate is the builder for updating LoginSession entities.
type LoginSessionUpdate struct {
	config
	hooks    []Hook
	mutation *LoginSessionMutation
}

// Where appends a list predicates to the LoginSessionUpdate builder.
func (lsu *LoginSessionUpdate) Where(ps ...predicate.LoginSession) *LoginSessionUpdate {
	lsu.mutation.Where(ps...)
	return lsu
}

// SetOwnerID sets the "owner_id" field.
func (lsu *LoginSessionUpdate) SetOwnerID(i int) *LoginSessionUpdate {
	lsu.mutation.SetOwnerID(i)
	return lsu
}

// SetLastLoginTime sets the "last_login_time" field.
func (lsu *LoginSessionUpdate) SetLastLoginTime(t time.Time) *LoginSessionUpdate {
	lsu.mutation.SetLastLoginTime(t)
	return lsu
}

// SetOwner sets the "owner" edge to the User entity.
func (lsu *LoginSessionUpdate) SetOwner(u *User) *LoginSessionUpdate {
	return lsu.SetOwnerID(u.ID)
}

// Mutation returns the LoginSessionMutation object of the builder.
func (lsu *LoginSessionUpdate) Mutation() *LoginSessionMutation {
	return lsu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (lsu *LoginSessionUpdate) ClearOwner() *LoginSessionUpdate {
	lsu.mutation.ClearOwner()
	return lsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lsu *LoginSessionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lsu.sqlSave, lsu.mutation, lsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lsu *LoginSessionUpdate) SaveX(ctx context.Context) int {
	affected, err := lsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lsu *LoginSessionUpdate) Exec(ctx context.Context) error {
	_, err := lsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsu *LoginSessionUpdate) ExecX(ctx context.Context) {
	if err := lsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lsu *LoginSessionUpdate) check() error {
	if _, ok := lsu.mutation.OwnerID(); lsu.mutation.OwnerCleared() && !ok {
		return errors.New(`entgenerated: clearing a required unique edge "LoginSession.owner"`)
	}
	return nil
}

func (lsu *LoginSessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(loginsession.Table, loginsession.Columns, sqlgraph.NewFieldSpec(loginsession.FieldID, field.TypeInt))
	if ps := lsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lsu.mutation.LastLoginTime(); ok {
		_spec.SetField(loginsession.FieldLastLoginTime, field.TypeTime, value)
	}
	if lsu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loginsession.OwnerTable,
			Columns: []string{loginsession.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lsu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loginsession.OwnerTable,
			Columns: []string{loginsession.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loginsession.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lsu.mutation.done = true
	return n, nil
}

// LoginSessionUpdateOne is the builder for updating a single LoginSession entity.
type LoginSessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LoginSessionMutation
}

// SetOwnerID sets the "owner_id" field.
func (lsuo *LoginSessionUpdateOne) SetOwnerID(i int) *LoginSessionUpdateOne {
	lsuo.mutation.SetOwnerID(i)
	return lsuo
}

// SetLastLoginTime sets the "last_login_time" field.
func (lsuo *LoginSessionUpdateOne) SetLastLoginTime(t time.Time) *LoginSessionUpdateOne {
	lsuo.mutation.SetLastLoginTime(t)
	return lsuo
}

// SetOwner sets the "owner" edge to the User entity.
func (lsuo *LoginSessionUpdateOne) SetOwner(u *User) *LoginSessionUpdateOne {
	return lsuo.SetOwnerID(u.ID)
}

// Mutation returns the LoginSessionMutation object of the builder.
func (lsuo *LoginSessionUpdateOne) Mutation() *LoginSessionMutation {
	return lsuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (lsuo *LoginSessionUpdateOne) ClearOwner() *LoginSessionUpdateOne {
	lsuo.mutation.ClearOwner()
	return lsuo
}

// Where appends a list predicates to the LoginSessionUpdate builder.
func (lsuo *LoginSessionUpdateOne) Where(ps ...predicate.LoginSession) *LoginSessionUpdateOne {
	lsuo.mutation.Where(ps...)
	return lsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lsuo *LoginSessionUpdateOne) Select(field string, fields ...string) *LoginSessionUpdateOne {
	lsuo.fields = append([]string{field}, fields...)
	return lsuo
}

// Save executes the query and returns the updated LoginSession entity.
func (lsuo *LoginSessionUpdateOne) Save(ctx context.Context) (*LoginSession, error) {
	return withHooks(ctx, lsuo.sqlSave, lsuo.mutation, lsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lsuo *LoginSessionUpdateOne) SaveX(ctx context.Context) *LoginSession {
	node, err := lsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lsuo *LoginSessionUpdateOne) Exec(ctx context.Context) error {
	_, err := lsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsuo *LoginSessionUpdateOne) ExecX(ctx context.Context) {
	if err := lsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lsuo *LoginSessionUpdateOne) check() error {
	if _, ok := lsuo.mutation.OwnerID(); lsuo.mutation.OwnerCleared() && !ok {
		return errors.New(`entgenerated: clearing a required unique edge "LoginSession.owner"`)
	}
	return nil
}

func (lsuo *LoginSessionUpdateOne) sqlSave(ctx context.Context) (_node *LoginSession, err error) {
	if err := lsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(loginsession.Table, loginsession.Columns, sqlgraph.NewFieldSpec(loginsession.FieldID, field.TypeInt))
	id, ok := lsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entgenerated: missing "LoginSession.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loginsession.FieldID)
		for _, f := range fields {
			if !loginsession.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entgenerated: invalid field %q for query", f)}
			}
			if f != loginsession.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lsuo.mutation.LastLoginTime(); ok {
		_spec.SetField(loginsession.FieldLastLoginTime, field.TypeTime, value)
	}
	if lsuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loginsession.OwnerTable,
			Columns: []string{loginsession.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lsuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loginsession.OwnerTable,
			Columns: []string{loginsession.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LoginSession{config: lsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loginsession.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lsuo.mutation.done = true
	return _node, nil
}
