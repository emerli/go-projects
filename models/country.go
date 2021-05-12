package models

type Country struct {
	Continent      string      `json:"continent" gorm:"column:continent"`
	Capital        int         `json:"capital" gorm:"column:capital"`
	Code2          string      `json:"code2" gorm:"column:code2"`
	Code           string      `json:"code" gorm:"column:code"`
	Localname      string      `json:"localname" gorm:"column:localname"`
	Gnp            float64     `json:"gnp" gorm:"column:gnp"`
	Headofstate    string      `json:"headofstate" gorm:"column:headofstate"`
	Population     int         `json:"population" gorm:"column:population"`
	Lifeexpectancy float64     `json:"lifeexpectancy" gorm:"column:lifeexpectancy"`
	Governmentform string      `json:"governmentform" gorm:"column:governmentform"`
	Name           string      `json:"name" gorm:"column:name"`
	Indepyear      int         `json:"indepyear" gorm:"column:indepyear"`
	Gnpold         interface{} `json:"gnpold" gorm:"column:gnpold"`
	Region         string      `json:"region" gorm:"column:region"`
	Surfacearea    int         `json:"surfacearea" gorm:"column:surfacearea"`
}
