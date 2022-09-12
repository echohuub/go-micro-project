package model

type ProductSeo struct {
	ID          int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	Title       string `json:"title"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Code        string `gorm:"unique_index;not_null" json:"code"`
	ProductId   int64  `json:"product_id"`
}
