package test

import (
	"fmt"
	"os"
	"log"
	"time"
	"testing"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformWebserverExample(t *testing.T){
	t.Parallel()
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: "../",

	})

	terraform.InitAndApply(t, terraformOptions)

	defer terraform.Destroy(t, terraformOptions)

	publicIp := terraform.Output(t, terraformOptions, "public_ip")

	url := fmt.Sprintf("http://%s", publicIp)

	http_helper.HttpGetWithRetry(t, url, nil, 200, "I Made a Terraform Module", 30, 5*time.Second)

	t.Logf("Making HTTP request to %s is success", url)

	// Write the response body to a file
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer f.Close()

}