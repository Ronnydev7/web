// Code generated by ent, DO NOT EDIT.

package entgenerated

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	SuperuserProfileID *int
	EmailCredentialID  *int
	LoginSessionIDs    []int
	PublicProfileID    *int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	if v := i.SuperuserProfileID; v != nil {
		m.SetSuperuserProfileID(*v)
	}
	if v := i.EmailCredentialID; v != nil {
		m.SetEmailCredentialID(*v)
	}
	if v := i.LoginSessionIDs; len(v) > 0 {
		m.AddLoginSessionIDs(v...)
	}
	if v := i.PublicProfileID; v != nil {
		m.SetPublicProfileID(*v)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}