package domain

type User struct {	
	Id       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}