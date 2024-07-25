package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"greenlight.kbennani.net/internal/data"
	// "greenlight.kbennani.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Declare anonymous struct to hold information expected to be in HTTP request body
	var input struct {
		Title string `json: "title"`
		Year int32 `json: "year"`
		Runtime int32 `json: "runtime"`
		Genres []string `json: "genres"`
	}

		err := app.readJSON(w, r, &input)
		if err != nil {
			app.errorResponse(w, r, http.StatusBadRequest, err.Error())
			return
		}
	// Reads from request body and Decode() method decodes body contents into input struct
	err := json.NewDecoder(r.Body).Decode(&input) // pointer passed to input struct as target decode destination
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

// When httprouter is parsing a reuquest, ay interpolated URL params stored in request context
// Retrieves a slice containing param names and values
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)

	// Create new instance of Movie struct containing id extracted from url
	movie := data.Movie{
		ID: id,
		CreatedAt: time.Now(),
		Title: "Casblanca",
		Runtime: 102,
		Genres: []string{"drama", "romance", "war"},
		Version: 1,
	}

	// Edncode struct to JSON and send as HTTP response
	err = app.writeJSON(w, http.StatusOK, envelope{"movie":movie}, nil)
	if err != nil{
		app.ServerErrorResponse(w, r, err)
	}

}