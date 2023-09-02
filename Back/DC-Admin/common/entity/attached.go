package entity

type Attached struct {
	EntityBase
	Name      string `gorm:"not null" json:"name" validate:"required"`
	Url       string `gorm:"unique;not null" json:"url"`
	Extension string `gorm:"not null" json:"extension" validate:"required"`
	Data      []byte `gorm:"-" json:"data" validate:"required"`
}

func (Attached) TableName() string {
	return "admin.Attached"
}
