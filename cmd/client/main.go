package main

import (
	"fmt"
	"log"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/service/user"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=agency sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: false})
	if err != nil {
		log.Fatalln(err)
	}

	uwork.Conn = db
	u := uwork.New().WithRole(models.AdminRole)
	s := user.NewService(u)
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
