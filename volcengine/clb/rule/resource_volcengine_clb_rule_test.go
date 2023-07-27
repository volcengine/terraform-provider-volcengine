package rule_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/rule"
	"testing"
)

const testAccVolcengineClbRuleCreateConfig = `
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

resource "volcengine_clb" "foo" {
  type = "public"
  subnet_id = "${volcengine_subnet.foo.id}"
  load_balancer_spec = "small_1"
  description = "acc0Demo"
  load_balancer_name = "acc-test-create"
  eip_billing_config {
    isp = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth = 1
  }
}

resource "volcengine_server_group" "foo" {
  load_balancer_id = "${volcengine_clb.foo.id}"
  server_group_name = "acc-test-create"
  description = "hello demo11"
}

resource "volcengine_listener" "foo" {
  load_balancer_id = "${volcengine_clb.foo.id}"
  listener_name = "acc-test-listener"
  protocol = "HTTP"
  port = 90
  server_group_id = "${volcengine_server_group.foo.id}"
  health_check {
    enabled = "on"
    interval = 10
    timeout = 3
    healthy_threshold = 5
    un_healthy_threshold = 2
    domain = "volcengine.com"
    http_code = "http_2xx"
    method = "GET"
    uri = "/"
  }
  enabled = "on"
}
resource "volcengine_clb_rule" "foo" {
  listener_id = "${volcengine_listener.foo.id}"
  server_group_id = "${volcengine_server_group.foo.id}"
  domain = "test-volc123.com"
  url = "/yyyy"
}
`

const testAccVolcengineClbRuleUpdateConfig = `
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

resource "volcengine_clb" "foo" {
  type = "public"
  subnet_id = "${volcengine_subnet.foo.id}"
  load_balancer_spec = "small_1"
  description = "acc0Demo"
  load_balancer_name = "acc-test-create"
  eip_billing_config {
    isp = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth = 1
  }
}

resource "volcengine_server_group" "foo" {
  load_balancer_id = "${volcengine_clb.foo.id}"
  server_group_name = "acc-test-create"
  description = "hello demo11"
}

resource "volcengine_listener" "foo" {
  load_balancer_id = "${volcengine_clb.foo.id}"
  listener_name = "acc-test-listener"
  protocol = "HTTP"
  port = 90
  server_group_id = "${volcengine_server_group.foo.id}"
  health_check {
    enabled = "on"
    interval = 10
    timeout = 3
    healthy_threshold = 5
    un_healthy_threshold = 2
    domain = "volcengine.com"
    http_code = "http_2xx"
    method = "GET"
    uri = "/"
  }
  enabled = "on"
}
resource "volcengine_clb_rule" "foo" {
  listener_id = "${volcengine_listener.foo.id}"
  server_group_id = "${volcengine_server_group.foo.id}"
  domain = "acc-test-volc123.com"
  url = "/accyyyy"
}
`

func TestAccVolcengineClbRuleResource_Basic(t *testing.T) {
	resourceName := "volcengine_clb_rule.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &rule.VolcengineRuleService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineClbRuleCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "domain", "test-volc123.com"),
					resource.TestCheckResourceAttr(acc.ResourceId, "url", "/yyyy"),
				),
			},
		},
	})
}

func TestAccVolcengineClbRuleResource_Update(t *testing.T) {
	resourceName := "volcengine_clb_rule.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &rule.VolcengineRuleService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineClbRuleCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "domain", "test-volc123.com"),
					resource.TestCheckResourceAttr(acc.ResourceId, "url", "/yyyy"),
				),
			},
			{
				Config: testAccVolcengineClbRuleUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "domain", "acc-test-volc123.com"),
					resource.TestCheckResourceAttr(acc.ResourceId, "url", "/accyyyy"),
				),
			},
			{
				Config:             testAccVolcengineClbRuleUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
