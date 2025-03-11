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
resource "volcengine_veecp_edge_node_pool" "foo" {
  cluster_id = ""
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The ClusterId of NodePool.
* `kubernetes_config` - (Required) The KubernetesConfig of NodeConfig.
* `billing_configs` - (Optional) The billing configuration of the node pool.
* `client_token` - (Optional) The ClientToken of NodePool.
* `elastic_config` - (Optional) Elastic scaling configuration. This field takes effect only when the node_pool_type is edge-machine-pool.
* `name` - (Optional) The Name of NodePool.
* `node_pool_type` - (Optional) Node pool type, with the default being a static node pool. edge - machine - set: Static node pool. edge - machine - pool: Elastic node poolNode pool type, which is static node pool by default. edge-machine-set: static node pool
edge-machine-pool: elastic node pool.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional, ForceNew) The VpcId of NodePool.

The `auto_scaling` object supports the following:

* `desired_replicas` - (Optional) The DesiredReplicas of AutoScaling, default 0, range in min_replicas to max_replicas.
* `enabled` - (Optional) Whether to enable the auto scaling function of the node pool. When a node needs to be manually added to the node pool, the value of this field must be `false`.
* `max_replicas` - (Optional) The MaxReplicas of AutoScaling, default 10, range in 1~2000. This field is valid when the value of `enabled` is `true`.
* `min_replicas` - (Optional) The MinReplicas of AutoScaling, default 0. This field is valid when the value of `enabled` is `true`.
* `priority` - (Optional) The Priority of AutoScaling, default 10, rang in 0~100. This field is valid when the value of `enabled` is `true` and the value of `subnet_policy` is `Priority`.

The `billing_configs` object supports the following:

* `pre_paid_period_number` - (Required) Prepaid period number.
* `pre_paid_period` - (Required) The pre-paid period of the node pool, in months. The value range is 1-9. This parameter takes effect only when the billing_type is PrePaid.
* `auto_renew` - (Optional) Whether to automatically renew the node pool.

The `elastic_config` object supports the following:

* `cloud_server_identity` - (Required) The ID of the edge service corresponding to the elastic node. On the edge computing node's edge service page, obtain the edge service ID.
* `auto_scaling` - (Optional) The node pool elastic scaling configuration information.
* `instance_area` - (Optional) 

The `kubernetes_config` object supports the following:

* `labels` - (Optional) The Labels of KubernetesConfig.
* `taints` - (Optional) The Taints of KubernetesConfig.

The `labels` object supports the following:

* `key` - (Optional) The Key of Labels.
* `value` - (Optional) The Value of Labels.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

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

