# Compute Example

Use terraform azure module "compute" to deploy one virtual machine on azure. Then use terratest to ssh to it and also send http request.

## Compute

These terraform files enable users to deploy one or more virtual machines on azure, as well as virtual network. To use these files, you should provide path to ssh public key file in terraform.tfvars. You can just test the infrastructure code manually without terratest.

## SSH

This folder includes three files. Most importantly, 'terraform_ssh_example_test.go' is the main go test file which represents the whole process of testing the module. First, it uses terraform compute module to deploy virtual machines on azure. After that, it calls functions from other two files, so as to ssh to these virtual machines and check whether they are running properly. Next, everything will be cleaned up after validation. Of course you can write your own test code. Finally, in order to make this program work, you should provide your own ssh private key.

## Reference

[Terraform Azure Compute Module](https://registry.terraform.io/modules/Azure/compute/azurerm/)

[Terratest SSH Source Code](https://github.com/gruntwork-io/terratest/blob/master/test/terraform_ssh_example_test.go)

[SSH Golang Document](https://godoc.org/golang.org/x/crypto/ssh)

[SSH Client Connection Example in Golang](http://blog.ralch.com/tutorial/golang-ssh-connection/)

[Azure Linux Virtual Machine Document](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/)

[Azure Virtual Network Document](https://docs.microsoft.com/en-us/azure/virtual-network/)