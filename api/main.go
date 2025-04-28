//Package where this file belongs
package main

//Import necessary packages
import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"coptic_dictionary/api/models"
	"coptic_dictionary/api/routers"
)

//Load the .env with dotEnv and make viper retrieve the variables
func loadConfig() {
	//If loading the file fails, err will contain the error message
	//Otherwise, err will be nil (0 value for many types)
	err := godotenv.Load(".env")

	//If err is not nil, then loading the file failed
	if err != nil {
		//Log that it failed and display the error
		log.Fatal("Error loading .env file:", err)
	}

	//Viper reads the env variables
	viper.AutomaticEnv()
}

//global variable; pointer, a reference to the actual database connection
//(rather than = gorm.DB which a copy)
var db *gorm.DB

//Database connection
func connectDatabase() {
	//Load the environment variables
	loadConfig()

	//Build the PostgreSQL DSN using env variables
	//Format the one-line DSN string using env variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_SSLMODE"),
	)

	//Connect to the database
	//declare variable named err as the type error (built-in)
	//gorm.Open opens database connection using dsn previously formatted
	//gorm.Config loads GORM config settings (&gorm --> passing a pointer, not a value)
	//db stores the database connection
	//if err is not nil (meaning error has occured, return the error)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Successfully connected to the database!")

	//Auto-migrate the database (create/update tables as model.go is updated)
	//& used as pointer to point to CopticDictionary type in model.go
	err = db.AutoMigrate(&models.CopticDictionary{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully!")
}

//Entry point of the app
func main() {
	//Connect to database
	connectDatabase()

	//Initialize Gin router
	r := gin.Default()

    //CORS configuration
    corsConfig := cors.Config{
        AllowOrigins:     []string{"https://copticscribe.com"},
        AllowMethods:     []string{"GET", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        AllowCredentials: true,
    }

    //Apply CORS config
    r.Use(cors.New(corsConfig))

	//Set up routes using modular router package
	routers.SetupRoutes(r, db)

	//Start server on port 8080
	r.Run(":8080")
}
