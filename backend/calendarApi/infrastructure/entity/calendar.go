package entity

type Calendar struct {
	ID          int    `json:"ID" gorm:"primaryKey"`
	Date        string `json:"Date"`
	BeginTime   string `json:"BeginTime"`
	EndTime     string `json:"EndTime"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}
