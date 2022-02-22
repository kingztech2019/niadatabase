package model

import "gorm.io/gorm"


type PersonalDetails struct{
	gorm.Model
  	UserID float64  `json:"userid"` 
	MeansOfId string `json:"means_of_id"`
	Id string `json:"id"`
	Title string `json:"title"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	PhoneNumber string `json:"phone_number"`
	Email string `json:"email"`
	State string `json:"state"`
	Lga string `json:"lga"`
	Address string `json`
	 
}