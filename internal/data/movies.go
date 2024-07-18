package data


import (
	"time"
)

// Between ``, customize key names that appear in JSON object by annotating fields with struct tags 
type Movie struct {
	ID int64 `json:"id"`// unique integer id for the movie
	CreatedAt time.Time `json:"-"`// Timestamp for when movie is added ti db
	Title string `json:"title"`// Movie title
	Year int32 `json:"year,omitempty"`
	Runtime int32 `json:"runtime,omitempty"`// movie runtime in mins
	Genres []string `json:"genres,omitempty"`
	Version int32 `json:"version"`// incremented if movie is updated
}





