package test

import (
	"api/ent/entgenerated"
	"api/graphql"
	"api/lib"
	"context"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("graphql_refresh_auth_token_extension", func() {
	Describe("GraphQLRefreshAuthTokenExtension", func() {
		var queryContext context.Context
		var subject lib.GraphQLRefreshAuthTokenExtension

		var entClient *entgenerated.Client
		var graphqlClient *client.Client

		BeforeEach(func() {
			queryContext = context.Background()
			subject = *lib.NewGraphQLRefreshAuthTokenExtension()

			entClient = createEntClientForTest()
			httpHandler := handler.NewDefaultServer(graphql.NewSchema(entClient))
			httpHandler.Use(subject)
			graphqlClient = client.New(httpHandler)
		})

		AfterEach(func() {
			entClient.Close()
		})

		doViewerQuery := func() (*client.Response, error) {
			return graphqlClient.RawPost(
				`
				{
					viewer {
						id
					}
				}
			`,
				requestGraphQLWithContext(func(parent context.Context) context.Context {
					return queryContext
				}),
			)
		}

		Context("when context has auth token", func() {
			BeforeEach(func() {
				queryContext = lib.NewContextWithAuthToken(queryContext, "test token")
			})

			It("set the extension field in the query response", func() {
				actual, err := doViewerQuery()
				Expect(err).To(BeNil())
				Expect(actual.Extensions).To(HaveKeyWithValue(lib.AUTH_TOKEN_PAYLOAD_EXTENSION_KEY, "test token"))
			})
		})

		Context("when context does not have auth token", func() {
			It("does not set the extension field in the query response", func() {
				actual, err := doViewerQuery()
				Expect(err).To(BeNil())
				Expect(actual.Extensions).NotTo(HaveKeyWithValue(lib.AUTH_TOKEN_PAYLOAD_EXTENSION_KEY, "test token"))
			})
		})
	})
})
