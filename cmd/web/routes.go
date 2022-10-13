package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// Static file serving routers.
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	// Snippet routers.
	router.Handler(http.MethodGet, "/", app.sessionManager.LoadAndSave(http.HandlerFunc(app.home)))
	router.Handler(http.MethodGet, "/snippet/view/:id", app.sessionManager.LoadAndSave(http.HandlerFunc(app.snippetView)))
	router.Handler(http.MethodGet, "/snippet/create", app.sessionManager.LoadAndSave(app.requireAuthentication(http.HandlerFunc(app.snippetCreate))))
	router.Handler(http.MethodPost, "/snippet/create", app.sessionManager.LoadAndSave(app.requireAuthentication(http.HandlerFunc(app.snippetCreatePost))))
	// User routers.
	router.Handler(http.MethodGet, "/user/login", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userLogin)))
	router.Handler(http.MethodPost, "/user/login", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userLoginPost)))
	router.Handler(http.MethodPost, "/user/logout", app.sessionManager.LoadAndSave(app.requireAuthentication(http.HandlerFunc(app.userLogoutPost))))
	router.Handler(http.MethodGet, "/user/signup", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userSignup)))
	router.Handler(http.MethodPost, "/user/signup", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userSignupPost)))

	// First RecoverPanic then log request and then secure headers
	// then router handlers.
	return app.recoverPanic(app.logRequest(secureHeaders(router)))
}
