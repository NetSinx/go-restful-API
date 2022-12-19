package domain

type Product struct {
	ID    int    `json:"id" gorm:"primaryKey, autoIncrement"`
    Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}