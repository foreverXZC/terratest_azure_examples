package test

import (
	"fmt"
	"io/ioutil"
	"net"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/test-structure"
	"golang.org/x/crypto/ssh"
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

	sshConfig := &ssh.ClientConfig{
		User: "azureuser",
		Auth: []ssh.AuthMethod{
			PublicKeyFile("id_rsa"),
		},
		HostKeyCallback: func(string, net.Addr, ssh.PublicKey) error {
			return nil
		},
	}

	host := publicIP
	port := "22"
	target := fmt.Sprintf("%s:%s", host, port)

	retry.DoWithRetry(t, description, maxRetries, timeBetweenRetries, func() (string, error) {
		connection, err := ssh.Dial("tcp", target, sshConfig)
		if err != nil {
			return "", err
		}

		session, err2 := connection.NewSession()
		if err2 != nil {
			return "", err2
		}

		modes := ssh.TerminalModes{
			ssh.ECHO:          0,     // disable echoing
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		}

		err3 := session.RequestPty("xterm", 80, 40, modes)
		if err3 != nil {
			session.Close()
			return "", err3
		}

		text := "Hello, World!"
		command := fmt.Sprintf("echo -n '%s'", text)

		err4 := session.Run(command)
		if err4 != nil {
			return "", err4
		}
		fmt.Println(text)

		return "", nil
	})
}

//PublicKeyFile ssh
func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}
