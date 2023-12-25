package entviewer

import (
	"api/ent/entgenerated"
	"api/privacy/viewer"
)

type UserViewer struct {
	viewer.Viewer
	id    int
	hasId bool
}

func NewUserViewerFromUser(user *entgenerated.User) viewer.Viewer {
	if user == nil {
		return NewUserViewerFromId(0, false)
	}

	return NewUserViewerFromId(user.ID, true)
}

func NewUserViewerFromId(id int, hasId bool) viewer.Viewer {
	return &UserViewer{
		id:    id,
		hasId: hasId,
	}
}

func NewAnonymouseUserViewer() viewer.Viewer {
	return NewUserViewerFromUser(nil)
}

func (viewer UserViewer) GetId() (int, bool) {
	return viewer.id, viewer.hasId
}

func (viewer UserViewer) IsAdmin() bool {
	return false
}

func (viewer UserViewer) IsSuperuser() bool {
	return false
}
