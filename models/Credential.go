package models

// Admin ...
type Admin struct {
	BaseModel
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	BranchID  uint   `json:"branch_id"`
	SessionID uint   `json:"session_id"`
	Token     string `json:"token"`
}

// Session ...
type Session struct {
	BaseModel
	Admin Admin  `json:"admin"`
	Token string `json:"-"`
}
