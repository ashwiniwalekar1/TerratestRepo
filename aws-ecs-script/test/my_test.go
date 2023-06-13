package test

import (
	"fmt"
	"net/http"
	"time"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	// "github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"

	awsSDK "github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsEcsExample(t *testing.T) {
	t.Parallel()
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})
	awsRegion := "ap-south-1"
	expectedClusterName := "my-ecs-cluster"
	expectedServiceName := "my-service"
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	dns_name := terraform.Output(t, terraformOptions, "dns_name")

	url := fmt.Sprintf("http://%s", dns_name)

	// Send an HTTP GET request
	maxRetries := 20
	retryInterval := 5 * time.Second

	for retry := 1; retry <= maxRetries; retry++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Attempt %d: HTTP access failed with error: %v\n", retry, err)
			time.Sleep(retryInterval)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Println("HTTP access successful")
			break
		} else {
			fmt.Printf("Attempt %d: HTTP access failed with status code: %d\n", retry, resp.StatusCode)
			time.Sleep(retryInterval)
			continue
		}
	}

	taskDefinition := terraform.Output(t, terraformOptions, "task_definition")

	// Look up the ECS cluster by name
	cluster := aws.GetEcsCluster(t, awsRegion, expectedClusterName)

	assert.Equal(t, int64(1), awsSDK.Int64Value(cluster.ActiveServicesCount))

	// Look up the ECS service by name
	service := aws.GetEcsService(t, awsRegion, expectedClusterName, expectedServiceName)

	assert.Equal(t, int64(1), awsSDK.Int64Value(service.DesiredCount))
	assert.Equal(t, "FARGATE", awsSDK.StringValue(service.LaunchType))

	// Look up the ECS task definition by ARN
	task := aws.GetEcsTaskDefinition(t, awsRegion, taskDefinition)

	assert.Equal(t, "256", awsSDK.StringValue(task.Cpu))
	assert.Equal(t, "512", awsSDK.StringValue(task.Memory))
	assert.Equal(t, "awsvpc", awsSDK.StringValue(task.NetworkMode))
}