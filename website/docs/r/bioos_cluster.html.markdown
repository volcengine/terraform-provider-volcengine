---
subcategory: "BIOOS"
layout: "volcengine"
page_title: "Volcengine: volcengine_bioos_cluster"
sidebar_current: "docs-volcengine-resource-bioos_cluster"
description: |-
  Provides a resource to manage bioos cluster
---
# volcengine_bioos_cluster
Provides a resource to manage bioos cluster
## Example Usage
```hcl
resource "volcengine_bioos_cluster" "foo" {
  name        = "test-cluster"     //必填
  description = "test-description" //选填
  #  vke_config { //选填，和shared_config二者中必填一个
  #    cluster_id = "ccerdh8fqtofh16uf6q60" //也可替换成volcengine_vke_cluster.example.id
  #    storage_class = "ebs-ssd"
  #  }
  shared_config {
    enable = true
  }
}
```
## Argument Reference
The following arguments are supported:
* `name` - (Required, ForceNew) The name of the cluster.
* `description` - (Optional, ForceNew) The description of the cluster.
* `shared_config` - (Optional, ForceNew) The configuration of the shared cluster.
* `vke_config` - (Optional, ForceNew) The configuration of the vke cluster. This cluster type is not recommended. It is recommended to use a shared cluster.

The `shared_config` object supports the following:

* `enable` - (Required, ForceNew) Whether to enable a shared cluster. This value must be `true`.

The `vke_config` object supports the following:

* `cluster_id` - (Required, ForceNew) The id of the vke cluster.
* `storage_class` - (Required, ForceNew) The name of the StorageClass that the vke cluster has installed.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `cluster_id` - The id of the bioos cluster.


## Import
Cluster can be imported using the id, e.g.
```
$ terraform import volcengine_bioos_cluster.default *****
```

