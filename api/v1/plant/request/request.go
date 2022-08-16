package request

import (
	"plant-api/business/native"
	"plant-api/business/plant"
)

type Request struct {
	Name          string          `json:"name"`
	BotanicalName string          `json:"botanical_name"`
	Type          string          `json:"type"`
	Difficulty    string          `json:"difficulty"`
	Description   string          `json:"description"`
	Natives       []NativeRequest `json:"natives"`
	WateringTime  string          `json:"watering_time"`
	HowToGrow     string          `json:"how_to_grow"`
	Soil          string          `json:"soil"`
}

type NativeRequest struct {
	ID uint `json:"id"`
}

func MapToNative(request NativeRequest) *native.Native {
	return &native.Native{
		ID: request.ID,
	}
}

func MapToNatives(request []NativeRequest) (natives []*native.Native) {
	for _, native := range request {
		natives = append(natives, MapToNative(native))
	}
	return
}

func (req *Request) MapToModel() plant.Plant {
	return plant.Plant{
		Name:          req.Name,
		BotanicalName: req.BotanicalName,
		Type:          req.Type,
		Difficulty:    req.Difficulty,
		Description:   req.Description,
		Natives:       MapToNatives(req.Natives),
		WateringTime:  req.WateringTime,
		HowToGrow:     req.HowToGrow,
		Soil:          req.Soil,
	}
}
