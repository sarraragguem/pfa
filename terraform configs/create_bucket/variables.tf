variable "project_id" {
  description = "The GCP project ID."
  type        = string
}

variable "bucket_name" {
  description = "The name of the bucket. It must be globally unique."
  type        = string
}

variable "location" {
  description = "The location of the bucket. Default to US multi-region."
  type        = string
  default     = "US"
}

variable "storage_class" {
  description = "The storage class of the bucket. Default to STANDARD."
  type        = string
  default     = "STANDARD"
}

variable "versioning_enabled" {
  description = "Enable versioning for the bucket."
  type        = bool
  default     = false
}

variable "force_destroy" {
  description = "Set to true to force bucket deletion even if it contains objects."
  type        = bool
  default     = false
}

variable "logging" {
  description = "The logging configuration."
  type        = map(string)
  default     = {
    log_bucket        = ""
    log_object_prefix = ""
  }
}

