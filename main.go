package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/jul-cesar/learning-golang/internal/app"
	"github.com/jul-cesar/learning-golang/internal/routes"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 8080, "Port to run the application on")
	flag.Parse()

	application, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	defer  application.Db.Close()
	
	r := routes.SetupRoutes(application)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	application.Logger.Printf("Server running on port %d", port)

	err = server.ListenAndServe()
	if err != nil {
		application.Logger.Fatal(err)
	}
}
