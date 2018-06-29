package test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformHttpExample(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../webapp",

		Vars: map[string]interface{}{},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	webappURL := terraform.Output(t, terraformOptions, "webapp_url")

	maxRetries := 15
	timeBetweenRetries := 5 * time.Second
	description := fmt.Sprintf("HTTP to %s", webappURL)

	retry.DoWithRetry(t, description, maxRetries, timeBetweenRetries, func() (string, error) {
		response, err := http.Get(webappURL)
		if err != nil {
			return "", err
		}

		defer response.Body.Close()
		fmt.Println("HTTP succeed")

		return "", nil
	})
}
