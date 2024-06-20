package models

import (
	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	Text string
}
