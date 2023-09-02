package database

import (
	"github.com/c-go-m/DC-Admin/common/entity"
	"github.com/c-go-m/DC-Admin/common/utility/useError"
)

func (d *DataBase) initialDatabase() error {
	if err := d.crateSchemas(); err != nil {
		return useError.ErrorConnectionDataBase
	}

	if err := d.initialMigration(); err != nil {
		return useError.ErrorConnectionDataBase
	}

	if err := d.insertDataMenus(); err != nil {
		return useError.ErrorConnectionDataBase
	}

	return nil
}

func (d *DataBase) initialMigration() error {
	err := d.Db.AutoMigrate(entity.Entities()...)

	if err != nil {
		return useError.ErrorConnectionDataBase
	}
	return nil
}

func (d *DataBase) crateSchemas() error {
	return d.Db.Exec("CREATE SCHEMA IF NOT EXISTS admin").Error
}

func (d *DataBase) insertDataMenus() error {
	var menus = []entity.Menu{
		{Name: "Inicio", Rute: "/client/initial", Type: "client", Position: 1},
		{Name: "Productos", Rute: "/client/product", Type: "client", Position: 2},
		{Name: "Servicios", Rute: "/client/service", Type: "client", Position: 3},
		{Name: "Contactenos", Rute: "/client/contact", Type: "client", Position: 4},
		{Name: "Banners", Rute: "/admin/banner", Type: "admin", Position: 1},
		{Name: "Categorias", Rute: "/admin/category", Type: "admin", Position: 2},
		{Name: "Productos", Rute: "/admin/product", Type: "admin", Position: 3},
		{Name: "Presentaciones", Rute: "/admin/presentation", Type: "admin", Position: 4},
		{Name: "Información", Rute: "/admin/information", Type: "admin", Position: 5},
	}

	d.Db.Create(&menus)
	return nil
}
