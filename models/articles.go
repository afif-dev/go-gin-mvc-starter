package models

import (
	"time"

	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	ID			uint 			`gorm:"primaryKey;autoIncrement"`
	Title		string 			`gorm:"not null;size:255"`
	Short_Desc	*string			`gorm:"size:255"`
	Content		*string			`gorm:"type:TEXT"`
	Tags		*string			`gorm:"size:255"`
	CreatedAt   time.Time		`gorm:"autoCreateTime"`
 	UpdatedAt	time.Time		`gorm:"autoCreateTime;autoUpdateTime"`
	DeletedAt 	*time.Time		
}

type Articles_Select struct {
	ID   uint
	Title string
	Short_Desc	string
	Tags		string
	CreatedAt   time.Time
 	UpdatedAt	time.Time
}