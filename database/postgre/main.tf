terraform {
  required_version = "~> 0.11.1"
}

provider "azurerm" {
  version = "~> 1.0.1"
}

resource "azurerm_resource_group" "resource_group" {
  location = "${var.location}"
  name     = "${var.service_name}-${var.environment}-postgresql-rg"
  count    = "${var.resource_group_name == "" ? 1 : 0}"
}

resource "azurerm_postgresql_server" "postgresql_server" {
  name                = "${var.service_name}${var.environment}pgsql"
  location            = "${var.location}"
  resource_group_name = "${var.resource_group_name == "" ? azurerm_resource_group.resource_group.name : var.resource_group_name}"

  sku {
    name     = "PGSQL${substr(var.azure_postgres_sku_tier, 0, 1)}${var.sku_compute_units}"
    capacity = "${var.sku_compute_units}"
    tier     = "${var.azure_postgres_sku_tier}"
  }

  administrator_login          = "${var.db_admin_username}"
  administrator_login_password = "${var.db_admin_password}"
  version                      = "${var.postgres_version}"
  storage_mb                   = "${var.db_disk_size_mb}"
  ssl_enforcement              = "${var.enforce_ssl}"
}

resource "azurerm_postgresql_firewall_rule" "postgres_firewall_rule" {
  name                = "${var.rule_name}"
  resource_group_name = "${azurerm_resource_group.resource_group.name}"
  server_name         = "${azurerm_postgresql_server.postgresql_server.name}"
  start_ip_address    = "${var.start_ip_address}"
  end_ip_address      = "${var.end_ip_address}"
}