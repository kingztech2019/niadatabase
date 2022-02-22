package model

import "gorm.io/gorm"


type PolicyDetails struct{
	gorm.Model
  	UserID float64  `json:"userid"` 
	PolicyHolder string `json:"policy_holder"`
	PhoneNumber string `json:"phone_number"`
	Email string `json:"email"`
	Company string `json:"company"`
	Nin string `json:"nin"`
	State string `json:"state"`
	Lga string `json:"lga"`
	Address string `json:"address"`
	 
	 
}