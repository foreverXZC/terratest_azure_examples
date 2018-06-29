# Configure the Azure provider
provider "azurerm" { }

resource "azurerm_resource_group" "slotDemo" {
    name = "${var.resource_group_name}"
    location = "${var.location}"
}

resource "azurerm_app_service_plan" "slotDemo" {
    name                = "${var.service_plan_name == "" ? replace(var.name, "/[^a-z0-9]/", "") : var.service_plan_name}"
    location            = "${azurerm_resource_group.slotDemo.location}"
    resource_group_name = "${azurerm_resource_group.slotDemo.name}"
    sku {
        tier = "Standard"
        size = "S1"
    }
}

resource "azurerm_app_service" "slotDemo" {
    name                = "${var.name}"
    location            = "${azurerm_resource_group.slotDemo.location}"
    resource_group_name = "${azurerm_resource_group.slotDemo.name}"
    app_service_plan_id = "${azurerm_app_service_plan.slotDemo.id}"
}
