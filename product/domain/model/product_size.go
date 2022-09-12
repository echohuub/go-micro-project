package model

type ProductSize struct {
	ID        int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	Name      string `json:"name"`
	Code      string `gorm:"unique_index;not_null" json:"code"`
	ProductId int64  `json:"product_id"`
}
