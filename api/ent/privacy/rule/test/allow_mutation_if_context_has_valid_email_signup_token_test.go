package test

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/privacy"
	"api/ent/internal"
	"api/ent/privacy/rule"
	"api/ent/privacy/token"
	"context"
	"errors"
	"fmt"
	"testing"
)

func setupAllowMutationIfContextHasValidEmailSignupTokenTest(t *testing.T) (*entgenerated.Client, rule.MutationEmailGetter) {
	client := internal.CreateEntClientForTest(t)
	type EmailCredentialMutation interface {
		Email() (email string, exists bool)
	}
	emailGetter := func(mutation entgenerated.Mutation) (string, error) {
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

	return client, emailGetter
}

func TestAllowMutationIfContextHasValidEmailSignupTokenAllowIfTokenIsValid(t *testing.T) {
	client, emailGetter := setupAllowMutationIfContextHasValidEmailSignupTokenTest(t)
	defer client.Close()
	emailMutation := client.EmailCredential.Create().SetEmail("test@example.com").Mutation()
	rule := rule.AllowMutationIfContextHasValidEmailSignupToken(emailGetter)
	emailSignupToken := token.EmailSignupToken{
		Email: "test@example.com",
	}
	ctx := token.NewContextWithSignupToken(context.Background(), &emailSignupToken)
	actual := rule.EvalMutation(ctx, emailMutation)
	internal.TestingMatchPrivacyDecision(t, actual, privacy.Allow)
}

func TestAllowMutationIfContextHasValidEmailSignupTokenSkipIfTokenNotFound(t *testing.T) {
	client, emailGetter := setupAllowMutationIfContextHasValidEmailSignupTokenTest(t)
	defer client.Close()
	emailMutation := client.EmailCredential.Create().SetEmail("test@example.com").Mutation()
	rule := rule.AllowMutationIfContextHasValidEmailSignupToken(emailGetter)
	actual := rule.EvalMutation(context.Background(), emailMutation)
	internal.TestingMatchPrivacyDecision(t, actual, privacy.Skip)
}

func TestAllowMutationIfContextHasValidEmailSignupTokenSkipIfEmailIsNotSetInMutation(t *testing.T) {
	client, emailGetter := setupAllowMutationIfContextHasValidEmailSignupTokenTest(t)
	defer client.Close()
	emailMutation := client.EmailCredential.Create().Mutation()
	rule := rule.AllowMutationIfContextHasValidEmailSignupToken(emailGetter)
	emailSignupToken := token.EmailSignupToken{
		Email: "test2@example.com",
	}
	ctx := token.NewContextWithSignupToken(context.Background(), &emailSignupToken)
	actual := rule.EvalMutation(ctx, emailMutation)
	internal.TestingMatchPrivacyDecision(t, actual, privacy.Skip)
}

func TestAllowMutationIfContextHasValidEmailSignupTokenSkipIfEmailIsSetButUnmatched(t *testing.T) {
	client, emailGetter := setupAllowMutationIfContextHasValidEmailSignupTokenTest(t)
	defer client.Close()
	emailMutation := client.EmailCredential.Create().SetEmail("test@example.com").Mutation()
	rule := rule.AllowMutationIfContextHasValidEmailSignupToken(emailGetter)
	emailSignupToken := token.EmailSignupToken{
		Email: "test2@example.com",
	}
	ctx := token.NewContextWithSignupToken(context.Background(), &emailSignupToken)
	actual := rule.EvalMutation(ctx, emailMutation)
	internal.TestingMatchPrivacyDecision(t, actual, privacy.Skip)
}
