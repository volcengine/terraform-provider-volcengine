package cluster_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/bioos/cluster"
)

const testAccVolcengineBioosClusterCreateConfig = `
resource "volcengine_bioos_cluster" "foo" {
    name = "acc-test-cluster"
  	description = "acc-test-description"
  	//vke_config {
	//	cluster_id = "cckrr899c5ami65led1hg"
	//	storage_class = "ebs-ssd"
	//}
	shared_config {
		enable = true
	}
}
`

func TestAccVolcengineBioosClusterResource_Basic(t *testing.T) {
	resourceName := "volcengine_bioos_cluster.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return cluster.NewVolcengineBioosClusterService(client)
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
				Config: testAccVolcengineBioosClusterCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-description"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-cluster"),
					resource.TestCheckResourceAttr(acc.ResourceId, "shared_config.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "vke_config.#", "0"),
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
