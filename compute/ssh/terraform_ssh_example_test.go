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

	exampleFolder := "../compute"

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

		testSSHToPublicHost(t, terraformOptions, "ubuntu_ip_address")
		testSSHToPublicHost(t, terraformOptions, "debian_ip_address")
	})

}

func configureTerraformOptions(t *testing.T, exampleFolder string) *terraform.Options {

	terraformOptions := &terraform.Options{
		TerraformDir: exampleFolder,

		Vars: map[string]interface{}{},
	}

	return terraformOptions
}

func testSSHToPublicHost(t *testing.T, terraformOptions *terraform.Options, address string) {

	publicIP := terraform.Output(t, terraformOptions, address)

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

		text := "Hello, World!"
		commands := fmt.Sprintf("echo -n '%s'", text)

		err2 := runSSHCommands(t, commands, session)
		if err2 != nil {
			return "", err2
		}
		fmt.Println(text)

		return "", nil
	})
}
