# Terratest Azure Examples

The following examples use terraform to deploy things on azure and use terratest to check whether the infrastructure works properly. For more information, please read reference and see details in each example. Any issues and pull request encouraged.

## Available Examples

### Compute

Use terraform azure module "compute" to deploy one or more virtual machines on azure. Then use terratest to ssh to virtual machines.

### Database

Use terraform azure module "database" to deploy a Microsoft SQL Database on azure. Then use terratest to connect to the database.

### Webapp

Use resources from terraform azure module "webapp" to deploy a web application on azure. Then use terratest to send http request.

### Composite

Use terraform azure module "compute" to deploy one virtual machine on azure. Then use terratest to ssh to it and also send http request.

## Reference

### Azure

[Azure Portal](https://portal.azure.com/)

[Azure Cloud Shell](https://shell.azure.com/)

### Terraform

[Terraform Index Page](https://www.terraform.io/)

[Terraform Modules for Azure](https://registry.terraform.io/browse?provider=azurerm)

[Terraform Tutorial on Azure](https://docs.microsoft.com/en-us/azure/terraform/)

### Terratest

[Terratest Source Code & Document](https://github.com/gruntwork-io/terratest/)
