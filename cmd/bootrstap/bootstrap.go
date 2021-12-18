package bootrstap

import (
	"database/sql"
	"majo_test/handler"
	"majo_test/midleware"
	"majo_test/repository"
	"majo_test/services"
	"majo_test/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	token  utils.Maker
	config utils.Config
	db     *sql.DB
	router *gin.Engine
}

func NewServer(config utils.Config, store *sql.DB, token utils.Maker) (*Server, error) {

	server := &Server{
		config: config,
		db:     store,
		token:  token,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	userRepository := repository.NewUserRepository(server.db)
	userService := services.NewUserServicec(userRepository, server.token, server.config)
	userHandler := handler.NewUserHandler(userService)

	reportRepository := repository.NewReportRepository(server.db)
	reportService := services.NewReportService(reportRepository)
	reportHandler := handler.NewReportHandler(userService, reportService)

	router.POST("/auth/login", userHandler.SignIn)
	rg := router.Group("/api/v1")
	rg.Use(midleware.AuthMiddleware(server.token))
	rg.POST("/report", reportHandler.MonthlyReport)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
