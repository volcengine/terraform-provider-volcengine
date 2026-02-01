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
  enable          = "false"
  from_time       = 1750649545
  name            = "tf-test-etl-task"
  script          = ""
  source_topic_id = "9b756385-1dfb-4306-a094-0c88e04b34a5"
  to_time         = 1750735958
  target_resources {
    alias    = "tf-test-1"
    topic_id = "a690a9b8-72c1-40a3-b8c6-f89a81d3748e"
  }
  target_resources {
    alias    = "tf-test-2"
    topic_id = "bdf4f23b-a889-456c-ac5f-09d727427557"
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

