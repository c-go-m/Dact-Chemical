resource "azurerm_key_vault" "key-vault" {
  name     = var.rg-name
  location = var.rg-location
  tags = var.rg-tags
}