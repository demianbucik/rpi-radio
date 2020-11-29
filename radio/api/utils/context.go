package utils

// contextKey is a value for use with context.WithValue.
// It's used as a pointer so it fits in an interface{} without allocation.
// This technique for defining context keys was copied from Go 1.7's new use of context in net/http.
type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return k.name
}

var (
	RequestCtxKey       = &contextKey{name: "RequestContext"}
	PlaylistCtxKey      = &contextKey{name: "Playlist"}
	TrackCtxKey         = &contextKey{name: "Track"}
	PlaylistTrackCtxKey = &contextKey{name: "PlaylistTrack"}
)

type RequestContext struct {
	Error      interface{}
	Panic      interface{}
	StackTrace string
}
