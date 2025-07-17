package server

import (
	"database/sql"
	"frdy-api/config"
	"frdy-api/internal/auth"
	"frdy-api/internal/items"
	"frdy-api/internal/purchases"
	"frdy-api/internal/sales"
	"frdy-api/internal/stock"
	"frdy-api/internal/users"
	"frdy-api/middleware"

	_ "frdy-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           FRDY API
// @version         1.0
// @description     API per gestionar compres, ventes i stocks bàsica
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8125
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

type Server struct {
	router *gin.Engine
	cfg    *config.Config
	db     *sql.DB
}

func NewServer(cfg *config.Config, db *sql.DB) *Server {
	return &Server{
		router: gin.Default(),
		cfg:    cfg,
		db:     db,
	}
}

func (s *Server) Setup() error {
	// CORS middleware
	s.router.Use(middleware.SetupCORS())
	

	// JWT middleware
	authMiddleware, err := middleware.SetupJWT(s.cfg)
	if err != nil {
		return err
	}

	// Action log middleware
	//actionLogMiddleware := middleware.NewActionLogMiddleware(s.db)
	
	// Inicialitzar repositoris
	userRepo := users.NewUserRepository(s.db)
	itemRepo := items.NewItemRepository(s.db)
	salesRepo := sales.NewSalesRepository(s.db)
	stockRepo := stock.NewStockRepository(s.db)
	purchaseRepo := purchases.NewPurchaseRepository(s.db)



	// Inicialitzar serveis
	userService := users.NewUserService(userRepo)
	authService := auth.NewAuthService(userRepo, authMiddleware)
	itemService := items.NewItemService(itemRepo)
	salesService := sales.NewSalesService(salesRepo)
	stockService := stock.NewStockService(stockRepo)
	purchaseService := purchases.NewPurchaseService(purchaseRepo)



	// Inicialitzar handlers
	userHandler := users.NewUserHandler(userService)
	authHandler := auth.NewAuthHandler(authService, authMiddleware)
	itemHandler := items.NewItemHandler(itemService)
	salesHandler := sales.NewSalesHandler(salesService)
	stocksHandler := stock.NewStockHandler(stockService)
	purchaseHandler := purchases.NewPurchasesHandler(purchaseService)

	
	// Configurar les rutes públiques (sense autenticació)
	public := s.router.Group("/auth")
	//public.Use(actionLogMiddleware.LogAction())
	users.RegisterPublicRoutes(public, userHandler)
	auth.RegisterRoutes(public, authHandler, authMiddleware)
	public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	// Configurar les rutes protegides (amb autenticació JWT)
	protected := s.router.Group("/api")
	protected.Use(authMiddleware.MiddlewareFunc())

	

	// Registrar les rutes protegides
	users.RegisterRoutes(protected, userHandler)
	items.RegisterRoutes(protected, itemHandler)
	sales.RegisterRoutes(protected, salesHandler)
	stock.RegisterRoutes(protected, stocksHandler)
	purchases.RegisterRoutes(protected, purchaseHandler)

	
	return nil
}

func (s *Server) Run() error {
	//return s.router.RunTLS(":" + s.cfg.ApiPort, "./certs/cert.pem", "./certs/key.pem")
	return s.router.Run(":" + s.cfg.ApiPort)
}