package model

import "gorm.io/gorm"


type UploadStatus struct{
	gorm.Model
	UploadStatus string `json:upload_status`
	UserID float64  `json:"userid"` 
	IdentityCode string `json:identitycode`
	 
}