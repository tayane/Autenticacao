package main

import (
	"autenticacao/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/register", api.RegisterHandler)
	http.HandleFunc("/login", api.LoginHandler)

	log.Println("Iniciando Servidor: 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Erro ao iniciar a porta 8081 %v", err)
	}
}
