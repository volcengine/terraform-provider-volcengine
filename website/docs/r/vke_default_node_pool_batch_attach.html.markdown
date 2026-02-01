---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_default_node_pool_batch_attach"
sidebar_current: "docs-volcengine-resource-vke_default_node_pool_batch_attach"
description: |-
  Provides a resource to manage vke default node pool batch attach
---
# volcengine_vke_default_node_pool_batch_attach
Provides a resource to manage vke default node pool batch attach
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-project1"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-test-2"
  cidr_block  = "172.16.0.0/24"
  zone_id     = "cn-beijing-a"
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo" {
  vpc_id              = volcengine_vpc.foo.id
  security_group_name = "acc-test-security-group2"
}

resource "volcengine_ecs_instance" "foo" {
  image_id             = "image-ybqi99s7yq8rx7mnk44b"
  instance_type        = "ecs.g1ie.large"
  instance_name        = "acc-test-ecs-name2"
  password             = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type   = "ESSD_PL0"
  system_volume_size   = 40
  subnet_id            = volcengine_subnet.foo.id
  security_group_ids   = [volcengine_security_group.foo.id]
  lifecycle {
    ignore_changes = [security_group_ids, instance_name]
  }
}

resource "volcengine_ecs_instance" "foo2" {
  image_id             = "image-ybqi99s7yq8rx7mnk44b"
  instance_type        = "ecs.g1ie.large"
  instance_name        = "acc-test-ecs-name2"
  password             = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type   = "ESSD_PL0"
  system_volume_size   = 40
  subnet_id            = volcengine_subnet.foo.id
  security_group_ids   = [volcengine_security_group.foo.id]
  lifecycle {
    ignore_changes = [security_group_ids, instance_name]
  }
}

resource "volcengine_vke_cluster" "foo" {
  name                      = "acc-test-1"
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

resource "volcengine_vke_default_node_pool" "foo" {
  cluster_id = volcengine_vke_cluster.foo.id
  node_config {
    security {
      login {
        password = "amw4WTdVcTRJVVFsUXpVTw=="
      }
      security_group_ids  = [volcengine_security_group.foo.id]
      security_strategies = ["Hids"]
    }
    initialize_script = "ISMvYmluL2Jhc2gKZWNobyAx"

  }
  kubernetes_config {
    labels {
      key   = "tf-key1"
      value = "tf-value1"
    }
    labels {
      key   = "tf-key2"
      value = "tf-value2"
    }
    taints {
      key    = "tf-key3"
      value  = "tf-value3"
      effect = "NoSchedule"
    }
    taints {
      key    = "tf-key4"
      value  = "tf-value4"
      effect = "NoSchedule"
    }
    cordon = true
  }
  tags {
    key   = "tf-k1"
    value = "tf-v1"
  }
}

resource "volcengine_vke_default_node_pool_batch_attach" "foo" {
  cluster_id           = volcengine_vke_cluster.foo.id
  default_node_pool_id = volcengine_vke_default_node_pool.foo.id
  instances {
    instance_id                          = volcengine_ecs_instance.foo.id
    keep_instance_name                   = true
    additional_container_storage_enabled = false
  }
  instances {
    instance_id                          = volcengine_ecs_instance.foo2.id
    keep_instance_name                   = true
    additional_container_storage_enabled = false
  }
  kubernetes_config {
    labels {
      key   = "tf-key1"
      value = "tf-value1"
    }
    labels {
      key   = "tf-key2"
      value = "tf-value2"
    }
    taints {
      key    = "tf-key3"
      value  = "tf-value3"
      effect = "NoSchedule"
    }
    taints {
      key    = "tf-key4"
      value  = "tf-value4"
      effect = "NoSchedule"
    }
    cordon = true
  }
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The ClusterId of NodePool.
* `default_node_pool_id` - (Required, ForceNew) The default NodePool ID.
* `instances` - (Optional) The ECS InstanceIds add to NodePool.
* `kubernetes_config` - (Optional, ForceNew) The KubernetesConfig of NodeConfig. Please note that this field is the configuration of the node. The same key is subject to the config of the node pool. Different keys take effect together.

The `instances` object supports the following:

* `instance_id` - (Required) The instance id.
* `additional_container_storage_enabled` - (Optional) The flag of additional container storage enable, the value is `true` or `false`..Default is `false`.
* `container_storage_path` - (Optional) The container storage path.When additional_container_storage_enabled is `false` will ignore.
* `image_id` - (Optional) The Image Id to the ECS Instance.
* `keep_instance_name` - (Optional) The flag of keep instance name, the value is `true` or `false`.Default is `false`.

The `kubernetes_config` object supports the following:

* `cordon` - (Optional, ForceNew) The Cordon of KubernetesConfig.
* `labels` - (Optional, ForceNew) The Labels of KubernetesConfig.
* `taints` - (Optional, ForceNew) The Taints of KubernetesConfig.

The `labels` object supports the following:

* `key` - (Required, ForceNew) The Key of Labels.
* `value` - (Optional, ForceNew) The Value of Labels.

The `taints` object supports the following:

* `key` - (Required, ForceNew) The Key of Taints.
* `effect` - (Optional, ForceNew) The Effect of Taints. The value can be one of the following: `NoSchedule`, `NoExecute`, `PreferNoSchedule`, default value is `NoSchedule`.
* `value` - (Optional, ForceNew) The Value of Taints.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `is_import` - Is import of the DefaultNodePool. It only works when imported, set to true.
* `node_config` - The Config of NodePool.
    * `ecs_tags` - Tags for Ecs.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `initialize_script` - The initializeScript of NodeConfig.
    * `name_prefix` - The NamePrefix of NodeConfig.
    * `pre_script` - The PreScript of NodeConfig.
    * `security` - The Security of NodeConfig.
        * `login` - The Login of Security.
            * `password` - The Password of Security.
            * `ssh_key_pair_name` - The SshKeyPairName of Security.
        * `security_group_ids` - The SecurityGroupIds of Security.
        * `security_strategies` - The SecurityStrategies of Security.
* `tags` - Tags.
    * `key` - The Key of Tags.
    * `value` - The Value of Tags.


