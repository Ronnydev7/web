//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("graphql/schemas/ent.graphqls"),
		entgql.WithConfigPath("gqlgen.yml"),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(
			ex,
		),
		entc.FeatureNames("privacy", "entql", "schema/snapshot"),
	}

	if err := entc.Generate(
		"./ent/schema",
		&gen.Config{
			Target:  "ent/entgenerated",
			Package: "api/ent/entgenerated",
		},
		opts...,
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
