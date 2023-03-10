package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/Slimo300/Bookings/internal/config"
)

var app *config.AppConfig

// NewHelpers intitializes helpers package with app config
func NewHelpers(a *config.AppConfig) {
	app = a
}

// ClientError handles errors on client side
func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of ", status)
	http.Error(w, http.StatusText(status), status)
}

// ServerError handles errors on server side
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
