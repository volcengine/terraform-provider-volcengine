package vmp_contact_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vmp/vmp_contact"
)

const testAccVolcengineVmpContactCreateConfig = `
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

`

const testAccVolcengineVmpContactUpdateConfig = `
resource "volcengine_vmp_contact" "foo" {
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

`

func TestAccVolcengineVmpContactResource_Basic(t *testing.T) {
	resourceName := "volcengine_vmp_contact.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return vmp_contact.NewService(client)
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
				Config: testAccVolcengineVmpContactCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "ding_talk_bot_webhook.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "ding_talk_bot_webhook.0.address", "https://www.dingacctest1.com"),
					resource.TestCheckResourceAttr(acc.ResourceId, "email", "acctest1@tftest.com"),
					resource.TestCheckResourceAttr(acc.ResourceId, "email_active", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lark_bot_webhook.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lark_bot_webhook.0.address", "https://www.acctest1.com"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-contact"),
					resource.TestCheckResourceAttr(acc.ResourceId, "phone_number.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "phone_number.0.country_code", "+86"),
					resource.TestCheckResourceAttr(acc.ResourceId, "phone_number.0.number", "18310101010"),
					resource.TestCheckResourceAttr(acc.ResourceId, "webhook.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "webhook.0.address", "https://www.acctest1.com"),
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

func TestAccVolcengineVmpContactResource_Update(t *testing.T) {
	resourceName := "volcengine_vmp_contact.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return vmp_contact.NewService(client)
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
				Config: testAccVolcengineVmpContactCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "ding_talk_bot_webhook.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "ding_talk_bot_webhook.0.address", "https://www.dingacctest1.com"),
					resource.TestCheckResourceAttr(acc.ResourceId, "email", "acctest1@tftest.com"),
					resource.TestCheckResourceAttr(acc.ResourceId, "email_active", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lark_bot_webhook.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lark_bot_webhook.0.address", "https://www.acctest1.com"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-contact"),
					resource.TestCheckResourceAttr(acc.ResourceId, "phone_number.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "phone_number.0.country_code", "+86"),
					resource.TestCheckResourceAttr(acc.ResourceId, "phone_number.0.number", "18310101010"),
					resource.TestCheckResourceAttr(acc.ResourceId, "webhook.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "webhook.0.address", "https://www.acctest1.com"),
				),
			},
			{
				Config: testAccVolcengineVmpContactUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "ding_talk_bot_webhook.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "ding_talk_bot_webhook.0.address", "https://www.dingacctest2.com"),
					resource.TestCheckResourceAttr(acc.ResourceId, "email", "acctest2@tftest.com"),
					resource.TestCheckResourceAttr(acc.ResourceId, "email_active", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lark_bot_webhook.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lark_bot_webhook.0.address", "https://www.acctest2.com"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-contact2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "phone_number.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "phone_number.0.country_code", "+86"),
					resource.TestCheckResourceAttr(acc.ResourceId, "phone_number.0.number", "18310101011"),
					resource.TestCheckResourceAttr(acc.ResourceId, "webhook.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "webhook.0.address", "https://www.acctest2.com"),
				),
			},
			{
				Config:             testAccVolcengineVmpContactUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
