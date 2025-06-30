package vmp_rule_file_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vmp/vmp_rule_file"
)

const testAccVolcengineVmpRuleFileCreateConfig = `
resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-1"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-1"
  username                  = "admin123"
  password                  = "admin1239A82"
}

resource "volcengine_vmp_rule_file" "foo" {
  name         = "acc-test-1"
  workspace_id = volcengine_vmp_workspace.foo.id
  description  = "acc-test-1"
  content      = <<EOF
groups:
    - interval: 10s
      name: recording_rules
      rules:
        - expr: sum(irate(container_cpu_usage_seconds_total{image!=""}[5m])) by (pod) *100
          labels:
            team: operations
          record: pod:cpu:useage
EOF
}

`

const testAccVolcengineVmpRuleFileUpdateConfig = `
resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-1"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-1"
  username                  = "admin123"
  password                  = "admin1239A82"
}

resource "volcengine_vmp_rule_file" "foo" {
  name         = "acc-test-2"
  workspace_id = volcengine_vmp_workspace.foo.id
  description  = "acc-test-2"
  content      = <<EOF
groups:
    - interval: 15s
      name: recording_rules
      rules:
        - expr: sum(irate(container_cpu_usage_seconds_total{image!=""}[5m])) by (pod) *100
          labels:
            team: operations
          record: pod:cpu:useage
EOF
}
`

func TestAccVolcengineVmpRuleFileResource_Basic(t *testing.T) {
	resourceName := "volcengine_vmp_rule_file.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return vmp_rule_file.NewService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVmpRuleFileCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVolcengineVmpRuleFileResource_Update(t *testing.T) {
	resourceName := "volcengine_vmp_rule_file.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return vmp_rule_file.NewService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVmpRuleFileCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-1"),
				),
			},
			{
				Config: testAccVolcengineVmpRuleFileUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-2"),
				),
			},
			{
				Config:             testAccVolcengineVmpRuleFileUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
