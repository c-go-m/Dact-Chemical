package entity

type Menu struct {
	EntityBase
	Name     string `validate:"required" gorm:"not null" json:"name"`
	Rute     string `validate:"required" gorm:"unique;not null" json:"rute"`
	Type     string `validate:"required" gorm:"not null" json:"type"`
	Position int    `validate:"required" gorm:"not null" json:"position"`
}

func (Menu) TableName() string {
	return "admin.Menu"
}
