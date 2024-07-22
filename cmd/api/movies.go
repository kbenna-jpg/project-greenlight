package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.kbennani.net/internal/data"
	// "greenlight.kbennani.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new movie")
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