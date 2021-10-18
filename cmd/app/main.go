package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/DmiAS/bd_course/internal/app/config"
	my_http "github.com/DmiAS/bd_course/internal/app/delivery/http"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/handler"
	"github.com/DmiAS/bd_course/internal/app/service"
	"github.com/DmiAS/bd_course/internal/app/uwork"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	router, err := createNewRouter(ctx, cfg)
	if err != nil {
		log.Fatalln("can't create router -", err.Error())
	}

	server := my_http.NewServer(router, *cfg)
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("server starts on port -", cfg.HTTP.Port)
	waitForInterrupt()
	defer cancel()
	if err := server.Stop(ctx); err != nil {
		log.Fatalf("server Shutdown Failed:%+v", err)
	}
	log.Print("server Exited Properly")
}

func createNewRouter(ctx context.Context, cfg *config.Config) (*handler.Handler, error) {
	unit, err := uwork.New(ctx, cfg.DB)
	if err != nil {
		return nil, err
	}
	wf := service.NewWorkerFactory(unit)
	af := service.NewAuthFactory(unit)
	uf := service.NewUserFactory(unit)
	pf := service.NewProjectFactory(unit)
	tf := service.NewThreadFactory(unit)
	cmpf := service.NewCampaignsFactory(unit)
	sf := service.NewStatsFactory(unit)
	router := handler.NewHandler(wf, af, uf, pf, tf, cmpf, sf)
	return router, nil
}

func waitForInterrupt() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
