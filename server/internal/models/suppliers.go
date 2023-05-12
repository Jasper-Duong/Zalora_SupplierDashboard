package models

type Suppliers struct {
	ID            uint32 `gorm:"primary_key"`
	Name          string
	Email         string
	ContactNumber string
	Status        bool
}
