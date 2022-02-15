package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tiagompalte/go-rest-api/controllers"
	"github.com/tiagompalte/go-rest-api/middleware"
)

func HandleResponse() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/api/personalidades", controllers.FindAll).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.FindById).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.Insert).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.Delete).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.Update).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
