package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

)

// Need to extract id parameters from URL
	// Get value of "id" parameter from slice
	// All movies will have unique positive integer ID. but
	// value returned ByName() is always a string, so convert it to base 10 i
func (app *application) readIDParam(r *http.Request) (int64, error) {
	// Decode request body into target destination 
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		// if error during decoding, start triage
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		
		switch {
		case errors.As(err, &syntaxError): // checks whether error has type *json,Syntax error, if so, return plain-english error msg
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.ls(err, io.ErrUnexpectedEOF): // if Decode() returns io.ErrUnexpectedEOF error
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Offset)

		case errors.ls(err, io.EOF):
			return errors.New(body must not be empty)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
			return err 
		
		}
	

	}
	params := httprouter.ParamsFromContext(r.Context()) // retrieve "id" UR: param rom current request context 

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)  // convert to integer
	
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter") // error returns 0
	}
	return id, nil 

}

type envelope map[string]interface{}

// Helper method which sends responses
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t") // adds whitespace to encoded JSON
	if err != nil {
		return err
	}
	js = append(js, '\n')

	// Note: if provided header map is nil, it will not throw error
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}