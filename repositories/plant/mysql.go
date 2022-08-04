package plant

import (
	"plant-api/api/v1/plant/response"
	"plant-api/business/native"
	"plant-api/business/plant"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// Generate mysql DB repository
func NewMysqlRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Type used to update plant
type plantModel plant.Plant

// Create new plant and store into database
func (repo *repository) Create(plant *plant.Plant) (uint, error) {
	if err := repo.db.Create(plant).Error; err != nil {
		return 0, err
	}
	return plant.ID, nil
}

/*
Get all plants by given name
it will return all plants if name is null
*/
func (repo *repository) GetAll(name string) ([]response.Plant, error) {
	plants := []response.Plant{}
	if err := repo.db.Preload("Natives").
		Where("deleted_at IS NULL AND name LIKE '%" + name + "%'").
		Find(&plants).
		Error; err != nil {
		return nil, err
	}
	return plants, nil
}

// Get plant by given id
func (repo *repository) GetDetail(id int) (*response.PlantDetail, error) {
	plant := &response.PlantDetail{}
	if err := repo.db.Table("plants").
		Where("deleted_at IS NULL").First(&plant, id).Error; err != nil {
		return nil, err
	}
	return plant, nil
}

// GetAllNativesByPlantID Get all natives by given plant id
func (repo *repository) GetAllNativesByPlantID(id int) ([]*response.Native, error) {
	natives := []*response.Native{}
	if err := repo.db.Table("natives").
		Joins("JOIN plant_natives ON natives.id = plant_natives.native_id").
		Where("plant_id", id).
		Find(&natives).
		Error; err != nil {
		return nil, err
	}
	return natives, nil
}

// Update plant and store it into database
func (repo *repository) Update(id int, plant plant.Plant) error {
	if err := repo.db.Model(&plant).
		Where("id", id).
		Updates(
			plantModel{
				Name:          plant.Name,
				BotanicalName: plant.BotanicalName,
				Type:          plant.Type,
				Difficulty:    plant.Difficulty,
				Description:   plant.Description,
				WateringTime:  plant.WateringTime,
				HowToGrow:     plant.HowToGrow,
				Soil:          plant.Soil,
			},
		).Error; err != nil {
		return err
	}
	return nil
}

// UpdatePlantNatives will update plant natives and store it into database
func (repo *repository) UpdatePlantNatives(id int, plantNatives []*native.Native) error {
	plant := &plant.Plant{ID: uint(id)}
	if err := repo.db.Model(plant).Association("Natives").Replace(plantNatives); err != nil {
		return err
	}
	return nil
}

/*
GetNativeByID get native by native id
it will return nil if native is found
*/
func (repo *repository) GetNativeByID(id int) error {
	return repo.db.First(new(native.Native), id).Error
}

// Delete plant by given id
func (repo *repository) Delete(id int) error {
	plant := &plant.Plant{}
	if err := repo.db.Delete(&plant, id).Error; err != nil {
		return err
	}
	return nil
}
