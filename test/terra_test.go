package test

import (
  "fmt"
  "testing"
  "time"

  http_helper "github.com/gruntwork-io/terratest/modules/http-helper"

  "github.com/gruntwork-io/terratest/modules/terraform"
  "github.com/gruntwork-io/terratest/modules/aws"
  "github.com/stretchr/testify/assert"
	"github.com/gruntwork-io/terratest/modules/random"
)

func TestTerraformAwsHelloWorldExample(t *testing.T) {
  t.Parallel()
  approvedRegions := []string{"us-east-2"}
  awsRegion := aws.GetRandomRegion(t, approvedRegions, nil)
  expectedName := fmt.Sprintf("terratest-%s", random.UniqueId())
  terraformOptions := &terraform.Options{
    TerraformDir: "../",

         Vars: map[string]interface{}{
          "instance_name": expectedName,
          "test_label": "yes",
          "region":awsRegion,
       },
       EnvVars: map[string]string{
        "AWS_DEFAULT_REGION": awsRegion,
     },

  }


  defer terraform.Destroy(t, terraformOptions)

  terraform.InitAndApply(t, terraformOptions)

  actualInstanceId := []string{terraform.Output(t, terraformOptions, "instance_id")}
  tagName := "Name"
  exptectedInstanceId := aws.GetEc2InstanceIdsByTag(t, awsRegion, tagName, expectedName)
  assert.Equal(t, exptectedInstanceId, actualInstanceId)

  publicIp := terraform.Output(t, terraformOptions, "public_ip")
  url := fmt.Sprintf("http://%s:8080", publicIp)
  http_helper.HttpGetWithRetry(t, url, nil, 200, "website is online", 30, 5*time.Second)
}
