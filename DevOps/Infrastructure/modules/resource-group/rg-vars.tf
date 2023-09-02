# Resource group
variable "rg-name"{
  description = "Nombre del grupo de recursos"
  nullable = false
  type = string
}

variable "rg-location"{
  description = "Localizacion del grupo de recursos"
  nullable = false
  type = string
  default = "East Us"
  validation {
    condition = contains(["East Us", "East US 2"],var.rg-location)
    error_message = "La localización del grupo de recursos solo esta permitida en las siguientes zonas: ['East Us', 'East US 2'] ."
  }
}

variable "rg-tags" {    
    description = "Tags de la aplicacion"
    nullable = false
    type = map(string)
    default = {}
}