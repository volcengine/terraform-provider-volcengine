package cluster_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/cluster"
	"testing"
)

const testAccVolcengineVkeClusterCreateConfig = `
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

resource "volcengine_security_group" "foo" {
    vpc_id = volcengine_vpc.foo.id
    security_group_name = "acc-test-security-group2"
}

resource "volcengine_vke_cluster" "foo" {
    name = "acc-test-1"
    description = "created by terraform"
    delete_protection_enabled = false
    cluster_config {
        subnet_ids = [volcengine_subnet.foo.id]
        api_server_public_access_enabled = true
        api_server_public_access_config {
            public_access_network_config {
                billing_type = "PostPaidByBandwidth"
                bandwidth = 1
            }
        }
        resource_public_access_default_enabled = true
    }
    pods_config {
        pod_network_mode = "VpcCniShared"
        vpc_cni_config {
            subnet_ids = [volcengine_subnet.foo.id]
        }
    }
    services_config {
        service_cidrsv4 = ["172.30.0.0/18"]
    }
    tags {
        key = "tf-k1"
        value = "tf-v1"
    }
}

`

const testAccVolcengineVkeClusterUpdateConfig = `
resource "volcengine_vpc" "foo" {
    vpc_name = "acc-test-project1"
    cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
    subnet_name = "acc-subnet-test-2"
    cidr_block = "172.16.0.0/24"
    zone_id = "cn-beijing-a"
    vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo" {
    vpc_id = volcengine_vpc.foo.id
    security_group_name = "acc-test-security-group2"
}

resource "volcengine_vke_cluster" "foo" {
    name = "acc-test-2"
    description = "created by terraform update"
    delete_protection_enabled = false
    cluster_config {
        subnet_ids = [volcengine_subnet.foo.id]
        api_server_public_access_enabled = false
        api_server_public_access_config {
            public_access_network_config {
                billing_type = "PostPaidByBandwidth"
                bandwidth = 2
            }
        }
        resource_public_access_default_enabled = true
    }
    pods_config {
        pod_network_mode = "VpcCniShared"
        vpc_cni_config {
            subnet_ids = [volcengine_subnet.foo.id]
        }
    }
    services_config {
        service_cidrsv4 = ["172.30.0.0/18"]
    }
    tags {
        key = "tf-k1"
        value = "tf-v1"
    }
    tags {
        key = "tf-k2"
        value = "tf-v2"
    }
}

`

func TestAccVolcengineVkeClusterResource_Basic(t *testing.T) {
	resourceName := "volcengine_vke_cluster.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &cluster.VolcengineVkeClusterService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVkeClusterCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "cluster_config.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "cluster_config.0.api_server_public_access_enabled", "true"),
					resource.TestCheckResourceAttr(acc.ResourceId, "delete_protection_enabled", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "created by terraform"),
					resource.TestCheckResourceAttr(acc.ResourceId, "logging_config.#", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "pods_config.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "pods_config.0.pod_network_mode", "VpcCniShared"),
					resource.TestCheckResourceAttr(acc.ResourceId, "services_config.#", "1"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "services_config.0.service_cidrsv4.*", "172.30.0.0/18"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "tf-k1",
						"value": "tf-v1",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "cluster_config.0.api_server_public_access_config.0.public_access_network_config.0.bandwidth", "1"),
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

func TestAccVolcengineVkeClusterResource_Update(t *testing.T) {
	resourceName := "volcengine_vke_cluster.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &cluster.VolcengineVkeClusterService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVkeClusterCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "cluster_config.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "cluster_config.0.api_server_public_access_enabled", "true"),
					resource.TestCheckResourceAttr(acc.ResourceId, "delete_protection_enabled", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "created by terraform"),
					resource.TestCheckResourceAttr(acc.ResourceId, "logging_config.#", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "pods_config.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "pods_config.0.pod_network_mode", "VpcCniShared"),
					resource.TestCheckResourceAttr(acc.ResourceId, "services_config.#", "1"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "services_config.0.service_cidrsv4.*", "172.30.0.0/18"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "tf-k1",
						"value": "tf-v1",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "cluster_config.0.api_server_public_access_config.0.public_access_network_config.0.bandwidth", "1"),
				),
			},
			{
				Config: testAccVolcengineVkeClusterUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "cluster_config.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "cluster_config.0.api_server_public_access_enabled", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "delete_protection_enabled", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "created by terraform update"),
					resource.TestCheckResourceAttr(acc.ResourceId, "logging_config.#", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "pods_config.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "pods_config.0.pod_network_mode", "VpcCniShared"),
					resource.TestCheckResourceAttr(acc.ResourceId, "services_config.#", "1"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "services_config.0.service_cidrsv4.*", "172.30.0.0/18"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "2"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "tf-k1",
						"value": "tf-v1",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "tf-k2",
						"value": "tf-v2",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "cluster_config.0.api_server_public_access_config.0.public_access_network_config.0.bandwidth", "0"),
				),
			},
			{
				Config:             testAccVolcengineVkeClusterUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
