# Composite Example

Use terraform azure module "compute" to deploy one virtual machine on azure. Then use terratest to ssh to it and also send http request.

## Linuxserver

Just as compute example does, these terraform files enable users to deploy one linux virtual machines on azure, as well as virtual network. To use these files, you should provide path to ssh public key file in terraform.tfvars. You can just test the infrastructure code manually without terratest.

## SSH_HTTP

This folder includes two files. First, the go test file uses terraform compute module to deploy one linux virtual machine on azure. After that, it tries to ssh to it and execute shell commands to install nginx. Then the program opens port 80 and sends HTTP request to the server. Next, everything will be cleaned up after validation. Be aware that some part of code is hardcoded corresponding to terraform file, and of course you can write your own test code. Finally, in order to make this program work, you should provide ssh private key in id_rsa.

## Reference

[Terraform Azure Compute Module](https://registry.terraform.io/modules/Azure/compute/azurerm/)

[Terratest SSH Source Code](https://github.com/gruntwork-io/terratest/blob/master/test/terraform_ssh_example_test.go)

[SSH Golang Document](https://godoc.org/golang.org/x/crypto/ssh)

[SSH Client Connection Example in Golang](http://blog.ralch.com/tutorial/golang-ssh-connection/)

[HTTP Golang Document](https://golang.org/pkg/net/http/)

[HTTP Request Example in Golang](https://gist.github.com/ijt/950790/fca88967337b9371bb6f7155f3304b3ccbf3946f)

[Shell Commands in Golang](https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/)

[Azure Linux Virtual Machine Document](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/)