output "fqdn" {
  description = "The fqdn of MySQL Server."
  value       = "${azurerm_mysql_server.mysql.fqdn}"
}

output "server_name" {
  description = "The server name of MySQL Server."
  value       = "${azurerm_mysql_server.mysql.name}"
}

output "admin_username" {
  description = "The administrator username of MySQL Database."
  value       = "${var.admin_username}"
}

output "password" {
  description = "The administrator password of the MySQL Database."
  value       = "${var.password}"
}

output "database_name" {
  description = "The database name of MySQL Database."
  value       = "${var.db_name}"
}