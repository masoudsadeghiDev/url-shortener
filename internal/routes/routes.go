package routes

import "net/http"

func NewRoutes() http.Handler {

	mux := http.NewServeMux()

	// Serve static files (React build)
	fs := http.FileServer(http.Dir("ui/build"))
	mux.Handle("/", fs)

	// Short routes
	mux.HandleFunc("/short/create", CreateShortUrl)
	mux.HandleFunc("/short", RedirectUrl)

	return mux
}
