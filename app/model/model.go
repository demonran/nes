package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Company string `json:"company"`
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
}
