package models

import (
	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	id   string
	text string
}
