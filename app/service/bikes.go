package service

import (
	"github.com/fernandoporazzi/bike-shop/app/entity"
	"github.com/fernandoporazzi/bike-shop/app/repository"
)

type BikeService interface {
	GetBikes() (entity.Stock, error)
	AddBike(bike *entity.Bike) (*entity.Bike, error)
	UpdateBike(bike *entity.Bike) (*entity.Bike, error)
	AddImages(bikeId string, paths []string) error
}

type bikeService struct {
	bikesRepository repository.BikesRepository
}

func NewBikesService(bikesRepository repository.BikesRepository) BikeService {
	return &bikeService{bikesRepository}
}

func (s *bikeService) GetBikes() (entity.Stock, error) {
	stock, err := s.bikesRepository.GetBikes()

	if err != nil {
		return entity.Stock{}, err
	}

	return stock, nil
}

func (s *bikeService) AddBike(bike *entity.Bike) (*entity.Bike, error) {
	b, _ := s.bikesRepository.AddBike(bike)
	return b, nil
}

func (s *bikeService) UpdateBike(bike *entity.Bike) (*entity.Bike, error) {
	b, _ := s.bikesRepository.UpdateBike(bike)
	return b, nil
}

func (s *bikeService) AddImages(bikeId string, paths []string) error {
	bike := s.bikesRepository.FindById(bikeId)

	bike.Images = paths

	return s.bikesRepository.AddImages(bike)
}
