package test

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/enttest"
	"api/ent/entgenerated/migrate"
	"api/ent/privacy/entviewer"
	"api/graphql"
	"api/intl"
	"api/lib"
	"api/lib/testutils"
	"api/privacy/viewer"
	"context"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

func startValidNewEntClientMock() (*entgenerated.Client, testutils.MockWrapper[lib.NewEntClientFunc]) {
	client := createEntClientForTest()
	return client, testutils.StartMockWrapper[lib.NewEntClientFunc](
		func() (*entgenerated.Client, intl.IntlError) {
			return client, nil
		},
		lib.NewEntClient,
		func(aNew lib.NewEntClientFunc) {
			lib.NewEntClient = aNew
		},
	)
}

func createTestUser(client *entgenerated.Client) *entgenerated.User {
	context := viewer.NewContext(context.Background(), testViewer{})
	return client.User.Create().SaveX(context)
}

func createTestUserViewerContext(client *entgenerated.Client) context.Context {
	userViewer := entviewer.NewUserViewerFromUser(createTestUser(client))
	return viewer.NewContext(context.Background(), userViewer)
}

func createTestAdminViewer() viewer.Viewer {
	return &testViewer{}
}

func createGraphqlClient(entClient *entgenerated.Client) *client.Client {
	httpHandler := handler.NewDefaultServer(graphql.NewSchema(entClient))
	return client.New(httpHandler)
}

func requestGraphQLWithContext(applyContext func(parent context.Context) context.Context) client.Option {
	return func(bd *client.Request) {
		ctx := applyContext(bd.HTTP.Context())
		bd.HTTP = bd.HTTP.WithContext(ctx)
	}
}

func TestLib(t *testing.T) {
	TheT = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graphql Suite")
}
