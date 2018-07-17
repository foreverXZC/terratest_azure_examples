variable "location" {
  description = "(Required) Azure location string - see 'az account list-locations' output for valid strings"
  default     = "westeurope"
}

variable "service_name" {
  description = "(Required) Service this resource belongs to"
  default     = "server12345"
}

variable "environment" {
  description = "(Optional) The working environment this resouce belongs to"
  default     = "development"
}

variable "resource_group_name" {
  description = "(Optional) Name of the resource group to place the database into. Optional and will create it's own if omitted"
  default     = ""
}

variable "db_admin_username" {
  description = "(Optional) Username for the database admin user"
  default     = "pgadmin"
}

variable "db_admin_password" {
  description = "(Required) Password for the database admin user"
  default     = "P@ssw0rd12345!"
}

variable "azure_postgres_sku_tier" {
  description = "(Required) Azure SKU tier reference for the DB (Preview - Basic and Standard available)"
  default     = "Standard"
}

variable "sku_compute_units" {
  description = "(Required) Azure compute units. 100 ~= 1 core. Default 100"
  default     = 100
}

variable "db_disk_size_mb" {
  description = "(Required) Size of the DB storage in MB - allowed values (basic) 51200, 179200, 307200. 435200 / (Standard) 128000, 256000, 384000 etc..."
  default     = 128000
}

variable "postgres_version" {
  description = "(Required) Version of postgres to use.  Currently 9.5 and 9.6 supported. Defaults to 9.6"
  default     = "9.6"
}

variable "enforce_ssl" {
  description = "(Optional) Should the server enforce SSL on connections to the database. Defaults to Enabled"
  default     = "Enabled"
}

variable "rule_name" {
  description = "Meaningful (and unique) name to describe what we're allowing"
  default     = "firewall_rule"
}

variable "start_ip_address" {
    description = "Start IP of the range to allow"
    default     = "0.0.0.0"
}

variable "end_ip_address" {
    description = "End IP of the range to allow"
    default     = "255.255.255.255"
}