package models

type Products struct {
	ID     uint32 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name   string `gorm:"column:name;not null" json:"name"`
	Brand  string `gorm:"column:brand;not null" json:"brand"`
	Sku    string `gorm:"column:SKU;not null" json:"sku"`
	Size   string `gorm:"column:size;not null" json:"size"`
	Color  string `gorm:"column:color;not null" json:"color"`
	Status bool   `gorm:"column:status;not null" json:"status"`
	Stock  uint32 `gorm:"column:stock;not null" json:"stock"`
}
