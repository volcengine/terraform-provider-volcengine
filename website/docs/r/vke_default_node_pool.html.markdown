---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_default_node_pool"
sidebar_current: "docs-volcengine-resource-vke_default_node_pool"
description: |-
  Provides a resource to manage vke default node pool
---
# volcengine_vke_default_node_pool
Provides a resource to manage vke default node pool
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
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The ClusterId of NodePool.
* `kubernetes_config` - (Required) The KubernetesConfig of NodeConfig.
* `node_config` - (Required) The Config of NodePool.
* `instances` - (Optional) The ECS InstanceIds add to NodePool.
* `tags` - (Optional) Tags.

The `ecs_tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

The `instances` object supports the following:

* `instance_id` - (Required) The instance id.
* `additional_container_storage_enabled` - (Optional) The flag of additional container storage enable, the value is `true` or `false`..Default is `false`.
* `container_storage_path` - (Optional) The container storage path.When additional_container_storage_enabled is `false` will ignore.
* `image_id` - (Optional) The Image Id to the ECS Instance.
* `keep_instance_name` - (Optional) The flag of keep instance name, the value is `true` or `false`.Default is `false`.

The `kubernetes_config` object supports the following:

* `cordon` - (Required) The Cordon of KubernetesConfig.
* `labels` - (Optional) The Labels of KubernetesConfig.
* `name_prefix` - (Optional) The NamePrefix of node metadata.
* `taints` - (Optional) The Taints of KubernetesConfig.

The `labels` object supports the following:

* `key` - (Optional) The Key of Labels.
* `value` - (Optional) The Value of Labels.

The `login` object supports the following:

* `password` - (Optional) The Password of Security.
* `ssh_key_pair_name` - (Optional) The SshKeyPairName of Security.

The `node_config` object supports the following:

* `security` - (Required) The Security of NodeConfig.
* `ecs_tags` - (Optional) Tags for Ecs.
* `initialize_script` - (Optional) The initializeScript of NodeConfig.
* `name_prefix` - (Optional) The NamePrefix of NodeConfig.

The `security` object supports the following:

* `login` - (Required) The Login of Security.
* `security_group_ids` - (Optional) The SecurityGroupIds of Security.
* `security_strategies` - (Optional) The SecurityStrategies of Security.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

The `taints` object supports the following:

* `effect` - (Optional) The Effect of Taints.
* `key` - (Optional) The Key of Taints.
* `value` - (Optional) The Value of Taints.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `is_import` - Is import of the DefaultNodePool. It only works when imported, set to true.


## Import
VKE default node can be imported using the node id, e.g.
```
$ terraform import volcengine_vke_default_node.default nc5t5epmrsf****
```

