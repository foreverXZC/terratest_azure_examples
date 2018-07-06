# Composite Example

Use terraform azure module "compute" to deploy one virtual machine on azure. Then use terratest to ssh to it and also send http request.

## Linuxserver

Just as compute example does, these terraform files enable users to deploy one linux virtual machines on azure, as well as virtual network. To use these files, you should provide path to ssh public key file in [linuxserver/terraform.tfvars](/composite/linuxserver/terraform.tfvars). You can just test the infrastructure code manually without terratest.

## SSH_HTTP

This folder includes four files. Two of them are deprecated, but we are still using HTTP functions. Essentially, [ssh_http/terraform_ssh_http_example_test.go](/composite/ssh_http/terraform_ssh_http_example_test.go) is the main go test file which represents the whole process of testing the module. First, the go test file uses terraform compute module to deploy one linux virtual machine on azure. After that, it calls functions from terratest ssh section, so as to ssh to the virtual machine and execute shell commands to install nginx. Then the program opens port 80 and sends http request to the server. Next, everything will be cleaned up after validation. Of course you can write your own test code, or even take advantage of deprecated methods. Finally, in order to make this program work, you should provide your own ssh private key.

## Running this module manually

1. Sign up for [Azure](https://portal.azure.com/).

1. Configure your Azure credentials. For instance, you may use [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) and execute `az login`.

1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.

1. Fill in blank of your ssh public key in [linuxserver/terraform.tfvars](/composite/linuxserver/terraform.tfvars) and make sure your configuration is correct.

1. Direct to folder [linuxserver](/composite/linuxserver) and run `terraform init`.

1. Run `terraform apply`.

1. When you're done, run `terraform destroy`.

## Running automated tests against this module

1. Sign up for [Azure](https://portal.azure.com/).

1. Configure your Azure credentials. For instance, you may use [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) and execute `az login`.

1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.

1. Install [Golang](https://golang.org/) and make sure this code is checked out into your `GOPATH`.

1. Fill in blank of your ssh public key in [linuxserver/terraform.tfvars](/composite/linuxserver/terraform.tfvars) and make sure your configuration is correct.

1. Direct to folder [ssh_http](/composite/ssh_http) and make sure all packages are installed, such as executing `go get github.com/gruntwork-io/terratest/modules/terraform`, etc.

1. Run `go test -timeout timelimit -args username path/to/your/private/key`. For example, `go test -timeout 20m -args azureuser id_rsa`. Be aware that `-timeout` is set to 10 minutes by default and can be omitted, but it should be defined before `-args`.

## Reference

[Terraform Azure Compute Module](https://registry.terraform.io/modules/Azure/compute/azurerm/)

[Terratest SSH Source Code](https://github.com/gruntwork-io/terratest/blob/master/test/terraform_ssh_example_test.go)

[SSH Golang Document](https://godoc.org/golang.org/x/crypto/ssh)

[SSH Client Connection Example in Golang](http://blog.ralch.com/tutorial/golang-ssh-connection/)

[HTTP Golang Document](https://golang.org/pkg/net/http/)

[HTTP Request Example in Golang](https://gist.github.com/ijt/950790/fca88967337b9371bb6f7155f3304b3ccbf3946f)

[Shell Commands in Golang](https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/)

[Azure Linux Virtual Machine Document](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/)