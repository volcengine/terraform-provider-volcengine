package node_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/node"
)

const testAccVolcengineVkeNodeCreateConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
	vpc_name   = "acc-test-vpc"
  	cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  	subnet_name = "acc-test-subnet"
  	cidr_block = "172.16.0.0/24"
  	zone_id = "${data.volcengine_zones.foo.zones[0].id}"
	vpc_id = "${volcengine_vpc.foo.id}"
}

resource "volcengine_security_group" "foo" {
  	security_group_name = "acc-test-security-group"
  	vpc_id = "${volcengine_vpc.foo.id}"
}

data "volcengine_images" "foo" {
  name_regex = "veLinux 1.0 CentOS兼容版 64位"
}

resource "volcengine_vke_cluster" "foo" {
    name = "acc-test-cluster"
    description = "created by terraform"
    delete_protection_enabled = false
    cluster_config {
        subnet_ids = ["${volcengine_subnet.foo.id}"]
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
            subnet_ids = ["${volcengine_subnet.foo.id}"]
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

resource "volcengine_vke_node_pool" "foo" {
	cluster_id = "${volcengine_vke_cluster.foo.id}"
	name = "acc-test-node-pool"
	auto_scaling {
        enabled = false
    }
	node_config {
		instance_type_ids = ["ecs.g1ie.xlarge"]
        subnet_ids = ["${volcengine_subnet.foo.id}"]
		image_id = [for image in data.volcengine_images.foo.images : image.image_id if image.image_name == "veLinux 1.0 CentOS兼容版 64位"][0]
		system_volume {
			type = "ESSD_PL0"
            size = "50"
		}
        data_volumes {
            type = "ESSD_PL0"
            size = "50"
			mount_point = "/tf"
        }
		initialize_script = "ZWNobyBoZWxsbyB0ZXJyYWZvcm0h"
		security {
            login {
                 password = "UHdkMTIzNDU2"
            }
			security_strategies = ["Hids"]
            security_group_ids = ["${volcengine_security_group.foo.id}"]
        }
		additional_container_storage_enabled = true
        instance_charge_type = "PostPaid"
        name_prefix = "acc-test"
        ecs_tags {
            key = "ecs_k1"
            value = "ecs_v1"
        }
	}
	kubernetes_config {
        labels {
            key   = "label1"
            value = "value1"
        }
        taints {
            key   = "taint-key/node-type"
            value = "taint-value"
			effect = "NoSchedule"
        }
        cordon = true
    }
	tags {
        key = "node-pool-k1"
        value = "node-pool-v1"
    }
}

resource "volcengine_ecs_instance" "foo" {
 	instance_name = "acc-test-ecs"
	host_name = "tf-acc-test"
  	image_id = [for image in data.volcengine_images.foo.images : image.image_id if image.image_name == "veLinux 1.0 CentOS兼容版 64位"][0]
  	instance_type = "ecs.g1ie.xlarge"
  	password = "93f0cb0614Aab12"
  	instance_charge_type = "PostPaid"
  	system_volume_type = "ESSD_PL0"
  	system_volume_size = 50
	data_volumes {
    	volume_type = "ESSD_PL0"
    	size = 50
    	delete_with_instance = true
  	}
	subnet_id = "${volcengine_subnet.foo.id}"
	security_group_ids = ["${volcengine_security_group.foo.id}"]
	project_name = "default"
	tags {
    	key = "k1"
    	value = "v1"
  	}
	lifecycle {
    	ignore_changes = [security_group_ids, tags, instance_name]
  	}
}

resource "volcengine_vke_node" "foo" {
    cluster_id = "${volcengine_vke_cluster.foo.id}"
    instance_id = "${volcengine_ecs_instance.foo.id}"
	node_pool_id = "${volcengine_vke_node_pool.foo.id}"
}
`

func TestAccVolcengineVkeNodeResource_Basic(t *testing.T) {
	resourceName := "volcengine_vke_node.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return node.NewVolcengineVkeNodeService(client)
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
				Config: testAccVolcengineVkeNodeCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "additional_container_storage_enabled", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "container_storage_path", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "initialize_script", "ZWNobyBoZWxsbyB0ZXJyYWZvcm0h"),
					resource.TestCheckResourceAttr(acc.ResourceId, "keep_instance_name", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.labels.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.labels.0.key", "label1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.labels.0.value", "value1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.taints.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.taints.0.key", "taint-key/node-type"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.taints.0.value", "taint-value"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.taints.0.effect", "NoSchedule"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.cordon", "true"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "cluster_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "instance_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "image_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "node_pool_id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"keep_instance_name"},
			},
		},
	})
}

const testAccVolcengineVkeNodeCreateDefaultNodePoolConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
	vpc_name   = "acc-test-vpc"
  	cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  	subnet_name = "acc-test-subnet"
  	cidr_block = "172.16.0.0/24"
  	zone_id = "${data.volcengine_zones.foo.zones[0].id}"
	vpc_id = "${volcengine_vpc.foo.id}"
}

resource "volcengine_security_group" "foo" {
  	security_group_name = "acc-test-security-group"
  	vpc_id = "${volcengine_vpc.foo.id}"
}

data "volcengine_images" "foo" {
  name_regex = "veLinux 1.0 CentOS兼容版 64位"
}

resource "volcengine_vke_cluster" "foo" {
    name = "acc-test-cluster"
    description = "created by terraform"
    delete_protection_enabled = false
    cluster_config {
        subnet_ids = ["${volcengine_subnet.foo.id}"]
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
            subnet_ids = ["${volcengine_subnet.foo.id}"]
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

resource "volcengine_ecs_instance" "foo" {
 	instance_name = "acc-test-ecs"
	host_name = "tf-acc-test"
  	image_id = [for image in data.volcengine_images.foo.images : image.image_id if image.image_name == "veLinux 1.0 CentOS兼容版 64位"][0]
  	instance_type = "ecs.g1ie.xlarge"
  	password = "93f0cb0614Aab12"
  	instance_charge_type = "PostPaid"
  	system_volume_type = "ESSD_PL0"
  	system_volume_size = 50
	data_volumes {
    	volume_type = "ESSD_PL0"
    	size = 50
    	delete_with_instance = true
  	}
	subnet_id = "${volcengine_subnet.foo.id}"
	security_group_ids = ["${volcengine_security_group.foo.id}"]
	project_name = "default"
	tags {
    	key = "k1"
    	value = "v1"
  	}
	lifecycle {
    	ignore_changes = [security_group_ids, tags]
  	}
}

resource "volcengine_vke_default_node_pool" "foo" {
  	cluster_id = "${volcengine_vke_cluster.foo.id}"
  	node_config {
		security {
      		login {
        		password = "amw4WTdVcTRJVVFsUXpVTw=="
      		}
		security_group_ids = ["${volcengine_security_group.foo.id}"]
      	security_strategies = ["Hids"]
    	}
    	initialize_script = "ZWNobyBoZWxsbyB0ZXJyYWZvcm0h"
  	}
	kubernetes_config {
        labels {
            key   = "label1"
            value = "value1"
        }
        taints {
            key   = "taint-key/node-type"
            value = "taint-value"
			effect = "NoSchedule"
        }
        cordon = true
    }
	tags {
        key = "node-pool-k1"
        value = "node-pool-v1"
    }
}

resource "volcengine_vke_node" "foo" {
    cluster_id = "${volcengine_vke_cluster.foo.id}"
    instance_id = "${volcengine_ecs_instance.foo.id}"
	image_id = [for image in data.volcengine_images.foo.images : image.image_id if image.image_name == "veLinux 1.0 CentOS兼容版 64位"][0]
	initialize_script = "ZWNobyBoZWxsbyB2a2Uh"
	keep_instance_name = true
	additional_container_storage_enabled = false
	kubernetes_config {
        labels {
            key   = "label1"
            value = "value3"
        }
		labels {
            key   = "label2"
            value = "value2"
        }
        taints {
            key   = "taint-key/node-type"
            value = "taint-value-3"
			effect = "PreferNoSchedule"
        }
		taints {
            key   = "taint-key/node-type-2"
            value = "taint-value-2"
			effect = "PreferNoSchedule"
        }
        cordon = true
    }
	depends_on = ["volcengine_vke_default_node_pool.foo"]
}
`

func TestAccVolcengineVkeNodeResource_DefaultNodePool(t *testing.T) {
	resourceName := "volcengine_vke_node.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return node.NewVolcengineVkeNodeService(client)
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
				Config: testAccVolcengineVkeNodeCreateDefaultNodePoolConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "additional_container_storage_enabled", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "container_storage_path", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "initialize_script", "ZWNobyBoZWxsbyB2a2Uh"),
					resource.TestCheckResourceAttr(acc.ResourceId, "keep_instance_name", "true"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.labels.#", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.labels.0.key", "label1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.labels.0.value", "value3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.labels.1.key", "label2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.labels.1.value", "value2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.taints.#", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.taints.0.key", "taint-key/node-type"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.taints.0.value", "taint-value-3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.taints.0.effect", "PreferNoSchedule"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.taints.1.key", "taint-key/node-type-2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.taints.1.value", "taint-value-2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.taints.1.effect", "PreferNoSchedule"),
					resource.TestCheckResourceAttr(acc.ResourceId, "kubernetes_config.0.cordon", "true"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "cluster_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "instance_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "image_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "node_pool_id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"keep_instance_name"},
			},
		},
	})
}
