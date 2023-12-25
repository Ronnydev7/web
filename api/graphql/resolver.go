package graphql

import (
	"api/ent/entgenerated"
	"api/graphql/gqlgenerated"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *entgenerated.Client
}

func NewSchema(client *entgenerated.Client) graphql.ExecutableSchema {
	return gqlgenerated.NewExecutableSchema(gqlgenerated.Config{
		Resolvers: &Resolver{client},
	})
}
