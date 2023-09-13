package entity

type Payment struct {
	ID          string `gorm:"primaryKey"`
	OrderID     string `gorm:"foreignKey"`
	UserID      string `gorm:"foreignKey"`
	PaymentType string
}
