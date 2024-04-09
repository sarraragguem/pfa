provider "google" {
  
  project = var.project_id
  region  = var.region
}

resource "google_sql_database_instance" "default" {
  name             = var.instance_name
  region           = var.region
  database_version = var.db_version

  settings {
    tier             = var.tier
    activation_policy = "ALWAYS"
    availability_type = var.availability_type

    disk_autoresize = true
    disk_size       = var.storage_size
    disk_type       = var.storage_type

    backup_configuration {
      enabled            = true
      start_time         = var.backup_start_time
      binary_log_enabled = var.db_version == "MYSQL_8_0" ? true : false
    }

    maintenance_window {
      day          = var.maintenance_window_day
      hour         = var.maintenance_window_hour
      update_track = "stable"
    }

    insights_config {
      query_insights_enabled = var.insights_config_query_insights_enabled
    }

    dynamic "database_flags" {
      for_each = var.db_version == "MYSQL_8_0" ? [1] : []
      content {
        name  = "log_bin_trust_function_creators"
        value = "on"
      }
    }
  }
}
