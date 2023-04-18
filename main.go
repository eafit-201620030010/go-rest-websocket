package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"jjchavarrg.com/go/rest-ws/handlers"
	"jjchavarrg.com/go/rest-ws/middleware"
	"jjchavarrg.com/go/rest-ws/server"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatalf("Error creating server %v\n", err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {

	api := r.PathPrefix("/api/v1").Subrouter()

	api.Use(middleware.CheckAuthMiddleware(s))

	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)

	// auth
	r.HandleFunc("/api/v1/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)

	// post
	api.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/posts/{postId}", handlers.DeletePostHandler(s)).Methods(http.MethodDelete)
	api.HandleFunc("/posts/{postId}", handlers.UpdatePostHandler(s)).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/posts/{postId}", handlers.GetPostByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/posts", handlers.ListPostHandler(s)).Methods(http.MethodGet)

	// websocket
	r.HandleFunc("/ws", s.Hub().HandleWebSocket)
}
