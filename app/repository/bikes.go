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

func (r *bikesRepository) UpdateBike(bike *entity.Bike) (*entity.Bike, error) {
	// find bike to be updated and the index
	var bikeToUpdate entity.Bike
	var index int
	for i, b := range r.stockEntity.Bikes {
		if b.ID == bike.ID {
			bikeToUpdate = b
			index = i
			break
		}
	}

	bikeToUpdate.UpdatedAt = time.Now()
	bikeToUpdate.Stock = bike.Stock
	bikeToUpdate.Name = bike.Name
	bikeToUpdate.Color = bike.Color

	// update bike by index
	r.stockEntity.Bikes[index] = bikeToUpdate

	return &bikeToUpdate, nil
}
