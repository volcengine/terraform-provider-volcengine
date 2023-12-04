---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_alarm_notify_group"
sidebar_current: "docs-volcengine-resource-tls_alarm_notify_group"
description: |-
  Provides a resource to manage tls alarm notify group
---
# volcengine_tls_alarm_notify_group
Provides a resource to manage tls alarm notify group
## Example Usage
```hcl
resource "volcengine_tls_alarm_notify_group" "foo" {
  iam_project_name        = "yyy"
  alarm_notify_group_name = "tf-test"
  notify_type             = ["Trigger"]
  receivers {
    receiver_type     = "User"
    receiver_names    = ["vke-qs"]
    receiver_channels = ["Email", "Sms"]
    start_time        = "23:00:00"
    end_time          = "23:59:59"
  }
}
```
## Argument Reference
The following arguments are supported:
* `alarm_notify_group_name` - (Required) The name of the notify group.
* `notify_type` - (Required) The notify type.
Trigger: Alarm Trigger
Recovery: Alarm Recovery.
* `receivers` - (Required) List of IAM users to receive alerts.
* `iam_project_name` - (Optional) The name of the iam project.

The `receivers` object supports the following:

* `end_time` - (Required) The end time.
* `receiver_channels` - (Required) The list of the receiver channels. Currently supported channels: Email, Sms, Phone.
* `receiver_names` - (Required) List of the receiver names.
* `receiver_type` - (Required) The receiver type, Can be set as: `User`(The id of user).
* `start_time` - (Required) The start time.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `alarm_notify_group_id` - The alarm notification group id.


## Import
tls alarm notify group can be imported using the id, e.g.
```
$ terraform import volcengine_tls_alarm_notify_group.default fa************
```

