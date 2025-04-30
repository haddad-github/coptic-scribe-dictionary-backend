//Package where this file belongs
package test

//Import necessary packages
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"coptic_dictionary/api/models"
	"coptic_dictionary/api/routers"
)

//Helper to initialize in-memory test DB and seed test data
func setupTestRouter() (*gin.Engine, *gorm.DB) {
	//Disable Gin debug logs during tests
	gin.SetMode(gin.TestMode)

	//Initialize Gin
	r := gin.Default()

	//Create in-memory SQLite DB
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}

	//Auto-migrate schema
	db.AutoMigrate(&models.CopticDictionary{})

	//Seed test data
	db.Create(&models.CopticDictionary{
		CopticWord:         "ⲡⲣⲟⲥⲕⲩⲛⲉⲓ",
		EnglishTranslation: "worship",
	})

	//Set up routes using test DB
	routers.SetupRoutes(r, db)

	return r, db
}

//Test GET /words
func TestGetCopticWords(t *testing.T) {
	r, _ := setupTestRouter()

	//Simulate GET request
	req, _ := http.NewRequest("GET", "/words", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	//Expect HTTP 200
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", w.Code)
	}

	//Expect JSON array with at least 1 item
	var response []map[string]interface{}
	json.NewDecoder(bytes.NewReader(w.Body.Bytes())).Decode(&response)
	if len(response) == 0 {
		t.Errorf("Expected at least 1 word in response")
	}
}

//Test GET /word?coptic=ⲡⲣⲟⲥⲕⲩⲛⲉⲓ
func TestGetOneCopticWord_Success(t *testing.T) {
	r, _ := setupTestRouter()

	req, _ := http.NewRequest("GET", "/word?coptic=ⲡⲣⲟⲥⲕⲩⲛⲉⲓ", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 but got %d", w.Code)
	}
}

//Test GET /word?coptic=ⲡⲣⲟⲥⲕⲩⲛⲏ (misspelled)
func TestGetOneCopticWord_Suggestion(t *testing.T) {
	r, _ := setupTestRouter()

	req, _ := http.NewRequest("GET", "/word?coptic=ⲡⲣⲟⲥⲕⲩⲛⲏ", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404 but got %d", w.Code)
	}

	//Check suggestion in body
	var response map[string]interface{}
	json.NewDecoder(bytes.NewReader(w.Body.Bytes())).Decode(&response)
	if _, ok := response["suggestion"]; !ok {
		t.Errorf("Expected suggestion in response")
	}
}