package plant

import (
	"errors"
	"plant-api/api/v1/plant/response"
	"plant-api/business"
	"plant-api/business/native"
	plantnative "plant-api/business/plant_native"
)

type service struct {
	repository     Repository
	nativeService  native.Service
	pNativeService plantnative.Service
}

// Construct plant service object
func NewService(
	repo Repository,
	nativeService native.Service,
	pNativeRepo plantnative.Repository,
) Service {
	return &service{
		repo,
		nativeService,
		pNativeRepo,
	}
}

// Create new plant and store into database
func (s *service) Create(plant *Plant) (uint, error) {
	/*
		Find native by name
		create new native if native not found
	*/
	for _, nativeRequest := range plant.Natives {
		if err := s.repository.GetNativeByID(int(nativeRequest.ID)); err != nil {
			if err.Error() == "record not found" {
				return nativeRequest.ID, business.ErrNotFound
			}
			return 0, err
		}
	}
	id, err := s.repository.Create(plant)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Get all plants by given name
func (s *service) GetAll(name string) ([]response.Plant, error) {
	return s.repository.GetAll(name)
}

/*
Get detail plant by given id
will return ErrNotFound if plant is not exist
*/
func (s *service) GetDetail(id int) (*response.PlantDetail, error) {
	plant, err := s.repository.GetDetail(id)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, business.ErrNotFound
		}
		return nil, err
	}

	natives, err := s.repository.GetAllNativesByPlantID(id)
	if err != nil {
		return nil, err
	}
	plant.Natives = natives
	return plant, nil
}

/*
Update existing plant and plant natives in database
will return ErrNotFound when plant is not exist
*/
func (s *service) Update(id int, plant Plant) error {
	_, err := s.repository.GetDetail(id)
	if err != nil {
		if err.Error() == "record not found" {
			return business.ErrNotFound
		}
		return err
	}
	if plant.Natives != nil {
		for _, native := range plant.Natives {
			err := s.repository.GetNativeByID(int(native.ID))
			if err != nil {
				if err.Error() == "record not found" {
					return errors.New("native not found")
				}
				return err
			}
		}
		if err := s.repository.UpdatePlantNatives(id, plant.Natives); err != nil {
			return err
		}
		plant.Natives = nil
	}
	if err := s.repository.Update(id, plant); err != nil {
		return err
	}
	return nil
}

/*
Delete plant native and plant
will return ErrNotFound when plant is not exist
*/
func (s *service) Delete(id int) error {
	_, err := s.repository.GetDetail(id)
	if err != nil {
		if err.Error() == "record not found" {
			return business.ErrNotFound
		}
		return err
	}
	if err := s.pNativeService.Delete(id); err != nil {
		return err
	}
	if err := s.repository.Delete(id); err != nil {
		return err
	}
	return nil
}
