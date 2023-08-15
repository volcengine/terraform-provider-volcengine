package scaling_policy_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/autoscaling/scaling_policy"
	"testing"
)

const testAccVolcengineScalingPoliciesDatasourceConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block = "172.16.0.0/24"
  zone_id = "${data.volcengine_zones.foo.zones[0].id}"
  vpc_id = "${volcengine_vpc.foo.id}"
}

resource "volcengine_scaling_group" "foo" {
  scaling_group_name = "acc-test-scaling-group"
  subnet_ids = ["${volcengine_subnet.foo.id}"]
  multi_az_policy = "BALANCE"
  desire_instance_number = 0
  min_instance_number = 0
  max_instance_number = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown = 10
}

resource "volcengine_scaling_policy" "foo" {
  count = 3
  active = false
  scaling_group_id = "${volcengine_scaling_group.foo.id}"
  scaling_policy_name = "acc-tf-sg-policy-test-${count.index}"
  scaling_policy_type = "Alarm"
  adjustment_type = "QuantityChangeInCapacity"
  adjustment_value = 100
  cooldown = 10
  alarm_policy_rule_type = "Static"
  alarm_policy_evaluation_count = 1
  alarm_policy_condition_metric_name = "Instance_CpuBusy_Avg"
  alarm_policy_condition_metric_unit = "Percent"
  alarm_policy_condition_comparison_operator = "="
  alarm_policy_condition_threshold = 100
}

data "volcengine_scaling_policies" "foo"{
    scaling_group_id = "${volcengine_scaling_group.foo.id}"
	ids = volcengine_scaling_policy.foo[*].scaling_policy_id
}
`

func TestAccVolcengineScalingPoliciesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_scaling_policies.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &scaling_policy.VolcengineScalingPolicyService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineScalingPoliciesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "scaling_policies.#", "3"),
				),
			},
		},
	})
}
