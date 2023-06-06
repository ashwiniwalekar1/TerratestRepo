package test

import (
	"database/sql"
	"testing"
	"fmt"

	"github.com/gruntwork-io/terratest/modules/aws"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestTerraformWebserverExample(t *testing.T){
	t.Parallel()
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: "../",

	})

	expectedPort := int64(3306)
	expectedDBname := "mydatabase"
	username := "admin"
	password := "Admin!1994"
	awsRegion := "ap-south-1"
	// engineVersion := "8.0.32"
	// instanceType := "db.t2.micro"

	terraform.InitAndApply(t, terraformOptions)

	

	// Run `terraform output` to get the value of an output variable
	dbInstanceID := terraform.Output(t, terraformOptions, "db_instance_id")

	// Look up the endpoint address and port of the RDS instance
	address := aws.GetAddressOfRdsInstance(t, dbInstanceID, awsRegion)
	t.Logf("%s",address)
	port := aws.GetPortOfRdsInstance(t, dbInstanceID, awsRegion)
	t.Logf("%d",port)
	schemaExistsInRdsInstance := aws.GetWhetherSchemaExistsInRdsMySqlInstance(t, address, port, username, password, expectedDBname)
	t.Logf("%t",schemaExistsInRdsInstance)
	// Verify that the address is not null
	assert.NotNil(t, address)
	// Verify that the DB instance is listening on the port mentioned
	assert.Equal(t, expectedPort, port)
	// Verify that the table/schema requested for creation is actually present in the database
	assert.True(t, schemaExistsInRdsInstance)

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, address, port, expectedDBname)
	t.Logf("database URL : %s", dbURL)

    // Use the appropriate database driver for your codebase (e.g., MySQL)
    db, err := sql.Open("mysql", dbURL)
    if err != nil {
        t.Fatalf("Failed to open database connection: %s", err)
    }
    defer db.Close()

    // Ping the database to test the connection
    err = db.Ping()
    if err != nil {
		t.Errorf("Failed to connect to the database: %s", err)
	} else {
		t.Logf("Successfully connected to the database")
	}

	defer terraform.Destroy(t, terraformOptions)

}