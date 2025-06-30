package vmp_rule_file_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vmp/vmp_rule_file"
)

const testAccVolcengineVmpRuleFilesDatasourceConfig = `
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

data "volcengine_vmp_rule_files" "foo"{
    ids = [volcengine_vmp_rule_file.foo.rule_file_id]
	workspace_id = volcengine_vmp_workspace.foo.id
}
`

func TestAccVolcengineVmpRuleFilesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_vmp_rule_files.foo"

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
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVmpRuleFilesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "files.#", "1"),
				),
			},
		},
	})
}
