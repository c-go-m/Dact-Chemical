# General

variable "location" {
  description = "Localizacion general del los recursos creados"
  nullable = false
  type = string
  default = "East Us"
  validation {
    condition = contains(["East Us", "East US 2"],var.location)
    error_message = "La localización solo esta permitida en las siguientes zonas: ['East Us', 'East US 2'] ."
  }
}

variable "proyect" {
  description = "Nombre del proyecto"
  nullable = false
  type = string
  default = "proyect"
}

variable "environment" {
  description = "Nombre del ambiente"
  nullable = false
  type = string
  default = "dllo"
}

variable "tags" {    
    description = "Tags de la aplicacion"
    nullable = false
    type = map(string)
    default = {}
}