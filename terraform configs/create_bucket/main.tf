provider "google" {
  project = var.project_id
}

resource "google_storage_bucket" "advanced_bucket" {
  name          = var.bucket_name
  location      = var.location
  storage_class = var.storage_class
  force_destroy = var.force_destroy

   versioning {
    enabled = var.versioning_enabled
  }

 
  

  logging {
    log_bucket        = var.logging["log_bucket"]
    log_object_prefix = var.logging["log_object_prefix"]
  }
}
