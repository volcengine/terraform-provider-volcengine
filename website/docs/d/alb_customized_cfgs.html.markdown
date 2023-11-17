---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_customized_cfgs"
sidebar_current: "docs-volcengine-datasource-alb_customized_cfgs"
description: |-
  Use this data source to query detailed information of alb customized cfgs
---
# volcengine_alb_customized_cfgs
Use this data source to query detailed information of alb customized cfgs
## Example Usage
```hcl
data "volcengine_alb_customized_cfgs" "foo" {}
```
## Argument Reference
The following arguments are supported:
* `customized_cfg_name` - (Optional) The name of the CustomizedCfg.
* `ids` - (Optional) A list of CustomizedCfg IDs.
* `listener_id` - (Optional) The id of the listener.
* `name_regex` - (Optional) A Name Regex of CustomizedCfg.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of the CustomizedCfg.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `cfgs` - The collection of CustomizedCfg query.
    * `create_time` - The create time of CustomizedCfg.
    * `customized_cfg_content` - The content of CustomizedCfg.
    * `customized_cfg_id` - The ID of CustomizedCfg.
    * `customized_cfg_name` - The name of CustomizedCfg.
    * `description` - The description of CustomizedCfg.
    * `id` - The ID of CustomizedCfg.
    * `listeners` - The listeners of CustomizedCfg.
        * `listener_id` - The ID of Listener.
        * `listener_name` - The Name of Listener.
        * `port` - The port info of listener.
        * `protocol` - The protocol info of listener.
    * `project_name` - The project name of CustomizedCfg.
    * `status` - The status of CustomizedCfg.
    * `update_time` - The update time of CustomizedCfg.
* `total_count` - The total count of CustomizedCfg query.


