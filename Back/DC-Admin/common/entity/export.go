package entity

func Entities() []interface{} {
	var entities []interface{}

	entities = append(entities,
		&Attached{},
		&Banner{},
		&Menu{},
		&Category{},
		&Product{},
		&Presentation{},
		&Information{},
	)

	return entities
}
