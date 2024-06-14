// package main

// import (
// 	"log"
// 	"os"
// 	"strings"

// 	"context"
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
// 	if !strings.EqualFold(environment, "production") { // case insensitive comparison
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

// 	// setup context update  middleware
// 	// dbService := db_service.NewMongoService[ambulance_wl.Ambulance](db_service.MongoServiceConfig{})
// 	// setup context update middleware
// 	dbService := db_service.NewMongoService[medcare.AllergyRecord](db_service.MongoServiceConfig{
// 		ServerHost: "localhost",
// 		ServerPort: 27017,
// 		UserName:   "root",
// 		Password:   "neUhaDnes",
// 		DbName:     "medcare-db",
// 		Collections: map[string]string{
// 			"allergy":     "allergy_records",
// 			"labResults":  "lab_results",
// 			"medical":     "medical_records",
// 			"vaccination": "vaccination_records",
// 		},
// 		Timeout: 10 * time.Second,
// 	})
// 	defer dbService.Disconnect(context.Background())
// 	engine.Use(func(ctx *gin.Context) {
// 		ctx.Set("db_service", dbService)
// 		ctx.Next()
// 	})

// 	// request routings
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
		ServerHost: "localhost",
		ServerPort: 27017,
		UserName:   "root",      // replace with your actual username
		Password:   "neUhaDnes", // replace with your actual password
		DbName:     "medcare-db",
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
