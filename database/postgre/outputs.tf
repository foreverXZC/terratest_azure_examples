output "fqdn" {
  value = "${azurerm_postgresql_server.postgresql_server.fqdn}"
}

output "resource_group_name" {
  value = "${azurerm_postgresql_server.postgresql_server.resource_group_name}"
}

output "server_name" {
  value = "${azurerm_postgresql_server.postgresql_server.name}"
}

output "admin_username" {
  value = "${var.db_admin_username}"
}

output "password" {
  value = "${var.db_admin_password}"
}