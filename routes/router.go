package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaiminbhaduri/codeprac/controllers"
)

func SetupRoutes() http.Handler {
	router := mux.NewRouter()

	// Static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// HTML views
	router.HandleFunc("/", controllers.Dashboard).Methods("GET")

	// API endpoints
	router.HandleFunc("/api/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/execute", controllers.ExecuteCode).Methods("POST")

	return router
}
