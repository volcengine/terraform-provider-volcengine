package prefix_list_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/prefix_list"
)

const testAccVolcengineVpcPrefixListsDatasourceConfig = `
resource "volcengine_vpc_prefix_list" "foo" {
  prefix_list_name = "acc-test-prefix"
  max_entries = 3
  description = "acc test description"
  ip_version = "IPv4"
  prefix_list_entries {
    cidr = "192.168.4.0/28"
    description = "acc-test-1"
  }
  prefix_list_entries {
    cidr = "192.168.5.0/28"
    description = "acc-test-2"
  }
  tags {
    key = "tf-key1"
    value = "tf-value1"
  }
}

data "volcengine_vpc_prefix_lists" "foo" {
  ids = [volcengine_vpc_prefix_list.foo.id]
}
`

func TestAccVolcengineVpcPrefixListsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_vpc_prefix_lists.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return prefix_list.NewVpcPrefixListService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVpcPrefixListsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "prefix_lists.#", "1"),
				),
			},
		},
	})
}
