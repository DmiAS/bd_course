package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/DmiAS/bd_course/internal/app/config"
	"github.com/DmiAS/bd_course/internal/app/delivery/http"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/handler"
	"github.com/DmiAS/bd_course/internal/app/service"
	"github.com/DmiAS/bd_course/internal/app/uwork"
)

func main() {
	unit := uwork.New()
	wf := service.NewWorkerFactory(unit)
	af := service.NewAuthFactory(unit)
	cf := service.NewClientFactory(unit)
	pf := service.NewProjectFactory(unit)
	tf := service.NewThreadFactory(unit)
	router := handler.NewHandler(wf, af, cf, pf, tf)
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
