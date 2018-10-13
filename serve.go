package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const DEV_PORT = "5000"
const DEV_ADDR = ""
const DEV_STATIC = "./static_pages"

func main() {
	var env string
	var port string
	var addr string
	var static string

	flag.StringVar(&env, "env", "dev", "Server Env")
	flag.StringVar(&port, "port", "", "Port to listen on")
	flag.StringVar(&addr, "addr", "", "Host IP Address")
	flag.StringVar(&static, "static", "", "Static files location")

	flag.Parse()

	if env == "dev" {
		if port == "" {
			port = DEV_PORT
		}
		if addr == "" {
			addr = DEV_ADDR
		}
		if static == "" {
			static = DEV_STATIC
		}
	}

	if env == "prod" && (port == "" || addr == "" || static == "") {
		log.Fatal("Failed to start server. Missing required config value(s)")
	}

	writeTimeout := 15 * time.Second
	readTimeout := 15 * time.Second

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/notes/{id:[0-9]+}", GetNoteHandler).Methods("GET")
	api.HandleFunc("/notes", GetNotesHandler).Methods("GET")
	api.HandleFunc("/games", GetAllGames).Methods("GET")

	// everything below to the DEV_END_HERE are for local dev only. Prod has caddy doing proxy servy stuff
	r.PathPrefix("/note-app").Handler(http.FileServer(http.Dir(static)))
	r.PathPrefix("/games").Handler(http.FileServer(http.Dir(static)))

	r.HandleFunc("/", IndexHandler)

	// opening up CORS if we are in dev env, for JS apps
	frontendHost := ""
	if env == "dev" {
		// vuejs apps uses this host for dev
		frontendHost = "http://localhost:8080"
	}
	// DEV_END_HERE

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{frontendHost},
	}).Handler(r)

	srv := &http.Server{
		Addr:         addr + ":" + port,
		Handler:      handlers.LoggingHandler(os.Stdout, handler),
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}
	log.Println("Listening at " + addr + ":" + port)
	log.Fatal(srv.ListenAndServe())
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is under construction")
}

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{{
		"id":    "1",
		"title": "note 1",
		"body":  "this is a note",
	}}

	j, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write(j)
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	data := map[string]interface{}{
		"id":    vars["id"],
		"title": "note 1",
		"body":  "this is a note",
	}

	j, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write(j)
}

type Game struct {
	location string
	name     string
}

func NewGame(location string) *Game {
	g := Game{
		location: location,
		name:     location,
	}

	return &g
}

func GetAllGames(w http.ResponseWriter, r *http.Request) {
	dirs, err := ioutil.ReadDir("./static_pages/games")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := []map[string]interface{}{}
	for _, d := range dirs {
		if d.IsDir() {
			game := NewGame(d.Name())
			data = append(data, map[string]interface{}{
				"location": game.location,
				"name":     game.name,
			})
		}

	}
	j, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(j)
}
