package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"api/config"
	"api/ent/entgenerated"
	"api/graphql/gqlgenerated"
	"api/graphql/resolverlib"
	"api/intl/intlgenerated"
	"api/lib"
	"api/privacy/viewer"
	"context"
	"errors"
	"net/http"

	"entgo.io/ent/entql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	cookieReader, err := lib.RequireCookieReaderFromContext(ctx)
	if err != nil {
		return resolverlib.UnknownLogoutDeath(ctx, err)
	}

	httpErrorWriter, err := lib.RequireHttpErrorWriterFromContext(ctx)
	if err != nil {
		return resolverlib.UnknownLogoutDeath(ctx, err)
	}

	userTokenFactory := lib.NewUserTokenFactory()
	loginSessionFactory := lib.NewLoginSessionFactory(ctx, r.client)
	loginSession, err := loginSessionFactory.FromRefreshTokenCookie(cookieReader, userTokenFactory)
	if err != nil {
		return resolverlib.UnauthorizedLogoutCall(ctx, httpErrorWriter, err)
	}

	err = r.client.LoginSession.DeleteOne(loginSession).Exec(ctx)
	if err != nil {
		return resolverlib.UnauthorizedLogoutCall(ctx, httpErrorWriter, err)
	}
	return true, nil
}

// CreateProfilePhotoUploadURL is the resolver for the createProfilePhotoUploadUrl field.
func (r *mutationResolver) CreateProfilePhotoUploadURL(ctx context.Context, userID int, md5 string) (*string, error) {
	profileQuery := r.client.UserPublicProfile.Query()
	profileQuery.Filter().WhereOwnerID(entql.IntEQ(userID))
	profile, err := profileQuery.Only(ctx)
	if err != nil {
		if !entgenerated.IsNotFound(err) {
			resolverlib.NewUserResolverLogger().LogError(err)
		}
		return nil, graphql.ErrorOnPath(ctx, errors.New(intlgenerated.BLOB_STORAGE__UNABLE_TO_CREATE_UPLOAD_URL))
	}

	photoBlobKeyString := profile.PhotoBlobKey
	if photoBlobKeyString == "" {
		photoBlobKey, err := uuid.NewRandom()
		if err != nil {
			resolverlib.NewUserResolverLogger().LogError(err)
			return nil, graphql.ErrorOnPath(ctx, errors.New(intlgenerated.BLOB_STORAGE__UNABLE_TO_CREATE_UPLOAD_URL))
		}
		photoBlobKeyString = photoBlobKey.String()
	}

	// Even if there is an existing blob key, call the mutation as a way to check for permissions
	profile, err = profile.Update().SetPhotoBlobKey(photoBlobKeyString).Save(ctx)
	if err != nil {
		resolverlib.NewUserResolverLogger().LogError(err)
		return nil, graphql.ErrorOnPath(ctx, errors.New(intlgenerated.BLOB_STORAGE__UNABLE_TO_CREATE_UPLOAD_URL))
	}

	storage := lib.NewBlobStorage(&lib.BlobStorageConfig{
		AwsConfig: config.GetAwsConfig(),
	})
	url, err := storage.GetSignedExternalMediaUploadUrl(&lib.BlobUploadSpec{
		Key: profile.PhotoBlobKey,
		Md5: md5,
	})
	if err != nil {
		resolverlib.NewUserResolverLogger().LogError(err)
		return nil, graphql.ErrorOnPath(ctx, errors.New(intlgenerated.BLOB_STORAGE__UNABLE_TO_CREATE_UPLOAD_URL))
	}

	return &url, nil
}

// Viewer is the resolver for the viewer field.
func (r *queryResolver) Viewer(ctx context.Context) (*entgenerated.User, error) {
	v := viewer.FromContext(ctx)
	if v != nil {
		id, exist := v.GetId()
		if exist {
			user, err := r.client.User.Get(ctx, id)
			if entgenerated.IsNotFound(err) {
				return nil, nil
			} else if err != nil {
				resolverlib.NewUserResolverLogger().LogError(err)
				return nil, graphql.ErrorOnPath(ctx, errors.New(intlgenerated.COMMON_STRINGS__UNKNOWN_SERVER_ERROR))
			}
			return user, nil
		}
	}
	return nil, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*entgenerated.User, error) {
	return r.client.User.Get(ctx, id)
}

// RefreshUserToken is the resolver for the refreshUserToken field.
func (r *queryResolver) RefreshUserToken(ctx context.Context) (*gqlgenerated.UserToken, error) {
	cookieReader, err := lib.RequireCookieReaderFromContext(ctx)
	if err != nil {
		return resolverlib.UnknownRefreshUserTokenDeath(ctx, err)
	}
	httpErrorWriter, err := lib.RequireHttpErrorWriterFromContext(ctx)
	if err != nil {
		return resolverlib.UnknownRefreshUserTokenDeath(ctx, err)
	}

	refreshTokenCookie, err := cookieReader.ReadRefreshToken()
	if err == http.ErrNoCookie {
		return nil, nil
	} else if err != nil {
		return resolverlib.UnauthorizedRefreshUserTokenCall(ctx, httpErrorWriter, nil)
	}

	userTokenFactory := lib.NewUserTokenFactory()
	loginSession, err := resolverlib.GetLoginSession(ctx, r.client, cookieReader, userTokenFactory)
	if err != nil {
		return resolverlib.UnauthorizedRefreshUserTokenCall(ctx, httpErrorWriter, nil)
	}
	owner, err := r.client.LoginSession.QueryOwner(loginSession).Only(ctx)
	if err != nil {
		return resolverlib.UnauthorizedRefreshUserTokenCall(ctx, httpErrorWriter, nil)
	}

	authToken, err := userTokenFactory.CreateAuthToken(owner)
	if err != nil {
		return resolverlib.UnauthorizedRefreshUserTokenCall(ctx, httpErrorWriter, err)
	}

	return &gqlgenerated.UserToken{
		AuthToken:    authToken,
		RefreshToken: refreshTokenCookie.Value,
	}, nil
}