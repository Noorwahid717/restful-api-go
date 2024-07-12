package model

import (
	"time"
)

type Event struct {
	BaseModel
	Title       string    `gorm:"type:string;size:100;not null"`
	Description string    `gorm:"type:string;size:255;null"`
	StartDate   time.Time `gorm:"type:datetime;not null"`
	EndDate     time.Time `gorm:"type:datetime;not null"`
	Location    string    `gorm:"type:string;size:255;null"`
	Company     Company
	CompanyId   int
}
