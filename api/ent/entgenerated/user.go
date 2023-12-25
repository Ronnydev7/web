// Code generated by ent, DO NOT EDIT.

package entgenerated

import (
	"api/ent/entgenerated/emailcredential"
	"api/ent/entgenerated/superuserprofile"
	"api/ent/entgenerated/user"
	"api/ent/entgenerated/userpublicprofile"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// SuperuserProfile holds the value of the superuser_profile edge.
	SuperuserProfile *SuperuserProfile `json:"superuser_profile,omitempty"`
	// EmailCredential holds the value of the email_credential edge.
	EmailCredential *EmailCredential `json:"email_credential,omitempty"`
	// LoginSessions holds the value of the login_sessions edge.
	LoginSessions []*LoginSession `json:"login_sessions,omitempty"`
	// PublicProfile holds the value of the public_profile edge.
	PublicProfile *UserPublicProfile `json:"public_profile,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedLoginSessions map[string][]*LoginSession
}

// SuperuserProfileOrErr returns the SuperuserProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SuperuserProfileOrErr() (*SuperuserProfile, error) {
	if e.loadedTypes[0] {
		if e.SuperuserProfile == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: superuserprofile.Label}
		}
		return e.SuperuserProfile, nil
	}
	return nil, &NotLoadedError{edge: "superuser_profile"}
}

// EmailCredentialOrErr returns the EmailCredential value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) EmailCredentialOrErr() (*EmailCredential, error) {
	if e.loadedTypes[1] {
		if e.EmailCredential == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: emailcredential.Label}
		}
		return e.EmailCredential, nil
	}
	return nil, &NotLoadedError{edge: "email_credential"}
}

// LoginSessionsOrErr returns the LoginSessions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LoginSessionsOrErr() ([]*LoginSession, error) {
	if e.loadedTypes[2] {
		return e.LoginSessions, nil
	}
	return nil, &NotLoadedError{edge: "login_sessions"}
}

// PublicProfileOrErr returns the PublicProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) PublicProfileOrErr() (*UserPublicProfile, error) {
	if e.loadedTypes[3] {
		if e.PublicProfile == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: userpublicprofile.Label}
		}
		return e.PublicProfile, nil
	}
	return nil, &NotLoadedError{edge: "public_profile"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QuerySuperuserProfile queries the "superuser_profile" edge of the User entity.
func (u *User) QuerySuperuserProfile() *SuperuserProfileQuery {
	return NewUserClient(u.config).QuerySuperuserProfile(u)
}

// QueryEmailCredential queries the "email_credential" edge of the User entity.
func (u *User) QueryEmailCredential() *EmailCredentialQuery {
	return NewUserClient(u.config).QueryEmailCredential(u)
}

// QueryLoginSessions queries the "login_sessions" edge of the User entity.
func (u *User) QueryLoginSessions() *LoginSessionQuery {
	return NewUserClient(u.config).QueryLoginSessions(u)
}

// QueryPublicProfile queries the "public_profile" edge of the User entity.
func (u *User) QueryPublicProfile() *UserPublicProfileQuery {
	return NewUserClient(u.config).QueryPublicProfile(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("entgenerated: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteByte(')')
	return builder.String()
}

// NamedLoginSessions returns the LoginSessions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedLoginSessions(name string) ([]*LoginSession, error) {
	if u.Edges.namedLoginSessions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedLoginSessions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedLoginSessions(name string, edges ...*LoginSession) {
	if u.Edges.namedLoginSessions == nil {
		u.Edges.namedLoginSessions = make(map[string][]*LoginSession)
	}
	if len(edges) == 0 {
		u.Edges.namedLoginSessions[name] = []*LoginSession{}
	} else {
		u.Edges.namedLoginSessions[name] = append(u.Edges.namedLoginSessions[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User