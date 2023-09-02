package entity

type Information struct {
	EntityBase
	Name       string   `gorm:"unique;not null" json:"name"`
	Value      string   `gorm:"not null" json:"value"`
	AttachedID int      `json:"attachedId"`
	Attached   Attached `json:"attached"`
	MenuID     int      `json:"menuId"`
	Menu       Menu     `json:"menu"`
	Position   int      `gorm:"unique;not null" json:"position"`
}

func (Information) TableName() string {
	return "admin.Information"
}
