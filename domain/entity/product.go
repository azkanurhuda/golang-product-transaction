package entity

type Product struct {
	ID                string `gorm:"primaryKey"`
	ProductCategoryID string `gorm:"foreignKey"`
	Name              string
	Description       string
	Price             int64
	Stock             int64
}
