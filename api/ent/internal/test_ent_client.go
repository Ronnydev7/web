package internal

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/enttest"
	"api/ent/entgenerated/migrate"
	"api/privacy/viewer"
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

type bootstrapViewer struct {
	viewer.Viewer
	user *entgenerated.User
}

func (viewer bootstrapViewer) GetId() (int, bool) {
	if viewer.user == nil {
		return 0, false
	}
	return viewer.user.ID, true
}

func (viewer bootstrapViewer) IsAdmin() bool {
	return true
}

func (viewer bootstrapViewer) IsSuperuser() bool {
	return true
}

func CreateEntClientForTest(t *testing.T) *entgenerated.Client {
	opts := []enttest.Option{
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	return enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
}

func CreateTestUser(client *entgenerated.Client) *entgenerated.User {
	context := viewer.NewContext(context.Background(), bootstrapViewer{})
	user, err := client.User.Create().Save(context)
	if err != nil {
		panic(err)
	}
	return user
}

func CreateTestSuperuser(client *entgenerated.Client, name string) (*entgenerated.User, *entgenerated.SuperuserProfile) {
	user := CreateTestUser(client)
	context := viewer.NewContext(context.Background(), bootstrapViewer{user: user})
	superuserProfile, err := client.SuperuserProfile.Create().SetOwner(user).Save(context)
	if err != nil {
		panic(err)
	}

	return client.User.GetX(context, user.ID), superuserProfile
}

func CreateTestEmailCredentialForUser(client *entgenerated.Client, user *entgenerated.User, address string) *entgenerated.EmailCredential {
	context := viewer.NewContext(context.Background(), bootstrapViewer{})
	emailCredential, err := client.EmailCredential.
		Create().
		SetEmail(address).
		SetAlgorithm("bcrypt").
		SetPasswordHash([]byte("hash")).
		SetOwner(user).
		Save(context)
	if err != nil {
		panic(err)
	}
	return emailCredential
}
