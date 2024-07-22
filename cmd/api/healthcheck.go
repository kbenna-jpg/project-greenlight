package main

import (
	// "encoding/json"
	"net/http"
)

// Writes a plain-text response with information about the
// application status, operating environment and version

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//Envelope map holds all information we want to send in response, encoded as JSON object
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version": version,
		},
		
	}
		
	
	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
	}
}
// METHOD WITHOUT HANDLER:
// Map passed to json.Marshal() which returns byte slice including encoded JSON
// 	js, err := json.Marshal(data)
// 	if err != nil {
// 		app.logger.Print(err)
// 		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
// 		return
// 	}
	
// 	js = append(js, '\n')
// 	w.Header().Set("Content-Type", "application/json") // client knows its recieving JSON
// 	// w.WriteHeader(status)
// 	w.Write(js)


// }
	// fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env) // dependencies can be included as a field in application struct
	// fmt.Fprintf(w, "version: %s\n", version)

	// ANOTHER METHOD TO ACHIEVE THE SAME GOAL
	// Can write JSON response from GO handlers in same way as any other text response:
	// using w.Write()
// 	js := `{"status": "available", "environment": %q, "version": %q}`
// 	js = fmt.Sprintf(js, app.config.env, version)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(js))
// 