package movie

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID        int            `gorm:"primaryKey, AUTO_INCREMENT" json:"id"`
	Title     string         `json:"Title"`
	Year      string         `json:"Year"`
	Rated     string         `json:"Rated"`
	Genre     []Category     `gorm:"many2many:movie_category;" json:"Genre"`
	Director  string         `json:"Director"`
	Writer    string         `json:"Writer"`
	Actors    string         `json:"Actors"`
	Plot      string         `json:"Plot"`
	Poster    string         `json:"Poster"`
	ImdbId    string         `json:"ImdbId" gorm:"unique"`
	Type      string         `json:"Type"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
