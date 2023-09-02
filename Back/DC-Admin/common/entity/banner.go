package entity

type Banner struct {
	EntityBase
	Name       string   `gorm:"unique;not null" json:"name" validate:"required"`
	AttachedID int      `json:"attachedId"`
	Attached   Attached `json:"attached"`
	Position   int      `gorm:"unique;not null" json:"position" validate:"required"`
}

func (Banner) TableName() string {
	return "admin.Banner"
}
