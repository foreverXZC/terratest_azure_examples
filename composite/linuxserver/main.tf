provider "random" {
  version = "~> 1.0"
}

resource "random_id" "ip_dns" {
  byte_length = 8
}

module "linuxservers" {
  source                       = "Azure/compute/azurerm"
  location                     = "${var.location}"
  vm_os_simple                 = "${var.vm_os_simple}"
  public_ip_dns                = ["linsimplevmips-${random_id.ip_dns.hex}"]
  vnet_subnet_id               = "${module.network.vnet_subnets[0]}"
  admin_username               = "${var.admin_username}"
  admin_password               = "${var.admin_password}"
  ssh_key                      = "${var.ssh_key}"
  resource_group_name          = "${var.resource_group_name}"
  public_ip_address_allocation = "static"
}

module "network" {
  version             = "2.0.0"
  source              = "Azure/network/azurerm"
  location            = "${var.location}"
  resource_group_name = "${var.resource_group_name}"
}