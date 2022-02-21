package model

import "gorm.io/gorm"

type VerifyCode struct{
	gorm.Model
	IdentityCode string `json:identitycode`
	UserID float64  `json:"userid"` 
	 
}