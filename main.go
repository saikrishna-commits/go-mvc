package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/justinas/alice"
	"github.com/justinas/nosurf"
	db "github.com/saikrishna-commits/go-mvc/dbCon"

	models "github.com/saikrishna-commits/go-mvc/models"
	services "github.com/saikrishna-commits/go-mvc/services"
	utils "github.com/saikrishna-commits/go-mvc/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	_id         string     `bson:"id,omitempty"`
	Title       string     `bson:"title,omitempty"json:"title"`
	Cast        []string   `bson:"cast,omitempty"`
	Directors   []string   `bson:"directors"`
	Genres      []string   `bson:"genres,omitempty"`
	Year        int64      `bson:"year,omitempty"`
	Released    time.Time  `bson:"released,omitempty"`
	LastUpdated time.Time  `bson:"lastUpdated,omitempty"`
	ImdbRating  ImdbRating `bson:"imdb"`
}

type ImdbRating struct {
	Rating float64 `bson:"rating,omitempty"`
	Votes  int64   `bson:"votes"`
	// id     int
}

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

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func init() {
	godotenv.Load(".env") // load env variables from specific file , alterantive we can use viper package
}

const defaultPort = "8080"

func main() {

	handler := http.NewServeMux()

	db.ConnectDatabaseMongo() //connect to mongo
	db.CreatePgConnection()   //connect to postgres

	handler.HandleFunc("/addPost", addPostHandler)
	handler.HandleFunc("/covid", covidSummaryHandler)
	handler.HandleFunc("/hello", hello)
	handler.Handle("/protectedRoute", services.IsAuthorized(hello))
	handler.HandleFunc("/createAndTestJWT", createAndTest)
	handler.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		notAllowedMethods := []string{"POST", "PATCH", "DELETE", "PUT"}
		methodType := r.Method
		isNotAllowed := utils.Contains(notAllowedMethods, methodType)
		if isNotAllowed == true {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		id := strings.SplitN(r.URL.Path, "/", 3)[2]
		if len(id) == 0 {
			http.Error(w, "Missing Input", http.StatusInternalServerError)
			return
		}

		data, err := queryToByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	wrappedHandler := alice.New(LoggingMiddleware, recoverHandler, nosurf.NewPure).Then(handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      wrappedHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// Start our HTTP server
	server.ListenAndServe()

}

func createAndTest(w http.ResponseWriter, r *http.Request) {
	validToken, err := services.GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/protectedRoute", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	fmt.Fprintf(w, string(body))
}

func hello(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	// Getting the instance of the collection from MongoDB Database
	collection := db.MongoClient.Database("sample_mflix").Collection("movies")

	//find options

	findOptions := options.Find()
	findOptions.SetLimit(2)
	findOptions.SetSort(bson.D{{"title", -1}})

	// Writing query to fetch the Data from the `movies` collection
	databaseCursor, err := collection.Find(context.Background(), bson.D{
		{"year", bson.D{
			{"$gt", 2010},
		}}}, findOptions)

	var movies []Movie

	defer databaseCursor.Close(context.Background())

	if err = databaseCursor.All(context.Background(), &movies); err != nil {
		log.Fatalln(err)
	}

	json.NewEncoder(w).Encode(movies)

}

func addPostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.AddPost
	json.NewDecoder(r.Body).Decode(&post)
	jsonValue, _ := json.Marshal(post)
	data := bytes.NewReader(jsonValue)
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", data)
	req.Header.Set("Content-Type", "application/json")
	// create a Client
	client := &http.Client{}
	// Do sends an HTTP request and
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err.Error())
		w.WriteHeader(400)
		//w.Write(err)
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var respData models.AddPost
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		fmt.Println(err)
	}

	resJSON, err := json.Marshal(&respData)
	w.Write(resJSON)

}

func queryToByID(id string) (models.PostData, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + id)
	if err != nil {
		return models.PostData{}, err
	}
	defer resp.Body.Close()
	var postDataR models.PostData
	if err := json.NewDecoder(resp.Body).Decode(&postDataR); err != nil {
		return models.PostData{}, err
	}

	return postDataR, nil
}

func covidSummaryHandler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	data, err := getCovidSummary()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func getCovidSummary() (models.GlobalCovid, error) {
	resp, err := http.Get("https://api.covid19api.com/summary")
	if err != nil {
		return models.GlobalCovid{}, err
	}
	defer resp.Body.Close()
	var gCovidData models.GlobalCovid
	if err := json.NewDecoder(resp.Body).Decode(&gCovidData); err != nil {
		return models.GlobalCovid{}, err
	}
	return gCovidData, nil
}

// func authRequired(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		next.ServeHTTP(w, r)

// 	}
// }

// func withHeader(key, value string) Adapter {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			w.Header.Add(key, value)
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

// catch panics, log them and keep the application running
func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// LoggingMiddleware logs the incoming HTTP request & its duration.
func LoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(
					"err", err,
					"trace", debug.Stack(),
				)
			}
		}()

		start := time.Now()

		next.ServeHTTP(w, r)

		log.Println(
			"method", r.Method,
			"path", r.URL.EscapedPath(),
			"duration", time.Since(start),
		)
	}

	return http.HandlerFunc(fn)
}
