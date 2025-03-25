package main

import (
	"beginnerGo/internal/app"
	"beginnerGo/internal/routes"
	"fmt"
	"net/http"
	"time"
)

const PORT = 8080

func main() {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("We are running our app!")

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", PORT),
		Handler:      routes.SetupRoutes(app),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	serverErr := server.ListenAndServe()

	if serverErr != nil {
		app.Logger.Fatal(serverErr)
	}
}
