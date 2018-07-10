module "sql-database" {
    source              = "github.com/Azure/terraform-azurerm-database"
    resource_group_name = "${var.resource_group_name}"
    location            = "${var.location}"
    db_name             = "${var.db_name}"
    sql_admin_username  = "${var.sql_admin_username}"
    sql_password        = "${var.sql_password}"
    start_ip_address    = "${var.start_ip_address}"
    end_ip_address      = "${var.end_ip_address}"
}