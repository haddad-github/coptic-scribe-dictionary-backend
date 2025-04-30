//Package where this file belongs
package handlers

//Import necessary packages
import (
	"log"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"gorm.io/gorm"
	"coptic_dictionary/api/models"
)

//API endpoint GET all Coptic words
//c manages request, response and metadata (equivalent to request.json, request.args)
func GetCopticWords(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Initialize an empty list of models.CopticDictionary
		var words []models.CopticDictionary

		//Query the database results
		//db.Find(&words) equivalent to query all
		result := db.Find(&words)
		log.Println("Query Result:", words) //Debugging

		//Check for errors after query
		//gin.H creates a JSON object, so the response returns a json
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve words"})
			return
		}

		//words already returns JSON, no need for gin.H
		c.JSON(http.StatusOK, words)
	}
}
func GetOneCopticWord(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		copticWord := c.Query("coptic")
		log.Printf("Received /word?coptic=%s\n", copticWord)

		// DEBUG SHORT-CIRCUIT RESPONSE
		c.JSON(http.StatusOK, gin.H{"status": "received", "input": copticWord})
	}
}

// //API endpoint GET one Coptic word
// func GetOneCopticWord(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		//Query path is /word?coptic=[word_here]
// 		//Fetch the parameter's value
// 		copticWord := c.Query("coptic")
//
// 		//If no parameter (no word), returns error
// 		if copticWord == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'coptic' query parameter"})
// 			return
// 		}
//
// 		//Initialize the variable "word" which follows the CopticDictionary type
// 		var word models.CopticDictionary
//
// 		//Query the database for "coptic_word = [copticWord from parameter]"
// 		//get the first result .First and store it (&) in the variable "word" defined previously
// 		result := db.Where("coptic_word = ?", copticWord).First(&word)
//
// 		//If word found, return it
// 		if result.Error == nil {
// 			c.JSON(http.StatusOK, word)
// 			return
// 		}
//
// 		//WORD SUGGESTION ALGORITHM
// 		//If word not found, find closest word using Levenshtein distance
// 		//List of words of the CopticDictionary type
// 		var words []models.CopticDictionary
//
// 		//Query all Coptic words (only the coptic_word column, for efficiency) and store it in words
// 		db.Select("coptic_word").Find(&words)
//
// 		//Initialize closestWord
// 		//Set threshold for minimum distance (arbitrarily huge so that it never reaches it)
// 		closestWord := ""
// 		minDistance := 1000
//
// 		//Loop through all words but ignore first value (index), only use w (word)
// 		//rune --> converts string to characters instead of its byte format
// 		for _, w := range words {
// 			distance := levenshtein.DistanceForStrings([]rune(strings.ToLower(copticWord)), []rune(strings.ToLower(w.CopticWord)), levenshtein.DefaultOptions)
// 			if distance < minDistance {
// 				minDistance = distance
// 				closestWord = w.CopticWord
// 			}
// 		}
//
// 		//Define a threshold for a "close enough" match
// 		//rune --> unicode character, converts the string to list of runes (characters)
// 		//len gives the number of characters instead of bytes as it does by default
// 		threshold := max(2, len([]rune(copticWord))/3)
//
// 		//Return closest word or no match found
// 		if minDistance <= threshold {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Word not found", "suggestion": closestWord})
// 		} else {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Word not found", "suggestion": "No close match found"})
// 		}
// 	}
// }

//Find max value between 2 numbers function
//Needed because math.Max() only works with float64 values, not int..
//..so would need to cast as float64 then turn back to int; unnecessary
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
