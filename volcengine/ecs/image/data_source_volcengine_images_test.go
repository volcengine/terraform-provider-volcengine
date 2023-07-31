package image_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/image"
	"testing"
)

const testAccVolcengineImagesDatasourceConfig = `
data "volcengine_images" "foo" {
	  os_type = "Linux"
	  visibility = "public"
	  instance_type_id = "ecs.g1.large"
}
`

func TestAccVolcengineImagesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_images.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &image.VolcengineImageService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineImagesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "images.#", "26"),
				),
			},
		},
	})
}
