package test

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/enttest"
	"api/ent/entgenerated/migrate"
	"api/graphql"
	"api/lib"
	"api/lib/libmocks"
	"api/privacy/viewer"
	"context"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

type testViewer struct {
	viewer.Viewer
}

func (testViewer) IsAdmin() bool {
	return true
}

var TheT *testing.T

func createEntClientForTest() *entgenerated.Client {
	opts := []enttest.Option{
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	return enttest.Open(TheT, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
}

func createTestUser(client *entgenerated.Client) *entgenerated.User {
	context := viewer.NewContext(context.Background(), &testViewer{})
	return client.User.Create().SaveX(context)
}

func CreateTestEmailCredentialForUser(
	client *entgenerated.Client,
	user *entgenerated.User,
	address string,
) *entgenerated.EmailCredential {
	context := viewer.NewContext(context.Background(), &testViewer{})
	return client.EmailCredential.
		Create().
		SetEmail(address).
		SetAlgorithm("bcrypt").
		SetPasswordHash([]byte("password")).
		SetOwner(user).
		SaveX(context)
}

func createGraphqlClient(entClient *entgenerated.Client) *client.Client {
	httpHandler := handler.NewDefaultServer(graphql.NewSchema(entClient))
	return client.New(httpHandler)
}

func requestWithContext(applyContext func(parent context.Context) context.Context) client.Option {
	return func(bd *client.Request) {
		ctx := applyContext(bd.HTTP.Context())
		bd.HTTP = bd.HTTP.WithContext(ctx)
	}
}

func TestGraphql(t *testing.T) {
	TheT = t
	RegisterFailHandler(Fail)
	lib.NewLogger = func(name string) lib.Logger {
		result := &libmocks.Logger{}
		result.On("LogError", mock.Anything).Return(nil)
		return result
	}
	RunSpecs(t, "Graphql Suite")
}
