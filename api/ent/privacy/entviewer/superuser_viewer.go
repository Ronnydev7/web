package entviewer

import (
	"api/ent/entgenerated"
	"api/privacy/viewer"
	"context"
	"fmt"
)

type (
	SuperuserViewer struct {
		viewer.Viewer
		User             *entgenerated.User
		SuperuserProfile *entgenerated.SuperuserProfile
	}

	UnauthorizedSuperuserRequestError struct {
		error
	}

	UnknownSuperuserRetrievalError struct {
		error
		Inner error
	}

	NewSuperuserViewerFunc = func(*entgenerated.User) (viewer.Viewer, error)
)

var NewSuperuserViewerFromUser NewSuperuserViewerFunc = func(
	user *entgenerated.User,
) (viewer.Viewer, error) {
	ctx := viewer.NewContext(context.Background(), NewUserViewerFromUser(user))
	superuserProfile, err := user.QuerySuperuserProfile().Only(ctx)
	if err != nil {
		return nil, &UnknownSuperuserRetrievalError{
			Inner: err,
		}
	}
	if superuserProfile == nil {
		return nil, &UnauthorizedSuperuserRequestError{}
	}

	return &SuperuserViewer{
		User:             user,
		SuperuserProfile: superuserProfile,
	}, nil
}

func (v SuperuserViewer) GetId() (id int, exists bool) {
	if v.User == nil {
		return 0, false
	}
	return v.User.ID, true
}

func (v SuperuserViewer) IsAdmin() bool {
	return v.IsSuperuser()
}

func (v SuperuserViewer) IsSuperuser() bool {
	return v.SuperuserProfile != nil
}

func (UnauthorizedSuperuserRequestError) Error() string {
	return "unable to find superuser profile for user"
}

func (err UnknownSuperuserRetrievalError) Error() string {
	return fmt.Sprintf("error loading superuser profile: %v", err.Inner)
}
