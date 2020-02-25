package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bagus123/go-date-example2"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	// our srvdate service
	srv := srvdate.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := srvdate.Endpoints{
		GetEndpoint:      srvdate.MakeGetEndpoint(srv),
		StatusEndpoint:   srvdate.MakeStatusEndpoint(srv),
		ValidateEndpoint: srvdate.MakeValidateEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("srvdate is listening on port:", *httpAddr)
		handler := srvdate.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}