// Code generated by ent, DO NOT EDIT.

package entgenerated

import (
	"api/ent/entgenerated/emailcredential"
	"api/ent/entgenerated/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmailCredentialCreate is the builder for creating a EmailCredential entity.
type EmailCredentialCreate struct {
	config
	mutation *EmailCredentialMutation
	hooks    []Hook
}

// SetOwnerID sets the "owner_id" field.
func (ecc *EmailCredentialCreate) SetOwnerID(i int) *EmailCredentialCreate {
	ecc.mutation.SetOwnerID(i)
	return ecc
}

// SetEmail sets the "email" field.
func (ecc *EmailCredentialCreate) SetEmail(s string) *EmailCredentialCreate {
	ecc.mutation.SetEmail(s)
	return ecc
}

// SetAlgorithm sets the "algorithm" field.
func (ecc *EmailCredentialCreate) SetAlgorithm(e emailcredential.Algorithm) *EmailCredentialCreate {
	ecc.mutation.SetAlgorithm(e)
	return ecc
}

// SetPasswordHash sets the "password_hash" field.
func (ecc *EmailCredentialCreate) SetPasswordHash(b []byte) *EmailCredentialCreate {
	ecc.mutation.SetPasswordHash(b)
	return ecc
}

// SetOwner sets the "owner" edge to the User entity.
func (ecc *EmailCredentialCreate) SetOwner(u *User) *EmailCredentialCreate {
	return ecc.SetOwnerID(u.ID)
}

// Mutation returns the EmailCredentialMutation object of the builder.
func (ecc *EmailCredentialCreate) Mutation() *EmailCredentialMutation {
	return ecc.mutation
}

// Save creates the EmailCredential in the database.
func (ecc *EmailCredentialCreate) Save(ctx context.Context) (*EmailCredential, error) {
	return withHooks(ctx, ecc.sqlSave, ecc.mutation, ecc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ecc *EmailCredentialCreate) SaveX(ctx context.Context) *EmailCredential {
	v, err := ecc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecc *EmailCredentialCreate) Exec(ctx context.Context) error {
	_, err := ecc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecc *EmailCredentialCreate) ExecX(ctx context.Context) {
	if err := ecc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecc *EmailCredentialCreate) check() error {
	if _, ok := ecc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`entgenerated: missing required field "EmailCredential.owner_id"`)}
	}
	if _, ok := ecc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`entgenerated: missing required field "EmailCredential.email"`)}
	}
	if v, ok := ecc.mutation.Email(); ok {
		if err := emailcredential.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`entgenerated: validator failed for field "EmailCredential.email": %w`, err)}
		}
	}
	if _, ok := ecc.mutation.Algorithm(); !ok {
		return &ValidationError{Name: "algorithm", err: errors.New(`entgenerated: missing required field "EmailCredential.algorithm"`)}
	}
	if v, ok := ecc.mutation.Algorithm(); ok {
		if err := emailcredential.AlgorithmValidator(v); err != nil {
			return &ValidationError{Name: "algorithm", err: fmt.Errorf(`entgenerated: validator failed for field "EmailCredential.algorithm": %w`, err)}
		}
	}
	if _, ok := ecc.mutation.PasswordHash(); !ok {
		return &ValidationError{Name: "password_hash", err: errors.New(`entgenerated: missing required field "EmailCredential.password_hash"`)}
	}
	if _, ok := ecc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`entgenerated: missing required edge "EmailCredential.owner"`)}
	}
	return nil
}

func (ecc *EmailCredentialCreate) sqlSave(ctx context.Context) (*EmailCredential, error) {
	if err := ecc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ecc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ecc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ecc.mutation.id = &_node.ID
	ecc.mutation.done = true
	return _node, nil
}

func (ecc *EmailCredentialCreate) createSpec() (*EmailCredential, *sqlgraph.CreateSpec) {
	var (
		_node = &EmailCredential{config: ecc.config}
		_spec = sqlgraph.NewCreateSpec(emailcredential.Table, sqlgraph.NewFieldSpec(emailcredential.FieldID, field.TypeInt))
	)
	if value, ok := ecc.mutation.Email(); ok {
		_spec.SetField(emailcredential.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := ecc.mutation.Algorithm(); ok {
		_spec.SetField(emailcredential.FieldAlgorithm, field.TypeEnum, value)
		_node.Algorithm = value
	}
	if value, ok := ecc.mutation.PasswordHash(); ok {
		_spec.SetField(emailcredential.FieldPasswordHash, field.TypeBytes, value)
		_node.PasswordHash = value
	}
	if nodes := ecc.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EmailCredentialCreateBulk is the builder for creating many EmailCredential entities in bulk.
type EmailCredentialCreateBulk struct {
	config
	err      error
	builders []*EmailCredentialCreate
}

// Save creates the EmailCredential entities in the database.
func (eccb *EmailCredentialCreateBulk) Save(ctx context.Context) ([]*EmailCredential, error) {
	if eccb.err != nil {
		return nil, eccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(eccb.builders))
	nodes := make([]*EmailCredential, len(eccb.builders))
	mutators := make([]Mutator, len(eccb.builders))
	for i := range eccb.builders {
		func(i int, root context.Context) {
			builder := eccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmailCredentialMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, eccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, eccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eccb *EmailCredentialCreateBulk) SaveX(ctx context.Context) []*EmailCredential {
	v, err := eccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eccb *EmailCredentialCreateBulk) Exec(ctx context.Context) error {
	_, err := eccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eccb *EmailCredentialCreateBulk) ExecX(ctx context.Context) {
	if err := eccb.Exec(ctx); err != nil {
		panic(err)
	}
}