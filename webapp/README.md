# Webapp Example

Use resources from terraform azure module "webapp" to deploy a web application on azure. Then use terratest to send http request.

## Webapp

Because webapp module is not verified and has some problems currently, we use corresponding resources to deploy a simple web application on azure. You can just test the infrastructure code manually without terratest.

## HTTP

The test code here is relatively simple. First, it uses specific resources to deploy a template web application on azure. After that, it gets the web app URL and tries to make http request to see whether the application runs properly. Finally, everything will be cleaned up after validation. This example only sends a request and does not do anything else. For a more complex one, see composite example.

## Reference

[Terraform Azure Webapp Module](https://registry.terraform.io/modules/rahulkhengare/webapp/azurerm/)

[HTTP Golang Document](https://golang.org/pkg/net/http/)

[HTTP Request Example in Golang](https://gist.github.com/ijt/950790/fca88967337b9371bb6f7155f3304b3ccbf3946f)

[Azure Webapp Deployment Using Terraform](https://docs.microsoft.com/en-us/azure/terraform/terraform-slot-walkthru)