package quote

import (
	//"encoding/json"
	"net/http"

	"encoding/json"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type Response struct {
	Message string `json:"message"`
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := appengine.NewContext(r)

	var quotes []*Quote
	inspirationalQuotes := GetInspirationalQuotes()
	movieQuotes := GetMovieQuotes()
	programmingQuotes := GetProgrammingQuotes()
	quotes = append(quotes, inspirationalQuotes...)
	quotes = append(quotes, movieQuotes...)
	quotes = append(quotes, programmingQuotes...)

	keys := make([]*datastore.Key, len(quotes))
	for i := range quotes {
		keys[i] = datastore.NewIncompleteKey(ctx, "Quote", nil)
	}

	if _, err := datastore.PutMulti(ctx, keys, quotes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	successRes := Response{"Success"}
	response, err := json.Marshal(successRes)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
