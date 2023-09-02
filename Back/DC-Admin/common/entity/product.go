package entity

type Product struct {
	EntityBase
	Name         string         `gorm:"unique;not null" json:"name" validate:"required"`
	AttachedID   int            `json:"attachedId"`
	Attached     Attached       `json:"attached"`
	CategoryID   int            `json:"categoryId" validate:"required"`
	Category     Category       `json:"category"`
	Presentation []Presentation `json:"presentation"`
}

func (Product) TableName() string {
	return "admin.Product"
}
