---
subcategory: "CLOUD_MONITOR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_monitor_contact_group"
sidebar_current: "docs-volcengine-resource-cloud_monitor_contact_group"
description: |-
  Provides a resource to manage cloud monitor contact group
---
# volcengine_cloud_monitor_contact_group
Provides a resource to manage cloud monitor contact group
## Example Usage
```hcl
resource "volcengine_cloud_monitor_contact_group" "foo" {
  name             = "tfgroup"
  description      = "tftest"
  contacts_id_list = ["1737376113733353472", "1737375997680111616"]
}
```
## Argument Reference
The following arguments are supported:
* `name` - (Required) The name of the contact group.
* `contacts_id_list` - (Optional) When creating a contact group, contacts should be added with their contact ID. The maximum number of IDs allowed is 10, meaning that the maximum number of members in a single contact group is 10.
* `description` - (Optional) The description of the contact group.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
CloudMonitorContactGroup can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_monitor_contact_group.default resource_id
```

