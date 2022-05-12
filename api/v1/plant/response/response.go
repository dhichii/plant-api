package response

type Plant struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name"`
	BotanicalName string         `json:"botanical_name"`
	Type          string         `json:"type"`
	Difficulty    string         `json:"difficulty"`
	Description   string         `json:"description"`
}

type PlantDetail struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name"`
	BotanicalName string         `json:"botanical_name"`
	Type          string         `json:"type"`
	Difficulty    string         `json:"difficulty"`
	Description   string         `json:"description"`
	Natives       []*Native      `gorm:"many2many:plant_natives" json:"natives"`
	WateringTime  string         `json:"watering_time"`
	HowToGrow     string         `json:"how_to_grow"`
	Soil          string         `json:"soil"`
}

type Native struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
}
