package server

import (
	"fmt"
	"log"
	"os"

	"github.com/Osas997/go-portfolio/internal/domains/auth"
	"github.com/Osas997/go-portfolio/internal/domains/auth/controller"
	"github.com/Osas997/go-portfolio/internal/domains/auth/entity"
	"github.com/Osas997/go-portfolio/internal/domains/auth/repository"
	"github.com/Osas997/go-portfolio/internal/domains/auth/service"
	"github.com/Osas997/go-portfolio/internal/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Server struct {
	db       *gorm.DB
	router   *gin.Engine
	validate *validator.Validate
}

func NewServer() *Server {
	db := database.InitDB()
	router := gin.Default()
	validate := validator.New()

	server := &Server{
		db:       db,
		router:   router,
		validate: validate,
	}

	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	routes := s.router.Group("/api")

	userRepo := repository.NewUserRepository(s.db)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService, s.validate)
	auth.RegisterRoutes(routes, authController)

	// Add other domain routes here
}

func (s *Server) Start() error {
	// Run database migrations
	s.runMigrations()

	// Start the server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	return s.router.Run(fmt.Sprintf(":%s", port))
}

func (s *Server) runMigrations() {
	s.db.Migrator().DropTable(entity.User{})
	s.db.AutoMigrate(&entity.User{})
	user := &entity.User{Username: "admin", Password: "admin"}
	s.db.Create(user)
	log.Print(user)
	// Add other model migrations here
}
