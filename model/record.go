package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Record struct {
	gorm.Model
	UserId    uint
	EventTime time.Time
	Comment   string
	Amount    float64
	Tags      string
}
