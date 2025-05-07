---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_kubeconfig"
sidebar_current: "docs-volcengine-resource-veecp_kubeconfig"
description: |-
  Provides a resource to manage veecp kubeconfig
---
# volcengine_veecp_kubeconfig
Provides a resource to manage veecp kubeconfig
## Example Usage
```hcl
resource "volcengine_veecp_kubeconfig" "foo" {
  cluster_id = ""
  type       = ""
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The cluster id of the Kubeconfig.
* `type` - (Required, ForceNew) The type of the Kubeconfig, the value of type should be Public or Private.
* `valid_duration` - (Optional, ForceNew) The ValidDuration of the Kubeconfig, the range of the ValidDuration is 1 hour to 43800 hour.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VeecpKubeconfig can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_kubeconfig.default resource_id
```

