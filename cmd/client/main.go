package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DmiAS/bd_course/internal/app/config"
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/service/user"
	"github.com/DmiAS/bd_course/internal/app/uwork"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln("err = ", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	unit, err := uwork.New(ctx, cfg.DB)
	if err != nil {
		log.Fatalln("err = ", err)
	}
	newUnit := unit.WithRole(models.AdminRole)
	s := user.NewService(newUnit)
	if data, err := s.Create(&models.User{
		FirstName: "dima",
		LastName:  "antsibor",
		VkLink:    "_",
		TgLink:    "_",
		Role:      models.AdminRole,
	}); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(*data)
	}
}
