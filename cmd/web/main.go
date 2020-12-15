package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
	"tkircsi/restful-template/pkg/models"
	"tkircsi/restful-template/pkg/models/mock"
	"tkircsi/restful-template/pkg/models/mongodb"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	products models.ProductHandler
}

var app *application
var (
	addr *string
	dsn  *string
)

func init() {

	// set up command-line parameters
	addr = flag.String("addr", ":5000", "HTTP Server address")
	dsn = flag.String("dsn", "mock", "Datasource name/connection string")
	flag.Parse()

	// set up loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// set up application context
	app = &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// set up datasource according to dsn
	switch *dsn {
	case "mock":
		app.products = mock.NewProductModel()
	case "mongo":
		model, err := mongodb.NewProductModel("mongodb://root:example@localhost:27017")
		if err != nil {
			app.errorLog.Fatalf("error connecting database: %v\n", err)
		}
		app.infoLog.Printf("successfully connected to %s database\n", *dsn)
		app.products = model
	default:
		app.products = mock.NewProductModel()
	}
}

func main() {

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     app.errorLog,
		Handler:      app.Router(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.infoLog.Printf("Starting server on %s\n", *addr)
	err := srv.ListenAndServe()
	app.errorLog.Fatal(err)

}
