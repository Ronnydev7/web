package schema

import (
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type UserOwnedMixin struct {
	mixin.Schema
	Ref      string
	Optional bool
}

func (userOwned UserOwnedMixin) Fields() []ent.Field {
	ownerIdField := field.Int("owner_id").Annotations(
		entgql.Skip(),
	)
	if userOwned.Optional {
		ownerIdField.Optional()
	}
	return []ent.Field{
		ownerIdField,
	}
}

func (userOwned UserOwnedMixin) Edges() []ent.Edge {
	if userOwned.Ref == "" {
		panic(errors.New("ref must be non-empty string"))
	}

	ownerEdge := edge.
		From("owner", User.Type).
		Field("owner_id").
		Ref(userOwned.Ref).
		Unique()

	if !userOwned.Optional {
		ownerEdge.Required()
	}

	return []ent.Edge{
		ownerEdge,
	}
}
