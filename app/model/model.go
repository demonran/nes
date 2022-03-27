package model

import (
	"gorm.io/gorm"
	"time"
)

type Entity interface {
	SetId(id string)
	GetId() string
}

type BModel struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *BModel) SetId(id string) {
	m.ID = id
}

func (m *BModel) GetId() string {
	return m.ID
}

type Customer struct {
	BModel
	Company string `json:"company"`
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
}
