package controller

import (
	"encoding/json"
	"net/http"

	"github.com/fernandoporazzi/bike-shop/app/entity"
	"github.com/fernandoporazzi/bike-shop/app/errors"
	"github.com/fernandoporazzi/bike-shop/app/service"
	"github.com/go-chi/chi/v5"
)

type BikesController interface {
	GetBikes(response http.ResponseWriter, request *http.Request)
	AddBike(response http.ResponseWriter, request *http.Request)
	UpdateBike(response http.ResponseWriter, request *http.Request)
	UploadImages(response http.ResponseWriter, request *http.Request)
}

type bikesController struct {
	bikeService   service.BikeService
	uploadService service.UploadService
}

func NewBikesController(bikeService service.BikeService, uploadService service.UploadService) BikesController {
	return &bikesController{
		bikeService:   bikeService,
		uploadService: uploadService,
	}
}

func (c *bikesController) GetBikes(response http.ResponseWriter, request *http.Request) {
	bikes, err := c.bikeService.GetBikes()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error returning bikes"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(bikes)

}

func (c *bikesController) AddBike(response http.ResponseWriter, request *http.Request) {
	var bike entity.Bike

	err := json.NewDecoder(request.Body).Decode(&bike)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	b, err := c.bikeService.AddBike(&bike)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error returning bikes"})
		return
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(b)

}

func (c *bikesController) UpdateBike(response http.ResponseWriter, request *http.Request) {
	var bike entity.Bike

	err := json.NewDecoder(request.Body).Decode(&bike)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	b, err := c.bikeService.UpdateBike(&bike)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error returning bikes"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(b)
}

func (c *bikesController) UploadImages(response http.ResponseWriter, request *http.Request) {
	bikeId := chi.URLParam(request, "id")

	err := request.ParseMultipartForm(200000)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	formdata := request.MultipartForm

	paths, err := c.uploadService.Upload(formdata)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	// add the paths to the repository
	err = c.bikeService.AddImages(bikeId, paths)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(paths)
}
