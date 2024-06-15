// package main

// import (
// 	"context"
// 	"log"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/DanielCok17/medcare-webapi/api"
// 	"github.com/DanielCok17/medcare-webapi/internal/db_service"
// 	"github.com/DanielCok17/medcare-webapi/internal/medcare"
// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	log.Printf("Server started")
// 	port := os.Getenv("AMBULANCE_API_PORT")
// 	if port == "" {
// 		port = "8080"
// 	}
// 	environment := os.Getenv("AMBULANCE_API_ENVIRONMENT")
// 	if !strings.EqualFold(environment, "production") {
// 		gin.SetMode(gin.DebugMode)
// 	}
// 	engine := gin.New()
// 	engine.Use(gin.Recovery())
// 	corsMiddleware := cors.New(cors.Config{
// 		AllowOrigins:     []string{"*"},
// 		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
// 		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
// 		ExposeHeaders:    []string{""},
// 		AllowCredentials: false,
// 		MaxAge:           12 * time.Hour,
// 	})
// 	engine.Use(corsMiddleware)

// 	// Setup the database service with the correct collection names
// 	config := db_service.MongoServiceConfig{
// 		ServerHost: "localhost",
// 		ServerPort: 27017,
// 		UserName:   "root",      // replace with your actual username
// 		Password:   "neUhaDnes", // replace with your actual password
// 		DbName:     "medcare-db",
// 		Collections: map[string]string{
// 			"allergy_records":     "allergy_records",
// 			"lab_results":         "lab_results",
// 			"medical_records":     "medical_records",
// 			"vaccination_records": "vaccination_records",
// 		},
// 		Timeout: 10 * time.Second,
// 	}

// 	// Initialize DbService for each type
// 	allergyDbService := db_service.NewMongoService[medcare.AllergyRecord](config)
// 	defer allergyDbService.Disconnect(context.Background())

// 	labResultDbService := db_service.NewMongoService[medcare.LabResult](config)
// 	defer labResultDbService.Disconnect(context.Background())

// 	medicalRecordDbService := db_service.NewMongoService[medcare.MedicalRecord](config)
// 	defer medicalRecordDbService.Disconnect(context.Background())

// 	vaccinationRecordDbService := db_service.NewMongoService[medcare.VaccinationRecord](config)
// 	defer vaccinationRecordDbService.Disconnect(context.Background())

// 	// Set the db_service in the context
// 	engine.Use(func(ctx *gin.Context) {
// 		ctx.Set("allergy_db_service", allergyDbService)
// 		ctx.Set("lab_result_db_service", labResultDbService)
// 		ctx.Set("medical_record_db_service", medicalRecordDbService)
// 		ctx.Set("vaccination_record_db_service", vaccinationRecordDbService)
// 		ctx.Next()
// 	})

// 	// Add routes
// 	medcare.AddRoutes(engine)
// 	engine.GET("/openapi", api.HandleOpenApi)
// 	engine.Run(":" + port)
// }

package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/DanielCok17/medcare-webapi/api"
	"github.com/DanielCok17/medcare-webapi/internal/db_service"
	"github.com/DanielCok17/medcare-webapi/internal/medcare"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Server started")
	port := os.Getenv("AMBULANCE_API_PORT")
	if port == "" {
		port = "8080"
	}
	environment := os.Getenv("AMBULANCE_API_ENVIRONMENT")
	if !strings.EqualFold(environment, "production") {
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{""},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
	engine.Use(corsMiddleware)

	// Setup the database service with the correct collection names
	config := db_service.MongoServiceConfig{
		UserName: os.Getenv("AMBULANCE_API_MONGODB_USERNAME"), // replace with your actual username pls nech to ide
		Password: os.Getenv("AMBULANCE_API_MONGODB_PASSWORD"), // replace with your actual password
		DbName:   os.Getenv("AMBULANCE_API_MONGODB_DATABASE"),
		Collections: map[string]string{
			"allergy_records":     "allergy_records",
			"lab_results":         "lab_results",
			"medical_records":     "medical_records",
			"vaccination_records": "vaccination_records",
		},
		Timeout: 10 * time.Second,
	}

	// Initialize DbService for each type
	allergyDbService := db_service.NewMongoService[medcare.AllergyRecord](config)
	defer allergyDbService.Disconnect(context.Background())

	labResultDbService := db_service.NewMongoService[medcare.LabResult](config)
	defer labResultDbService.Disconnect(context.Background())

	medicalRecordDbService := db_service.NewMongoService[medcare.MedicalRecord](config)
	defer medicalRecordDbService.Disconnect(context.Background())

	vaccinationRecordDbService := db_service.NewMongoService[medcare.VaccinationRecord](config)
	defer vaccinationRecordDbService.Disconnect(context.Background())

	// Set the db_service in the context
	engine.Use(func(ctx *gin.Context) {
		ctx.Set("allergy_db_service", allergyDbService)
		ctx.Set("lab_result_db_service", labResultDbService)
		ctx.Set("medical_record_db_service", medicalRecordDbService)
		ctx.Set("vaccination_record_db_service", vaccinationRecordDbService)
		ctx.Next()
	})

	// Add routes
	medcare.AddRoutes(engine)
	engine.GET("/openapi", api.HandleOpenApi)
	engine.Run(":" + port)
}
