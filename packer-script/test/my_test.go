package test

import (
	"testing"

	// "github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/packer"
	// terratest_aws "github.com/gruntwork-io/terratest/modules/aws"
	// "github.com/stretchr/testify/assert"
)


func TestPackerHelloWorldExample(t *testing.T) {
	packerOptions := &packer.Options{
		Template: "../rhel-ami.pkr.hcl",

	}
	// awsRegion := "ap-south-1"

	amiID := packer.BuildArtifact(t, packerOptions)

	// Clean up the AMI after we're done
	// defer terratest_aws.DeleteAmiAndAllSnapshots(t, awsRegion, amiID)

}