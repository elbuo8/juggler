package middlewares

import (
	"context"
	"github.com/elbuo8/juggler/app"
	"github.com/urfave/negroni"
	"net/http"
)

func Authorize(app *app.App) negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		ctx := r.Context()
		context.WithValue(ctx, "user", "user")
		next(w, r.WithContext(ctx))
	})
}
