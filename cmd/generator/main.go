package main

import (
	"log"
	"math"
	random "math/rand"
	"strconv"
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=agency sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: false})
	if err != nil {
		log.Fatalln(err)
	}
	generateCampaignStat(db, "aea9c20d-1b43-479b-9806-12ea2d697172")
	//generateCampaigns(db)
}

func generateCampaigns(db *gorm.DB) {
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

func generateCampaignStat(db *gorm.DB, campaignID string) {
	date := time.Now().Add(5 * -time.Hour * 24)
	const border = 10000
	id := uuid.MustParse(campaignID)
	for date.Day() <= time.Now().Day() {
		subs, unsubs := generateSubSlices()
		db.
			Create(&models.CampaignStat{
				CampID:      id,
				Date:        date,
				Spent:       float64(random.Intn(border)) * random.Float64(),
				Impressions: random.Intn(border),
				Conversion:  random.Intn(border),
				Subs:        subs,
				Unsubs:      unsubs,
				Sales:       random.Intn(border),
			})
		date = date.Add(time.Hour * 24)
	}
}

func generateSubSlices() (pq.Int64Array, pq.Int64Array) {
	const border = 10
	subCap := 1 + random.Int63n(border)
	subs := make(pq.Int64Array, subCap)
	unsubCap := 1 + random.Int63n(subCap)
	unsubs := make(pq.Int64Array, unsubCap)
	for i := range subs {
		subs[i] = 1 + random.Int63n(math.MaxInt)
	}

	for i := range unsubs {
		index := random.Int63n(subCap)
		unsubs[i] = subs[index]
	}

	return subs, unsubs
}
