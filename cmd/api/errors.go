package main

import (
	"fmt"
	"net/http"
)

// Generic helped for logging error msg
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// Genertic error for sending JSON-formatted error msgs to client with given status code
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	// If returns error, log and send client emoty response with Internal Server Error code
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

	// Used when applicaiton encounters unexpected error at runtime, logs detailed error msg and uses errorResponse() method to send 500 Internal Error c0de and JSON response to client
func (app *application) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// Sends 404 MEthod Not Allowed status code and JSON response to client
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// Sends 405 Method Not Allowed status code and JSON response to client
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this response", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}


