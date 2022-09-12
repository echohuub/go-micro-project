package model

type Product struct {
	ID          int64          `gorm:"primary_key;not_null;auto_increment" json:"id"`
	Name        string         `json:"name"`
	Sku         string         `gorm:"unique_index:not_null" json:"sku"`
	Price       float32        `json:"price"`
	Description string         `json:"description"`
	Image       []ProductImage `gorm:"ForeignKey:ImageProductID" json:"image"`
	Size        []ProductSize  `gorm:"ForeignKey:SizeProductID" json:"size"`
	Seo         ProductSeo     `gorm:"ForeignKey:SeoProductID" json:"seo"`
}
