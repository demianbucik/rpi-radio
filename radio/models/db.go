package models

import (
	"radio/app/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.Env.DB_FILE), nil)
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Playlist{}, &Track{}, &PlaylistTrack{}); err != nil {
		return nil, err
	}

	return db, nil
}
