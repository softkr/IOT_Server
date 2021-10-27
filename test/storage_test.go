package test

import (
	"fmt"
	"iot/grpc/client"
	"testing"
)

func TestProject(t *testing.T) {
	result := client.GetProject("21IHPA0000A")
	fmt.Println(result.Project)
}
