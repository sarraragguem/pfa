provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "example" {
  name     = var.resource_group_name
  location = var.location
}

resource "azurerm_mssql_server" "example" {
  name                         = var.sql_server_name
  resource_group_name          = azurerm_resource_group.example.name
  location                     = azurerm_resource_group.example.location
  version                      = "12.0"
  administrator_login          = var.sql_server_admin_login
  administrator_login_password = var.sql_server_admin_password
}

resource "azurerm_mssql_database" "example" {
  name                = var.sql_database_name
  server_id           = azurerm_mssql_server.example.id
  collation           = var.sql_database_collation
  max_size_gb         = var.sql_database_max_size_gb
  sku_name            = var.sql_database_sku_name  # Make sure this is aligned with Azure's valid SKUs
  zone_redundant      = var.sql_database_zone_redundant
}

resource "azurerm_mssql_firewall_rule" "example" {
  count               = length(var.sql_firewall_rules)
  name                = var.sql_firewall_rules[count.index]["name"]
  server_id           = azurerm_mssql_server.example.id
  start_ip_address    = var.sql_firewall_rules[count.index]["start_ip_address"]
  end_ip_address      = var.sql_firewall_rules[count.index]["end_ip_address"]
}
