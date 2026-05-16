package router

import "net/http"

func Register(handler) *http.ServeMux {
	mux := http.NewServeMux()

	RegisterUsersRoute(mux, handler)

	return mux
}
