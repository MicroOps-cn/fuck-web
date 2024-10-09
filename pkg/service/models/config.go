package models

type SystemConfig struct {
	Name  string `json:"name" gorm:"primary_key;type:varchar(100)"`
	Type  string `json:"type" gorm:"type:varchar(10)"`
	Value string `json:"value" gorm:"type:varchar(50)"`
}
