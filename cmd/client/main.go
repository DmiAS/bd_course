package main

import "C"
import (
	"log"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/pkg/gen"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=agency sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: false})
	if err != nil {
		log.Fatalln(err)
	}
	id := uuid.New()
	ids := models.IDs{ID: id, Role: "admin"}
	db.Model(models.IDs{}).Create(ids)
	salt := []byte("salt")
	password := []byte("1234")
	enc, err := gen.PasswordWithSalt(password, salt)
	if err != nil {
		log.Fatalln(enc)
	}
	if err := db.Create(&models.Auth{
		Login:    "advwolf",
		Password: enc,
		Salt:     "salt",
		UserID:   id,
	}).Error; err != nil {
		log.Fatalln(err)
	}
}
