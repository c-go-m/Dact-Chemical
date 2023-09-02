# Storage Account
variable "sa-name"{
  description = "Nombre del grupo de recursos"
  nullable = false
  type = string
}

variable "sa-resource-group"{
  description = "Nombre del grupo de recursos"
  nullable = false
  type = string
}

variable "sa-location"{
  description = "Localizacion del grupo de recursos"
  nullable = false
  type = string
  default = "East Us"
  validation {
    condition = contains(["East Us", "East US 2"],var.rg-location)
    error_message = "La localización del grupo de recursos solo esta permitida en las siguientes zonas: ['East Us', 'East US 2'] ."
  }
}

variable "sa-tags" {    
    description = "Tags de la aplicacion"
    nullable = false
    type = map(string)
    default = {}
}