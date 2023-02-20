---
subcategory: "BIOOS"
layout: "volcengine"
page_title: "Volcengine: volcengine_bioos_cluster_bind"
sidebar_current: "docs-volcengine-resource-bioos_cluster_bind"
description: |-
  Provides a resource to manage bioos cluster bind
---
# volcengine_bioos_cluster_bind
Provides a resource to manage bioos cluster bind
## Example Usage
```hcl
resource "volcengine_bioos_cluster_bind" "example" {
  workspace_id = "wcfhp1vdeig48u8ufv8sg"
  cluster_id   = "ucfhp1nteig48u8ufv8s0" //必填
  type         = "workflow"              //必填, workflow 或 notebook
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The id of the cluster.
* `type` - (Required, ForceNew) The type of the cluster bind.
* `workspace_id` - (Required, ForceNew) The id of the workspace.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Cluster binder can be imported using the workspace id and cluster id, e.g.
```
$ terraform import volcengine_bioos_cluster_bind.default wc*****:uc***
```

