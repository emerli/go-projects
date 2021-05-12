package models

type CountryLanguage struct {
	Countrycode string  `json:"countrycode" gorm:"column:countrycode"`
	Percentage  float64 `json:"percentage" gorm:"column:percentage"`
	Language    string  `json:"language" gorm:"column:language"`
	Isofficial  bool    `json:"isofficial" gorm:"column:isofficial"`
}
