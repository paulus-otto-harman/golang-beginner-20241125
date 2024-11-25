package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	gola "github.com/paulus-otto-harman/golang-module/web"
	"go.uber.org/zap"
	"project/config"
	"project/database"
	"project/handler"
	"project/middleware"
	"project/repository"
	"project/service"
	"project/util"
)

func Init(config config.AppConfig) (*chi.Mux, *zap.Logger, *sql.DB) {
	logger := util.InitLog(config)
	db := database.Open(config.Database)

	services := service.InitServices(repository.InitRepositories(db, logger, config))
	handle := handler.InitHandlers(util.InitValidator(services, logger))
	middlewares := middleware.InitMiddlewares(services, logger)

	r := chi.NewRouter()
	r.With(middlewares.AccessLog.Middleware).Route("/api", func(r chi.Router) {
		r.Use(gola.JsonResponse())

		r.Post("/register", handle.Customer.Registration)
		r.Post("/login", handle.Auth.Login)

		r.Get("/banners", handle.Banner.Get)

		r.Route("/categories", func(r chi.Router) {
			r.Get("/", handle.Category.Get)
			r.Get("/{category_id}/products", handle.Product.All)
		})

		r.Get("/products", handle.Product.All)
		r.Get("/products/{id}", handle.Product.Get)

		r.Get("/best-sellers", handle.Bestseller.All)
		r.Get("/weeklies", handle.Weekly.Get)
		r.Get("/recommendations", handle.Recommendation.Get)

		r.With(middlewares.Auth.Middleware).Group(func(r chi.Router) {
			r.Route("/wishlist", func(r chi.Router) {
				r.Post("/", handle.Wishlist.Store)
				r.Delete("/", handle.Wishlist.Destroy)
			})

			r.Route("/cart", func(r chi.Router) {
				r.Get("/", handle.Cart.Get)
				r.Post("/", handle.Cart.Create)
				r.Put("/", handle.Cart.Update)
				r.Delete("/{product_id}", handle.Cart.Delete)

				r.Post("/checkout", handle.Order.Create)
			})

			r.Route("/addresses", func(r chi.Router) {
				r.Get("/", handle.Address.All)
				r.Post("/", handle.Address.Create)
				r.Put("/{id}", handle.Address.Update)
				r.Put("/{id}/default", handle.Address.Default)
				r.Delete("/{id}", handle.Address.Delete)
			})

			r.Get("/profile", handle.Customer.Profile)
		})

	})

	return r, logger, db
}
