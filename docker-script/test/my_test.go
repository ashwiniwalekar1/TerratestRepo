package test

import (
        "testing"

        "github.com/gruntwork-io/terratest/modules/docker"
        "github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
        buildOptions := &docker.BuildOptions{
                Tags:       []string{"node-app"},
        }
        docker.Build(t, "../", buildOptions)

        runOptions := &docker.RunOptions{
         OtherOptions: []string{"-p", "8000:8000"},
        }

        // Run the Docker container and get the output
		output := docker.Run(t, "node-app:latest", runOptions)

        // Assert the expected outcome
        expectedOutput := "Data existed in mongoDB"
        assert.Equal(t, output, expectedOutput)
}