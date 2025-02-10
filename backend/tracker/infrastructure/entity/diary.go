package entity

type DailyProductsConsumed struct {
	ID       uint              `json:"ID" gorm:"primaryKey;autoIncrement"`
	Date     string            `json:"Date" gorm:"uniqueIndex"`
	Products []ConsumedProduct `json:"Products" gorm:"foreignKey:DailyID"`
}
