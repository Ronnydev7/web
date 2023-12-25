package lib

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

type (
	GraphQLRefreshAuthTokenExtension struct {
		graphql.HandlerExtension
		graphql.ResponseInterceptor
	}

	graphqlAuthTokenExtensionContextKey struct{}

	NewGraphQLRefreshAuthTokenExtensionFunc func() *GraphQLRefreshAuthTokenExtension
)

const AUTH_TOKEN_PAYLOAD_EXTENSION_KEY = "authToken"

var NewGraphQLRefreshAuthTokenExtension NewGraphQLRefreshAuthTokenExtensionFunc = func() *GraphQLRefreshAuthTokenExtension {
	return &GraphQLRefreshAuthTokenExtension{}
}

func (GraphQLRefreshAuthTokenExtension) ExtensionName() string {
	return "GraphQLRefreshAuthTokenExtension"
}

func (GraphQLRefreshAuthTokenExtension) Validate(_ graphql.ExecutableSchema) error {
	return nil
}

func (GraphQLRefreshAuthTokenExtension) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	authToken := GetAuthTokenFromContext(ctx)
	response := next(ctx)

	if authToken != nil {
		response.Extensions[AUTH_TOKEN_PAYLOAD_EXTENSION_KEY] = *authToken
	}

	return response
}

func NewContextWithAuthToken(parent context.Context, authToken string) context.Context {
	return context.WithValue(parent, graphqlAuthTokenExtensionContextKey{}, authToken)
}

func GetAuthTokenFromContext(ctx context.Context) *string {
	value := ctx.Value(graphqlAuthTokenExtensionContextKey{})
	if value == nil {
		return nil
	}

	authToken, ok := value.(string)
	if ok {
		return &authToken
	}
	return nil
}
