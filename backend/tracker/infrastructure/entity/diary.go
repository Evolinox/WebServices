package entity

type DailyProductsConsumed struct {
	ID       uint              `json:"ID" gorm:"primaryKey"`
	Date     string            `json:"Date" gorm:"uniqueIndex"`
	Products []ConsumedProduct `json:"Products" gorm:"foreignKey:DailyProductsConsumedID;constraint:OnDelete:CASCADE"`
}
