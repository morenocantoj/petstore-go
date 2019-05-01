package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/morenocantoj/petstore-go/internal/app/controllers"
	"github.com/morenocantoj/petstore-go/internal/app/types/responses"
	"github.com/morenocantoj/petstore-go/internal/pkg/utils/errors"
)

// Controllers
var PetsController = controllers.PetsController{}

func welcome(writter http.ResponseWriter, request *http.Request) {
	fmt.Println("POST /")
	response := responses.WelcomeJSON{
		Message:  "Welcome to my PetStore API server",
		LoginURL: fmt.Sprintf("%s/auth", request.URL.Host),
	}
	responseJSON, err := json.Marshal(&response)
	errors.Check(err)
	writter.Write(responseJSON)
}

func defineRoutes(router *mux.Router) {
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/pets", PetsController.Index).Methods("GET")
}

func launchServer(address string, port string) {
	listeningAddress := address + port
	// suscripción SIGINT
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	router := mux.NewRouter()
	defineRoutes(router)

	server := &http.Server{
		Addr:         listeningAddress,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second}

	go func() {
		fmt.Printf("Server listening at %s...\n", listeningAddress)
		err := server.ListenAndServe()
		errors.Check(err)
	}()

	<-stopChan // espera señal SIGINT
	log.Println("Shutting down server ...")

	// apagar servidor de forma segura
	ctx, fnc := context.WithTimeout(context.Background(), 5*time.Second)
	fnc()
	server.Shutdown(ctx)

	log.Println("Servidor detenido correctamente")
}

// Server: Runs a MUX instance to handle client requests
func Server() {
	fmt.Println("### PetStore API Server ###")
	fmt.Println("Server launching...")
	launchServer("127.0.0.1", ":9080")
}
