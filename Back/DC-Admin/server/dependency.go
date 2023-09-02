package server

import (
	Database "github.com/c-go-m/DC-Admin/common/utility/database"
	Storage "github.com/c-go-m/DC-Admin/common/utility/storage/azureStorage"
	FileSystem "github.com/c-go-m/DC-Admin/common/utility/storage/fileSystem"
	ControlError "github.com/c-go-m/DC-Admin/common/utility/useError"
	AttachedRepository "github.com/c-go-m/DC-Admin/repository/attached"
	BannerRepository "github.com/c-go-m/DC-Admin/repository/banner"
	CategoryRepository "github.com/c-go-m/DC-Admin/repository/category"
	InformationRepository "github.com/c-go-m/DC-Admin/repository/information"
	MenuRepository "github.com/c-go-m/DC-Admin/repository/menu"
	PresentationRepository "github.com/c-go-m/DC-Admin/repository/presentation"
	ProductRepository "github.com/c-go-m/DC-Admin/repository/product"

	AttachedService "github.com/c-go-m/DC-Admin/service/attached"
	BannerService "github.com/c-go-m/DC-Admin/service/banner"
	CategoryService "github.com/c-go-m/DC-Admin/service/category"
	InformationService "github.com/c-go-m/DC-Admin/service/information"
	MenuService "github.com/c-go-m/DC-Admin/service/menu"
	PresentationService "github.com/c-go-m/DC-Admin/service/presentation"
	ProductService "github.com/c-go-m/DC-Admin/service/product"
)

var database *Database.DataBase
var fileSystem *FileSystem.FileSystem
var storage *Storage.AzureStorage
var bannerRepository *BannerRepository.BannerRepository
var bannerService *BannerService.BannerService
var categoryRepository *CategoryRepository.CategoryRepository
var categoryService *CategoryService.CategoryService
var menuRepository *MenuRepository.MenuRepository
var menuService *MenuService.MenuService
var attachedRepository *AttachedRepository.AttachedRepository
var attachedService *AttachedService.AttachedService
var productRepository *ProductRepository.ProductRepository
var productService *ProductService.ProductService
var presentationRepository *PresentationRepository.PresentationRepository
var presentationService *PresentationService.PresentationService
var informationRepository *InformationRepository.InformationRepository
var informationService *InformationService.InformationService

func initDependency() {
	initDatabase()
	initStorage()
	initRepositories()
	initServices()
}

func initDatabase() {
	var err error
	database, err = Database.New()
	ControlError.ControlError(err)
}

func initStorage() {
	var err error
	storage, err = Storage.New()
	ControlError.ControlError(err)
	fileSystem = FileSystem.New()
}

func initRepositories() {
	bannerRepository = BannerRepository.New(database)
	attachedRepository = AttachedRepository.New(database)
	categoryRepository = CategoryRepository.New(database)
	menuRepository = MenuRepository.New(database)
	productRepository = ProductRepository.New(database)
	presentationRepository = PresentationRepository.New(database)
	informationRepository = InformationRepository.New(database)
}

func initServices() {
	attachedService = AttachedService.New(attachedRepository, storage)
	menuService = MenuService.New(menuRepository)
	bannerService = BannerService.New(bannerRepository, attachedService)
	categoryService = CategoryService.New(categoryRepository, attachedService)
	productService = ProductService.New(productRepository, attachedService)
	presentationService = PresentationService.New(presentationRepository, attachedService)
	informationService = InformationService.New(informationRepository, attachedService)
}
