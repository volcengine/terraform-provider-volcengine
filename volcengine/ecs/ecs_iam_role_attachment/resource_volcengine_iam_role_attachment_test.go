package ecs_iam_role_attachment_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_iam_role_attachment"
)

const testAccVolcengineIamRoleAttachmentCreateConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block = "172.16.0.0/24"
  zone_id = data.volcengine_zones.foo.zones[0].id
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-security-group"
  vpc_id = volcengine_vpc.foo.id
}

data "volcengine_images" "foo" {
  os_type = "Linux"
  visibility = "public"
  instance_type_id = "ecs.g1ie.large"
}

resource "volcengine_ecs_instance" "foo" {
  instance_name = "acc-test-ecs"
  description = "acc-test"
  host_name = "tf-acc-test"
  image_id = data.volcengine_images.foo.images[0].image_id
  instance_type = "ecs.g1ie.large"
  password = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type = "ESSD_PL0"
  system_volume_size = 40
  data_volumes {
    volume_type = "ESSD_PL0"
    size = 50
    delete_with_instance = true
  }
  subnet_id = volcengine_subnet.foo.id
  security_group_ids = [volcengine_security_group.foo.id]
  project_name = "default"
  tags {
    key = "k1"
    value = "v1"
  }
}

resource "volcengine_ecs_instance" "foo1" {
  instance_name = "acc-test-ecs-1"
  description = "acc-test"
  host_name = "tf-acc-test"
  image_id = data.volcengine_images.foo.images[0].image_id
  instance_type = "ecs.g1ie.large"
  password = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type = "ESSD_PL0"
  system_volume_size = 40
  data_volumes {
    volume_type = "ESSD_PL0"
    size = 50
    delete_with_instance = true
  }
  subnet_id = volcengine_subnet.foo.id
  security_group_ids = [volcengine_security_group.foo.id]
  project_name = "default"
  tags {
    key = "k1"
    value = "v1"
  }
}

resource "volcengine_iam_role" "foo" {
  role_name = "acc-test-role"
  display_name = "acc-test"
  description = "acc-test"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"ecs\"]}}]}"
  max_session_duration = 36000
}

resource "volcengine_iam_role_attachment" "foo" {
  iam_role_name = volcengine_iam_role.foo.id
  instance_id = volcengine_ecs_instance.foo.id
}
`

func TestAccVolcengineIamRoleAttachmentResource_Basic(t *testing.T) {
	resourceName := "volcengine_iam_role_attachment.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return ecs_iam_role_attachment.NewIamRoleAttachmentService(client)
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
				Config: testAccVolcengineIamRoleAttachmentCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "iam_role_name", "acc-test-role"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "instance_id"),
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
