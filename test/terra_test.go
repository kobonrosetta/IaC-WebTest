package test

import (
  "fmt"
  "testing"
  "time"

  http_helper "github.com/gruntwork-io/terratest/modules/http-helper"

  "github.com/gruntwork-io/terratest/modules/terraform"
)
//////////////


const regionOverrideEnvVarName = "us-east-2"

func TestAWSregion (t *testing.T) {
  t.Parallel()
  terraformOptions := &terraform.Options{
    TerraformDir: "../",
 }

 AWS_region = regionOverrideEnvVarName


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
