//Package where this file belongs
package routers

//Import necessary packages
import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"coptic_dictionary/api/handlers"
)

//Initialize all API routes
//r is the Gin router instance
//db is the connected database instance
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	//Define the GET routes
	//router.[TYPE_OF_REQUEST]("/[ROUTE]", function)
	r.GET("/words", handlers.GetCopticWords(db))
	r.GET("/word", handlers.GetOneCopticWord(db))
}