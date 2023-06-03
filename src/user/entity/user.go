package entity

import (
	"time"

	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type User struct {
	gorm.Model
	Nickname        string
	Email           string `gorm:"uniqueIndex;size:50"`
	Password        string
	ProfileImageUrl string
	Gender          Gender `gorm:"column:gender;type:varchar(20)"`
	Popularity      int64  `gorm:"default:0"`
	IsVerified      bool   `gorm:"default:false"`
	Details         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
