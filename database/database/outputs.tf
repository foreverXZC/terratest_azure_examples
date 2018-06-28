output "database_name" {
    value = "${module.sql-database.database_name}"
}

output "sql_server_name" {
    value = "${module.sql-database.sql_server_name}"
}

output "sql_server_fqdn" {
    value = "${module.sql-database.sql_server_fqdn}"
}