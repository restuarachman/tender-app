package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username           string `gorm:"uniqueIndex;size:50"`
	Password           string
	Email              string `gorm:"uniqueIndex;size:50"`
	Name               string
	ProfileImageUrl    string
	BackgroundImageUrl string
	Bio                string
	Profession         string
	ValidEmail         bool `gorm:"default:false"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
