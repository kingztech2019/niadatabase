package model

import "gorm.io/gorm"


type ImagesUrl struct{
	gorm.Model
	FirstImage string `json:firstimage`
	UserID float64  `json:"userid"` 
	SecondImage string `json:"secondimage"`
	ThirdImage  string `json:"thirdimage"`
	FourthImage string `json:"fourthimage"`
}