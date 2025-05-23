package server

import (
	"fmt"
	"os"

	"github.com/Osas997/go-portfolio/internal/domains/auth"
	authController "github.com/Osas997/go-portfolio/internal/domains/auth/controller"
	authRepository "github.com/Osas997/go-portfolio/internal/domains/auth/repository"
	authService "github.com/Osas997/go-portfolio/internal/domains/auth/service"

	"github.com/Osas997/go-portfolio/internal/domains/projects"
	projectsController "github.com/Osas997/go-portfolio/internal/domains/projects/controller"
	projectsRepository "github.com/Osas997/go-portfolio/internal/domains/projects/repository"
	projectsService "github.com/Osas997/go-portfolio/internal/domains/projects/service"

	"github.com/Osas997/go-portfolio/internal/pkg/database"
	"github.com/Osas997/go-portfolio/internal/pkg/uploadfile"
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
	uploadfile.RegisterCustomValidators(validate)
	router.Static("/uploads", "./uploads")

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

	// auth domain
	userRepo := authRepository.NewUserRepository(s.db)
	authServ := authService.NewAuthService(userRepo)
	authCtrl := authController.NewAuthController(authServ, s.validate)
	auth.RegisterRoutes(routes, authCtrl)

	// projects domain
	projectsRepo := projectsRepository.NewProjectRepository(s.db)
	projectsImageRepo := projectsRepository.NewProjectImagesRepository(s.db)
	projectsServ := projectsService.NewProjectService(projectsRepo, projectsImageRepo)
	projectsCtrl := projectsController.NewProjectController(projectsServ, s.validate)
	projects.RegisterRoutes(routes, projectsCtrl)
}

func (s *Server) Start() error {
	s.runMigrations()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	return s.router.Run(fmt.Sprintf(":%s", port))
}

func (s *Server) runMigrations() {
	// user
	// s.db.Migrator().DropTable(entity.User{})
	// s.db.AutoMigrate(&entity.User{})
	// s.db.AutoMigrate(&entity.Projects{})
	// s.db.AutoMigrate(&entity.ProjectImages{})
	// password, _ := hash.HashPassword("password")

	// user := &entity.User{Username: "admin", Password: password}
	// s.db.Create(user)

	// tokens, err := token.GenerateToken(user)
	// if err != nil {
	// 	panic(err)
	// }

	// log.Println("access token: ", tokens.AccessToken)
	// log.Println("refresh token: ", tokens.RefreshToken)

	// Add other model migrations here
}
