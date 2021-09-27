package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/fernandoporazzi/bike-shop/app/entity"
	"github.com/fernandoporazzi/bike-shop/app/errors"
	"github.com/fernandoporazzi/bike-shop/app/service"
)

type BikesController interface {
	GetBikes(response http.ResponseWriter, request *http.Request)
	AddBike(response http.ResponseWriter, request *http.Request)
	UpdateBike(response http.ResponseWriter, request *http.Request)
	UploadImages(response http.ResponseWriter, request *http.Request)
}

type bikesController struct {
	bikeService service.BikeService
}

func NewBikesController(bikeService service.BikeService) BikesController {
	return &bikesController{
		bikeService: bikeService,
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
	err := request.ParseMultipartForm(200000) // grab the multipart form
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error on ParseMultipartForm"})
		return
	}

	formdata := request.MultipartForm

	files := formdata.File["files"]

	for i, _ := range files {
		file, err := files[i].Open()

		defer file.Close()

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: err.Error()})
			return
		}

		// rename file
		newName := fmt.Sprintf("./static/%d%s", time.Now().UnixNano(), filepath.Ext(files[i].Filename))
		destination, err := os.Create(newName)

		defer destination.Close()

		if err != nil {
			fmt.Println(err)
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error creating destination"})
			return
		}

		if _, err := io.Copy(destination, file); err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error copying"})
			return
		}
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode("Success")
}
