package models

import (
	"gorm.io/gorm"
)

type City struct {
	Countrycode string `json:"countrycode" gorm:"column:countrycode"`
	District    string `json:"district" gorm:"column:district"`
	Name        string `json:"name" gorm:"column:name"`
	ID          int    `json:"id" gorm:"column:id;primaryKey;default: nextval('city_id_seq')""`
	Population  int    `json:"population" gorm:"column:population"`


}

func (a *City) Insert(value *City, db *gorm.DB) (int, error) {
	var city City
	result := db.Model(&city).Create(value)
	return value.ID, result.Error
}
func (a *City) GetAll(db *gorm.DB) ([]City, error) {
	var cities []City

	table := db.Table("city").Scan(&cities)
	return cities, table.Error
}

func (a *City) Get(id int, db *gorm.DB) (City, error) {
	var city City

	table := db.Table("city").Where("id = ?", id).Scan(&city)
	return city, table.Error
}

func (a *City) Update(value *City, db *gorm.DB) error {
	result := db.Updates(value)
	return result.Error
}

func (a *City) Delete(id int, db *gorm.DB) error {
	var city City
	result := db.Delete(city, id)
	return result.Error
}
