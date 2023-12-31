// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EmailCredentialsColumns holds the columns for the "email_credentials" table.
	EmailCredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "algorithm", Type: field.TypeEnum, Enums: []string{"bcrypt"}},
		{Name: "password_hash", Type: field.TypeBytes},
		{Name: "owner_id", Type: field.TypeInt, Unique: true},
	}
	// EmailCredentialsTable holds the schema information for the "email_credentials" table.
	EmailCredentialsTable = &schema.Table{
		Name:       "email_credentials",
		Columns:    EmailCredentialsColumns,
		PrimaryKey: []*schema.Column{EmailCredentialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "email_credentials_users_email_credential",
				Columns:    []*schema.Column{EmailCredentialsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// LoginSessionsColumns holds the columns for the "login_sessions" table.
	LoginSessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "last_login_time", Type: field.TypeTime},
		{Name: "owner_id", Type: field.TypeInt},
	}
	// LoginSessionsTable holds the schema information for the "login_sessions" table.
	LoginSessionsTable = &schema.Table{
		Name:       "login_sessions",
		Columns:    LoginSessionsColumns,
		PrimaryKey: []*schema.Column{LoginSessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "login_sessions_users_login_sessions",
				Columns:    []*schema.Column{LoginSessionsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SuperuserProfilesColumns holds the columns for the "superuser_profiles" table.
	SuperuserProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "owner_id", Type: field.TypeInt, Unique: true},
	}
	// SuperuserProfilesTable holds the schema information for the "superuser_profiles" table.
	SuperuserProfilesTable = &schema.Table{
		Name:       "superuser_profiles",
		Columns:    SuperuserProfilesColumns,
		PrimaryKey: []*schema.Column{SuperuserProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "superuser_profiles_users_superuser_profile",
				Columns:    []*schema.Column{SuperuserProfilesColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserPublicProfilesColumns holds the columns for the "user_public_profiles" table.
	UserPublicProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "handle_name", Type: field.TypeString, Unique: true},
		{Name: "photo_blob_key", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "owner_id", Type: field.TypeInt, Unique: true},
	}
	// UserPublicProfilesTable holds the schema information for the "user_public_profiles" table.
	UserPublicProfilesTable = &schema.Table{
		Name:       "user_public_profiles",
		Columns:    UserPublicProfilesColumns,
		PrimaryKey: []*schema.Column{UserPublicProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_public_profiles_users_public_profile",
				Columns:    []*schema.Column{UserPublicProfilesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EmailCredentialsTable,
		LoginSessionsTable,
		SuperuserProfilesTable,
		UsersTable,
		UserPublicProfilesTable,
	}
)

func init() {
	EmailCredentialsTable.ForeignKeys[0].RefTable = UsersTable
	LoginSessionsTable.ForeignKeys[0].RefTable = UsersTable
	SuperuserProfilesTable.ForeignKeys[0].RefTable = UsersTable
	UserPublicProfilesTable.ForeignKeys[0].RefTable = UsersTable
}
