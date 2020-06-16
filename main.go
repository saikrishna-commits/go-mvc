package main

import (
	"encoding/json"
	"net/http"
	"strings"

	utils "github.com/saikrishna-commits/go-mvc/lib/utils"
)

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {

		notAllowedMethods := [...]string{"POST", "PATCH", "DELETE", "PUT"}

		methodType := r.Method

		isAllowed := utils.Contains(notAllowedMethods, methodType)

		if isAllowed == false {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}

		id := strings.SplitN(r.URL.Path, "/", 3)[2]
		if len(id) == 0 {
			http.Error(w, "Missing Input", http.StatusInternalServerError)
			return
		}

		data, err := queryToById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func queryToById(id string) (postData, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + id)
	if err != nil {
		return postData{}, err
	}

	defer resp.Body.Close()

	var d postData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return postData{}, err
	}

	return d, nil
}

type postData struct {
	ID        int    `json:"userId"`
	TodoTitle string `json:"title"`
}
