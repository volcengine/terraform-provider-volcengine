---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_nodes"
sidebar_current: "docs-volcengine-datasource-vke_nodes"
description: |-
  Use this data source to query detailed information of vke nodes
---
# volcengine_vke_nodes
Use this data source to query detailed information of vke nodes
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-security-group"
  vpc_id              = volcengine_vpc.foo.id
}

data "volcengine_images" "foo" {
  name_regex = "veLinux 1.0 CentOS兼容版 64位"
}

resource "volcengine_vke_cluster" "foo" {
  name                      = "acc-test-cluster"
  description               = "created by terraform"
  delete_protection_enabled = false
  cluster_config {
    subnet_ids                       = [volcengine_subnet.foo.id]
    api_server_public_access_enabled = true
    api_server_public_access_config {
      public_access_network_config {
        billing_type = "PostPaidByBandwidth"
        bandwidth    = 1
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
    key   = "tf-k1"
    value = "tf-v1"
  }
}

resource "volcengine_vke_node_pool" "foo" {
  cluster_id = volcengine_vke_cluster.foo.id
  name       = "acc-test-node-pool"
  auto_scaling {
    enabled = false
  }
  node_config {
    instance_type_ids = ["ecs.g1ie.xlarge"]
    subnet_ids        = [volcengine_subnet.foo.id]
    image_id          = [for image in data.volcengine_images.foo.images : image.image_id if image.image_name == "veLinux 1.0 CentOS兼容版 64位"][0]
    system_volume {
      type = "ESSD_PL0"
      size = "50"
    }
    data_volumes {
      type        = "ESSD_PL0"
      size        = "50"
      mount_point = "/tf"
    }
    initialize_script = "ZWNobyBoZWxsbyB0ZXJyYWZvcm0h"
    security {
      login {
        password = "UHdkMTIzNDU2"
      }
      security_strategies = ["Hids"]
      security_group_ids  = [volcengine_security_group.foo.id]
    }
    additional_container_storage_enabled = true
    instance_charge_type                 = "PostPaid"
    name_prefix                          = "acc-test"
    ecs_tags {
      key   = "ecs_k1"
      value = "ecs_v1"
    }
  }
  kubernetes_config {
    labels {
      key   = "label1"
      value = "value1"
    }
    taints {
      key    = "taint-key/node-type"
      value  = "taint-value"
      effect = "NoSchedule"
    }
    cordon = true
  }
  tags {
    key   = "node-pool-k1"
    value = "node-pool-v1"
  }
}

resource "volcengine_ecs_instance" "foo" {
  instance_name        = "acc-test-ecs-${count.index}"
  host_name            = "tf-acc-test"
  image_id             = [for image in data.volcengine_images.foo.images : image.image_id if image.image_name == "veLinux 1.0 CentOS兼容版 64位"][0]
  instance_type        = "ecs.g1ie.xlarge"
  password             = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type   = "ESSD_PL0"
  system_volume_size   = 50
  data_volumes {
    volume_type          = "ESSD_PL0"
    size                 = 50
    delete_with_instance = true
  }
  subnet_id          = volcengine_subnet.foo.id
  security_group_ids = [volcengine_security_group.foo.id]
  project_name       = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  lifecycle {
    ignore_changes = [security_group_ids, tags, instance_name]
  }
  count = 2
}

resource "volcengine_vke_node" "foo" {
  cluster_id   = volcengine_vke_cluster.foo.id
  instance_id  = volcengine_ecs_instance.foo[count.index].id
  node_pool_id = volcengine_vke_node_pool.foo.id
  count        = 2
}

data "volcengine_vke_nodes" "foo" {
  ids = volcengine_vke_node.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `cluster_ids` - (Optional) A list of Cluster IDs.
* `create_client_token` - (Optional) The Create Client Token.
* `ids` - (Optional) A list of Node IDs.
* `name_regex` - (Optional) A Name Regex of Node.
* `name` - (Optional) The Name of Node.
* `node_pool_ids` - (Optional) The Node Pool IDs.
* `output_file` - (Optional) File name where to save data source results.
* `statuses` - (Optional) The Status of filter.
* `zone_ids` - (Optional) The Zone IDs.

The `statuses` object supports the following:

* `conditions_type` - (Optional) The Type of Node Condition, the value is `Progressing` or `Ok` or `Unschedulable` or `InitilizeFailed` or `Unknown` or `NotReady` or `Security` or `Balance` or `ResourceCleanupFailed`.
* `phase` - (Optional) The Phase of Node, the value is `Creating` or `Running` or `Updating` or `Deleting` or `Failed` or `Starting` or `Stopping` or `Stopped`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `nodes` - The collection of Node query.
    * `additional_container_storage_enabled` - Is Additional Container storage enables.
    * `cluster_id` - The cluster id of node.
    * `condition_types` - The Condition of Node.
    * `container_storage_path` - The Storage Path.
    * `cordon` - The Cordon of KubernetesConfig.
    * `create_client_token` - The create client token of node.
    * `create_time` - The create time of Node.
    * `id` - The ID of Node.
    * `image_id` - The ImageId of NodeConfig.
    * `initialize_script` - The InitializeScript of NodeConfig.
    * `instance_id` - The instance id of node.
    * `is_virtual` - Is virtual node.
    * `labels` - The Label of KubernetesConfig.
        * `key` - The Key of KubernetesConfig.
        * `value` - The Value of KubernetesConfig.
    * `name` - The name of Node.
    * `node_pool_id` - The node pool id.
    * `phase` - The Phase of Node.
    * `pre_script` - The PreScript of NodeConfig.
    * `roles` - The roles of node.
    * `taints` - The Taint of KubernetesConfig.
        * `effect` - The Effect of Taint.
        * `key` - The Key of Taint.
        * `value` - The Value of Taint.
    * `update_time` - The update time of Node.
    * `zone_id` - The zone id.
* `total_count` - The total count of Node query.


