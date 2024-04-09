variable "project_id" {
  description = "The GCP project ID."
  type        = string
}

variable "region" {
  description = "The region for the Cloud SQL instance."
  type        = string
  default     = "us-central1"
}

variable "instance_name" {
  description = "Name of the Cloud SQL instance."
  type        = string
  default     = "my-sql-instance"
}

variable "db_version" {
  description = "The version of the database engine."
  type        = string
  default     = "MYSQL_8_0"
}

variable "tier" {
  description = "The tier (machine type) for the instance."
  type        = string
  default     = "db-f1-micro"
}

variable "storage_type" {
  description = "The storage type for the instance."
  type        = string
  default     = "SSD"
}

variable "storage_size" {
  description = "Initial storage size in GB."
  type        = number
  default     = 10
}

variable "availability_type" {
  description = "The availability type of the SQL instance."
  type        = string
  default     = "ZONAL"
}

variable "maintenance_window_day" {
  description = "Day of the week for maintenance window (1-7)."
  type        = number
  default     = 7  # Sunday
}

variable "maintenance_window_hour" {
  description = "Hour of the day for maintenance window (0-23)."
  type        = number
  default     = 0
}

variable "backup_start_time" {
  description = "Start time for the daily backup configuration in HH:MM format."
  type        = string
  default     = "05:00"
}

variable "insights_config_query_insights_enabled" {
  description = "Enables query insights feature."
  type        = bool
  default     = true
}
