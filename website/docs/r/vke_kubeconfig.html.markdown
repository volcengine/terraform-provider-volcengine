---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_kubeconfig"
sidebar_current: "docs-volcengine-resource-vke_kubeconfig"
description: |-
  Provides a resource to manage vke kubeconfig
---
# volcengine_vke_kubeconfig
Provides a resource to manage vke kubeconfig
## Example Usage
```hcl
resource "volcengine_vke_kubeconfig" "foo" {
  cluster_id = "cce7hb97qtofmj1oi4udg"
  type       = "Private"
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
VkeKubeconfig can be imported using the id, e.g.
```
$ terraform import volcengine_vke_kubeconfig.default kce8simvqtofl0l6u4qd0
```

