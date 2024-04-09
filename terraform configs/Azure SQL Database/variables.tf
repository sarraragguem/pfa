variable "resource_group_name" {
  description = "The name of the resource group."
  type        = string
}

variable "location" {
  description = "The location for the resources."
  type        = string
}

variable "sql_server_name" {
  description = "The name of the SQL server."
  type        = string
}

variable "sql_server_admin_login" {
  description = "The admin login for the SQL server."
  type        = string
}

variable "sql_server_admin_password" {
  description = "The admin password for the SQL server."
  type        = string
  sensitive   = true
}

variable "sql_database_name" {
  description = "The name of the SQL database."
  type        = string
}

variable "sql_database_collation" {
  description = "The collation for the SQL database."
  type        = string
}

variable "sql_database_max_size_gb" {
  description = "The max size of the SQL database in gigabytes."
  type        = number
}

variable "sql_database_sku_name" {
  description = "The SKU name for the SQL database."
  type        = string
}

variable "sql_database_zone_redundant" {
  description = "Whether the SQL database is zone redundant."
  type        = bool
}

variable "sql_firewall_rules" {
  description = "A list of firewall rules for the SQL server."
  type        = list(object({
    name              = string
    start_ip_address  = string
    end_ip_address    = string
  }))
  default     = []
}
