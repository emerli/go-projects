package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/emerli/go-projects/connectionsfactory"
	"github.com/emerli/go-projects/handlers"
	"github.com/emerli/go-projects/middlewares"
	gorillaHandlers "github.com/gorilla/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func main() {
	db, err := connectionsfactory.NewDB()
	if err != nil {
		log.Printf("exceptionsError while initializing context: %s\n", err.Error())
	} else {
		log.Printf("database connection Ok! ")
	}
	srv := CreateService(db)

	log.Fatal(srv.ListenAndServe())
}

func CreateService(db *gorm.DB) *http.Server {
	r := mux.NewRouter()
	r.Use(middlewares.SetContentTypeMiddleware)
	r.Use(middlewares.LogAuditMiddleware)

	// appointment

	r.HandleFunc("/city", handlers.CreateCity(db)).Methods("POST")
	r.HandleFunc("/city/{id}", handlers.GetCity(db)).Methods("GET")
	r.HandleFunc("/city", handlers.GetAllCity(db)).Methods("GET")
	r.HandleFunc("/city/{id}", handlers.UpdateCity(db)).Methods("PATCH")
	r.HandleFunc("/city/{id}", handlers.DeleteCity(db)).Methods("DELETE")

	port := 8085

	headersOk := gorillaHandlers.AllowedHeaders([]string{
		"X-Requested-With",
		"Access-Control-Allow-Origin",
		"Authorization",
		"Access-Control-Request-Headers",
		"Access-Control-Request-Method",
		"Origin",
		"Referer",
		"User-Agent",
		"Content-Type",
		"user-id",
	})
	originsOk := gorillaHandlers.AllowedOrigins([]string{"*"})
	methodsOk := gorillaHandlers.AllowedMethods([]string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodHead,
	})

	srv := &http.Server{
		Handler:      gorillaHandlers.CORS(headersOk, originsOk, methodsOk)(r),
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		WriteTimeout: 180 * time.Second,
		ReadTimeout:  180 * time.Second,
		IdleTimeout:  200 * time.Second,
	}
	log.Printf("Server started on port %d\n", port)
	return srv
}
