package main

import (
	"log"
	"os"

	"github.com/ceperic/backend/internal/config"
	"github.com/ceperic/backend/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	// Inicializar configuración
	cfg := config.New()

	// Conectar a base de datos (opcional en desarrollo)
	var db *gorm.DB
	if cfg.DBHost != "" && cfg.DBHost != "localhost" {
		var err error
		db, err = config.NewDatabase(cfg)
		if err != nil {
			log.Printf("⚠️  Database connection failed (continuing without DB): %v", err)
		} else {
			// Ejecutar migraciones automáticas
			if err := config.AutoMigrate(db); err != nil {
				log.Printf("⚠️  Migrations failed: %v", err)
			}
		}
	} else {
		log.Println("📝 Running without database (development mode)")
	}

	// Inicializar Firebase (opcional en desarrollo)
	var firebaseApp interface{}
	// firebaseApp, err = config.NewFirebase(cfg)
	// if err != nil {
	// 	log.Printf("⚠️  Firebase initialization skipped: %v", err)
	// }

	// Crear aplicación Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
		AppName:      "CEPERIC API v1.0",
		ServerHeader: "CEPERIC",
	})

	// Middlewares globales
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path} (${latency})\n",
		TimeFormat: "15:04:05",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AllowedOrigins,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Setup de rutas
	router.Setup(app, db, firebaseApp)

	// Puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor
	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("🌍 Environment: %s", cfg.Environment)
	log.Printf("📝 API Docs: http://localhost:%s/api/v1/health", port)
	
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}

// Error handler personalizado
func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	log.Printf("❌ Error: %s", err.Error())

	return ctx.Status(code).JSON(fiber.Map{
		"success": false,
		"error":   message,
		"code":    code,
	})
}
