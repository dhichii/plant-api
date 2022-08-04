package plant

import (
	"plant-api/business/native"
	"time"

	"gorm.io/gorm"
)

type Plant struct {
	ID            uint             `json:"id"`
	Name          string           `json:"name"`
	BotanicalName string           `json:"botanical_name"`
	Type          string           `json:"type"`
	Difficulty    string           `json:"difficulty"`
	Description   string           `json:"description"`
	Natives       []*native.Native `gorm:"many2many:plant_natives" json:"natives"`
	WateringTime  string           `json:"watering_time"`
	HowToGrow     string           `json:"how_to_grow"`
	Soil          string           `json:"soil"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	DeletedAt     gorm.DeletedAt   `gorm:"index" json:"deleted_at"`
}
