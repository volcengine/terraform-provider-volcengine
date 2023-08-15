---
subcategory: "EBS"
layout: "volcengine"
page_title: "Volcengine: volcengine_volumes"
sidebar_current: "docs-volcengine-datasource-volumes"
description: |-
  Use this data source to query detailed information of volumes
---
# volcengine_volumes
Use this data source to query detailed information of volumes
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_volume" "foo" {
  volume_name        = "acc-test-volume-${count.index}"
  volume_type        = "ESSD_PL0"
  description        = "acc-test"
  kind               = "data"
  size               = 60
  zone_id            = data.volcengine_zones.foo.zones[0].id
  volume_charge_type = "PostPaid"
  project_name       = "default"
  count              = 3
}

data "volcengine_volumes" "foo" {
  ids = volcengine_volume.foo[*].id
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
* `volume_status` - (Optional) The Status of Volume, the value can be `available` or `attaching` or `attached` or `detaching` or `creating` or `deleting` or `error` or `extending`.
* `volume_type` - (Optional) The type of Volume.
* `zone_id` - (Optional) The Id of Zone.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of Volume query.
* `volumes` - The collection of Volume query.


