# Database Example

Using terraform azure module "database" to deploy a Microsoft SQL Dababase on azure.

## Database

These terraform files enable users to deploy a Microsoft SQL Dababase on azure. Because database, server and firewall rules are all indicated in the module, it is possible to connect to the database from local and execute SQL commands. You can just test the infrastructure code manually without terratest.

## SQL

This folder only includes one file. First, the go test file uses terraform database module to deploy a database on azure. After that, it tries to connect to the database and execute several SQL commands to check whether the infrastructure runs correctly. Eventually, everything will be cleaned up after validation. You can write your own test code, for instance, to create an independent SQL file and use it in test code.

## Reference

[Terraform Azure Database Module](https://registry.terraform.io/modules/Azure/database/azurerm)

[Go Database Tutorial](http://go-database-sql.org)

[SQL Server Connection in Golang](https://mathaywardhill.com/2017/04/27/get-started-with-golang-and-sql-server-in-visual-studio-code)

[Azure SQL Database Document](https://docs.microsoft.com/en-us/azure/sql-database)