resource_group_name         = "sarsourataggurm" #uniqur
location                    = "East US"
sql_server_name             = "sarouraserver"  #unique 
sql_server_admin_login      = "adminuser"
sql_server_admin_password   = "SuperSecretPassword123!"  # Be sure to change this to a secure password!
sql_database_name           = "example-sql-db"
sql_database_collation      = "SQL_Latin1_General_CP1_CI_AS"
sql_database_max_size_gb    = 10
sql_database_sku_name       = "GP_Gen5_2"
sql_database_sku_capacity   = 2
sql_database_sku_tier       = "GeneralPurpose"
sql_database_sku_family     = "Gen5"
sql_database_zone_redundant = false
sql_firewall_rules          = [
  {
    name              = "AllowAllWindowsAzureIps"
    start_ip_address  = "0.0.0.0"
    end_ip_address    = "0.0.0.0"
  },
  # Add additional firewall rules as needed
]
