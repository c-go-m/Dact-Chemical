# Proveedor Azure
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.29.1"
    }
  }
}

# Configuracion Microsoft Azure Provider
provider "azurerm" {
  # Configuration options
  features {}
}

# Create a resource group
module "resource-group" {
  source = "./modules/resource-group"
  rg-name     = "rg-${var.proyect}-${var.environment}"
  rg-location = var.location
  rg-tags = var.tags
}

