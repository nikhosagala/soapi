package models

// Company ...
type Company struct {
	BaseModel
	BasicInfo
	Branches []Branch `json:"branches"`
}

// Branch ...
type Branch struct {
	BaseModel
	BasicInfo
	CompanyID uint  `json:"company_id" binding:"required"`
	Admin     Admin `json:"admin"`
}
