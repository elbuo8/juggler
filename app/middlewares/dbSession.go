package middlewares

import (
	"context"
	"github.com/elbuo8/juggler/app"
	"github.com/elbuo8/juggler/app/models"
	"github.com/urfave/negroni"
	"net/http"
	"os"
)

func SetDBSession(app *app.App) negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		dbSession := app.DBSession.Copy()
		db := models.DB{dbSession.DB(os.Getenv("DB_NAME"))}
		ctx := context.WithValue(r.Context(), "database", &db)
		next(w, r.WithContext(ctx))
		dbSession.Close()
	})
}
