package repository

import (
	"time"

	"github.com/fernandoporazzi/bike-shop/app/entity"
	"github.com/google/uuid"
)

type BikesRepository interface {
	GetBikes() (entity.Stock, error)
	AddBike(bike *entity.Bike) (*entity.Bike, error)
	UpdateBike(bike *entity.Bike) (*entity.Bike, error)
	FindById(id string) entity.Bike
	AddImages(bike entity.Bike) error
}

type bikesRepository struct {
	stockEntity entity.Stock
}

func NewBikesRepository() BikesRepository {
	return &bikesRepository{}
}

func (r *bikesRepository) GetBikes() (entity.Stock, error) {
	return r.stockEntity, nil
}

func (r *bikesRepository) AddBike(bike *entity.Bike) (*entity.Bike, error) {
	bike.ID = uuid.NewString()
	bike.CreatedAt = time.Now()

	r.stockEntity.Bikes = append(r.stockEntity.Bikes, *bike)

	return bike, nil
}

func (r *bikesRepository) FindById(id string) entity.Bike {
	bike, _ := r.find(id)

	return bike
}

func (r *bikesRepository) UpdateBike(bike *entity.Bike) (*entity.Bike, error) {
	bikeToUpdate, index := r.find(bike.ID)

	bikeToUpdate.UpdatedAt = time.Now()
	bikeToUpdate.Stock = bike.Stock
	bikeToUpdate.Name = bike.Name
	bikeToUpdate.Color = bike.Color

	r.stockEntity.Bikes[index] = bikeToUpdate

	return &bikeToUpdate, nil
}

func (r *bikesRepository) AddImages(bike entity.Bike) error {
	bikeToUpdate, index := r.find(bike.ID)

	bikeToUpdate.Images = bike.Images
	r.stockEntity.Bikes[index] = bikeToUpdate

	return nil
}

func (r *bikesRepository) find(id string) (entity.Bike, int) {
	var bikeToUpdate entity.Bike
	var index int
	for i, b := range r.stockEntity.Bikes {
		if b.ID == id {
			bikeToUpdate = b
			index = i
			break
		}
	}

	return bikeToUpdate, index
}
