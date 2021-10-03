package main

import (
	"context"
	"github.com/DmiAS/bd_course/internal/app/config"
	"github.com/DmiAS/bd_course/internal/app/delivery/http"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/handler"
	"github.com/DmiAS/bd_course/internal/app/service"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"os"
	"os/signal"
)

func main() {
	unit := uwork.New()
	wf := service.NewWorkerFactory(unit)
	af := service.NewAuthFactory(unit)
	cf := service.NewClientFactory(unit)
	router := handler.NewHandler(wf, af, cf)
	cfg := config.Config{HTTP: config.HTTP{Port: "80"}}
	server := http.NewServer(router, cfg)
	go server.Start()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	ctx, cancel := context.WithCancel(context.Background())
	server.Stop(ctx)
	cancel()
}
