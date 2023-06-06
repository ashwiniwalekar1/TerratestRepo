
package test

import (
	"fmt"
	"testing"
	
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformWebserverExample(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	bucketName := terraform.Output(t, terraformOptions, "bucket_name")
	region := terraform.Output(t, terraformOptions, "region")

	// awsSession, err := session.NewSession()
	// if err != nil {
	// 	t.Fatalf("Failed to create AWS session: %v", err)
	// }
	awsSession, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
		},
	})

	s3Client := s3.New(awsSession)

	objectName := "index.html"
	objectExists := false

	// objectNames := []string{"index.html", "buildspec.yml"}
	// objectsFound := make(map[string]bool)

	// for _, objectName := range objectNames {
	// 	err = s3Client.ListObjectsV2Pages(&s3.ListObjectsV2Input{
	// 		Bucket: &bucketName,
	// 		Prefix: aws.String(objectName),
			
	err = s3Client.ListObjectsV2Pages(&s3.ListObjectsV2Input{
		Bucket: &bucketName,
			 
	}, func(page *s3.ListObjectsV2Output, lastPage bool) bool {
		for _, object := range page.Contents {
			if *object.Key == objectName {
				objectExists = true
				// objectsFound[objectName] = true
				break
			}
		}
		return !lastPage
	})
	if objectExists {
		t.Logf("Object %s exists in S3 bucket %s", objectName, bucketName)
	}
	// for _, objectName := range objectNames {
	// 	if objectsFound[objectName] {
	// 		t.Logf("Object %s exists in S3 bucket %s", objectName, bucketName)
	// 	} 
	// }


	if err != nil {
		t.Errorf("Failed to list objects in S3 bucket %s: %v", bucketName, err)
	}

	if !objectExists {
		t.Errorf("Object %s does not exist in S3 bucket %s", objectName, bucketName)
	}

// 	for _, objectName := range objectNames {
// 		if !objectsFound[objectName] {
// 			t.Errorf("Object %s does not exist in S3 bucket %s", objectName, bucketName)
// 		}
// 	}

   // Specify the S3 bucket and object you want to get the contents of
   objectKey := "index.html"

   // Get the object contents
   objectOutput, err := s3Client.GetObject(&s3.GetObjectInput{
	   Bucket: &bucketName,
	   Key:    &objectKey,
   })
   assert.NoError(t, err)

   // Read the contents of the object
   buffer := make([]byte, int(aws.Int64Value(objectOutput.ContentLength)))
   _, err = objectOutput.Body.Read(buffer)
   assert.NoError(t, err)

   // Assert that the object contents are not empty
   assert.NotEmpty(t, buffer)

   fmt.Printf("Object contents:\n%s\n", string(buffer))
 }




// package test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/s3"
// 	"github.com/gruntwork-io/terratest/modules/terraform"
// )

// func TestTerraformWebserverExample(t *testing.T) {
// 	t.Parallel()

// 	terraformOptions := &terraform.Options{
// 		TerraformDir: "../",
// 	}

// 	defer terraform.Destroy(t, terraformOptions)

// 	terraform.InitAndApply(t, terraformOptions)

// 	bucketName := terraform.Output(t, terraformOptions, "bucket_name")
// 	// region := terraform.Output(t, terraformOptions, "region")

// 	awsSession, err := session.NewSession()
// 	if err != nil {
// 		t.Fatalf("Failed to create AWS session: %v", err)
// 	}

// 	s3Client := s3.New(awsSession)

// 	for _, obj := range []string{"index.html"} {
// 		key := fmt.Sprintf("%s", obj)

// 		_, err := s3Client.GetObject(&s3.GetObjectInput{
// 			Bucket: &bucketName,
// 			Key:    &key,
// 		})

// 		if err != nil {
// 			t.Errorf("Object %s does not exist in S3 bucket %s", key,  bucketName)
// 		}
// 	}
// }