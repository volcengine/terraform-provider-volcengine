package vmp_notify_policy_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vmp/vmp_notify_policy"
)

const testAccVolcengineVmpNotifyPoliciesDatasourceConfig = `
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

data "volcengine_vmp_notify_policies" "foo"{
    ids = [volcengine_vmp_notify_policy.foo.id]
}
`

func TestAccVolcengineVmpNotifyPoliciesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_vmp_notify_policies.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return vmp_notify_policy.NewService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVmpNotifyPoliciesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "notify_policies.#", "1"),
				),
			},
		},
	})
}
