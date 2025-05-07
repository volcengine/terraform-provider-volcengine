---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_edge_node_pool"
sidebar_current: "docs-volcengine-resource-veecp_edge_node_pool"
description: |-
  Provides a resource to manage veecp edge node pool
---
# volcengine_veecp_edge_node_pool
Provides a resource to manage veecp edge node pool
## Example Usage
```hcl
#data "volcengine_zones" "foo"{
#}
#
#resource "volcengine_vpc" "foo" {
#    vpc_name = "acc-test-project1"
#    cidr_block = "172.16.0.0/16"
#}
#
#resource "volcengine_subnet" "foo" {
#    subnet_name = "acc-subnet-test-2"
#    cidr_block = "172.16.0.0/24"
#    zone_id = data.volcengine_zones.foo.zones[0].id
#    vpc_id = volcengine_vpc.foo.id
#}
#
#resource "volcengine_security_group" "foo" {
#    vpc_id = volcengine_vpc.foo.id
#    security_group_name = "acc-test-security-group2"
#}
#
#resource "volcengine_veecp_cluster" "foo" {
#    name = "acc-test-1"
#    description = "created by terraform"
#    delete_protection_enabled = false
#    profile = "Edge"
#    cluster_config {
#        subnet_ids = [volcengine_subnet.foo.id]
#        api_server_public_access_enabled = true
#        api_server_public_access_config {
#            public_access_network_config {
#                billing_type = "PostPaidByBandwidth"
#                bandwidth = 1
#            }
#        }
#        resource_public_access_default_enabled = true
#    }
#    pods_config {
#        pod_network_mode = "Flannel"
#        flannel_config {
#            pod_cidrs = ["172.22.224.0/20"]
#            max_pods_per_node = 64
#        }
#    }
#    services_config {
#        service_cidrsv4 = ["172.30.0.0/18"]
#    }
#}

resource "volcengine_veecp_edge_node_pool" "foo" {
  cluster_id     = "ccvmb0c66t101fnob3dhg"
  name           = "acc-test-tf"
  node_pool_type = "edge-machine-pool"
  vpc_id         = "vpc-l9sz9qlf2t"
  elastic_config {
    cloud_server_identity = "cloudserver-47vz7k929cp9xqb"
    auto_scale_config {
      enabled          = true
      max_replicas     = 2
      desired_replicas = 0
      min_replicas     = 0
      priority         = 10
    }
    instance_area {
      cluster_name = "bdcdn-zzcu02"
      vpc_identity = "vpc-l9sz9qlf2t"
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The ClusterId of NodePool.
* `billing_configs` - (Optional) The billing configuration of the node pool.
* `client_token` - (Optional) The ClientToken of NodePool.
* `elastic_config` - (Optional) Elastic scaling configuration.
* `kubernetes_config` - (Optional) The KubernetesConfig of NodeConfig.
* `name` - (Optional) The Name of NodePool.
* `node_pool_type` - (Optional, ForceNew) Node pool type, with the default being a static node pool. edge-machine-set: Static node pool. edge-machine-pool: Elastic node poolNode pool type, which is static node pool by default. edge-machine-set: static node pool
edge-machine-pool: elastic node pool.
* `vpc_id` - (Optional, ForceNew) The VpcId of NodePool.

The `auto_scale_config` object supports the following:

* `desired_replicas` - (Required) The DesiredReplicas of AutoScaling, default 0, range in min_replicas to max_replicas.
* `enabled` - (Required) Whether to enable the auto scaling function of the node pool. When a node needs to be manually added to the node pool, the value of this field must be `false`.
* `max_replicas` - (Required) The MaxReplicas of AutoScaling, default 10, range in 1~2000. This field is valid when the value of `enabled` is `true`.
* `min_replicas` - (Required) The MinReplicas of AutoScaling, default 0. This field is valid when the value of `enabled` is `true`.
* `priority` - (Required) The Priority of AutoScaling, default 10, rang in 0~100. This field is valid when the value of `enabled` is `true` and the value of `subnet_policy` is `Priority`.

The `billing_configs` object supports the following:

* `pre_paid_period_number` - (Required) Prepaid period number.
* `pre_paid_period` - (Required) The pre-paid period of the node pool, in months. The value range is 1-9. This parameter takes effect only when the billing_type is PrePaid.
* `auto_renew` - (Optional) Whether to automatically renew the node pool.

The `elastic_config` object supports the following:

* `cloud_server_identity` - (Required) The ID of the edge service corresponding to the elastic node. On the edge computing node's edge service page, obtain the edge service ID.
* `auto_scale_config` - (Optional) The node pool elastic scaling configuration information.
* `instance_area` - (Optional) 

The `kubernetes_config` object supports the following:

* `labels` - (Optional) The Labels of KubernetesConfig.
* `taints` - (Optional) The Taints of KubernetesConfig.

The `labels` object supports the following:

* `key` - (Optional) The Key of Labels.
* `value` - (Optional) The Value of Labels.

The `taints` object supports the following:

* `effect` - (Optional) The Effect of Taints, the value can be `NoSchedule` or `NoExecute` or `PreferNoSchedule`.
* `key` - (Optional) The Key of Taints.
* `value` - (Optional) The Value of Taints.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VeecpNodePool can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_node_pool.default resource_id
```

