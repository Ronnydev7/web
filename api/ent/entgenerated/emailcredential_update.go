// Code generated by ent, DO NOT EDIT.

package entgenerated

import (
	"api/ent/entgenerated/emailcredential"
	"api/ent/entgenerated/predicate"
	"api/ent/entgenerated/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmailCredentialUpdate is the builder for updating EmailCredential entities.
type EmailCredentialUpdate struct {
	config
	hooks    []Hook
	mutation *EmailCredentialMutation
}

// Where appends a list predicates to the EmailCredentialUpdate builder.
func (ecu *EmailCredentialUpdate) Where(ps ...predicate.EmailCredential) *EmailCredentialUpdate {
	ecu.mutation.Where(ps...)
	return ecu
}

// SetOwnerID sets the "owner_id" field.
func (ecu *EmailCredentialUpdate) SetOwnerID(i int) *EmailCredentialUpdate {
	ecu.mutation.SetOwnerID(i)
	return ecu
}

// SetEmail sets the "email" field.
func (ecu *EmailCredentialUpdate) SetEmail(s string) *EmailCredentialUpdate {
	ecu.mutation.SetEmail(s)
	return ecu
}

// SetAlgorithm sets the "algorithm" field.
func (ecu *EmailCredentialUpdate) SetAlgorithm(e emailcredential.Algorithm) *EmailCredentialUpdate {
	ecu.mutation.SetAlgorithm(e)
	return ecu
}

// SetPasswordHash sets the "password_hash" field.
func (ecu *EmailCredentialUpdate) SetPasswordHash(b []byte) *EmailCredentialUpdate {
	ecu.mutation.SetPasswordHash(b)
	return ecu
}

// SetOwner sets the "owner" edge to the User entity.
func (ecu *EmailCredentialUpdate) SetOwner(u *User) *EmailCredentialUpdate {
	return ecu.SetOwnerID(u.ID)
}

// Mutation returns the EmailCredentialMutation object of the builder.
func (ecu *EmailCredentialUpdate) Mutation() *EmailCredentialMutation {
	return ecu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (ecu *EmailCredentialUpdate) ClearOwner() *EmailCredentialUpdate {
	ecu.mutation.ClearOwner()
	return ecu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ecu *EmailCredentialUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ecu.sqlSave, ecu.mutation, ecu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecu *EmailCredentialUpdate) SaveX(ctx context.Context) int {
	affected, err := ecu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ecu *EmailCredentialUpdate) Exec(ctx context.Context) error {
	_, err := ecu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecu *EmailCredentialUpdate) ExecX(ctx context.Context) {
	if err := ecu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecu *EmailCredentialUpdate) check() error {
	if v, ok := ecu.mutation.Email(); ok {
		if err := emailcredential.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`entgenerated: validator failed for field "EmailCredential.email": %w`, err)}
		}
	}
	if v, ok := ecu.mutation.Algorithm(); ok {
		if err := emailcredential.AlgorithmValidator(v); err != nil {
			return &ValidationError{Name: "algorithm", err: fmt.Errorf(`entgenerated: validator failed for field "EmailCredential.algorithm": %w`, err)}
		}
	}
	if _, ok := ecu.mutation.OwnerID(); ecu.mutation.OwnerCleared() && !ok {
		return errors.New(`entgenerated: clearing a required unique edge "EmailCredential.owner"`)
	}
	return nil
}

func (ecu *EmailCredentialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ecu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(emailcredential.Table, emailcredential.Columns, sqlgraph.NewFieldSpec(emailcredential.FieldID, field.TypeInt))
	if ps := ecu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecu.mutation.Email(); ok {
		_spec.SetField(emailcredential.FieldEmail, field.TypeString, value)
	}
	if value, ok := ecu.mutation.Algorithm(); ok {
		_spec.SetField(emailcredential.FieldAlgorithm, field.TypeEnum, value)
	}
	if value, ok := ecu.mutation.PasswordHash(); ok {
		_spec.SetField(emailcredential.FieldPasswordHash, field.TypeBytes, value)
	}
	if ecu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   emailcredential.OwnerTable,
			Columns: []string{emailcredential.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   emailcredential.OwnerTable,
			Columns: []string{emailcredential.OwnerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ecu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailcredential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ecu.mutation.done = true
	return n, nil
}

// EmailCredentialUpdateOne is the builder for updating a single EmailCredential entity.
type EmailCredentialUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmailCredentialMutation
}

// SetOwnerID sets the "owner_id" field.
func (ecuo *EmailCredentialUpdateOne) SetOwnerID(i int) *EmailCredentialUpdateOne {
	ecuo.mutation.SetOwnerID(i)
	return ecuo
}

// SetEmail sets the "email" field.
func (ecuo *EmailCredentialUpdateOne) SetEmail(s string) *EmailCredentialUpdateOne {
	ecuo.mutation.SetEmail(s)
	return ecuo
}

// SetAlgorithm sets the "algorithm" field.
func (ecuo *EmailCredentialUpdateOne) SetAlgorithm(e emailcredential.Algorithm) *EmailCredentialUpdateOne {
	ecuo.mutation.SetAlgorithm(e)
	return ecuo
}

// SetPasswordHash sets the "password_hash" field.
func (ecuo *EmailCredentialUpdateOne) SetPasswordHash(b []byte) *EmailCredentialUpdateOne {
	ecuo.mutation.SetPasswordHash(b)
	return ecuo
}

// SetOwner sets the "owner" edge to the User entity.
func (ecuo *EmailCredentialUpdateOne) SetOwner(u *User) *EmailCredentialUpdateOne {
	return ecuo.SetOwnerID(u.ID)
}

// Mutation returns the EmailCredentialMutation object of the builder.
func (ecuo *EmailCredentialUpdateOne) Mutation() *EmailCredentialMutation {
	return ecuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (ecuo *EmailCredentialUpdateOne) ClearOwner() *EmailCredentialUpdateOne {
	ecuo.mutation.ClearOwner()
	return ecuo
}

// Where appends a list predicates to the EmailCredentialUpdate builder.
func (ecuo *EmailCredentialUpdateOne) Where(ps ...predicate.EmailCredential) *EmailCredentialUpdateOne {
	ecuo.mutation.Where(ps...)
	return ecuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ecuo *EmailCredentialUpdateOne) Select(field string, fields ...string) *EmailCredentialUpdateOne {
	ecuo.fields = append([]string{field}, fields...)
	return ecuo
}

// Save executes the query and returns the updated EmailCredential entity.
func (ecuo *EmailCredentialUpdateOne) Save(ctx context.Context) (*EmailCredential, error) {
	return withHooks(ctx, ecuo.sqlSave, ecuo.mutation, ecuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecuo *EmailCredentialUpdateOne) SaveX(ctx context.Context) *EmailCredential {
	node, err := ecuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ecuo *EmailCredentialUpdateOne) Exec(ctx context.Context) error {
	_, err := ecuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecuo *EmailCredentialUpdateOne) ExecX(ctx context.Context) {
	if err := ecuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecuo *EmailCredentialUpdateOne) check() error {
	if v, ok := ecuo.mutation.Email(); ok {
		if err := emailcredential.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`entgenerated: validator failed for field "EmailCredential.email": %w`, err)}
		}
	}
	if v, ok := ecuo.mutation.Algorithm(); ok {
		if err := emailcredential.AlgorithmValidator(v); err != nil {
			return &ValidationError{Name: "algorithm", err: fmt.Errorf(`entgenerated: validator failed for field "EmailCredential.algorithm": %w`, err)}
		}
	}
	if _, ok := ecuo.mutation.OwnerID(); ecuo.mutation.OwnerCleared() && !ok {
		return errors.New(`entgenerated: clearing a required unique edge "EmailCredential.owner"`)
	}
	return nil
}

func (ecuo *EmailCredentialUpdateOne) sqlSave(ctx context.Context) (_node *EmailCredential, err error) {
	if err := ecuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(emailcredential.Table, emailcredential.Columns, sqlgraph.NewFieldSpec(emailcredential.FieldID, field.TypeInt))
	id, ok := ecuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entgenerated: missing "EmailCredential.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ecuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emailcredential.FieldID)
		for _, f := range fields {
			if !emailcredential.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entgenerated: invalid field %q for query", f)}
			}
			if f != emailcredential.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ecuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecuo.mutation.Email(); ok {
		_spec.SetField(emailcredential.FieldEmail, field.TypeString, value)
	}
	if value, ok := ecuo.mutation.Algorithm(); ok {
		_spec.SetField(emailcredential.FieldAlgorithm, field.TypeEnum, value)
	}
	if value, ok := ecuo.mutation.PasswordHash(); ok {
		_spec.SetField(emailcredential.FieldPasswordHash, field.TypeBytes, value)
	}
	if ecuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   emailcredential.OwnerTable,
			Columns: []string{emailcredential.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   emailcredential.OwnerTable,
			Columns: []string{emailcredential.OwnerColumn},
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
	_node = &EmailCredential{config: ecuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ecuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailcredential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ecuo.mutation.done = true
	return _node, nil
}
