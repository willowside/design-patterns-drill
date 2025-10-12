package main

import (
	"flag"
	"fmt"
	"go-designpattern/models"
	"html/template"
	"log"
	"net/http"
	"time"
)

// A simple web server that listens for requests
const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	// DB          *sql.DB // use Repository in models
	Models models.Models
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}
	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.Parse()

	// get db connection
	db, err := initMySQLDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}
	// app.DB = db
	app.Models = *models.New(db)

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello, world!")
	// })

	fmt.Println("Starting web application on port", port)

	// err := http.ListenAndServe(port, nil)
	// if err != nil {
	// 	log.Panic(err)
	// }

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
