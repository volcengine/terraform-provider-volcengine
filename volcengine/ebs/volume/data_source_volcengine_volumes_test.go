package volume_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/volume"
	"testing"
)

const testAccVolcengineVolumesDatasourceConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_volume" "foo" {
	volume_name = "acc-test-volume-${count.index}"
    volume_type = "ESSD_PL0"
	description = "acc-test"
    kind = "data"
    size = 60
    zone_id = "${data.volcengine_zones.foo.zones[0].id}"
	volume_charge_type = "PostPaid"
	project_name = "default"
	count = 3
}

data "volcengine_volumes" "foo"{
    ids = volcengine_volume.foo[*].id
}
`

func TestAccVolcengineVolumesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_volumes.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &volume.VolcengineVolumeService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVolumesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "volumes.#", "3"),
				),
			},
		},
	})
}
