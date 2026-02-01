---
subcategory: "VEENEDGE"
layout: "volcengine"
page_title: "Volcengine: volcengine_veenedge_instance"
sidebar_current: "docs-volcengine-resource-veenedge_instance"
description: |-
  Provides a resource to manage veenedge instance
---
# volcengine_veenedge_instance
Provides a resource to manage veenedge instance
## Example Usage
```hcl
resource "volcengine_veenedge_instance" "foo" {
  cloudserver_id = "cloudserver-x92*****jcc8f"
  area_name      = "*****"
  isp            = "CMCC"
}

resource "volcengine_veenedge_instance" "foo1" {
  instance_id = "veen0*****0111112"
}
```
## Argument Reference
The following arguments are supported:
* `area_name` - (Optional, ForceNew) The area name.
* `cloudserver_id` - (Optional, ForceNew) The id of cloud server.
* `cluster_name` - (Optional, ForceNew) The name of cluster.
* `default_isp` - (Optional, ForceNew) The default isp for multi line node.
* `instance_id` - (Optional, ForceNew) Import an exist instance, usually for import a default instance generated with cloud server creating.
* `isp` - (Optional, ForceNew) The isp info.
* `name` - (Optional) The name of instance, only effected in update scene.
* `secret_data` - (Optional) The data of secret, only effected in update scene.
* `secret_type` - (Optional) The type of secret, only effected in update scene. The value can be `KeyPair` or `Password`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Instance can be imported using the id, e.g.
```
$ terraform import volcengine_veenedge_instance.default veenn769ewmjjqyqh5dv
```

