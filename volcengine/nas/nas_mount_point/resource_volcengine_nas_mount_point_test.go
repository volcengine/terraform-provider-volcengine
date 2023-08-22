package nas_mount_point_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nas/nas_mount_point"
)

const testAccVolcengineNasMountPointCreateConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
    vpc_name = "acc-test-project1"
    cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
    subnet_name = "acc-subnet-test-2"
    cidr_block = "172.16.0.0/24"
    zone_id = data.volcengine_zones.foo.zones[0].id
    vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_nas_permission_group" "foo" {
  permission_group_name = "acc-test"
  description = "acctest"
  permission_rules {
    cidr_ip = "*"
    rw_mode = "RW"
    use_mode = "All_squash"
  }
  permission_rules {
    cidr_ip = "192.168.0.0"
    rw_mode = "RO"
    use_mode = "All_squash"
  }
}

data "volcengine_nas_zones" "foo" {

}

resource "volcengine_nas_file_system" "foo" {
    file_system_name = "acc-test-fs"
  	description = "acc-test"
  	zone_id = "${data.volcengine_nas_zones.foo.zones[0].id}"
  	capacity = 103
  	project_name = "default"
  	tags {
    	key = "k1"
    	value = "v1"
  	}
}

resource "volcengine_nas_mount_point" "foo" {
    file_system_id = volcengine_nas_file_system.foo.id
    mount_point_name = "acc-test"
    permission_group_id = volcengine_nas_permission_group.foo.id
    subnet_id = volcengine_subnet.foo.id
}
`

const testAccVolcengineNasMountPointUpdateConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
    vpc_name = "acc-test-project1"
    cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
    subnet_name = "acc-subnet-test-2"
    cidr_block = "172.16.0.0/24"
    zone_id = data.volcengine_zones.foo.zones[0].id
    vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_nas_permission_group" "foo" {
  permission_group_name = "acc-test"
  description = "acctest"
  permission_rules {
    cidr_ip = "*"
    rw_mode = "RW"
    use_mode = "All_squash"
  }
  permission_rules {
    cidr_ip = "192.168.0.0"
    rw_mode = "RO"
    use_mode = "All_squash"
  }
}

data "volcengine_nas_zones" "foo" {

}

resource "volcengine_nas_file_system" "foo" {
    file_system_name = "acc-test-fs"
  	description = "acc-test"
  	zone_id = "${data.volcengine_nas_zones.foo.zones[0].id}"
  	capacity = 103
  	project_name = "default"
  	tags {
    	key = "k1"
    	value = "v1"
  	}
}

resource "volcengine_nas_mount_point" "foo" {
    file_system_id = volcengine_nas_file_system.foo.id
    mount_point_name = "acc-test1"
    permission_group_id = volcengine_nas_permission_group.foo.id
    subnet_id = volcengine_subnet.foo.id
}
`

func TestAccVolcengineNasMountPointResource_Basic(t *testing.T) {
	resourceName := "volcengine_nas_mount_point.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return nas_mount_point.NewVolcengineNasMountPointService(client)
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
				Config: testAccVolcengineNasMountPointCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "mount_point_name", "acc-test"),
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

func TestAccVolcengineNasMountPointResource_Update(t *testing.T) {
	resourceName := "volcengine_nas_mount_point.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return nas_mount_point.NewVolcengineNasMountPointService(client)
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
				Config: testAccVolcengineNasMountPointCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "mount_point_name", "acc-test"),
				),
			},
			{
				Config: testAccVolcengineNasMountPointUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "mount_point_name", "acc-test1"),
				),
			},
			{
				Config:             testAccVolcengineNasMountPointUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
