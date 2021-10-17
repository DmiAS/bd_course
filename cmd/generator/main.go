package main

import (
	"log"
	"strconv"

	"github.com/DmiAS/bd_course/internal/app/models"
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

	var camp models.Campaign
	db.Model(&models.Campaign{}).Last(&camp)
	threadID := uuid.MustParse("8d366c42-cf0a-4413-a897-e4061758a760")
	targetID := uuid.MustParse("876d02ed-f06f-404a-8533-440d2f9c9602")
	for i := 1; i < 11; i++ {
		db.Create(&models.Campaign{
			ID:           uuid.New(),
			ThreadID:     threadID,
			TargetologID: targetID,
			CabinetID:    camp.CabinetID + i,
			ClientID:     camp.ClientID + i,
			VkCampID:     camp.VkCampID + i,
			Name:         strconv.Itoa(i),
		})
	}
}
