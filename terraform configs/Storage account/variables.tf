# variables.tf for Storage Account

variable "resource_group_name" {
  description = "The name of the resource group."
  type        = string
}

variable "location" {
  description = "The location for the resources."
  type        = string
}

variable "storage_account_name" {
  description = "Specifies the name of the storage account."
  type        = string
}

variable "account_tier" {
  description = "Defines the Tier to use for this storage account."
  type        = string
}

variable "replication_type" {
  description = "Defines the type of replication to use for this storage account."
  type        = string
}

variable "tags" {
  description = "A mapping of tags to assign to the resource."
  type        = map(string)
  default     = {}
}
