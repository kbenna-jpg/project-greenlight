package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

)

// Need to extract id parameters from URL
	// Get value of "id" parameter from slice
	// All movies will have unique positive integer ID. but
	// value returned ByName() is always a string, so convert it to base 10 i
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context()) // retrieve "id" UR: param rom current request context 

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)  // convert to integer
	
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter") // error returns 0
	}
	return id, nil 

}
