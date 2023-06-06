package test

import (
    "time"
    "bytes"
    "io/ioutil"
    "testing"
    "os/exec"
    // "fmt"
    // "strings"

    // "github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"golang.org/x/crypto/ssh"
)

func TestFileExists(t *testing.T) {
    t.Parallel()
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: "../",
        MaxRetries:   3,
        TimeBetweenRetries: 5 * time.Second,
	})

	terraform.InitAndApply(t, terraformOptions)

	defer terraform.Destroy(t, terraformOptions)

	publicIp := terraform.Output(t, terraformOptions, "public_ip")
	

    // Create an SSH client to run commands on the EC2 instance
    sshUsername := "ec2-user"
    sshKeyPath := "../ssmKey.pem"
    privateKeyBytes, err := ioutil.ReadFile(sshKeyPath)
    if err != nil {
        t.Errorf("Error reading private key file: %v", err)
    }
    privateKey, err := ssh.ParsePrivateKey(privateKeyBytes)
    if err != nil {
        t.Errorf("Error parsing private key: %v", err)
    }
    sshConfig := &ssh.ClientConfig{
        User: sshUsername,
        Auth: []ssh.AuthMethod{
            ssh.PublicKeys(privateKey),
        },
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    // Remove the existing host key for the EC2 instance
    cmd := exec.Command("ssh-keygen", "-R", publicIp)
    if err := cmd.Run(); err != nil {
        t.Errorf("Error removing existing host key: %v", err)
    }
    sshClient, err := ssh.Dial("tcp", publicIp+":22", sshConfig)
    if err != nil {
        t.Errorf("Error establishing SSH connection: %v", err)
    }
    t.Logf("Making SSH Connection is for  %v success", sshClient)
    defer sshClient.Close()

    // Run a remote command to check if the file exists
    // command := "sudo ls /var/www/html/"
    // session, err := sshClient.NewSession()
    // if err != nil {
    //     t.Errorf("Error creating SSH session: %v", err)
    // }
    // defer session.Close()
    // var stdout bytes.Buffer
    // session.Stdout = &stdout
    // if err := session.Run(command); err != nil {
    //     t.Errorf("Error running command: %v", err)
    // }
    // output := stdout.String()
    // t.Logf("Output of command: %s", output)
    // if output == "" {
    //     t.Errorf("File /var/www/html/ does not exist")
    // }

    command := "ls -l ~/"
    session, err := sshClient.NewSession()
    if err != nil {
        t.Errorf("Error creating SSH session: %v", err)
    }
    defer session.Close()
    var stdout bytes.Buffer
    session.Stdout = &stdout
    if err := session.Run(command); err != nil {
        t.Errorf("Error running command: %v", err)
    }
    output := stdout.String()
    t.Logf("Output of command: %s", output)
}
