package entity

import (
	"time"
)

type InteractionType string

const (
	Like InteractionType = "like"
	Pass InteractionType = "pass"
)

type Interaction struct {
	UserGivenId     uint            `gorm:"primaryKey"`
	UserReceivedId  uint            `gorm:"primaryKey"`
	InteractionType InteractionType `gorm:"column:interactin_type;type:varchar(10)"`
	CreatedAt       time.Time       `gorm:"autoCreateTime:true"`
}
