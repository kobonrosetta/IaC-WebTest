package test

import (
  "fmt"
  "testing"
  "time"

  http_helper "github.com/gruntwork-io/terratest/modules/http-helper"

  "github.com/gruntwork-io/terratest/modules/terraform"

  "github.com/stretchr/testify/assert"
)
//////////////

func GetRandomRegion(t *Testing.T, approvedRegions []string, forbiddenRegions []string) string {
	region, err := GetRandomRegionE(t, approvedRegions, forbiddenRegions)
	if err != nil {
		t.Fatal(err)
	}
	return region
}

func TestGetRandomRegionExcludesForbiddenRegions(t *testing.T) {
	t.Parallel()

	approvedRegions := []string{"us-east-2"}
	forbiddenRegions := []string{"us-west-2", "ap-northeast-2", "ca-central-1", "us-east-1", "us-west-1", "us-west-2", "eu-west-1", "eu-west-2", "eu-central-1", "ap-southeast-1", "ap-northeast-1", "ap-northeast-2", "ap-south-1"}

	for i := 0; i < 1000; i++ {
		randomRegion := GetRandomRegion(t, approvedRegions, forbiddenRegions)
		assert.NotContains(t, forbiddenRegions, randomRegion)
	}
}


///////////

func TestTerraformAwsHelloWorldExample(t *testing.T) {
  t.Parallel()

  terraformOptions := &terraform.Options{
    TerraformDir: "../",
  }


  defer terraform.Destroy(t, terraformOptions)

  terraform.InitAndApply(t, terraformOptions)

  publicIp := terraform.Output(t, terraformOptions, "public_ip")

  url := fmt.Sprintf("http://%s:8080", publicIp)
  http_helper.HttpGetWithRetry(t, url, nil, 200, "Hello, World!", 30, 5*time.Second)
}
