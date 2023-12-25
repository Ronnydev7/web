package schema

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/privacy"
	"api/ent/privacy/rule"
	"api/ent/privacy/token"
	"errors"
	"fmt"
	"regexp"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

type EmailCredential struct {
	ent.Schema
}

const emailRegexpStr = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])"

func (EmailCredential) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Match(regexp.MustCompile(emailRegexpStr)).
			Unique(),
		field.Enum("algorithm").Values(
			"bcrypt",
		).Annotations(
			entgql.Skip(),
		),
		field.Bytes("password_hash").Annotations(
			entgql.Skip(),
		),
	}
}

func (EmailCredential) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UserOwnedMixin{
			Ref: "email_credential",
		},
	}
}

func (EmailCredential) Policy() ent.Policy {
	emailGetter := func(mutation entgenerated.Mutation) (email string, err error) {
		type EmailCredentialMutation interface {
			Email() (email string, exists bool)
		}
		emailCredentialMutation, ok := mutation.(EmailCredentialMutation)
		if !ok {
			return "", fmt.Errorf("unexpected mutation type %T", mutation)
		}
		email, exists := emailCredentialMutation.Email()
		if !exists {
			return "", errors.New("email is not set")
		}
		return email, nil
	}

	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.AllowIfAdmin(),
			rule.AllowIfOwnedByViewer(),
			rule.AllowAfterApplyingPrivacyTokenFilter(
				&token.PerformingLoginToken{},
				func(t token.PrivacyToken, filter privacy.Filter) {
					actualToken := t.(*token.PerformingLoginToken)
					emailFilter := filter.(*entgenerated.EmailCredentialFilter)
					emailFilter.WhereEmail(entql.StringEQ(actualToken.Email))
				},
			),
			rule.AllowAfterApplyingPrivacyTokenFilter(
				&token.PerformingResetPasswordToken{},
				func(t token.PrivacyToken, filter privacy.Filter) {
					actualToken := t.(*token.PerformingResetPasswordToken)
					emailFilter := filter.(*entgenerated.EmailCredentialFilter)
					emailFilter.WhereEmail(entql.StringEQ(actualToken.Email))
				},
			),
			privacy.AlwaysDenyRule(),
		},
		Mutation: privacy.MutationPolicy{
			privacy.OnMutationOperation(
				privacy.MutationPolicy{
					rule.AllowIfAdmin(),
					rule.AllowMutationIfContextHasValidEmailSignupToken(emailGetter),
					privacy.AlwaysDenyRule(),
				},
				ent.OpCreate,
			),
			privacy.OnMutationOperation(
				privacy.MutationPolicy{
					rule.AllowIfAdmin(),
					rule.AllowMutationAfterApplyingOwnerFilter(),
					privacy.AlwaysDenyRule(),
				},
				ent.OpUpdateOne|ent.OpUpdate|ent.OpDeleteOne|ent.OpDelete,
			),
		},
	}
}
