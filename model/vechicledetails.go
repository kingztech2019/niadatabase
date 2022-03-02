package model

import "gorm.io/gorm"


type VechicleDetails struct{
	gorm.Model
  	UserID float64  `json:"userid"` 
	PlateNo string `json:"plate_no"`
	Vin string `json:"vin"`
	Engine string `json:"engine"`
	VechicleColor string `json:"vechicle_color"`
	Make string `json:"make"`
	Modell string `json:"model"`
	Value string `json:"phone_number"`
	Capacity string `json:"capacity"`
	 
	 
}