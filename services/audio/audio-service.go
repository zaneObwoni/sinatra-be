package service

import (
	"sinatra/entity"
)

type AudioService interface {
	Save(entity.Audio) entity.Audio
	FindAll() entity.Audio
}

type audioService struct {
	audio []entity.Audio
}

func New() AudioService {
	return &audioService{}
}

func (service *audioService) Save(entity.Audio) entity.Audio {
	return &audioService{}
}

func (service *audioService) FindAll() []entity.Audio {
	return nil
}