package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	UserId  uint
	TagName string
}
