package instance_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/escloud/instance"
)

const testAccVolcengineEscloudInstanceCreateConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet_new"
  description = "tfdesc"
  cidr_block = "172.16.0.0/24"
  zone_id = "${data.volcengine_zones.foo.zones[0].id}"
  vpc_id = "${volcengine_vpc.foo.id}"
}

resource "volcengine_escloud_instance" "foo" {
  instance_configuration {
    version            = "V6_7"
    zone_number        = 1
    enable_https       = true
    admin_user_name    = "admin"
    admin_password     = "Password@@"
    charge_type        = "PostPaid"
    configuration_code = "es.standard"
    enable_pure_master = true
    instance_name      = "acc-test-0"
    node_specs_assigns {
      type               = "Master"
      number             = 3
      resource_spec_name = "es.x4.medium"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 100
    }
    node_specs_assigns {
      type               = "Hot"
      number             = 2
      resource_spec_name = "es.x4.large"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 100
    }
    node_specs_assigns {
      type               = "Kibana"
      number             = 1
      resource_spec_name = "kibana.x2.small"
    }
    subnet_id = volcengine_subnet.foo.id
    project_name = "default"
    force_restart_after_scale = false
  }
}
`

const testAccVolcengineEscloudInstanceUpdateConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet_new"
  description = "tfdesc"
  cidr_block = "172.16.0.0/24"
  zone_id = "${data.volcengine_zones.foo.zones[0].id}"
  vpc_id = "${volcengine_vpc.foo.id}"
}

resource "volcengine_escloud_instance" "foo" {
  instance_configuration {
    version            = "V6_7"
    zone_number        = 1
    enable_https       = true
    admin_user_name    = "admin"
    admin_password     = "Password@@acc"
    charge_type        = "PostPaid"
    configuration_code = "es.standard"
    enable_pure_master = true
    instance_name      = "acc-test-1"
    node_specs_assigns {
      type               = "Master"
      number             = 3
      resource_spec_name = "es.x4.medium"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 100
    }
	node_specs_assigns {
      type               = "Hot"
      number             = 4
      resource_spec_name = "es.x4.large"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 120
    }
    node_specs_assigns {
      type               = "Kibana"
      number             = 1
      resource_spec_name = "kibana.x2.small"
    }
    subnet_id = volcengine_subnet.foo.id
    project_name = "default"
    force_restart_after_scale = false
  }
}
`

func TestAccVolcengineEscloudInstanceResource_Basic(t *testing.T) {
	resourceName := "volcengine_escloud_instance.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return instance.NewESCloudInstanceService(client)
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
				Config: testAccVolcengineEscloudInstanceCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.#", "3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.version", "V6_7"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.zone_number", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.enable_https", "true"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.admin_user_name", "admin"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.admin_password", "Password@@"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.charge_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.configuration_code", "es.standard"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.enable_pure_master", "true"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.instance_name", "acc-test-0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.type", "Master"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.number", "3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.resource_spec_name", "es.x4.medium"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.storage_spec_name", "es.volume.essd.pl0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.storage_size", "100"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.type", "Hot"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.number", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.resource_spec_name", "es.x4.large"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.storage_spec_name", "es.volume.essd.pl0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.storage_size", "100"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.2.type", "Kibana"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.2.number", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.2.resource_spec_name", "kibana.x2.small"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance_configuration.0.admin_password"},
			},
		},
	})
}

func TestAccVolcengineEscloudInstanceResource_Update(t *testing.T) {
	resourceName := "volcengine_escloud_instance.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return instance.NewESCloudInstanceService(client)
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
				Config: testAccVolcengineEscloudInstanceCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.#", "3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.version", "V6_7"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.zone_number", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.enable_https", "true"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.admin_user_name", "admin"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.admin_password", "Password@@"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.charge_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.configuration_code", "es.standard"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.enable_pure_master", "true"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.instance_name", "acc-test-0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.type", "Master"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.number", "3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.resource_spec_name", "es.x4.medium"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.storage_spec_name", "es.volume.essd.pl0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.storage_size", "100"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.type", "Hot"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.number", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.resource_spec_name", "es.x4.large"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.storage_spec_name", "es.volume.essd.pl0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.storage_size", "100"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.2.type", "Kibana"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.2.number", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.2.resource_spec_name", "kibana.x2.small"),
				),
			},
			{
				Config: testAccVolcengineEscloudInstanceUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.#", "3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.version", "V6_7"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.zone_number", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.enable_https", "true"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.admin_user_name", "admin"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.admin_password", "Password@@acc"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.charge_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.configuration_code", "es.standard"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.enable_pure_master", "true"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.instance_name", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.type", "Master"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.number", "3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.resource_spec_name", "es.x4.medium"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.storage_spec_name", "es.volume.essd.pl0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.0.storage_size", "100"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.type", "Hot"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.number", "4"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.resource_spec_name", "es.x4.large"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.storage_spec_name", "es.volume.essd.pl0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.1.storage_size", "120"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.2.type", "Kibana"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.2.number", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_configuration.0.node_specs_assigns.2.resource_spec_name", "kibana.x2.small"),
				),
			},
			{
				Config:             testAccVolcengineEscloudInstanceUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
