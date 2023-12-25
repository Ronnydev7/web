package rule

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/privacy"
	"api/ent/privacy/token"
	"context"
)

type MutationEmailGetter func(entgenerated.Mutation) (string, error)

func AllowMutationIfContextHasValidEmailSignupToken(getEmail MutationEmailGetter) privacy.MutationRule {
	return privacy.MutationRuleFunc(
		func(ctx context.Context, mutation entgenerated.Mutation) error {
			emailSignupToken := token.EmailSignupTokenFromContext(ctx)
			if emailSignupToken == nil {
				return privacy.Skipf("email signup token not found in context")
			}

			email, err := getEmail(mutation)
			if err != nil {
				return privacy.Skipf("unable to obtain email from mutation")
			}
			if email != emailSignupToken.Email {
				return privacy.Skipf("email sign up token does not match mutation result")
			}
			return privacy.Allow
		},
	)
}
