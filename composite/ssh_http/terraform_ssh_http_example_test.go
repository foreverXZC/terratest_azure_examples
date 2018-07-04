package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestTerraformSshExample(t *testing.T) {
	t.Parallel()

	exampleFolder := "../linuxserver"

	defer test_structure.RunTestStage(t, "teardown", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, exampleFolder)
		terraform.Destroy(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "setup", func() {
		terraformOptions := configureTerraformOptions(t, exampleFolder)

		test_structure.SaveTerraformOptions(t, exampleFolder, terraformOptions)

		terraform.InitAndApply(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "validate", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, exampleFolder)

		testSSHToPublicHost(t, terraformOptions)
		testHTTPToServer(t, terraformOptions)
	})

}

func configureTerraformOptions(t *testing.T, exampleFolder string) *terraform.Options {

	terraformOptions := &terraform.Options{
		TerraformDir: exampleFolder,

		Vars: map[string]interface{}{},
	}

	return terraformOptions
}

func testSSHToPublicHost(t *testing.T, terraformOptions *terraform.Options) {

	publicIP := terraform.Output(t, terraformOptions, "ip_address")

	maxRetries := 15
	timeBetweenRetries := 5 * time.Second
	description := fmt.Sprintf("SSH to public host %s", publicIP)

	sshConfig := createSSHConfig(t)
	target := createSSHTarget(t, publicIP)

	retry.DoWithRetry(t, description, maxRetries, timeBetweenRetries, func() (string, error) {
		session, err1 := createSSHSession(t, target, sshConfig)
		if err1 != nil {
			return "", err1
		}

		command1 := "sudo apt-get -y update"
		command2 := "sudo apt-get -y install nginx"
		commands := fmt.Sprintf("%s\n %s", command1, command2)

		err2 := runSSHCommands(t, commands, session)
		if err2 != nil {
			return "", err2
		}
		fmt.Println("Already installed nginx.")

		return "", nil
	})
}

func testHTTPToServer(t *testing.T, terraformOptions *terraform.Options) {
	publicIP := terraform.Output(t, terraformOptions, "ip_address")

	maxRetries := 15
	timeBetweenRetries := 5 * time.Second
	description1 := "Open port 80 to allow HTTP request."
	description2 := fmt.Sprintf("HTTP to %s", publicIP)

	retry.DoWithRetry(t, description1, maxRetries, timeBetweenRetries, func() (string, error) {
		err := openPort80(t)
		if err != nil {
			return "", err
		}
		fmt.Println("Already opened port 80.")

		return "", nil
	})

	retry.DoWithRetry(t, description2, maxRetries, timeBetweenRetries, func() (string, error) {
		response, err1 := getHTTPResponse(t, publicIP)
		if err1 != nil {
			return "", err1
		}

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
