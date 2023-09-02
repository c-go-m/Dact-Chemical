package entity

type Presentation struct {
	EntityBase
	Name       string   `gorm:"unique;not null" json:"name"`
	AttachedID int      `json:"attachedId"`
	Attached   Attached `json:"attached"`
	ProductID  int      `json:"productId"`
	Product    Product  `json:"product"`
	Position   int      `gorm:"unique;not null" json:"position"`
	Cost       int      `json:"cost"`
}

func (Presentation) TableName() string {
	return "admin.Presentation"
}
