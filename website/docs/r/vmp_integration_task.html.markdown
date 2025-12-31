---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_integration_task"
sidebar_current: "docs-volcengine-resource-vmp_integration_task"
description: |-
  Provides a resource to manage vmp integration task
---
# volcengine_vmp_integration_task
Provides a resource to manage vmp integration task
## Example Usage
```hcl
# 创建一个 VMP 集成任务
resource "volcengine_vmp_integration_task" "foo1" {
  name         = "tf_test_integration_task3"
  type         = "CloudMonitor"
  environment  = "Managed"
  workspace_id = "38046480-06a4-4d0e-bbd0-fed4eb358385"
  params = jsonencode({
    JobName        = "test1"
    ScrapeInterval = "1m"
    Regions = {
      cn-beijing = {
        VCM_ALB = {}
        VCM_ECS = {}
      }
    }
    MetaLabelWithNewMetric = false
  })
}

resource "volcengine_vmp_integration_task" "foo2" {
  name         = "tf_test_integration_task4"
  type         = "CloudMonitor"
  environment  = "Managed"
  workspace_id = "38046480-06a4-4d0e-bbd0-fed4eb358385"
  params = jsonencode({
    JobName        = "test1"
    ScrapeInterval = "1m"
    Regions = {
      cn-beijing = {
        VCM_ALB = {}
        VCM_ECS = {}
      }
    }
    MetaLabelWithNewMetric = false
  })
}


# 查询 VMP 集成任务列表
data "volcengine_vmp_integration_tasks" "foo" {
  ids = [volcengine_vmp_integration_task.foo1.id, volcengine_vmp_integration_task.foo2.id]
}

output "integration_task_id" {
  value = join(",", [volcengine_vmp_integration_task.foo1.id, volcengine_vmp_integration_task.foo2.id])
}

output "integration_task_name" {
  value = join(",", [volcengine_vmp_integration_task.foo1.name, volcengine_vmp_integration_task.foo2.name])
}
```
## Argument Reference
The following arguments are supported:
* `name` - (Required) The name of the integration task. Length: 1-40 characters. Supports Chinese, English, numbers, and underscores.
* `type` - (Required) The type of the integration task. For example, `CloudMonitor` indicates a cloud monitoring integration task.
* `environment` - (Optional) The deployment environment. Valid values: `Vke` or `Managed`.
* `params` - (Optional) The parameters of the integration task. Must be a JSON-escaped string.
* `vke_cluster_id` - (Optional) The ID of the VKE cluster. Required when Environment is `Vke`.
* `workspace_id` - (Optional) The workspace ID. Required when Environment is `Managed`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of the integration task.
* `status` - The status of the integration task. Valid values: `Creating`, `Updating`, `Active`, `Error`, `Deleting`.
* `update_time` - The update time of the integration task.


## Import
VMP Integration Task can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_integration_task.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
```

