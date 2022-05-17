---
subcategory: "EBS"
layout: "vestack"
page_title: "Vestack: vestack_volumes"
sidebar_current: "docs-vestack-datasource-volumes"
description: |-
  Use this data source to query detailed information of volumes
---
# vestack_volumes
Use this data source to query detailed information of volumes
## Example Usage
```hcl
data "vestack_volumes" "default" {
  ids = ["vol-3tzg6y5imn3b9fchkedb"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Volume IDs.
* `instance_id` - (Optional) The Id of instance.
* `kind` - (Optional) The Kind of Volume.
* `name_regex` - (Optional) A Name Regex of Volume.
* `output_file` - (Optional) File name where to save data source results.
* `volume_name` - (Optional) The name of Volume.
* `volume_status` - (Optional) The Status of Volume.
* `volume_type` - (Optional) The type of Volume.
* `zone_id` - (Optional) The Id of Zone.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of Volume query.
* `volumes` - The collection of Volume query.


