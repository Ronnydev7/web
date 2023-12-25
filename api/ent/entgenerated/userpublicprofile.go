// Code generated by ent, DO NOT EDIT.

package entgenerated

import (
	"api/ent/entgenerated/user"
	"api/ent/entgenerated/userpublicprofile"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserPublicProfile is the model entity for the UserPublicProfile schema.
type UserPublicProfile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID int `json:"owner_id,omitempty"`
	// HandleName holds the value of the "handle_name" field.
	HandleName string `json:"handle_name,omitempty"`
	// PhotoBlobKey holds the value of the "photo_blob_key" field.
	PhotoBlobKey string `json:"photo_blob_key,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserPublicProfileQuery when eager-loading is set.
	Edges        UserPublicProfileEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserPublicProfileEdges holds the relations/edges for other nodes in the graph.
type UserPublicProfileEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserPublicProfileEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserPublicProfile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userpublicprofile.FieldID, userpublicprofile.FieldOwnerID:
			values[i] = new(sql.NullInt64)
		case userpublicprofile.FieldHandleName, userpublicprofile.FieldPhotoBlobKey:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserPublicProfile fields.
func (upp *UserPublicProfile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userpublicprofile.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			upp.ID = int(value.Int64)
		case userpublicprofile.FieldOwnerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				upp.OwnerID = int(value.Int64)
			}
		case userpublicprofile.FieldHandleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field handle_name", values[i])
			} else if value.Valid {
				upp.HandleName = value.String
			}
		case userpublicprofile.FieldPhotoBlobKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo_blob_key", values[i])
			} else if value.Valid {
				upp.PhotoBlobKey = value.String
			}
		default:
			upp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserPublicProfile.
// This includes values selected through modifiers, order, etc.
func (upp *UserPublicProfile) Value(name string) (ent.Value, error) {
	return upp.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the UserPublicProfile entity.
func (upp *UserPublicProfile) QueryOwner() *UserQuery {
	return NewUserPublicProfileClient(upp.config).QueryOwner(upp)
}

// Update returns a builder for updating this UserPublicProfile.
// Note that you need to call UserPublicProfile.Unwrap() before calling this method if this UserPublicProfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (upp *UserPublicProfile) Update() *UserPublicProfileUpdateOne {
	return NewUserPublicProfileClient(upp.config).UpdateOne(upp)
}

// Unwrap unwraps the UserPublicProfile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (upp *UserPublicProfile) Unwrap() *UserPublicProfile {
	_tx, ok := upp.config.driver.(*txDriver)
	if !ok {
		panic("entgenerated: UserPublicProfile is not a transactional entity")
	}
	upp.config.driver = _tx.drv
	return upp
}

// String implements the fmt.Stringer.
func (upp *UserPublicProfile) String() string {
	var builder strings.Builder
	builder.WriteString("UserPublicProfile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", upp.ID))
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", upp.OwnerID))
	builder.WriteString(", ")
	builder.WriteString("handle_name=")
	builder.WriteString(upp.HandleName)
	builder.WriteString(", ")
	builder.WriteString("photo_blob_key=")
	builder.WriteString(upp.PhotoBlobKey)
	builder.WriteByte(')')
	return builder.String()
}

// UserPublicProfiles is a parsable slice of UserPublicProfile.
type UserPublicProfiles []*UserPublicProfile
