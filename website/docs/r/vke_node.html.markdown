---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_node"
sidebar_current: "docs-volcengine-resource-vke_node"
description: |-
  Provides a resource to manage vke node
---
# volcengine_vke_node
Provides a resource to manage vke node
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
  instance_name        = "acc-test-ecs"
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
}

resource "volcengine_vke_node" "foo" {
  cluster_id   = volcengine_vke_cluster.foo.id
  instance_id  = volcengine_ecs_instance.foo.id
  node_pool_id = volcengine_vke_node_pool.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The cluster id.
* `instance_id` - (Required, ForceNew) The instance id.
* `additional_container_storage_enabled` - (Optional, ForceNew) The flag of additional container storage enable, the value is `true` or `false`. This field is valid only when adding an existing instance to the default node pool.
* `client_token` - (Optional, ForceNew) The client token.
* `container_storage_path` - (Optional, ForceNew) The container storage path. This field is valid only when adding an existing instance to the default node pool.
* `image_id` - (Optional, ForceNew) The ImageId of NodeConfig. This field is valid only when adding an existing instance to the default node pool.
* `initialize_script` - (Optional, ForceNew) The initializeScript of Node. This field is valid only when adding an existing instance to the default node pool.
* `keep_instance_name` - (Optional) The flag of keep instance name, the value is `true` or `false`.
* `kubernetes_config` - (Optional, ForceNew) The KubernetesConfig of Node. This field is valid only when adding an existing instance to the default node pool.
* `node_pool_id` - (Optional, ForceNew) The node pool id. This field is used to specify the custom node pool to which you want to add nodes. If not filled in, it means added to the default node pool.
* `pre_script` - (Optional, ForceNew) The PreScript of Node. This field is valid only when adding an existing instance to the default node pool.

The `kubernetes_config` object supports the following:

* `cordon` - (Optional, ForceNew) The Cordon of KubernetesConfig.
* `labels` - (Optional, ForceNew) The Labels of KubernetesConfig.
* `taints` - (Optional, ForceNew) The Taints of KubernetesConfig.

The `labels` object supports the following:

* `key` - (Optional, ForceNew) The Key of Labels.
* `value` - (Optional, ForceNew) The Value of Labels.

The `taints` object supports the following:

* `effect` - (Optional, ForceNew) The Effect of Taints, the value can be `NoSchedule` or `NoExecute` or `PreferNoSchedule`.
* `key` - (Optional, ForceNew) The Key of Taints.
* `value` - (Optional, ForceNew) The Value of Taints.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VKE node can be imported using the node id, e.g.
```
$ terraform import volcengine_vke_node.default nc5t5epmrsf****
```

