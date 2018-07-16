# Configure the Azure provider
provider "azurerm" {}

resource "azurerm_resource_group" "webapp" {
  name     = "${var.resource_group_name}"
  location = "${var.location}"
}

resource "azurerm_app_service_plan" "webapp" {
  name                = "${var.service_plan_name == "" ? replace(var.name, "/[^a-z0-9]/", "") : var.service_plan_name}"
  location            = "${azurerm_resource_group.webapp.location}"
  resource_group_name = "${azurerm_resource_group.webapp.name}"

  sku {
    tier = "Standard"
    size = "S1"
  }
}

resource "azurerm_app_service" "webapp" {
  name                = "${var.name}"
  location            = "${azurerm_resource_group.webapp.location}"
  resource_group_name = "${azurerm_resource_group.webapp.name}"
  app_service_plan_id = "${azurerm_app_service_plan.webapp.id}"
}
