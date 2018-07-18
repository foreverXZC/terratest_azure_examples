variable "resource_group_name" {
  description = "Default resource group name that the database will be created in."
  default     = "mysqlResourceGroup"
}

variable "location" {
  description = "Specifies the supported Azure location where the resource exists."
}

variable "db_name" {
  description = "The name of the database to be created."
}

variable "admin_username" {
  description = "The Administrator Login for the MySQL Server."
}

variable "password" {
  description = "The Password associated with the administrator_login for the MySQL Server."
}

variable "start_ip_address" {
  description = "Defines the start IP address used in your database firewall rule."
  default     = "0.0.0.0"
}

variable "end_ip_address" {
  description = "Defines the end IP address used in your database firewall rule."
  default     = "255.255.255.255"
}

variable "db_version" {
  description = "Specifies the version of MySQL to use. Valid values are 5.6 and 5.7."
  default     = "5.7"
}

variable "ssl_enforcement" {
  description = "Specifies if SSL should be enforced on connections. Possible values are Enforced and Disabled."
  default     = "Enabled"
}

variable "sku_name" {
  description = "Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the tier + family + cores pattern (e.g. B_Gen4_1, GP_Gen5_8)."
  default     = "B_Gen4_2"
}

variable "sku_capacity" {
  description = "The scale up/out capacity, representing server's compute units."
  default     = 2
}

variable "sku_tier" {
  description = "The tier of the particular SKU. Possible values are Basic, GeneralPurpose, and MemoryOptimized."
  default     = "Basic"
}

variable "sku_family" {
  description = "The family of hardware Gen4 or Gen5, before selecting your family check the product documentation for availability in your region."
  default     = "Gen4"
}

variable "storage_mb" {
  description = "Max storage allowed for a server. Possible values are between 5120 MB(5GB) and 1048576 MB(1TB) for the Basic SKU and between 5120 MB(5GB) and 4194304 MB(4TB) for General Purpose/Memory Optimized SKUs."
  default     = 5120
}

variable "backup_retention_days" {
  description = "Backup retention days for the server, supported values are between 7 and 35 days."
  default     = 7
}

variable "geo_redundant_backup" {
  description = "Enable Geo-redundant or not for server backup. Valid values for this property are Enabled or Disabled, not supported for the basic tier."
  default     = "Disabled"
}

variable "charset" {
  description = "Specifies the Charset for the MySQL Database, which needs to be a valid MySQL Charset."
  default     = "utf8"
}

variable "collation" {
  description = "Specifies the Collation for the MySQL Database, which needs to be a valid MySQL Collation."
  default     = "utf8_unicode_ci"
}
