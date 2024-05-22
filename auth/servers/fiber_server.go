package servers

import (
	"fmt"
	"log"

	"auth.services/config"
	"auth.services/handlers"
	"auth.services/repositories"
	"auth.services/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type server struct {
	app *fiber.App
	cfg *config.Config
}

func NewFiberServer(cfg *config.Config) Server {
	return &server{
		app: fiber.New(),
		cfg: cfg,
	}
}

func (s *server) Start() {
	url := fmt.Sprintf("%v:%d", s.cfg.Server.Host, s.cfg.Server.Port)

	// middleware
	s.app.Use(logger.New(logger.Config{
		Format:     "${time} ${status} - ${method} ${path}\n",
		TimeFormat: "2006/01/02 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))

	// healthcheck
	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("User server is running . . .")
	})

	// connect all layer
	repository := repositories.NewAuthRepository(s.cfg)
	service := services.NewAuthService(repository, s.cfg)
	hander := handlers.NewAuthHttpHandler(service)

	// router
	router := s.app.Group("/auth")
	router.Post("/login", hander.Login)

	// log
	log.Printf("Auth server is starting on %v", url)
	if err := s.app.Listen(url); err != nil {
		log.Fatalf("Error while starting user server: %v", err)
	}
}
