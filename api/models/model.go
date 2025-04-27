package models

//Equivalent to PostgreSQL's database structure
//uint64 = unsigned 64 bit integer (no negatives)
//`key:"value"` = format for struct tags
//key value must be CamelCase; Go automatically transform it..
//..however force the name with struct tags, such as: `json: "coptic_word"`..
//..returns response in JSON format
type CopticDictionary struct {
	ID                      uint64 `gorm:"primaryKey" json:"id"`
	CopticWord              string `gorm:"uniqueIndex" json:"coptic_word"`
	ArabicTranslation       string `json:"arabic_translation"`
	EnglishTranslation      string `json:"english_translation"`
	GreekScriptCopticWord   string `json:"greek_script_coptic_word"`
	GrammaticalModification string `json:"grammatical_modification"`
	WordCategory            string `json:"word_category"`
	WordGender              string `json:"word_gender"`
	WordCategory2           string `json:"word_category_2"`
	WordGender2             string `json:"word_gender_2"`
	GreekWord               string `json:"greek_word"`
	CopticWordAlt           string `json:"coptic_word_alt"`
}

//Force the correct name of the table
func (CopticDictionary) TableName() string {
	return "coptic_dictionary"
}