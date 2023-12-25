package viewer

import "context"

type Viewer interface {
	GetId() (id int, exists bool)
	IsAdmin() bool
	IsSuperuser() bool
}

type ctxKey struct{}

func FromContext(ctx context.Context) Viewer {
	v, _ := ctx.Value(ctxKey{}).(Viewer)
	return v
}

func NewContext(parent context.Context, v Viewer) context.Context {
	return context.WithValue(parent, ctxKey{}, v)
}
