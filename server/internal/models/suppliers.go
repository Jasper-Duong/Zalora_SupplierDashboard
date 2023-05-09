package models

type Suppliers struct {
	ID            uint64 `gorm:"primary_key"`
	Name          string
	Email         string
	ContactNumber string
	Status        bool
	Products      []*Products `gorm:"many2many:products_suppliers"`
}
