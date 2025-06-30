---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_rule_files"
sidebar_current: "docs-volcengine-datasource-vmp_rule_files"
description: |-
  Use this data source to query detailed information of vmp rule files
---
# volcengine_vmp_rule_files
Use this data source to query detailed information of vmp rule files
## Example Usage
```hcl
resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-1"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-1"
  username                  = "admin123"
  password                  = "*********"
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

data "volcengine_vmp_rule_files" "foo" {
  ids          = [volcengine_vmp_rule_file.foo.rule_file_id]
  workspace_id = volcengine_vmp_workspace.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `workspace_id` - (Required) The id of workspace.
* `ids` - (Optional) A list of Rule File IDs.
* `name` - (Optional) The name of rule file.
* `output_file` - (Optional) File name where to save data source results.
* `status` - (Optional) The status of rule file.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `files` - The collection of query.
    * `content` - The content of rule file.
    * `create_time` - The create time of rule file.
    * `description` - The description of rule file.
    * `id` - The ID of rule file.
    * `last_update_time` - The last update time of rule file.
    * `name` - The name of rule file.
    * `rule_count` - The rule count number of rule file.
    * `status` - The status of rule file.
* `total_count` - The total count of query.


