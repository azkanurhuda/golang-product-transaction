package entity

type Cart struct {
	ID        string `gorm:"primaryKey"`
	UserID    string `gorm:"foreignKey"`
	ProductID string `gorm:"foreignKey"`
}

type Order struct {
	ID     string `gorm:"primaryKey"`
	CartID string `gorm:"foreignKey"`
	UserID string `gorm:"foreignKey"`
}
