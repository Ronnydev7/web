package schema

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/emailcredential"
	"api/ent/entgenerated/privacy"
	"api/ent/entgenerated/user"
	"api/ent/privacy/rule"
	"api/ent/privacy/token"

	"entgo.io/ent"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

type LoginSession struct {
	ent.Schema
}

func (LoginSession) Fields() []ent.Field {
	return []ent.Field{
		field.Time("last_login_time"),
	}
}

func (LoginSession) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UserOwnedMixin{
			Ref: "login_sessions",
		},
		UserOwnedMutationPolicyMixin{
			AllowAdminMutation: true,
		},
	}
}

func (LoginSession) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.AllowIfAdmin(),
			rule.AllowAfterApplyingPrivacyTokenFilter(
				&token.PerformingLoginToken{},
				func(t token.PrivacyToken, filter privacy.Filter) {
					actualToken := t.(*token.PerformingLoginToken)
					loginSessionFilter := filter.(*entgenerated.LoginSessionFilter)
					loginSessionFilter.WhereHasOwnerWith(
						user.HasEmailCredentialWith(
							emailcredential.EmailEQ(actualToken.Email),
						),
					)
				},
			),
			rule.AllowAfterApplyingPrivacyTokenFilter(
				&token.PerformingAuthRefreshToken{},
				func(t token.PrivacyToken, filter privacy.Filter) {
					actualToken := t.(*token.PerformingAuthRefreshToken)
					loginSessionFilter := filter.(*entgenerated.LoginSessionFilter)
					loginSessionFilter.WhereID(entql.IntEQ(actualToken.LoginSessionId))
				},
			),
			rule.AllowIfOwnedByViewer(),
			privacy.AlwaysDenyRule(),
		},
	}
}
