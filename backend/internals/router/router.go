package router

import (
	"os"

	"github.com/Diaku49/FoodOrderSystem/backend/Redis"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/handler"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/middleware"
	"github.com/Diaku49/FoodOrderSystem/backend/mq"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, mqc *mq.MQClient) *chi.Mux {
	r := chi.NewRouter()
	keyStr := os.Getenv("JWT_SECRET")
	key := []byte(keyStr)

	//setup Redis
	rdbc := Redis.NewRedisClient()

	//setup CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true, // Important for cookies/auth
		MaxAge:           300,  // Cache results (seconds)
	}))

	// User routes
	userHandlers := handler.NewUH(db)
	r.Route("/user", func(r chi.Router) {
		r.Post("/login", userHandlers.Login)
		r.Post("/signup", userHandlers.Signup)
		r.Get("/profile", middleware.Auth(userHandlers.GetProfile, key))
	})

	// Video routes
	videoHandlers := handler.NewVH(db, mqc, rdbc)
	r.Route("/video", func(r chi.Router) {
		r.Get("", videoHandlers.GetAllVideos)
		r.Post("/upload", middleware.Auth(videoHandlers.UploadHandler, key))
		r.Get("/upload/{videoId}", middleware.Auth(videoHandlers.GetVideoInfoHandler, key))
	})

	return r
}
