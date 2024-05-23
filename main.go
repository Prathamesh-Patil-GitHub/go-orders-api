package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	router := chi.NewRouter()     // Created a Chi Router
	router.Use(middleware.Logger) // Middleware for Logging the requests

	// Routes
	router.Get("/helloworld", helloWorldController)
	router.Get("/", homePageController)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server", err)
	}
}

// Route Controllers
func helloWorldController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func homePageController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page API"))
}
