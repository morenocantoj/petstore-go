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

	"github.com/morenocantoj/petstore-go/internal/app/types/responses"
	"github.com/morenocantoj/petstore-go/internal/pkg/utils/errors"
)

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

func router(router *http.ServeMux) {
	router.HandleFunc("/", http.HandlerFunc(welcome))
}

func launchServer() {

	// suscripción SIGINT
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	mux := http.NewServeMux()
	router(mux)

	server := &http.Server{Addr: ":9080", Handler: mux}

	go func() {
		err := server.ListenAndServe()
		errors.Check(err)
		log.Printf("listen: %s\n", err)
	}()

	<-stopChan // espera señal SIGINT
	log.Println("Apagando servidor ...")

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
	launchServer()
}
