---
subcategory: "CLOUD_MONITOR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_monitor_contact"
sidebar_current: "docs-volcengine-resource-cloud_monitor_contact"
description: |-
  Provides a resource to manage cloud monitor contact
---
# volcengine_cloud_monitor_contact
Provides a resource to manage cloud monitor contact
## Example Usage
```hcl
resource "volcengine_cloud_monitor_contact" "default" {
  name  = "tf-acc"
  email = "192*****72@****.com"
  phone = "180****27812"
}
```
## Argument Reference
The following arguments are supported:
* `email` - (Required) The email of contact.
* `name` - (Required) The name of contact.
* `phone` - (Optional) The phone of contact.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
CloudMonitor Contact can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_monitor_contact.default 145258255725730****
```

