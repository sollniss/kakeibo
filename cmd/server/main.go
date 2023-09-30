package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sollniss/graceful"
	"github.com/sollniss/kakeibo/repository/xsyncmap"
	"github.com/sollniss/kakeibo/transport/http/html/plain"
	"github.com/sollniss/kakeibo/transport/http/router/httprouter"
	"github.com/sollniss/kakeibo/usecases/transfer"
)

func main() {
	transferRepository := xsyncmap.NewTransferRepository()
	transferRepository.AddDummyData()

	transferUsecase := transfer.NewUsecase(nil, transferRepository, transferRepository)
	frontend := plain.NewFrontend()

	router := httprouter.NewRouter(transferUsecase, frontend)

	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	ctx := graceful.NotifyShutdown()
	err := graceful.ListenAndServe(ctx, server, 60*time.Second)
	if err != nil {
		log.Printf("error during shutdown: %s", err)
		return
	}
	log.Print("gracefully shut down")
}
