package model

type Auth struct {
	Email    string `form:"email" binding:"required" json:"email"`
	Password string `form:"password" binding:"required" json:"password"`
}
