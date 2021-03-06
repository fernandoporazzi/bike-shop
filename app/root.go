package app

import (
	"net/http"
	"time"

	"github.com/fernandoporazzi/bike-shop/app/controller"
	"github.com/fernandoporazzi/bike-shop/app/repository"
	"github.com/fernandoporazzi/bike-shop/app/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Start() {
	bikesRepository := repository.NewBikesRepository()
	uploadService := service.NewUploadService()
	bikesService := service.NewBikesService(bikesRepository)
	bikesController := controller.NewBikesController(bikesService, uploadService)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The bike shop API"))
	})

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Route("/bikes", func(r chi.Router) {
		r.Get("/", bikesController.GetBikes)
		r.Get("/{id}", bikesController.GetBikeById)
		r.Put("/{id}", bikesController.UpdateBike)
		r.Post("/", bikesController.AddBike)
		r.Post("/upload/{id}", bikesController.UploadImages)
	})

	http.ListenAndServe(":8080", r)
}
