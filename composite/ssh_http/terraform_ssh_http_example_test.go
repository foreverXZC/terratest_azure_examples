package test

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ssh"
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

		command1 := "sudo apt-get -y update"
		command2 := "sudo apt-get -y install nginx"
		command := fmt.Sprintf("%s\n %s", command1, command2)

		err4 := session.Run(command)
		if err4 != nil {
			return "", err4
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
		args := []string{"vm", "open-port", "--port", "80", "--resource-group", "linuxResourceGroup", "--name", "myvm0"}
		cmd := exec.Command("az", args...)
		err := cmd.Run()
		if err != nil {
			return "", err
		}
		fmt.Println("Already opened port 80.")

		return "", nil
	})

	retry.DoWithRetry(t, description2, maxRetries, timeBetweenRetries, func() (string, error) {
		target := "http://" + publicIP
		response, err2 := http.Get(target)
		if err2 != nil {
			return "", err2
		}

		defer response.Body.Close()
		contents, err3 := ioutil.ReadAll(response.Body)
		if err3 != nil {
			return "", err3
		}
		document := string(contents)
		nginx := strings.Contains(document, "nginx")
		assert.Equal(t, true, nginx)
		fmt.Println("HTTP found nginx.")

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
