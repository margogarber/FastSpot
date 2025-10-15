package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/fastspot/backend/configs"
	"github.com/fastspot/backend/internal/handlers"
	"github.com/fastspot/backend/internal/middleware"
	"github.com/fastspot/backend/internal/repository"
	"github.com/fastspot/backend/internal/services/ai"
	"github.com/fastspot/backend/internal/services/payments"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration
	config := configs.LoadConfig()

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(ctx)

	// Ping MongoDB
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}
	log.Println("✅ Connected to MongoDB successfully")

	// Initialize database
	db := client.Database("fastspot")

	// Initialize repositories
	repos := &repository.Repositories{
		Users:         repository.NewUserRepository(db),
		Categories:    repository.NewCategoryRepository(db),
		Products:      repository.NewProductRepository(db),
		Promotions:    repository.NewPromotionRepository(db),
		Carts:         repository.NewCartRepository(db),
		Orders:        repository.NewOrderRepository(db),
		MoodQuestions: repository.NewMoodQuestionRepository(db),
		AISessions:    repository.NewAISessionRepository(db),
	}

	// Initialize services
	geminiService := ai.NewGeminiService(config.GeminiAPIKey)
	paymentService := payments.NewStubProvider()

	// Initialize Gin router
	if config.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Session-ID"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "FastSpot API is running"})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Auth routes
		authHandler := handlers.NewAuthHandler(repos, config)
		auth := v1.Group("/auth")
		{
			auth.POST("/guest", authHandler.CreateGuestSession)
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
			auth.GET("/me", middleware.AuthMiddleware(config.JWTSecret), authHandler.GetCurrentUser)
			auth.POST("/logout", middleware.AuthMiddleware(config.JWTSecret), authHandler.Logout)
		}

		// Categories routes (public read, admin write)
		categoryHandler := handlers.NewCategoryHandler(repos)
		categories := v1.Group("/categories")
		{
			categories.GET("", categoryHandler.GetAll)
			categories.GET("/:slug", categoryHandler.GetBySlug)
		}
		adminCategories := v1.Group("/admin/categories", middleware.AuthMiddleware(config.JWTSecret), middleware.AdminMiddleware())
		{
			adminCategories.POST("", categoryHandler.Create)
			adminCategories.GET("/:id", categoryHandler.GetByID)
			adminCategories.PUT("/:id", categoryHandler.Update)
			adminCategories.DELETE("/:id", categoryHandler.Delete)
		}

		// Products routes
		productHandler := handlers.NewProductHandler(repos)
		products := v1.Group("/products")
		{
			products.GET("", productHandler.GetAll)
			products.GET("/:slug", productHandler.GetBySlug)
		}
		adminProducts := v1.Group("/admin/products", middleware.AuthMiddleware(config.JWTSecret), middleware.AdminMiddleware())
		{
			adminProducts.GET("", productHandler.GetAll)
			adminProducts.POST("", productHandler.Create)
			adminProducts.GET("/:id", productHandler.GetByID)
			adminProducts.PUT("/:id", productHandler.Update)
			adminProducts.DELETE("/:id", productHandler.Delete)
		}

		// Promotions routes
		promotionHandler := handlers.NewPromotionHandler(repos)
		promotions := v1.Group("/promotions")
		{
			promotions.GET("", promotionHandler.GetAll)
			promotions.GET("/:id", promotionHandler.GetByID)
		}
		adminPromotions := v1.Group("/admin/promotions", middleware.AuthMiddleware(config.JWTSecret), middleware.AdminMiddleware())
		{
			adminPromotions.POST("", promotionHandler.Create)
			adminPromotions.PUT("/:id", promotionHandler.Update)
			adminPromotions.DELETE("/:id", promotionHandler.Delete)
		}

		// Cart routes
		cartHandler := handlers.NewCartHandler(repos)
		cart := v1.Group("/cart", middleware.OptionalAuthMiddleware(config.JWTSecret))
		{
			cart.GET("", cartHandler.Get)
			cart.POST("/items", cartHandler.AddItem)
			cart.PUT("/items/:productId", cartHandler.UpdateItem)
			cart.DELETE("/items/:productId", cartHandler.RemoveItem)
			cart.DELETE("", cartHandler.Clear)
		}

		// Orders routes
		orderHandler := handlers.NewOrderHandler(repos, paymentService)
		orders := v1.Group("/orders", middleware.OptionalAuthMiddleware(config.JWTSecret))
		{
			orders.POST("", orderHandler.Create)
			orders.GET("", orderHandler.GetAll)
			orders.GET("/:id", orderHandler.GetByID)
			orders.POST("/:id/cancel", orderHandler.Cancel)
		}
		adminOrders := v1.Group("/admin/orders", middleware.AuthMiddleware(config.JWTSecret), middleware.AdminMiddleware())
		{
			adminOrders.GET("", orderHandler.GetAllAdmin)
			adminOrders.PUT("/:id/status", orderHandler.UpdateStatus)
		}

		// Mood Quiz routes
		moodHandler := handlers.NewMoodHandler(repos, geminiService)
		mood := v1.Group("/mood")
		{
			mood.GET("/questions", moodHandler.GetQuestions)
			mood.POST("/recommend", moodHandler.GetRecommendations) // No auth required for guest users
		}
		// Mood Quiz Management (Questions only - AI decides recommendations)
		adminMood := v1.Group("/admin/mood", middleware.AuthMiddleware(config.JWTSecret), middleware.AdminMiddleware())
		{
			adminMood.GET("/questions", moodHandler.GetAllQuestions)
			adminMood.GET("/questions/:id", moodHandler.GetQuestionByID)
			adminMood.POST("/questions", moodHandler.CreateQuestion)
			adminMood.PUT("/questions/:id", moodHandler.UpdateQuestion)
			adminMood.DELETE("/questions/:id", moodHandler.DeleteQuestion)
		}

		// Admin dashboard
		adminHandler := handlers.NewAdminHandler(repos)
		admin := v1.Group("/admin", middleware.AuthMiddleware(config.JWTSecret), middleware.AdminMiddleware())
		{
			admin.GET("/analytics", adminHandler.GetAnalytics)
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
