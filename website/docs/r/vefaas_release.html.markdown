---
subcategory: "VEFAAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_vefaas_release"
sidebar_current: "docs-volcengine-resource-vefaas_release"
description: |-
  Provides a resource to manage vefaas release
---
# volcengine_vefaas_release
Provides a resource to manage vefaas release
## Example Usage
```hcl
resource "volcengine_vefaas_release" "foo" {
  function_id           = "9p5emxxxx"
  revision_number       = 0
  target_traffic_weight = 30
  lifecycle {
    ignore_changes = [revision_number]
  }
}
```
## Argument Reference
The following arguments are supported:
* `function_id` - (Required, ForceNew) The ID of Function.
* `revision_number` - (Required, ForceNew) When the RevisionNumber to be released is 0, the Latest code (Latest) will be released and a new version will be created. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `description` - (Optional) The description of released this time.
* `max_instance` - (Optional, ForceNew) Upper limit of the number of function instances.
* `rolling_step` - (Optional, ForceNew) Percentage of grayscale step size.
* `target_traffic_weight` - (Optional) Target percentage of published traffic.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `current_traffic_weight` - The current percentage of current published traffic.
* `error_code` - Error code when the release fails.
* `failed_instance_logs` - Download link for the failed instance log.
* `new_revision_number` - The version number of the newly released version.
* `old_revision_number` - The version number of the old version.
* `release_record_id` - The ID of Release record.
* `stable_revision_number` - The current version number that is stably running online.
* `start_time` - The current release start time.
* `status_message` - Detailed information of the function release status.
* `status` - The status of function release.


## Import
VefaasRelease can be imported using the id, e.g.
```
$ terraform import volcengine_vefaas_release.default FunctionId:ReleaseRecordId
```

