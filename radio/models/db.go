package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), nil)
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Playlist{}, &Track{}, &PlaylistTrack{}); err != nil {
		return nil, err
	}

	// seed(db)

	return db, nil

}

func seed(db *gorm.DB) {
	log.Println("Inside")
	db.Create(&Playlist{
		Name: "Default",
	})

	var p Playlist
	db.Find(&p, "name = ?", "Default")

	db.Create(&Track{
		Name: "Matter/persons from porlock ~ MANA",
		Url:  "https://www.youtube.com/watch?pt=i2NlVQi9XUE",
	})
	db.Create(&Track{
		Name: "Matter ~ Sedimenti",
		Url:  "https://www.youtube.com/watch?pt=lzp4EXoOtSo",
	})

	var mana, sedimenti Track
	db.Find(&mana, "name = ?", "Matter/persons from porlock ~ MANA")
	db.Find(&sedimenti, "name = ?", "Matter ~ Sedimenti")

	var p0 uint = 0
	var p1 uint = 1
	db.Create(&PlaylistTrack{
		Playlist: &p,
		Track:    &sedimenti,
		Position: &p0,
	})
	db.Create(&PlaylistTrack{
		Playlist: &p,
		Track:    &mana,
		Position: &p1,
	})

	// var pts []PlaylistTrack
	// // var s []Track
	// err := db.Preload("Track").Order("position").Find(&pts, &PlaylistTrack{PlaylistID: p.ID}).Error
	// // err := db.Preload("PlaylistTrack").Find(&s, &PlaylistTrack{PlaylistID: p.ID}).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, pt := range pts {
	// 	j, _ := json.MarshalIndent(pt, "", "  ")
	// 	log.Println(string(j))
	// }

	// db.Model(&pts).Find(&s)
	// log.Println(s)
}
