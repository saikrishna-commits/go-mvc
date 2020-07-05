package model


type ImdbRating struct {
	Rating float64 `json:"rating,omitempty"`
	Votes  int64   `json:"votes"`
}

type Movie struct {
	_id         string     `json:"id,omitempty" bson:"_id"`
	Title       string     `json:"title,omitempty"json:"title"`
	Cast        []string   `json:"cast,omitempty"`
	Directors   []string   `json:"directors"`
	Genres      []string   `json:"genres,omitempty"`
	Year        int64      `json:"year,omitempty"`
    // Released    time.Time  `json:"released,omitempty"`
    //  LastUpdated time.Time  `json:"lastUpdated,omitempty"`
	ImdbRating  ImdbRating `bson:"imdb" json:"imdb"`  // nested elements
}

