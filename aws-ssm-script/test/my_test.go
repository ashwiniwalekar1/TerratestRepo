package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	// "github.com/stretchr/testify/require"
)

func TestTerraformAwsSsmExample(t *testing.T) {
	t.Parallel()
	region := "ap-south-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		
	})
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	instanceID := terraform.Output(t, terraformOptions, "instance_id")
	timeout := 5 * time.Minute

	aws.WaitForSsmInstance(t, region, instanceID, timeout)

	result := aws.CheckSsmCommand(t, region, instanceID, "echo Hello, World", timeout)
	// require.Equal(t, result.Stdout, "Hello, World\n")
	// require.Equal(t, result.Stderr, "")
	// require.Equal(t, int64(0), result.ExitCode)
	t.Logf("the output of command : %s", result.Stdout)

	result, err := aws.CheckSsmCommandE(t, region, instanceID, "cat /wrong/file", timeout)
	// require.Error(t, err)
	// require.Equal(t, "Failed", err.Error())
	// require.Equal(t, "cat: /wrong/file: No such file or directory\nfailed to run commands: exit status 1", result.Stderr)
	// require.Equal(t, "", result.Stdout)
	// require.Equal(t, int64(1), result.ExitCode)
	t.Logf("the output of command : %s and %s", result.Stderr, err.Error())
}