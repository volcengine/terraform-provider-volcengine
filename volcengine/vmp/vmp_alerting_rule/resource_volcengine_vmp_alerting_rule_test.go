package vmp_alerting_rule_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vmp/vmp_alerting_rule"
)

const testAccVolcengineVmpAlertingRuleCreateConfig = `
resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-1"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-1"
  username                  = "admin123"
  password                  = "admin1239A82"
}

resource "volcengine_vmp_contact" "foo" {
  name  = "acc-test-contact"
  email = "acctest1@tftest.com"

  webhook {
    address = "https://www.acctest1.com"
  }

  lark_bot_webhook {
    address = "https://www.acctest1.com"
  }

  ding_talk_bot_webhook {
    address    = "https://www.dingacctest1.com"
    at_mobiles = ["18046891812"]
  }
  phone_number {
    country_code = "+86"
    number = "18310101010"
  }
}

resource "volcengine_vmp_contact" "foo1" {
  name  = "acc-test-contact2"
  email = "acctest2@tftest.com"

  webhook {
    address = "https://www.acctest2.com"
  }

  lark_bot_webhook {
    address = "https://www.acctest2.com"
  }

  ding_talk_bot_webhook {
    address    = "https://www.dingacctest2.com"
    at_mobiles = ["18046891813"]
  }
  phone_number {
    country_code = "+86"
    number = "18310101011"
  }
}

resource "volcengine_vmp_contact_group" "foo" {
  name = "acc-test"
  contact_ids = [volcengine_vmp_contact.foo.id]
}

resource "volcengine_vmp_contact_group" "foo1" {
  name = "acc-test-1"
  contact_ids = [volcengine_vmp_contact.foo1.id]
}

resource "volcengine_vmp_notify_policy" "foo" {
  name = "acc-test-1"
  description = "acc-test-1"
  levels {
    level = "P1"
    contact_group_ids = [volcengine_vmp_contact_group.foo.id]
    channels = ["Email", "Webhook"]
  }
  levels {
    level = "P0"
    contact_group_ids = [volcengine_vmp_contact_group.foo1.id]
    channels = ["LarkBotWebhook"]
  }
}

resource "volcengine_vmp_notify_group_policy" "foo" {
  name = "acc-test-1"
  description = "acc-test-1"
  levels {
    level = "P2"
    group_by = ["__rule__"]
    group_wait = "35"
    group_interval = "30"
    repeat_interval = "30"
  }
  levels {
    level = "P0"
    group_by = ["__rule__"]
    group_wait = "30"
    group_interval = "30"
    repeat_interval = "30"
  }
  levels {
    level = "P1"
    group_by = ["__rule__"]
    group_wait = "40"
    group_interval = "45"
    repeat_interval = "30"
  }
}

resource "volcengine_vmp_alerting_rule" "foo" {
  name = "acc-test-1"
  description = "acc-test-1"
  notify_policy_id = volcengine_vmp_notify_policy.foo.id
  notify_group_policy_id = volcengine_vmp_notify_group_policy.foo.id
  query {
    workspace_id = volcengine_vmp_workspace.foo.id
    prom_ql = "sum(up)"
  }
  levels {
    level = "P0"
    for = "0s"
    comparator = ">="
    threshold = 2.0
  }
  levels {
    level = "P1"
    for = "0s"
    comparator = ">="
    threshold = 1.0
  }
  levels {
    level = "P2"
    for = "0s"
    comparator = ">="
    threshold = 0.5
  }
}

`

const testAccVolcengineVmpAlertingRuleUpdateConfig = `
resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-1"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-1"
  username                  = "admin123"
  password                  = "admin1239A82"
}

resource "volcengine_vmp_contact" "foo" {
  name  = "acc-test-contact"
  email = "acctest1@tftest.com"

  webhook {
    address = "https://www.acctest1.com"
  }

  lark_bot_webhook {
    address = "https://www.acctest1.com"
  }

  ding_talk_bot_webhook {
    address    = "https://www.dingacctest1.com"
    at_mobiles = ["18046891812"]
  }
  phone_number {
    country_code = "+86"
    number = "18310101010"
  }
}

resource "volcengine_vmp_contact" "foo1" {
  name  = "acc-test-contact2"
  email = "acctest2@tftest.com"

  webhook {
    address = "https://www.acctest2.com"
  }

  lark_bot_webhook {
    address = "https://www.acctest2.com"
  }

  ding_talk_bot_webhook {
    address    = "https://www.dingacctest2.com"
    at_mobiles = ["18046891813"]
  }
  phone_number {
    country_code = "+86"
    number = "18310101011"
  }
}

resource "volcengine_vmp_contact_group" "foo" {
  name = "acc-test"
  contact_ids = [volcengine_vmp_contact.foo.id]
}

resource "volcengine_vmp_contact_group" "foo1" {
  name = "acc-test-1"
  contact_ids = [volcengine_vmp_contact.foo1.id]
}

resource "volcengine_vmp_notify_policy" "foo" {
  name = "acc-test-1"
  description = "acc-test-1"
  levels {
    level = "P1"
    contact_group_ids = [volcengine_vmp_contact_group.foo.id]
    channels = ["Email", "Webhook"]
  }
  levels {
    level = "P0"
    contact_group_ids = [volcengine_vmp_contact_group.foo1.id]
    channels = ["LarkBotWebhook"]
  }
}

resource "volcengine_vmp_notify_group_policy" "foo" {
  name = "acc-test-1"
  description = "acc-test-1"
  levels {
    level = "P2"
    group_by = ["__rule__"]
    group_wait = "35"
    group_interval = "30"
    repeat_interval = "30"
  }
  levels {
    level = "P0"
    group_by = ["__rule__"]
    group_wait = "30"
    group_interval = "30"
    repeat_interval = "30"
  }
  levels {
    level = "P1"
    group_by = ["__rule__"]
    group_wait = "40"
    group_interval = "45"
    repeat_interval = "30"
  }
}

resource "volcengine_vmp_alerting_rule" "foo" {
  name = "acc-test-2"
  description = "acc-test-2"
  notify_policy_id = volcengine_vmp_notify_policy.foo.id
  notify_group_policy_id = volcengine_vmp_notify_group_policy.foo.id
  query {
    workspace_id = volcengine_vmp_workspace.foo.id
    prom_ql = "count(up)"
  }
  levels {
    level = "P1"
    for = "0s"
    comparator = ">="
    threshold = 2.0
  }
  levels {
    level = "P2"
    for = "0s"
    comparator = ">="
    threshold = 1.0
  }
  levels {
    level = "P0"
    for = "0s"
    comparator = ">="
    threshold = 0.5
  }
}

`

func TestAccVolcengineVmpAlertingRuleResource_Basic(t *testing.T) {
	resourceName := "volcengine_vmp_alerting_rule.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return vmp_alerting_rule.NewVmpAlertingRuleService(client)
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
				Config: testAccVolcengineVmpAlertingRuleCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "levels.#", "3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "query.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "query.0.prom_ql", "sum(up)"),
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

func TestAccVolcengineVmpAlertingRuleResource_Update(t *testing.T) {
	resourceName := "volcengine_vmp_alerting_rule.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return vmp_alerting_rule.NewVmpAlertingRuleService(client)
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
				Config: testAccVolcengineVmpAlertingRuleCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "levels.#", "3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "query.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "query.0.prom_ql", "sum(up)"),
				),
			},
			{
				Config: testAccVolcengineVmpAlertingRuleUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "levels.#", "3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "query.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "query.0.prom_ql", "count(up)"),
				),
			},
			{
				Config:             testAccVolcengineVmpAlertingRuleUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
