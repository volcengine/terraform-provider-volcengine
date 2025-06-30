---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_etl_task"
sidebar_current: "docs-volcengine-resource-tls_etl_task"
description: |-
  Provides a resource to manage tls etl task
---
# volcengine_tls_etl_task
Provides a resource to manage tls etl task
## Example Usage
```hcl
resource "volcengine_tls_etl_task" "foo" {
  dsl_type        = "NORMAL"
  description     = "for-tf-test"
  enable          = "true"
  from_time       = 1750649545
  name            = "tf-test-etl-task-1"
  script          = ""
  source_topic_id = "8ba48bd7-2493-4300-b1d0-cb7xxxxxxx"
  to_time         = 1750735958
  target_resources {
    alias    = "tf-test-1"
    topic_id = "b966e41a-d6a6-4999-bd75-39962xxxxxx"
  }
  target_resources {
    alias    = "tf-test-2"
    topic_id = "0ed72ac8-9531-4967-b216-ac3xxxxx"
  }
  task_type = "Resident"

}
```
## Argument Reference
The following arguments are supported:
* `dsl_type` - (Required, ForceNew) DSL type, fixed as NORMAL. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `enable` - (Required, ForceNew) Whether to enable the data processing task.
* `name` - (Required) The name of the processing task.
* `script` - (Required) Processing rules.
* `source_topic_id` - (Required, ForceNew) The log topic where the log to be processed is located.
* `target_resources` - (Required) Output the relevant information of the target.
* `task_type` - (Required, ForceNew) The task type is fixed as Resident.
* `description` - (Optional) A simple description of the data processing task.
* `from_time` - (Optional, ForceNew) The start time of the data to be processed.
* `to_time` - (Optional, ForceNew) The end time of the data to be processed.

The `target_resources` object supports the following:

* `alias` - (Required) Customize the name of the output target, which needs to be used to refer to the output target in the data processing rules.
* `topic_id` - (Required) Log topics used for storing processed logs.
* `role_trn` - (Optional, ForceNew) Cross-account authorized character names.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
EtlTask can be imported using the id, e.g.
```
$ terraform import volcengine_etl_task.default resource_id
```

