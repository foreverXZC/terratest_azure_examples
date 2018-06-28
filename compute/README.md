# Compute Example

Using terraform azure module "compute" to deploy one or more virtual machines on azure.

## Compute

These terraform files enable users to deploy one or more virtual machines on azure, as well as virtual network. In order to use these files, you should provide path to ssh public key file in terraform.tfvars. You can just test the infrastructure code manually without terratest.

### Reference

[Terraform Azure Compute Module](https://registry.terraform.io/modules/Azure/compute/azurerm)
[Source Code](https://github.com/Azure/terraform-azurerm-compute)

## ssh

This folder includes two files. First, the go test file uses terraform compute module to deploy virtual machines on azure. After that, it tries to ssh to virtual machines and check whether they are running properly. Next, everthing will be cleaned up after validation. Of course you can write you own test code. Finally, in order to make this program work, you should provide ssh private key in id_rsa.

[Terratest SSH Source Code](https://github.com/gruntwork-io/terratest/blob/master/test/terraform_ssh_example_test.go)
[SSH Client Connection in Golang](http://blog.ralch.com/tutorial/golang-ssh-connection)