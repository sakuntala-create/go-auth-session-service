package main

import (
	"log"
	"net/http"

	"example.com/config"
	"example.com/internal/handlers"
	"example.com/internal/middleware"
	"example.com/internal/utils"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	utils.ConnectMongo(cfg.MongoURI, cfg.MongoDB)
	utils.InitJWT(cfg.JWTSecret)

	mux := http.NewServeMux()
	mux.HandleFunc("/auth/signup", handlers.Signup)
	mux.HandleFunc("/auth/login", handlers.Login)

	protected := middleware.JWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("protected route"))
	}))
	mux.Handle("/me", protected)

	log.Println("🚀 running on :" + cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, mux))
}
