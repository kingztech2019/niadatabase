package model

import "gorm.io/gorm"

type PasswordToken struct{
	gorm.Model
	 
	Token string `json:"token"`
	UserID uint  `json:"userid"` 
	Used bool
	 
	 	 
}