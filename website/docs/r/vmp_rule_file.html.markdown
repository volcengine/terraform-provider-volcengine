---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_rule_file"
sidebar_current: "docs-volcengine-resource-vmp_rule_file"
description: |-
  Provides a resource to manage vmp rule file
---
# volcengine_vmp_rule_file
Provides a resource to manage vmp rule file
## Example Usage
```hcl
resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-1"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-1"
  username                  = "admin123"
  password                  = "**********"
}

resource "volcengine_vmp_rule_file" "foo" {
  name         = "acc-test-1"
  workspace_id = volcengine_vmp_workspace.foo.id
  description  = "acc-test-1"
  content      = <<EOF
groups:
    - interval: 10s
      name: recording_rules
      rules:
        - expr: sum(irate(container_cpu_usage_seconds_total{image!=""}[5m])) by (pod) *100
          labels:
            team: operations
          record: pod:cpu:useage
EOF
}
```
## Argument Reference
The following arguments are supported:
* `content` - (Required) The content of the rule file.
* `name` - (Required, ForceNew) The name of the rule file.
* `workspace_id` - (Required, ForceNew) The id of the workspace.
* `description` - (Optional) The description of the rule file.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of workspace.
* `last_update_time` - The last update time of rule file.
* `rule_count` - The rule count number of rule file.
* `rule_file_id` - The id of rule file.
* `status` - The status of workspace.


## Import
VMP Rule File can be imported using the workspace_id:rule_file_id, e.g.
(We can only get rule file by WorkspaceId and RuleFileId)
```
$ terraform import volcengine_vmp_rule_file.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6:d6f72bd9-674e-4651-b98c-3797657d9614
```

