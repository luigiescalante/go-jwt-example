package model

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id        int32          `json:"id,primary_key" gorm:"Column:id"`
	FirstName string         `json:"firstName" gorm:"size:100;Column:first_name"`
	LastName  string         `json:"lastName" gorm:"size:100;Column:last_name"`
	Email     string         `json:"email" gorm:"size:50;Column:email"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:null"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
