package main

import (
	"log"
	"login-user/controller"
	"login-user/middleware"
	"login-user/prisma/db"
	"login-user/service"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		log.Fatalf("Erro ao conectar Prisma: %v", err)
	}

	defer client.Prisma.Disconnect()

	userService := service.UserService{Client: client}
	userController := controller.UserController{UserService: userService}

	r := mux.NewRouter()

	r.HandleFunc("/register", userController.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userController.LoginUser).Methods("POST")

	// Rotas
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/protected-endpoint", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Você acessou uma rota protegida!"))
	}).Methods("GET")

	log.Println("Servidor rodando na porta 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
