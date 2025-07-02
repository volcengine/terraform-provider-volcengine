---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_traffic_mirror_filter"
sidebar_current: "docs-volcengine-resource-traffic_mirror_filter"
description: |-
  Provides a resource to manage traffic mirror filter
---
# volcengine_traffic_mirror_filter
Provides a resource to manage traffic mirror filter
## Example Usage
```hcl
resource "volcengine_traffic_mirror_filter" "foo" {
  traffic_mirror_filter_name = "acc-test-traffic-mirror-filter"
  description                = "acc-test"
  project_name               = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `description` - (Optional) The description of the traffic mirror filter.
* `project_name` - (Optional) The project name of the traffic mirror filter.
* `tags` - (Optional) Tags.
* `traffic_mirror_filter_name` - (Optional) The name of the traffic mirror filter.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `created_at` - The create time of traffic mirror filter.
* `status` - The status of traffic mirror filter.
* `updated_at` - The last update time of traffic mirror filter.


## Import
TrafficMirrorFilter can be imported using the id, e.g.
```
$ terraform import volcengine_traffic_mirror_filter.default resource_id
```

