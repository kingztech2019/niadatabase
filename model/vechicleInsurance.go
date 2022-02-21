package model

import "gorm.io/gorm"

type VechicleInsurance struct{
	gorm.Model
	ClassOfInsurance string `json:classofinsurance`
	UserID float64  `json:"userid"` 
	Type string `json:"type"`
	VechicleUse  string `json:"vechicleuse"`
}