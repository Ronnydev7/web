package main

import (
	"api/graphql"
	"api/lib"
	"api/privacy/viewer"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func main() {
	// Create ent.Client and run the schema migration.
	client, err := lib.NewEntClient()
	if err != nil {
		panic(fmt.Errorf("opening ent client: %+v", err))
	}
	defer client.Close()

	router := chi.NewRouter()

	router.Use(lib.ApplyViewerMiddleware(
		func(r *http.Request) viewer.Viewer {
			viewerFactory := lib.NewViewerFactory()
			return viewerFactory.FromHttpRequestOrElseAnonymous(client, r)
		},
	))
	router.Use(lib.ApplyCookieReaderWriterMiddleware())
	router.Use(lib.ApplyHttpErrorWriterMiddleware())

	// Define paths
	router.Handle("/playground", playground.Handler("Graphql Playground Server", "/query"))

	queryHandler := handler.NewDefaultServer(graphql.NewSchema(client))
	queryHandler.Use(lib.NewGraphQLRefreshAuthTokenExtension())
	corsConfig := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	router.Handle("/query", corsConfig.Handler(queryHandler))

	port := fmt.Sprintf(":%d", 8080)
	fmt.Printf("listening on %s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		panic(fmt.Errorf("http server terminated: %v", err))
	}
}
