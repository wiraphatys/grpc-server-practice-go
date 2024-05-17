package servers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"user.services/config"
	"user.services/handlers"
	"user.services/repositories"
	"user.services/services"
)

type fiberServer struct {
	app *fiber.App
	cfg *config.Config
	db  *gorm.DB
}

func NewFiberServer(cfg *config.Config, db *gorm.DB) Server {
	return &fiberServer{
		app: fiber.New(),
		cfg: cfg,
		db:  db,
	}
}

func (s *fiberServer) Start() {
	url := fmt.Sprintf("%v:%d", s.cfg.Server.Host, s.cfg.Server.Port)

	// healthcheck
	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("User server is running . . .")
	})

	// connect all layer
	repository := repositories.NewUserRepository(s.db)
	service := services.NewUserService(repository)
	hander := handlers.NewUserHandler(service)

	// router
	router := s.app.Group("/users")
	router.Post("/", hander.CreateUser)
	router.Get("/:email", hander.GetUserByEmail)

	// log
	log.Printf("User server is starting on %v", url)
	if err := s.app.Listen(url); err != nil {
		log.Fatalf("Error while starting user server: %v", err)
	}
}
