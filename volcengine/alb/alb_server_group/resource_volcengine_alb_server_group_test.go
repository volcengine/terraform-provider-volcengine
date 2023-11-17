package alb_server_group_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_server_group"
)

const testAccVolcengineAlbServerGroupCreateConfig = `
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_alb_server_group" "foo" {
  vpc_id = volcengine_vpc.foo.id
  server_group_name = "acc-test-server-group"
  description = "acc-test"
  server_group_type = "instance"
  scheduler = "wlc"
  project_name = "default"
}
`

func TestAccVolcengineAlbServerGroupResource_Basic(t *testing.T) {
	resourceName := "volcengine_alb_server_group.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return alb_server_group.NewAlbServerGroupService(client)
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
				Config: testAccVolcengineAlbServerGroupCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "server_group_name", "acc-test-server-group"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "server_group_type", "instance"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "scheduler", "wlc"),
					resource.TestCheckResourceAttr(acc.ResourceId, "server_count", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Active"),
					resource.TestCheckResourceAttr(acc.ResourceId, "listeners.#", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "health_check.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "health_check.*", map[string]string{
						"enabled":             "on",
						"healthy_threshold":   "3",
						"http_code":           "http_2xx,http_3xx",
						"http_version":        "HTTP1.0",
						"interval":            "2",
						"method":              "HEAD",
						"timeout":             "2",
						"unhealthy_threshold": "3",
						"uri":                 "/",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "sticky_session_config.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "sticky_session_config.*", map[string]string{
						"sticky_session_enabled": "off",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "create_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "update_time"),
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

const testAccVolcengineAlbServerGroupUpdateConfig = `
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_alb_server_group" "foo" {
  vpc_id = volcengine_vpc.foo.id
  server_group_name = "acc-test-server-group-new"
  description = "acc-test-new"
  server_group_type = "instance"
  scheduler = "sh"
  project_name = "default"
  health_check {
    enabled = "on"
    interval = 3
    timeout = 3
    method = "GET"
  }
  sticky_session_config {
    sticky_session_enabled = "on"
    sticky_session_type = "insert"
    cookie_timeout = "1100"
  }
}
`

func TestAccVolcengineAlbServerGroupResource_Update(t *testing.T) {
	resourceName := "volcengine_alb_server_group.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return alb_server_group.NewAlbServerGroupService(client)
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
				Config: testAccVolcengineAlbServerGroupCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "server_group_name", "acc-test-server-group"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "server_group_type", "instance"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "scheduler", "wlc"),
					resource.TestCheckResourceAttr(acc.ResourceId, "server_count", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Active"),
					resource.TestCheckResourceAttr(acc.ResourceId, "listeners.#", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "health_check.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "health_check.*", map[string]string{
						"enabled":             "on",
						"healthy_threshold":   "3",
						"http_code":           "http_2xx,http_3xx",
						"http_version":        "HTTP1.0",
						"interval":            "2",
						"method":              "HEAD",
						"timeout":             "2",
						"unhealthy_threshold": "3",
						"uri":                 "/",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "sticky_session_config.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "sticky_session_config.*", map[string]string{
						"sticky_session_enabled": "off",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "create_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "update_time"),
				),
			},
			{
				Config: testAccVolcengineAlbServerGroupUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "server_group_name", "acc-test-server-group-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "server_group_type", "instance"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "scheduler", "sh"),
					resource.TestCheckResourceAttr(acc.ResourceId, "server_count", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Active"),
					resource.TestCheckResourceAttr(acc.ResourceId, "listeners.#", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "health_check.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "health_check.*", map[string]string{
						"enabled":             "on",
						"healthy_threshold":   "3",
						"http_code":           "http_2xx,http_3xx",
						"http_version":        "HTTP1.0",
						"interval":            "3",
						"method":              "GET",
						"timeout":             "3",
						"unhealthy_threshold": "3",
						"uri":                 "/",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "sticky_session_config.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "sticky_session_config.*", map[string]string{
						"sticky_session_enabled": "on",
						"sticky_session_type":    "insert",
						"cookie_timeout":         "1100",
						"cookie":                 "",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "create_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "update_time"),
				),
			},
			{
				Config:             testAccVolcengineAlbServerGroupUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
