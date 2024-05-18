package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/polyhistor2050/microservices/handlers"
)

func main() {
    l := log.New(os.Stdout, "product-api", log.LstdFlags)
    ph := handlers.NewProduct(l)

    sm := http.NewServeMux()
    sm.Handle("/", ph)

    svr := &http.Server{
        Addr: ":8080",
        Handler: sm,
        IdleTimeout: 120*time.Second,
        ReadTimeout: 1*time.Second,
        WriteTimeout: 1*time.Second,
    }

    go func ()  {
        l.Println("\nServer is running on port 8080")

        err := svr.ListenAndServe()
        if err != nil {
            l.Printf("Error starting server: %s\n", err)
            os.Exit(1)
        }
    }()

    // Successfully terminate the server
    sigChan := make(chan os.Signal)
    signal.Notify(sigChan, os.Interrupt)
    signal.Notify(sigChan, os.Kill)

    sig := <- sigChan
    l.Println("Recieved terminate, graceful shutdown", sig)

    tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
    svr.Shutdown(tc)

}


