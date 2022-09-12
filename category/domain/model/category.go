package model

type Category struct {
	ID          int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	Name        string `gorm:"unique_index,not_null" json:"name"`
	Level       uint32 `json:"level"`
	Parent      int64  `json:"parent"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
