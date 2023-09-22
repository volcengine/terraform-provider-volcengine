package nas_file_system_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nas/nas_file_system"
)

const testAccVolcengineNasFileSystemsDatasourceConfig = `
data "volcengine_nas_zones" "foo" {

}

resource "volcengine_nas_file_system" "foo" {
    file_system_name = "acc-test-fs-${count.index}"
  	description = "acc-test"
  	zone_id = "${data.volcengine_nas_zones.foo.zones[0].id}"
  	capacity = 103
  	project_name = "default"
  	tags {
    	key = "k1"
    	value = "v1"
  	}
	count = 3
}

data "volcengine_nas_file_systems" "foo"{
    ids = volcengine_nas_file_system.foo[*].id
}
`

func TestAccVolcengineNasFileSystemsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_nas_file_systems.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return nas_file_system.NewNasFileSystemService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineNasFileSystemsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "file_systems.#", "3"),
				),
			},
		},
	})
}
