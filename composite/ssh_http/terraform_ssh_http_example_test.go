package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/ssh"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestTerraformSshExample(t *testing.T) {
	t.Parallel()

	exampleFolder := "../linuxserver"

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer test_structure.RunTestStage(t, "teardown", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, exampleFolder)
		terraform.Destroy(t, terraformOptions)
	})

	// Deploy the example
	test_structure.RunTestStage(t, "setup", func() {
		terraformOptions := configureTerraformOptions(t, exampleFolder)

		// Save the options so later test stages can use them
		test_structure.SaveTerraformOptions(t, exampleFolder, terraformOptions)

		// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
		terraform.InitAndApply(t, terraformOptions)
	})

	// Make sure we can SSH to the virtual machine and send HTTP request
	test_structure.RunTestStage(t, "validate", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, exampleFolder)

		// Make sure we can SSH to virtual machines directly from the public Internet
		testSSHToPublicHost(t, terraformOptions)

		// Make sure we can send HTTP request to the server
		testHTTPToServer(t, terraformOptions)
	})

}

func configureTerraformOptions(t *testing.T, exampleFolder string) *terraform.Options {

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: exampleFolder,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{},
	}

	return terraformOptions
}

func testSSHToPublicHost(t *testing.T, terraformOptions *terraform.Options) {
	// Run `terraform output` to get the value of an output variable
	publicIP := terraform.Output(t, terraformOptions, "ip_address")

	// Read private key from given file
	buffer, err := ioutil.ReadFile(os.Args[len(os.Args)-1])
	if err != nil {
		t.Fatal(err)
	}
	keyPair := ssh.KeyPair{PrivateKey: string(buffer)}

	// We're going to try to SSH to the virtual machine, using our local key pair and specific username
	publicHost := ssh.Host{
		Hostname:    publicIP,
		SshKeyPair:  &keyPair,
		SshUserName: os.Args[len(os.Args)-2],
	}

	// It can take a minute or so for the virtual machine to boot up, so retry a few times
	maxRetries := 15
	timeBetweenRetries := 5 * time.Second
	description := fmt.Sprintf("SSH to public host %s", publicIP)

	// Run command that installs nginx on the server
	command1 := "sudo apt-get -y update"
	command2 := "sudo apt-get -y install nginx"
	command := fmt.Sprintf("%s\n %s", command1, command2)

	// Verify that we can SSH to the virtual machine and run commands
	retry.DoWithRetry(t, description, maxRetries, timeBetweenRetries, func() (string, error) {
		// Run the command to install nginx
		_, err := ssh.CheckSshCommandE(t, publicHost, command)
		if err != nil {
			return "", err
		}
		fmt.Println("Already installed nginx.")

		return "", nil
	})
}

func testHTTPToServer(t *testing.T, terraformOptions *terraform.Options) {
	// Run `terraform output` to get the value of an output variable
	publicIP := terraform.Output(t, terraformOptions, "ip_address")

	// It can take a minute or so for the web server to boot up, so retry a few times
	maxRetries := 15
	timeBetweenRetries := 5 * time.Second
	description1 := "Open port 80 to allow HTTP request."
	description2 := fmt.Sprintf("HTTP to %s", publicIP)

	// Try several times to open port 80 to allow HTTP request
	retry.DoWithRetry(t, description1, maxRetries, timeBetweenRetries, func() (string, error) {
		err := openPort80(t)
		if err != nil {
			return "", err
		}
		fmt.Println("Already opened port 80.")

		return "", nil
	})

	// Verify that we can send HTTP request
	retry.DoWithRetry(t, description2, maxRetries, timeBetweenRetries, func() (string, error) {
		// Get HTTP response from server
		response, err1 := getHTTPResponse(t, publicIP)
		if err1 != nil {
			return "", err1
		}

		// Check whether the content of HTTP response contains nginx
		defer response.Body.Close()
		substring := "nginx"
		err2 := checkContents(t, response, substring)
		if err2 != nil {
			return "", err2
		}
		fmt.Printf("HTTP found %s.\n", substring)

		return "", nil
	})
}
